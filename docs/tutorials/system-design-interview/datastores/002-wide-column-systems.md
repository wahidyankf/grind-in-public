---
tldr: "Models wide-column data from known partition-local queries, explaining hot keys, compaction, and consistency."
when_to_use: "Use for very large, write-heavy datasets with predictable key-and-time access patterns."
---

# Wide-Column Systems

Wide-column systems distribute rows by a partition key and sort within a partition. They trade flexible joins and
cross-partition transactions for predictable horizontal scale when queries are known in advance.

## Query-first modelling

For "read an account's events for a day in time order":

```text
partition key: (tenant_id, account_id, event_day)
clustering:    event_time DESC, event_id

(T1, A9, 2026-09-26)
  12:04 e103
  12:01 e102
  11:58 e101
```

The day bucket bounds partition size. The account preserves locality and order. A query across all tenants needs a
separate table or analytical path; server-side filtering across the cluster defeats the model.

## Cassandra-style trade-offs

Replication factor and per-operation consistency level determine how many replicas participate. Quorum reads and writes
can provide strong overlap within one healthy region, but repair, hinted handoff, tombstones, and clock/order semantics
still matter. Lightweight transactions coordinate contention but cost more; do not use them as a universal replacement
for relational transactions.

## HBase-style trade-offs

Rows are lexicographically ordered by row key in regions. This supports range scans but sequential prefixes can hotspot
one region. Salt or reverse a naturally increasing prefix when the access pattern allows it. Single-row operations are
the natural atomic boundary; design related values into the row when that invariant matters.

## Operating costs

Compaction rewrites data, tombstones make deletion deferred, repairs reconcile replicas, and unbounded partitions hurt
latency. Capacity planning must include these background activities, not only foreground requests.

## Reject when

Reject a wide-column store for ad hoc joins, rapidly changing query shapes, or multi-entity invariants. Also reject it
when expected scale fits one relational system and the team lacks operational expertise; theoretical scale is not free.

## References

- [Apache Cassandra documentation](https://cassandra.apache.org/doc/latest/)
- [Apache HBase Reference Guide](https://hbase.apache.org/book.html)
