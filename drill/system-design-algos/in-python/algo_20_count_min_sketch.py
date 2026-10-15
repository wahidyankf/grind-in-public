"""Count-min sketch with heavy hitters.

Used in: estimating how often items appear in a stream with fixed memory. Examples are trending hashtags, hot keys in a
cache, and top talkers in network monitoring.

Complexity: add and estimate are O(depth). Memory is width * depth counters.

The sketch never undercounts. Collisions can only add to a count, so the estimate is the smallest value across rows.
"""

import hashlib
import heapq
import math


class CountMinSketch:
    def __init__(self, width: int, depth: int) -> None:
        if width < 1 or depth < 1:
            raise ValueError("width and depth must be at least 1")
        self._width = width
        self._depth = depth
        self._table: list[list[int]] = [[0] * width for _ in range(depth)]

    @classmethod
    def from_error(cls, epsilon: float, delta: float) -> "CountMinSketch":
        """Overcount is at most epsilon times the total count, with probability 1 - delta."""
        if not (0 < epsilon < 1 and 0 < delta < 1):
            raise ValueError("epsilon and delta must be between 0 and 1")
        return cls(math.ceil(math.e / epsilon), math.ceil(math.log(1 / delta)))

    def add(self, item: str, count: int = 1) -> None:
        for row, column in enumerate(self._columns(item)):
            self._table[row][column] += count

    def estimate(self, item: str) -> int:
        return min(self._table[row][column] for row, column in enumerate(self._columns(item)))

    def _columns(self, item: str) -> list[int]:
        columns: list[int] = []
        for row in range(self._depth):
            # Each row salts the hash with its own number, so the rows pick different columns.
            digest = hashlib.blake2b(f"{row}:{item}".encode(), digest_size=8).digest()
            columns.append(int.from_bytes(digest, "big") % self._width)
        return columns


class HeavyHitters:
    """Keep the k items with the largest estimated counts, using a min-heap of candidates."""

    def __init__(self, sketch: CountMinSketch, k: int) -> None:
        if k < 1:
            raise ValueError("k must be at least 1")
        self._sketch = sketch
        self._k = k
        self._current: dict[str, int] = {}
        # Entries can go stale when an item's estimate grows. They are skipped when they reach the top.
        self._heap: list[tuple[int, str]] = []

    def observe(self, item: str, count: int = 1) -> None:
        self._sketch.add(item, count)
        estimate = self._sketch.estimate(item)
        if item in self._current or len(self._current) < self._k:
            self._set(item, estimate)
            return
        self._drop_stale()
        smallest_estimate, smallest_item = self._heap[0]
        if estimate > smallest_estimate:
            heapq.heappop(self._heap)
            del self._current[smallest_item]
            self._set(item, estimate)

    def top(self) -> list[tuple[str, int]]:
        return sorted(self._current.items(), key=lambda pair: (-pair[1], pair[0]))

    def _set(self, item: str, estimate: int) -> None:
        self._current[item] = estimate
        heapq.heappush(self._heap, (estimate, item))

    def _drop_stale(self) -> None:
        while self._heap and self._current.get(self._heap[0][1]) != self._heap[0][0]:
            heapq.heappop(self._heap)
