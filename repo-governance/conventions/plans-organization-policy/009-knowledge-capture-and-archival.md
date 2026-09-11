---
tldr: "Requires every learning to reach a durable owner or a reasoned discard before a plan is archived."
when_to_use: "Use while executing a plan and when running its final Knowledge Capture phase."
---

# Knowledge Capture and Archival

A plan that teaches something and then archives it has wasted the lesson. This repository exists to learn, so capture is
a gate rather than a courtesy.

## learnings.md Is Transient

`learnings.md` is a holding area, not a destination. It exists so that a discovery made mid-execution has somewhere to
go immediately, without stopping to decide where it belongs.

It is written in the moment something is noticed — a surprise, a wrong assumption, a rule that failed to prevent the
failure it targets — not reconstructed from memory afterwards, because reconstruction records what the author already
believed rather than what happened. Each entry is one short paragraph: what happened, and what a future reader should do
differently.

Deciding where it belongs is deferred, not skipped. Before a plan is archived, every entry is resolved: promoted to a
durable owner, or discarded with a reason. "Interesting, keep it here" is not a resolution — it archives the entry into
a folder nobody will open again.

## Durable Owners

The final phase of every substantive plan, immediately before archival, is Knowledge Capture. It promotes each entry to
exactly one owner:

| Owner                   | Fits when the learning is                                            |
| ----------------------- | -------------------------------------------------------------------- |
| governance              | a rule that should now bind future work                              |
| specification or test   | a behaviour that should fail loudly if it regresses                  |
| code comment            | context a future reader needs at that exact line                     |
| permanent documentation | something a reader under `docs/` needs to know                       |
| role instruction        | a lesson that changes how a skill or agent behaves                   |
| idea brief              | work worth considering but out of this plan's scope                  |
| discarded, with reason  | not generalizable, already covered, or simply turned out to be wrong |

One owner, not several. An entry copied into three places creates three things that can drift, and no rule for which is
authoritative when they do. A governance entry is integrated through the automatically triggered
[rules-propagation](../../workflows/rules/rules-propagation.md) workflow rather than written straight into a rule file.

Before routing, every surviving entry passes two checks: it holds no secret or sensitive detail, and it is relevant to
this repository rather than to one incident. A silent deletion is not a discard, because the next person to discover the
same thing has no way to learn it was already considered.

## Execution Review Precedes Archival

Before a plan is archived, its execution is reviewed against its substantive phases in a fixed order, and the review
records a terminal verdict. Archival is not permitted while any acceptance criterion or delivery unit is unresolved: a
filed plan reads as a finished one, and filing an unfinished plan destroys that signal.

Archival is also blocked until every `learnings.md` entry has reached a terminal state, or the plan records the explicit
escape `No generalizable learnings — <reason>`.

Substantive completion and archival stay separate. A plan can be finished and not yet filed; the reverse must not be
possible.

## The Archival Sequence

Archival is one transaction, in this order:

1. confirm the recorded execution verdict permits it;
2. move the plan folder to `plans/done/YYYY-MM-DD__<slug>`, using the completion date;
3. update every lifecycle index and every live reference to the old path;
4. run the repository's complete validation from the archived state; and
5. commit the move.

Step 3 is the one most often missed and the one most worth doing. A reference to `plans/in-progress/<slug>` that
survives archival points at nothing, and the reader who follows it concludes the plan was deleted rather than finished.

Step 4 runs _after_ the move, not before, because moving the folder is exactly what breaks links.

`learnings.md` moves with the folder and may be deleted from `done/` later, so nothing durable may depend on it
surviving. The [execution record](012-execution-record.md) is what stays.
