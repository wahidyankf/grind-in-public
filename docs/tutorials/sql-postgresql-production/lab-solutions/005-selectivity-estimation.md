---
tldr: "Shows a 12-times underestimate caused by correlated tenant and status values."
when_to_use: "Use after completing Lab 005."
---

# Solution 005: Selectivity Estimation

Independence estimates `10m * 0.40 * 0.05 = 200,000`; reality is `10m * 0.40 * 0.60 = 2.4m`, a 12-times underestimate.
That can select a nested loop or undersized hash/sort. Add multicolumn most-common-value/dependency statistics on
`(tenant_id, status)` and `ANALYZE`; isolate a very large tenant only when broader capacity and blast-radius evidence
also supports it.
