---
tldr: "Indexes the detail behind the task tracking policy."
when_to_use:
  "Use when opening a progress ledger, spawning a delegated agent, taking new direction mid-task, or looking up how
  concurrent edits by another task are reconciled."
---

# Task Tracking Policy Details

Detail behind the [task tracking policy](../task-tracking-policy.md). Filenames are unnumbered; these are consulted
individually rather than performed in order, and the [document naming policy](../document-naming-policy.md) says why.

## Contents

- [Concurrent Ownership](concurrent-ownership.md) — refreshing before relying, attributing an unrecognized change to its
  owner, and staging, committing, and pushing only this task's work.
- [Delegated Agent Cap](delegated-agent-cap.md) — at most three live delegated agents besides the main thread, counted
  at any depth and in every harness, with the rest waiting.
- [New Direction](new-direction.md) — reconciling new, follow-on, or changed direction against every open item before
  acting on it.
- [Progress Ledger](progress-ledger.md) — the `local-tmp/` ledger that holds a task's list outside plan execution and
  lets other tasks see its claimed paths.
