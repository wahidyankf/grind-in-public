---
tldr: "Maps every course section to PostgreSQL 18, Psycopg, Kubernetes, and PostgreSQL source-tree references."
when_to_use:
  "Use to verify a mechanism at its primary source or deepen one lesson without reading documentation randomly."
---

# Primary References

Use the course to build a mental model, then verify the exact guarantee in the versioned source. Read one reference to
answer one question and summarize: workload, mechanism, guarantee, cost, and rejection condition.

## SQL, schema, and concurrency

- [SQL language](https://www.postgresql.org/docs/18/sql.html) is the canonical statement reference.
- [Data definition](https://www.postgresql.org/docs/18/ddl.html) covers types, constraints, generated identity,
  privileges, row security, inheritance, and partitioning.
- [Concurrency control](https://www.postgresql.org/docs/18/mvcc.html) covers isolation, locking, serialization failures,
  and caveats.
- [Explicit locking](https://www.postgresql.org/docs/18/explicit-locking.html) defines table/row lock conflicts and
  deadlock behaviour.

## Plans, indexes, and physical storage

- [Using `EXPLAIN`](https://www.postgresql.org/docs/18/using-explain.html) defines plan nodes, estimates, and runtime
  instrumentation.
- [Planner statistics](https://www.postgresql.org/docs/18/planner-stats.html) explains histograms, common values,
  extended statistics, and statistics targets.
- [Index types](https://www.postgresql.org/docs/18/indexes-types.html),
  [multicolumn indexes](https://www.postgresql.org/docs/18/indexes-multicolumn.html), and
  [index-only scans](https://www.postgresql.org/docs/18/indexes-index-only-scans.html) document operator support and
  limitations.
- [Page layout](https://www.postgresql.org/docs/18/storage-page-layout.html),
  [TOAST](https://www.postgresql.org/docs/18/storage-toast.html), and
  [visibility map](https://www.postgresql.org/docs/18/storage-vm.html) connect tuple versions to storage and index-only
  behaviour.
- [Routine vacuuming](https://www.postgresql.org/docs/18/routine-vacuuming.html) covers reuse, freezing, statistics, and
  autovacuum.

## Durability, replication, and operations

- [WAL introduction](https://www.postgresql.org/docs/18/wal-intro.html) and
  [checkpoint configuration](https://www.postgresql.org/docs/18/wal-configuration.html) define recovery ordering and
  checkpoint trade-offs.
- [Continuous archiving and point-in-time recovery](https://www.postgresql.org/docs/18/continuous-archiving.html)
  defines the base-backup plus WAL recovery chain.
- [High availability and load balancing](https://www.postgresql.org/docs/18/high-availability.html) covers physical
  replication, synchronous modes, failover, and read-only standbys.
- [Logical replication](https://www.postgresql.org/docs/18/logical-replication.html) covers publication/subscription
  semantics and restrictions.
- [Monitoring statistics](https://www.postgresql.org/docs/18/monitoring-stats.html) and
  [`pg_stat_statements`](https://www.postgresql.org/docs/18/pgstatstatements.html) define the views used by performance
  lessons.
- [PostgreSQL versioning policy](https://www.postgresql.org/support/versioning/) explains supported major/minor updates.

## Application and deployment boundaries

- [Psycopg 3](https://www.psycopg.org/psycopg3/docs/) documents parameter binding, transaction contexts, row factories,
  COPY, async use, and pools.
- [Kubernetes workload concepts](https://kubernetes.io/docs/concepts/workloads/) and
  [pod lifecycle](https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/) define the orchestration boundary;
  they do not replace database fencing, backup, or restore design.

## Source-aware deep dives

These source-tree READMEs explain implementation intent without requiring line-by-line C study:

- [B-tree implementation](https://github.com/postgres/postgres/blob/REL_18_STABLE/src/backend/access/nbtree/README)
- [HOT updates](https://github.com/postgres/postgres/blob/REL_18_STABLE/src/backend/access/heap/README.HOT)
- [Planner architecture](https://github.com/postgres/postgres/blob/REL_18_STABLE/src/backend/optimizer/README)
- [Executor architecture](https://github.com/postgres/postgres/blob/REL_18_STABLE/src/backend/executor/README)
- [Buffer manager](https://github.com/postgres/postgres/blob/REL_18_STABLE/src/backend/storage/buffer/README)
- [Transaction and WAL implementation][transaction-and-wal-implementation]

Treat source details as version-specific. Carry the invariant and trade-off into an interview; verify implementation
details again before making an operational decision on another major version.

[transaction-and-wal-implementation]:
  https://github.com/postgres/postgres/blob/REL_18_STABLE/src/backend/access/transam/README
