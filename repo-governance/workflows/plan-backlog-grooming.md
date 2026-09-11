---
tldr: "Re-judges every backlog plan against the current repository and orders the survivors."
when_to_use: "Use when plans/backlog/ holds plans written before the repository changed under them."
---

# Plan Backlog Grooming

## Purpose

Establish, for every plan in `plans/backlog/`, whether it is still true — and in what order the true ones should run.

## When to Use

Use it when `plans/backlog/` holds at least one formal plan and the repository has changed since they were written. Run
it on explicit owner direction.

## Prerequisites

Local `main` is current, because the question this workflow asks is about the repository as it is now.

## Steps

1. **Freeze the list.** Enumerate every plan in `plans/backlog/` first. Plans added mid-pass belong to the next one.
2. **Re-read each plan against the current repository**, not against the repository it was written for. The question is
   not whether the plan is well written; it is whether it is still true.
3. **Assign exactly one disposition:**

   | Disposition | Means                                                                 |
   | ----------- | --------------------------------------------------------------------- |
   | valid       | assumptions still hold; ready to execute as written                   |
   | revise      | the goal holds but specifics have gone stale; name what must change   |
   | remove      | the goal no longer holds, or the work has since been done another way |

4. **Name the stale part** for every `revise`. "Needs updating" is not a finding. Which acceptance criterion, which
   path, which command.
5. **Record the order** among `valid` plans and what determined it — dependency, risk, or value. An order with no stated
   basis is re-litigated every pass.
6. **Delete removed plans.** History keeps them, and a removed plan left in `backlog/` will be picked up by someone
   reading the directory rather than this record.

## Verification

```sh
rtk ./hippo run --class ephemeral --disk-path . -- ./rhino plan validate
rtk ./hippo run --class ephemeral --disk-path . -- npm run check:markdown-links
```

Grooming is complete when every plan on the frozen list carries a disposition, every `revise` names what is stale, the
`valid` plans carry a stated order, and `plans/backlog/README.md` matches what is on disk.

## Recovery

If a plan cannot be judged without doing part of the work, that is a `revise` whose stale part is the unknown itself.
Record the unknown and what would resolve it; do not start the work to settle a grooming question.

## What This Does Not Do

It does not revise plans. Grooming identifies staleness; correcting it is authoring, and it runs through
[plan-planning](plan-planning.md) with the gates that implies.

It does not start execution. Selecting a plan and executing it are separate decisions, and collapsing them means the
selection was never made deliberately.
