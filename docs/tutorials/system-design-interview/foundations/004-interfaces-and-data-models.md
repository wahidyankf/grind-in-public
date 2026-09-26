---
tldr: "Designs versioned APIs, events, identifiers, idempotency, and data models around domain invariants."
when_to_use: "Use before the component diagram so contracts and ownership shape the architecture."
---

# Interfaces and Data Models

Interfaces outlive deployments. Define the smallest stable contract at the business boundary, validate it once, and
preserve raw input separately when audit or reprocessing requires it.

## Command, query, and event

```text
command: "please decide"  -> may succeed or fail
query:   "show decision"  -> returns current representation
event:   "decision made"  -> immutable fact in past tense
```

Commands need idempotency because clients retry uncertain outcomes. Events need stable identity, schema version,
producer time, observed time, tenant, trace context, and an ordering key when order matters.

## Complete Python boundary model

```python
from dataclasses import dataclass
from datetime import UTC, datetime
from decimal import Decimal, InvalidOperation
from typing import Any
from uuid import UUID


class ContractError(ValueError):
    """Raised when an external command violates the public contract."""


@dataclass(frozen=True, slots=True)
class PaymentCommand:
    command_id: UUID
    tenant_id: UUID
    account_id: UUID
    amount: Decimal
    currency: str
    occurred_at: datetime


def parse_payment(payload: dict[str, Any]) -> PaymentCommand:
    try:
        occurred_at = datetime.fromisoformat(str(payload["occurred_at"]))
        amount = Decimal(str(payload["amount"]))
        command = PaymentCommand(
            command_id=UUID(str(payload["command_id"])),
            tenant_id=UUID(str(payload["tenant_id"])),
            account_id=UUID(str(payload["account_id"])),
            amount=amount,
            currency=str(payload["currency"]).upper(),
            occurred_at=occurred_at,
        )
    except (KeyError, InvalidOperation, TypeError, ValueError) as error:
        raise ContractError("invalid payment command") from error

    if command.amount <= 0:
        raise ContractError("amount must be positive")
    if len(command.currency) != 3:
        raise ContractError("currency must be a three-letter code")
    if command.occurred_at.tzinfo is None:
        raise ContractError("occurred_at must include an offset")
    return command


example = parse_payment(
    {
        "command_id": "5bf0d8e5-08dc-4ace-b40b-2c52f78645ae",
        "tenant_id": "4edc64b7-4c3d-4701-9ff2-b5494da7a7ce",
        "account_id": "311c91f5-31ec-4144-9b2c-9f4104bc3115",
        "amount": "19.95",
        "currency": "usd",
        "occurred_at": datetime.now(UTC).isoformat(),
    }
)
print(example.currency)  # USD
```

`Decimal` preserves money semantics that binary floating point cannot. Frozen, slotted data objects prevent accidental
mutation and reduce per-object overhead. Reject this in-memory model as the sole evidence format: persistence still
needs a schema, provenance, and migration policy.

## Idempotency state machine

```text
                 same key + same request
NEW ----------------------------------------> COMPLETED -> replay result
 |                                                ^
 | reserve key                                    |
 v                                                |
PROCESSING -- effect committed + result stored ---+
 |     |
 |     +-- lease expires --> retry or reconcile
 +-- same key + different request --> 409 conflict
```

Store a request hash with the key. Returning the first result for a different payload would hide client corruption.
Choose an expiry longer than the maximum retry horizon and business dispute window.

## Schema evolution

Prefer additive changes: add an optional field, deploy readers that tolerate it, then deploy writers. Renaming or
changing meaning needs a new field or version. A compatibility registry protects syntax; contract tests and shadow
readers protect semantics.

## Data model questions

- Which entity owns the invariant?
- Which identifiers are stable across imports and retries?
- What must be immutable, and what is a current projection?
- Which query paths need an index?
- Which history must be reconstructable after software changes?

Do not normalize or denormalize by ideology. Normalize when transactionally consistent updates dominate; denormalize
when bounded read latency matters and asynchronous repair is acceptable.
