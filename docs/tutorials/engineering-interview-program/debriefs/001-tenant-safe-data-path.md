---
tldr: "Combines composite constraints, RLS, idempotency, keyset pagination, and measured rollout."
when_to_use: "Use after attempting Capstone 001."
---

# Debrief 001: Tenant-Safe Data Path

```text
authenticated tenant -> API validation -> transaction
                                         +-> idempotency unique key
                                         +-> composite account FK
                                         +-> RLS tenant check
search cursor -> signed tenant/filter/last-key -> matching B-tree
```

Use `(tenant_id, idempotency_key)` uniqueness and `(tenant_id, account_id)` foreign keys. Set tenant context locally in
the transaction and run the application role under RLS; keep a separate audited administrative role. Search with
`(tenant_id, occurred_at DESC, transaction_id DESC)` keyset pagination. The B-tree provides logarithmic seek plus output
walk; the unique index provides the persistent hash/set-like membership invariant under concurrency.

Tests cover another tenant's account, same/different retry payload, equal timestamps, cursor tampering, null/empty
filters, concurrent inserts, and old/new schema clients. Shadow the query, build indexes concurrently, compare result
hashes and plans, then enable gradually. Roll back application routing without dropping the additive schema.
