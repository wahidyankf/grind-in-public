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

There is no plan-fixer. The maker repairs its own work, because a separate fixer would let the checker hand off a
finding and consider itself finished.
