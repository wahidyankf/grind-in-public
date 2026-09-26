---
tldr: "Applies deadlines, retry budgets, circuit breakers, bulkheads, canaries, and cells to bounded failures."
when_to_use: "Use when dependencies, deploys, or shared resources can spread one failure across the system."
---

# Resilience and Safe Delivery

## Deadline propagation

The caller owns the end-to-end deadline. Each hop receives the remaining budget and reserves time to handle failure.

```text
200 ms client deadline
  edge 20 -> service A 130 -> dependency B 80
                          \-> 30 ms reserved for fallback/response
```

Reject fixed independent timeouts whose sum exceeds the caller deadline; abandoned work continues consuming capacity.

## Retry budget

Retry only errors likely to succeed on another attempt and only for idempotent operations. A retry budget caps retries
as a fraction of normal traffic. Add exponential backoff and full jitter. Reject retries after overload rejection unless
the server supplies a useful delay.

## Circuit breaker and bulkhead

```text
closed -- failure threshold --> open -- cool-down --> half-open -- success --> closed
                                                    `-- failure --> open
```

A circuit breaker fails fast when a dependency is predictably unavailable. A bulkhead limits the capacity one
dependency, tenant, or workload class can occupy. Breakers can synchronize and hide recovery; use metrics and jitter.
Reject a breaker where a bounded concurrency limiter and deadline already handle the risk more predictably.

## Cell architecture

Partition tenants into self-contained cells, each with compute and data dependencies, so one cell failure affects a
bounded population. A global control plane assigns tenants and distributes signed configuration.

```text
control plane -> directory/config
                   |
        +----------+----------+
        v                     v
   cell A stack            cell B stack
```

Cells increase fleet management, capacity fragmentation, and cross-cell query complexity. Adopt them when blast-radius
or residency requirements justify that cost.

## Progressive delivery

Release to internal traffic, then a small tenant/traffic cohort, then expand. Automated analysis must include domain
outcomes, not only CPU and `5xx`. Keep rollback safe by retaining compatible schemas and old consumers through the
observation window.

## Chaos and recovery exercises

Test stated failures: kill a Pod, deny dependency traffic, exhaust a pool, delay a partition, restore a backup, and fail
a region in a controlled environment. An exercise succeeds when detection, mitigation, data integrity, and human
coordination match the runbook—not merely when Kubernetes replaces a Pod.
