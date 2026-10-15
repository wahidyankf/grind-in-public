"""Work queue and append-only log.

Used in: message brokers (Kafka, RabbitMQ), job systems, and event sourcing. A plain queue hands each message to one
worker and forgets it. A log keeps every record, so each consumer group tracks its own offset and can replay.

Complexity: enqueue and dequeue are O(1). Append is O(1). Read of k records is O(k).
"""

from collections import deque


class WorkQueue[T]:
    def __init__(self) -> None:
        self._items: deque[T] = deque()

    def __len__(self) -> int:
        return len(self._items)

    def enqueue(self, item: T) -> None:
        self._items.append(item)

    def dequeue(self) -> T:
        if not self._items:
            raise IndexError("dequeue from empty queue")
        return self._items.popleft()


class AppendOnlyLog[T]:
    def __init__(self) -> None:
        self._records: list[T] = []
        self._committed: dict[str, int] = {}

    def append(self, record: T) -> int:
        """Store a record and return its offset."""
        self._records.append(record)
        return len(self._records) - 1

    def read(self, offset: int, limit: int) -> list[T]:
        """Read up to limit records starting at offset. Reading does not move any consumer offset."""
        return self._records[offset : offset + limit]

    def committed(self, group: str) -> int:
        """Return the next offset the group should read. Groups that never committed start at 0."""
        return self._committed.get(group, 0)

    def commit(self, group: str, next_offset: int) -> None:
        self._committed[group] = next_offset

    def poll(self, group: str, limit: int) -> list[T]:
        """Read from the group's committed offset. The caller must commit to move on."""
        return self.read(self.committed(group), limit)
