---
tldr: "Uses joins and set operations while controlling cardinality and preserving unmatched rows."
when_to_use: "Use for multi-table queries and interview questions about relationship semantics."
---

# Joins and Set Operations

Start with relationship cardinality. A one-to-many join multiplies rows; aggregation after an accidental many-to-many
join multiplies money.

```sql
SELECT t.transaction_id, t.amount, c.full_name
FROM transactions AS t
JOIN accounts AS a
  ON a.tenant_id = t.tenant_id AND a.account_id = t.account_id
JOIN customers AS c
  ON c.tenant_id = a.tenant_id AND c.customer_id = a.customer_id
WHERE t.tenant_id = 1;
```

Keep predicates for the optional side in `ON` when a left join must preserve unmatched rows:

```sql
SELECT c.case_id, count(ca.alert_id) AS alert_count
FROM cases AS c
LEFT JOIN case_alerts AS ca
  ON ca.tenant_id = c.tenant_id
 AND ca.case_id = c.case_id
 AND ca.added_at >= now() - interval '7 days'
WHERE c.tenant_id = 1
GROUP BY c.case_id;
```

`EXISTS` is a semi-join: it asks whether any match exists without multiplying the left row. `NOT EXISTS` is safer than
`NOT IN` when nulls may occur.

```sql
SELECT a.account_id
FROM accounts AS a
WHERE a.tenant_id = 1
  AND NOT EXISTS (
      SELECT 1 FROM transactions AS t
      WHERE t.tenant_id = a.tenant_id AND t.account_id = a.account_id
  );
```

`UNION` removes duplicates through hashing or sorting; `UNION ALL` concatenates and is cheaper. Use set operations when
inputs have compatible meaning, not to hide a broken model.
