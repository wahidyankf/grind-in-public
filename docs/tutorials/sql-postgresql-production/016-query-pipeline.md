---
tldr: "Traces SQL through parsing, rewriting, planning, execution, and instrumentation."
when_to_use: "Use to connect query text with plans, runtime operators, and observable performance."
---

# Query Pipeline

PostgreSQL transforms a declarative statement through distinct stages:

```text
SQL text
   |
   v
parser -> parse tree -> analyser -> typed query tree -> rewriter
                                                       |
                                                       v
executor <- physical plan <- cost-based planner <--- rewritten tree
   |
   +-> rows, locks, WAL, buffers, temporary files, statistics
```

Parsing checks grammar. Analysis resolves names, types, operators, and permissions. Rewriting expands views and rules.
The planner enumerates viable paths and chooses the lowest estimated cost. The executor pulls tuples through plan nodes,
often one at a time, although some nodes must consume input first.

```sql
EXPLAIN (ANALYZE, BUFFERS, WAL, SETTINGS, FORMAT TEXT)
SELECT account_id, sum(amount)
FROM transactions
WHERE tenant_id = 1
GROUP BY account_id;
```

Read a plan from the most indented node outward. Compare estimated `rows` with actual rows before blaming the selected
algorithm. `cost` is an abstract estimate, not milliseconds. `actual time` is measured execution time, and loops
multiply work. `BUFFERS` exposes cache and storage activity; temporary reads/writes signal spills; `WAL` shows write
amplification for mutations.

Prepared statements may use custom plans per parameter or a generic plan. A generic plan can be poor for skewed tenants.
Diagnose parameter sensitivity before forcing plan behaviour.
