---
tldr: "Builds a small package entry point that maps arguments and expected errors to process behaviour."
when_to_use: "Use when turning Python application logic into a command that a person or script can run."
---

# Packages and CLI

The CLI is a boundary: it receives strings from the operating system, creates `Path`, invokes application functions, and
decides what reaches stdout, stderr, and the exit status. Keep parsing and reporting free of `argparse` so they remain
easy to test.

```text
argv --> argparse --> Path --> read_transactions --> render_report --> stdout
                  |          |          |
                  +----------+----------+--> InputError --> stderr, status 2
```

Put this code in `src/balance_report/__main__.py`:

```python
import argparse
import sys
from pathlib import Path

from balance_report.parsing import InputError
from balance_report.reader import read_transactions
from balance_report.report import render_report


def build_parser() -> argparse.ArgumentParser:
    parser = argparse.ArgumentParser(description="Print a balance report from CSV transactions.")
    parser.add_argument("transactions", type=Path, help="CSV with date, description, and amount columns")
    return parser


def main(argv: list[str] | None = None) -> int:
    arguments = build_parser().parse_args(argv)
    try:
        transactions = read_transactions(arguments.transactions)
    except InputError as error:
        print(f"error: {error}", file=sys.stderr)
        return 2

    print(render_report(transactions))
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
```

Point the generated command at that function in `pyproject.toml`:

```toml
[project.scripts]
balance-report = "balance_report.__main__:main"
```

`argparse` owns a missing positional argument: it prints usage and returns status 2 before `main` can read a file. Your
code owns failures after the argument has become a `Path`.

## Checkpoint

Run both `uv run balance-report transactions.csv` and `uv run python -m balance_report transactions.csv`. Each valid
invocation prints the same report to stdout and exits zero. A missing file prints one `error:` line to stderr and exits
status 2. Keep error text concise; command users need the cause and file, not a traceback for expected input mistakes.
