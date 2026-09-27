---
tldr: "Uses row_number for an exact row limit per account."
when_to_use: "Use after Drill 003."
---

# Solution 003: Ranking and Top-N

```sql
WITH ranked AS (
    SELECT t.*,
           row_number() OVER (
               PARTITION BY tenant_id, account_id
               ORDER BY amount DESC, transaction_id
           ) AS rn
    FROM transactions AS t
)
SELECT tenant_id, account_id, transaction_id, amount
FROM ranked WHERE rn <= 3
ORDER BY tenant_id, account_id, rn;
```

Use `dense_rank` only when all tied values must survive, because it can return more than three rows.
