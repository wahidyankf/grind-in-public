---
tldr: "Reviews each outgoing range privately before a push, and fixes how a leak is remediated before and after it."
when_to_use: "Use immediately before every push, and whenever a screen or review reports a finding."
---

# Push Review

A pushed commit is on the remote, and every clone keeps it. Every push is therefore reviewed before it leaves,
privately, with nothing posted.

## Range

For each ref the push updates, the range runs from the remote's current tip of that ref to the local head; for this
repository that is `origin/main..main`. A ref the remote lacks starts where it leaves every remote ref.

## Sequence

1. **List the range's commits.** Every commit in the range, oldest first, including merges.
2. **Run the repository's screen on the range.** It screens commit by commit per [Enforcement](003-enforcement.md); a
   scan error blocks exactly as a finding does. Before the push, run it from the repository root:

   ```sh
   OSE_GATE_SURFACE=ci PUBLIC_SAFETY_BASE="$(git rev-parse origin/main)" PUBLIC_SAFETY_HEAD="$(git rev-parse main)" \
     ./hippo run --class ephemeral --resource-tier standard --disk-path . -- bash scripts/public-safety/check.sh
   ```

3. **Read every commit.** Each commit's added lines, its file names, its message, and the ref name. A merge contributes
   what it resolved beyond the automatic merge. A summary or memory of the change is not a reading, and no file is
   skipped because a screen covers it.
4. **Judge against the [leak classes](001-leak-classes.md).** No candidate is copied into notes, commands, or logs.
5. **Decide.** No finding: push. Any finding: do not push; remediate, then start again from step 1.

## Remediation

Before the push, the fix is to the history, not the tree. A later commit that deletes the value does not pass, because
the commit that added it would still be published. Rewrite the unpushed commits so that none carries the value: amend
the latest commit, or rebuild the range without it.

| Class                            | Replace the value with                                                      |
| -------------------------------- | --------------------------------------------------------------------------- |
| `secret_or_private_value`        | a reference to environment or secret storage; rotate it if it left the host |
| `protected_environment_property` | an environment variable declared in the committed template                  |
| `machine_specific_absolute_path` | a `~/` path, a repository-relative path, or a documented placeholder        |

After the push, the value is disclosed. Stop, rotate any credential, and report it to the repository owner without the
value. Rewriting published history requires the owner's explicit approval under the
[commit hook policy](../../development/commit-hook-policy.md#public-repository-safety); correcting the tree alone is
never the whole remedy.
