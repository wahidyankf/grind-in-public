---
tldr: "Indexes the detail behind the task tracking policy."
when_to_use: "Use when opening a progress ledger, or looking up how concurrent edits by another task are reconciled."
---

# Task Tracking Policy Details

Detail behind the [task tracking policy](../task-tracking-policy.md). Filenames are unnumbered; these are consulted
individually rather than performed in order, and the [document naming policy](../document-naming-policy.md) says why.

## Contents

- [Concurrent Ownership](concurrent-ownership.md) — refreshing before relying, attributing an unrecognized change to its
  owner, and staging, committing, and pushing only this task's work.
- [Progress Ledger](progress-ledger.md) — the `local-tmp/` ledger that holds a task's list outside plan execution and
  lets other tasks see its claimed paths.
