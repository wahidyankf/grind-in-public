"""Leaky bucket queue: a bounded queue that drains at a fixed rate.

Used in: traffic shaping (network shapers), throttled outbound notifications, and protecting a slow downstream service
from bursts. The output rate stays steady even when the input is bursty.

Complexity: offer and drain are O(1) each, plus one step per released item.

Items wait in the queue until drain releases them. When the queue is full, offer returns False.
"""

import time
from collections import deque
from collections.abc import Callable


class LeakyBucketQueue[T]:
    def __init__(
        self,
        capacity: int,
        leak_rate: float,
        clock: Callable[[], float] = time.monotonic,
    ) -> None:
        if capacity < 1 or leak_rate <= 0:
            raise ValueError("capacity must be at least 1 and leak_rate must be positive")
        self._capacity = capacity
        self._leak_rate = leak_rate
        self._clock = clock
        self._items: deque[T] = deque()
        self._credit = 0.0
        self._last = clock()

    def __len__(self) -> int:
        return len(self._items)

    def offer(self, item: T) -> bool:
        if not self._items:
            # An idle bucket must not save up credit, so the timer starts when the first item arrives.
            self._credit = 0.0
            self._last = self._clock()
        if len(self._items) >= self._capacity:
            return False
        self._items.append(item)
        return True

    def drain(self) -> list[T]:
        """Release the items the elapsed time has earned, oldest first."""
        self._accrue()
        released: list[T] = []
        while self._items and self._credit >= 1.0:
            released.append(self._items.popleft())
            self._credit -= 1.0
        if not self._items:
            self._credit = 0.0
        return released

    def _accrue(self) -> None:
        now = self._clock()
        self._credit += (now - self._last) * self._leak_rate
        self._last = now
