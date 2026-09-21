from typing import Any, Final

from forum_be_python.schemas import PostOut

POSTS_SELECT: Final = (
    "SELECT p.id, p.thread_id, p.parent_id, u.username, p.body, p.created_at, p.updated_at"
    " FROM posts p JOIN users u ON u.id = p.author_id"
)


def to_post(row: tuple[Any, ...]) -> PostOut:
    return PostOut(
        id=row[0],
        thread_id=row[1],
        parent_id=row[2],
        author=row[3],
        body=row[4],
        created_at=row[5],
        updated_at=row[6],
    )
