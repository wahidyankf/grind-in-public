---
tldr: "Explains in-memory sort, top-N heaps, external merge sort, hash aggregation, and spill behaviour."
when_to_use: "Use when ORDER BY, DISTINCT, windows, or grouping dominate a plan."
---

# Sort and Aggregation Algorithms

PostgreSQL chooses among sorting strategies from required output and memory:

```text
full order, fits memory -> quicksort             O(n log n)
ORDER BY ... LIMIT k  -> bounded top-N heap      O(n log k)
does not fit memory   -> sorted disk runs + merge O(n log n), storage-bound
```

A top-N heap keeps only the best `k` rows. External merge sort creates sorted runs that fit memory, writes them to
temporary files, then merges the run heads. More `work_mem` can prevent a spill, but it applies per eligible plan node,
per worker, per concurrent query—not per server.

```sql
EXPLAIN (ANALYZE, BUFFERS)
SELECT transaction_id, amount
FROM transactions
WHERE tenant_id = 1
ORDER BY amount DESC, transaction_id
LIMIT 20;
```

Hash aggregation stores one entry per group and is expected `O(n)`; high group cardinality can spill or use large
memory. Group aggregation consumes ordered rows and can stream groups, but may require a sort. `DISTINCT` likewise uses
hashing or sorting.

Fix the access path before globally raising memory. A matching index can avoid sorting, an earlier selective predicate
can shrink it, and a rollup can avoid repeating it. Related learning:
[heaps](../python-algorithms-interview/008-heaps-and-streaming-statistics.md).
