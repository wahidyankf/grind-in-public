---
tldr: "Implement a complete idempotent Python transaction and reliable outbox publisher design."
when_to_use: "Use after transactions, Psycopg, queues, and messaging consistency lessons."
---

# Capstone 002: Idempotent Outbox

## Functional requirements

- FR-1: Create one case and one publish intent atomically.
- FR-2: Return the original result on an identical request retry.
- FR-3: Publish events at least once and process each consumer effect once logically.
- FR-4: Quarantine poison events with repair and replay.

## Non-functional requirements

- NFR-1: API p99 below 200 ms at 1,000 requests/s.
- NFR-2: No acknowledged case is missing its event after crash/failover.
- NFR-3: Publisher backlog normally remains below 30 seconds.
- NFR-4: Workers scale in Kubernetes without duplicate side effects.

Provide complete Psycopg transaction and worker slices, schema constraints, `SKIP LOCKED` reasoning, retries, ordering,
metrics, failure sequences, tests, deployment, and team ownership.

[Debrief](../debriefs/002-idempotent-outbox.md)
