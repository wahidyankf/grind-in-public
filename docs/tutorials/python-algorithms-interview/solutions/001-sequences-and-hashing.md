---
tldr: "Solves stable event deduplication and first-seen grouped totals with explicit hash-table invariants."
when_to_use: "Use after attempting drill 001 without the reference implementation."
---

# Sequences and Hashing Solutions

## Stable event deduplication

The dictionary maps each identifier to its accepted payload. The output list contains exactly the first accepted event
for each identifier in encounter order.

```python
from collections.abc import Sequence
from dataclasses import dataclass


@dataclass(frozen=True)
class Event:
    event_id: str
    payload: str


def stable_events(events: Sequence[Event]) -> list[Event]:
    accepted_payloads: dict[str, str] = {}
    accepted: list[Event] = []
    for event in events:
        previous = accepted_payloads.get(event.event_id)
        if previous is None:
            accepted_payloads[event.event_id] = event.payload
            accepted.append(event)
        elif previous != event.payload:
            raise ValueError(f"conflicting payload for {event.event_id}")
    return accepted


assert stable_events([Event("a", "x"), Event("b", "y"), Event("a", "x")]) == [
    Event("a", "x"),
    Event("b", "y"),
]

try:
    stable_events([Event("a", "x"), Event("a", "z")])
except ValueError as error:
    assert str(error) == "conflicting payload for a"
else:
    raise AssertionError("conflicting duplicate was accepted")
```

Expected time is `O(n)` and additional space is `O(k)`. Cross-batch correctness still needs a durable unique key and a
stored content hash or payload comparison inside the authoritative transaction.

## First-seen grouped totals

```python
from collections.abc import Iterable


def totals_by_tenant(entries: Iterable[tuple[str, int]]) -> dict[str, int]:
    totals: dict[str, int] = {}
    for tenant, amount in entries:
        if amount < 0:
            raise ValueError("amount must not be negative")
        totals[tenant] = totals.get(tenant, 0) + amount
    return totals


assert totals_by_tenant([("north", 3), ("south", 2), ("north", 4)]) == {
    "north": 7,
    "south": 2,
}
```

The invariant is that each total equals the consumed prefix for that tenant. This is a batch aggregate, not an
authoritative balance: concurrent writers and retries require a transaction or idempotent event ledger.
