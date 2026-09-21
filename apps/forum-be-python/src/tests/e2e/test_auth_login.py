from typing import Any

import psycopg
import pytest
from fastapi.testclient import TestClient

from forum_be_python.tokens import hash_token

_ALICE = {"username": "test-user-alice", "password": "synthetic-pass-1"}


def test_login_returns_a_bearer_token_for_valid_credentials(client: TestClient) -> None:
    client.post("/auth/register", json=_ALICE)

    response = client.post("/auth/login", json=_ALICE)

    body: dict[str, Any] = response.json()
    assert response.status_code == 200
    assert body["token_type"] == "bearer"
    assert len(body["access_token"]) == 43
    assert "expires_at" in body


def test_login_stores_only_a_hash_of_the_token(client: TestClient, database_url: str) -> None:
    client.post("/auth/register", json=_ALICE)
    token: str = client.post("/auth/login", json=_ALICE).json()["access_token"]

    with psycopg.connect(database_url) as conn:
        rows = conn.execute("SELECT token_hash FROM auth_tokens").fetchall()

    assert rows == [(hash_token(token),)]


@pytest.mark.parametrize(
    "attempt",
    [
        {"username": "test-user-alice", "password": "wrong-pass-123"},
        {"username": "test-user-nobody", "password": "synthetic-pass-1"},
    ],
    ids=["wrong-password", "unknown-user"],
)
def test_login_rejects_bad_credentials_without_saying_which_part_was_wrong(
    client: TestClient, attempt: dict[str, str]
) -> None:
    client.post("/auth/register", json=_ALICE)

    response = client.post("/auth/login", json=attempt)

    assert response.status_code == 401
    assert response.json() == {"detail": "Invalid username or password"}
