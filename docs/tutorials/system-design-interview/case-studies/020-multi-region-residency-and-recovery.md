---
tldr: "Design regional placement, replication, failover, and recovery under residency and conflict constraints."
when_to_use: "Use to practise RPO/RTO, active-passive versus active-active, DNS/routing, and reconciliation."
---

# Case 020: Multi-Region Residency and Recovery

Design regional deployment for tenants contracted to specific data regions. Some require regional disaster recovery;
others allow only in-region backup locations.

### Functional requirements

| ID   | Requirement                                                                                          |
| ---- | ---------------------------------------------------------------------------------------------------- |
| FR-1 | Place each tenant's write authority, replicas, backups, and derived data only in approved locations. |
| FR-2 | Route clients to their tenant's active region and reject conflicting placement metadata.             |
| FR-3 | Declare a region unavailable, promote an eligible recovery region, and fence the old writer.         |
| FR-4 | Reconcile durable inputs, database state, projections, and evidence after recovery.                  |
| FR-5 | Exercise failover/failback and produce proof of achieved RPO/RTO.                                    |

### Non-functional requirements

| ID    | Requirement                                                                         |
| ----- | ----------------------------------------------------------------------------------- |
| NFR-1 | Premium tenants require RPO 5 minutes and RTO 30 minutes after regional loss.       |
| NFR-2 | Ordinary in-region API latency remains below 200 ms p99.                            |
| NFR-3 | Split brain must not create two accepted write authorities for one tenant.          |
| NFR-4 | Control-plane and DNS/routing dependencies have documented stale/failure behaviour. |
| NFR-5 | Recovery procedures are executable under partial telemetry and staff availability.  |

## Assumptions and exclusions

Not every tenant requires active-active writes. A recovery region is pre-provisioned enough to meet its RTO.

## Interview prompts

1. Compare backup/restore, warm standby, active-passive, and active-active per data class.
2. How is a writer lease/epoch fenced across partitions?
3. Which acknowledgements determine RPO and latency?
4. How are failover, failback, replay, conflict detection, and exercises controlled?

Solve before reading [the worked solution](../solutions/020-multi-region-residency-and-recovery.md).
