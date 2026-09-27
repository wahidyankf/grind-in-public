---
tldr: "Forms islands by subtracting row_number from each distinct date."
when_to_use: "Use after Drill 005."
---

# Solution 005: Gaps and Islands

```sql
WITH days AS (
    SELECT DISTINCT tenant_id, account_id, occurred_at::date AS active_day
    FROM transactions
), labelled AS (
    SELECT *, active_day - row_number() OVER (
        PARTITION BY tenant_id, account_id ORDER BY active_day
    )::integer AS island
    FROM days
)
SELECT tenant_id, account_id, min(active_day), max(active_day), count(*) AS active_days
FROM labelled
GROUP BY tenant_id, account_id, island
ORDER BY tenant_id, account_id, min(active_day);
```

Consecutive dates share the same shifted value, which becomes the grouping invariant.
