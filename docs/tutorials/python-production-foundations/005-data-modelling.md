---
tldr: "Models a transaction as an immutable record with explicit date and money values."
when_to_use: "Use when related fields form one value in a Python application."
---

# Data Modelling

A transaction has no identity separate from the values in this course, so model it as a frozen dataclass. A frozen value
cannot accidentally change after parsing, which makes reporting and testing easier to reason about.

```python
from dataclasses import dataclass
from datetime import date
from decimal import Decimal


@dataclass(frozen=True)
class Transaction:
    booked_on: date
    description: str
    amount: Decimal
```

Use an `Enum` when a closed vocabulary carries meaning beyond a string. Do not make an enum just to avoid two clear
string literals.

```python
from enum import StrEnum


class Direction(StrEnum):
    CREDIT = "credit"
    DEBIT = "debit"
```

The balance report does not need `Direction`: the sign of `amount` is sufficient. This is deliberate. Good production
code models what the current rule needs and grows only when a new rule demands it.

```text
CSV row
  |
  +--> date ---------> booked_on: date
  +--> description --> description: str
  +--> amount -------> amount: Decimal
                       |
                       v
                  Transaction (immutable)
```

## Checkpoint

Create `src/balance_report/model.py` with `Transaction`. Instantiate it with `date(2026, 9, 26)`, `"Coffee"`, and
`Decimal("-4.50")`. Attempting to assign a new `amount` raises `FrozenInstanceError`, showing that parsed values stay
stable.
