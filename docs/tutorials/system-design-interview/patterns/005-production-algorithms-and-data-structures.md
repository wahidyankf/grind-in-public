---
tldr:
  "Maps interview algorithms and data structures to rate limits, deduplication, matching, ranking, graphs, and
  scheduling."
when_to_use: "Use to connect algorithm exercises to production architecture and capacity decisions."
---

# Production Algorithms and Data Structures

The [Python algorithms course](../../python-algorithms-interview/README.md) teaches implementation. This lesson explains
where the structures enter a distributed design.

| Technique                  | Production use                             | Main cost / rejection case                                |
| -------------------------- | ------------------------------------------ | --------------------------------------------------------- |
| Hash map/set               | idempotency, joins, feature lookup         | memory; reject if exact set exceeds affordable RAM        |
| Heap                       | top-k alerts, timers, merge sorted streams | `O(log k)` updates; reject for full global sort           |
| Trie                       | prefix rules, route matching               | node overhead; reject for arbitrary substring             |
| Bloom filter               | avoid expensive negative lookups           | false positives; reject if false negatives allowed? Never |
| Count-Min Sketch           | approximate heavy hitters                  | overestimation; reject for exact billing                  |
| HyperLogLog                | approximate distinct entities              | no membership; reject for exact small sets                |
| Union-find                 | batch connected components                 | weak for deletions; reject for dynamic traversal          |
| Dijkstra / A*              | weighted risk or routing paths             | graph size; reject negative weights for Dijkstra          |
| Sliding window / deque     | velocity and rolling extrema               | per-key state; reject if coarse buckets suffice           |
| Token bucket               | burst-tolerant rate limiting               | distributed atomicity; reject for exact windows           |
| Consistent/rendezvous hash | stable placement across changing nodes     | skew/membership; reject if directory is required          |

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
