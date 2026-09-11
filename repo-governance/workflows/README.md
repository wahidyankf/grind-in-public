---
tldr: "Indexes repeatable repository procedures."
when_to_use: "Use when a task has a defined sequence, required checks, or recovery steps."
---

# Repository Workflows

This directory contains repeatable procedures for working in Grind in Public. Use a workflow when a task has a defined
sequence, required checks, or recovery steps that should be performed consistently by contributors and agents.

## Adding a Workflow

Create one Markdown file per procedure. The [document naming policy](../conventions/document-naming-policy.md) owns what
to name it, and how a workflow that outgrows one file splits into a child directory. Keep each workflow narrowly scoped;
link to related governance guidance instead of duplicating it.

## Available Workflows

- [Rule workflows](rules/README.md) index the automatically triggered procedures for rule-path changes.
- [Harness Alignment](harness-alignment.md) verifies that every supported harness receives the same rules through its
  instruction file, config, and subagents. Its detail lives in [`harness-alignment/`](harness-alignment/README.md).
- [Rules Propagation](rules/rules-propagation.md) automatically starts for a rule-path change, remains the sole writer,
  and consumes `NEEDS_PROPAGATION` ledgers. Its detail lives in
  [`rules/rules-propagation/`](rules/rules-propagation/README.md).
- [README Refresh](readme-refresh.md) keeps root, project, documentation, and governance READMEs accurate before a
  thematic commit.
- [Rules Quality Gate](rules-quality-gate.md) runs only on explicit owner direction and cannot end blocked; every
  non-pass hands its ledger to Rules Propagation.
- [Rules Grooming](rules-grooming.md) runs only on explicit owner direction, never writes, and hands each approved
  reduction to Rules Propagation.
- [Plan Ideas Grooming](plan-ideas-grooming.md) gives every brief in `plans/ideas/` one disposition, with a reason for
  every keep and retire.
- [Plan Backlog Grooming](plan-backlog-grooming.md) re-judges every backlog plan against the current repository and
  orders the survivors.
- [Plan Planning](plan-planning.md) turns a described change into a validated six-document plan under `plans/`.
- [Plan Quality Gate](plan-quality-gate.md) runs only on explicit owner direction and uses a frozen snapshot and finite
  ledger to return one bounded semantic verdict after at most two cycles.
- [Gherkin Implementation Review](gherkin-implementation-review.md) inspects each expanded scenario and applicable
  adapter for substantive Given-When-Then evidence.
- [Exploratory and Usability Testing](exploratory-and-usability-testing.md) separates spec-aware probing from a fresh,
  spec-blind usability pass for UI-affecting plans.
- [Red-Green-Refactor](red-green-refactor.md) defines the evidenced TDD cycle for application and library behaviour.
- [Plan Execution](plan-execution.md) executes a plan phase by phase, delivering to `main` at each gate, then archives
  it.
- [Plan Execution Check](plan-execution-check.md) judges finished execution in a fixed order and returns one terminal
  verdict that archival depends on.
- [Dev Artifact Clean-Up](dev-artifact-clean-up.md) removes exactly the development artifacts one piece of work created
  — its regenerable build output — and leaves local `main` level with `origin/main`.

Plan Planning and Plan Execution keep detail in [`plan-planning/`](plan-planning/README.md) and
[`plan-execution/`](plan-execution/README.md).

These seven — ideas grooming, backlog grooming, planning, execution, the quality gate, the execution check, and artifact
clean-up — are the complete lifecycle roster named in
[Workflows and Skills](../conventions/plans-organization-policy/006-workflows-and-skills.md). A missing one is a gap in
the lifecycle, not a preference.

## Workflow Template

Each workflow should include:

1. **Purpose** — What outcome the procedure produces.
2. **When to use** — The task or condition that triggers it.
3. **Prerequisites** — Required tools, repository state, or access.
4. **Steps** — Ordered commands and actions, including expected results.
5. **Verification** — Checks that prove the outcome is complete.
6. **Recovery** — Safe next actions if a step fails, when applicable.

Use exact commands and paths where possible. Keep instructions current with the repository tooling, including the
formatting, governance, and dependency checks defined in `package.json`.

## Maintenance

Update a workflow whenever its procedure changes. Move universally required, short rules to the root `AGENTS.md`; keep
the detailed, conditional procedure here to preserve progressive disclosure.
