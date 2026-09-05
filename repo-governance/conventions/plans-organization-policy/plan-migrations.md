---
tldr: "Makes formal plans preserve, verify, and recover data, configuration, and dependency transitions."
when_to_use:
  "Use when a formal plan moves, copies, normalizes, replaces, or retires data, configuration, APIs, protocols, or
  dependencies."
---

# Plan Migrations

Use `tech-docs/migration-design.md` whenever a formal plan changes a stored, runtime, browser, schema, API, protocol,
configuration, or dependency representation. Add `tech-docs/data-contracts.md` when exact shapes change. Both documents
are mapped from `tech-docs/README.md`; do not create either for work that has no transition.

Inventory every affected source with its safe location or key, readers, writers, accepted shape or version, owner,
destination, compatibility behaviour, and disposition proof. Never place private values, credentials, or user data in
the plan.

## Data-Model Changes

Every plan that adds, changes, or removes a persistent schema places a terminal-readable ASCII data-model diagram beside
the exact old and new contracts. A relational diagram shows the affected entities, primary, foreign, and unique keys,
relationships, cardinalities, and adjacent ownership context. A non-relational diagram shows the equivalent records,
keys, references, and ownership boundaries. The diagram aids comprehension and never replaces exact types, constraints,
defaults, validation, compatibility, migration, rollback, or tests.

The data contract also gives a field-by-field guide for every resulting column, property, or key. Explain its purpose,
producer or owner, value shape, unit or timezone where relevant, required, null, and default behaviour, keys and
references, sensitivity, and creation, update, clearing, and retention lifecycle. Keep the exact schema authoritative;
the guide explains intent that a name or type cannot.

Describe the transition in this order:

1. **Expand** — add the reader, writer, schema, or location without removing the old one.
2. **Migrate** — copy an immutable source through an idempotent identity, validation, and observable outcome.
3. **Verify** — use the normal product flow and rehearse restoration from the verified recovery source.
4. **Contract** — retain compatibility for a stated window; schedule destructive archival or deletion in a separately
   authorized later plan.

For an authority or retirement cutover, verification boots a fresh process from persisted target configuration, makes
the prior source unavailable in an isolated fixture, and exercises every affected critical reader and writer through
normal product boundaries. It proves the target remains authoritative and fails closed instead of falling back. Counts,
schema presence, migration summaries, adapter-only calls, and same-process state are insufficient.

State mixed-version boundaries, retry behaviour, rollback reader and writer behaviour, recovery evidence, and manual
verification. Preserve unknown or malformed input as an opaque record with a reported outcome; never coerce, discard, or
overwrite it to make a migration appear successful. Delivery checkpoints prove each inventoried item is retained, copied
and verified, or safely reported for retry.
