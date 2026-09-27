---
tldr: "Design an auditable alert and case platform with transactional workflow and immutable evidence."
when_to_use: "Use after schema, evidence, case-management, and Kubernetes system-design lessons."
---

# Capstone 004: Alert and Case Platform

## Functional requirements

- FR-1: Create scored alerts referencing exact policy/model versions.
- FR-2: Correlate alerts into cases, assign reviewers, and record every transition.
- FR-3: Search queues and export a complete evidence package.
- FR-4: Retain immutable evidence and replay derived projections.

## Non-functional requirements

- NFR-1: 20,000 alerts/s peak; case operations p99 below 300 ms.
- NFR-2: Tenant-scoped authorization, encryption, residency, and audit.
- NFR-3: 99.95% availability; no acknowledged workflow loss.
- NFR-4: Seven-year evidence retention and four-hour export deadline.

Deliver full interview sequence: estimates, APIs/data, architecture, algorithms, database/indexes, consistency,
Kubernetes, security, observability, failures, capacity evolution, and Engineering Manager delivery plan.

[Debrief](../debriefs/004-alert-case-platform.md)
