"""Binary min-heap and top-k.

Used in: schedulers, timers, rate-limited job queues, and "top N" queries (trending posts, largest files).

Complexity: push and pop are O(log n). Peek is O(1). top_k over n items with heap size k is O(n log k).
"""

from typing import Protocol, Self


class SupportsOrder(Protocol):
    def __lt__(self, other: Self, /) -> bool: ...


class MinHeap[T: SupportsOrder]:
    def __init__(self) -> None:
        self._items: list[T] = []

    def __len__(self) -> int:
        return len(self._items)

    def peek(self) -> T:
        if not self._items:
            raise IndexError("peek at empty heap")
        return self._items[0]

    def push(self, item: T) -> None:
        self._items.append(item)
        self._sift_up(len(self._items) - 1)

    def pop(self) -> T:
        if not self._items:
            raise IndexError("pop from empty heap")
        top = self._items[0]
        last = self._items.pop()
        if self._items:
            self._items[0] = last
            self._sift_down(0)
        return top

    def _sift_up(self, index: int) -> None:
        items = self._items
        while index > 0:
            parent = (index - 1) // 2
            if not items[index] < items[parent]:
                break
            items[index], items[parent] = items[parent], items[index]
            index = parent

    def _sift_down(self, index: int) -> None:
        items = self._items
        size = len(items)
        while True:
            smallest = index
            left = 2 * index + 1
            right = left + 1
            if left < size and items[left] < items[smallest]:
                smallest = left
            if right < size and items[right] < items[smallest]:
                smallest = right
            if smallest == index:
                return
            items[index], items[smallest] = items[smallest], items[index]
            index = smallest


def top_k[T: SupportsOrder](items: list[T], k: int) -> list[T]:
    """Return the k largest items, largest first. A min-heap of size k drops the smallest as it goes."""
    if k <= 0:
        return []
    heap: MinHeap[T] = MinHeap()
    for item in items:
        heap.push(item)
        if len(heap) > k:
            heap.pop()
    result: list[T] = []
    while len(heap) > 0:
        result.append(heap.pop())
    result.reverse()
    return result
