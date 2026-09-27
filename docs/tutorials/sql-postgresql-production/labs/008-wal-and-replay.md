---
tldr: "Orders WAL, dirty-page writes, checkpoints, and crash replay."
when_to_use: "Use after the WAL lesson to test durability reasoning."
---

# Lab 008: WAL and Replay

Order these events: dirty a data page, emit WAL, flush commit WAL, acknowledge commit, write the data page. Which order
is forbidden? Explain recovery from the last checkpoint and why a replica is not a backup.

Compare with the [solution](../lab-solutions/008-wal-and-replay.md).
