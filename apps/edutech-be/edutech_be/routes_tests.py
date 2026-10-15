from collections.abc import Iterator

import pytest
from fastapi.testclient import TestClient

from edutech_be.fake_store import FakeAuthStore
from edutech_be.main import app
from edutech_be.routes import get_store

PASSWORD = "correct horse"


def test_health_returns_ok() -> None:
    response = TestClient(app).get("/health")

    assert response.status_code == 200
    assert response.json() == {"status": "ok"}


@pytest.fixture
def client() -> Iterator[TestClient]:
    store = FakeAuthStore()
    app.dependency_overrides[get_store] = lambda: store
    yield TestClient(app)
    app.dependency_overrides.clear()


def _register_and_login(client: TestClient) -> dict[str, str]:
    client.post("/auth/register", json={"email": "a@example.com", "password": PASSWORD})
    response = client.post(
        "/auth/login", json={"email": "a@example.com", "password": PASSWORD}
    )
    return response.json()


def test_register_returns_201_with_normalized_email(client: TestClient) -> None:
    response = client.post(
        "/auth/register", json={"email": " A@Example.com ", "password": PASSWORD}
    )

    assert response.status_code == 201
    assert response.json() == {"email": "a@example.com"}


def test_register_with_taken_email_is_409(client: TestClient) -> None:
    body = {"email": "a@example.com", "password": PASSWORD}
    client.post("/auth/register", json=body)

    response = client.post("/auth/register", json=body)

    assert response.status_code == 409
    assert response.json() == {"detail": "Email is already registered"}


def test_register_with_weak_password_is_422(client: TestClient) -> None:
    response = client.post(
        "/auth/register", json={"email": "a@example.com", "password": "short"}
    )

    assert response.status_code == 422


def test_register_with_bad_email_is_422(client: TestClient) -> None:
    response = client.post(
        "/auth/register", json={"email": "not-an-email", "password": PASSWORD}
    )

    assert response.status_code == 422


def test_login_returns_bearer_tokens_and_expiry(client: TestClient) -> None:
    tokens = _register_and_login(client)

    assert tokens["token_type"] == "bearer"
    assert tokens["expires_in"] == 900
    assert tokens["access_token"] and tokens["refresh_token"]


def test_login_failures_return_the_same_401(client: TestClient) -> None:
    client.post("/auth/register", json={"email": "a@example.com", "password": PASSWORD})

    wrong_password = client.post(
        "/auth/login", json={"email": "a@example.com", "password": "wrong password"}
    )
    unknown_email = client.post(
        "/auth/login", json={"email": "nobody@example.com", "password": PASSWORD}
    )

    assert wrong_password.status_code == unknown_email.status_code == 401
    assert wrong_password.json() == unknown_email.json()


def test_me_returns_the_current_user(client: TestClient) -> None:
    tokens = _register_and_login(client)

    response = client.get(
        "/auth/me", headers={"Authorization": f"Bearer {tokens['access_token']}"}
    )

    assert response.status_code == 200
    assert response.json()["email"] == "a@example.com"
    assert set(response.json()) == {"id", "email"}


def test_me_without_token_is_401_with_bearer_challenge(client: TestClient) -> None:
    response = client.get("/auth/me")

    assert response.status_code == 401
    assert response.headers["WWW-Authenticate"] == "Bearer"


def test_me_with_refresh_token_is_401(client: TestClient) -> None:
    tokens = _register_and_login(client)

    response = client.get(
        "/auth/me", headers={"Authorization": f"Bearer {tokens['refresh_token']}"}
    )

    assert response.status_code == 401


def test_refresh_returns_a_new_access_token(client: TestClient) -> None:
    tokens = _register_and_login(client)

    response = client.post(
        "/auth/refresh", json={"refresh_token": tokens["refresh_token"]}
    )

    assert response.status_code == 200
    assert response.json()["token_type"] == "bearer"
    assert response.json()["expires_in"] == 900
    assert "refresh_token" not in response.json()


def test_refresh_with_bad_token_is_401(client: TestClient) -> None:
    response = client.post("/auth/refresh", json={"refresh_token": "nope"})

    assert response.status_code == 401


def test_logout_returns_204_and_revokes_the_tokens(client: TestClient) -> None:
    tokens = _register_and_login(client)
    headers = {"Authorization": f"Bearer {tokens['access_token']}"}

    response = client.post(
        "/auth/logout", headers=headers, json={"refresh_token": tokens["refresh_token"]}
    )

    assert response.status_code == 204
    assert client.get("/auth/me", headers=headers).status_code == 401
    assert (
        client.post(
            "/auth/refresh", json={"refresh_token": tokens["refresh_token"]}
        ).status_code
        == 401
    )
