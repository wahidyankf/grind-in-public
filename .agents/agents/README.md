---
description: >-
  Indexes the canonical agents, the role each one plays, and the checker and fixer pair of each quality-gate family.
when_to_use: >-
  Use to find which canonical agent fits a task, or which agent replaced a retired agent name.
---

# Canonical Agents

This directory contains the canonical prompt and semantic capability contract for each shared custom agent. Harness
adapters contain only native metadata and a route to one definition here.

## Available Agents

- [`drill-reviewer.md`](drill-reviewer.md) — reviews a completed owner-solved drill without supplying a solution.
- [`repo-explorer.md`](repo-explorer.md) — locates repository evidence without changing repository state.
- [`plan-maker.md`](plan-maker.md) — authors a formal plan end to end and runs both decision gates.
- [`plan-execution-checker.md`](plan-execution-checker.md) — audits finished execution in fixed order and returns the
  verdict archival depends on.
- [`swe-orchestrator.md`](swe-orchestrator.md) — decomposes a deterministic goal and dispatches the swe family until its
  checks pass, editing nothing.
- [`swe-architect.md`](swe-architect.md) — designs boundaries before a build, reviews it after, and serves as the
  architecture lens of a review pass.
- [`swe-developer.md`](swe-developer.md) — builds behaviour test-first and applies re-validated findings.
- [`swe-debugger.md`](swe-debugger.md) — repairs failing type checks, lint, and tests at the cause.
- [`swe-reviewer.md`](swe-reviewer.md) — audits code, component source, and scenario bindings against the adopted
  standards and returns rated findings, changing nothing.
- [`swe-web-tester.md`](swe-web-tester.md) — judges a running web interface by spec, design, or exploratory charter, and
  is the ui-web gate's judge.
- [`swe-usability-tester.md`](swe-usability-tester.md) — judges first use of a running web interface or command-line
  tool without reading its specification.
- [`swe-api-tester.md`](swe-api-tester.md) — judges a running request-based interface by contract or exploratory
  charter, and is the api-http gate's judge.

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

The ui-web and api-http gates name their own judge and repairer instead: [`swe-web-tester.md`](swe-web-tester.md) or
[`swe-api-tester.md`](swe-api-tester.md) judges, and [`swe-developer.md`](swe-developer.md) repairs in apply-findings
mode.

The plan maker authors and the plan fixer repairs only the rows of a frozen ledger, so each audit stays independent of
the repair before it.

## Old-to-New Map

Each agent the swe family replaced, and where its work went. A reader holding an old name finds its replacement here.

| Replaced agent                    | New agent              | Mode or charter   |
| --------------------------------- | ---------------------- | ----------------- |
| `swe-code-maker`                  | `swe-developer`        | build             |
| `swe-ui-maker`                    | `swe-developer`        | build (UI skills) |
| `swe-code-fixer`, `swe-ui-fixer`  | `swe-developer`        | apply findings    |
| `ui-web-fixer`, `api-http-fixer`  | `swe-developer`        | apply findings    |
| `bugs-solver`                     | `swe-debugger`         | —                 |
| `swe-code-checker`                | `swe-reviewer`         | code              |
| `swe-ui-checker`                  | `swe-reviewer`         | interface         |
| `gherkin-implementation-reviewer` | `swe-reviewer`         | scenario trace    |
| `ui-web-checker`                  | `swe-web-tester`       | spec              |
| `web-design-tester`               | `swe-web-tester`       | design            |
| `web-exploratory-tester`          | `swe-web-tester`       | exploratory       |
| `web-usability-tester`            | `swe-usability-tester` | —                 |
| `api-http-checker`                | `swe-api-tester`       | contract          |
| `api-exploratory-tester`          | `swe-api-tester`       | exploratory       |
| `pr-review-architecture-checker`  | `swe-architect`        | lens              |
| `apps-*-deployer` (per app)       | `swe-releaser`         | deploy            |
