---
tldr: "Uses lag and a time-based RANGE frame."
when_to_use: "Use after Drill 004."
---

# Solution 004: Window Functions

```sql
SELECT tenant_id, account_id, transaction_id, occurred_at, amount,
       lag(amount) OVER account_order AS previous_amount,
       sum(amount) FILTER (WHERE direction = 'debit') OVER (
           PARTITION BY tenant_id, account_id
           ORDER BY occurred_at
           RANGE BETWEEN interval '7 days' PRECEDING AND CURRENT ROW
       ) AS debit_7d
FROM transactions
WINDOW account_order AS (
    PARTITION BY tenant_id, account_id ORDER BY occurred_at, transaction_id
);
```

`RANGE` represents elapsed time and includes all peers at the same timestamp. Add explicit semantics if duplicate event
times require ingestion-order cutoffs.
