---
tldr: "Selects a promotion candidate from replay positions and requires fencing."
when_to_use: "Use after replication to practise correctness-first failover."
---

# Lab 010: Replication and Failover

Primary last acknowledged LSN is `0/900`; standby A replayed `0/900`, B `0/8F0`. The primary is unreachable. State
candidate, possible loss, fencing steps, read consistency, and how to treat the old primary if it returns.

Compare with the [solution](../lab-solutions/010-replication-and-failover.md).
