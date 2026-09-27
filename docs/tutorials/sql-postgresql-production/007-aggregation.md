---
tldr: "Builds grouped metrics and explains cardinality, conditional aggregation, and production rollups."
when_to_use: "Use for dashboards, summaries, and aggregate interview queries."
---

# Aggregation

Aggregation changes grain. Write the output grain before the query: “one row per tenant and UTC day.”

```sql
SELECT tenant_id,
       date_trunc('day', occurred_at) AS day,
       count(*) AS transaction_count,
       sum(amount) FILTER (WHERE direction = 'debit') AS debit_total,
       count(*) FILTER (WHERE amount >= 1000) AS large_count
FROM transactions
WHERE occurred_at >= '2026-01-01T00:00:00Z'
  AND occurred_at <  '2026-02-01T00:00:00Z'
GROUP BY tenant_id, date_trunc('day', occurred_at)
ORDER BY tenant_id, day;
```

Use half-open time ranges so adjacent windows neither overlap nor leave a gap. `WHERE` filters input rows; `HAVING`
filters groups:

```sql
SELECT tenant_id, account_id, sum(amount) AS debit_total
FROM transactions
WHERE direction = 'debit'
GROUP BY tenant_id, account_id
HAVING sum(amount) >= 1000;
```

PostgreSQL can use hash aggregation (`O(n)` expected, memory-sensitive) or sort/group aggregation (`O(n log n)`, can
spill). For repeated expensive dashboards, consider incremental rollups or a materialized view, but define freshness,
rebuild, and correction semantics. Never make a mutable cache the sole financial authority.
