---
tldr: "Uses scalar, correlated, and common-table expressions without treating syntax as an optimization boundary."
when_to_use: "Use when decomposing complex queries or reasoning about repeated work."
---

# Subqueries and CTEs

A scalar subquery must return at most one row. A correlated subquery can execute per outer row unless the planner
decorrelates it. Express existence with `EXISTS` and inspect the plan.

```sql
SELECT c.customer_id, c.full_name
FROM customers AS c
WHERE c.tenant_id = 1
  AND EXISTS (
      SELECT 1
      FROM accounts AS a
      JOIN transactions AS t
        ON t.tenant_id = a.tenant_id AND t.account_id = a.account_id
      WHERE a.tenant_id = c.tenant_id
        AND a.customer_id = c.customer_id
        AND t.amount >= 1000
  );
```

A CTE names a relation and often improves reasoning:

```sql
WITH account_totals AS (
    SELECT tenant_id, account_id, sum(amount) AS total
    FROM transactions
    WHERE direction = 'debit'
    GROUP BY tenant_id, account_id
)
SELECT a.account_id, at.total
FROM accounts AS a
JOIN account_totals AS at USING (tenant_id, account_id)
WHERE a.tenant_id = 1 AND at.total >= 1000;
```

Modern PostgreSQL may inline a side-effect-free CTE referenced once. `MATERIALIZED` forces one evaluated result;
`NOT MATERIALIZED` permits joint optimization. Force either only after measuring: materialization avoids repeated work
but consumes memory or temporary storage and may prevent predicate pushdown.
