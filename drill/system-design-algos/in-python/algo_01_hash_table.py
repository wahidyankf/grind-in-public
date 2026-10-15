"""Hash table with separate chaining.

Used in: caches, session stores, and key-to-server maps. Almost every system keeps one somewhere.

Complexity: put, get, and delete are O(1) on average. The table doubles its bucket count when the load factor passes
0.75, so the average bucket stays short.
"""

from typing import Final

LOAD_FACTOR_LIMIT: Final[float] = 0.75
INITIAL_BUCKETS: Final[int] = 8


class HashTable[K, V]:
    def __init__(self, capacity: int = INITIAL_BUCKETS) -> None:
        self._buckets: list[list[tuple[K, V]]] = [[] for _ in range(capacity)]
        self._size = 0

    def __len__(self) -> int:
        return self._size

    @property
    def bucket_count(self) -> int:
        return len(self._buckets)

    def _index(self, key: K) -> int:
        return hash(key) % len(self._buckets)

    def put(self, key: K, value: V) -> None:
        bucket = self._buckets[self._index(key)]
        for i, (existing_key, _) in enumerate(bucket):
            if existing_key == key:
                bucket[i] = (key, value)
                return
        bucket.append((key, value))
        self._size += 1
        if self._size / len(self._buckets) > LOAD_FACTOR_LIMIT:
            self._resize(len(self._buckets) * 2)

    def get(self, key: K) -> V | None:
        for existing_key, value in self._buckets[self._index(key)]:
            if existing_key == key:
                return value
        return None

    def delete(self, key: K) -> bool:
        bucket = self._buckets[self._index(key)]
        for i, (existing_key, _) in enumerate(bucket):
            if existing_key == key:
                del bucket[i]
                self._size -= 1
                return True
        return False

    def _resize(self, new_capacity: int) -> None:
        old_pairs = [pair for bucket in self._buckets for pair in bucket]
        self._buckets = [[] for _ in range(new_capacity)]
        self._size = 0
        for key, value in old_pairs:
            self.put(key, value)
