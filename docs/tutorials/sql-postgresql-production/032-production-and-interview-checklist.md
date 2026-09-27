---
tldr: "Synthesizes SQL, PostgreSQL internals, production operations, and interview communication into one checklist."
when_to_use: "Use for final review, design interviews, and production database decisions."
---

# Production and Interview Checklist

For a SQL problem, say the output grain, edge cases, query, supporting index, complexity, and test cases. For a database
design, trace requirements to invariants, access paths, consistency, capacity, failure recovery, and evolution.

```text
functional requirements -> schema + transactions + APIs
non-functional targets  -> indexes + topology + capacity + operations
failure model            -> retries + backups + replication + fencing
evolution                -> migrations + backfills + compatibility
evidence                 -> plans + metrics + load tests + restore drills
```

## Production review

- Are tenant identity and authorization enforced at every data boundary?
- Which constraints own correctness, and which transaction owns each invariant?
- What are the top query shapes, result cardinalities, indexes, and worst-case plans?
- Are timeouts, pools, retries, and idempotency bounded under overload?
- Can autovacuum keep pace with updates, and are long transactions controlled?
- What write amplification comes from indexes, WAL, replicas, and denormalized projections?
- What are RPO, RTO, promotion/fencing steps, and the last successful restore drill?
- Can migrations run with old and new application versions concurrently?
- What happens at 10x data, one hot tenant, a zone loss, and a stale replica?

## Interview communication

State assumptions before selecting technology. Use numbers: rows/day, bytes/row, retention, read/write QPS, concurrency,
p99, availability, RPO, and RTO. Explain one rejected alternative. End with the next bottleneck and an evolution
trigger, not an imaginary infinitely scalable design.

Continue with [SQL drills](drills/README.md), [production cases](cases/README.md), and the
[unified programme](../engineering-interview-program/README.md).
