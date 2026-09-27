---
tldr: "Uses NOT EXISTS for a tenant-safe anti-join."
when_to_use: "Use after Drill 002."
---

# Solution 002: Joins and Anti-Joins

```sql
SELECT a.account_id, c.full_name
FROM accounts AS a
JOIN customers AS c
  ON c.tenant_id = a.tenant_id AND c.customer_id = a.customer_id
WHERE a.tenant_id = 1 AND a.status = 'open'
  AND NOT EXISTS (
      SELECT 1 FROM transactions AS t
      WHERE t.tenant_id = a.tenant_id AND t.account_id = a.account_id
        AND t.occurred_at >= '2026-01-01T00:00:00Z'
        AND t.occurred_at <  '2026-02-01T00:00:00Z'
  );
```

An index on `(tenant_id, account_id, occurred_at)` supports the existence probe. A null inside `NOT IN` makes its result
unknown for every candidate.
