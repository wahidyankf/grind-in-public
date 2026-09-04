---
tldr: "Defines how a plan is promoted, started, completed, archived, and reopened."
when_to_use: "Use when moving a plan between plans/ stages."
---

# Lifecycle Moves

A plan moves in one direction until it is done, and each move is a committed change that updates both indexes.

```text
ideas/ --promote--> backlog/ --start--> in-progress/ --complete--> done/
                                            ^                        |
                                            +--------reopen----------+
```

## Promoting an Idea

After the owner explicitly requests a plan and its open questions are answered, create `plans/backlog/<identifier>/`,
author the five core documents through the [plan-planning](../../workflows/plan-planning.md) workflow, delete the idea
file, and update both maps. The idea's prior art and non-goals carry into `brd.md` rather than being rewritten from
scratch.

## Starting Work

1. Move the folder from `backlog/<identifier>/` to `in-progress/<identifier>/`. No rename: neither stage carries a date.
2. Update `backlog/README.md` and `in-progress/README.md`.
3. Commit and push the move before executing any checklist item, so the repository states what is active before it
   changes.

A plan is never executed out of `backlog/`.

## Completing Work

1. Re-run the quality gate and reconcile every acceptance criterion, specification, README, gate, learning, and
   conditional task with the delivered system.
2. Record a dated, evidence-backed `Not triggered` disposition for every dormant recovery task, then run the Knowledge
   Capture phase; see [knowledge capture](knowledge-capture.md).
3. Refuse an already-existing `plans/done/YYYY-MM-DD__<identifier>/` destination; never merge, overwrite, or invent a
   suffix.
4. Rename the folder to `YYYY-MM-DD__<identifier>` using the completion date and move it to `done/`.
5. Update maps, resolve archived internal links directly, confirm the source is absent and the destination occurs once,
   then commit the move with a message naming the plan.

## Checkbox Lockstep

Tick a checkbox only after the change it describes exists. Ticking ahead of the work turns the checklist into a wish and
makes a resumed session trust a state that was never reached.

## Reopening

If a defect surfaces after archival, move the folder back to `in-progress/`, strip the date prefix, and add a dated note
in `README.md` stating what broke. A reopened plan is honest history; a quietly edited `done/` plan is not.

The [plan quality gate](../../workflows/plan-quality-gate.md) reads these canonical moves directly; no harness copy is
maintained.
