---
tldr: "Selects linked nodes, stacks, queues, and deques from insertion, removal, and ordering requirements."
when_to_use: "Use for LIFO work, FIFO work, bounded buffers, or pointer-rewiring exercises."
---

# Linked Structures, Stacks, and Queues

A stack answers “what was most recently opened?” A queue answers “what arrived first?” A deque supports efficient work
at both ends. Linked lists matter in interviews because pointer invariants are visible; in production Python, `deque`
usually beats a hand-written linked queue in clarity and memory efficiency.

## Production motivation: bounded work queue

An unbounded in-memory queue turns downstream slowness into memory exhaustion. A bounded deque makes the overload policy
explicit. This example rejects new work rather than silently dropping old work:

```python
from collections import deque
from dataclasses import dataclass, field


class QueueFullError(Exception):
    """The local work buffer has reached its configured capacity."""


@dataclass
class WorkQueue:
    capacity: int
    _items: deque[str] = field(default_factory=deque)

    def enqueue(self, item: str) -> None:
        if len(self._items) >= self.capacity:
            raise QueueFullError("work queue is full")
        self._items.append(item)

    def dequeue(self) -> str | None:
        return self._items.popleft() if self._items else None


queue = WorkQueue(capacity=2)
queue.enqueue("a")
queue.enqueue("b")
assert queue.dequeue() == "a"
```

```text
producer -> [ a | b ] -> consumer
             ^ full
             +-- reject c; caller decides retry or shed
```

Each operation is `O(1)`. The invariant is `0 <= len(items) <= capacity`, with FIFO encounter order. This queue is
process-local and disappears on restart; durable jobs need a broker or database-backed lease.

## Stack example

Use a list as a stack for nested delimiters or workflow undo steps: `append` pushes and `pop` removes from the end in
amortized constant time. Never use `pop(0)` for a queue: it shifts the remaining list and costs `O(n)`.

## Linked-list reality

Linked lists offer constant-time insertion when the node is already known, but finding that node remains linear. Their
extra objects and poor cache locality make them uncommon for ordinary Python application data. They remain useful for
understanding pointer ownership and are embedded inside structures such as LRU caches.

## Checkpoint

Choose an overload policy for an investigator-notification queue: reject, drop oldest, block, or spill to durable
storage. State which requirement makes the choice correct.
