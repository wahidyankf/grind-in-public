---
tldr: "Matches isolation levels and retry loops to dirty reads, non-repeatable reads, write skew, and serialization."
when_to_use: "Use when concurrent transactions can violate a business invariant."
---

# Isolation and Anomalies

PostgreSQL maps the SQL levels to three distinct behaviours; `READ UNCOMMITTED` behaves as `READ COMMITTED`.

| Level           | Snapshot                                             | Main production concern                |
| --------------- | ---------------------------------------------------- | -------------------------------------- |
| Read committed  | New snapshot per statement                           | Non-repeatable reads; check/write race |
| Repeatable read | Stable transaction snapshot                          | Serialization failures; write skew     |
| Serializable    | Serializable snapshot isolation with conflict checks | Abort and retry whole transaction      |

Two reviewers can each observe “another reviewer remains” and both go off duty: different rows are updated, but the
cross-row invariant breaks. Lock a materialized invariant row or use serializable isolation with bounded retries.

```sql
BEGIN TRANSACTION ISOLATION LEVEL SERIALIZABLE;
-- Read all rows that establish the invariant, then perform the write.
COMMIT;
```

Retry SQLSTATE `40001` and deadlock `40P01` around the whole transaction, with jitter and a finite attempt limit. The
body must be safe to repeat; do not send email or call another service inside it. Isolation does not solve stale replica
reads, external side effects, or missing constraints.
