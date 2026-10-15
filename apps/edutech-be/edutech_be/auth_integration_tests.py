from collections.abc import Iterator
from datetime import UTC, datetime, timedelta
from uuid import uuid4

import pytest
from alembic import command
from fastapi.testclient import TestClient
from sqlalchemy import Connection, func, select, update
from sqlalchemy.exc import DBAPIError

from edutech_be import auth
from edutech_be.config import Settings, build_settings, settings
from edutech_be.db import create_db_engine
from edutech_be.main import app
from edutech_be.migrations import alembic_config
from edutech_be.routes import get_store
from edutech_be.sql_store import SqlAuthStore, auth_tokens, users
from edutech_be.tokens import hash_token

# Needs the Postgres container from `npx nx db:up edutech-be`.
pytestmark = pytest.mark.integration

PASSWORD = "correct horse"
SETTINGS: Settings = build_settings({})


@pytest.fixture
def connection() -> Iterator[Connection]:
    command.upgrade(alembic_config(), "head")
    engine = create_db_engine(settings)
    with engine.connect() as connection:
        # Everything here is rolled back, so tests do not leave rows behind.
        transaction = connection.begin()
        yield connection
        transaction.rollback()
    engine.dispose()


def _email() -> str:
    # Unique per test, so reruns never clash even if a rollback fails.
    return f"user-{uuid4().hex}@example.com"


def test_register_and_login_work_against_postgres(connection: Connection) -> None:
    store = SqlAuthStore(connection)
    now = datetime.now(UTC)
    email = _email()

    auth.register(store, email, PASSWORD, now)
    pair = auth.login(store, SETTINGS, email, PASSWORD, now)

    assert auth.get_current_user(store, pair.access_token, now).email == email


def test_duplicate_email_in_any_case_creates_one_row(connection: Connection) -> None:
    store = SqlAuthStore(connection)
    now = datetime.now(UTC)
    email = _email()

    auth.register(store, email, PASSWORD, now)
    with pytest.raises(auth.EmailTaken):
        auth.register(store, email.upper(), "another password", now)

    count = connection.execute(
        select(func.count()).select_from(users).where(users.c.email == email)
    ).scalar_one()
    assert count == 1


def test_database_rejects_delete_on_users(connection: Connection) -> None:
    store = SqlAuthStore(connection)
    email = _email()
    auth.register(store, email, PASSWORD, datetime.now(UTC))

    # A savepoint keeps the outer transaction usable after the trigger error.
    with pytest.raises(DBAPIError, match="hard delete"):
        with connection.begin_nested():
            connection.execute(users.delete())

    assert store.find_user_by_email(email) is not None


def test_database_rejects_delete_on_auth_tokens(connection: Connection) -> None:
    store = SqlAuthStore(connection)
    now = datetime.now(UTC)
    email = _email()
    auth.register(store, email, PASSWORD, now)
    auth.login(store, SETTINGS, email, PASSWORD, now)

    with pytest.raises(DBAPIError, match="hard delete"):
        with connection.begin_nested():
            connection.execute(auth_tokens.delete())


def test_trigger_sets_updated_at_on_update(connection: Connection) -> None:
    store = SqlAuthStore(connection)
    now = datetime.now(UTC)
    email = _email()
    auth.register(store, email, PASSWORD, now)
    pair = auth.login(store, SETTINGS, email, PASSWORD, now)
    token = store.find_token_by_hash(hash_token(pair.refresh_token))
    assert token is not None

    # Move updated_at back, then revoke. The trigger should set it to a fresh time.
    old = datetime(2000, 1, 1, tzinfo=UTC)
    connection.execute(
        update(auth_tokens).where(auth_tokens.c.id == token.id).values(updated_at=old)
    )
    store.revoke_token(token.id, now, token.user_id)

    updated_at = connection.execute(
        select(auth_tokens.c.updated_at).where(auth_tokens.c.id == token.id)
    ).scalar_one()
    assert updated_at > old


def test_http_flow_works_against_postgres(connection: Connection) -> None:
    app.dependency_overrides[get_store] = lambda: SqlAuthStore(connection)
    try:
        client = TestClient(app)
        email = _email()
        body = {"email": email, "password": PASSWORD}

        assert client.post("/auth/register", json=body).status_code == 201
        tokens = client.post("/auth/login", json=body).json()
        headers = {"Authorization": f"Bearer {tokens['access_token']}"}

        assert client.get("/auth/me", headers=headers).json()["email"] == email
        assert (
            client.post(
                "/auth/logout",
                headers=headers,
                json={"refresh_token": tokens["refresh_token"]},
            ).status_code
            == 204
        )
        assert client.get("/auth/me", headers=headers).status_code == 401
    finally:
        app.dependency_overrides.clear()


def test_expired_access_token_is_rejected_by_postgres(connection: Connection) -> None:
    store = SqlAuthStore(connection)
    now = datetime.now(UTC)
    email = _email()
    auth.register(store, email, PASSWORD, now)
    pair = auth.login(store, SETTINGS, email, PASSWORD, now)

    later = now + timedelta(seconds=SETTINGS.access_token_ttl_seconds + 1)
    with pytest.raises(auth.InvalidToken):
        auth.get_current_user(store, pair.access_token, later)
