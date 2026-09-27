---
tldr: "Restores service by bounding admission and repairing the first plan-estimate divergence."
when_to_use: "Use after attempting Case 001."
---

# Debrief 001: Query-Plan Regression

```text
request -> pool wait -> query id -> plan/parameters -> first row-estimate error
   |                                                     |
   +-> admission cap / timeout                    stats, index, query
```

Correlate release, statistics, and load timing; compare good/bad query id, parameter class, estimates/actuals, buffers,
spills, locks, and I/O. Cap concurrency and use the last safe query path as reversible mitigation. Refresh or extend
statistics for correlated tenant/status data and add a matching partial index only if its write cost is acceptable.
Load-test FR results plus p99/write throughput, deploy gradually, and retain plan/metric evidence in the post-incident
review.
