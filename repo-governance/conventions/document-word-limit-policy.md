---
tldr: "Sets the word limit every governed document lives under and how a document that reaches it is fixed."
when_to_use: "Use when a governed document nears or exceeds the word limit, or when deciding how to shorten one."
---

# Document Word Limit Policy

## Scope

`npm run check:governance` holds `AGENTS.md`, `CLAUDE.md`, every `repo-governance/` document, and every harness
directory README to 750 words. Agent and command definitions are prompts, not indexes, and stay unmeasured. This policy
governs every document that check measures, including a Rules Propagation transaction.

## The Limit

Every governed document stays within 750 words. A document of 700 words or more sits in the headroom band: 50 or fewer
words are left under the cap.

The point is to plan a split rather than force one: a document with no room left is where a later fix reaches for
compression, and compression is how a clause gets deleted.

## Closing the Finding

The finding is MEDIUM, and only relocation closes it: text moved into a linked document, whole. Shortening the document
back under the band does not, because that is the edit the finding exists to prevent. An instruction file that cannot be
split still relocates, into a focused document it links to.

An index that reaches the limit is not too wordy; its directory has too many peers. Group them: create a subdirectory,
move the related documents into it, give it a README, and register that child from the parent, which then carries one
line instead of many. The [documentation index policy](../documentation-index-policy.md) owns what an index must never
do to fit.

The deterministic repository gate reports the breach. The read-only
[rules quality gate](../workflows/rules-quality-gate.md) consumes that finding, and Rules Propagation performs any
required relocation without dropping unique intent.

## Verification

```sh
npm run check:governance
```

[Workspace commands](../development/workspace-commands.md#repository-checks) lists this check beside the repository's
other deterministic validations.
