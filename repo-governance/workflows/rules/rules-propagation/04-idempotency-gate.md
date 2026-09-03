---
tldr: "Stops rules propagation without edits when the existing rule already has one objectively clear effect."
when_to_use: "Use after inventorying guidance and before choosing a canonical home or editing any rule."
---

# Idempotency Gate

Compare the intended rule with the current rule corpus before editing. Existing guidance is clear enough only when every
applicable condition leads to one required or prohibited action without relying on an unstated assumption, and all of
these facts can be cited in the corpus:

- one canonical source states the scope, trigger, and required behavior;
- authority and precedence are explicit wherever multiple sources apply;
- references are accurate, and no equivalent, overlapping, inverse, or contradictory guidance produces a different
  action for the same condition; and
- the required action has an observable outcome, and any named verification method accurately checks it.

If every criterion passes, record the supporting sources in the task report and stop. Make no change to rule-bearing
files: alternate wording, stylistic consistency, or personal preference does not justify an edit.

If a criterion fails, name the unmet criterion before editing. Make only the smallest change that satisfies it, then
rerun the gate. A successful rerun ends the workflow; do not continue into unrelated cleanup.
