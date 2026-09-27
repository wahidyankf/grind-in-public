---
tldr: "Builds precise SELECT queries and connects predicates, ordering, pagination, and indexes."
when_to_use: "Use to learn retrieval syntax and avoid ambiguous or unstable production queries."
---

# Select, Filter, and Order

SQL's logical evaluation order differs from its written order:

```text
FROM -> JOIN -> WHERE -> GROUP BY -> HAVING -> SELECT -> DISTINCT -> ORDER BY -> LIMIT
```

Retrieve recent large debits for one tenant:

```sql
SELECT transaction_id, account_id, amount, occurred_at
FROM transactions
WHERE tenant_id = 1
  AND direction = 'debit'
  AND amount >= 1000
ORDER BY occurred_at DESC, transaction_id DESC
LIMIT 50;
```

The identity tie-breaker makes pagination deterministic. Avoid `SELECT *` across service boundaries: schema evolution
then changes response shape and may fetch wide TOAST values unnecessarily. Avoid `OFFSET` for deep pages; PostgreSQL
must still visit and discard earlier rows, and concurrent writes shift positions. Use a cursor:

```sql
SELECT transaction_id, amount, occurred_at
FROM transactions
WHERE tenant_id = 1
  AND (occurred_at, transaction_id) < ('2026-01-10T09:05:00Z', 2)
ORDER BY occurred_at DESC, transaction_id DESC
LIMIT 50;
```

The production index follows equality columns, then range/order columns:

```sql
CREATE INDEX transactions_tenant_occurred_idx
    ON transactions (tenant_id, occurred_at DESC, transaction_id DESC)
    INCLUDE (account_id, amount, direction);
```

`INCLUDE` can enable index-only scans, but visibility-map state still determines whether heap checks are needed. Every
extra column increases storage and write cost. Verify the actual access pattern before creating it.
