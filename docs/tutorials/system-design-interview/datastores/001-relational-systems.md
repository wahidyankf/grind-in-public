---
tldr: "Uses relational transactions, constraints, indexes, isolation, partitioning, and replicas for strong invariants."
when_to_use: "Use for systems with multi-row invariants, evolving queries, and authoritative workflow state."
---

# Relational Systems

Relational databases are a strong default for authoritative business state because constraints and transactions keep
related changes atomic. They also support flexible indexed queries and mature operational tooling.

## Production fit

Use a relational store for case workflow, configuration versions, idempotency reservations, ownership, and an outbox
when the write-side invariant spans several records.

```sql
CREATE TABLE decisions (
    tenant_id UUID NOT NULL,
    decision_id UUID NOT NULL,
    status TEXT NOT NULL CHECK (status IN ('allow', 'review', 'reject')),
    policy_version BIGINT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    PRIMARY KEY (tenant_id, decision_id)
);

CREATE INDEX decisions_tenant_created_idx
    ON decisions (tenant_id, created_at DESC);
```

The primary key prevents cross-tenant identity collision. The secondary index supports a named query: recent decisions
for one tenant. Every index speeds some reads but adds write amplification, storage, vacuum/compaction work, and lock or
cache pressure.

## Isolation

- read committed prevents dirty reads but allows repeated queries to see new committed data;
- repeatable read gives a stable transaction snapshot but may permit write skew depending on the engine;
- serializable aims to make concurrent outcomes equivalent to a serial order and may require retry.

Choose isolation per invariant and handle serialization failures. `SELECT ... FOR UPDATE` can coordinate a small hot set
but becomes a bottleneck under contention.

## Replicas and partitioning

Read replicas improve read capacity and failure options, but asynchronous replicas may be stale. Route read-after-write
queries to the leader or carry a consistency token. Table partitioning helps retention and pruning; it does not remove
the need for good indexes. Sharding is a major ownership and transaction change—delay it until vertical scaling,
partitioning, and workload tuning are insufficient.

## Reject when

Reject a single relational cluster as the entire high-volume event history when sustained append scale, retention, and
geographic distribution exceed its tested envelope. Keep authoritative metadata relational while moving immutable event
payloads to a log or object store.

## References

- [PostgreSQL documentation](https://www.postgresql.org/docs/current/)
