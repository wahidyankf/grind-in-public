---
tldr: "Finishes the scratch application with strict typing, formatting, linting, tests, and CLI smoke checks."
when_to_use: "Use after completing the balance-report implementation and before treating its behaviour as proved."
---

# Quality Gates

A production foundation includes the feedback loops that keep the code honest. Configure Pyright in `pyproject.toml` to
check source and tests in strict mode:

```toml
[tool.pyright]
include = ["src", "tests"]
typeCheckingMode = "strict"
reportUnnecessaryTypeIgnoreComment = true
```

Keep the CLI test's fixture type precise so strict checking can verify its public capture methods:

```python
from pathlib import Path

import pytest

from balance_report.__main__ import main


def test_main_prints_report_to_stdout(tmp_path: Path, capsys: pytest.CaptureFixture[str]) -> None:
    source = tmp_path / "transactions.csv"
    source.write_text("date,description,amount\n2026-09-01,Salary,100.00\n", encoding="utf-8")

    assert main([str(source)]) == 0
    captured = capsys.readouterr()
    assert captured.out == "Transactions: 1\nBalance: 100.00\n"
    assert captured.err == ""
```

Run the gates in the cheapest useful order:

```sh
uv run ruff format --check .
uv run ruff check .
uv run pyright
uv run pytest tests/unit
uv run pytest tests/integration
```

Then exercise the process boundary itself:

```sh
uv run balance-report transactions.csv
uv run python -m balance_report transactions.csv
uv run balance-report missing.csv
```

The first two commands print identical reports and exit zero. The last prints one concise line to stderr and exits 2.
This manual smoke check reaches packaging, argument parsing, filesystem access, rendering, streams, and process status
at once; the narrower automated tests explain which boundary broke when it fails.

```text
source -> format -> lint -> type check -> unit -> integration -> process smoke
   |                                                                  |
   +---------------- stop at the first failed contract ---------------+
```

## Production connection

The sequence scales beyond this small CLI. Fast local checks protect feedback time, boundary tests protect adapters, and
a final public-interface smoke check catches wiring mistakes. More tests are not automatically safer: each layer must
target a distinct failure that a cheaper layer cannot expose.

## Final checkpoint

All gates pass, both valid commands agree, invalid input produces no traceback, and the course's final tree matches the
files on disk. You are ready to use the same Python contracts in the algorithms course.
