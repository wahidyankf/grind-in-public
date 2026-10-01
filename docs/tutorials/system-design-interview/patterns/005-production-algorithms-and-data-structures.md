---
tldr:
  "Maps interview algorithms and data structures to rate limits, deduplication, matching, ranking, graphs, and
  scheduling."
when_to_use: "Use to connect algorithm exercises to production architecture and capacity decisions."
---

# Production Algorithms and Data Structures

The [Python algorithms course](../../python-algorithms-interview/README.md) teaches implementation. This lesson explains
where the structures enter a distributed design.

- **Hash map/set**
  - Production use: idempotency, joins, feature lookup
  - Main cost / rejection case: memory; reject if exact set exceeds affordable RAM
- **Heap**
  - Production use: top-k alerts, timers, merge sorted streams
  - Main cost / rejection case: `O(log k)` updates; reject for full global sort
- **Trie**
  - Production use: prefix rules, route matching
  - Main cost / rejection case: node overhead; reject for arbitrary substring
- **Bloom filter**
  - Production use: avoid expensive negative lookups
  - Main cost / rejection case: false positives; reject if false negatives allowed? Never
- **Count-Min Sketch**
  - Production use: approximate heavy hitters
  - Main cost / rejection case: overestimation; reject for exact billing
- **HyperLogLog**
  - Production use: approximate distinct entities
  - Main cost / rejection case: no membership; reject for exact small sets
- **Union-find**
  - Production use: batch connected components
  - Main cost / rejection case: weak for deletions; reject for dynamic traversal
- **Dijkstra / A***
  - Production use: weighted risk or routing paths
  - Main cost / rejection case: graph size; reject negative weights for Dijkstra
- **Sliding window / deque**
  - Production use: velocity and rolling extrema
  - Main cost / rejection case: per-key state; reject if coarse buckets suffice
- **Token bucket**
  - Production use: burst-tolerant rate limiting
  - Main cost / rejection case: distributed atomicity; reject for exact windows
- **Consistent/rendezvous hash**
  - Production use: stable placement across changing nodes
  - Main cost / rejection case: skew/membership; reject if directory is required

## Exact versus approximate

At large scale, the first design question is often whether exactness changes the business outcome. Approximation is
appropriate for monitoring, prefiltering, and candidate generation when a precise second stage exists. It is usually
inappropriate for final monetary, access-control, or evidence decisions.

```text
cheap approximate filter -> candidate set -> exact authoritative check
```

Bloom filters can say "possibly present" or "definitely absent." Use them to avoid disk/network lookups; always verify
positive results. Size them from expected elements and acceptable false-positive rate, and rebuild before saturation.

## Online versus batch

Online algorithms keep bounded incremental state and respond within a per-event budget. Batch algorithms can sort, scan,
and recompute globally. A common production design uses online approximation for immediate action and batch exact
reconciliation for assurance.

## Adversarial input

Complexity assumptions become security boundaries. Bound graph depth and fan-out, regex/rule evaluation, request size,
heap cardinality, and hash-controlled keys. A theoretically linear parser can still allocate unbounded memory.
