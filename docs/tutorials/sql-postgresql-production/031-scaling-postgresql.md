---
tldr: "Scales PostgreSQL in a reversible order from measurement and tuning through partitioning, cells, and regions."
when_to_use: "Use for capacity planning and system-design interviews involving relational growth."
---

# Scaling PostgreSQL

Scale in the least complex order that meets measured requirements:

```text
measure
  -> repair queries/statistics/indexes
  -> scale primary CPU/RAM/storage/IOPS
  -> bound/pool connections
  -> cache or precompute named read paths
  -> add replicas for stale-tolerant reads and recovery
  -> isolate analytical/background workloads
  -> partition for pruning/lifecycle
  -> shard or place tenants into cells
  -> add regional topology from RPO/RTO/residency
```

Vertical scaling preserves transactions and operational simplicity. Read replicas do not scale writes and introduce lag
semantics. Caches need keys, freshness, invalidation, stampede control, and fallback capacity. Table partitioning
divides one logical table: it helps pruning, bulk retention, and maintenance, but too many partitions increase planning
and metadata overhead. Every unique constraint on a partitioned table must be enforceable with the partition key.

Sharding moves correctness into the system architecture:

```text
tenant directory -> router -> cell A: app + primary + replicas
                         \-> cell B: app + primary + replicas

global control: placement, schema version, capacity, backup, failover policy
cross-cell analytics: event/log export -> analytical store
```

Prefer tenant/cell placement when most transactions are tenant-local. Define rebalancing, hot tenants, idempotent move
orchestration, schema rollout, global identity, cross-cell reports, and failure isolation before claiming “horizontal
scale.” Distributed transactions are usually a signal to revisit boundaries; sagas trade atomicity for explicit
compensation and intermediate states.

Kubernetes should run stateless application replicas with bounded pools and disruption budgets. A database operator can
automate failover and backups, but persistent-volume topology, fencing, WAL archive, restore, monitoring, and capacity
still need human-owned runbooks. A pod restart is not a database recovery strategy.
