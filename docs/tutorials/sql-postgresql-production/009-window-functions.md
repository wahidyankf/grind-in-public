---
tldr: "Computes ranks, running metrics, and latest-per-group results without collapsing row detail."
when_to_use: "Use for analytical SQL interviews and production reports that retain individual rows."
---

# Window Functions

Unlike `GROUP BY`, a window function preserves input rows. Rank transactions per account:

```sql
SELECT tenant_id, account_id, transaction_id, amount,
       row_number() OVER (
           PARTITION BY tenant_id, account_id
           ORDER BY amount DESC, transaction_id
       ) AS amount_rank,
       sum(amount) OVER (
           PARTITION BY tenant_id, account_id
           ORDER BY occurred_at, transaction_id
           ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW
       ) AS running_total
FROM transactions;
```

Specify a frame. The default `RANGE` frame groups ordering peers, which can surprise a running total. `ROWS` advances
one physical row at a time. Latest row per group is a reusable interview pattern:

```sql
WITH ranked AS (
    SELECT t.*,
           row_number() OVER (
               PARTITION BY tenant_id, account_id
               ORDER BY occurred_at DESC, transaction_id DESC
           ) AS rn
    FROM transactions AS t
)
SELECT tenant_id, account_id, transaction_id, occurred_at
FROM ranked
WHERE rn = 1;
```

The supporting index is `(tenant_id, account_id, occurred_at DESC, transaction_id DESC)`. A window still may sort the
whole qualifying set; bound the input first and inspect spills.
