---
tldr: "Diagnose table growth and vacuum lag caused by long-lived snapshots."
when_to_use: "Use for MVCC, vacuum, and operational ownership practice."
---

# Case 005: Long Transactions and Bloat

## Functional requirements

- FR-1: Keep case searches and updates available.
- FR-2: Preserve a nightly export's complete logical result.
- FR-3: Retain 90 days online and archive older evidence.

## Non-functional requirements

- NFR-1: Recover 500 GB of avoidable growth without a long exclusive lock.
- NFR-2: Keep write p99 below 250 ms during recovery.
- NFR-3: Prevent transaction-ID wraparound risk.
- NFR-4: Detect recurrence before disk exceeds 75%.

A six-hour repeatable-read export coincides with update-heavy case workflow. Explain evidence, immediate action, export
redesign, vacuum/bloat recovery, retention, and alerts.

[Debrief](../debriefs/005-long-transactions-and-bloat.md)
