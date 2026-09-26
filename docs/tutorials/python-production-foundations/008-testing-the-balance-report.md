---
tldr: "Completes and tests a small CSV balance-report CLI with pure report and file-boundary coverage."
when_to_use:
  "Use when finishing the course or checking that a Python CLI behaves correctly for valid and invalid input."
---

# Testing the Balance Report

Keep the report pure: it accepts transactions and returns text. This is a unit-test boundary because it reads no file
and no process state.

```python
from collections.abc import Sequence
from decimal import Decimal

from balance_report.model import Transaction


def render_report(transactions: Sequence[Transaction]) -> str:
    balance = sum((transaction.amount for transaction in transactions), start=Decimal("0"))
    return "\n".join((f"Transactions: {len(transactions)}", f"Balance: {balance:.2f}"))
```

Save it as `src/balance_report/report.py`. The source tree should now be:

```text
balance-report/
+-- src/balance_report/
|   +-- __init__.py
|   +-- __main__.py
|   +-- model.py
|   +-- parsing.py
|   +-- reader.py
|   `-- report.py
+-- tests/
|   +-- test_reader.py
|   `-- test_report.py
+-- pyproject.toml
`-- transactions.csv
```

Put the parser from Lesson 3 into `parsing.py`, the file reader from Lesson 6 into `reader.py`, and the CLI from Lesson
7 into `__main__.py`. Add these tests.

```python
from datetime import date
from decimal import Decimal

from balance_report.model import Transaction
from balance_report.report import render_report


def test_render_report_counts_transactions_and_sums_money() -> None:
    transactions = [
        Transaction(date(2026, 9, 1), "Salary", Decimal("100.00")),
        Transaction(date(2026, 9, 2), "Coffee", Decimal("-4.50")),
    ]

    assert render_report(transactions) == "Transactions: 2\nBalance: 95.50"
```

```python
from pathlib import Path

import pytest

from balance_report.parsing import InputError
from balance_report.reader import read_transactions


def test_read_transactions_rejects_a_bad_amount(tmp_path: Path) -> None:
    source = tmp_path / "transactions.csv"
    source.write_text("date,description,amount\n2026-09-01,Coffee,tea\n", encoding="utf-8")

    with pytest.raises(InputError, match="amount is not decimal"):
        read_transactions(source)
```

The first test is unit-level: it supplies values directly. The second touches a real temporary file, so it proves the
file boundary. Run both with `uv run pytest`.

## Final checkpoint

Use this input file:

```text
date,description,amount
2026-09-01,Salary,100.00
2026-09-02,Coffee,-4.50
```

Then `uv run python -m balance_report transactions.csv` prints:

```text
Transactions: 2
Balance: 95.50
```

You now have the Python tools that appear repeatedly in application code and interview solutions: carefully chosen
collections, typed functions, immutable values, explicit validation, standard-library boundaries, and tests. Next,
practice them by solving focused data-structure and algorithm drills yourself; defer a richer ledger until the domain
rules, not Python syntax, are the learning goal.
