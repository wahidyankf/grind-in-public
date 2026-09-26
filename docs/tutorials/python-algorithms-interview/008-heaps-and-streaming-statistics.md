---
tldr: "Maintains the most urgent or extreme items without fully sorting every value."
when_to_use: "Use for priority scheduling, top-K queries, merged streams, or bounded running statistics."
---

# Heaps and Streaming Statistics

A heap keeps its smallest item at the root while leaving the rest partially ordered. It is ideal when the next item or
the best `k` items matter, but a fully sorted collection does not.

## Production motivation: bounded alert queue

Keep the three highest scores with a min-heap. The root is the weakest retained alert.

```python
import heapq
from collections.abc import Iterable


def top_scores(scores: Iterable[int], limit: int) -> list[int]:
    if limit < 1:
        raise ValueError("limit must be positive")
    selected: list[int] = []
    for score in scores:
        if len(selected) < limit:
            heapq.heappush(selected, score)
        elif score > selected[0]:
            heapq.heapreplace(selected, score)
    return sorted(selected, reverse=True)


assert top_scores([40, 90, 20, 75, 95], 3) == [95, 90, 75]
```

```text
40 -> [40]
90 -> [40,90]
20 -> [20,90,40]
75 -> replace 20 -> [40,90,75]
95 -> replace 40 -> [75,90,95]
```

The algorithm uses `O(n log k)` time and `O(k)` memory. The invariant is that the heap contains the highest `k` values
from the consumed prefix.

## Production limits

A score alone is not a stable queue contract. Real alert routing needs deterministic tie-breakers, tenant fairness,
aging, persistence, cancellation, and concurrent claims. A database index or durable broker may be the source of truth;
the heap can remain a local scheduling optimization.

Two heaps can maintain an exact running median, but approximate quantile sketches are usually preferable at large
distributed scale because merging exact heaps across partitions is expensive.

## Checkpoint

Add a sequence number to make equal scores deterministic, then state which ordering the operator will observe.
