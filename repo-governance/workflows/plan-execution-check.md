---
tldr: "Judges finished execution in a fixed order and returns one terminal verdict before archival."
when_to_use: "Use when every substantive delivery item is terminal and archival has not started."
---

# Plan Execution Check

## Purpose

Decide whether finished execution actually matched its plan, and record one terminal verdict that archival then depends
on.

## When to Use

Use it when every substantive item in `delivery.md` is terminal and archival has not started. Run it on explicit owner
direction. It is a review, not a repair pass.

## Prerequisites

The plan sits in `plans/in-progress/`, its working tree is clean, and its declared gates have run. Reviewing an
in-flight plan measures a moving target and says nothing.

## Steps

The evaluation runs in this fixed order. The order matters: each step assumes the previous one held, and reordering
produces findings that contradict each other.

1. **Scope.** Did the executed work match the plan's stated scope — nothing quietly added, nothing quietly dropped?
2. **Requirements.** Is every acceptance criterion terminal, and does its evidence actually establish it?
3. **Checklist evidence.** Does every completed item carry a result, and does the result match what the repository now
   contains?
4. **Gates.** Did every declared gate run and return a terminal result, including the ones that returned findings?
5. **Cleanup.** Are the task-owned artifacts gone, and is their absence proven rather than assumed?
6. **Knowledge capture.** Is every `learnings.md` entry resolved to a durable owner or discarded with a reason?

Record the verdict — `PASS`, `PASS_WITH_FINDINGS`, or `FAIL` — in the plan's execution record with the date and the
evidence each step read.

## Verification

```sh
rtk ./hippo run --class ephemeral --disk-path . -- ./rhino plan validate
rtk ./hippo run --class ephemeral --disk-path . -- npm run test:scheduled
```

The check is complete when all six steps have been evaluated and exactly one terminal verdict is recorded.

## Recovery

A `FAIL` is resolved by fixing the item and re-running this check from step 1, not by re-reading the same state more
sympathetically. There is no partial archival and no archival with a note explaining what was left open.

## Archival Is Blocked, Not Warned

An unresolved acceptance criterion, an unproven cleanup, or an unrouted learning blocks archival outright.

This is stricter than it looks, and deliberately so. A plan in `plans/done/` reads as finished — that is the entire
signal the lifecycle root carries. Filing an unfinished plan does not merely record something inaccurate; it destroys
the meaning of the root for every plan already in it.
