---
tldr: "Turns a described change into a five-core-document plan ready for an explicitly requested quality gate."
when_to_use: "Use only after the owner explicitly requests a plan."
---

# Plan Planning

## Purpose

Turn a described change into a five-core-document plan that another session could execute without inventing a decision.
The plan lands in `plans/backlog/` or `plans/in-progress/` ready for a separately authorized
[plan-quality-gate](plan-quality-gate.md) run.

## When to Use

Use it only after the owner explicitly requests a plan. The
[plans organization policy](../conventions/plans-organization-policy.md) states that authorization boundary. Do not use
it for application work, infrastructure work, substantial rule work, or a drill unless the owner asks for a plan.

## Prerequisites

Know which stage the plan targets. `backlog/` means prepared but not started and is the default. Use `in-progress/` only
when separately authorized execution will follow after an explicitly requested quality-gate `PASS`; commit and push
remain separate permissions.

## Steps

1. [Explore before asking](plan-planning/01-exploration.md). Read the code, the governance documents, and the Git
   history that bear on the change. A question the repository already answers must not reach the owner.
2. [Grill the open decisions](plan-planning/02-grilling.md) with structured options, as the
   [grilling-with-options policy](../conventions/grilling-with-options-policy.md) requires. Unresolved decisions become
   guesses baked into a checklist.
3. [Author the five core documents](plan-planning/03-plan-authoring.md) into `plans/<stage>/<identifier>/`, following
   the [five-document structure](../conventions/plans-organization-policy/five-document-structure.md).
4. [Review the structure](plan-planning/04-structural-review.md) against the checks listed there before handing off.
5. Update the stage's `README.md` index and run the deterministic structural checks. Do not start the semantic quality
   gate unless the owner explicitly requests that checkpoint.
6. Commit and push the plan only under separate authorization. A local plan remains unready for execution until an
   explicitly requested gate returns `PASS`.

## Verification

```sh
npm run format:check
npm run check:markdown-links
```

The plan is ready for semantic review when every document exists, every checklist item names a path, command, and
acceptance criterion, every phase ends with a gate, and the deterministic structural checks pass.

## Recovery

If an owner-requested plan turns out to need a different shape or scope, ask the owner whether to amend or cancel it
rather than deleting it based on an inferred size threshold. If a decision cannot be resolved, leave the plan in
`backlog/` with the open question written into `README.md` rather than guessing past it.
