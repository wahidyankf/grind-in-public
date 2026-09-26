---
tldr: "Combines partition routing, cache policy, workload isolation, and tenant placement without losing correctness."
when_to_use: "Use when scale or noisy-neighbour behaviour differs substantially across keys or tenants."
---

# Partitioning, Caching, and Multi-Tenancy

## Consistent hashing

Consistent hashing maps keys and nodes onto a ring so adding a node moves only a fraction of keys. Virtual nodes improve
balance and represent unequal capacity.

```text
          N1
       /      \
    kA          N2
    |            |
    N4          kB
       \      /
          N3
```

Production need: client-side cache sharding or routing to stateful workers while minimizing movement. Cost: membership
coordination, uneven hot keys, and difficult cross-key operations. Reject it when the datastore already provides robust
partitioning or when central directory placement is needed for residency.

## Rendezvous hashing

Compute a score for every candidate node and choose the highest. It is simple, stable under membership changes, and can
support weights, but scoring every node costs `O(nodes)` unless candidates are reduced.

## Tenant placement directory

```text
request tenant T -> placement cache -> cluster/region/shard
                         |
                         `-> authoritative directory on miss
```

Version placements and include the version in requests/events during moves. A migration state can route reads to old and
new while allowing exactly one writer. Reject transparent live movement when the business can tolerate a scheduled
tenant freeze; simpler mechanisms are safer.

## Cache hierarchy

```text
process LRU -> distributed cache -> canonical store
  microseconds       milliseconds       durable
```

Local caches reduce network latency but multiply invalidation and memory. Use immutable versioned configuration keys and
bounded LRU eviction. Never key only by resource id in a multi-tenant system; include tenant and authorization- relevant
version.

## Noisy-neighbour controls

- per-tenant token bucket at admission;
- weighted fair queues for asynchronous work;
- concurrency limits around scarce dependencies;
- quotas for storage and expensive queries;
- dedicated cells for outlier tenants.

Global limits alone let one tenant consume the budget. Per-tenant limits alone may waste idle shared capacity. Use a
hierarchy: global safety ceiling, tenant entitlement, and optional borrowing.
