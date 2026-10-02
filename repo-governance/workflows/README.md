---
tldr: "Indexes repeatable repository procedures in their plan, quality, and maintenance groups."
when_to_use: "Use when a task has a defined sequence, required checks, or recovery steps."
---

# Repository Workflows

This directory contains repeatable procedures for working in Grind in Public. Use a workflow when a task has a defined
sequence, required checks, or recovery steps that should be performed consistently by contributors and agents.

## Groups

Every workflow sits in one of three groups. No other entry belongs at this level.

- [Plan](plan/README.md) holds the plan lifecycle: ideas and backlog grooming, planning, execution, and the execution
  check.
- [Quality](quality/README.md) holds every quality gate with its propagation, the automatically triggered Rules and Docs
  Propagation, and every single-pass review: Gherkin implementation, exploratory usability, harness parity, and the push
  leak review.
- [Maintenance](maintenance/README.md) holds upkeep: rules grooming and development artifact clean-up.

Every quality gate follows the [Quality Gate Contract](../development/workflow/quality-gate-contract.md): it runs only
on an explicit owner request, never edits, hands its blocking findings to its family's propagation, stops after at most
three cycles, and returns an advisory verdict. The
[Quality Gate Adapter](../development/workflow/quality-gate-adapter.md) records how this repository adopted the
contract.

Ideas grooming, backlog grooming, planning, execution, the plan quality gate, the execution check, and artifact clean-up
are the complete lifecycle roster named in
[Workflows and Skills](../conventions/plans-organization-policy/006-workflows-and-skills.md). A missing one is a gap in
the lifecycle, not a preference.

## Adding a Workflow

Create one Markdown file per procedure in the group that owns its purpose. The
[document naming policy](../conventions/document-naming-policy.md) owns what to name it, and how a workflow that
outgrows one file splits into a child directory of numbered modules. Keep each workflow narrowly scoped; link to related
governance guidance instead of duplicating it.

## Workflow Template

Each workflow should include:

1. **Purpose** — What outcome the procedure produces.
2. **When to use** — The task or condition that triggers it.
3. **Prerequisites** — Required tools, repository state, or access.
4. **Steps** — Ordered commands and actions, including expected results.
5. **Verification** — Checks that prove the outcome is complete.
6. **Recovery** — Safe next actions if a step fails, when applicable.

A quality gate and a propagation instead carry the section shapes their contracts fix. Use exact commands and paths
where possible. Keep instructions current with the repository tooling, including the formatting, governance, and
dependency checks defined in `package.json`.

## Maintenance

Update a workflow whenever its procedure changes. Move universally required, short rules to the root `AGENTS.md`; keep
the detailed, conditional procedure here to preserve progressive disclosure.
