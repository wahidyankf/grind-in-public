---
tldr: "Removes the snapshot blocker, restores vacuum progress, and moves exports off long primary transactions."
when_to_use: "Use after attempting Case 005."
---

# Debrief 005: Long Transactions and Bloat

Correlate transaction age/backend, dead tuples, vacuum progress, relation/index sizes, WAL, replication slots, and disk.
Cancel the export only through incident authority, then confirm vacuum advances. Export from a replica or consistent
snapshot copied into an analytical system, with bounded chunks and a documented consistency point. Tune per-table
autovacuum and fillfactor from churn. Use concurrent index rebuild or controlled online rewrite tooling when reclaim is
required; avoid `VACUUM FULL` on the hot path. Partition future time-based evidence so retention detaches whole ranges.
