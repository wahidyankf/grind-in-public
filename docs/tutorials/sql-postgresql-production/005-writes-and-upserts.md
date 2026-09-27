---
tldr: "Implements safe inserts, updates, deletes, idempotency, and optimistic concurrency."
when_to_use: "Use when building mutation paths that must behave correctly under retries and concurrent writers."
---

# Writes and Upserts

Use `RETURNING` to receive database-generated state in the same round trip:

```sql
INSERT INTO cases (tenant_id, title, status)
VALUES (1, 'Review unusual transfer', 'open')
RETURNING tenant_id, case_id, version, created_at;
```

An upsert needs a real uniqueness invariant. Idempotency is not “ignore every conflict”; the same key with a different
payload must be detected.

```sql
INSERT INTO transactions (
    tenant_id, account_id, idempotency_key, amount,
    direction, counterparty, occurred_at
)
VALUES (1, 1, 'n-004', 85.00, 'debit', 'Pine Store', now())
ON CONFLICT (tenant_id, idempotency_key) DO NOTHING
RETURNING transaction_id;
```

If no row returns, query the existing record and compare canonical request fields. Blind `DO UPDATE` can turn a retry
into mutation.

Optimistic concurrency prevents lost updates without holding a lock during user think time:

```sql
UPDATE cases
SET status = 'review', version = version + 1, updated_at = now()
WHERE tenant_id = 1 AND case_id = 1 AND version = 3
RETURNING version;
```

Zero rows means stale input; return a conflict and let the caller reload. For deletes, default to an explicit retention
workflow when records are evidence. A `deleted_at` flag is not free: every query, uniqueness rule, index, purge job, and
foreign key must define its semantics.
