---
tldr: "A PostgreSQL 18 course from SQL fundamentals through internals, Python, performance, operations, and scaling."
when_to_use: "Use after Python foundations to become interview-ready and production-capable with relational data."
---

# SQL and PostgreSQL Production

This course uses PostgreSQL 18.6 and one fictional multi-tenant investigation platform. It teaches each feature from
four angles: the query you write, the mechanism underneath it, the production problem it solves, and the condition under
which it is the wrong choice. Type and run the examples; this repository contains documentation, not a hidden
application.

## Quick start

```text
client -> port 54329 -> PostgreSQL 18.6 container -> investigation database
```

Create a disposable database outside this repository:

```bash
docker run --name pg-study --rm \
  -e POSTGRES_PASSWORD=study \
  -e POSTGRES_DB=investigation \
  -p 127.0.0.1:54329:5432 \
  -d postgres:18.6-alpine

docker exec -it pg-study psql -U postgres -d investigation
```

Set a shell-local URL for later lessons. Do not commit credentials.

```bash
export STUDY_DATABASE_URL='postgresql://postgres:study@127.0.0.1:54329/investigation'
```

## Reading order

1. SQL: [setup](001-setup-and-study-schema.md), [relational thinking](002-relational-thinking.md),
   [queries](003-select-filter-and-order.md), [DDL](004-types-constraints-and-ddl.md),
   [writes](005-writes-and-upserts.md), [joins](006-joins-and-set-operations.md), [aggregation](007-aggregation.md),
   [CTEs](008-subqueries-and-ctes.md), [windows](009-window-functions.md), and [recursion](010-recursive-ctes.md).
2. Correctness: [normalization](011-modelling-and-normalization.md),
   [multi-tenancy](012-multi-tenant-data-modelling.md), [transactions](013-transactions-and-acid.md),
   [isolation](014-isolation-and-anomalies.md), and [locks](015-locks-and-deadlocks.md).
3. Internals: [query pipeline](016-query-pipeline.md), [storage](017-storage-pages-and-toast.md),
   [buffers](018-buffer-cache-and-io.md), [B-trees](019-btree-indexes.md),
   [specialized indexes](020-specialized-indexes.md), [joins](021-join-algorithms.md),
   [sorts](022-sort-and-aggregation-algorithms.md), [planning](023-planner-statistics-and-costs.md),
   [MVCC](024-mvcc-and-hot.md), [vacuum](025-vacuum-and-freezing.md), [WAL](026-wal-and-recovery.md), and
   [replication](027-replication-and-failover.md).
4. Production: [performance](028-performance-workflow.md), [Psycopg](029-python-and-psycopg.md),
   [operations](030-migrations-security-and-backups.md), [scaling](031-scaling-postgresql.md), and
   [synthesis](032-production-and-interview-checklist.md). Use the ordered
   [primary references](033-primary-references.md) to verify mechanisms at their versioned source.
5. Practise with the [internals labs](labs/README.md) and [lab solutions](lab-solutions/README.md),
   [SQL drills](drills/README.md) and [SQL solutions](solutions/README.md), then [production cases](cases/README.md) and
   [case debriefs](debriefs/README.md). Finish with the
   [cross-course capstones](../engineering-interview-program/README.md).

## Learning contract

Before opening a solution, state the access pattern, invariant, expected cardinality, failure behaviour, and complexity.
Use `EXPLAIN (ANALYZE, BUFFERS)` only on safe test data: `ANALYZE` executes the statement. Timings from a laptop teach
method, not production capacity.

## Where this course fits

Prerequisite: [Python Production Foundations](../python-production-foundations/README.md). Pair internals with
[Python Algorithms Interview](../python-algorithms-interview/README.md), architecture with
[System Design Interview](../system-design-interview/README.md), and operational decisions with
[Engineering Manager Interview](../engineering-manager-interview/README.md). The
[Engineering Interview Programme](../engineering-interview-program/README.md) connects all five courses.

## Primary references

- [PostgreSQL 18 documentation](https://www.postgresql.org/docs/18/)
- [PostgreSQL release and versioning policy](https://www.postgresql.org/support/versioning/)
- [Psycopg 3 documentation](https://www.psycopg.org/psycopg3/docs/)
