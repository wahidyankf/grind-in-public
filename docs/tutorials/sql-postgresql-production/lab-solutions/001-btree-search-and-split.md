---
tldr: "Shows the B-tree split, range walk, complexity, and production index."
when_to_use: "Use after completing Lab 001."
---

# Solution 001: B-Tree Search and Split

```text
             [30 | 50]
            /    |    \
    [10 20] <-> [30 40] <-> [50 60]
```

Exact split placement varies by implementation; the invariant is ordered occupancy and equal leaf depth. Seek `25` in
`O(log_b n)`, then follow linked leaves for `O(k)` output. Random inserts cause scattered page reads and splits. Use
`(tenant_id, occurred_at DESC, transaction_id DESC)`; its useful leading prefixes start with `tenant_id`.
