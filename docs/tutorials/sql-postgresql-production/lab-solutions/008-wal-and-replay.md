---
tldr: "Applies the WAL-before-data and WAL-before-durable-acknowledgement rules."
when_to_use: "Use after completing Lab 008."
---

# Solution 008: WAL and Replay

Change buffers and generate WAL; flush commit WAL before a durable acknowledgement; data pages may be written later, but
never before their WAL is durable. Recovery loads checkpoint state and replays later WAL to consistency. Replication
copies valid changes and mistakes; a backup provides retained independent history and point-in-time restore.
