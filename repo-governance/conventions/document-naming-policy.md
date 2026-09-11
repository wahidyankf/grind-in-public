---
tldr: "Fixes how Markdown documents and their child directories are named."
when_to_use:
  "Use when creating, renaming, or splitting any Markdown document under docs/, repo-governance/, plans/, or specs/."
---

# Document Naming Policy

## Scope

This policy governs the filenames of Markdown documents in `docs/`, `repo-governance/`, `plans/`, and `specs/`, and the
names of the child directories a split document creates. Code identifiers are a separate concern and belong to the
[code style policy](../development/code-style-policy.md). Plan folder names are stage-dependent and belong to the
[plans organization policy](plans-organization-policy/001-lifecycle-and-folders.md).

## Rules

**Lowercase and hyphens.** Use lowercase letters, digits, and hyphens. No spaces, underscores, or capitals. The filename
is a path a link has to reproduce exactly, so a name that needs escaping or shifting is a name that gets mistyped.

**Name the subject, not the action.** A document is named for what it covers: `markdown-style-policy.md`,
`commit-hook-policy.md`, `001-lifecycle-and-folders.md`. A workflow is named the same way, as a domain-prefixed noun
phrase — `rules-propagation.md`, `harness-alignment.md`, `plan-quality-gate.md` — not as an imperative such as
`propagate-rules.md`. The noun form sorts into families as a directory grows, and the family prefix is what makes six
workflows readable at a glance.

**Number children when the entrypoint declares a reading order.** A document split into a directory of children names
each child for what it covers, and carries a `NNN-` prefix when, and only when, its entrypoint declares the order they
are read in. The prefix is a claim: these documents are meant to be read in this sequence, and the third assumes the
first. Every child directory under `repo-governance/workflows/` is numbered, without exception — a workflow is a
procedure, its children are its steps, and a child that reads like reference material is still consulted at a point in
the run, so it takes the number of that point rather than escaping the sequence. A convention whose modules build on
each other declares its order and is numbered on the same grounds. Leave a directory of genuinely independent topics
unnumbered: there the prefix asserts a dependency that does not exist. Inserting a module renumbers what follows;
`002a-` encodes the edit history into the reading order, and the reading order is the only thing these names are for.

**Number in three digits.** An ordinal prefix is written zero-padded to three digits: `001-inventory.md`, never `01-` or
`1-`. Ordinals are contiguous and start at `001`. The sibling repositories number the ordered companions of a split
document at the same width, so one workspace-wide width means a reader never has to remember which repository they are
in. The width is not permission to number: a directory with no declared reading order still takes plain names.

**Match the directory to the document it splits.** A child directory takes its parent document's name without the
extension, so `plans-organization-policy.md` splits into `plans-organization-policy/`. A reader who sees one can predict
the other.

## Verification

No automated check reads a filename. This policy is verified in review, and by `npm run check:markdown-links`, which
fails when a rename leaves a link pointing at the old name.
