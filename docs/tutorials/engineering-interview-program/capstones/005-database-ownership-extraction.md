---
tldr: "Extract a database-owning service from a Python modular monolith without a big-bang cutover."
when_to_use: "Use after migration, messaging, and modular-monolith extraction lessons."
---

# Capstone 005: Database Ownership Extraction

## Functional requirements

- FR-1: Extract alert ownership while case workflows continue.
- FR-2: Preserve existing API behaviour and historical queries during transition.
- FR-3: Move writes and data with reconciliation and rollback.
- FR-4: Decommission old write paths and tables after proof.

## Non-functional requirements

- NFR-1: No planned downtime; p99 degradation below 10%.
- NFR-2: No distributed transaction between old and new services.
- NFR-3: Data mismatch below 0.001% before cutover and zero unexplained mismatch.
- NFR-4: Migration remains supportable by two teams with clear ownership.

Present readiness criteria, target boundary, contract, data movement, change capture/outbox, shadowing, cutover state
machine, Kubernetes rollout, rollback, observability, team topology, and stop conditions.

[Debrief](../debriefs/005-database-ownership-extraction.md)
