---
tldr: "Uses local main as the sole integration path while preserving governed production promotion branches."
when_to_use:
  "Use before creating a branch or worktree, integrating or pushing work, or deciding whether to open a pull request."
---

# Integration Path Policy

## Integration

Integrate repository work only by committing on local `main` and pushing it directly to `origin/main`. Do not open a
pull request or create a task, feature, delivery, or integration branch or worktree as an alternate path. Commit and
push remain separate owner-authorized actions under the [commit hook policy](../development/commit-hook-policy.md); this
policy chooses the route and grants neither permission.

Other tasks share the same local `main`, working tree, and index; stage, commit, and push only this task's work, as
[concurrent ownership](task-tracking-policy/concurrent-ownership.md) requires.

When a push is rejected because `origin/main` has moved, run `git fetch origin`, then `git merge --no-edit origin/main`,
then push again so the hooks rerun. Merge rather than rebase: a rebase refuses a working tree holding another task's
uncommitted edits, and its autostash would move them. If the merge refuses or conflicts outside this task's paths, run
`git merge --abort` and ask the owner. Never force-push past it.

`main` is the only persistent development branch. If an external tool creates a temporary branch or worktree for a
non-integration purpose, place it only at `{repository location}/worktrees/<task>`, give it one explicit purpose, and
remove it immediately after that purpose completes or is abandoned, following
[dev artifact clean-up](../workflows/dev-artifact-clean-up.md). A sibling `*-worktrees/` path is forbidden. Never retain
the worktree as undeclared backlog or use it to bypass the direct-main route.

## Promotion Exception

Branches named `prod-<project>` are persistent deployment pointers, not development or integration branches. Create,
move, or remove one only under the [deployment policy](../development/deployment-policy.md) and separate owner
authorization. Never develop on, merge through, or open a pull request from a promotion branch.

If a future external system requires pull requests or another integration branch, stop and obtain an explicit rule
change before using it.
