---
tldr: "Maintains connectivity with union-find and chooses minimum-cost routes with weighted shortest paths."
when_to_use: "Use for components, clustering edges, cycle detection, and non-negative weighted paths."
---

# Connectivity and Shortest Paths

Union-find answers whether nodes belong to the same evolving component. Dijkstra's algorithm finds minimum-cost paths
when edge weights are non-negative.

## Production motivation: linked-account components

Merge accounts that share a verified device. Path compression and union by size keep operations nearly constant in
practice.

```python
from dataclasses import dataclass


@dataclass
class DisjointSet:
    parent: list[int]
    size: list[int]

    @classmethod
    def create(cls, count: int) -> "DisjointSet":
        return cls(parent=list(range(count)), size=[1] * count)

    def find(self, item: int) -> int:
        while item != self.parent[item]:
            self.parent[item] = self.parent[self.parent[item]]
            item = self.parent[item]
        return item

    def union(self, left: int, right: int) -> None:
        left_root = self.find(left)
        right_root = self.find(right)
        if left_root == right_root:
            return
        if self.size[left_root] < self.size[right_root]:
            left_root, right_root = right_root, left_root
        self.parent[right_root] = left_root
        self.size[left_root] += self.size[right_root]


groups = DisjointSet.create(4)
groups.union(0, 1)
groups.union(2, 3)
assert groups.find(0) == groups.find(1)
assert groups.find(0) != groups.find(2)
```

```text
before: 0  1  2  3
union:  0--1  2--3
query:  root(0)=0, root(1)=0
```

Union-find does not support cheap deletions or explain the path of evidence. Production investigations usually need the
edges and timestamps as well as the component identifier.

Dijkstra uses a heap and costs `O((V + E) log V)` for adjacency lists. Do not use it with negative weights; do not run
an unbounded graph search inside a strict-latency request without limits or precomputation.

## Checkpoint

Describe how expiring a device link breaks the union-find model and which recomputation strategy could restore correct
components.
