---
tldr: "Reads a UTF-8 CSV through pathlib and validates every row before it reaches the report."
when_to_use: "Use when an application accepts file-based input from a user or another system."
---

# Files and CSV

Turn a file argument into `Path` immediately. Open it with a context manager so the handle closes on success and on
failure. CSV is still untrusted text: check its header, row width, date, description, and decimal amount before
constructing `Transaction`.

```python
import csv
from datetime import date
from pathlib import Path

from balance_report.model import Transaction
from balance_report.parsing import InputError, parse_amount


def parse_row(row: list[str], line_number: int) -> Transaction:
    if len(row) != 3:
        raise InputError(f"line {line_number}: expected three columns")

    date_text, description, amount_text = row
    try:
        booked_on = date.fromisoformat(date_text)
    except ValueError as error:
        raise InputError(f"line {line_number}: invalid date") from error

    cleaned_description = description.strip()
    if not cleaned_description:
        raise InputError(f"line {line_number}: description must not be blank")

    try:
        amount = parse_amount(amount_text)
    except InputError as error:
        raise InputError(f"line {line_number}: {error}") from error

    return Transaction(booked_on, cleaned_description, amount)


def read_transactions(path: Path) -> list[Transaction]:
    try:
        with path.open(encoding="utf-8", newline="") as source:
            rows = csv.reader(source)
            if next(rows, None) != ["date", "description", "amount"]:
                raise InputError("expected date,description,amount header")
            return [parse_row(row, line_number) for line_number, row in enumerate(rows, start=2)]
    except OSError as error:
        raise InputError(f"cannot read {path}") from error
```

Keep `InputError` and `parse_amount` in `src/balance_report/parsing.py`; move the `Transaction` import only after
creating that package layout. The final lesson shows the complete file list.

```text
transactions.csv --Path--> open UTF-8 file --csv.reader--> validate row --> Transaction
       |                         |                 |               |
       +-------------------------+-----------------+---------------+--> InputError on invalid input
```

## Checkpoint

Create `transactions.csv` beside the scratch project's `pyproject.toml` with the required header and two valid rows. Run
`read_transactions(Path("transactions.csv"))` in the REPL and check that it returns two `Transaction` values. Change one
amount to `tea`; the read must raise `InputError` containing `line 2: amount is not decimal`, so the user can locate the
bad external value without seeing a traceback.
