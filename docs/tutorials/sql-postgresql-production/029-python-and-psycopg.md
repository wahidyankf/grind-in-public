---
tldr: "Builds a complete typed Psycopg transaction boundary with pooling, safe SQL, retries, and an outbox write."
when_to_use: "Use when connecting a Python service to PostgreSQL in production."
---

# Python and Psycopg

Create a disposable project and install binary Psycopg plus its pool:

```bash
uv init --app --python 3.14 pg-client
cd pg-client
uv add 'psycopg[binary,pool]>=3.2,<4'
```

Save this complete example as `main.py`. Parameters are bound separately from SQL; never interpolate untrusted values.

```python
from __future__ import annotations

import os
import random
import time
from dataclasses import dataclass
from decimal import Decimal
from typing import Any, Callable, TypeAlias, TypeVar

import psycopg
from psycopg import sql
from psycopg.errors import DeadlockDetected, SerializationFailure
from psycopg.rows import class_row
from psycopg.types.json import Jsonb
from psycopg_pool import ConnectionPool


@dataclass(frozen=True, slots=True)
class CreatedCase:
    tenant_id: int
    case_id: int
    version: int


T = TypeVar("T")
Pool: TypeAlias = ConnectionPool[tuple[Any, ...]]


def retry_transaction(operation: Callable[[], T], attempts: int = 4) -> T:
    for attempt in range(attempts):
        try:
            return operation()
        except (SerializationFailure, DeadlockDetected):
            if attempt == attempts - 1:
                raise
            time.sleep((2**attempt) * 0.025 + random.random() * 0.025)
    raise AssertionError("unreachable")


def create_case(pool: Pool, tenant_id: int, title: str) -> CreatedCase:
    def transaction() -> CreatedCase:
        with pool.connection() as connection:
            with connection.transaction():
                with connection.cursor(row_factory=class_row(CreatedCase)) as cursor:
                    cursor.execute(
                        """
                        INSERT INTO cases (tenant_id, title, status)
                        VALUES (%s, %s, 'open')
                        RETURNING tenant_id, case_id, version
                        """,
                        (tenant_id, title),
                    )
                    created = cursor.fetchone()
                    if created is None:
                        raise RuntimeError("case insert returned no row")

                    cursor.execute(
                        """
                        INSERT INTO outbox_events (
                            tenant_id, aggregate_type, aggregate_id, event_type, payload
                        )
                        VALUES (%s, 'case', %s, 'case.opened', %s)
                        """,
                        (tenant_id, created.case_id, Jsonb({"case_id": created.case_id})),
                    )
                    return created

    return retry_transaction(transaction)


def recent_transactions(pool: Pool, tenant_id: int) -> None:
    allowed_order = {"occurred_at", "amount"}
    requested_order = "occurred_at"
    if requested_order not in allowed_order:
        raise ValueError("unsupported order")

    query = sql.SQL(
        """
        SELECT transaction_id, amount, occurred_at
        FROM transactions
        WHERE tenant_id = %s AND amount >= %s
        ORDER BY {} DESC, transaction_id DESC
        LIMIT %s
        """
    ).format(sql.Identifier(requested_order))

    with pool.connection() as connection, connection.cursor() as cursor:
        cursor.execute(query, (tenant_id, Decimal("100.00"), 20))
        for row in cursor:
            print(row)


def main() -> None:
    database_url = os.environ["STUDY_DATABASE_URL"]
    with ConnectionPool(
        conninfo=database_url,
        min_size=1,
        max_size=4,
        timeout=2.0,
        kwargs={"autocommit": False},
    ) as pool:
        print(create_case(pool, 1, "Review idempotent payment"))


if __name__ == "__main__":
    main()
```

The example intentionally keeps pool capacity small. In Kubernetes, total possible connections are approximately
`replicas * processes * pool max`, plus jobs and operator connections. Bound that number below database capacity and
keep headroom for failover. A transaction-pooling proxy can absorb connection churn, but session features and prepared
state need compatibility review.

Use `COPY` for bulk ingestion, server-side cursors or chunked reads for large results, cancellation/timeouts for bounded
work, and separate pools for workloads with different latency budgets. Never retry every error; only retry transient
whole-transaction failures whose side effects are repeatable.
