---
tldr: "Defines what counts as a rule here and how strongly each wording binds."
when_to_use: "Use when writing a rule, or when judging how strongly an existing sentence binds."
---

# Rule Definition Policy

## Scope

This policy defines what a rule is and how strongly it binds. The
[rule change trigger policy](../development/rule-change-trigger-policy.md) owns which _paths_ carry rules and what a
change to one announces; this policy owns what makes a _sentence_ one. Both feed the
[rules propagation](../workflows/rules/rules-propagation.md) workflow, which integrates the rule once it exists.

## What a Rule Is

A rule is a statement that directs or constrains a decision, behaviour, standard, or procedure within a stated scope. It
tells a reader — person or agent — what is required, prohibited, expected, permitted, or forbidden.

Its scope may be named outright, or carried by where it sits: a sentence in the
[testing policy](../development/testing-policy.md) binds testing without repeating the word.

Rationale, examples, and explanation support a rule without being one. They become rules only when they carry their own
direction, which is why an example that quietly introduces a requirement is a defect rather than an illustration.

## Strength

Strength is read from the words, never from the tone:

- **Must** and **must not** state a mandatory requirement or prohibition.
- **Should** and **should not** state expected behaviour. A deviation is allowed and needs a stated reason.
- **May** states permission. It never creates an obligation, and it is not a polite "must".
- A bare imperative — "run the check", "read the policy first" — is mandatory unless its context marks it optional.

The last entry matters most here, because this repository's governance is written in the imperative throughout. A reader
who treats bare imperatives as advice is reading almost every rule at the wrong strength.

Two documents that state the same requirement at different strengths are a contradiction, not duplication: a reader
following either one behaves differently, which is the test a contradiction has to meet.

## Canonical Source

Every rule has exactly one canonical source, and that source's location fixes its level in the
[precedence order](../principles/README.md#precedence). A reference at the point of use links to that source rather than
restating it, so strength cannot drift between the copy and the original.

## Verification

No automated check reads a sentence for strength. Run the read-only
[rules quality gate](../workflows/rules-quality-gate.md) only on explicit owner direction; its non-passing ledger hands
off to propagation, which owns every repair and blocker.
