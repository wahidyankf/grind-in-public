"""Consistent hashing with virtual nodes.

Used in: sharding keys across cache and database nodes (Dynamo, Cassandra, Memcached clients, CDN routing). When a node
joins or leaves, only the keys on its arcs of the ring move.

Complexity: get_node is O(log v) for v virtual nodes, using binary search. add_node and remove_node are O(v log v).

Hashes come from hashlib, not the built-in hash(), because built-in hash() changes between Python processes.
"""

import hashlib
from bisect import bisect_right


DEFAULT_VIRTUAL_NODES = 100


def _ring_hash(text: str) -> int:
    return int.from_bytes(hashlib.sha256(text.encode()).digest()[:8], "big")


class ConsistentHashRing:
    def __init__(self, virtual_nodes: int = DEFAULT_VIRTUAL_NODES) -> None:
        if virtual_nodes < 1:
            raise ValueError("virtual_nodes must be at least 1")
        self._virtual_nodes = virtual_nodes
        self._points: list[int] = []
        self._owner: dict[int, str] = {}

    def __len__(self) -> int:
        return len(set(self._owner.values()))

    def add_node(self, name: str) -> None:
        for i in range(self._virtual_nodes):
            point = _ring_hash(f"{name}#{i}")
            if point not in self._owner:
                self._owner[point] = name
                self._points.append(point)
        self._points.sort()

    def remove_node(self, name: str) -> None:
        kept = {point: owner for point, owner in self._owner.items() if owner != name}
        self._owner = kept
        self._points = sorted(kept)

    def get_node(self, key: str) -> str:
        if not self._points:
            raise LookupError("the ring has no nodes")
        index = bisect_right(self._points, _ring_hash(key)) % len(self._points)
        return self._owner[self._points[index]]
