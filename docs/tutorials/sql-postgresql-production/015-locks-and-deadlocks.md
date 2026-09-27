---
tldr: "Controls row and table locks, diagnoses waits, and prevents deadlocks through consistent acquisition order."
when_to_use: "Use when writes block, workers compete, or DDL must be introduced safely."
---

# Locks and Deadlocks

MVCC lets readers and writers coexist, but conflicting writes and DDL still lock. Claim queue rows without making
workers wait behind each other:

```sql
WITH claimed AS (
    SELECT tenant_id, event_id
    FROM outbox_events
    WHERE published_at IS NULL
    ORDER BY created_at, event_id
    FOR UPDATE SKIP LOCKED
    LIMIT 100
)
SELECT o.*
FROM outbox_events AS o
JOIN claimed AS c USING (tenant_id, event_id);
```

`SKIP LOCKED` is suitable for work queues, not user queries requiring a complete result. Multiple workers may process
different batches, so effects remain idempotent.

```text
T1 locks case 7 ----waits for alert 9
                         ^             |
                         |             v
T2 locks alert 9 --waits for case 7 ---+
```

PostgreSQL detects the cycle and aborts one transaction. Prevent it by locking shared resources in the same canonical
order, touching fewer rows, and shortening transactions. Diagnose blockers:

```sql
SELECT pid, wait_event_type, wait_event, pg_blocking_pids(pid), query
FROM pg_stat_activity
WHERE wait_event IS NOT NULL;
```

Set `lock_timeout` for migrations and `statement_timeout` per workload. Do not “fix” contention by raising timeouts;
that increases queued work and tail latency.
