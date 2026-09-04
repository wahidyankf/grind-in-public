---
tldr: "Describes the four plans/ lifecycle stages and what belongs in each."
when_to_use: "Use when deciding which plans/ stage a document belongs in."
---

# Folder Structure

`plans/` holds four stages, and a plan occupies exactly one of them:

```text
plans/
+-- ideas/          quadrant-classified two-pager briefs, not yet plans
+-- backlog/        full plans prepared but not started
+-- in-progress/    full plans under active execution
+-- done/           completed plans, kept as history
```

## Stage Purposes

**`ideas/`** holds a two-pager per idea as `ideas/q<1-4>-<priority>/<slug>.md`, never an idea folder. Choose the
quadrant from dated urgency and importance evidence. An idea is a problem worth solving that has not earned a plan yet.
Promotion to `backlog/` turns it into a five-core-document plan; see the [two-pager template](two-pager-template.md).

**`backlog/`** holds prepared plans that nobody is executing. A backlog plan is complete enough to start without further
authoring: it has passed the [plan-quality-gate](../../workflows/plan-quality-gate.md) workflow.

**`in-progress/`** holds active plans. Keep the count small, because a second active plan splits attention rather than
doubling output. Execution reads and ticks the plan's `delivery.md` in place.

**`done/`** holds finished plans as a historical record. A done plan is not casually rewritten: it records what
happened, including the parts that went badly, and its value comes from being accurate rather than tidy.

## Directory Maps

Every directory recursively under `plans/`, including each technical and asset directory, carries a `README.md` with a
`## Directory Map` that links every direct sibling file and directory other than itself. The
[documentation index policy](../../documentation-index-policy.md) does not reach `plans/`, so this rule makes plans
discoverable. A move updates its source and destination maps in the same change, so no index describes an artifact that
moved.

The [plan quality gate](../../workflows/plan-quality-gate.md) verifies this index requirement.
