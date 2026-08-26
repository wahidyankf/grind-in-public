---
tldr: "Makes formal plans preserve, verify, and recover data, configuration, and dependency transitions."
when_to_use: "Use when a formal plan moves, copies, normalizes, replaces, or retires data, configuration, APIs, protocols, or dependencies."
---

# Plan Migrations

Use `tech-docs/migration-design.md` whenever a formal plan changes a stored, runtime, browser, schema, API, protocol, configuration, or dependency representation. Add `tech-docs/data-contracts.md` when exact shapes change. Both documents are mapped from `tech-docs/README.md`; do not create either for work that has no transition.

Inventory every affected source with its safe location or key, readers, writers, accepted shape or version, owner, destination, compatibility behavior, and disposition proof. Never place private values, credentials, or user data in the plan.

Describe the transition in this order:

1. **Expand** — add the reader, writer, schema, or location without removing the old one.
2. **Migrate** — copy an immutable source through an idempotent identity, validation, and observable outcome.
3. **Verify** — use the normal product flow and rehearse restoration from the verified recovery source.
4. **Contract** — retain compatibility for a stated window; schedule destructive archival or deletion in a separately authorized later plan.

State mixed-version boundaries, retry behavior, rollback reader and writer behavior, recovery evidence, and manual verification. Preserve unknown or malformed input as an opaque record with a reported outcome; never coerce, discard, or overwrite it to make a migration appear successful. Delivery checkpoints prove each inventoried item is retained, copied and verified, or safely reported for retry.
