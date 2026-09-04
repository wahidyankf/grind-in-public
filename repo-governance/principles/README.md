---
tldr: "Indexes the foundational principles and states the precedence order governance documents follow."
when_to_use: "Use before creating a governance document or resolving a conflict between two of them."
---

# Governance Principles

These principles are the foundation of `repo-governance/`. All governance policies, workflows, and future documents must
follow them. When a conflict is found, resolve it before adopting the lower-level guidance; do not leave contradictory
rules in place.

## Principles

- [Maintenance Value](maintenance-value.md) — require every maintained surface to prove more recurring benefit than
  upkeep cost.
- [Governance Continuity](governance-continuity.md) — preserve workflow state and authorization through compaction,
  handoff, and resumed work.
- [Minimum Sufficiency](minimum-sufficiency.md) — choose the smallest responsible change and stop after required
  verification passes.
- [Progressive Disclosure](progressive-disclosure.md) — keep guidance focused and load only what the task needs.
- [Root Cause Orientation](root-cause-orientation.md) — fix the responsible cause rather than suppressing its symptoms.

## Precedence

Governance is ordered. A lower level may add detail to a higher one, and may never contradict it:

```text
principles > conventions > development > workflows
```

A document at the root of `repo-governance/`, other than the entry index, ranks with conventions. A child directory
ranks with the document it splits, because it holds that document's own detail rather than a weaker rule.

The order exists so that most conflicts settle without a round trip. When two documents at different levels disagree,
the higher one stands and the lower one changes: that resolution is mechanical, and an agent performs it. When they sit
at the same level, nothing ranks them, so choosing between them is a decision only the owner can make.
[Conflict resolution](../workflows/rules/rules-propagation/03-conflict-resolution.md) owns both paths and states what
each one requires.

Precedence settles which document changes. It never settles whether a rule is right: a higher-level document that is
wrong is amended deliberately, not worked around by an exception written below it.
