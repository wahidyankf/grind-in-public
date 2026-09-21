from collections.abc import Callable
from typing import Any

from fastapi.testclient import TestClient

SignIn = Callable[[str], dict[str, str]]


def _create_topic(client: TestClient, headers: dict[str, str], title: str) -> dict[str, Any]:
    body: dict[str, Any] = client.post("/topics", json={"title": title}, headers=headers).json()
    return body


def test_create_topic_records_the_signed_in_user_as_its_author(client: TestClient, sign_in: SignIn) -> None:
    headers = sign_in("test-user-alice")

    response = client.post(
        "/topics", json={"title": "Databases", "description": "Postgres and friends"}, headers=headers
    )

    body: dict[str, Any] = response.json()
    assert response.status_code == 201
    assert body["title"] == "Databases"
    assert body["description"] == "Postgres and friends"
    assert body["created_by"] == "test-user-alice"
    assert isinstance(body["id"], int)


def test_create_topic_requires_a_token(client: TestClient) -> None:
    response = client.post("/topics", json={"title": "Databases"})

    assert response.status_code == 401


def test_create_topic_rejects_an_empty_title(client: TestClient, sign_in: SignIn) -> None:
    response = client.post("/topics", json={"title": ""}, headers=sign_in("test-user-alice"))

    assert response.status_code == 422


def test_list_topics_is_public_and_keeps_creation_order(client: TestClient, sign_in: SignIn) -> None:
    headers = sign_in("test-user-alice")
    _create_topic(client, headers, "First")
    _create_topic(client, headers, "Second")

    response = client.get("/topics")

    assert response.status_code == 200
    assert [topic["title"] for topic in response.json()] == ["First", "Second"]


def test_get_topic_returns_one_topic(client: TestClient, sign_in: SignIn) -> None:
    created = _create_topic(client, sign_in("test-user-alice"), "Databases")

    response = client.get(f"/topics/{created['id']}")

    assert response.status_code == 200
    assert response.json()["title"] == "Databases"


def test_get_topic_returns_404_for_a_missing_topic(client: TestClient) -> None:
    response = client.get("/topics/999999")

    assert response.status_code == 404


def test_update_topic_changes_only_the_given_fields(client: TestClient, sign_in: SignIn) -> None:
    headers = sign_in("test-user-alice")
    created = client.post("/topics", json={"title": "Old", "description": "Keep me"}, headers=headers).json()

    response = client.patch(f"/topics/{created['id']}", json={"title": "New"}, headers=headers)

    body: dict[str, Any] = response.json()
    assert response.status_code == 200
    assert body["title"] == "New"
    assert body["description"] == "Keep me"


def test_update_topic_is_forbidden_for_another_user(client: TestClient, sign_in: SignIn) -> None:
    created = _create_topic(client, sign_in("test-user-alice"), "Mine")

    response = client.patch(f"/topics/{created['id']}", json={"title": "Taken"}, headers=sign_in("test-user-bob"))

    assert response.status_code == 403
    assert client.get(f"/topics/{created['id']}").json()["title"] == "Mine"


def test_update_topic_requires_a_token(client: TestClient, sign_in: SignIn) -> None:
    created = _create_topic(client, sign_in("test-user-alice"), "Mine")

    response = client.patch(f"/topics/{created['id']}", json={"title": "Anonymous"})

    assert response.status_code == 401


def test_update_topic_returns_404_for_a_missing_topic(client: TestClient, sign_in: SignIn) -> None:
    response = client.patch("/topics/999999", json={"title": "Nothing"}, headers=sign_in("test-user-alice"))

    assert response.status_code == 404


def test_delete_topic_removes_it(client: TestClient, sign_in: SignIn) -> None:
    headers = sign_in("test-user-alice")
    created = _create_topic(client, headers, "Doomed")

    response = client.delete(f"/topics/{created['id']}", headers=headers)

    assert response.status_code == 204
    assert client.get(f"/topics/{created['id']}").status_code == 404


def test_delete_topic_is_forbidden_for_another_user(client: TestClient, sign_in: SignIn) -> None:
    created = _create_topic(client, sign_in("test-user-alice"), "Mine")

    response = client.delete(f"/topics/{created['id']}", headers=sign_in("test-user-bob"))

    assert response.status_code == 403
    assert client.get(f"/topics/{created['id']}").status_code == 200


def test_delete_topic_requires_a_token(client: TestClient, sign_in: SignIn) -> None:
    created = _create_topic(client, sign_in("test-user-alice"), "Mine")

    response = client.delete(f"/topics/{created['id']}")

    assert response.status_code == 401


def test_delete_topic_returns_404_for_a_missing_topic(client: TestClient, sign_in: SignIn) -> None:
    response = client.delete("/topics/999999", headers=sign_in("test-user-alice"))

    assert response.status_code == 404
