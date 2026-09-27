---
tldr: "Prevent a Kubernetes rollout from multiplying PostgreSQL connections beyond capacity."
when_to_use: "Use for application, platform, and database capacity interviews."
---

# Case 006: Kubernetes Connection Storm

## Functional requirements

- FR-1: Serve case and alert APIs throughout rolling deployment.
- FR-2: Run migrations once, before incompatible code starts.
- FR-3: Keep operator access available during incidents.

## Non-functional requirements

- NFR-1: PostgreSQL permits 600 sessions; reserve 60 for operations/system roles.
- NFR-2: Scale from 20 to 80 pods without connection refusal.
- NFR-3: Absorb a full-zone pod reschedule without retry amplification.
- NFR-4: API availability is 99.95% and p99 below 500 ms.

Each pod has four processes and a pool max of 10. Design budgets, rollout controls, proxy use, probes, backoff, and
monitoring.

[Debrief](../debriefs/006-kubernetes-connection-storm.md)
