---
tldr: "Explains write-ahead logging, commit durability, full-page images, checkpoints, and crash recovery."
when_to_use: "Use for durability, write amplification, recovery-time, and backup reasoning."
---

# WAL and Recovery

Write-ahead logging records redo information before dirty data pages reach durable storage.

```text
transaction changes shared buffers
          |
          +-> WAL records -> WAL flush -> COMMIT acknowledged
                                   |
crash -> last checkpoint -> replay WAL -> consistent database
                                   |
dirty data pages may reach storage later
```

PostgreSQL assigns log sequence numbers (LSNs). Commit durability normally requires WAL through the commit record to be
flushed. `synchronous_commit = off` can reduce latency but permits recent acknowledged transactions to disappear after
an operating-system crash; it does not corrupt the database. Use only when that explicit loss window is acceptable.

A checkpoint establishes a recovery starting region and drives dirty-page writes. Too-frequent checkpoints increase
write pressure and full-page images after each checkpoint; too-infrequent checkpoints need more WAL and can lengthen
recovery. Smooth checkpoint I/O and monitor generated WAL, checkpoint duration, storage saturation, and recovery goals.

Full-page images protect against torn pages after the first post-checkpoint modification. WAL is not a logical audit log
and not a backup by itself. Point-in-time recovery requires a base backup plus a continuous, tested WAL archive.
