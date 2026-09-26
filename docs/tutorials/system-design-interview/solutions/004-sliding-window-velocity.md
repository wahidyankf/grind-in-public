---
tldr: "Keyed stream state uses exact short-window buckets and mergeable approximate long-window distinct sketches."
when_to_use: "Use after attempting Case 004 to compare window algorithms, lateness, checkpoints, and hot-key handling."
---

# Solution 004: Sliding-Window Velocity

## Requirement traceability

| Requirements | Design response                                                                                 |
| ------------ | ----------------------------------------------------------------------------------------------- |
| FR-1, NFR-2  | Durable source partitioned by tenant/entity; inbox deduplication and checkpointed keyed state.  |
| FR-2, NFR-1  | Materialized feature rows serve queries; stream workers update time buckets incrementally.      |
| FR-3         | Event-time watermarks hold ten-minute lateness and reopen affected buckets before finalization. |
| FR-4, NFR-5  | Tenant/range replay into a shadow namespace, then verified pointer swap.                        |
| FR-5, NFR-4  | Metadata marks freshness and algorithm; short count/sum exact, long distinct approximate.       |
| NFR-3        | Tiered windows, state TTL, hot-key sub-buckets, and isolated large tenants bound state.         |

## Architecture

```text
durable events -> keyed stream workers -> checkpoint/state store -> serving KV
                         |                       |
                         +-> late/correction ----+
query API -------------------------------------> feature row
```

Use one-second buckets for one minute, one-minute buckets for one hour/day, and one-hour buckets for 30 days. Each
bucket stores count, decimal/minor-unit sum, max, and a mergeable distinct sketch. Exact short-window distinct can use a
counted set only when measured cardinality fits memory.

```text
now
 |  [s][s][s]... 60       exact 1-minute ring
 |  [ minute buckets ]    exact count/sum/max
 `--[ hour sketches ]     approximate 30-day distinct
```

## Complete exact rolling-count core

```python
from collections import deque
from dataclasses import dataclass, field


@dataclass(slots=True)
class RollingCount:
    window_seconds: int
    timestamps: deque[int] = field(default_factory=deque)

    def add(self, event_second: int) -> None:
        self.timestamps.append(event_second)

    def count(self, now_second: int) -> int:
        first_in_window = now_second - self.window_seconds + 1
        while self.timestamps and self.timestamps[0] < first_in_window:
            self.timestamps.popleft()
        return len(self.timestamps)


counter = RollingCount(window_seconds=60)
for second in (100, 120, 159, 160):
    counter.add(second)
print(counter.count(now_second=160))  # 3: seconds 120, 159, and 160
```

The deque is `O(1)` amortized and exact, but storing every event is too costly for 100 million keys and 30 days. Rings
bound memory by bucket count; a monotonic deque can maintain rolling maximum when fine-grained exact max is required.

## Lateness and query semantics

Workers advance a watermark from observed partition progress minus ten minutes. Events behind the watermark go to an
offline correction stream. A query returns `computed_through`, watermark, and exactness. Previously executed decisions
remain immutable; correction affects later features and produces audit lineage.

## Recovery and skew

Checkpoint state and source offsets atomically from the stream processor's perspective. On restart, restore checkpoint
then replay. For hot entities, route subkeys to partial aggregators and merge only algorithms that are associative; if
strict per-entity order is required, isolate the key on a larger worker instead.

## Alternatives rejected

A database query over raw events cannot meet 15 ms at this scale. One deque per key violates memory bounds. Prefix sums
are excellent for immutable batch arrays but expensive to update for a live sliding horizon.
