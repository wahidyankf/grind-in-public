---
tldr: "Explains heap pages, tuple versions, line pointers, free space, visibility, and TOAST storage."
when_to_use: "Use to understand table bloat, wide rows, update cost, and index-only scans."
---

# Storage, Pages, and TOAST

Tables are heap files divided into fixed-size pages, commonly 8 KiB at build time. A page has a header, line pointers,
tuple data growing from the opposite end, and free space between them.

```text
+--------------------------- 8 KiB page ----------------------------+
| page header | item pointers ->        free       <- tuple bodies |
+------------------------------------------------------------------+
                    pointer -> tuple header + null map + values
```

A heap tuple header carries transaction visibility metadata such as inserting and deleting transaction identifiers. An
index points to a tuple location `(block, offset)`, not a permanent logical row. Updating normally creates a new tuple
version and leaves the old version until vacuum can reclaim it.

The free-space map helps find pages with room for inserts and updates. The visibility map records pages known all-
visible or all-frozen; index-only scans can skip heap visibility checks only for all-visible pages. Heavy churn or
lagging vacuum reduces that benefit.

TOAST compresses or moves large variable-length values to a side table. Fetching `SELECT *` may detoast payloads the
caller never needs. A large `jsonb` update rewrites a value and generates WAL; PostgreSQL is not doing in-place field
mutation inside an arbitrary document.

Inspect safely:

```sql
SELECT pg_size_pretty(pg_relation_size('transactions')) AS heap,
       pg_size_pretty(pg_indexes_size('transactions')) AS indexes,
       pg_size_pretty(pg_total_relation_size('transactions')) AS total;
```
