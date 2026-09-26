---
tldr: "Provides interview arithmetic for throughput, storage, bandwidth, concurrency, partitions, and headroom."
when_to_use: "Use after requirements to expose the scale and the components that need a deeper design."
---

# Estimation and Capacity

Estimate to choose an architecture class, not to demonstrate false precision. State assumptions, show units, and round.
A tenfold error should not invalidate the design.

## Workload worksheet

Suppose a platform receives 300 million events per day, peaks at five times average, stores a 1.5 KiB canonical event,
retains searchable data for 30 days, and retains compressed evidence for seven years.

```text
average writes/s = 300,000,000 / 86,400 ~= 3,500
peak writes/s    = 3,500 * 5             ~= 17,500
raw/day          = 300,000,000 * 1.5 KiB ~= 450 GB
30-day raw       = 450 GB * 30           ~= 13.5 TB
```

Now add replicas, indexes, log overhead, temporary compaction space, and headroom. A rough factor of three to six is
often more honest than reporting raw payload size as provisioned storage.

## Little's Law

For a stable system:

```text
concurrency = arrival rate * average time in system
```

At 17,500 requests/s and 80 ms average service time:

```text
17,500/s * 0.080 s = 1,400 in-flight requests
```

This informs connection pools, worker concurrency, queue limits, and memory. Tail latency needs additional headroom; an
average is not a safe p99 capacity target.

## Bandwidth

If each request is 1.5 KiB and replication sends it to three copies:

```text
17,500/s * 1.5 KiB ~= 26 MiB/s ingress
replicated payload ~= 79 MiB/s before protocol and index overhead
```

Cross-region synchronous replication adds both bandwidth cost and wide-area latency. Only pay it when the RPO and
failover semantics require it.

## Partition estimate

If one tested consumer partition sustains 2,000 events/s at the required processing cost:

```text
minimum partitions = ceil(17,500 / 2,000) = 9
planned partitions = 18
```

The planned count allows failures and growth. But excessive partitions increase metadata, open files, rebalances, and
operating complexity. Benchmark the real payload and handler; vendor maximums are not workload guarantees.

## Queue backlog and recovery

A consumer handling 15,000 events/s during a 17,500 events/s peak accumulates 2,500 events/s. In 20 minutes:

```text
backlog = 2,500 * 1,200 = 3,000,000 events
```

After traffic returns to 10,000/s, a 15,000/s consumer has 5,000/s spare and needs about ten minutes to catch up. Track
both backlog count and oldest-message age; count alone hides whether a tenant or partition is stuck.

## Capacity envelope

```text
normal <= tested steady state < autoscale ceiling < overload limit
                                                  |
                                                  +-> shed optional work
```

Autoscaling is not instantaneous. Keep warm capacity for startup time, model loading, and node provisioning. Reject
CPU-only scaling when the bottleneck is queue depth, database connections, or an external quota.

## Checkpoint

For any case, calculate peak requests/s, in-flight concurrency, daily raw storage, replication/index factor, and backlog
recovery time. Identify the assumption that most changes the design.
