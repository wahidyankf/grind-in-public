---
tldr: "Gives every idea brief one disposition, with a reason for every keep and retire."
when_to_use: "Use when plans/ideas/ has accumulated briefs and someone needs to know which still matter."
---

# Plan Ideas Grooming

## Purpose

Reduce `plans/ideas/` to briefs that are still worth something, each carrying a stated disposition.

## When to Use

Use it when `plans/ideas/` holds at least one brief and nobody can say which of them are live. Run it on explicit owner
direction; it is not triggered by adding a brief.

## Prerequisites

The working tree is clean enough to commit the deletions this produces. Nothing else is required: grooming reads briefs
and writes dispositions, and it consults no build.

## Steps

1. **Freeze the list.** Enumerate every brief under `plans/ideas/q*/` before judging any of them. A brief added during
   grooming belongs to the next pass — otherwise the list never closes.
2. **Read each brief once.** A brief that cannot be understood in one reading has a defect worth recording; that is
   itself a grooming outcome.
3. **Assign exactly one disposition** per brief:

   | Disposition | Means                                                           |
   | ----------- | --------------------------------------------------------------- |
   | promote     | worth planning now; it becomes a formal plan in the backlog     |
   | keep        | still worth doing eventually, and here is what would trigger it |
   | retire      | not worth doing, and here is why                                |

4. **Record the reason** for every `keep` and every `retire`. A `keep` without a trigger is indistinguishable from
   indecision, and it will be re-read at every future pass at the same cost.
5. **Re-file a misplaced brief.** A `keep` whose urgency or importance has changed moves to the quadrant its dated
   evidence now supports, and the move updates both quadrant indexes.
6. **Promote by authoring, not by moving.** A promoted brief is the input to [plan-planning](plan-planning.md); the
   brief itself does not become the plan. Delete the brief only once the plan exists.
7. **Delete retired briefs.** Git is the archive. A retired brief left in place will be reconsidered by someone who does
   not know it was already rejected.

## Verification

```sh
rtk ./hippo run --class ephemeral --disk-path . -- npm run check:markdown-links
```

Grooming is complete when every brief on the frozen list carries a disposition, every `keep` and `retire` carries a
reason, and every quadrant `README.md` matches what is on disk.

## Recovery

If a disposition cannot be reached because the brief is too vague to judge, retire it with that as the reason. Rewriting
it is authoring — the same work as writing it the first time — and it belongs to whoever wants the idea back.

## What This Does Not Do

It does not size, schedule, or sequence work. A promoted brief has been judged worth planning; whether it is planned
next is [plan-backlog-grooming](plan-backlog-grooming.md)'s question.
