---
tldr:
  "Tenant-scoped active-passive authority, fenced epochs, durable replication, and rehearsed reconciliation meet
  recovery goals."
when_to_use: "Use after attempting Case 020 to compare regional strategies and exact recovery sequencing."
---

# Solution 020: Multi-Region Residency and Recovery

## Requirement traceability

| Requirements | Design response                                                                                         |
| ------------ | ------------------------------------------------------------------------------------------------------- |
| FR-1, NFR-2  | Residency-aware placement maps every canonical, derived, backup, key, and telemetry class.              |
| FR-2, NFR-4  | Signed tenant placement cached at routers/cells; conflict or expired authority rejects writes.          |
| FR-3, NFR-3  | Quorum-controlled monotonically increasing writer epoch fences old region before promotion.             |
| FR-4         | Restore/replay canonical inputs, rebuild projections, compare counts/hashes/invariants.                 |
| FR-5, NFR-5  | Automated exercises record detection, decisions, epoch, RPO/RTO, reconciliation, and follow-ups.        |
| NFR-1        | Async cross-region log/database replication within five-minute lag and warm capacity within 30 minutes. |

## Per-data-class strategy

```text
active region                 approved recovery region
API + workers                 warm API/workers
primary DB  -- async -------> replica / continuous backup
durable log -- replicate ---> recovery log
objects     -- encrypted ---> replicated objects
indexes                      rebuilt from canonical sources
```

Use active-passive writes per tenant. It avoids application conflict resolution and wide-area latency while meeting the
stated RPO, not zero RPO. Tenants that forbid another region use in-region multi-zone availability plus approved local
backup; their regional-loss RTO must differ explicitly.

## Failover sequence

```text
1 detect and declare region unavailable
2 stop routing writes; obtain recovery decision quorum
3 increment tenant writer epoch in independent placement authority
4 promote recovery database/log at a recorded durable watermark
5 start data plane with new epoch; reject old-epoch writes
6 smoke-test and gradually route clients
7 replay to watermark; rebuild/validate projections
8 reconcile accepted ids, counts, hashes, sequence gaps, evidence
```

Credentials and database write fences must enforce the epoch; an HTTP header alone cannot stop an isolated old region.
Any events acknowledged after the replicated watermark are inside the declared RPO loss window and recovered from client
retry or source reconciliation when available.

## Failback

Treat failback as another planned migration: seed from current authority, replicate changes, shadow, fence the recovery
writer, advance epoch, and switch. Never simply point DNS back to a stale former primary.

## Operations and alternatives

Measure replication lag, last durable cross-region offset, epoch conflicts, routing convergence, recovery capacity, and
reconciliation gaps. Exercise dependency and human communication failure, not only infrastructure scripts.

Active-active writes are rejected because requirements do not justify conflict semantics and every decision/evidence
record wants one authority. Backup-only is rejected for the 30-minute RTO unless restore tests prove it.
