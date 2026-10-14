---
tldr:
  "Enforces the push leak review through a commit-by-commit pre-push screen and the fresh-context push review alone."
when_to_use: "Use when changing the leak screen or the pre-push hook, or when deciding how the push review is enforced."
---

# Enforcement

A precondition stated only in prose is a convention someone forgets. This repository enforces the push leak review with
one mechanical layer beneath the review itself; it fails closed.

## The Screen

[`scripts/public-safety/`](../../../../scripts/public-safety/README.md) runs on the `pre-push` surface, which RHINO
feeds Git's ref-update lines. It screens each pushed range commit by commit: each commit's added lines at the line
numbers they occupy, its file names, its message, and the ref names, then the tracked tree. A merge contributes what it
resolved beyond the automatic merge. Content the range did not add is not screened again, so the screen binds from
adoption onward. Its shapes include absolute home paths on macOS, Linux, and Windows, private addresses, and internal
hostnames, beside a credential scanner.

The screen matches shapes; the review reads context. Neither replaces the other.

## No Hosted Replay

No hosted check replays a push. The push review is the only review: a fresh-context reviewer subagent, run by the
harness before every push, refuses the push on any finding or failed spawn. A local hook can be skipped, and no hosted
check remains to catch a push afterwards. A leak found after the push is a disclosure: follow the after-push remediation
in [Push Review](002-push-review.md#remediation).

## Fresh-Context Judgement

No posted record exists to prove who judged, because this repository posts nothing. Freshness is therefore unenforced by
decision, with this reason: the harness spawns the reviewer, and a spawn failure or an unreadable range blocks the push
exactly as a finding does, so the mechanical screen still fails closed on every outcome a machine can observe. A
self-judged verdict is a rule breach, not a passing review with a caveat.

## Adopter Decisions

- **Integration path:** direct push to `main`; no pull request.
- **Hosted check:** none; the push review before each push is the only review.
- **Not adopted:** the posted review record, its marker and reviewer identity, the `leak-review` record status, and its
  `scripts/leak-review/` publisher. Each exists to gate a pull request's merge, and this repository has no pull request
  to carry or gate a record.
