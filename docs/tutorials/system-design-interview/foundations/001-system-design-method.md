---
tldr: "Frames system design as explicit reasoning from requirements to trade-offs, failure handling, and evolution."
when_to_use: "Use at the start of preparation or when an interview answer becomes a list of technologies."
---

# System Design as a Method

System design is the practice of choosing boundaries and mechanisms that satisfy known requirements under incomplete
information. The interview tests whether you can reduce ambiguity, quantify the load, find the risky parts, and explain
trade-offs. It does not test whether you can draw every component used by a large company.

## The reasoning chain

```text
 business outcome
       |
       v
+----------------+     +-------------+     +---------------+
| FR and NFR     | --> | estimates   | --> | contracts     |
+----------------+     +-------------+     +---------------+
       |                                          |
       v                                          v
+----------------+     +-------------+     +---------------+
| failure model  | --> | components  | --> | verification  |
+----------------+     +-------------+     +---------------+
```

- Functional requirements (`FR-*`) say what actors can do.
- Non-functional requirements (`NFR-*`) quantify latency, throughput, availability, durability, security, recovery, and
  operability.
- Constraints say what cannot change now: residency, an existing database, Kubernetes, or a migration boundary.
- Out-of-scope items protect the design from silently expanding.

Each component must earn its place. If `NFR-2` needs asynchronous processing after a 50 ms response, a durable queue may
be justified. "Modern systems use queues" is not justification.

## Start with a walking skeleton

For a new service, first draw the smallest end-to-end path:

```text
client -> API -> application -> primary database
```

Then add a component only when a requirement creates pressure:

```text
high read repetition  -> cache
slow optional work    -> queue + worker
full-text retrieval   -> search index
immutable replay      -> event log
independent scaling   -> separate runtime boundary
```

This prevents premature distribution. A single deployable with strong internal module boundaries is often the safest
initial design: local calls are fast, transactions are simple, and debugging crosses fewer processes.

## Breadth, then depth

Spend the middle of the interview on two risky paths rather than narrating every box. Good deep dives include:

- deduplication and ordering in ingestion;
- hot-key control in a rate limiter;
- rule and model versioning in decisioning;
- tenant isolation and evidence integrity;
- backfill, reconciliation, and rollback during extraction.

For each, state the normal path, degraded path, retry behaviour, observability, and owner response.

## Reject unjustified complexity

Do not introduce microservices merely because the expected organization is large. Reject the split when the proposed
boundary has no independent scaling, security, release, ownership, or availability requirement. Distribution adds
network failure, version skew, partial success, duplicated data, and operational load.

## Checkpoint

Take a product you know and explain one component using this sentence:

> Because `NFR-x` requires ___ under ___ load, use ___, accepting ___; reject it if ___.

If a clause is missing, the choice is not grounded yet.
