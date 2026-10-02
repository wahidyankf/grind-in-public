---
tldr:
  "Indexes every quality gate with its propagation, the automatically triggered writers, and the single-pass reviews."
when_to_use: "Use when judging or repairing a change, a plan, the rules, the docs, or a running surface."
---

# Quality Workflows

Each gate follows the [Quality Gate Contract](../../development/workflow/quality-gate-contract.md), and each propagation
is its family's sole writer under [Sole-Writer Propagation](../../development/workflow/sole-writer-propagation.md). The
[Quality Gate Adapter](../../development/workflow/quality-gate-adapter.md) records the local decisions.

## Directory Map

Gates and their propagations, one pair per family:

- [Plan Quality Gate](plan-quality-gate.md) and [Plan Propagation](plan-propagation.md) judge and repair one plan draft.
- [Docs Quality Gate](docs-quality-gate.md) and [Docs Propagation](docs-propagation.md) judge and repair human-facing
  documents; the propagation also starts automatically before a documentation-changing commit.
- [Rules Quality Gate](rules-quality-gate.md) and [Rules Propagation](rules-propagation.md) judge and write the rules;
  the propagation starts automatically for a rule-path change. Its detail lives in
  [`rules-propagation/`](rules-propagation/README.md).
- [Harness Quality Gate](harness-quality-gate.md) and [Harness Propagation](harness-propagation.md) judge and repair
  drift from each harness's upstream conventions.
- [CI Quality Gate](ci-quality-gate.md) and [CI Propagation](ci-propagation.md) judge and repair test targets, hooks,
  and hosted workflow wiring.
- [PR Review Quality Gate](pr-review-quality-gate.md) and [PR Review Propagation](pr-review-propagation.md) judge and
  repair one local commit range, through [PR Review](pr-review.md). Its detail lives in
  [`pr-review-quality-gate/`](pr-review-quality-gate/README.md).
- [Specs Quality Gate](specs-quality-gate.md) and [Specs Propagation](specs-propagation.md) judge and repair
  specification folders.
- [UI Web Quality Gate](ui-web-quality-gate.md) and [UI Web Propagation](ui-web-propagation.md) judge and repair a
  running web interface.
- [API HTTP Quality Gate](api-http-quality-gate.md) and [API HTTP Propagation](api-http-propagation.md) judge and repair
  a running HTTP interface.

Single-pass reviews and the test-first cycle:

- [Harness Parity Verification](harness-parity-verification.md) verifies that every supported harness receives the same
  rules through its instruction file, config, and subagents. Its detail lives in
  [`harness-parity-verification/`](harness-parity-verification/README.md).
- [Push Leak Review](push-leak-review.md) reviews every outgoing commit for a leak, privately, before each push; its
  classes, review, and enforcement live in [`push-leak-review/`](push-leak-review/README.md).
- [Gherkin Implementation Review](gherkin-implementation-review.md) inspects each expanded scenario and applicable
  adapter for substantive Given-When-Then evidence.
- [Exploratory Usability Review](exploratory-usability-review.md) separates spec-aware probing from a fresh, spec-blind
  usability pass for UI-affecting plans.
- [Red-Green-Refactor](red-green-refactor.md) defines the evidenced TDD cycle for application and library behaviour.
