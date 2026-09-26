---
tldr: "Uses explicit values and validation branches to keep small Python decisions readable."
when_to_use: "Use when choosing value types, checking conditions, or protecting an application rule."
---

# Values and Control Flow

Prefer the value that describes the data. Money uses `Decimal`, not binary `float`; a file path becomes `Path` at the
boundary; a missing value is `None`, checked with `is None`. These choices prevent a surprising amount of application
code from becoming ambiguous.

```python
from decimal import Decimal

balance = Decimal("100.00")
withdrawal = Decimal("12.50")

if withdrawal > balance:
    print("declined")
else:
    balance -= withdrawal
    print(f"balance: {balance}")
```

The result is `balance: 87.50`. `Decimal("12.50")` preserves the exact decimal value written by the input; do not
construct it from a float such as `Decimal(12.50)`.

Use a guard clause when the invalid case should stop the function early:

```python
def require_description(description: str) -> str:
    cleaned = description.strip()
    if not cleaned:
        raise ValueError("description must not be blank")
    return cleaned
```

## Mutability matters

Integers and strings are immutable. Lists and dictionaries are mutable: two names can refer to the same list, so a
mutation through one name is visible through the other.

```text
names ----+--> ["Ada", "Lin"]
aliases --+

aliases.append("Mina")

names ----+--> ["Ada", "Lin", "Mina"]
```

Keep ownership clear. Return a new value when callers should not share mutation; mutate a local list only when that
mutation is the point of the function.

## Checkpoint

Run `require_description` with `" groceries "` and with `"   "`. The first call returns `"groceries"`; the second raises
`ValueError`. The caller will later turn that expected input failure into a useful CLI error message.
