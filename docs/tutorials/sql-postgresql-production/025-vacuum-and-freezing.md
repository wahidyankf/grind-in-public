---
tldr: "Explains pruning, vacuum, autovacuum thresholds, freezing, bloat, and transaction-ID wraparound prevention."
when_to_use: "Use when tables bloat, index-only scans regress, or autovacuum cannot keep up."
---

# Vacuum and Freezing

Vacuum does not normally shrink the table file. It marks dead tuple space reusable, cleans index references, advances
visibility information, and freezes sufficiently old transaction metadata. Page pruning may shorten HOT chains earlier.

```text
updates -> dead tuple versions -> vacuum eligibility -> reusable space
                ^                       |
                |                       +-> visibility map / index cleanup
old snapshot ---+ prevents cleanup
```

Transaction IDs are finite and compared by age. Freezing replaces old transactional visibility with a permanent form,
preventing wraparound from making ancient rows appear new. Anti-wraparound vacuum is a safety mechanism, not optional
maintenance.

```sql
SELECT s.relname, s.n_live_tup, s.n_dead_tup,
       s.last_autovacuum, s.autovacuum_count,
       age(c.relfrozenxid) AS xid_age
FROM pg_stat_user_tables AS s
JOIN pg_class AS c ON c.oid = s.relid
ORDER BY s.n_dead_tup DESC;
```

Default scale-factor thresholds can be too slow for a huge hot table: a small percentage is still millions of dead rows.
Tune per table from update rate and observed cleanup. Ensure workers have I/O budget. Find and terminate or fix
idle-in-transaction clients before adding vacuum workers.

`VACUUM FULL` rewrites and exclusively locks the table; it is exceptional recovery, not routine bloat management. Prefer
preventing bloat, rebuilding indexes concurrently when justified, and rehearsing any rewrite.
