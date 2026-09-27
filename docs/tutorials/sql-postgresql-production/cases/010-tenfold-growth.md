---
tldr: "Evolve a multi-tenant PostgreSQL platform through tenfold growth with explicit triggers."
when_to_use: "Use for end-to-end database scaling interviews."
---

# Case 010: Tenfold Growth

## Functional requirements

- FR-1: Preserve case workflow transactions and tenant-scoped search.
- FR-2: Support event ingestion, audit retention, exports, and policy configuration.
- FR-3: Move an exceptional tenant without changing customer-facing identity.

## Non-functional requirements

- NFR-1: Grow from 3,000 to 30,000 writes/s and 1 to 10 PB historical evidence.
- NFR-2: Online p99 below 300 ms and availability 99.95%.
- NFR-3: In-region RPO zero/RTO 5 minutes; regional RPO 5 minutes/RTO 60 minutes.
- NFR-4: Limit one cell failure to at most 5% of tenants.

Present the current design, measurement, ordered scaling stages, cell placement/moves, analytical history, regional
recovery, Kubernetes responsibilities, and organizational ownership.

[Debrief](../debriefs/010-tenfold-growth.md)
