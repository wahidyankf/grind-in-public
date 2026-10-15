"""Graph search: BFS, topological order, and Dijkstra.

Used in: routing and shortest paths (maps, network routing), social graphs ("friends of friends"), and build or
dependency ordering (topological sort).

Complexity: BFS and topological sort are O(V + E). Dijkstra with a binary heap is O((V + E) log V).
"""

import heapq
from collections import deque


class Graph:
    """A directed graph. Edges carry a weight, which defaults to 1."""

    def __init__(self) -> None:
        self._edges: dict[str, dict[str, float]] = {}

    def add_node(self, node: str) -> None:
        self._edges.setdefault(node, {})

    def add_edge(self, source: str, target: str, weight: float = 1.0) -> None:
        self.add_node(source)
        self.add_node(target)
        self._edges[source][target] = weight

    def nodes(self) -> list[str]:
        return list(self._edges)

    def neighbours(self, node: str) -> dict[str, float]:
        return dict(self._edges[node])


def bfs_shortest_path(graph: Graph, start: str, goal: str) -> list[str] | None:
    """Shortest path by number of edges, or None when the goal cannot be reached."""
    previous: dict[str, str | None] = {start: None}
    queue: deque[str] = deque([start])
    while queue:
        node = queue.popleft()
        if node == goal:
            path: list[str] = []
            current: str | None = node
            while current is not None:
                path.append(current)
                current = previous[current]
            path.reverse()
            return path
        for neighbour in graph.neighbours(node):
            if neighbour not in previous:
                previous[neighbour] = node
                queue.append(neighbour)
    return None


def topological_order(graph: Graph) -> list[str]:
    """Order the nodes so every edge points forward. Raises ValueError if the graph has a cycle."""
    # 1 means on the current DFS path, 2 means finished.
    state: dict[str, int] = {}
    order: list[str] = []

    def visit(node: str) -> None:
        if state.get(node) == 1:
            raise ValueError(f"cycle through {node}")
        if node in state:
            return
        state[node] = 1
        for neighbour in graph.neighbours(node):
            visit(neighbour)
        state[node] = 2
        order.append(node)

    for node in graph.nodes():
        visit(node)
    order.reverse()
    return order


def dijkstra(graph: Graph, start: str) -> dict[str, float]:
    """Cheapest total weight from start to every reachable node. Weights must not be negative."""
    distance: dict[str, float] = {start: 0.0}
    heap: list[tuple[float, str]] = [(0.0, start)]
    while heap:
        cost, node = heapq.heappop(heap)
        if cost > distance[node]:
            continue
        for neighbour, weight in graph.neighbours(node).items():
            candidate = cost + weight
            if candidate < distance.get(neighbour, float("inf")):
                distance[neighbour] = candidate
                heapq.heappush(heap, (candidate, neighbour))
    return distance
