---
tldr: "Defines how a plan is promoted, started, returned, completed, archived, and reopened."
when_to_use: "Use when moving a plan between plans/ stages."
---

# Lifecycle Moves

Each move is a committed change that updates both indexes.

```text
ideas/ --promote--> backlog/ --start--> in-progress/ --complete--> done/
                       ^                    ^                         |
                       +------not-ready-----+--------reopen-----------+
```

**Promoting an idea.** After the owner explicitly requests a plan and its open questions are answered, create
`plans/backlog/<slug>/`, author the required documents through the [plan-planning](../../workflows/plan-planning.md)
workflow, delete the idea file, and update both maps. Promotion is authorship, not renaming: the brief is the input to
writing the plan, and its prior art and non-goals carry into `brd.md` rather than being rewritten from scratch.

**Starting work.** Move the folder from `backlog/<slug>/` to `in-progress/<slug>/`, update both stage READMEs, and —
when separately authorized — commit and push the move before executing any checklist item, so the repository states what
is active before it changes. Otherwise stop before execution.

**Not ready after all.** A plan that turns out not to be ready may return from `in-progress/` to `backlog/`. Returning
to `ideas/` is not a move; an idea brief and a formal plan are different documents, and going back means writing the
brief again.

**Completing work.** Require explicit owner direction for a fresh completion quality-gate run and continue only on
`PASS`, reconciling every acceptance criterion, specification, README, gate, learning, and conditional task with the
delivered system. Record a dated, evidence-backed `Not triggered` disposition for every dormant recovery task, then run
the Knowledge Capture phase; see [Knowledge Capture and Archival](009-knowledge-capture-and-archival.md). Refuse an
already-existing `plans/done/YYYY-MM-DD__<slug>/` destination — never merge, overwrite, or invent a suffix. Rename with
the completion date, move to `done/`, update maps, resolve archived internal links directly, confirm the source is
absent and the destination occurs once, then commit the move with a message naming the plan.

**Reopening.** If a defect surfaces after archival, move the folder back to `in-progress/`, strip the date prefix, and
add a dated note in `README.md` stating what broke. A reopened plan is honest history; a quietly edited `done/` plan is
not.

## Checkbox Lockstep

Tick a checkbox only after the change it describes exists. Ticking ahead of the work turns the checklist into a wish and
makes a resumed session trust a state that was never reached.

## Directory Maps

Every directory recursively under `plans/`, including each technical and asset directory, carries a `README.md` with a
`## Directory Map` that links every direct sibling file and directory other than itself. The
[documentation index policy](../../documentation-index-policy.md) does not reach `plans/`, so this rule makes plans
discoverable. A move updates its source and destination maps in the same change, so no index describes an artifact that
moved.

The [plan quality gate](../../workflows/plan-quality-gate.md) reads these canonical moves and verifies both the
stage-aware naming rule and this index requirement directly; no harness copy is maintained.
