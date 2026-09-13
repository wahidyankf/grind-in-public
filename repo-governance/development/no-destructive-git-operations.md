---
tldr: "Requires per-instance approval for any Git command that destroys work or history, and names the additive route."
when_to_use: "Use before resetting, cleaning, force-removing, deleting a branch, rewriting commits, or force-pushing."
---

# No Destructive Git Operations

Git keeps no undo for work that was never committed, and a rewrite of pushed history reaches every clone that holds it.
Assume other agents and processes are using the same checkout and remote at the same moment.

Adapted from the catalog rule of the same name. What changed: work lands on local `main` and is pushed straight to
`origin/main` under the [integration path policy](../conventions/integration-path-policy.md), so rewriting pushed
history means rewriting `main` itself; and because the stash stack is shared, as
[dev artifact clean-up](../workflows/dev-artifact-clean-up.md) warns, the rows for uncommitted work point at a commit
rather than a stash.

## The Rule Is the Effect

Any Git invocation whose effect is to discard uncommitted changes, destroy work this actor did not create, rewrite
history others may hold, or remove the means of recovering any of those needs explicit owner approval for that one
instance. The spelling is irrelevant: an alias, a script, or an unlisted flag with the same effect is covered. Ask what
the command destroys and who made it, not whether it appears below.

## Common Cases

| Operation                                                             | Destroys                                        | Use instead                                                             |
| --------------------------------------------------------------------- | ----------------------------------------------- | ----------------------------------------------------------------------- |
| `git push --force`, or `--force-with-lease` without an expected value | remote commits absent locally                   | `--force-with-lease=<ref>:<expected-sha> --force-if-includes`, approved |
| rebasing or amending commits already pushed                           | history others built on                         | `git revert`                                                            |
| `git reset --hard`, `git checkout -f`, `git switch --discard-changes` | uncommitted changes                             | commit first                                                            |
| `git checkout -- <path>` or `git restore <path>` over edits           | the unstaged edits at those paths               | commit first                                                            |
| `git clean -fd` or `git clean -fdx`                                   | untracked and ignored files                     | `git clean -n` to preview, then delete named paths                      |
| `git stash drop`, `git stash clear`                                   | stash entries, which then become prunable       | leave the entries                                                       |
| `git branch -D`, `git update-ref -d`                                  | a branch, skipping the merged check             | `git branch -d`                                                         |
| expiring the reflog and pruning at once                               | the recovery path itself                        | let automatic maintenance run                                           |
| `git worktree remove --force`, deleting a worktree folder             | a working tree and everything uncommitted in it | plain `git worktree remove`                                             |

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

A secret found in history is no exception: the [commit hook policy](commit-hook-policy.md#public-repository-safety) puts
rotation first and requires separate owner authorization before any rewrite of shared history.

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
