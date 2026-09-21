from collections.abc import Callable
from datetime import datetime
from typing import Any

from fastapi.testclient import TestClient

SignIn = Callable[[str], dict[str, str]]

_OPENING = "When should I add an index?"


def _make_thread(client: TestClient, headers: dict[str, str]) -> dict[str, Any]:
    topic: dict[str, Any] = client.post("/topics", json={"title": "Databases"}, headers=headers).json()
    thread: dict[str, Any] = client.post(
        f"/topics/{topic['id']}/threads", json={"title": "Indexes", "body": _OPENING}, headers=headers
    ).json()
    return thread


def test_reply_adds_a_top_level_post_to_the_thread(client: TestClient, sign_in: SignIn) -> None:
    thread = _make_thread(client, sign_in("test-user-alice"))

    response = client.post(
        f"/threads/{thread['id']}/posts",
        json={"body": "Add one on the filter column."},
        headers=sign_in("test-user-bob"),
    )

    body: dict[str, Any] = response.json()
    assert response.status_code == 201
    assert body["author"] == "test-user-bob"
    assert body["thread_id"] == thread["id"]
    assert body["parent_id"] is None
    posts = client.get(f"/threads/{thread['id']}").json()["posts"]
    assert [post["body"] for post in posts] == [_OPENING, "Add one on the filter column."]


def test_reply_requires_a_token(client: TestClient, sign_in: SignIn) -> None:
    thread = _make_thread(client, sign_in("test-user-alice"))

    response = client.post(f"/threads/{thread['id']}/posts", json={"body": "Anonymous"})

    assert response.status_code == 401


def test_reply_returns_404_for_a_missing_thread(client: TestClient, sign_in: SignIn) -> None:
    response = client.post("/threads/999999/posts", json={"body": "Hello"}, headers=sign_in("test-user-alice"))

    assert response.status_code == 404


def test_reply_rejects_an_empty_body(client: TestClient, sign_in: SignIn) -> None:
    alice = sign_in("test-user-alice")
    thread = _make_thread(client, alice)

    response = client.post(f"/threads/{thread['id']}/posts", json={"body": ""}, headers=alice)

    assert response.status_code == 422


def _opening_post_id(client: TestClient, thread_id: int) -> int:
    opening_id: int = client.get(f"/threads/{thread_id}").json()["posts"][0]["id"]
    return opening_id


def test_reply_can_answer_another_post_and_nests_under_it(client: TestClient, sign_in: SignIn) -> None:
    alice = sign_in("test-user-alice")
    bob = sign_in("test-user-bob")
    thread = _make_thread(client, alice)
    url = f"/threads/{thread['id']}/posts"

    reply = client.post(url, json={"body": "Reply", "parent_id": _opening_post_id(client, thread["id"])}, headers=bob)
    nested = client.post(url, json={"body": "Nested", "parent_id": reply.json()["id"]}, headers=alice)

    assert reply.status_code == 201
    assert reply.json()["parent_id"] == _opening_post_id(client, thread["id"])
    assert nested.status_code == 201
    roots = client.get(f"/threads/{thread['id']}").json()["posts"]
    assert len(roots) == 1
    assert [post["body"] for post in roots[0]["replies"]] == ["Reply"]
    assert [post["body"] for post in roots[0]["replies"][0]["replies"]] == ["Nested"]


def test_reply_rejects_a_parent_from_another_thread(client: TestClient, sign_in: SignIn) -> None:
    alice = sign_in("test-user-alice")
    first = _make_thread(client, alice)
    second = _make_thread(client, alice)

    response = client.post(
        f"/threads/{second['id']}/posts",
        json={"body": "Across threads", "parent_id": _opening_post_id(client, first["id"])},
        headers=alice,
    )

    assert response.status_code == 422


def test_reply_rejects_a_parent_that_does_not_exist(client: TestClient, sign_in: SignIn) -> None:
    alice = sign_in("test-user-alice")
    thread = _make_thread(client, alice)

    response = client.post(
        f"/threads/{thread['id']}/posts", json={"body": "Orphan", "parent_id": 999999}, headers=alice
    )

    assert response.status_code == 422


def test_get_post_returns_the_post_with_its_whole_reply_subtree_and_nothing_else(
    client: TestClient, sign_in: SignIn
) -> None:
    alice = sign_in("test-user-alice")
    bob = sign_in("test-user-bob")
    thread = _make_thread(client, alice)
    url = f"/threads/{thread['id']}/posts"
    opening_id = _opening_post_id(client, thread["id"])
    branch = client.post(url, json={"body": "A", "parent_id": opening_id}, headers=bob).json()
    first = client.post(url, json={"body": "A1", "parent_id": branch["id"]}, headers=alice).json()
    client.post(url, json={"body": "A1a", "parent_id": first["id"]}, headers=bob)
    client.post(url, json={"body": "A2", "parent_id": branch["id"]}, headers=alice)
    client.post(url, json={"body": "B", "parent_id": opening_id}, headers=alice)

    response = client.get(f"/posts/{branch['id']}")

    body: dict[str, Any] = response.json()
    assert response.status_code == 200
    assert body["body"] == "A"
    assert [reply["body"] for reply in body["replies"]] == ["A1", "A2"]
    assert [reply["body"] for reply in body["replies"][0]["replies"]] == ["A1a"]
    assert body["replies"][1]["replies"] == []


def test_get_post_returns_404_for_a_missing_post(client: TestClient) -> None:
    response = client.get("/posts/999999")

    assert response.status_code == 404


def _reply(client: TestClient, headers: dict[str, str], thread_id: int, body: str, parent_id: int | None) -> int:
    created: dict[str, Any] = client.post(
        f"/threads/{thread_id}/posts", json={"body": body, "parent_id": parent_id}, headers=headers
    ).json()
    post_id: int = created["id"]
    return post_id


def test_update_post_changes_the_body_and_moves_updated_at_forward(client: TestClient, sign_in: SignIn) -> None:
    alice = sign_in("test-user-alice")
    thread = _make_thread(client, alice)
    post_id = _reply(client, alice, thread["id"], "Before", None)

    response = client.patch(f"/posts/{post_id}", json={"body": "After"}, headers=alice)

    body: dict[str, Any] = response.json()
    assert response.status_code == 200
    assert body["body"] == "After"
    assert datetime.fromisoformat(body["updated_at"]) > datetime.fromisoformat(body["created_at"])


def test_update_post_is_forbidden_for_anyone_but_its_author(client: TestClient, sign_in: SignIn) -> None:
    alice = sign_in("test-user-alice")
    thread = _make_thread(client, alice)
    post_id = _reply(client, sign_in("test-user-bob"), thread["id"], "Bob's", None)

    response = client.patch(f"/posts/{post_id}", json={"body": "Hijacked"}, headers=alice)

    assert response.status_code == 403
    assert client.get(f"/posts/{post_id}").json()["body"] == "Bob's"


def test_update_post_requires_a_token(client: TestClient, sign_in: SignIn) -> None:
    alice = sign_in("test-user-alice")
    post_id = _reply(client, alice, _make_thread(client, alice)["id"], "Mine", None)

    response = client.patch(f"/posts/{post_id}", json={"body": "Anonymous"})

    assert response.status_code == 401


def test_update_post_returns_404_for_a_missing_post(client: TestClient, sign_in: SignIn) -> None:
    response = client.patch("/posts/999999", json={"body": "Nothing"}, headers=sign_in("test-user-alice"))

    assert response.status_code == 404


def test_update_post_rejects_an_empty_body(client: TestClient, sign_in: SignIn) -> None:
    alice = sign_in("test-user-alice")
    post_id = _reply(client, alice, _make_thread(client, alice)["id"], "Mine", None)

    response = client.patch(f"/posts/{post_id}", json={"body": ""}, headers=alice)

    assert response.status_code == 422


def test_delete_post_removes_it_together_with_all_its_replies_and_keeps_its_siblings(
    client: TestClient, sign_in: SignIn
) -> None:
    alice = sign_in("test-user-alice")
    thread = _make_thread(client, alice)
    opening_id = _opening_post_id(client, thread["id"])
    branch = _reply(client, alice, thread["id"], "Branch", opening_id)
    _reply(client, alice, thread["id"], "Nested", branch)
    _reply(client, alice, thread["id"], "Keep", opening_id)

    response = client.delete(f"/posts/{branch}", headers=alice)

    assert response.status_code == 204
    roots = client.get(f"/threads/{thread['id']}").json()["posts"]
    assert [post["body"] for post in roots[0]["replies"]] == ["Keep"]
    assert client.get(f"/posts/{branch}").status_code == 404


def test_delete_post_is_forbidden_for_anyone_but_its_author(client: TestClient, sign_in: SignIn) -> None:
    alice = sign_in("test-user-alice")
    post_id = _reply(client, sign_in("test-user-bob"), _make_thread(client, alice)["id"], "Bob's", None)

    response = client.delete(f"/posts/{post_id}", headers=alice)

    assert response.status_code == 403
    assert client.get(f"/posts/{post_id}").status_code == 200


def test_delete_post_requires_a_token(client: TestClient, sign_in: SignIn) -> None:
    alice = sign_in("test-user-alice")
    post_id = _reply(client, alice, _make_thread(client, alice)["id"], "Mine", None)

    response = client.delete(f"/posts/{post_id}")

    assert response.status_code == 401


def test_delete_post_returns_404_for_a_missing_post(client: TestClient, sign_in: SignIn) -> None:
    response = client.delete("/posts/999999", headers=sign_in("test-user-alice"))

    assert response.status_code == 404
