---
tldr: "Design ordered screening of transfer parties before execution with bounded external dependencies and repair."
when_to_use: "Use to practise workflow states, entity matching, policy snapshots, and fail-safe semantics."
---

# Case 008: Pre-Execution Transfer Screening

Design a system that screens originator, beneficiary, institutions, countries, and free-text remittance data before a
transfer may execute.

### Functional requirements

| ID   | Requirement                                                                                     |
| ---- | ----------------------------------------------------------------------------------------------- |
| FR-1 | Accept an idempotent screening request with all parties and transfer context.                   |
| FR-2 | Normalize and match parties against the tenant's active reference-list versions.                |
| FR-3 | Apply corridor, amount, party, and exception policy and return clear/review/block.              |
| FR-4 | Place uncertain or configured matches into human review with the evidence snapshot.             |
| FR-5 | Support rescreening a pending transfer after list or policy change without duplicate execution. |

### Non-functional requirements

| ID    | Requirement                                                                                            |
| ----- | ------------------------------------------------------------------------------------------------------ |
| NFR-1 | Return an automatic result within 500 ms p99 at 2,000 requests/s.                                      |
| NFR-2 | Never return clear when required list data or policy is unavailable or unverifiable.                   |
| NFR-3 | Preserve the exact input, normalization, candidates, versions, thresholds, and result for seven years. |
| NFR-4 | Maintain per-transfer state consistency across retry, timeout, review, and rescreening.                |
| NFR-5 | Meet 99.95% monthly availability; degraded paths may return pending/review.                            |

## Assumptions and exclusions

Actual transfer execution is owned elsewhere and requires a single-use clearance token. Cross-organization protocol
standardization is out of scope.

## Interview prompts

1. Model the transfer-screening state machine.
2. How are candidate generation and exact scoring kept inside 500 ms?
3. How does a clearance token prevent stale or replayed execution?
4. What happens during list activation, policy change, or matcher outage?

Solve before reading [the worked solution](../solutions/008-pre-execution-transfer-screening.md).
