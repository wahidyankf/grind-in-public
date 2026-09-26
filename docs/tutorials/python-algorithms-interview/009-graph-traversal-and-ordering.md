---
tldr: "Traverses relationships with BFS or DFS and orders dependency graphs with topological sorting."
when_to_use: "Use when entities have arbitrary relationships, reachability, dependencies, or cycles."
---

# Graph Traversal and Ordering

Graphs model relationships that are not constrained to one parent. BFS explores by distance in unweighted edges; DFS
dives through a branch; topological sorting orders a directed acyclic dependency graph.

## Production motivation: rule dependency order

Derived features may depend on other features. Kahn's algorithm processes nodes whose prerequisites are satisfied and
detects a cycle when not every node can be emitted.

```python
from collections import defaultdict, deque
from collections.abc import Mapping, Sequence


def dependency_order(dependencies: Mapping[str, Sequence[str]]) -> list[str]:
    indegree = {name: len(required) for name, required in dependencies.items()}
    outgoing: defaultdict[str, list[str]] = defaultdict(list)
    for name, required in dependencies.items():
        for prerequisite in required:
            if prerequisite not in indegree:
                raise ValueError(f"unknown prerequisite: {prerequisite}")
            outgoing[prerequisite].append(name)

    ready = deque(sorted(name for name, degree in indegree.items() if degree == 0))
    ordered: list[str] = []
    while ready:
        current = ready.popleft()
        ordered.append(current)
        for dependent in sorted(outgoing[current]):
            indegree[dependent] -= 1
            if indegree[dependent] == 0:
                ready.append(dependent)

    if len(ordered) != len(indegree):
        raise ValueError("dependency cycle")
    return ordered


assert dependency_order({"amount": [], "velocity": ["amount"], "risk": ["velocity"]}) == [
    "amount",
    "velocity",
    "risk",
]
```

```text
amount -> velocity -> risk
ready     emitted     newly ready
amount    amount      velocity
velocity  velocity    risk
```

Time and space are `O(V + E)`. The invariant is that every ready node has zero unmet prerequisites.

## Production limits

Large graphs need pagination, partition-aware traversal, authorization, and protection against supernodes. A graph
database makes certain traversals convenient but does not remove worst-case expansion. Precompute bounded features when
an online decision path cannot afford arbitrary traversal.

## Checkpoint

Explain why BFS finds the fewest edges in an unweighted graph but not the cheapest path in a weighted graph.
