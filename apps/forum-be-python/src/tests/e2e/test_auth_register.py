from typing import Any

import psycopg
import pytest
from fastapi.testclient import TestClient

_ALICE = {"username": "test-user-alice", "password": "synthetic-pass-1"}


def test_register_creates_a_user_without_exposing_the_password(client: TestClient) -> None:
    response = client.post("/auth/register", json=_ALICE)

    body: dict[str, Any] = response.json()
    assert response.status_code == 201
    assert body["username"] == "test-user-alice"
    assert isinstance(body["id"], int)
    assert "password" not in body
    assert "password_hash" not in body


def test_register_stores_only_a_hash_of_the_password(client: TestClient, database_url: str) -> None:
    client.post("/auth/register", json=_ALICE)

    with psycopg.connect(database_url) as conn:
        row = conn.execute("SELECT password_hash FROM users WHERE username = %s", (_ALICE["username"],)).fetchone()

    assert row is not None
    assert row[0].startswith("scrypt$")
    assert _ALICE["password"] not in row[0]


def test_register_rejects_a_username_that_is_already_taken(client: TestClient) -> None:
    client.post("/auth/register", json=_ALICE)

    response = client.post("/auth/register", json=_ALICE)

    assert response.status_code == 409


@pytest.mark.parametrize(
    "payload",
    [
        {"username": "ab", "password": "synthetic-pass-1"},
        {"username": "Test-User-Upper", "password": "synthetic-pass-1"},
        {"username": "test-user-bob", "password": "short"},
    ],
    ids=["username-too-short", "username-has-uppercase", "password-too-short"],
)
def test_register_rejects_invalid_input(client: TestClient, payload: dict[str, str]) -> None:
    response = client.post("/auth/register", json=payload)

    assert response.status_code == 422
