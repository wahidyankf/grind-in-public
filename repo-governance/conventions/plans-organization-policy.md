---
tldr: "Defines how a plan is staged, named, structured, and archived in plans/."
when_to_use: "Use when creating, executing, reviewing, or archiving a plan under plans/."
---

# Plans Organization Policy

## Scope

This policy governs `plans/`, the repository's working record of change. A plan explains why work exists, what it
depends on, and what evidence proves it finished. Plans are temporary and belong to delivery; `docs/` serves readers and
`repo-governance/` holds rules, so neither is a home for a plan. The seven plan workflows named in
[Workflows and Skills](plans-organization-policy/006-workflows-and-skills.md) carry out what this policy defines.

## When a Plan Is Allowed

Create a plan only when the owner explicitly requests one. Do not infer authorization from the size or kind of work,
including application work, infrastructure work, or substantial rule work. When requested, a plan may cover any of those
changes. Drills and study are not planned unless the owner asks: the owner practices by hand and tracks the session in a
harness task list, as the [task tracking policy](task-tracking-policy.md) requires. Rules Propagation may edit multiple
governance surfaces without a delivery plan because its bounded transaction records the authorized rule outcome.

## Rules

Read the rule you need rather than the whole set. The modules are numbered in reading order and
[their index](plans-organization-policy/README.md) carries the full list; these are the entry points.

- [Lifecycle and Folders](plans-organization-policy/001-lifecycle-and-folders.md) — the four stages and their slugs.
- [Lifecycle Moves](plans-organization-policy/002-lifecycle-moves.md) — starting, completing, reopening.
- [Required Documents](plans-organization-policy/003-required-documents.md) — the six plan documents.
- [Technical Shape](plans-organization-policy/004-technical-shape-and-companions.md) — `tech-docs`, one shape only.
- [Delivery Contract](plans-organization-policy/005-delivery-contract.md) — granularity, clarity, executor tags.
- [Workflows and Skills](plans-organization-policy/006-workflows-and-skills.md) — the seven lifecycle capabilities.
- [Structural Validation](plans-organization-policy/007-structural-validation.md) — what is checked mechanically.
- [Evidence and Quality](plans-organization-policy/008-evidence-and-quality.md) — verdicts and bounded repair.
- [Knowledge Capture](plans-organization-policy/009-knowledge-capture-and-archival.md) — draining `learnings.md`.
- [Portability](plans-organization-policy/010-portability.md) — self-containment and recorded deviations.
- [Phases and Gates](plans-organization-policy/011-phases-and-gates.md) — natural pauses.
- [Execution Record](plans-organization-policy/012-execution-record.md) — the dated log of phases, gates, and retries.
- [Two-Pager Template](plans-organization-policy/013-two-pager-template.md) — what an idea contains.
- [Specification Changes](plans-organization-policy/014-specification-changes.md) — planned C4, Gherkin, binding, and
  File Impact deltas.
- [Plan Migrations](plans-organization-policy/015-plan-migrations.md) — safe data, configuration, and dependency
  transitions.
- [Plan UI Design](plans-organization-policy/016-plan-ui-design.md) — selected UI direction, accessible assets, and
  device proof.
- [Plan Document Safety](plans-organization-policy/017-plan-document-safety.md) — ASCII diagrams and secret-free plan
  records.

## Delivery

Plans deliver directly to `main`. This repository runs no pull-request flow, no worktrees, and no delivery modes: a
phase ends, its gate passes, and separately authorized commit and push actions deliver the work. The
[integration path policy](integration-path-policy.md) owns the repository-wide route, and the
[commit hook policy](../development/commit-hook-policy.md) still governs every commit.

## Verification

`plans/` is outside `repo-governance/`, so no word limit applies to a plan. Repository checks still validate its tracked
Markdown links and required indexes. The [plan quality gate](../workflows/plan-quality-gate.md) owns semantic review and
bounded repairs, then consumes those deterministic results before returning `PASS`.
