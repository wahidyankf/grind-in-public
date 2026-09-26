---
tldr: "Solves bounded breadth-first traversal and deterministic topological ordering."
when_to_use: "Use after attempting drill 004 without the reference implementation."
---

# Trees and Dependencies Solutions

## Bounded breadth-first traversal

```python
from collections import deque
from collections.abc import Mapping, Sequence


def breadth_first(children: Mapping[str, Sequence[str]], root: str, max_depth: int) -> list[str]:
    if max_depth < 0:
        raise ValueError("max_depth must not be negative")
    queue: deque[tuple[str, int]] = deque([(root, 0)])
    seen: set[str] = set()
    ordered: list[str] = []
    while queue:
        node, depth = queue.popleft()
        if node in seen:
            continue
        seen.add(node)
        ordered.append(node)
        if depth < max_depth:
            queue.extend((child, depth + 1) for child in children.get(node, ()))
    return ordered


graph = {"root": ["a", "b"], "a": ["c"], "b": ["root"]}
assert breadth_first(graph, "root", max_depth=2) == ["root", "a", "b", "c"]
```

The seen set prevents malformed cycles from looping. Time is `O(V + E)` within the reached depth.

## Deterministic feature order

```python
import heapq
from collections import defaultdict
from collections.abc import Mapping, Sequence


def feature_order(dependencies: Mapping[str, Sequence[str]]) -> list[str]:
    indegree = {name: len(required) for name, required in dependencies.items()}
    outgoing: defaultdict[str, list[str]] = defaultdict(list)
    for feature, required in dependencies.items():
        for prerequisite in required:
            if prerequisite not in indegree:
                raise ValueError(f"unknown prerequisite: {prerequisite}")
            outgoing[prerequisite].append(feature)

    ready = [name for name, degree in indegree.items() if degree == 0]
    heapq.heapify(ready)
    ordered: list[str] = []
    while ready:
        current = heapq.heappop(ready)
        ordered.append(current)
        for dependent in outgoing[current]:
            indegree[dependent] -= 1
            if indegree[dependent] == 0:
                heapq.heappush(ready, dependent)
    if len(ordered) != len(indegree):
        raise ValueError("dependency cycle")
    return ordered


assert feature_order({"amount": [], "device": [], "risk": ["amount", "device"]}) == [
    "amount",
    "device",
    "risk",
]
```

The heap adds deterministic `O(log V)` ready selection. Reproducible ordering makes result comparison, caching, and
incident analysis easier even if a production executor later runs independent nodes concurrently.
