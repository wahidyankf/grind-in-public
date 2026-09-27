---
tldr: "Explains selectivity estimates, histograms, common values, correlation, join ordering, costs, and GEQO."
when_to_use: "Use when actual rows diverge from estimates or the planner chooses a surprising path."
---

# Planner Statistics and Costs

The planner estimates rows, then costs possible paths. Per-column statistics include null fraction, distinct estimate,
most-common values, histogram bounds, and physical correlation. Independence assumptions fail when columns correlate.

```text
estimated rows = input rows * predicate selectivity
join rows       = outer rows * inner rows * join selectivity
```

If `tenant_id` and `account_id` are correlated, multiplying independent selectivities may be badly wrong. Extended
statistics capture dependencies, multicolumn distinct counts, or common value combinations:

```sql
CREATE STATISTICS transactions_tenant_account_stats
    (dependencies, ndistinct, mcv)
ON tenant_id, account_id
FROM transactions;

ANALYZE transactions;
```

Planning join order is combinatorial. PostgreSQL uses dynamic programming for manageable join counts, retaining good
paths for subsets; beyond configured thresholds the genetic query optimizer samples the search space. This is why a
twenty-table generated query may plan slowly or choose an unstable order.

Diagnosis order:

1. compare estimated and actual rows at the first divergence;
2. confirm fresh statistics and representative data;
3. increase a column statistics target or add extended statistics when justified;
4. rewrite a predicate only if its semantics are clearer or estimable;
5. change cost settings only from measured hardware-wide evidence.

Hints are not the first tool: they freeze a symptom while data distribution changes.
