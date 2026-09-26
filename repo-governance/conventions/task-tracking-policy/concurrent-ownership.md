---
tldr: "Assumes other tasks share the checkout, attributes unknown changes through their ledgers, and never takes them."
when_to_use: "Use before editing, staging, committing, or pushing, and whenever a change you did not make appears."
---

# Concurrent Ownership

[Task Tracking Policy](../task-tracking-policy.md)

More than one task can be working this repository at once — a second harness, a second session, a background process, or
the owner editing by hand. Every task integrates on the same local `main` under the
[integration path policy](../integration-path-policy.md), so they share one working tree, one index, and one branch.
Assume another task is present unless the evidence says otherwise.

## Refresh Before Relying

Any path can collide; `plans/`, `repo-governance/`, the harness directories, and root configuration such as
`package.json` collide most, because every task reaches for them. Re-read a file before relying on or editing it rather
than trusting what the task list says about it, since a list records what was intended and the file records what is
there.

At the start of a task, read the `active` [progress ledgers](progress-ledger.md) under `local-tmp/*/progress.md` and the
plans under `plans/in-progress/`. When another task already claims a path yours needs, that overlap is a decision for
the owner, not a race to win.

## Attributing an Unrecognized Change

A change this task did not make is another actor's work, never an error. Identify its owner:

1. An `active` ledger whose `Paths` covers it, or an in-progress plan whose `delivery.md` names it.
2. A commit on local `main` or `origin/main` that is newer than this task's `Base`.
3. Otherwise, the owner or a process that keeps no ledger.

Whatever the answer, preserve it: do not revert it, overwrite it, or fold it into your own commit. Record the change and
its attributed owner in this task's ledger log, and name it in the report to the owner so the concurrent work is visible
rather than silently absorbed. When it genuinely conflicts with what you were asked to do, that is a decision rather
than a merge, and it goes to the owner under the [grilling-with-options policy](../grilling-with-options-policy.md).

## Staging and Committing

The index is shared, and `git commit` records all of it. Stage only this task's paths, by name. Never use `git add -A`,
`git add .`, `git add -u`, or `git commit -a`, which sweep in another task's edits. Immediately before committing,
`git diff --cached --name-only` must list only this task's paths. If another task's paths are staged, do not unstage or
commit them; stop and ask the owner.

## Pushing

`git push` publishes every local commit `origin/main` lacks, whichever task made it. Before pushing,
`git log origin/main..HEAD` must list only commits this task is authorized to push; any other commit is a question for
the owner. The integration path policy states what to do when `origin/main` has moved.

A ledger or task list grants no authority over any of this. Committing and pushing remain governed by the
[commit hook policy](../../development/commit-hook-policy.md).
