---
tldr: "Uses row_number for cleanup and a unique constraint for prevention."
when_to_use: "Use after Drill 006."
---

# Solution 006: Deduplication and Latest State

```sql
WITH ranked AS (
    SELECT s.*, row_number() OVER (
        PARTITION BY tenant_id, idempotency_key
        ORDER BY recorded_at DESC, transaction_id DESC
    ) AS rn
    FROM transaction_staging AS s
)
SELECT * FROM ranked WHERE rn = 1;
```

The production invariant is `UNIQUE (tenant_id, idempotency_key)`. Cleanup SQL repairs history; it is not concurrency
control.
