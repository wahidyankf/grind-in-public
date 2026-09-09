---
tldr: "Removes exactly the development artifacts one piece of work created, and nothing else."
when_to_use: "Use once a piece of work has landed on origin/main or has been deliberately abandoned."
---

# Dev Artifact Clean-Up

## Purpose

Remove exactly the development artifacts one piece of work created, and leave local `main` level with `origin/main`.
Every rule this applies is stated elsewhere; this is the order, and the checks that make deleting safe.

## Scope

Two things here, and nothing else: the regenerable build output this work produced, and local `main`'s position against
`origin/main`.

Build output means the compiled and cached output a documented command rebuilds — `node_modules/`, `.nx/`,
`apps/*/dist/`, `.next/`, `coverage/`, `test-results/`, `playwright-report/`, `.features-gen/`, `*.tsbuildinfo`. It
never means a `.env*` file, `hippo.local.json`, `.claude/settings.local.json`, or any other local secret or machine
configuration: those are not build output, exist nowhere else, and are out of scope in every location.

The worktree and branch halves of this workflow do not apply. The
[integration path policy](../conventions/integration-path-policy.md) makes local `main` the sole integration path, so no
task branch or worktree exists to remove. Where an external tool created a temporary one for a non-integration purpose,
that policy already requires removing it the moment its purpose completes, and `prod-<project>` promotion branches are
never in scope.

Everything else on the machine belongs to someone else — another repository's state, an artifact this work did not
create. That holds even when it looks abandoned.

## When to Use

Once the work has landed on `origin/main`, or once it is deliberately abandoned. Not between commits of the same piece
of work, and never as a periodic sweep. Retain the output of a run that failed, and say so, rather than deleting the
evidence a diagnosis needs.

## Prerequisites

1. Nothing is unpushed: `git status --porcelain` is empty and `git log @{u}..HEAD` prints nothing.
2. Nothing is running against the tree — no dev server, no watch mode, no gate.
3. The work landed, or its abandonment is deliberate.

## Steps

1. Reconcile local `main`:

   ```sh
   git fetch origin --prune
   git merge --ff-only origin/main
   git rev-list --left-right --count HEAD...origin/main
   ```

2. Purge the build output this work produced, once nothing is using it. Prefer the tool's own command over `rm`, because
   the tool knows what it owns:

   ```sh
   rtk ./hippo run --class ephemeral --disk-path . -- npx nx reset
   ```

   Delete a specific directory only when it is named in the scope above and this work is what produced it.

3. Retain logs, traces, `generated-reports/` content, and any other non-regenerable evidence a failure would need.

## Verification

The divergence count reads `0 0`, the purged output is gone, and `git status --porcelain` shows no tracked file removed.
A rebuild from the documented command restores everything that was deleted.

## Recovery

If a purge removed something a later command cannot rebuild, it was not build output and this workflow was misapplied;
say so rather than reconstructing it silently. If `git merge --ff-only` refuses, local `main` carries a commit
`origin/main` lacks — stop and inspect it. Never force, and never delete `main` locally or on `origin`.

## Never

Never delete a `.env*` file or any other local secret. They are gitignored and unregenerable — nothing in the repository
reconstructs one — so deleting one is permanent loss of the owner's own configuration, not a reclaimed artifact.

Never delete an artifact another actor created. Never stash to clear the tree before purging; the stash stack is shared,
and a pop elsewhere takes an entry it did not create.
