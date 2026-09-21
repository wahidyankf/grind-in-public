from datetime import UTC, datetime

from forum_be_python.schemas import PostOut
from forum_be_python.tree import build_reply_tree

_NOW = datetime(2026, 1, 1, tzinfo=UTC)


def _post(post_id: int, parent_id: int | None = None) -> PostOut:
    return PostOut(
        id=post_id,
        thread_id=1,
        parent_id=parent_id,
        author="test-user-a",
        body=f"body {post_id}",
        created_at=_NOW,
        updated_at=_NOW,
    )


def test_replies_nest_under_their_parents_to_any_depth_and_keep_sibling_order() -> None:
    flat = [_post(1), _post(2, 1), _post(3, 2), _post(4, 1), _post(5)]

    roots = build_reply_tree(flat)

    assert [root.id for root in roots] == [1, 5]
    assert [reply.id for reply in roots[0].replies] == [2, 4]
    assert [reply.id for reply in roots[0].replies[0].replies] == [3]
    assert roots[1].replies == []
