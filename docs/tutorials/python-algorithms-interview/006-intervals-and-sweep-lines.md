---
tldr: "Models spans and ordered boundary events to merge coverage, detect overlap, and measure concurrency."
when_to_use: "Use when starts and ends define occupancy, availability, reservations, or time-based conflicts."
---

# Intervals and Sweep Lines

Intervals represent occupancy over a domain: time, identifiers, or numeric ranges. Before coding, declare whether an
endpoint is inclusive. Half-open intervals `[start, end)` compose well because adjacent spans do not overlap.

## Production motivation: peak concurrent investigations

Turn each case assignment into a `+1` start event and `-1` end event. For half-open intervals, process an end before a
start at the same timestamp.

```python
from collections.abc import Iterable


def peak_concurrency(intervals: Iterable[tuple[int, int]]) -> int:
    events: list[tuple[int, int]] = []
    for start, end in intervals:
        if start >= end:
            raise ValueError("interval must have positive duration")
        events.extend(((start, 1), (end, -1)))

    active = 0
    peak = 0
    for _, change in sorted(events, key=lambda event: (event[0], event[1])):
        active += change
        peak = max(peak, active)
    return peak


assert peak_concurrency([(1, 4), (2, 5), (4, 6)]) == 2
```

```text
time      1   2   4   4   5   6
change   +1  +1  -1  +1  -1  -1
active    1   2   1   2   1   0
```

Sorting dominates at `O(n log n)`; the sweep is `O(n)`. If events already arrive ordered, process them as a stream.

## Production constraints

Calendar bookings, rate windows, and database validity ranges may use different endpoint rules. Distributed event time
also introduces late arrivals: a correct online sweep may need watermarks and corrections. Use database range types or
window functions when persistence and concurrent updates are the real boundary.

## Checkpoint

Change the contract to closed intervals and explain why the tie-breaking order must reverse.
