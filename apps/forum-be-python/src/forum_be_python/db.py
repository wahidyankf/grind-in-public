from pathlib import Path
from typing import LiteralString, cast

from psycopg import Connection
from psycopg.rows import TupleRow
from psycopg_pool import ConnectionPool

Pool = ConnectionPool[Connection[TupleRow]]

_SCHEMA_FILE = Path(__file__).with_name("schema.sql")


def create_pool(database_url: str) -> Pool:
    pool: Pool = ConnectionPool(database_url, min_size=1, max_size=5, open=False)
    pool.open(wait=True)
    return pool


def init_schema(pool: Pool) -> None:
    # The file ships with the code and carries no user input, so it is safe to treat as a literal statement.
    # Every statement in it is idempotent, so the same file serves a fresh database and a restart.
    schema = cast(LiteralString, _SCHEMA_FILE.read_text())
    with pool.connection() as conn:
        conn.execute(schema)


def one[T](row: T | None) -> T:
    """Unwrap a row the statement itself guarantees, such as the result of INSERT ... RETURNING."""
    if row is None:
        raise LookupError("the statement returned no row")
    return row
