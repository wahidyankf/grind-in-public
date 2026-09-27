---
tldr: "Explains snapshots, tuple visibility, update chains, and heap-only tuples."
when_to_use: "Use to reason about concurrent reads, dead tuples, update amplification, and index maintenance."
---

# MVCC and HOT

Multi-version concurrency control gives a statement or transaction a snapshot. A tuple version is visible when its
creating transaction is visible and its deleting/replacing transaction is not visible to that snapshot.

```text
logical case 42

tuple v1: xmin=100, xmax=130 -> old snapshot can see
                              \
tuple v2: xmin=130, xmax=0   -> newer snapshot can see

index key -> v1 -> HOT chain -> v2
```

An `UPDATE` creates a new version. When indexed columns are unchanged and the page has room, a heap-only tuple (HOT)
update can avoid new index entries; the old tuple links to the new version. Lowering table `fillfactor` reserves page
space and can improve HOT rate for update-heavy tables, at the cost of larger heaps and fewer tuples per read.

```sql
SELECT relname, n_tup_upd, n_tup_hot_upd, n_dead_tup
FROM pg_stat_user_tables
WHERE relname IN ('cases', 'alerts', 'transactions');
```

Long transactions keep old snapshots alive, so vacuum cannot remove versions they might still see. That increases
heap/index work and replication/recovery pressure. MVCC reduces read/write blocking; it does not eliminate conflicting
writer locks or guarantee serializable business outcomes.
