from typing import Any, Final

from fastapi import APIRouter, HTTPException
from psycopg import Connection
from psycopg.rows import TupleRow

from forum_be_python.db import one
from forum_be_python.deps import CurrentUser, PoolDep
from forum_be_python.schemas import TopicIn, TopicOut, TopicPatch, UserOut

router = APIRouter(prefix="/topics", tags=["topics"])

_SELECT: Final = (
    "SELECT t.id, t.title, t.description, u.username, t.created_at FROM topics t JOIN users u ON u.id = t.created_by"
)


def _to_topic(row: tuple[Any, ...]) -> TopicOut:
    return TopicOut(id=row[0], title=row[1], description=row[2], created_by=row[3], created_at=row[4])


def _require_owner(conn: Connection[TupleRow], topic_id: int, user: UserOut) -> None:
    row = conn.execute("SELECT created_by FROM topics WHERE id = %s", (topic_id,)).fetchone()
    if row is None:
        raise HTTPException(status_code=404, detail="Topic not found")
    if row[0] != user.id:
        raise HTTPException(status_code=403, detail="Only the author can change this topic")


@router.post("", status_code=201)
def create_topic(topic: TopicIn, user: CurrentUser, pool: PoolDep) -> TopicOut:
    with pool.connection() as conn:
        row = one(
            conn.execute(
                "INSERT INTO topics (title, description, created_by) VALUES (%s, %s, %s) RETURNING id, created_at",
                (topic.title, topic.description, user.id),
            ).fetchone()
        )
    return TopicOut(
        id=row[0],
        title=topic.title,
        description=topic.description,
        created_by=user.username,
        created_at=row[1],
    )


@router.get("")
def list_topics(pool: PoolDep) -> list[TopicOut]:
    with pool.connection() as conn:
        rows = conn.execute(_SELECT + " ORDER BY t.id").fetchall()
    return [_to_topic(row) for row in rows]


@router.get("/{topic_id}")
def get_topic(topic_id: int, pool: PoolDep) -> TopicOut:
    with pool.connection() as conn:
        row = conn.execute(_SELECT + " WHERE t.id = %s", (topic_id,)).fetchone()
    if row is None:
        raise HTTPException(status_code=404, detail="Topic not found")
    return _to_topic(row)


@router.patch("/{topic_id}")
def update_topic(topic_id: int, patch: TopicPatch, user: CurrentUser, pool: PoolDep) -> TopicOut:
    with pool.connection() as conn:
        _require_owner(conn, topic_id, user)
        conn.execute(
            "UPDATE topics SET title = COALESCE(%s, title), description = COALESCE(%s, description) WHERE id = %s",
            (patch.title, patch.description, topic_id),
        )
        row = one(conn.execute(_SELECT + " WHERE t.id = %s", (topic_id,)).fetchone())
    return _to_topic(row)


@router.delete("/{topic_id}", status_code=204)
def delete_topic(topic_id: int, user: CurrentUser, pool: PoolDep) -> None:
    with pool.connection() as conn:
        _require_owner(conn, topic_id, user)
        conn.execute("DELETE FROM topics WHERE id = %s", (topic_id,))
