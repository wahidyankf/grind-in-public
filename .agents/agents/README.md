# Canonical Agents

This directory contains the canonical prompt and semantic capability contract for each shared custom agent. Harness
adapters contain only native metadata and a route to one definition here.

## Available Agents

- [`drill-reviewer.md`](drill-reviewer.md) — reviews a completed owner-solved drill without supplying a solution.
- [`repo-explorer.md`](repo-explorer.md) — locates repository evidence without changing repository state.
- [`plan-maker.md`](plan-maker.md) — authors a formal plan end to end and runs both decision gates.
- [`plan-execution-checker.md`](plan-execution-checker.md) — audits finished execution in fixed order and returns the
  verdict archival depends on.
- [`swe-code-maker.md`](swe-code-maker.md) — builds a project's behaviour test-first under the adopted standards and the
  stack packs its inventory entry lists.
- [`swe-code-checker.md`](swe-code-checker.md) — audits named projects against the adopted standards and returns rated
  findings, changing nothing.
- [`swe-code-fixer.md`](swe-code-fixer.md) — re-validates code checker findings and applies only the high-confidence
  ones, test first where a fix needs a test.

Each quality-gate family has one read-only checker and one fixer that executes its propagation, per the
[Quality Gate Contract](../../repo-governance/development/workflow/quality-gate-contract.md):

- [`plan-checker.md`](plan-checker.md) and [`plan-fixer.md`](plan-fixer.md) — a plan draft.
- [`docs-checker.md`](docs-checker.md) and [`docs-fixer.md`](docs-fixer.md) — human-facing documents.
- [`rules-checker.md`](rules-checker.md) and [`rules-fixer.md`](rules-fixer.md) — the repository's rules.
- [`harness-checker.md`](harness-checker.md) and [`harness-fixer.md`](harness-fixer.md) — drift from each harness's
  upstream conventions.
- [`ci-checker.md`](ci-checker.md) and [`ci-fixer.md`](ci-fixer.md) — test targets, hooks, and hosted workflows.
- [`pr-review-checker.md`](pr-review-checker.md) and [`pr-review-fixer.md`](pr-review-fixer.md) — one commit range.
- [`specs-checker.md`](specs-checker.md) and [`specs-fixer.md`](specs-fixer.md) — specification folders.
- [`ui-web-checker.md`](ui-web-checker.md) and [`ui-web-fixer.md`](ui-web-fixer.md) — a running web interface.
- [`api-http-checker.md`](api-http-checker.md) and [`api-http-fixer.md`](api-http-fixer.md) — a running HTTP interface.

The plan maker authors and the plan fixer repairs only the rows of a frozen ledger, so each audit stays independent of
the repair before it.
