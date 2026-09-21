from typing import Final

import psycopg.errors
from fastapi import APIRouter, HTTPException
from psycopg import Connection
from psycopg.rows import TupleRow

from forum_be_python.db import one
from forum_be_python.deps import CurrentUser, PoolDep
from forum_be_python.post_rows import POSTS_SELECT, to_post
from forum_be_python.schemas import PostIn, PostOut, PostPatch, UserOut
from forum_be_python.tree import build_reply_tree

router = APIRouter(tags=["posts"])

# Walks down from one post through its replies, replies of replies, and so on, whatever the depth.
_SUBTREE: Final = (
    "WITH RECURSIVE subtree AS ("
    " SELECT id FROM posts WHERE id = %s"
    " UNION ALL"
    " SELECT p.id FROM posts p JOIN subtree s ON p.parent_id = s.id"
    ") " + POSTS_SELECT + " WHERE p.id IN (SELECT id FROM subtree) ORDER BY p.id"
)


def _require_author(conn: Connection[TupleRow], post_id: int, user: UserOut) -> None:
    row = conn.execute("SELECT author_id FROM posts WHERE id = %s", (post_id,)).fetchone()
    if row is None:
        raise HTTPException(status_code=404, detail="Post not found")
    if row[0] != user.id:
        raise HTTPException(status_code=403, detail="Only the author can change this post")


@router.post("/threads/{thread_id}/posts", status_code=201)
def create_post(thread_id: int, post: PostIn, user: CurrentUser, pool: PoolDep) -> PostOut:
    try:
        with pool.connection() as conn:
            row = one(
                conn.execute(
                    "INSERT INTO posts (thread_id, parent_id, author_id, body) VALUES (%s, %s, %s, %s)"
                    " RETURNING id, created_at, updated_at",
                    (thread_id, post.parent_id, user.id, post.body),
                ).fetchone()
            )
    except psycopg.errors.ForeignKeyViolation as error:
        # The composite key is what ties a reply to a parent in the same thread; the only other key is the thread.
        if error.diag.constraint_name == "posts_parent_fk":
            raise HTTPException(status_code=422, detail="Parent post does not exist in this thread") from None
        raise HTTPException(status_code=404, detail="Thread not found") from None
    return PostOut(
        id=row[0],
        thread_id=thread_id,
        parent_id=post.parent_id,
        author=user.username,
        body=post.body,
        created_at=row[1],
        updated_at=row[2],
    )


@router.get("/posts/{post_id}")
def get_post(post_id: int, pool: PoolDep) -> PostOut:
    with pool.connection() as conn:
        rows = conn.execute(_SUBTREE, (post_id,)).fetchall()
    if not rows:
        raise HTTPException(status_code=404, detail="Post not found")
    # The requested post's own parent is outside the subtree, so it comes back as the only top-level node.
    return build_reply_tree([to_post(row) for row in rows])[0]


@router.patch("/posts/{post_id}")
def update_post(post_id: int, patch: PostPatch, user: CurrentUser, pool: PoolDep) -> PostOut:
    with pool.connection() as conn:
        _require_author(conn, post_id, user)
        conn.execute("UPDATE posts SET body = %s, updated_at = now() WHERE id = %s", (patch.body, post_id))
        row = one(conn.execute(POSTS_SELECT + " WHERE p.id = %s", (post_id,)).fetchone())
    return to_post(row)


@router.delete("/posts/{post_id}", status_code=204)
def delete_post(post_id: int, user: CurrentUser, pool: PoolDep) -> None:
    with pool.connection() as conn:
        _require_author(conn, post_id, user)
        # The composite foreign key cascades, so the replies go with the post.
        conn.execute("DELETE FROM posts WHERE id = %s", (post_id,))
