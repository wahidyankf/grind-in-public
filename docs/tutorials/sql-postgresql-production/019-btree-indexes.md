---
tldr: "Derives B-tree search, split, range-scan, multicolumn, covering, partial, and expression-index behaviour."
when_to_use: "Use for equality, ordered range, pagination, and uniqueness access paths."
---

# B-Tree Indexes

A PostgreSQL B-tree is a balanced, high-fan-out search tree. Internal pages route by separator keys; leaf pages contain
ordered keys and tuple identifiers. Leaves are linked for range scans.

```text
                         [ 400 | 900 ]
                        /      |       \
        [100 220 390] <-> [400 550 870] <-> [900 970]
             leaf               leaf              leaf
```

Searching is `O(log_b n)` page visits, where fan-out `b` is large. A range scan pays the search plus `O(k)` matching
entries. Inserts locate a leaf; a full page splits and propagates a separator upward. Random keys reduce locality and
can cause more page churn than mostly increasing keys, while purely increasing traffic concentrates writes at the right
edge.

Column order follows the usable prefix. For `(tenant_id, status, created_at DESC)`, equality on tenant and status plus a
created-time range is ideal. A created-time predicate alone cannot efficiently seek past all tenant/status prefixes.

```sql
CREATE INDEX alerts_open_by_tenant_idx
    ON alerts (tenant_id, created_at DESC, alert_id DESC)
    INCLUDE (score, transaction_id)
    WHERE status = 'open';
```

This partial covering index is small and fits the open-alert queue. The query predicate must imply `status = 'open'`.
Included columns are payload, not search keys. Deduplication, prefix compression, page splits, heap fetches, vacuum, and
write amplification mean an index is never free.

Related learning: [trees and tries](../python-algorithms-interview/007-trees-and-tries.md) and
[binary search](../python-algorithms-interview/005-binary-search-sorting-and-selection.md).
