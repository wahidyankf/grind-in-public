---
tldr: "Design time partitioning for retention without confusing it with sharding."
when_to_use: "Use for lifecycle, pruning, and operational partitioning practice."
---

# Case 009: Partition Lifecycle

## Functional requirements

- FR-1: Query one tenant over a time interval and retrieve by event identity.
- FR-2: Keep 90 days online, archive seven years, and delete expired online data.
- FR-3: Accept events up to 48 hours late.

## Non-functional requirements

- NFR-1: Sustain 30,000 inserts/s and 50 billion retained online rows.
- NFR-2: Retention removal must not produce a delete/WAL storm.
- NFR-3: p99 seven-day tenant query below 2 seconds.
- NFR-4: Partition creation failure must alert before ingestion stops.

Choose keys/granularity, indexes, uniqueness semantics, late arrivals, automation, archival, and future sharding
trigger.

[Debrief](../debriefs/009-partition-lifecycle.md)
