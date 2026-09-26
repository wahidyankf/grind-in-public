---
tldr: "Design a tamper-evident, queryable, retained ledger for decisions and administrative actions."
when_to_use: "Use to practise append-only storage, hash chains, anchoring, retention, and audit access."
---

# Case 006: Append-Only Evidence Ledger

Design an evidence ledger recording decisions, policy changes, privileged reads, exports, and operator actions. Auditors
must verify integrity and retrieve a tenant/time/resource trail.

### Functional requirements

| ID   | Requirement                                                                                                      |
| ---- | ---------------------------------------------------------------------------------------------------------------- |
| FR-1 | Append a canonical event with actor, authority, tenant, resource, action, versions, outcome, and correlation id. |
| FR-2 | Query by tenant and bounded time plus resource, actor, or action filters.                                        |
| FR-3 | Export a signed manifest and records whose integrity can be verified offline.                                    |
| FR-4 | Detect gaps, mutation, reordering within an integrity segment, and missing batch anchors.                        |
| FR-5 | Enforce retention, legal holds, and approved deletion with its own evidence record.                              |

### Non-functional requirements

| ID    | Requirement                                                                                                |
| ----- | ---------------------------------------------------------------------------------------------------------- |
| NFR-1 | Sustain 30,000 appends/s; acknowledge within 100 ms p99.                                                   |
| NFR-2 | Lose no acknowledged record after a zone failure and recover a region within RPO 5 minutes/RTO 30 minutes. |
| NFR-3 | Retain records for seven years while ordinary queries cover at most 90 days.                               |
| NFR-4 | Separate ledger administration from application and audit-reader roles.                                    |
| NFR-5 | Encrypt tenant data and make bulk export resource usage controllable.                                      |

## Assumptions and exclusions

The ledger proves recorded history has not changed; it cannot prove the original actor told the truth. A public
blockchain is neither required nor assumed.

## Interview prompts

1. What is acknowledged synchronously, and what is batched?
2. How are hash-chain segments partitioned without demanding one global writer?
3. How do hot query indexes remain rebuildable from immutable storage?
4. How do anchoring, key rotation, export verification, legal hold, and deletion interact?

Solve before reading [the worked solution](../solutions/006-append-only-evidence-ledger.md).
