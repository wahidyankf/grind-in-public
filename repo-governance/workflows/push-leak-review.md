---
tldr: "Defines a leak and reviews every commit bound for origin, privately, before each push to main."
when_to_use: "Use immediately before every push, and whenever the pre-push screen or the hosted leak screen fails."
---

# Push Leak Review

## Purpose

A **leak** is anything in outbound history that a reader of the remote could use to reach an environment or identify the
machine it came from. [Leak Classes](push-leak-review/001-leak-classes.md) defines the three classes and what is not
one. History is the subject, not the final tree: a value one commit adds and a later commit deletes is still in every
clone. The review binds from adoption onward; history published before it is out of scope.

This repository pushes straight to `origin/main` under the
[integration path policy](../conventions/integration-path-policy.md), so no pull request exists to carry a posted review
record. The private push review is therefore the only review, and it is no less required. The standard's merge review,
its record check, and the required `leak-review` status are not adopted, because nothing here merges a pull request; see
[Enforcement](push-leak-review/003-enforcement.md).

## When to Use

Before every push to a remote, with nothing posted. A finding blocks the push. Pushing stays a separate owner-authorized
act under the [commit hook policy](../development/commit-hook-policy.md#authorization); this review grants no
permission.

## Prerequisites

Local `main` holds the commits to push, and `origin/main` is fetched, so the outgoing range is known.

## Steps

1. Determine the range: `origin/main..main`, listed oldest first with `git log --reverse origin/main..main`.
2. Run the review in [Push Review](push-leak-review/002-push-review.md): screen the range, then read every commit's
   additions, file names, and message, and judge them against the three classes.
3. No finding: push. The pre-push hook replays the commit-by-commit screen and must pass; never bypass it.
4. Any finding: do not push. Remediate the history per [Push Review](push-leak-review/002-push-review.md#remediation),
   then start again from step 1.

## Verification

The push succeeds with the pre-push screen clean, and the hosted leak screen replays the same range on `main` and
passes. No candidate value appears in any note, command, log, or report the review produced.

## Recovery

A hosted leak screen failure, or any leak found after the push, means the value is disclosed: follow the after-push
remediation in [Push Review](push-leak-review/002-push-review.md#remediation). A scan error blocks exactly as a finding
does; fix the cause rather than retrying past it.

## Example Usage

```text
Run the push leak review for the commits local main is about to push.
```
