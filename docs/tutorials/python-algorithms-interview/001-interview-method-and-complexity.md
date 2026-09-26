---
tldr: "Turns an ambiguous algorithm prompt into a contract, invariant, complexity target, and tested Python solution."
when_to_use: "Use before solving any algorithm exercise or explaining why an implementation is correct."
---

# Interview Method and Complexity

An algorithm interview is a compressed design review. The result matters, but so do the questions that establish what
“correct” means. Start with the observable contract:

1. Name the input, output, and mutation policy.
2. Ask about empty input, duplicates, ordering, bounds, and invalid values.
3. State a simple solution and its cost.
4. Identify the operation that dominates that cost.
5. Choose a structure that makes that operation cheaper.
6. State an invariant before writing code.
7. Test ordinary, boundary, and adversarial cases.

## Production motivation

Suppose an ingestion worker must detect duplicate event identifiers within one batch. A nested comparison works, but its
quadratic growth turns a harmless validation step into the batch bottleneck. A set stores the identifiers already seen,
changing the repeated operation from a scan to average constant-time membership.

```python
from collections.abc import Sequence


def first_duplicate(values: Sequence[str]) -> str | None:
    """Return the first repeated value in encounter order."""
    seen: set[str] = set()
    for value in values:
        if value in seen:
            return value
        seen.add(value)
    return None


assert first_duplicate(["e1", "e2", "e1"]) == "e1"
assert first_duplicate([]) is None
assert first_duplicate(["e1", "e2"]) is None
```

The invariant is: before each iteration, `seen` contains exactly the earlier values. Therefore the first membership hit
is the first duplicate in encounter order.

```text
value   seen before   decision
e1      {}            add e1
e2      {e1}          add e2
e1      {e1,e2}       return e1
```

## Complexity that communicates

Use variables that name the input. For `n` identifiers, the function takes `O(n)` expected time and `O(n)` additional
space. Say “expected” because Python's hash-table operations are not a worst-case constant-time guarantee.

Complexity describes growth, not actual latency. Production decisions also consider object size, cache locality,
serialization, concurrency, and whether all data fits in memory. A database uniqueness constraint is still required when
duplicates may arrive across batches or workers; this in-memory set solves only the local contract.

## Checkpoint

Explain why sorting first gives `O(n log n)` time, why it may change encounter order, and when its lower auxiliary
memory could still be preferable.
