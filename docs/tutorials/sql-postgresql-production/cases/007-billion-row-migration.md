---
tldr: "Add a required derived key to a billion-row table without downtime."
when_to_use: "Use for expand-contract and backfill system-design practice."
---

# Case 007: Billion-Row Migration

## Functional requirements

- FR-1: Add `region_code` derived from immutable tenant placement.
- FR-2: New and old application versions must coexist for two releases.
- FR-3: Resume backfill after interruption and prove completeness.

## Non-functional requirements

- NFR-1: No table rewrite or lock longer than 2 seconds on the write path.
- NFR-2: Backfill uses at most 15% database I/O capacity.
- NFR-3: Replication lag remains below 30 seconds.
- NFR-4: Rollback does not discard accepted writes.

Design schema expansion, dual-write or derivation, chunking, validation, constraint/index rollout, cutover, and cleanup.

[Debrief](../debriefs/007-billion-row-migration.md)
