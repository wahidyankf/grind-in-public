from datetime import UTC, datetime, timedelta

import pytest

from edutech_be import auth
from edutech_be.auth import (
    EmailTaken,
    InvalidCredentials,
    InvalidEmail,
    InvalidToken,
    WeakPassword,
)
from edutech_be.config import Settings, build_settings
from edutech_be.fake_store import FakeAuthStore
from edutech_be.passwords import verify_password
from edutech_be.tokens import hash_token

NOW = datetime(2026, 10, 15, 12, 0, tzinfo=UTC)
SETTINGS: Settings = build_settings({})
PASSWORD = "correct horse"


def _register_and_login(store: FakeAuthStore) -> auth.TokenPair:
    auth.register(store, "a@example.com", PASSWORD, NOW)
    return auth.login(store, SETTINGS, "a@example.com", PASSWORD, NOW)


def test_register_stores_normalized_email_and_hashed_password() -> None:
    store = FakeAuthStore()

    auth.register(store, "  Alice@Example.COM ", PASSWORD, NOW)

    (user,) = store.users.values()
    assert user.email == "alice@example.com"
    assert user.password_hash != PASSWORD
    assert verify_password(user.password_hash, PASSWORD)


def test_register_returns_the_normalized_email() -> None:
    email = auth.register(FakeAuthStore(), "  Alice@Example.COM ", PASSWORD, NOW)

    assert email == "alice@example.com"


def test_register_with_taken_email_raises_and_creates_no_second_user() -> None:
    store = FakeAuthStore()
    auth.register(store, "a@example.com", PASSWORD, NOW)

    with pytest.raises(EmailTaken):
        auth.register(store, "A@example.com", "another password", NOW)

    assert len(store.users) == 1


def test_register_rejects_malformed_email() -> None:
    with pytest.raises(InvalidEmail):
        auth.register(FakeAuthStore(), "not-an-email", PASSWORD, NOW)


def test_register_rejects_short_password() -> None:
    with pytest.raises(WeakPassword):
        auth.register(FakeAuthStore(), "a@example.com", "short", NOW)


def test_register_rejects_common_password() -> None:
    with pytest.raises(WeakPassword):
        auth.register(FakeAuthStore(), "a@example.com", "password", NOW)


def test_login_returns_access_and_refresh_tokens() -> None:
    store = FakeAuthStore()

    pair = _register_and_login(store)

    assert pair.access_token and pair.refresh_token
    assert pair.access_expires_in == SETTINGS.access_token_ttl_seconds


def test_login_stores_only_token_hashes() -> None:
    store = FakeAuthStore()

    pair = _register_and_login(store)

    assert pair.access_token not in store.tokens
    assert hash_token(pair.access_token) in store.tokens
    assert hash_token(pair.refresh_token) in store.tokens


@pytest.mark.parametrize(
    ("email", "password"),
    [("a@example.com", "wrong password"), ("nobody@example.com", PASSWORD)],
)
def test_login_failure_is_the_same_generic_error(email: str, password: str) -> None:
    store = FakeAuthStore()
    auth.register(store, "a@example.com", PASSWORD, NOW)

    with pytest.raises(InvalidCredentials):
        auth.login(store, SETTINGS, email, password, NOW)


def test_current_user_from_valid_access_token() -> None:
    store = FakeAuthStore()
    pair = _register_and_login(store)

    user = auth.get_current_user(store, pair.access_token, NOW)

    assert user.email == "a@example.com"


def test_refresh_token_cannot_be_used_as_access_token() -> None:
    store = FakeAuthStore()
    pair = _register_and_login(store)

    with pytest.raises(InvalidToken):
        auth.get_current_user(store, pair.refresh_token, NOW)


def test_access_token_expires_after_its_ttl() -> None:
    store = FakeAuthStore()
    pair = _register_and_login(store)
    later = NOW + timedelta(seconds=SETTINGS.access_token_ttl_seconds + 1)

    with pytest.raises(InvalidToken):
        auth.get_current_user(store, pair.access_token, later)


def test_refresh_returns_a_new_working_access_token() -> None:
    store = FakeAuthStore()
    pair = _register_and_login(store)

    new_access = auth.refresh(store, SETTINGS, pair.refresh_token, NOW)

    assert auth.get_current_user(store, new_access.token, NOW).email == "a@example.com"


def test_refresh_with_expired_token_fails() -> None:
    store = FakeAuthStore()
    pair = _register_and_login(store)
    later = NOW + timedelta(seconds=SETTINGS.refresh_token_ttl_seconds + 1)

    with pytest.raises(InvalidToken):
        auth.refresh(store, SETTINGS, pair.refresh_token, later)


def test_refresh_with_access_token_fails() -> None:
    store = FakeAuthStore()
    pair = _register_and_login(store)

    with pytest.raises(InvalidToken):
        auth.refresh(store, SETTINGS, pair.access_token, NOW)


def test_logout_revokes_both_tokens() -> None:
    store = FakeAuthStore()
    pair = _register_and_login(store)
    user = auth.get_current_user(store, pair.access_token, NOW)

    auth.logout(store, pair.access_token, pair.refresh_token, NOW)

    with pytest.raises(InvalidToken):
        auth.get_current_user(store, pair.access_token, NOW)
    with pytest.raises(InvalidToken):
        auth.refresh(store, SETTINGS, pair.refresh_token, NOW)
    assert user.id in store.users


def test_logout_with_another_users_refresh_token_fails() -> None:
    store = FakeAuthStore()
    alice = _register_and_login(store)
    auth.register(store, "b@example.com", "another password", NOW)
    bob = auth.login(store, SETTINGS, "b@example.com", "another password", NOW)

    with pytest.raises(InvalidToken):
        auth.logout(store, alice.access_token, bob.refresh_token, NOW)
