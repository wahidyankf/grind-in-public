---
tldr: "Chooses PostgreSQL types and constraints that preserve meaning through schema evolution."
when_to_use: "Use when designing or reviewing tables and migrations."
---

# Types, Constraints, and DDL

Choose types by domain operations, not by their surface appearance.

| Need                   | Prefer              | Why                                            |
| ---------------------- | ------------------- | ---------------------------------------------- |
| Money                  | `numeric(p,s)`      | Exact decimal arithmetic                       |
| Measured approximation | `double precision`  | Fast floating-point operations                 |
| Global instant         | `timestamptz`       | Stored as an instant; rendered in session zone |
| Local calendar value   | `date`              | No accidental time-zone conversion             |
| Evolving sparse detail | `jsonb`             | Queryable document with controlled flexibility |
| Stable finite state    | `text` plus `CHECK` | Easy transactional evolution                   |

Do not store money in floating point, timestamps as strings, or comma-separated identifiers in text. Prefer identity
columns over sequence calls hidden in application code. Use domains only when one database-owned semantic type truly
recurs; otherwise a domain couples migrations more than it helps.

DDL example with an additive, low-risk evolution:

```sql
ALTER TABLE transactions ADD COLUMN reviewed_at timestamptz;

ALTER TABLE transactions
    ADD CONSTRAINT transactions_reviewed_after_recorded
    CHECK (reviewed_at IS NULL OR reviewed_at >= recorded_at)
    NOT VALID;

ALTER TABLE transactions
    VALIDATE CONSTRAINT transactions_reviewed_after_recorded;
```

`NOT VALID` avoids scanning existing rows while holding the initial stronger lock, while new writes are checked.
Validation later scans existing rows. Always test lock acquisition: even metadata-only DDL can wait behind a long
transaction and then block queued traffic.
