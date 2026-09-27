---
tldr: "Compares physical and logical replication, synchronous durability, lag, failover, and split-brain fencing."
when_to_use: "Use when designing read scale, upgrades, recovery, or regional database topology."
---

# Replication and Failover

Physical streaming replication sends WAL changes for the cluster; logical replication sends decoded table changes.

```text
clients -> writable primary -> WAL -> standby A -> read traffic
                    |          |
                    |          +----> standby B -> recovery candidate
                    +-> archive -> object storage -> point-in-time recovery

logical publication -> subscriber tables for migration or selective projection
```

Asynchronous standbys protect availability and scale stale-tolerant reads but can lose the latest acknowledged writes if
promoted. Synchronous replication can wait for selected standbys, reducing the loss window while adding network and
standby latency to commit. State the RPO and RTO before choosing.

Failover is a distributed-systems protocol, not “start the replica.” It must:

1. establish that the old primary cannot keep accepting writes;
2. fence it through infrastructure, credentials, routing, or storage ownership;
3. select and promote an eligible standby based on timeline and replay position;
4. route clients and invalidate pools;
5. rebuild former members instead of letting divergent histories rejoin;
6. verify data, lag, jobs, backups, and write capability.

Replication slots retain WAL until consumers advance and can fill disk. Logical replication does not automatically
reproduce every DDL operation or sequence state; migrations need coordinated rollout. Read-after-write paths stay on the
primary or wait for a required replay LSN.

Related learning:
[distributed correctness](../system-design-interview/foundations/005-scaling-and-distributed-correctness.md) and
[multi-region design](../system-design-interview/case-studies/020-multi-region-residency-and-recovery.md).
