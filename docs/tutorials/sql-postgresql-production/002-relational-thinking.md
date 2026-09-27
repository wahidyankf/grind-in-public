---
tldr: "Explains relations, keys, rows, nulls, and declarative query reasoning through production invariants."
when_to_use: "Use before writing SQL or translating an object model into tables."
---

# Relational Thinking

SQL describes the result, while PostgreSQL chooses a physical algorithm. A relation is conceptually an unordered set of
tuples; a table is its stored implementation. Therefore, no query promises order without `ORDER BY`.

```text
business invariant -> key/constraint -> valid relation -> declarative query -> chosen physical plan
```

A candidate key uniquely identifies a row; the primary key is the selected candidate. A foreign key says a referenced
identity must exist. The composite `(tenant_id, customer_id)` key makes tenant membership part of identity, allowing the
database—not every caller—to reject cross-tenant references.

`NULL` means unknown or inapplicable, not empty and not zero. SQL uses three-valued logic:

| Expression          | Result    |
| ------------------- | --------- |
| `NULL = NULL`       | `UNKNOWN` |
| `NULL IS NULL`      | `TRUE`    |
| `TRUE AND UNKNOWN`  | `UNKNOWN` |
| `FALSE AND UNKNOWN` | `FALSE`   |

`WHERE` keeps only `TRUE`, so `assignee <> 'lee'` also excludes unassigned cases. Write the intended predicate:

```sql
SELECT tenant_id, case_id, assignee
FROM cases
WHERE assignee IS NULL OR assignee <> 'lee';
```

Production rule: encode stable invariants in `NOT NULL`, `CHECK`, `UNIQUE`, foreign keys, and transactions. Application
validation improves errors but races with other writers. Reject a constraint only when the database cannot express the
rule without external or time-dependent state.

Checkpoint: explain why `customer_id = 1` is not a complete identity in this schema.
