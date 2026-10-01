---
tldr: "Enforces the push leak review through a commit-by-commit pre-push screen and a hosted replay of each push."
when_to_use: "Use when changing the leak screen, the pre-push hook, or the hosted leak screen workflow."
---

# Enforcement

A precondition stated only in prose is a convention someone forgets. This repository enforces the push leak review in
two mechanical layers beneath the review itself; each fails closed.

## The Screen

[`scripts/public-safety/`](../../../scripts/public-safety/README.md) runs on the `pre-push` surface, which RHINO feeds
Git's ref-update lines. It screens each pushed range commit by commit: each commit's added lines at the line numbers
they occupy, its file names, its message, and the ref names, then the tracked tree. A merge contributes what it resolved
beyond the automatic merge. Content the range did not add is not screened again, so the screen binds from adoption
onward. Its shapes include absolute home paths on macOS, Linux, and Windows, private addresses, and internal hostnames,
beside a credential scanner.

The screen matches shapes; the review reads context. Neither replaces the other.

## Hosted Replay

`.github/workflows/leak-screen.yml` replays the same screen over each push to `main`, from the event's previous tip to
its new one, because a local hook can be skipped and a hosted check cannot. When the previous tip is all zeros or absent
from the clone, it replays the new tip's own commit against its first parent, and screens the whole tree for a root
commit.

The push has already landed when it runs, so it is detective rather than preventive, and the push review stays
mandatory. A failure is a disclosure: follow the after-push remediation in
[Push Review](002-push-review.md#remediation).

## Adopter Decisions

- **Integration path:** direct push to `main`; no pull request.
- **Hosted check:** `Leak screen`, the range replay above, run on push rather than required by branch protection, since
  no pull request precedes the push.
- **Not adopted:** the posted review record, its marker and reviewer identity, the `leak-review` record status, and its
  `scripts/leak-review/` publisher. Each exists to gate a pull request's merge, and this repository has no pull request
  to carry or gate a record.
