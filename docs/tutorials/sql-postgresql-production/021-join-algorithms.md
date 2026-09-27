---
tldr: "Explains nested-loop, hash, and merge joins with complexity, memory, ordering, and cardinality trade-offs."
when_to_use: "Use when interpreting join plans or choosing indexes and query shapes."
---

# Join Algorithms

The logical join is one operation; the executor can implement it several ways.

```text
nested loop: for each outer row -> probe inner source
hash join:   build hash(smaller input) -> scan/probe larger input
merge join:  sort/scan both ordered inputs -> advance the smaller key
```

| Algorithm   | Approximate cost          | Strong fit                                     |
| ----------- | ------------------------- | ---------------------------------------------- |
| Nested loop | `O(n * probe)`            | Small outer input and indexed inner lookup     |
| Hash join   | `O(n + m)` expected       | Equality join; build side fits memory          |
| Merge join  | `O(n + m)` after ordering | Large ordered equality/range-compatible inputs |

A nested loop is not inherently slow: 20 outer rows with an `O(log m)` B-tree probe can beat building a large hash. A
hash join partitions/spills when its build table exceeds memory. Skewed keys create uneven buckets. A merge join can
reuse index order; otherwise sorting adds `O(n log n + m log m)`.

```sql
EXPLAIN (ANALYZE, BUFFERS)
SELECT t.transaction_id, a.status
FROM transactions AS t
JOIN accounts AS a
  ON a.tenant_id = t.tenant_id AND a.account_id = t.account_id
WHERE t.tenant_id = 1 AND t.occurred_at >= now() - interval '1 day';
```

Never force an algorithm before fixing cardinality estimates. The wrong row estimate makes the right algorithm appear
cheap. Related learning: [hashing](../python-algorithms-interview/002-arrays-strings-and-hashing.md) and the paired
[join lab](labs/003-join-algorithms.md).
