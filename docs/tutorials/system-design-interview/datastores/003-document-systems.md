---
tldr: "Applies document databases to aggregate-shaped data while controlling schema drift, transactions, and indexes."
when_to_use: "Use when most reads and writes address one evolving aggregate as a whole."
---

# Document Systems

A document store fits aggregates whose related fields are usually fetched and updated together: versioned policy
definitions, investigation forms, or external entity profiles. Flexible schema means validation moves into application
and database rules; it does not mean there is no schema.

## Embed or reference

```text
embed                         reference
+----------------------+      policy { rule_ids: [...] }
| policy               |                 |
|  rules: [{...}, ...] |                 +--> rules collection
+----------------------+
atomic aggregate read          shared or independently growing child
```

Embed bounded, owned children needed with the parent. Reference shared entities, independently updated data, or arrays
that can grow without bound. A single oversized document becomes a hot write and rewrite cost.

## Versioning

Keep an immutable published version and a mutable draft rather than editing live policy in place:

```json
{
  "tenant_id": "tenant-7",
  "policy_id": "payments",
  "version": 12,
  "status": "published",
  "schema_version": 3,
  "rules": [{ "kind": "amount_limit", "threshold": "10000.00" }]
}
```

Compound indexes must start with tenant scope for common tenant-local queries. Review index selectivity and write cost.
Multi-document transactions exist in modern document databases, but frequent cross-document transactions suggest the
aggregate boundary or datastore choice is wrong.

## Consistency and change streams

Choose read and write concern from durability and freshness needs. Change streams can feed projections, but consumers
still need checkpoints and idempotency. Treat resume-token expiry and full resynchronization as designed paths.

## Reject when

Reject document storage for graph traversal, analytical scans, or strict invariants across many aggregates. Reject
"schema flexibility" as a reason if ungoverned variants would make every consumer branch on historical shapes.

## References

- [MongoDB manual](https://www.mongodb.com/docs/manual/)
