---
tldr: "Selects the latest duplicate event deterministically and designs prevention."
when_to_use: "Use for deduplication and latest-state interview problems."
---

# Drill 006: Deduplication and Latest State

Assume a staging table can contain duplicate `(tenant_id, idempotency_key)` rows. Return the latest by `recorded_at`,
tie-breaking on transaction id. Then state how the production table prevents the duplicate.

[Reference solution](../solutions/006-deduplication-and-latest.md)
