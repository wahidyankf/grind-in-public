---
tldr: "Design a tenant-safe transaction search from API through schema, SQL, indexes, and tests."
when_to_use: "Use after PostgreSQL lessons 001-012 and algorithm lessons on hashing and search."
---

# Capstone 001: Tenant-Safe Data Path

## Functional requirements

- FR-1: Ingest an idempotent transaction for a tenant/account.
- FR-2: Search one tenant by time, direction, amount, and cursor.
- FR-3: Reject cross-tenant account references and reused keys with different payloads.

## Non-functional requirements

- NFR-1: 5,000 writes/s and 2,000 searches/s; search p99 below 150 ms.
- NFR-2: Tenant isolation survives an application predicate bug.
- NFR-3: Audit rejected writes without logging sensitive payloads.
- NFR-4: Schema changes remain compatible with two application versions.

Deliver composite keys/constraints, RLS role model, SQL and Psycopg functions, cursor design, B-tree choice, test
matrix, `EXPLAIN` expectations, deployment, rollback, and ownership. Relate B-tree search and hash idempotency to their
Python algorithm equivalents.

[Debrief](../debriefs/001-tenant-safe-data-path.md)
