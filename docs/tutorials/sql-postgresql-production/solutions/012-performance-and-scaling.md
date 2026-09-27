---
tldr: "Fixes estimates and access paths before pooling, replicas, partitioning, or tenant cells."
when_to_use: "Use after Drill 012."
---

# Solution 012: Performance and Scaling

Capture query id/parameters, pool and lock waits, estimates versus actuals, buffers, temp I/O, and statistics freshness.
Bound admission and cancel runaway work as reversible mitigation. Add extended statistics or a tenant-aware index, then
load-test. Next scale the primary/storage, cap connection multiplication, cache/precompute named reads, use replicas for
stale-tolerant reads, isolate analytics, partition for lifecycle, and finally place exceptional tenants into cells.
Every step has a metric and rollback trigger.
