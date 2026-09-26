---
tldr: "Teaches production-relevant Python foundations through a small CSV balance-report CLI."
when_to_use: "Use when learning Python syntax and standard-library patterns before algorithm or application work."
---

# Python Production Foundations

This course gives you a practical starting point for Python used in applications: clear values, explicit errors, typed
functions, small data models, file boundaries, and a command-line interface. It is not a framework course and does not
attempt to teach web development, databases, async code, or full ledger domain modelling.

You will create a disposable packaged application named `balance-report` outside this repository. Its final commands are
`uv run balance-report transactions.csv` and the equivalent `uv run python -m balance_report transactions.csv`. It reads
a UTF-8 CSV with `date`, `description`, and signed decimal `amount` columns, then prints a transaction count and final
balance.

## Prerequisites

Install `uv`; it can obtain Python when the requested version is absent. Create a scratch project with
`uv init --app --package --python 3.14 balance-report`, enter it with `cd balance-report`, and add the development tools
with `uv add --dev pyright pytest ruff`. The explicit `--package` flag records that the `src/` package and command entry
point are intentional; current `uv` application templates are packaged by default. Keep this project outside
`grind-in-public`: the course is documentation, not a new workspace application.

Every Python listing in this course uses a `python` fence. Terminal output and ASCII diagrams use `text` because they
are not Python source.

## Learning path

```text
setup -> values -> functions -> collections -> models -> CSV -> CLI -> tests -> gates
  |                                                                             |
  +----------------------------- balance-report -------------------------------+
```

Read and type the lessons in order:

1. [Toolchain and REPL](001-toolchain-and-repl.md) — run small Python experiments predictably.
2. [Values and Control Flow](002-values-and-control-flow.md) — state, truth, and validation branches.
3. [Functions and Errors](003-functions-and-errors.md) — typed contracts and narrow failures.
4. [Collections and Iteration](004-collections-and-iteration.md) — the containers used constantly in application and
   interview code.
5. [Data Modelling](005-data-modelling.md) — frozen records, enums, dates, and money.
6. [Files and CSV](006-files-and-csv.md) — turn untrusted text into validated values.
7. [Packages and CLI](007-packages-and-cli.md) — a small, usable command-line boundary.
8. [Testing the Balance Report](008-testing-the-balance-report.md) — prove the pure report and file boundary.
9. [Quality Gates](009-quality-gates.md) — run strict typing, formatting, linting, tests, and command smoke checks.

Each lesson has a checkpoint. Type the code yourself, run it, and compare its result with the expected behaviour before
moving on. The final program intentionally has one small responsibility; its structure prepares you for a richer ledger
later without pretending that a crash course solves the ledger's domain rules.

## References

- [Python 3.14 documentation](https://docs.python.org/3.14/)
- [`uv init` project templates](https://docs.astral.sh/uv/concepts/projects/init/)
- [Pyright configuration](https://microsoft.github.io/pyright/#/configuration)
- [pytest documentation](https://docs.pytest.org/en/stable/)
- [Ruff documentation](https://docs.astral.sh/ruff/)
