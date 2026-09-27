---
tldr: "Solves daily conditional aggregation with a half-open time range."
when_to_use: "Use after Drill 001."
---

# Solution 001: Filtering and Aggregation

```sql
SELECT tenant_id, date_trunc('day', occurred_at AT TIME ZONE 'UTC') AS utc_day,
       count(*) AS debit_count, sum(amount) AS debit_amount,
       count(*) FILTER (WHERE amount >= 1000) AS large_debit_count
FROM transactions
WHERE direction = 'debit'
  AND occurred_at >= '2026-01-01T00:00:00Z'
  AND occurred_at <  '2026-02-01T00:00:00Z'
GROUP BY tenant_id, utc_day
HAVING sum(amount) > 2000
ORDER BY tenant_id, utc_day;
```

The grain is tenant/day. `amount` is non-null by schema. Use an explicit business time zone if “day” is tenant-local.
