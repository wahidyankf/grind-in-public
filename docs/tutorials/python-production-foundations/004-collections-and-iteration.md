---
tldr: "Chooses Python collections by access pattern and processes them without hiding application rules."
when_to_use:
  "Use when storing related values, checking membership, or traversing data for application or interview code."
---

# Collections and Iteration

Use `list` for ordered, repeatable values; `tuple` for a fixed small sequence; `dict` for lookup by a unique key; and
`set` for membership without duplicates. Choose the collection from the operation you need, not from habit.

```text
list: ordered traversal       [first] -> [second] -> [third]
dict: lookup by a key         "food" -> Decimal("12.50")
set: membership only          {"coffee", "rent", "food"}
```

The report needs an ordered list of transactions and a total. A generator expression keeps the transformation local;
`sum` starts from a decimal zero so an empty list produces the right money type.

```python
from decimal import Decimal

amounts = [Decimal("100.00"), Decimal("-4.50"), Decimal("-12.00")]
balance = sum(amounts, start=Decimal("0"))
print(balance)
```

The result is `83.50`. Do not use `sum` to join strings or repeatedly concatenate a large list; use `"".join` and list
accumulation for those jobs.

Use a dictionary when you need to aggregate by a known key:

```python
from collections import defaultdict
from decimal import Decimal

spent_by_category: defaultdict[str, Decimal] = defaultdict(lambda: Decimal("0"))
for category, amount in [("food", Decimal("-12.00")), ("food", Decimal("-4.50"))]:
    spent_by_category[category] += amount
```

## Checkpoint

Add a third `("transport", Decimal("-3.00"))` row. Confirm the dictionary has `food` at `Decimal("-16.50")` and
`transport` at `Decimal("-3.00")`. Later lessons use a list for the report because input order remains meaningful.
