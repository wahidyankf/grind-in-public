from typing import Any, Final

import psycopg.errors
from fastapi import APIRouter, HTTPException
from psycopg import Connection
from psycopg.rows import TupleRow

from forum_be_python.db import one
from forum_be_python.deps import CurrentUser, PoolDep
from forum_be_python.post_rows import POSTS_SELECT, to_post
from forum_be_python.schemas import ThreadDetail, ThreadIn, ThreadOut, ThreadPatch, UserOut
from forum_be_python.tree import build_reply_tree

router = APIRouter(tags=["threads"])

_SELECT: Final = (
    "SELECT th.id, th.topic_id, th.title, u.username, th.created_at"
    " FROM threads th JOIN users u ON u.id = th.created_by"
)


def _to_thread(row: tuple[Any, ...]) -> ThreadOut:
    return ThreadOut(id=row[0], topic_id=row[1], title=row[2], created_by=row[3], created_at=row[4])


def _require_author(conn: Connection[TupleRow], thread_id: int, user: UserOut) -> None:
    row = conn.execute("SELECT created_by FROM threads WHERE id = %s", (thread_id,)).fetchone()
    if row is None:
        raise HTTPException(status_code=404, detail="Thread not found")
    if row[0] != user.id:
        raise HTTPException(status_code=403, detail="Only the author can change this thread")


@router.post("/topics/{topic_id}/threads", status_code=201)
def create_thread(topic_id: int, thread: ThreadIn, user: CurrentUser, pool: PoolDep) -> ThreadOut:
    # The thread and its opening post are written in one transaction, so a thread never exists without a post.
    # A missing topic surfaces as a foreign-key violation, which also covers it vanishing between two statements.
    try:
        with pool.connection() as conn:
            row = one(
                conn.execute(
                    "INSERT INTO threads (topic_id, title, created_by) VALUES (%s, %s, %s) RETURNING id, created_at",
                    (topic_id, thread.title, user.id),
                ).fetchone()
            )
            conn.execute(
                "INSERT INTO posts (thread_id, author_id, body) VALUES (%s, %s, %s)",
                (row[0], user.id, thread.body),
            )
    except psycopg.errors.ForeignKeyViolation:
        raise HTTPException(status_code=404, detail="Topic not found") from None
    return ThreadOut(
        id=row[0], topic_id=topic_id, title=thread.title, created_by=user.username, created_at=row[1]
    )


@router.get("/topics/{topic_id}/threads")
def list_threads(topic_id: int, pool: PoolDep) -> list[ThreadOut]:
    with pool.connection() as conn:
        if conn.execute("SELECT 1 FROM topics WHERE id = %s", (topic_id,)).fetchone() is None:
            raise HTTPException(status_code=404, detail="Topic not found")
        rows = conn.execute(_SELECT + " WHERE th.topic_id = %s ORDER BY th.id", (topic_id,)).fetchall()
    return [_to_thread(row) for row in rows]


@router.get("/threads/{thread_id}")
def get_thread(thread_id: int, pool: PoolDep) -> ThreadDetail:
    with pool.connection() as conn:
        thread = conn.execute(_SELECT + " WHERE th.id = %s", (thread_id,)).fetchone()
        if thread is None:
            raise HTTPException(status_code=404, detail="Thread not found")
        # Replies always get a larger id than their parent, so id order is also a valid tree order.
        posts = conn.execute(POSTS_SELECT + " WHERE p.thread_id = %s ORDER BY p.id", (thread_id,)).fetchall()
    return ThreadDetail(
        **_to_thread(thread).model_dump(),
        posts=build_reply_tree([to_post(row) for row in posts]),
    )


@router.patch("/threads/{thread_id}")
def update_thread(thread_id: int, patch: ThreadPatch, user: CurrentUser, pool: PoolDep) -> ThreadOut:
    with pool.connection() as conn:
        _require_author(conn, thread_id, user)
        conn.execute("UPDATE threads SET title = %s WHERE id = %s", (patch.title, thread_id))
        row = one(conn.execute(_SELECT + " WHERE th.id = %s", (thread_id,)).fetchone())
    return _to_thread(row)


@router.delete("/threads/{thread_id}", status_code=204)
def delete_thread(thread_id: int, user: CurrentUser, pool: PoolDep) -> None:
    with pool.connection() as conn:
        _require_author(conn, thread_id, user)
        conn.execute("DELETE FROM threads WHERE id = %s", (thread_id,))
