---
tldr: "Uses order to search boundaries, establish deterministic processing, and select ranked values efficiently."
when_to_use: "Use when input is ordered or can be ordered and the answer has a monotonic boundary."
---

# Binary Search, Sorting, and Selection

Binary search is not merely “find a number.” It locates a boundary in a monotonic predicate. Sorting pays an upfront
cost to make later search, merging, grouping, or deterministic processing cheaper.

## Production motivation: threshold lookup

A rule table maps ascending score thresholds to actions. Find the first threshold greater than the score, then choose
the previous action:

```python
from bisect import bisect_right
from collections.abc import Sequence
from dataclasses import dataclass


@dataclass(frozen=True)
class Threshold:
    minimum: int
    action: str


def action_for_score(rules: Sequence[Threshold], score: int) -> str:
    """Return the action for sorted, gap-free thresholds starting at zero."""
    positions = [rule.minimum for rule in rules]
    index = bisect_right(positions, score) - 1
    if index < 0:
        raise ValueError("score falls below the first threshold")
    return rules[index].action


rules = [Threshold(0, "accept"), Threshold(50, "review"), Threshold(80, "reject")]
assert action_for_score(rules, 67) == "review"
```

```text
thresholds: 0 -------- 50 -------- 80 -------->
score:                        67
boundary: first minimum > 67 is 80; choose previous rule
```

The search is `O(log n)`, but building `positions` inside every call is `O(n)`. Production code would validate and store
the immutable positions once when publishing the rule version. The example keeps construction visible so the cost cannot
hide behind the library call.

## Sorting and selection

Python's stable sort keeps equal-key items in their earlier order, which is valuable for deterministic tie-breaking.
Full sorting costs `O(n log n)`. If only the smallest `k` values matter, a heap can use `O(n log k)` time and `O(k)`
space. If data already lives in a database, an indexed `ORDER BY ... LIMIT` may avoid transferring irrelevant rows.

## Rejection cases

Do not use binary search on data that is not sorted under the same comparison contract. Do not sort a stream merely to
find a running maximum. Do not claim `O(log n)` for the example without including how the searchable structure is built
and updated.

## Checkpoint

Move `positions` into a validated immutable table object and state the construction, lookup, and update costs.
