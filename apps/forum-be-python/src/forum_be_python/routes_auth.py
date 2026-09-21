from datetime import UTC, datetime, timedelta

import psycopg.errors
from fastapi import APIRouter, HTTPException

from forum_be_python.config import TOKEN_TTL_HOURS
from forum_be_python.db import one
from forum_be_python.deps import CurrentUser, PoolDep
from forum_be_python.passwords import hash_password, verify_password
from forum_be_python.schemas import Credentials, LoginRequest, TokenOut, UserOut
from forum_be_python.tokens import hash_token, new_token

router = APIRouter(prefix="/auth", tags=["auth"])

# Verified against when the username does not exist, so a miss costs as much as a wrong password
# and response time does not reveal which usernames are registered.
_UNKNOWN_USER_HASH = hash_password("synthetic-placeholder")


@router.post("/register", status_code=201)
def register(credentials: Credentials, pool: PoolDep) -> UserOut:
    try:
        with pool.connection() as conn:
            row = one(
                conn.execute(
                    "INSERT INTO users (username, password_hash) VALUES (%s, %s) RETURNING id",
                    (credentials.username, hash_password(credentials.password)),
                ).fetchone()
            )
    except psycopg.errors.UniqueViolation:
        raise HTTPException(status_code=409, detail="Username is already taken") from None
    return UserOut(id=row[0], username=credentials.username)


@router.post("/login")
def login(request: LoginRequest, pool: PoolDep) -> TokenOut:
    with pool.connection() as conn:
        user = conn.execute(
            "SELECT id, password_hash FROM users WHERE username = %s", (request.username,)
        ).fetchone()
        stored: str = user[1] if user else _UNKNOWN_USER_HASH
        password_matches = verify_password(request.password, stored)
        if user is None or not password_matches:
            raise HTTPException(status_code=401, detail="Invalid username or password")
        token = new_token()
        expires_at = datetime.now(UTC) + timedelta(hours=TOKEN_TTL_HOURS)
        conn.execute(
            "INSERT INTO auth_tokens (token_hash, user_id, expires_at) VALUES (%s, %s, %s)",
            (hash_token(token), user[0], expires_at),
        )
    return TokenOut(access_token=token, expires_at=expires_at)


@router.get("/me")
def me(user: CurrentUser) -> UserOut:
    return user
