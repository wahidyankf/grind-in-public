from collections.abc import Callable
from typing import Any

import psycopg
import pytest
from fastapi.testclient import TestClient


def test_me_returns_the_signed_in_user(client: TestClient, sign_in: Callable[[str], dict[str, str]]) -> None:
    headers = sign_in("test-user-alice")

    response = client.get("/auth/me", headers=headers)

    body: dict[str, Any] = response.json()
    assert response.status_code == 200
    assert body["username"] == "test-user-alice"
    assert isinstance(body["id"], int)


@pytest.mark.parametrize(
    "headers",
    [{}, {"Authorization": "Bearer not-a-real-token"}, {"Authorization": "Basic dGVzdDp0ZXN0"}],
    ids=["no-header", "unknown-token", "wrong-scheme"],
)
def test_me_rejects_a_missing_or_unknown_token(client: TestClient, headers: dict[str, str]) -> None:
    response = client.get("/auth/me", headers=headers)

    assert response.status_code == 401
    assert response.headers["www-authenticate"] == "Bearer"


def test_me_rejects_an_expired_token(
    client: TestClient, sign_in: Callable[[str], dict[str, str]], database_url: str
) -> None:
    headers = sign_in("test-user-alice")
    with psycopg.connect(database_url, autocommit=True) as conn:
        conn.execute("UPDATE auth_tokens SET expires_at = now() - interval '1 hour'")

    response = client.get("/auth/me", headers=headers)

    assert response.status_code == 401
