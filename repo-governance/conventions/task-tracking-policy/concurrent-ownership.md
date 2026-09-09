---
tldr: "Refreshes shared areas before relying on them and reconciles around another task's concurrent edits."
when_to_use: "Use before relying on or editing plans/, repo-governance/, or the harness directories."
---

# Concurrent Ownership

[Task Tracking Policy](../task-tracking-policy.md)

More than one task can be working this repository at once — a second harness, a second session, or the owner editing by
hand — and `plans/`, `repo-governance/`, and the harness directories are where they collide, because those are the files
every task reaches for.

Refresh the state of those areas before relying on or editing them. Re-read the document rather than trusting what the
task list says about it, since a list records what was intended and the file records what is there.

Treat a change you do not recognize as another task's work, not as an error. Preserve it and reconcile around it: do not
revert it, do not overwrite it, and do not fold it into your own commit. When it genuinely conflicts with what you were
asked to do, that is a decision rather than a merge, and it goes to the owner under the
[grilling-with-options policy](../grilling-with-options-policy.md).

A task list grants no authority over any of this. It records intended work; committing and pushing it remain governed by
the [commit hook policy](../../development/commit-hook-policy.md).
