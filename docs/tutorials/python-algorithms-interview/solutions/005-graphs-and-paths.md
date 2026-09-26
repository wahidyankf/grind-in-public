---
tldr: "Solves connected-component retrieval and reconstructs a lowest-cost non-negative path."
when_to_use: "Use after attempting drill 005 without the reference implementation."
---

# Graphs and Paths Solutions

## Linked component

```python
from collections import defaultdict, deque
from collections.abc import Iterable, Set


def linked_component(entities: Set[str], links: Iterable[tuple[str, str]], start: str) -> list[str]:
    if start not in entities:
        raise ValueError("unknown start entity")
    neighbours: defaultdict[str, set[str]] = defaultdict(set)
    for left, right in links:
        if left not in entities or right not in entities:
            raise ValueError("link references an unknown entity")
        neighbours[left].add(right)
        neighbours[right].add(left)

    queue = deque([start])
    seen = {start}
    while queue:
        current = queue.popleft()
        for neighbour in neighbours[current]:
            if neighbour not in seen:
                seen.add(neighbour)
                queue.append(neighbour)
    return sorted(seen)


assert linked_component({"a", "b", "c"}, [("a", "b")], "a") == ["a", "b"]
assert linked_component({"a", "b", "c"}, [("a", "b")], "c") == ["c"]
```

## Lowest-cost explanation path

```python
import heapq
from collections.abc import Mapping, Sequence


def lowest_cost_path(
    graph: Mapping[str, Sequence[tuple[str, int]]], source: str, target: str
) -> tuple[int, list[str]] | None:
    if source not in graph or target not in graph:
        raise ValueError("source and target must be graph nodes")
    for edges in graph.values():
        if any(cost < 0 for _, cost in edges):
            raise ValueError("edge costs must not be negative")

    distances = {source: 0}
    previous: dict[str, str] = {}
    pending: list[tuple[int, str]] = [(0, source)]
    while pending:
        distance, current = heapq.heappop(pending)
        if distance != distances[current]:
            continue
        if current == target:
            path = [target]
            while path[-1] != source:
                path.append(previous[path[-1]])
            path.reverse()
            return distance, path
        for neighbour, cost in graph[current]:
            candidate = distance + cost
            if candidate < distances.get(neighbour, candidate + 1):
                distances[neighbour] = candidate
                previous[neighbour] = current
                heapq.heappush(pending, (candidate, neighbour))
    return None


network = {"a": [("b", 4), ("c", 1)], "b": [], "c": [("b", 1)]}
assert lowest_cost_path(network, "a", "b") == (2, ["a", "c", "b"])
```

The traversal costs `O((V + E) log V)`. An online service still needs traversal budgets, authorization, and protection
against high-degree nodes.
