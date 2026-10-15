"""Snowflake-style 64-bit IDs.

Used in: Twitter's original ID scheme and similar generators in large services. Each machine creates IDs without asking
a central service, the IDs are unique across machines, and they sort by creation time.

Layout, high bits to low: 1 unused bit, 41 bits of milliseconds since a fixed epoch, 10 bits of worker id, and 12 bits
of sequence.

Complexity: next_id is O(1). When one worker makes more than 4096 IDs in one millisecond, it waits for the next
millisecond.
"""

import time
from collections.abc import Callable

EPOCH_MS = 1_704_067_200_000  # 2024-01-01T00:00:00Z
WORKER_BITS = 10
SEQUENCE_BITS = 12
MAX_WORKER_ID = (1 << WORKER_BITS) - 1
MAX_SEQUENCE = (1 << SEQUENCE_BITS) - 1
TIME_SHIFT = WORKER_BITS + SEQUENCE_BITS


def _now_ms() -> int:
    return time.time_ns() // 1_000_000


class SnowflakeGenerator:
    def __init__(self, worker_id: int, clock: Callable[[], int] = _now_ms) -> None:
        if not 0 <= worker_id <= MAX_WORKER_ID:
            raise ValueError(f"worker_id must be between 0 and {MAX_WORKER_ID}")
        self._worker_id = worker_id
        self._clock = clock
        self._last_ms = -1
        self._sequence = 0

    def next_id(self) -> int:
        now = self._clock()
        if now < self._last_ms:
            raise RuntimeError("clock moved backwards")
        if now == self._last_ms:
            self._sequence = (self._sequence + 1) & MAX_SEQUENCE
            if self._sequence == 0:
                now = self._wait_past(self._last_ms)
        else:
            self._sequence = 0
        self._last_ms = now
        return ((now - EPOCH_MS) << TIME_SHIFT) | (self._worker_id << SEQUENCE_BITS) | self._sequence

    def _wait_past(self, last_ms: int) -> int:
        now = self._clock()
        while now <= last_ms:
            now = self._clock()
        return now


def decode_snowflake(snowflake: int) -> tuple[int, int, int]:
    """Return (unix milliseconds, worker id, sequence)."""
    timestamp_ms = (snowflake >> TIME_SHIFT) + EPOCH_MS
    worker_id = (snowflake >> SEQUENCE_BITS) & MAX_WORKER_ID
    sequence = snowflake & MAX_SEQUENCE
    return timestamp_ms, worker_id, sequence
