import re
from dataclasses import dataclass
from datetime import datetime, timedelta
from functools import cache
from uuid import UUID, uuid4

from edutech_be.config import Settings
from edutech_be.passwords import hash_password, is_acceptable_password, verify_password
from edutech_be.store import AuthStore, TokenKind, TokenRecord, User
from edutech_be.tokens import generate_token, hash_token

EMAIL_PATTERN = re.compile(r"[^@\s]+@[^@\s]+\.[^@\s]+")
MAX_EMAIL_LENGTH = 254


class AuthError(Exception):
    pass


class InvalidEmail(AuthError):
    pass


class WeakPassword(AuthError):
    pass


class EmailTaken(AuthError):
    pass


class InvalidCredentials(AuthError):
    # One error for unknown email and wrong password, so it does not leak which.
    pass


class InvalidToken(AuthError):
    pass


@dataclass(frozen=True)
class TokenPair:
    access_token: str
    refresh_token: str
    access_expires_in: int


@dataclass(frozen=True)
class AccessToken:
    token: str
    expires_in: int


def register(store: AuthStore, email: str, password: str, now: datetime) -> str:
    normalized = _normalize_email(email)
    if not is_acceptable_password(password):
        raise WeakPassword

    user = User(
        id=uuid4(),
        email=normalized,
        password_hash=hash_password(password),
        created_at=now,
    )
    if not store.insert_user(user):
        raise EmailTaken
    return normalized


def login(
    store: AuthStore, settings: Settings, email: str, password: str, now: datetime
) -> TokenPair:
    user = store.find_user_by_email(email.strip().lower())
    if user is None:
        # Hash anyway, so the response time does not reveal the email exists.
        verify_password(_dummy_hash(), password)
        raise InvalidCredentials
    if not verify_password(user.password_hash, password):
        raise InvalidCredentials

    return TokenPair(
        access_token=_issue(
            store, user.id, "access", settings.access_token_ttl_seconds, now
        ),
        refresh_token=_issue(
            store, user.id, "refresh", settings.refresh_token_ttl_seconds, now
        ),
        access_expires_in=settings.access_token_ttl_seconds,
    )


def refresh(
    store: AuthStore, settings: Settings, refresh_token: str, now: datetime
) -> AccessToken:
    record = _find_active_token(store, refresh_token, "refresh", now)
    if store.find_user_by_id(record.user_id) is None:
        raise InvalidToken

    token = _issue(
        store, record.user_id, "access", settings.access_token_ttl_seconds, now
    )
    return AccessToken(token=token, expires_in=settings.access_token_ttl_seconds)


def logout(
    store: AuthStore, access_token: str, refresh_token: str, now: datetime
) -> None:
    access = _find_active_token(store, access_token, "access", now)
    refresh_record = _find_active_token(store, refresh_token, "refresh", now)
    if access.user_id != refresh_record.user_id:
        raise InvalidToken

    store.revoke_token(access.id, now, access.user_id)
    store.revoke_token(refresh_record.id, now, access.user_id)


def get_current_user(store: AuthStore, access_token: str, now: datetime) -> User:
    record = _find_active_token(store, access_token, "access", now)
    user = store.find_user_by_id(record.user_id)
    if user is None:
        raise InvalidToken
    return user


def _normalize_email(email: str) -> str:
    normalized = email.strip().lower()
    if len(normalized) > MAX_EMAIL_LENGTH or not EMAIL_PATTERN.fullmatch(normalized):
        raise InvalidEmail
    return normalized


def _issue(
    store: AuthStore, user_id: UUID, kind: TokenKind, ttl_seconds: int, now: datetime
) -> str:
    token = generate_token()
    store.insert_token(
        TokenRecord(
            id=uuid4(),
            user_id=user_id,
            kind=kind,
            token_hash=hash_token(token),
            expires_at=now + timedelta(seconds=ttl_seconds),
            revoked_at=None,
        )
    )
    return token


def _find_active_token(
    store: AuthStore, raw_token: str, kind: TokenKind, now: datetime
) -> TokenRecord:
    record = store.find_token_by_hash(hash_token(raw_token))
    if (
        record is None
        or record.kind != kind
        or record.revoked_at is not None
        or record.expires_at <= now
    ):
        raise InvalidToken
    return record


@cache
def _dummy_hash() -> str:
    return hash_password("not-a-real-password")
