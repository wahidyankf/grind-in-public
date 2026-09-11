---
tldr: "Sets the bar a new repository-wide check must clear, and where each kind of check belongs."
when_to_use: "Use when adding, moving, or removing a recurring repository validation check."
---

# Repository Check Policy

A repository check is anything that runs on every change to keep the repository itself correct: a gate, a hook step, a
validation script. This policy decides whether one should exist, who should implement it, and where it runs.

## Restraint

**Do not create a check merely because a rule can be automated.** First reuse an existing command, hook, standard tool,
or clear manual review. A new check requires explicit repository-owner approval plus evidence that the failure is
recurring, that it materially affects repository-wide correctness, and that centralised enforcement returns more value
than its code, tests, documentation, execution time, false positives, upgrades, and eventual removal will cost.

That last item is the one usually left out of the estimate, and it is rarely the smallest. Every check written is a
check someone later has to prove is obsolete before they can delete it.

## Who Implements It

| Kind of check                               | Owner                                                                     |
| ------------------------------------------- | ------------------------------------------------------------------------- |
| reads documents — structure, links, budgets | a [RHINO](https://github.com/wahidyankf/rhino) entry in `repo-config.yml` |
| reads the repository's own shape or wiring  | a script under `scripts/`, beside its own test                            |
| runs an external tool at a pinned version   | a `tool` directive in `tools/go.mod`, or an equivalent pin                |

A check that reads documents is a configuration entry, not new code. Reach for a script only when the rule is about this
repository's own arrangement and no declared validator expresses it.

## How It Must Behave

Keep validation deterministic and offline, so it can run in a Git hook without a network. Inspect the Git-tracked state
when the check concerns committed content, rather than the working tree, so what is checked is what would be published.

A check that only reports must never block, and must never be able to. A notice that can fail closed will eventually
stop an unrelated commit, and the fix will be to remove the notice.

## Where It Runs

Wire a check that **blocks** into pre-push, scoped to the paths that can break it, so a push that cannot fail the check
does not pay for it. A check that no path narrows runs on every push, which is a cost worth stating out loud.

Wire a check that only **reports** into pre-commit, where the author can still act on what it says. A notice that
arrives at push time asks for an amend rather than an edit.

State the hook and its scope in the [workspace commands](workspace-commands.md) hook summary, which is canonical for
what each hook runs. A hook whose behaviour is documented only in the hook file is a hook nobody reads.

## Removing One

A check is removed the same way it was added: with evidence. Before deleting it, name what now owns each behaviour it
had, or state the ground on which that behaviour no longer needs an owner — the subject is gone, or another check
already covers it. "It seemed redundant" is not a ground.

Update every reference in the same change that removes the check. A rule that names a deleted command is worse than no
rule, because it reads as current.
