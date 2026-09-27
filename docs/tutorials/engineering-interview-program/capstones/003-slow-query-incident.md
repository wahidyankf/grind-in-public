---
tldr: "Lead a slow-query incident from customer symptom through plan internals and durable prevention."
when_to_use: "Use after performance, planner, and Engineering Manager incident lessons."
---

# Capstone 003: Slow-Query Incident

## Functional requirements

- FR-1: Preserve decision results and case search during mitigation.
- FR-2: Identify affected tenants and queries.
- FR-3: Produce a verified fix and post-incident actions.

## Non-functional requirements

- NFR-1: Restore p99 from 8 seconds to below 250 ms.
- NFR-2: Prevent retry-driven overload and retain operator capacity.
- NFR-3: Deploy a reversible mitigation within 20 minutes.
- NFR-4: No weakened isolation, constraint, or tenant boundary.

Run the incident as Engineering Manager: roles, communication, hypotheses, PostgreSQL evidence, algorithm explanation,
mitigation/rollback, load validation, learning review, and prioritized follow-up.

[Debrief](../debriefs/003-slow-query-incident.md)
