---
tldr: "Reuses ordered structure with two pointers, sliding windows, and prefix sums instead of rescanning input."
when_to_use: "Use for contiguous ranges, pair searches, rolling limits, and repeated range totals."
---

# Pointers, Windows, and Prefix Sums

These patterns turn repeated scans into maintained state. Two pointers exploit order, a sliding window tracks one
contiguous range, and prefix sums precompute cumulative values for repeated range queries.

## Production motivation: rolling velocity

A service must count events in the last 60 seconds for one account. If timestamps arrive in non-decreasing order, a
deque holds exactly the active window:

```python
from collections import deque
from dataclasses import dataclass, field


@dataclass
class VelocityWindow:
    seconds: int
    _timestamps: deque[int] = field(default_factory=deque)

    def add(self, timestamp: int) -> int:
        cutoff = timestamp - self.seconds
        while self._timestamps and self._timestamps[0] <= cutoff:
            self._timestamps.popleft()
        self._timestamps.append(timestamp)
        return len(self._timestamps)


window = VelocityWindow(seconds=60)
assert [window.add(value) for value in (10, 20, 69, 70)] == [1, 2, 3, 3]
```

```text
t=10  [10]
t=20  [10,20]
t=69  [10,20,69] -> cutoff=9, so 10 remains
t=70  [10 expires] [20,69,70]
```

Each timestamp enters and leaves once, giving amortized `O(1)` work and `O(w)` memory for the busiest window. The
invariant is that every stored timestamp is within `(current - seconds, current]`.

Production streams complicate the assumption: events can arrive late, be duplicated, or be processed by different
partitions. The system-design course adds event-time watermarks, idempotency, partition keys, and durable state.

## Prefix sums

For immutable daily totals, build `prefix[i + 1] = prefix[i] + value[i]`. A range sum from `left` through `right`
becomes `prefix[right + 1] - prefix[left]`: `O(n)` preprocessing, `O(1)` per query, and `O(n)` extra space.

## When not to use the pattern

Sliding windows need a well-defined order and eviction boundary. Prefix sums become expensive when values change
frequently. A database window function, time-series store, or streaming state engine may own the production problem
better than application memory.

## Checkpoint

Explain how the example changes if timestamps arrive out of order and the result may wait up to five seconds for late
events.
