---
tldr: "Scale a platform tenfold and recover from regional failure with explicit database and team trade-offs."
when_to_use: "Use as the final integrated system-design and Engineering Manager rehearsal."
---

# Capstone 006: Scale and Recover

## Functional requirements

- FR-1: Ingest events, decide, investigate cases, and export evidence.
- FR-2: Move tenants between cells and preserve stable identities.
- FR-3: Recover a failed region and reconcile delayed asynchronous work.

## Non-functional requirements

- NFR-1: 10x traffic/data; online p99 below 300 ms and 99.95% availability.
- NFR-2: In-region RPO zero/RTO 5 minutes; regional RPO 5 minutes/RTO 60 minutes.
- NFR-3: Data residency and tenant isolation remain enforceable.
- NFR-4: One cell or team failure affects at most 5% of tenants.

Produce capacity estimates, staged evolution, cell architecture, PostgreSQL/log/object roles, Kubernetes failure model,
regional recovery sequence, consistency, security, cost, SLOs, roadmap, staffing, incident command, and rehearsal plan.

[Debrief](../debriefs/006-scale-and-recover.md)
