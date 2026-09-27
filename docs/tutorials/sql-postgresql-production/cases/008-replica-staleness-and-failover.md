---
tldr: "Preserve user consistency and recover safely during replica lag and primary failure."
when_to_use: "Use for replication, read routing, RPO, RTO, and fencing practice."
---

# Case 008: Replica Staleness and Failover

## Functional requirements

- FR-1: A user opening a case must immediately see it on the next request.
- FR-2: Search may be 10 seconds stale and exports 5 minutes stale.
- FR-3: Promote a standby and resume writes after primary failure.

## Non-functional requirements

- NFR-1: RPO is zero for committed case changes in-region.
- NFR-2: RTO is 5 minutes; monthly availability is 99.95%.
- NFR-3: No split-brain writes or silent read regression.
- NFR-4: Failover is rehearsed quarterly.

Design synchronous topology, consistency tokens/routing, lag admission, fencing, promotion, client recovery, and drills.

[Debrief](../debriefs/008-replica-staleness-and-failover.md)
