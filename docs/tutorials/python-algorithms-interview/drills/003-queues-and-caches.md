---
tldr: "Practises explicit overload and eviction policies with deques and an LRU cache."
when_to_use: "Use after linked structures, queues, and production algorithms."
---

# Queues and Caches

## Drill A: bounded fair queue

Build a queue with a fixed total capacity and per-tenant FIFO order. Dequeue tenants round-robin so one noisy tenant
cannot starve another. Reject enqueue when total capacity is full.

- No tenant may appear twice in the active rotation.
- Empty tenant queues disappear from the rotation.
- State the cost of enqueue and dequeue.

## Drill B: LRU decision cache

Build a fixed-capacity cache mapping string keys to integer decisions. `get` refreshes recency; `put` updates or
inserts; inserting past capacity evicts the least recently used key. Reject non-positive capacity.

Explain why this local cache cannot be the source of truth after a rule-version change.

[Solution](../solutions/003-queues-and-caches.md)
