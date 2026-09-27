---
tldr: "Replace deep OFFSET pagination with stable cursor pagination on a changing feed."
when_to_use: "Use to practise API contracts and multicolumn B-tree design."
---

# Case 003: Pagination and Indexes

## Functional requirements

- FR-1: List a tenant's cases newest-first with optional status.
- FR-2: Return a continuation token and no duplicate row during normal forward traversal.
- FR-3: Support jumping to a supplied created-time boundary.

## Non-functional requirements

- NFR-1: p99 below 150 ms at page 10,000 over 200 million rows.
- NFR-2: Inserts continue at 5,000/s.
- NFR-3: Tokens are tamper-evident and tenant-bound.
- NFR-4: Schema/index rollout is online.

Design cursor contents, comparison predicate, index, concurrent changes semantics, rollout, and observability.

[Debrief](../debriefs/003-pagination-and-indexes.md)
