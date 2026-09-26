---
tldr: "Explores decision trees while restoring state and pruning branches that cannot produce a valid result."
when_to_use: "Use for combinations, constrained assignments, parsing, and exhaustive search with meaningful pruning."
---

# Recursion and Backtracking

Backtracking explores one decision, recurses, and restores state before trying the next. Its invariant is not merely
“recursive”: each call owns a valid partial solution.

## Production motivation: policy test combinations

Generate all non-empty combinations of enabled signals for a small rule-simulation tool:

```python
from collections.abc import Sequence


def signal_combinations(signals: Sequence[str]) -> list[tuple[str, ...]]:
    results: list[tuple[str, ...]] = []
    selected: list[str] = []

    def visit(index: int) -> None:
        if index == len(signals):
            if selected:
                results.append(tuple(selected))
            return
        visit(index + 1)
        selected.append(signals[index])
        visit(index + 1)
        selected.pop()

    visit(0)
    return results


assert signal_combinations(["device", "velocity"]) == [
    ("velocity",),
    ("device",),
    ("device", "velocity"),
]
```

```text
index 0
+-- exclude device -> exclude/include velocity
`-- include device -> exclude/include velocity
```

The algorithm produces `2^n - 1` results, so exponential time and output are unavoidable. This is acceptable only for
small controlled inputs. Production configuration systems should impose bounds, prune invalid combinations early, or
sample rather than accidentally creating an exponential job.

Python recursion depth also limits large traversals. Use an explicit stack for untrusted or deep input.

## Checkpoint

Add a maximum combination size and explain how the pruning changes explored branches without changing valid output.
