---
tldr: "Defines small typed functions and reports expected invalid input through precise exceptions."
when_to_use: "Use when turning a calculation or validation rule into a callable application contract."
---

# Functions and Errors

Give every function a narrow job, annotated parameters, and one understandable return shape. Put required inputs first;
use keyword-only arguments only when naming a caller's choice makes misuse less likely.

```python
from decimal import Decimal, InvalidOperation


class InputError(Exception):
    """An input value cannot become a valid application value."""


def parse_amount(text: str) -> Decimal:
    try:
        amount = Decimal(text)
    except InvalidOperation as error:
        raise InputError(f"amount is not decimal: {text!r}") from error

    if not amount.is_finite():
        raise InputError("amount must be finite")
    return amount
```

The parser validates at the boundary. An annotation such as `text: str` says what the function accepts; it does not
prove that a CSV field holds a usable decimal.

```text
raw CSV text --> parse_amount --> Decimal
      |              |
      |              +--> InputError for an invalid external value
      v
untrusted boundary
```

Catch the narrowest exception that you expect from the operation. Do not use a bare `except`, and do not turn a
programming mistake into a friendly input error. The `from error` preserves the original cause for debugging while the
user-facing message stays focused.

## Checkpoint

`parse_amount("-4.50")` returns `Decimal("-4.50")`. `parse_amount("coffee")` raises `InputError`, not
`InvalidOperation`. This makes the rest of the program independent from `decimal`'s implementation detail.
