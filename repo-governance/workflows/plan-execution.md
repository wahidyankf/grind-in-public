---
tldr: "Executes a plan phase by phase, delivering to main at each gate, then archives it."
when_to_use: "Use when starting or resuming execution of a plan under plans/."
---

# Plan Execution

## Purpose

Execute a plan's `delivery.md` phase by phase until every item is ticked, every gate has passed, and the plan is
archived in `plans/done/`.

## When to Use

Use it when a plan is ready to run. Resuming an interrupted plan uses this workflow too, starting from
[resuming](plan-execution/04-finalization.md#resuming).

## Prerequisites

Require a current `PASS` from an explicitly owner-directed [plan-quality-gate](plan-quality-gate.md) run. Execution
authority does not authorize that gate. If the result is absent or blocked, stop without starting or rerunning it. The
plan sits in `plans/in-progress/`: move it out of `backlog/`, update both indexes, and push that move only when
authorized before running any checklist item, per the
[lifecycle rules](../conventions/plans-organization-policy/lifecycle-moves.md).

## Steps

1. [Materialize the checklist into the harness task list](plan-execution/01-task-list-sync.md), one task per checkbox.
2. Run Phase 0 and record the baseline: dependencies installed, gates green before anything changes.
3. [Work the current phase](plan-execution/02-phase-loop.md) item by item, in checklist order, ticking each only after
   its change exists.
4. [Pass the phase gate, then deliver](plan-execution/03-gates-and-pushes.md) to `main` only when commit and push are
   separately authorized. Do not begin the next phase while a gate or required delivery item fails.
5. Repeat until the last delivery phase is complete.
6. [Run Knowledge Capture and archive](plan-execution/04-finalization.md) the plan to `plans/done/`.

## Verification

```sh
rtk ./hippo run --class ephemeral --disk-path . -- npm test
rtk ./hippo run --class ephemeral --disk-path . -- npm run format:check
rtk ./hippo run --class ephemeral --disk-path . -- npm run check:markdown-links
```

Execution is complete when every checkbox is ticked, every gate passed, every `learnings.md` entry reached a terminal
state, and the plan folder sits in `plans/done/` with a completion-date prefix and both indexes updated.

## Recovery

A failing gate is fixed inside its own phase, never by continuing. If the fix changes the approach, stop and amend the
plan through [plan-planning](plan-planning.md), then require explicit owner direction for a fresh
[plan-quality-gate](plan-quality-gate.md) run before resuming. Execution never authorizes that run.

If a step turns out to be impossible, do not tick it and do not silently drop it. Record what blocked it in
`learnings.md`, mark the item blocked in `delivery.md` with one line of explanation, and finish every item that does not
depend on it.
