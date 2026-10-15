"""LRU cache: hash map plus doubly linked list.

Used in: application caches, CDN edge caches, and the eviction policy of many key-value stores.

Complexity: get and put are O(1). Eviction removes the tail of the list, so it is O(1) too.

The map finds a node in O(1). The list keeps nodes in recency order, with the most recent at the head.
"""


class _Node[K, V]:
    def __init__(self, key: K, value: V) -> None:
        self.key = key
        self.value = value
        self.prev: _Node[K, V] | None = None
        self.next: _Node[K, V] | None = None


class LRUCache[K, V]:
    def __init__(self, capacity: int) -> None:
        if capacity < 1:
            raise ValueError("capacity must be at least 1")
        self._capacity = capacity
        self._index: dict[K, _Node[K, V]] = {}
        self._head: _Node[K, V] | None = None
        self._tail: _Node[K, V] | None = None

    def __len__(self) -> int:
        return len(self._index)

    def get(self, key: K) -> V | None:
        """Return the value and mark it most recent. Returns None on a miss."""
        node = self._index.get(key)
        if node is None:
            return None
        self._move_to_front(node)
        return node.value

    def put(self, key: K, value: V) -> None:
        node = self._index.get(key)
        if node is not None:
            node.value = value
            self._move_to_front(node)
            return
        if len(self._index) >= self._capacity:
            self._evict_least_recent()
        node = _Node(key, value)
        self._index[key] = node
        self._push_front(node)

    def _evict_least_recent(self) -> None:
        victim = self._tail
        if victim is None:
            return
        self._unlink(victim)
        del self._index[victim.key]

    def _move_to_front(self, node: _Node[K, V]) -> None:
        self._unlink(node)
        self._push_front(node)

    def _unlink(self, node: _Node[K, V]) -> None:
        if node.prev is None:
            self._head = node.next
        else:
            node.prev.next = node.next
        if node.next is None:
            self._tail = node.prev
        else:
            node.next.prev = node.prev
        node.prev = None
        node.next = None

    def _push_front(self, node: _Node[K, V]) -> None:
        node.next = self._head
        node.prev = None
        if self._head is not None:
            self._head.prev = node
        self._head = node
        if self._tail is None:
            self._tail = node
