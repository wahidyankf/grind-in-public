---
tldr: "Uses transactions to preserve invariants and an outbox to connect database state with asynchronous work."
when_to_use: "Use when several writes must succeed or fail as one business operation."
---

# Transactions and ACID

Atomicity groups effects, consistency preserves declared invariants, isolation controls concurrent observations, and
durability defines what survives a committed failure. Keep transactions short and free of network calls.

```sql
BEGIN;

WITH created AS (
    INSERT INTO cases (tenant_id, title, status)
    VALUES (1, 'Review transfer 2', 'open')
    RETURNING tenant_id, case_id
)
INSERT INTO outbox_events (
    tenant_id, aggregate_type, aggregate_id, event_type, payload
)
SELECT tenant_id, 'case', case_id, 'case.opened',
       jsonb_build_object('case_id', case_id)
FROM created;

COMMIT;
```

```text
request -> PostgreSQL transaction -> case row + outbox row -> COMMIT
                                                           |
publisher <- claim with SKIP LOCKED <- retry until publish -+
```

The outbox makes state and publish intent atomic, not end-to-end delivery exactly once. Consumers still need an inbox or
idempotent effect because a crash after publishing but before marking `published_at` causes redelivery.

Use savepoints sparingly for partial recovery. Never leave a transaction open while waiting for user input: it holds a
snapshot, may retain locks, blocks vacuum cleanup, and increases failover recovery work.

Related learning: [messaging and consistency](../system-design-interview/patterns/002-messaging-and-consistency.md).
