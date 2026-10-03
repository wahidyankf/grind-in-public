---
tldr: "Requires a granular task list before every task starts, updated as the work happens."
when_to_use: "Use before planning, executing, or reviewing any task."
---

# Task Tracking Policy

## Scope

This policy covers the task list kept while work is in progress: when one is required, how small its items must be, when
it must be updated, and how many delegated agents work on it at once. It applies to every supported harness; see the
[agent harness support policy](agent-harness-support.md), whatever each calls the feature.

## When a List Is Required

Before starting any task, create a task list, including work that has only one anticipated verifiable step. Mark its
first item in progress before the task's first action. A task may begin with one item and gain more as work is
discovered, but it must never begin without a current list. Outside plan execution the list lives in a
[progress ledger](task-tracking-policy/progress-ledger.md) under `local-tmp/`, where other tasks can read it; a plan's
`delivery.md` is its own record.

## Granularity

Write one item per outcome that can be checked on its own. An item is too coarse when judging it done requires accepting
several separate claims at once, and a plan step that names two verbs usually hides two items.

For new or changed application or library behaviour and bug fixes, represent each
[red-green-refactor](../workflows/quality/red-green-refactor.md) increment as separate RED, GREEN, and REFACTOR items.
Preserve the exact test path and Nx target, expected behavioural RED reason before production implementation, and
observed RED, GREEN, and REFACTOR-green results. Pure refactors follow the green-baseline and characterization rule in
[TDD](../development/tdd-policy.md).

```text
too coarse:  "Add the policy and wire it everywhere and run the checks"
granular:    "Write the policy document"
             "Link it from AGENTS.md and the category README"
             "Run the verification gates"
```

Prefer the smaller split when unsure. A list that is too fine costs a line of output; a list that is too coarse hides
how much work remains.

## Keeping It in Sync

The list must describe the present, not a plan written once and abandoned:

- Mark an item in progress before its first action, not after it succeeds.
- Mark it completed only when its outcome holds and has been verified. A failing gate leaves the item in progress.
- Record work discovered mid-task as new items instead of widening an existing one, so the count reflects the true
  remaining scope.
- Update the list as each item resolves. Marking several items complete in one batch at the end reports a state that was
  never observed.

## New Direction Mid-Task

New, follow-on, or changed direction reaches the list before it reaches the work: reconcile it against every open item,
record that reconciliation, then continue. Appending the new work alone is insufficient. See
[new direction](task-tracking-policy/new-direction.md).

## Delegated Agents

At most three delegated agents run at once besides the main thread, counted at any depth, foreground or background, in
every harness; work beyond the cap waits for one to return. See the
[delegated agent cap](task-tracking-policy/delegated-agent-cap.md).

## Concurrent Ownership

More than one task can be working this repository at once, sharing one working tree, index, and local `main`. Read the
active ledgers before starting, attribute any change you did not make to its owner instead of reverting it, and stage,
commit, and push only your own work. See [concurrent ownership](task-tracking-policy/concurrent-ownership.md).

## Why

The owner reads the list to know what is done, what is left, and what went wrong, and cannot see the reasoning behind
it. A stale or coarse list therefore misreports the work rather than merely describing it briefly. Granular items also
make an interrupted session resumable, because the first unfinished item states exactly where to restart.

## Verification

No automated gate can read a harness task list or an ignored ledger, since both are local state. This policy is verified
in review: the list is compared against the change and the commands actually run. Announcements that a rule change
occurred are separate; see the [rule change trigger policy](../development/rule-change-trigger-policy.md).
