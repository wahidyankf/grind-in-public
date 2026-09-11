---
tldr: "Fixes the four plans/ roots and the slug form each carries."
when_to_use: "Use when deciding which plans/ stage a document belongs in, or when naming a plan folder."
---

# Lifecycle and Folders

Plans live under `plans/` in exactly four roots. The spelling is fixed, not stylistic: a validator matches these names
literally, and an agent locating work relies on them without searching.

| Root                 | Holds                                         | Slug form            |
| -------------------- | --------------------------------------------- | -------------------- |
| `plans/ideas/`       | two-page briefs that are not yet formal plans | `<slug>.md`          |
| `plans/backlog/`     | formal plans that are ready but not started   | `<slug>`             |
| `plans/in-progress/` | the formal plans currently being executed     | `<slug>`             |
| `plans/done/`        | completed plans, kept as history              | `YYYY-MM-DD__<slug>` |

`in-progress` is hyphenated. `done` is not `completed`, `archive`, `archived`, or `finished`. A repository that prefers
different words does not have this convention; it has a different one, and its plans will not validate.

## Stage Purposes

**`ideas/`** holds a two-pager per idea, as a file rather than a folder. An idea is a problem worth solving that has not
earned a plan yet; see the [two-pager template](013-two-pager-template.md).

**`backlog/`** holds prepared plans that nobody is executing. A backlog plan is complete enough to start without further
authoring: it has passed the [plan-quality-gate](../../workflows/plan-quality-gate.md) workflow.

**`in-progress/`** holds active plans. Keep the count small, because a second active plan splits attention rather than
doubling output. Execution reads and ticks the plan's `delivery.md` in place, and a plan is never executed out of
`backlog/`.

**`done/`** holds finished plans as a historical record. A done plan is not casually rewritten: it records what
happened, including the parts that went badly, and its value comes from being accurate rather than tidy.

## Slug Rules

A slug is lowercase, alphanumeric, and hyphen-separated. It names the outcome, not the ticket, the quarter, or the
person: `wahidyankf-www-migration`, not `q3-cleanup` or `plan-3`.

Slugs carry no date while a plan is live. A dated slug in `backlog/` or `in-progress/` is wrong, because the date it
would carry — creation, target, estimate — is either meaningless or a commitment the plan system does not make. It also
means promoting a plan from `backlog/` to `in-progress/` is a pure move with no rename.

`done/` is the exception, and the date it carries is the **completion** date — the day the final commit landed — joined
to the slug by a double underscore: `2026-08-18__wahidyankf-www-migration`. The double underscore is what lets the date
be split off mechanically without guessing where the slug begins.

```text
good:  plans/backlog/wahidyankf-www-migration/
good:  plans/done/2026-08-18__wahidyankf-www-migration/
bad:   plans/backlog/2026-08-18__wahidyankf-www-migration/   date before completion
bad:   plans/done/2026-08-18_wahidyankf_www_migration/       single underscore, underscores
```

## Ideas Are Files, In Quadrants

An idea is a file, not a folder: `plans/ideas/q<1-4>-<priority>/<slug>.md`, kebab-case, no date. Its quadrant directory
is one of `q1-urgent-important`, `q2-not-urgent-important`, `q3-urgent-not-important`, or `q4-not-urgent-not-important`,
chosen from dated urgency and importance evidence.

This is a local extension of the canonical shape, which places briefs directly under `ideas/`. The quadrant is kept
because it is what makes a growing idea list survivable: an unsorted list of briefs is read once and then never again.

An idea is deliberately cheap, and it is exempt from the document requirements in
[Required Documents](003-required-documents.md).

## One Plan, One Root

A plan occupies exactly one root at a time. Moving it between stages is a move, not a copy: no stub, no forwarding
pointer, and no second folder left behind under the old stage. A slug appearing under two roots is a validation failure
rather than a merge to resolve, because the two copies will disagree and there is no rule for which one wins.

Movement between these roots, and the index duties each move carries, are in [Lifecycle Moves](002-lifecycle-moves.md).
