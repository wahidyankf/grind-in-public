---
tldr: "Design the final hybrid migration stage, tenfold growth response, regional recovery, and legacy decommissioning."
when_to_use: "Use as a capstone combining architecture, Kubernetes, data migration, operations, and leadership."
---

# Case 024: Hybrid Cutover, Recovery, and Decommissioning

Several capabilities have been extracted, but the modular monolith remains system of record for cases and customer
profiles. Traffic will grow tenfold, a regional recovery test is due, and leadership wants a defensible end state.

### Functional requirements

| ID   | Requirement                                                                                            |
| ---- | ------------------------------------------------------------------------------------------------------ |
| FR-1 | Map every capability, data writer, read model, event, dependency, owner, and migration state.          |
| FR-2 | Route tenant cohorts through old/new paths with shadow, canary, rollback, and reconciliation controls. |
| FR-3 | Absorb tenfold traffic while protecting synchronous decisions and pausing lower-priority work.         |
| FR-4 | Recover a region, fence writers, restore/replay state, and verify cross-system consistency.            |
| FR-5 | Decommission superseded code, tables, events, flags, permissions, infrastructure, and runbooks.        |
| FR-6 | Preserve a modular monolith for capabilities that lack a justified extraction pressure.                |

### Non-functional requirements

| ID    | Requirement                                                                               |
| ----- | ----------------------------------------------------------------------------------------- |
| NFR-1 | Meet existing SLOs through growth with documented capacity and load-test evidence.        |
| NFR-2 | Premium tenant recovery is RPO 5 minutes/RTO 30 minutes; no split-brain writers.          |
| NFR-3 | Every migration state has an owner, observable entry/exit criteria, and bounded lifetime. |
| NFR-4 | No cutover depends on destructive rollback or unverifiable dual writes.                   |
| NFR-5 | Reduce operational surfaces after migration rather than permanently running both stacks.  |

## Assumptions and exclusions

Budget permits capacity and engineering work, but not an all-at-once rewrite. The answer must prioritize by risk and
business pressure, not produce a calendar estimate.

## Interview prompts

1. Draw current and target authority maps, including intentionally retained monolith modules.
2. Identify the first three tenfold bottlenecks and the load-shedding order.
3. Give an exact regional recovery and reconciliation sequence.
4. Define decommission evidence and the leadership metrics that show migration value.

Solve before reading [the worked solution](../solutions/024-hybrid-cutover-and-decommissioning.md).
