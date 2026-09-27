---
tldr: "Provides a measurement-first workflow for query, connection, memory, I/O, WAL, and contention performance."
when_to_use: "Use for slow-query diagnosis, load testing, and production performance interviews."
---

# Performance Workflow

Performance is a workload property, not a configuration checklist. Begin with the service-level symptom and preserve a
before/after trace.

```text
slow request
   |
   +-> waiting for pool? -> connection saturation
   +-> waiting on lock? -> blocker / transaction design
   +-> executing CPU?   -> rows, joins, expressions, JIT
   +-> reading storage? -> access path, cache, random/sequential I/O
   +-> spilling temp?   -> cardinality, sort/hash memory
   +-> writing WAL?     -> indexes, updates, checkpoints, replication
```

Use this order:

1. capture endpoint, tenant, query identity, latency percentile, concurrency, and change window;
2. inspect pool wait and `pg_stat_activity` waits;
3. use `pg_stat_statements` for total time, mean, calls, rows, and I/O concentration;
4. reproduce with safe parameter classes and `EXPLAIN (ANALYZE, BUFFERS, WAL)`;
5. fix the first estimate or row-amplification error;
6. change query/index/schema or workload before broad configuration;
7. load-test p50/p95/p99, resource saturation, and regression queries;
8. deploy observably with a rollback condition.

The following view requires `pg_stat_statements` to be preloaded by the server and installed with
`CREATE EXTENSION pg_stat_statements`; use the managed-service equivalent when configuration is provider-owned.

```sql
SELECT queryid, calls, total_exec_time, mean_exec_time, rows,
       shared_blks_hit, shared_blks_read, temp_blks_written
FROM pg_stat_statements
ORDER BY total_exec_time DESC
LIMIT 20;
```

Reset statistics only through an agreed operational procedure. A faster average can hide worse p99 from locks or
checkpoints. Throughput without bounded queueing is not capacity: test overload, cancellations, and recovery after the
load stops.
