---
tldr: "Builds a funnel from one-row-per-customer milestone CTEs."
when_to_use: "Use after Drill 007."
---

# Solution 007: Funnels and Cohorts

```sql
WITH customer_events AS (
    SELECT a.tenant_id, a.customer_id, min(t.occurred_at) AS first_transaction
    FROM accounts AS a JOIN transactions AS t USING (tenant_id, account_id)
    GROUP BY a.tenant_id, a.customer_id
), first_alerts AS (
    SELECT a.tenant_id, a.customer_id, min(al.created_at) AS first_alert
    FROM accounts AS a
    JOIN transactions AS t USING (tenant_id, account_id)
    JOIN alerts AS al USING (tenant_id, transaction_id)
    GROUP BY a.tenant_id, a.customer_id
)
SELECT date_trunc('month', ce.first_transaction) AS cohort,
       count(*) AS customers,
       count(*) FILTER (
           WHERE fa.first_alert < ce.first_transaction + interval '7 days'
       ) AS alerted_within_7d
FROM customer_events AS ce
LEFT JOIN first_alerts AS fa USING (tenant_id, customer_id)
GROUP BY cohort ORDER BY cohort;
```
