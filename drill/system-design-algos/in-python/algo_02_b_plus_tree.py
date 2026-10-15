"""B+ tree: an ordered index where all values live in linked leaves.

Used in: database indexes. PostgreSQL and MySQL InnoDB use this shape for most of their indexes.

Complexity: search and insert are O(log n) node visits. A range scan costs O(log n) to find the start, then O(k) for k
results, because the leaves are linked.

This version stores each node's entries in plain lists and splits full nodes. It does not handle deletes or disk pages.
"""

from bisect import bisect_left, bisect_right
from typing import Protocol, Self


class SupportsLt(Protocol):
    def __lt__(self, other: Self, /) -> bool: ...


class _Leaf[K, V]:
    def __init__(self) -> None:
        self.keys: list[K] = []
        self.values: list[V] = []
        self.next: _Leaf[K, V] | None = None


class _Internal[K, V]:
    def __init__(self) -> None:
        self.keys: list[K] = []
        self.children: list[_Leaf[K, V] | _Internal[K, V]] = []


type _Node[K, V] = _Leaf[K, V] | _Internal[K, V]


class BPlusTree[K: SupportsLt, V]:
    def __init__(self, order: int = 4) -> None:
        if order < 3:
            raise ValueError("order must be at least 3")
        self._order = order
        self._root: _Node[K, V] = _Leaf()
        self._size = 0

    def __len__(self) -> int:
        return self._size

    def search(self, key: K) -> V | None:
        leaf = self._find_leaf(key)
        i = bisect_left(leaf.keys, key)
        if i < len(leaf.keys) and leaf.keys[i] == key:
            return leaf.values[i]
        return None

    def insert(self, key: K, value: V) -> None:
        split = self._insert(self._root, key, value)
        if split is not None:
            separator, new_node = split
            root: _Internal[K, V] = _Internal()
            root.keys = [separator]
            root.children = [self._root, new_node]
            self._root = root

    def range(self, low: K, high: K) -> list[tuple[K, V]]:
        """Return every pair with low <= key <= high, in key order."""
        result: list[tuple[K, V]] = []
        leaf: _Leaf[K, V] | None = self._find_leaf(low)
        while leaf is not None:
            for key, value in zip(leaf.keys, leaf.values):
                if high < key:
                    return result
                if not key < low:
                    result.append((key, value))
            leaf = leaf.next
        return result

    def items(self) -> list[tuple[K, V]]:
        leaf: _Leaf[K, V] | None = self._leftmost_leaf()
        result: list[tuple[K, V]] = []
        while leaf is not None:
            result.extend(zip(leaf.keys, leaf.values))
            leaf = leaf.next
        return result

    def _find_leaf(self, key: K) -> _Leaf[K, V]:
        node = self._root
        while isinstance(node, _Internal):
            node = node.children[bisect_right(node.keys, key)]
        return node

    def _leftmost_leaf(self) -> _Leaf[K, V]:
        node = self._root
        while isinstance(node, _Internal):
            node = node.children[0]
        return node

    def _insert(self, node: _Node[K, V], key: K, value: V) -> tuple[K, _Node[K, V]] | None:
        if isinstance(node, _Leaf):
            i = bisect_left(node.keys, key)
            if i < len(node.keys) and node.keys[i] == key:
                node.values[i] = value
                return None
            node.keys.insert(i, key)
            node.values.insert(i, value)
            self._size += 1
            if len(node.keys) <= self._order:
                return None
            return self._split_leaf(node)

        i = bisect_right(node.keys, key)
        split = self._insert(node.children[i], key, value)
        if split is None:
            return None
        separator, new_child = split
        node.keys.insert(i, separator)
        node.children.insert(i + 1, new_child)
        if len(node.children) <= self._order:
            return None
        return self._split_internal(node)

    def _split_leaf(self, leaf: _Leaf[K, V]) -> tuple[K, _Leaf[K, V]]:
        mid = len(leaf.keys) // 2
        right: _Leaf[K, V] = _Leaf()
        right.keys = leaf.keys[mid:]
        right.values = leaf.values[mid:]
        right.next = leaf.next
        leaf.keys = leaf.keys[:mid]
        leaf.values = leaf.values[:mid]
        leaf.next = right
        # The first key of the right leaf is copied up. Searches send equal keys to the right, so it stays findable.
        return right.keys[0], right

    def _split_internal(self, node: _Internal[K, V]) -> tuple[K, _Internal[K, V]]:
        mid = len(node.children) // 2
        separator = node.keys[mid - 1]
        right: _Internal[K, V] = _Internal()
        right.keys = node.keys[mid:]
        right.children = node.children[mid:]
        node.keys = node.keys[: mid - 1]
        node.children = node.children[:mid]
        return separator, right
