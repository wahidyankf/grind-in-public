---
tldr:
  "Explains partitioning, replication, consistency, time, retries, ordering, and idempotency as production contracts."
when_to_use: "Use when a design crosses processes, availability zones, or regions."
---

# Scaling and Distributed Correctness

A distributed system replaces local failure with partial failure: one participant may succeed while another times out.
Correctness therefore starts with explicit semantics, not retry middleware.

## Partitioning and routing

```text
hash(tenant_id) -> partition

tenant A ----+             +--> P0
tenant B ----+--> router --+--> P1
tenant C ----+             +--> P2
```

Hash partitioning spreads ordinary keys but cannot split one huge tenant. Range partitioning supports scans but risks
hot recent ranges. Directory-based routing isolates large tenants and residency but adds control-plane state. A common
evolution is shared hash partitions for most tenants plus dedicated placements for outliers.

The partition key must match the strongest locality requirement. An account key preserves per-account order; a random
event key maximizes spread but makes ordered stateful evaluation expensive.

## Replication and consistency

```text
client -> leader -> replica A
                  -> replica B
             ^
             +-- acknowledge after required quorum
```

Leader replication gives a clear write order but creates leader failover. Leaderless quorums can remain writable under
some failures but expose conflicts and read repair. "Eventually consistent" is incomplete: state the convergence
mechanism, staleness bound, conflict rule, and which invariant may temporarily fail.

CAP applies during a network partition: a system cannot provide both linearizable availability and a response from every
isolated side. It does not mean every database is permanently only two letters.

## Delivery is not effect semantics

Most durable brokers provide at-least-once delivery. Exactly-once business effect still requires cooperation between the
consumer and the effect store.

```text
receive event
    |
    v
begin transaction
    +-- insert processed(event_id) -- duplicate? --> return success
    +-- apply domain mutation
    +-- insert outbox event
commit
    |
    v
ack input
```

If the database commit succeeds but acknowledgement is lost, redelivery finds `processed(event_id)` and avoids a second
effect. Retain deduplication records for at least the maximum replay horizon.

## Ordering

Global order is expensive and rarely necessary. Define the entity whose order matters and serialize only that key.
Sequence numbers detect gaps; producer epochs fence an old writer after failover.

```text
key A: 41 -> 42 -> 43
key B:  8 ->  9

No claim is made about A:42 versus B:8.
```

Reject strict ordering if operations commute or version comparison is enough. It reduces parallelism and creates hot
partitions.

## Time

Wall clocks can jump and differ across hosts. Use:

- monotonic time for local durations and deadlines;
- UTC wall time for human records;
- database versions or logical sequence numbers for ordering;
- bounded lateness and watermarks for event-time windows.

## Retries and overload

Retry only transient, idempotent operations. Use exponential backoff, jitter, a deadline, and a finite attempt budget.
Unbounded retries amplify failure:

```text
dependency slows -> callers time out -> retries multiply -> dependency slows further
       ^                                                    |
       +-------------------- retry storm -------------------+
```

Bound queues and concurrency. Shed optional work before critical work, and surface retry-after guidance. Reject a
circuit breaker for domain validation failures; those are deterministic, not dependency instability.

## Checkpoint

For one write path, explain the acknowledgement point, durable copies, duplicate behaviour, ordering scope, partition
response, and reconciliation method.
