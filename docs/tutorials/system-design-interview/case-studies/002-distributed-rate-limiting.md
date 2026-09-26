---
tldr: "Design hierarchical rate limiting for global safety, tenant fairness, and burst tolerance."
when_to_use: "Use to practise token buckets, atomic state, approximate local leases, and failure policy."
---

# Case 002: Distributed Rate Limiting

Design rate limiting for public APIs and internal expensive operations across many Kubernetes replicas. Policies vary by
tenant, endpoint class, and purchased capacity.

### Functional requirements

| ID   | Requirement                                                                       |
| ---- | --------------------------------------------------------------------------------- |
| FR-1 | Evaluate global, tenant, and endpoint limits for each authenticated request.      |
| FR-2 | Permit configured bursts while enforcing a sustained rate.                        |
| FR-3 | Return remaining capacity and a retry-after hint on rejection.                    |
| FR-4 | Publish policy changes with versioned audit history and bounded propagation time. |
| FR-5 | Let critical decision traffic retain capacity when bulk traffic is saturated.     |

### Non-functional requirements

| ID    | Requirement                                                                                   |
| ----- | --------------------------------------------------------------------------------------------- |
| NFR-1 | Add at most 5 ms p99 in-region at 100,000 checks/s.                                           |
| NFR-2 | Keep over-admission below 1% during healthy operation; never promise an exact sliding window. |
| NFR-3 | Remain available during one limiter-node failure without resetting every bucket.              |
| NFR-4 | Bound tenant blast radius and prevent a hot tenant from exhausting limiter state.             |
| NFR-5 | Make the fail-open/fail-closed policy explicit for each endpoint class.                       |

## Assumptions and exclusions

Identity is already authenticated. Billing settlement and volumetric network attack protection are outside scope.

## Interview prompts

1. Compare fixed window, sliding log, sliding counters, token bucket, and leaky bucket.
2. Where is atomic bucket state kept, and what happens during partition?
3. Can Pods receive short local token leases without violating the over-admission budget?
4. How are policy rollout, clock behaviour, hot keys, and observability handled?

Solve before reading [the worked solution](../solutions/002-distributed-rate-limiting.md).
