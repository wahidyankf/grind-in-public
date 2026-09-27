---
tldr: "Chooses hash, GIN, GiST, BRIN, and bitmap combinations by operator and data distribution."
when_to_use:
  "Use when a B-tree does not match containment, range overlap, large ordered tables, or combined predicates."
---

# Specialized Indexes

Choose an index by the operator family the query uses.

| Index | Strong fit                                       | Important cost                        |
| ----- | ------------------------------------------------ | ------------------------------------- |
| Hash  | Equality only                                    | No ordering or range                  |
| GIN   | Inverted membership: arrays, documents, text     | Expensive writes; pending-list upkeep |
| GiST  | Extensible spatial/range/nearest-neighbour rules | Lossy matches may need recheck        |
| BRIN  | Huge tables correlated with physical order       | Approximate block ranges; false hits  |

```sql
CREATE INDEX transactions_metadata_gin_idx
    ON transactions USING gin (metadata jsonb_path_ops);

CREATE INDEX transactions_occurred_brin_idx
    ON transactions USING brin (occurred_at) WITH (pages_per_range = 64);
```

GIN maps a token/key to a posting list of matching rows—the same inverted-index idea used by search systems. BRIN stores
summaries for page ranges, often min/max; its index stays tiny, then qualifying ranges are rechecked against the heap.
BRIN is excellent for append-ordered event time and poor after physical order becomes random.

PostgreSQL can combine multiple indexes into a bitmap of heap locations. That helps `AND`/`OR` predicates but loses
index ordering and may become lossy under memory pressure. A well-designed multicolumn index is often better for a hot,
stable query. Measure write rate, index size, cache residency, and vacuum cost—not only one read plan.
