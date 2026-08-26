---
tldr: "Defines how a plan is staged, named, structured, and archived in plans/."
when_to_use: "Use when creating, executing, reviewing, or archiving a plan under plans/."
---

# Plans Organization Policy

## Scope

This policy governs `plans/`, the repository's working record of change. A plan explains why work exists, what it depends on, and what evidence proves it finished. Plans are temporary and belong to delivery; `docs/` serves readers and `repo-governance/` holds rules, so neither is a home for a plan. The three plan workflows — [plan-planning](../workflows/plan-planning.md), [plan-quality-gate](../workflows/plan-quality-gate.md), and [plan-execution](../workflows/plan-execution.md) — carry out what this policy defines.

## When a Plan Is Allowed

Create a plan only when the owner explicitly requests one. Do not infer authorization from the size or kind of work, including application work, infrastructure work, or substantial rule work. When requested, a plan may cover any of those changes. Drills and study are not planned unless the owner asks: the owner practices by hand and tracks the session in a harness task list, as the [task tracking policy](task-tracking-policy.md) requires. `rules-fixer` is exempt inside a [rules quality gate](../workflows/rules-quality-gate.md) run: it may create a receiving document and touch several documents in one pass without a plan, because it records both in the run's [findings report](../workflows/rules-quality-gate/05-findings-report.md).

## Rules

Read the rule you need rather than the whole set:

- [Folder Structure](plans-organization-policy/folder-structure.md) — the four lifecycle stages.
- [Plan Naming](plans-organization-policy/plan-naming.md) — stage-aware folder names.
- [Two-Pager Template](plans-organization-policy/two-pager-template.md) — what an idea contains.
- [Five-Document Structure](plans-organization-policy/five-document-structure.md) — the plan files.
- [Specification Changes](plans-organization-policy/specification-changes.md) — planned C4, Gherkin, binding, and File Impact deltas.
- [Plan Migrations](plans-organization-policy/plan-migrations.md) — safe data, configuration, and dependency transitions.
- [Plan UI Design](plans-organization-policy/plan-ui-design.md) — selected UI direction, accessible assets, and device proof.
- [Plan Document Safety](plans-organization-policy/plan-document-safety.md) — ASCII diagrams and secret-free plan records.
- [Delivery Checklists](plans-organization-policy/delivery-checklists.md) — granularity, clarity, executor tags.
- [Phases and Gates](plans-organization-policy/phases-and-gates.md) — natural pauses.
- [Knowledge Capture](plans-organization-policy/knowledge-capture.md) — draining `learnings.md`.
- [Lifecycle Moves](plans-organization-policy/lifecycle-moves.md) — starting, completing, reopening.

## Delivery

Plans deliver directly to `main`. This repository runs no pull-request flow, no worktrees, and no delivery modes: a phase ends, its gate passes, and the work is committed and pushed. The [commit hook policy](../development/commit-hook-policy.md) still governs every commit.

## Verification

`plans/` is outside `repo-governance/`, so no word limit applies to a plan and no governance check reads one. The Markdown link check does: it reads every Git-tracked document, a plan included. The [plan-quality-gate](../workflows/plan-quality-gate.md) workflow is the verification: `plan-checker` reports findings against these rules and `plan-fixer` resolves them.
