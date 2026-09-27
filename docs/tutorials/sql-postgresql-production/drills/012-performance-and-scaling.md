---
tldr: "Diagnoses a query regression and evolves a single primary toward tenfold load."
when_to_use: "Use for production and system-design interview synthesis."
---

# Drill 012: Performance and Scaling

After one tenant grows 100 times, p99 rises from 80 ms to 4 s, pool wait increases, and a plan changes from index scan
to nested loops over underestimated rows. Give the diagnostic order, immediate safe mitigation, durable fix, and 10x
architecture evolution without jumping directly to sharding.

[Reference solution](../solutions/012-performance-and-scaling.md)
