---
tldr: "Design secure investigation workflow, evidence access, collaboration, search, and complete audit."
when_to_use: "Use to practise relational workflow state, RBAC/ABAC, optimistic concurrency, and audit separation."
---

# Case 016: Case Management with Auditable Access

Design a multi-tenant case-management system used by investigators, reviewers, supervisors, and auditors.

### Functional requirements

| ID   | Requirement                                                                                  |
| ---- | -------------------------------------------------------------------------------------------- |
| FR-1 | Create, assign, prioritize, comment on, and transition a case through a configured workflow. |
| FR-2 | Attach immutable evidence references and versioned investigator conclusions.                 |
| FR-3 | Enforce role, tenant, team, assignment, sensitivity, and separation-of-duty policies.        |
| FR-4 | Search cases by permitted attributes and produce saved work queues.                          |
| FR-5 | Record every privileged read, mutation, export, reassignment, and policy override.           |
| FR-6 | Support concurrent users without silently overwriting changes.                               |

### Non-functional requirements

| ID    | Requirement                                                                                      |
| ----- | ------------------------------------------------------------------------------------------------ |
| NFR-1 | Serve ordinary reads within 300 ms p95 and writes within 500 ms p95.                             |
| NFR-2 | Provide 99.9% monthly availability and prevent unauthorized disclosure under dependency failure. |
| NFR-3 | Keep workflow and authorization state strongly consistent; search may lag by 30 seconds.         |
| NFR-4 | Retain case history and evidence lineage for seven years.                                        |
| NFR-5 | Support 10,000 concurrent users and large tenant-specific workload skew.                         |

## Assumptions and exclusions

Evidence blobs live in a separately controlled object store. Real-time collaborative text editing is out of scope.

## Interview prompts

1. Model workflow invariants and optimistic concurrency.
2. Where are authorization policies evaluated and cached?
3. How is the search projection secured against cross-tenant leakage?
4. How do audit, export, retention, legal hold, and evidence access work?

Solve before reading [the worked solution](../solutions/016-case-management-with-auditable-access.md).
