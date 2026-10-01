---
tldr: "Indexes the modules that define a leak, review a push, and enforce the review."
when_to_use:
  "Use to decide whether a value is a leak, to run the review before a push, or to change how the review is enforced."
---

# Push Leak Review Details

Detail behind the [push leak review](../push-leak-review.md) workflow. Read in order; filenames are numbered, and the
[document naming policy](../../conventions/document-naming-policy.md) says why.

## Contents

- [Leak Classes](001-leak-classes.md) — the three classes, what is not a leak, and why history is the subject.
- [Push Review](002-push-review.md) — the private review of each outgoing range, and remediation before and after a
  push.
- [Enforcement](003-enforcement.md) — the pre-push screen, the hosted replay, and this repository's adopter decisions.
