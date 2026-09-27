---
tldr: "Promotes the most advanced eligible standby only after preventing old-primary writes."
when_to_use: "Use after completing Lab 010."
---

# Solution 010: Replication and Failover

A is most advanced and shows no known loss relative to the stated acknowledgement point, but observability itself may be
stale. Fence the primary through routing, credentials, storage lease, or infrastructure power before promotion. Route
read-after-write traffic to the writer or wait for replay LSN. Rebuild the old primary from the new timeline; do not let
divergent writes rejoin.
