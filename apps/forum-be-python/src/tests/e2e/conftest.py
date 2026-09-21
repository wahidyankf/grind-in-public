import os
import uuid
from collections.abc import Callable, Iterator
from urllib.parse import urlsplit, urlunsplit

import psycopg
import pytest
from fastapi.testclient import TestClient
from psycopg import sql

from forum_be_python.main import create_app

_ADMIN_URL = os.environ.get("FORUM_TEST_ADMIN_URL", "postgresql://postgres@127.0.0.1:54329/postgres")
_LOOPBACK_HOSTS = {"127.0.0.1", "localhost", "::1"}


@pytest.fixture(scope="session")
def database_url() -> Iterator[str]:
    """Give the run its own database, and drop exactly that database afterwards.

    The name carries a per-run identifier, and a non-loopback admin host is refused, so a misconfigured
    environment fails instead of touching a database this run does not own.
    """
    parts = urlsplit(_ADMIN_URL)
    if parts.hostname not in _LOOPBACK_HOSTS:
        pytest.fail("refusing to create test databases on a non-loopback host")
    name = f"forum_test_{uuid.uuid4().hex[:12]}"
    with psycopg.connect(_ADMIN_URL, autocommit=True) as admin:
        admin.execute(sql.SQL("CREATE DATABASE {}").format(sql.Identifier(name)))
    try:
        yield urlunsplit(parts._replace(path=f"/{name}"))
    finally:
        with psycopg.connect(_ADMIN_URL, autocommit=True) as admin:
            admin.execute(sql.SQL("DROP DATABASE IF EXISTS {} WITH (FORCE)").format(sql.Identifier(name)))


@pytest.fixture(scope="session")
def client(database_url: str) -> Iterator[TestClient]:
    with TestClient(create_app(database_url)) as test_client:
        yield test_client


@pytest.fixture
def sign_in(client: TestClient) -> Callable[[str], dict[str, str]]:
    """Register and log in a synthetic user, returning the headers that authenticate their requests."""

    def _sign_in(username: str) -> dict[str, str]:
        credentials = {"username": username, "password": "synthetic-pass-1"}
        client.post("/auth/register", json=credentials)
        token: str = client.post("/auth/login", json=credentials).json()["access_token"]
        return {"Authorization": f"Bearer {token}"}

    return _sign_in


@pytest.fixture(autouse=True)
def _empty_tables(client: TestClient, database_url: str) -> None:  # pyright: ignore[reportUnusedFunction]
    """Start every test from empty tables so tests cannot depend on each other's rows."""
    with psycopg.connect(database_url, autocommit=True) as conn:
        tables = [row[0] for row in conn.execute("SELECT tablename FROM pg_tables WHERE schemaname = 'public'")]
        if tables:
            names = sql.SQL(", ").join(sql.Identifier(table) for table in tables)
            conn.execute(sql.SQL("TRUNCATE {} RESTART IDENTITY CASCADE").format(names))
