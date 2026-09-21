from collections.abc import Callable
from typing import Any

import psycopg
import pytest
from fastapi.testclient import TestClient

SignIn = Callable[[str], dict[str, str]]


def _create_topic(client: TestClient, headers: dict[str, str], title: str = "Databases") -> dict[str, Any]:
    body: dict[str, Any] = client.post("/topics", json={"title": title}, headers=headers).json()
    return body


def _create_thread(
    client: TestClient,
    headers: dict[str, str],
    topic_id: int,
    title: str = "Indexes",
    body: str = "When should I add one?",
) -> dict[str, Any]:
    created: dict[str, Any] = client.post(
        f"/topics/{topic_id}/threads", json={"title": title, "body": body}, headers=headers
    ).json()
    return created


def test_create_thread_records_the_signed_in_user_as_its_author(client: TestClient, sign_in: SignIn) -> None:
    headers = sign_in("test-user-alice")
    topic = _create_topic(client, headers)

    response = client.post(
        f"/topics/{topic['id']}/threads",
        json={"title": "Indexes", "body": "When should I add one?"},
        headers=headers,
    )

    body: dict[str, Any] = response.json()
    assert response.status_code == 201
    assert body["title"] == "Indexes"
    assert body["topic_id"] == topic["id"]
    assert body["created_by"] == "test-user-alice"
    assert isinstance(body["id"], int)


def test_create_thread_requires_a_token(client: TestClient, sign_in: SignIn) -> None:
    topic = _create_topic(client, sign_in("test-user-alice"))

    response = client.post(f"/topics/{topic['id']}/threads", json={"title": "Indexes", "body": "Text"})

    assert response.status_code == 401


def test_create_thread_returns_404_for_a_missing_topic(client: TestClient, sign_in: SignIn) -> None:
    response = client.post(
        "/topics/999999/threads",
        json={"title": "Indexes", "body": "Text"},
        headers=sign_in("test-user-alice"),
    )

    assert response.status_code == 404


@pytest.mark.parametrize(
    "payload",
    [{"title": "", "body": "Text"}, {"title": "Indexes", "body": ""}],
    ids=["empty-title", "empty-body"],
)
def test_create_thread_rejects_invalid_input(
    client: TestClient, sign_in: SignIn, payload: dict[str, str]
) -> None:
    headers = sign_in("test-user-alice")
    topic = _create_topic(client, headers)

    response = client.post(f"/topics/{topic['id']}/threads", json=payload, headers=headers)

    assert response.status_code == 422


def test_list_threads_returns_only_the_threads_of_that_topic_in_creation_order(
    client: TestClient, sign_in: SignIn
) -> None:
    headers = sign_in("test-user-alice")
    databases = _create_topic(client, headers, "Databases")
    python = _create_topic(client, headers, "Python")
    _create_thread(client, headers, databases["id"], title="First")
    _create_thread(client, headers, python["id"], title="Elsewhere")
    _create_thread(client, headers, databases["id"], title="Second")

    response = client.get(f"/topics/{databases['id']}/threads")

    assert response.status_code == 200
    assert [thread["title"] for thread in response.json()] == ["First", "Second"]


def test_list_threads_returns_404_for_a_missing_topic(client: TestClient) -> None:
    response = client.get("/topics/999999/threads")

    assert response.status_code == 404


def test_get_thread_returns_the_thread_with_its_opening_post(client: TestClient, sign_in: SignIn) -> None:
    headers = sign_in("test-user-alice")
    topic = _create_topic(client, headers)
    thread = _create_thread(client, headers, topic["id"], title="Indexes", body="When should I add one?")

    response = client.get(f"/threads/{thread['id']}")

    body: dict[str, Any] = response.json()
    assert response.status_code == 200
    assert body["title"] == "Indexes"
    assert body["created_by"] == "test-user-alice"
    assert len(body["posts"]) == 1
    opening = body["posts"][0]
    assert opening["body"] == "When should I add one?"
    assert opening["author"] == "test-user-alice"
    assert opening["parent_id"] is None
    assert opening["replies"] == []


def test_get_thread_returns_404_for_a_missing_thread(client: TestClient) -> None:
    response = client.get("/threads/999999")

    assert response.status_code == 404


def test_update_thread_changes_the_title(client: TestClient, sign_in: SignIn) -> None:
    headers = sign_in("test-user-alice")
    thread = _create_thread(client, headers, _create_topic(client, headers)["id"], title="Old")

    response = client.patch(f"/threads/{thread['id']}", json={"title": "New"}, headers=headers)

    assert response.status_code == 200
    assert response.json()["title"] == "New"


def test_update_thread_is_forbidden_for_another_user(client: TestClient, sign_in: SignIn) -> None:
    alice = sign_in("test-user-alice")
    thread = _create_thread(client, alice, _create_topic(client, alice)["id"], title="Mine")

    response = client.patch(f"/threads/{thread['id']}", json={"title": "Taken"}, headers=sign_in("test-user-bob"))

    assert response.status_code == 403
    assert client.get(f"/threads/{thread['id']}").json()["title"] == "Mine"


def test_update_thread_requires_a_token(client: TestClient, sign_in: SignIn) -> None:
    alice = sign_in("test-user-alice")
    thread = _create_thread(client, alice, _create_topic(client, alice)["id"])

    response = client.patch(f"/threads/{thread['id']}", json={"title": "Anonymous"})

    assert response.status_code == 401


def test_update_thread_returns_404_for_a_missing_thread(client: TestClient, sign_in: SignIn) -> None:
    response = client.patch("/threads/999999", json={"title": "Nothing"}, headers=sign_in("test-user-alice"))

    assert response.status_code == 404


def test_delete_thread_removes_it_together_with_its_posts(
    client: TestClient, sign_in: SignIn, database_url: str
) -> None:
    headers = sign_in("test-user-alice")
    thread = _create_thread(client, headers, _create_topic(client, headers)["id"])

    response = client.delete(f"/threads/{thread['id']}", headers=headers)

    assert response.status_code == 204
    assert client.get(f"/threads/{thread['id']}").status_code == 404
    with psycopg.connect(database_url) as conn:
        row = conn.execute("SELECT count(*) FROM posts").fetchone()
    assert row == (0,)


def test_delete_thread_is_forbidden_for_another_user(client: TestClient, sign_in: SignIn) -> None:
    alice = sign_in("test-user-alice")
    thread = _create_thread(client, alice, _create_topic(client, alice)["id"])

    response = client.delete(f"/threads/{thread['id']}", headers=sign_in("test-user-bob"))

    assert response.status_code == 403
    assert client.get(f"/threads/{thread['id']}").status_code == 200


def test_delete_thread_requires_a_token(client: TestClient, sign_in: SignIn) -> None:
    alice = sign_in("test-user-alice")
    thread = _create_thread(client, alice, _create_topic(client, alice)["id"])

    response = client.delete(f"/threads/{thread['id']}")

    assert response.status_code == 401


def test_delete_thread_returns_404_for_a_missing_thread(client: TestClient, sign_in: SignIn) -> None:
    response = client.delete("/threads/999999", headers=sign_in("test-user-alice"))

    assert response.status_code == 404


def test_deleting_a_topic_removes_its_threads(client: TestClient, sign_in: SignIn) -> None:
    headers = sign_in("test-user-alice")
    topic = _create_topic(client, headers)
    thread = _create_thread(client, headers, topic["id"])

    client.delete(f"/topics/{topic['id']}", headers=headers)

    assert client.get(f"/threads/{thread['id']}").status_code == 404
