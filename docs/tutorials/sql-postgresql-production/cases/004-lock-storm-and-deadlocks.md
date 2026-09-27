---
tldr: "Recover from lock queues and deadlocks caused by inconsistent update order."
when_to_use: "Use for concurrency incident and transaction-design practice."
---

# Case 004: Lock Storm and Deadlocks

## Functional requirements

- FR-1: Assign an alert and attach it to a case atomically.
- FR-2: Preserve optimistic version checks on the case.
- FR-3: Return a clear conflict for stale user state.

## Non-functional requirements

- NFR-1: p99 below 400 ms at 300 concurrent reviewers.
- NFR-2: Deadlocks below 0.01% of transactions.
- NFR-3: No partial assignments or duplicate case links.
- NFR-4: Incident mitigation must not disable constraints.

Some code locks case then alert; another path locks alert then case. Describe triage, safe mitigation, canonical
ordering, retries, timeouts, and tests.

[Debrief](../debriefs/004-lock-storm-and-deadlocks.md)
