---
tldr: "Fix a tenant-skewed join whose estimated row count is wrong by two orders of magnitude."
when_to_use: "Use to practise statistics, skew, and plan-choice reasoning."
---

# Case 002: Cardinality Misestimation

## Functional requirements

- FR-1: Produce the same account-risk report for every tenant.
- FR-2: Join accounts, transactions, and alerts over a configurable interval.
- FR-3: Permit export without blocking online writes.

## Non-functional requirements

- NFR-1: Interactive reports finish within 5 seconds p95 for 99% of tenants.
- NFR-2: One tenant owns 45% of rows and must not exhaust shared resources.
- NFR-3: Report data may be 60 seconds stale.
- NFR-4: Avoid per-tenant hand-authored SQL.

Explain how to identify the first estimate divergence, model correlations, choose custom/generic plans, isolate the
largest tenant if needed, and test representative distributions.

[Debrief](../debriefs/002-cardinality-misestimation.md)
