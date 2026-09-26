---
tldr: "Connects interview structures to caches, rate limits, approximate analytics, matching, and partition routing."
when_to_use: "Use after the core patterns to understand where algorithms sit inside production system boundaries."
---

# Production Algorithms

An algorithm becomes a production mechanism only after its scope, ownership, persistence, concurrency, and failure
semantics are explicit.

| Need                | Mechanism          | Production question                              |
| ------------------- | ------------------ | ------------------------------------------------ |
| bounded cache       | LRU/LFU            | is stale data acceptable after failover?         |
| admission control   | token bucket       | is the quota local, regional, or global?         |
| membership precheck | Bloom filter       | what false-positive rate fits the memory budget? |
| distinct estimate   | HyperLogLog        | how much error is acceptable?                    |
| frequency estimate  | Count-Min Sketch   | can overestimation change a decision?            |
| partition routing   | consistent hashing | how are membership changes coordinated?          |
| fuzzy candidates    | n-grams or trie    | which recall target precedes scoring?            |
| graph component     | union-find         | how are edge deletions and evidence handled?     |

## Grounded selection

Suppose a screening service wants to avoid remote lookups for keys definitely absent from a large immutable list. A
Bloom filter is useful because a negative result is definitive and memory is bounded; a positive result still requires
the authoritative lookup.

```text
key -> Bloom filter -- definitely absent --> stop
                   `-- maybe present -----> authoritative index
```

It is unsafe if a false positive directly blocks a customer. The filter is an optimization, not the decision source.

Likewise, consistent hashing reduces remapped keys when workers change, but it does not coordinate membership, migrate
state, replicate data, or prevent split-brain views. Those are distributed-system responsibilities around the local
algorithm.

## Interview answer frame

For any production algorithm, state:

1. the exact operation it optimizes;
2. input size and required bound;
3. correctness or approximation contract;
4. time and memory cost;
5. concurrency and persistence scope;
6. failure and recovery behaviour;
7. authoritative fallback;
8. the metric that proves it helps.

## Checkpoint

Choose one row from the table and explain a situation where the mechanism's local benefit is outweighed by operational
complexity.
