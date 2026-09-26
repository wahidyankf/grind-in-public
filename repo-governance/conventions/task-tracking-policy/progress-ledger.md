---
tldr: "Keeps each task's list in an ignored local-tmp ledger that every other task on the machine can read."
when_to_use: "Use before starting any task that is not executing a plan, and when finishing or abandoning one."
---

# Progress Ledger

[Task Tracking Policy](../task-tracking-policy.md)

A harness task list is private to one session, so a second session, harness, or process cannot see what the first is
doing. The ledger makes that work visible on the machine: it is where the task list lives, and it is what
[concurrent ownership](concurrent-ownership.md) reads to attribute a change.

## When and Where

Before the first action of any task, open `local-tmp/<task>/progress.md`, with `<task>` a short lowercase-hyphenated
name no other ledger uses. `local-tmp/` is ignored, so the ledger never enters a commit. The harness's own list may
mirror it; the ledger is authoritative because other tasks can read it.

Executing a plan is the exception. The plan's `delivery.md` is its record, and a plan under `plans/in-progress/` already
announces its work to every other task, so no ledger is opened for it.

## Contents

The ledger opens with a header another task can read at a glance:

```text
Status: active | done | abandoned
Actor: <harness and session, as precisely as it is known>
Started: <date>
Base: <HEAD commit when the task started>
Authorization: <commit and push permissions granted, or none>
Paths: <every repository path this task edits>
```

Below it come the task items, kept under the task tracking policy's granularity and sync rules, and a dated log of what
happened. Add a path to `Paths` before its first edit, not after, since the header is the claim another task reads.

## Closing

When the work lands on `origin/main` or is abandoned, set `Status` accordingly; the
[dev artifact clean-up](../../workflows/dev-artifact-clean-up.md) then removes the ledger with the rest of this work's
scratch. An `active` ledger left by a crashed session is reclaimed only as that workflow describes, never by another
task deciding it looks stale.
