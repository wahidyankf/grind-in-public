---
tldr: "Design a multi-tenant control plane and isolated Kubernetes data-plane cells."
when_to_use: "Use to practise cell architecture, configuration rollout, workload identity, and noisy-neighbour control."
---

# Case 019: Kubernetes Control and Data Planes

Design a SaaS platform whose global control plane manages tenant configuration and placement while regional data-plane
cells ingest and decide on tenant traffic.

### Functional requirements

| ID   | Requirement                                                                                       |
| ---- | ------------------------------------------------------------------------------------------------- |
| FR-1 | Provision a tenant into an approved region/cell with quotas, keys, policy, and service endpoints. |
| FR-2 | Publish signed, versioned configuration from control plane to assigned cells.                     |
| FR-3 | Route requests to the current placement and move a tenant through an explicit migration state.    |
| FR-4 | Isolate critical synchronous, asynchronous, and bulk workloads within each cell.                  |
| FR-5 | Report cell health, configuration convergence, capacity, and tenant-specific SLOs.                |

### Non-functional requirements

| ID    | Requirement                                                                                        |
| ----- | -------------------------------------------------------------------------------------------------- |
| NFR-1 | A control-plane outage must not stop a healthy cell using last-known-good configuration.           |
| NFR-2 | A cell failure affects no more than 5% of tenants and no other cell's data plane.                  |
| NFR-3 | Configuration activation completes within 60 seconds and never mixes one request across versions.  |
| NFR-4 | Respect region residency, workload identity, network isolation, and per-tenant encryption context. |
| NFR-5 | Scale from 10 to 100 cells without manual per-cell deployment drift.                               |

## Assumptions and exclusions

Managed regional databases and logs are available. The design need not implement the Kubernetes control plane itself.

## Interview prompts

1. Which responsibilities belong globally, regionally, and inside a cell?
2. How are placement and configuration cached, versioned, and recovered?
3. Which Deployments, Jobs, autoscaling signals, budgets, and topology controls are required?
4. How are fleet rollout, cell evacuation, and tenant move made observable and reversible?

Solve before reading [the worked solution](../solutions/019-kubernetes-control-and-data-planes.md).
