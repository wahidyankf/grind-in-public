---
tldr: "Uses caches for bounded acceleration and ephemeral coordination without confusing them with canonical state."
when_to_use: "Use when repeated reads, expensive computation, counters, or short-lived coordination dominate latency."
---

# Caches and Ephemeral State

A cache is a performance dependency with a correctness policy. Define the key, value, owner, freshness, invalidation,
miss path, capacity policy, and failure behaviour.

## Cache-aside

```text
read key -> cache hit -> return
    |
    `-> miss -> read canonical store -> populate with TTL -> return
```

Cache-aside is simple but allows a stale interval after a write. Delete the key after the canonical commit; repopulating
inside the write transaction can publish data before the commit. Versioned keys avoid ambiguous invalidation for
immutable configurations.

## Stampede control

When a popular key expires, many callers may recompute it. Use single-flight locking, probabilistic early refresh,
stale-while-revalidate, and TTL jitter. Locks require ownership tokens and bounded leases so one caller cannot release
another's lock.

## Useful ephemeral structures

- expiring key: idempotency reservation or session;
- atomic counter: approximate rate-limit bucket;
- sorted set: small sliding-window timestamps or priority ordering;
- bitmap: compact membership/state flags;
- stream: limited operational queue where its durability semantics fit.

Memory databases are fast because working state is memory-oriented, not because networks and failover disappear.
Persistence and replication modes change latency and loss windows.

## Hot keys and failure

A celebrity tenant or global configuration can saturate one shard. Replicate read-only keys, add local caches, split a
counter, or isolate the tenant. On cache outage, protect the database with bounded concurrency and load shedding; a 100%
miss storm can be worse than the original outage.

## Reject when

Reject caching if the canonical query already meets its budget and invalidation risk exceeds the benefit. Reject a cache
as the only record for auditable decisions or money movement.

## References

- [Redis documentation](https://redis.io/docs/latest/)
