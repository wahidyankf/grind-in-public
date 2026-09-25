# Canonical Agents

This directory contains the canonical prompt and semantic capability contract for each shared custom agent. Harness
adapters contain only native metadata and a route to one definition here.

## Available Agents

- [`drill-reviewer.md`](drill-reviewer.md) — reviews a completed owner-solved drill without supplying a solution.
- [`repo-explorer.md`](repo-explorer.md) — locates repository evidence without changing repository state.
- [`plan-maker.md`](plan-maker.md) — authors a formal plan end to end and repairs its own draft within a declared
  budget.
- [`plan-checker.md`](plan-checker.md) — audits a frozen draft and returns findings with a terminal verdict, changing
  nothing.
- [`plan-execution-checker.md`](plan-execution-checker.md) — audits finished execution in fixed order and returns the
  verdict archival depends on.
- [`swe-code-maker.md`](swe-code-maker.md) — builds a project's behaviour test-first under the adopted standards and the
  stack packs its inventory entry lists.
- [`swe-code-checker.md`](swe-code-checker.md) — audits named projects against the adopted standards and returns rated
  findings, changing nothing.
- [`swe-code-fixer.md`](swe-code-fixer.md) — re-validates code checker findings and applies only the high-confidence
  ones, test first where a fix needs a test.

There is no plan-fixer. The maker repairs its own work, because a separate fixer would let the checker hand off a
finding and consider itself finished.
