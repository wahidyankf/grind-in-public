---
tldr: "Solves a rolling breach and effective-version lookup while keeping preprocessing costs explicit."
when_to_use: "Use after attempting drill 002 without the reference implementation."
---

# Windows and Search Solutions

## Rolling threshold breach

```python
from collections import deque
from collections.abc import Iterable


def first_breach(timestamps: Iterable[int], window_seconds: int, threshold: int) -> int | None:
    if window_seconds <= 0:
        raise ValueError("window_seconds must be positive")
    if threshold <= 0:
        raise ValueError("threshold must be positive")

    active: deque[int] = deque()
    previous: int | None = None
    for timestamp in timestamps:
        if previous is not None and timestamp < previous:
            raise ValueError("timestamps must be non-decreasing")
        previous = timestamp
        cutoff = timestamp - window_seconds
        while active and active[0] <= cutoff:
            active.popleft()
        active.append(timestamp)
        if len(active) >= threshold:
            return timestamp
    return None


assert first_breach([10, 20, 30, 80], window_seconds=60, threshold=3) == 30
assert first_breach([10, 80], window_seconds=60, threshold=2) is None
```

The deque contains exactly the timestamps inside the current window. Every value enters and leaves once, so time is
`O(n)` and space is `O(w)` for the busiest window.

## Effective version lookup

```python
from bisect import bisect_right
from collections.abc import Sequence
from dataclasses import dataclass


@dataclass(frozen=True)
class RuleVersion:
    version: str
    effective_from: int


@dataclass(frozen=True)
class RuleTimeline:
    versions: tuple[RuleVersion, ...]
    starts: tuple[int, ...]

    @classmethod
    def build(cls, versions: Sequence[RuleVersion]) -> "RuleTimeline":
        copied = tuple(versions)
        starts = tuple(item.effective_from for item in copied)
        if not copied:
            raise ValueError("at least one version is required")
        if any(left >= right for left, right in zip(starts, starts[1:], strict=False)):
            raise ValueError("effective_from values must be strictly increasing")
        return cls(copied, starts)

    def at(self, timestamp: int) -> RuleVersion:
        index = bisect_right(self.starts, timestamp) - 1
        if index < 0:
            raise ValueError("no version is effective yet")
        return self.versions[index]


timeline = RuleTimeline.build([RuleVersion("v1", 10), RuleVersion("v2", 20)])
assert timeline.at(25) == RuleVersion("v2", 20)
```

Validation and construction cost `O(n)` once; each lookup is `O(log n)`.
