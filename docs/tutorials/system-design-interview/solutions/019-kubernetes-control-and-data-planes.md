---
tldr:
  "A small global control plane assigns tenants to independently operable regional cells using signed versioned state."
when_to_use: "Use after attempting Case 019 to compare cells, fleet delivery, Kubernetes controls, and isolation."
---

# Solution 019: Kubernetes Control and Data Planes

## Requirement traceability

| Requirements       | Design response                                                                                    |
| ------------------ | -------------------------------------------------------------------------------------------------- |
| FR-1, FR-3         | Authoritative placement directory uses versioned state machine and one writer epoch per tenant.    |
| FR-2, NFR-1, NFR-3 | Signed immutable configuration bundles cached in cells; request pins one version.                  |
| FR-4, NFR-4        | Separate workload classes, service accounts, network policies, quotas, keys, and tenant admission. |
| FR-5               | Cell/tenant SLO, capacity, configuration convergence, and placement telemetry.                     |
| NFR-2              | Cell contains no more than 5% of tenant blast radius and avoids synchronous cross-cell data calls. |
| NFR-5              | Declarative fleet templates, progressive waves, conformance checks, and drift detection.           |

## Architecture

```text
                   global control plane
         tenant directory | config signer | fleet status
                    /             |             \
              region A         region B        recovery metadata
            +---------+       +---------+
router ---> | cell A1 |       | cell B1 | <--- router
            | API/log |       | API/log |
            | workers |       | workers |
            | data    |       | data    |
            +---------+       +---------+
```

Cells do not call the control plane on the request path. Routers and cells cache signed placements/configurations with
versions and bounded validity. A stale but valid cell continues safely; provisioning and changes pause.

## Kubernetes mapping

- API and independent worker Deployments, spread across zones;
- Jobs/CronJobs for bounded migration, reconciliation, and maintenance;
- HPA on concurrency or queue age, with warm headroom and database connection caps;
- priority classes for synchronous decision, durable ingestion, projection, then bulk;
- Pod disruption budgets and topology spread for critical replicas;
- per-workload service accounts, network policy, secret references, namespace/resource quotas;
- startup/readiness/liveness aligned with application semantics.

Do not use namespaces alone as strong tenant isolation. High-risk/outlier tenants receive dedicated data placement or
cells; ordinary tenants share a cell with application-level and data-level scope.

## Tenant move

```text
STABLE(A) -> COPYING(A->B) -> DUAL_READ/SHADOW -> FENCED_A -> ACTIVE_B -> CLEANUP_A
```

Placement version and writer epoch travel with every write. Backfill and continuous changes converge; shadow reads are
compared; then the directory fences A before enabling B. Rollback before authority transfer returns routing to A.

## Fleet operations

Release one test cell, then small regional waves with automatic SLO/domain gates. Conformance checks verify image
digest, configuration, policy, quotas, and required observability. Capacity reserves allow evacuation of one cell
without overloading neighbours.

## Alternatives rejected

One global data plane creates wide blast radius and residency coupling. One cluster per small tenant creates an
unmanageable fleet. Request-time control-plane calls violate failure independence.
