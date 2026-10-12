---
tldr: "Requires per-instance approval for any Git command that destroys work or history, and names the additive route."
when_to_use: "Use before resetting, cleaning, force-removing, deleting a branch, rewriting commits, or force-pushing."
---

# No Destructive Git Operations

Git keeps no undo for work that was never committed, and a rewrite of pushed history reaches every clone that holds it.
Assume other agents and processes are using the same checkout and remote at the same moment.

Adapted for [direct-main delivery](../conventions/integration-path-policy.md): rewriting pushed history rewrites `main`.
The shared stash stack requires committing uncommitted work instead, per
[dev artifact clean-up](../workflows/maintenance/dev-artifact-clean-up.md).

## Authorized Task Delivery

Ongoing task authority follows [Commit Hook Policy](commit-hook-policy.md#authorization). The copied
[owned-branch sync contract](no-destructive-git-operations/001-authorized-task-branch-sync.md) is reference material:
direct-main delivery permits no integration branch/worktree or force push to `main`. Promotion pointers remain under
their separate deployment authorization.

## The Rule Is the Effect

Any Git invocation whose effect is to discard uncommitted changes, destroy work this actor did not create, rewrite
history others may hold, or remove the means of recovering any of those needs explicit owner approval for that one
instance. The spelling is irrelevant: an alias, a script, or an unlisted flag with the same effect is covered. Ask what
the command destroys and who made it, not whether it appears below.

## Common Cases

- `git push --force`, or `--force-with-lease` without an expected value
  - Destroys: remote commits absent locally
  - Use instead: `git revert` for `main`; other refs remain under their separately authorized owning policy
- **rebasing or amending commits already pushed**
  - Destroys: history others built on
  - Use instead: `git revert`
- `git reset --hard`, `git checkout -f`, `git switch --discard-changes`
  - Destroys: uncommitted changes
  - Use instead: commit first
- `git checkout -- <path>` or `git restore <path>` over edits
  - Destroys: the unstaged edits at those paths
  - Use instead: commit first
- `git clean -fd` or `git clean -fdx`
  - Destroys: untracked and ignored files
  - Use instead: `git clean -n` to preview, then delete named paths
- `git stash drop`, `git stash clear`
  - Destroys: stash entries, which then become prunable
  - Use instead: leave the entries
- `git branch -D`, `git update-ref -d`
  - Destroys: a branch, skipping the merged check
  - Use instead: `git branch -d`; for a tool's task branch,
    [dev artifact clean-up](../workflows/maintenance/dev-artifact-clean-up.md)
- **expiring the reflog and pruning at once**
  - Destroys: the recovery path itself
  - Use instead: let automatic maintenance run
- `git worktree remove --force`, deleting a worktree folder
  - Destroys: a working tree and everything uncommitted in it
  - Use instead: plain `git worktree remove`

The lease form is still a force push and still needs approval; it only refuses to overwrite commits nobody has seen.
`git clean -fdx` also removes the ignored local secrets and machine files — `.env*` files, `hippo.local.json`,
`.claude/settings.local.json` — that nothing in the repository can reconstruct.

## Asking for Approval

First look for a route that destroys nothing: a new commit, a revert, a removal without force. When none exists:

1. State the exact command as it will run.
2. State what it affects: the ref, the commits left unreachable where they can be determined, and the paths.
3. Ask a yes-or-no question and wait for the answer.
4. Run exactly what was approved. If a flag, ref, or target changes, ask again.

Approval never carries forward. The repository can change between two operations, and the earlier answer was about a
state that no longer exists.

A secret found in history is no exception:
[push review remediation](../workflows/quality/push-leak-review/002-push-review.md#remediation) puts rotation first and
requires separate owner authorization before any rewrite of shared history.

## Shared State

A temporary worktree an external tool creates shares the object database and every ref with the primary checkout, so
pruning and forced removal reach state the checkout depends on. Never pass `--ignore-other-worktrees`. A destructive
operation in a checkout or worktree this task does not use needs positive evidence that it is idle. A `prod-<project>`
branch moves only under the [deployment policy](deployment-policy.md), and moving one backwards is also a force push
under this rule.

## Prefer Additive

When a destructive and an additive operation reach the same end state, take the one that leaves a trail. A revert can be
reverted; an erased commit cannot be recalled. Before any bulk deletion, run its dry-run form.

## Verification

No automated check reads this. A hook cannot tell an approved force push from an unapproved one, so the rule is verified
in review and depends on attention.
