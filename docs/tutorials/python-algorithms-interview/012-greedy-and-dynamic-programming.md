---
tldr: "Distinguishes locally safe choices from problems that require storing overlapping subproblem results."
when_to_use: "Use after defining optimal substructure, state, transitions, and the proof obligation."
---

# Greedy and Dynamic Programming

A greedy algorithm commits to a locally best choice and needs an exchange-style proof that the choice cannot harm the
optimum. Dynamic programming stores overlapping subproblem results when future choices depend on earlier state.

## Production motivation: minimum review capacity

For interval scheduling, choosing the assignment that ends earliest leaves the most room for later work and is provably
optimal for the maximum number of non-overlapping intervals.

```python
from collections.abc import Iterable


def maximum_non_overlapping(intervals: Iterable[tuple[int, int]]) -> list[tuple[int, int]]:
    selected: list[tuple[int, int]] = []
    last_end: int | None = None
    for start, end in sorted(intervals, key=lambda interval: interval[1]):
        if start >= end:
            raise ValueError("interval must have positive duration")
        if last_end is None or start >= last_end:
            selected.append((start, end))
            last_end = end
    return selected


assert maximum_non_overlapping([(1, 4), (2, 3), (3, 5), (5, 7)]) == [(2, 3), (3, 5), (5, 7)]
```

```text
choose earliest end: (2,3) -> (3,5) -> (5,7)
reject overlap:      (1,4)
```

Sorting costs `O(n log n)`; selection is linear. The exchange argument replaces the first interval of any optimal
solution with the earliest-ending interval without reducing remaining capacity.

Dynamic programming instead requires explicit state, recurrence, base cases, and evaluation order. Do not label a
solution “DP” merely because it has a table. Production scheduling often adds weights, fairness, skills, and changing
availability; those constraints can turn a greedy solution into weighted interval scheduling or an optimization model.

## Checkpoint

Explain why choosing the shortest interval is not the same greedy rule, and give a counterexample.
