---
tldr:
  "Fixes the three terminal verdicts, the two-cycle repair bound, what an evidence record carries, and what resolves an
  item."
when_to_use: "Use when running a plan quality gate or recording the proof a delivery item claims."
---

# Evidence and Quality

## Terminal Verdicts

A quality gate returns exactly one of three results against a frozen snapshot of the plan:

| Verdict              | Means                                                            |
| -------------------- | ---------------------------------------------------------------- |
| `PASS`               | nothing outstanding                                              |
| `PASS_WITH_FINDINGS` | findings exist, are recorded, and none of them blocks proceeding |
| `FAIL`               | at least one finding blocks proceeding                           |

All three are terminal. "Almost passing", "passing pending a fix", and "re-run it and see" are not verdicts — they are
the absence of one, and they let work proceed on an unresolved question while appearing to have cleared a gate.

The snapshot is frozen because a gate that re-reads a changing plan is measuring a moving target and cannot say what it
verified.

## Bounded Repair

A gate does not run until it goes green. It runs once, and findings may then be repaired for at most two cycles.

At the ceiling the outcome is decided rather than retried: either the repaired plan is accepted on its merits or the
last known-good state is kept, whichever is better against the criteria that were declared before the first cycle. That
decision is recorded with its reasoning. Extending the budget because the next attempt feels close is how an unbounded
loop starts.

## Evidence Records

Evidence is a file, not a claim in conversation. Each record carries:

- the exact command that was run;
- the commit it was run against;
- when it ran;
- the result; and
- the findings, sanitized.

Sanitized means the finding is described without reproducing what made it a finding. A record that quotes a discovered
credential has published it a second time; a record that names a private host has leaked the thing the scan existed to
protect. Raw scanner output never becomes evidence. This repository publishes, so that rule is not a precaution here —
it is the difference between a scan and a disclosure.

Evidence that cannot be recorded safely is summarized structurally — how many findings, of what class, in what surface —
and the unsafe detail stays out.

Files a delivery item links live under the plan's optional `evidence/` directory.

## Manual Verification

Some claims cannot be established by automation, and a green pipeline does not become sufficient because it is
convenient.

A plan whose result has a user-facing surface routes evidence through distinct layers, records each one separately, and
remains incomplete while any applicable layer is unresolved. What each layer proves and what an assertion must state are
owned by the [exploratory and usability testing](../../workflows/exploratory-and-usability-testing.md) workflow rather
than restated here.

## Resolution Is Not a Tick

An item is resolved when its outcome is recorded, not when its box is ticked: a recovery trigger that never fires stays
unticked under a dated `Not triggered` disposition, and that has always counted as resolved.

Three dispositions resolve an item or a criterion — **met**, **not applicable**, and **accepted as permanently unmet**.
The third is for a plan finished except for evidence that no longer exists to be taken, such as a slot promoted and
retired. It holds only when all of these do:

1. the evidence is unobtainable rather than unobtained, and the disposition says why;
2. the plan's authority accepts it, as a dated decision naming what was put to them and what they chose, never the
   executor;
3. the requirement keeps its wording, so the gap stays legible;
4. the box stays unticked, because a tick means met; and
5. the execution review confirms each acceptance in its verdict.

An accepted item no longer blocks archival. The archived plan still shows the gap; it no longer implies unfinished work.
