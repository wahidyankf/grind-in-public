---
tldr:
  "Hierarchical token buckets use atomic regional state plus bounded local leases for low latency and fair overload."
when_to_use: "Use after attempting Case 002 to compare limiter algorithms and failure policies."
---

# Solution 002: Distributed Rate Limiting

## Requirement traceability

| Requirements      | Design response                                                                              |
| ----------------- | -------------------------------------------------------------------------------------------- |
| FR-1, FR-5, NFR-4 | Hierarchical global, tenant, endpoint, and priority-class buckets.                           |
| FR-2, NFR-2       | Token bucket permits bounded bursts; regional atomic updates bound healthy over-admission.   |
| FR-3              | Response carries limit, approximate remaining tokens, policy version, and retry time.        |
| FR-4, NFR-3       | Versioned policy snapshots and replicated regional limiter shards.                           |
| NFR-1, NFR-5      | Local token leases avoid a network check per request; endpoint class declares outage policy. |

## Algorithm

For capacity `C`, refill rate `r`, stored tokens `T`, last time `t0`, and monotonic current time `t`:

```text
refilled = min(C, T + r * (t - t0))
allow    = refilled >= cost
new T    = refilled - cost, if allowed; otherwise refilled
```

Token bucket matches burst-tolerant requirements. A fixed window has boundary spikes; an exact sliding log costs one
timestamp per request; a sliding counter is approximate; a leaky bucket smooths output rather than admission.

## Architecture

```text
request -> gateway Pod -> local lease cache
                           | enough tokens -> admit
                           `-> lease request -> regional shard -> atomic bucket state
                                                   ^
policy control -> signed version ------------------+
```

The shard atomically leases a small token block to one Pod. Maximum excess admission after a network partition is
bounded by outstanding leases, so lease size follows the 1% budget. Critical traffic owns a reserved bucket; bulk may
borrow unused shared capacity but cannot consume the reserve.

## Complete Python token bucket

```python
from dataclasses import dataclass


@dataclass(slots=True)
class TokenBucket:
    capacity: float
    refill_per_second: float
    tokens: float
    updated_at: float

    def try_consume(self, now: float, cost: float = 1.0) -> tuple[bool, float]:
        if now < self.updated_at:
            now = self.updated_at
        elapsed = now - self.updated_at
        self.tokens = min(self.capacity, self.tokens + elapsed * self.refill_per_second)
        self.updated_at = now
        if self.tokens >= cost:
            self.tokens -= cost
            return True, 0.0
        wait_seconds = (cost - self.tokens) / self.refill_per_second
        return False, wait_seconds


bucket = TokenBucket(capacity=10.0, refill_per_second=2.0, tokens=10.0, updated_at=100.0)
allowed, retry_after = bucket.try_consume(now=100.0, cost=3.0)
print(allowed, retry_after)  # True 0.0
```

Production shards run the state transition atomically and use a trusted server clock. The Python example explains the
math; it is not a distributed coordinator.

## Failure and operations

If the regional store is unavailable, critical write endpoints fail closed or consume only already leased reserve;
read-heavy low-risk endpoints may fail open under a hard local emergency ceiling. Policy fetch failure uses a signed
last-known-good version. Metrics include allowed/rejected by hierarchy level, lease waste, store latency, hot keys,
emergency mode, and observed over-admission.

Add TTL to inactive buckets and cardinality admission to resist random-key memory attacks. Use rendezvous hashing to
route bucket keys to shards and replicate state across zones.

## Alternatives rejected

One global cross-region counter violates the 5 ms budget and becomes a shared failure domain. Fully local counters have
unbounded aggregate over-admission as Pod count changes. Exact sliding logs are reserved for low-volume contractual
limits, not the main 100,000 checks/s path.
