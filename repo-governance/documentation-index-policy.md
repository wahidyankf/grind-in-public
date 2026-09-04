---
tldr: "Defines recursive README indexes and concise discovery metadata."
when_to_use:
  "Use when adding, moving, renaming, or reviewing content in docs/, repo-governance/, scripts/, plans/, specs/, or a
  harness directory."
---

# Documentation Index Policy

## Scope

Every directory in `docs/`, `repo-governance/`, `scripts/`, `plans/`, `specs/`, and each agent harness directory —
`.agents/`, `.claude/`, `.codex/`, and `.opencode/` — at every depth, must contain a `README.md`. The README is the
directory's concise entry point for both people and agents.

A harness directory holds tool configuration rather than prose, so its README indexes what the directory contains and
what each entry does. The `tldr` and `when_to_use` requirement below does not apply there, because those files carry the
frontmatter their tool defines.

Two exemptions apply. A skill directory needs none: its `SKILL.md` names the skill and when to use it, and
`skills/README.md` registers it. Second, a harness registers some directories by filename, so an index placed there
becomes a command or an agent; add the README only where the tool ignores it or offers a flag that keeps it inert, and
otherwise index that directory from its parent. The [agent harness support policy](conventions/agent-harness-support.md)
records the verified behaviour per directory; check it, and test the tool, first.

## Required Indexing

Each README must register its immediate Markdown documents, excluding itself, with a descriptive relative link. It must
also register each immediate child directory through that directory's README. Child READMEs own their descendants; do
not repeat the full recursive tree.

In `specs/`, an index registers every immediate entry whatever its file type, because the Gherkin corpus is the content
there and a `.feature` file left out of the map is behaviour nobody can find. Elsewhere the requirement covers Markdown
and child directories.

`plans/` and `specs/` gain the index and nothing else: no word limit reaches either tree, and neither needs the
frontmatter below. An archived plan keeps the index it was delivered with, and reopening one restores the duty to update
it.

Every Markdown document under `docs/` or `repo-governance/`, except `repo-governance/README.md`, must begin with YAML
frontmatter containing `tldr` and `when_to_use`. Keep both values short enough to help a reader decide whether to open
the document.

`repo-governance/README.md` is the exception because it is the governance entry index.

`badakmini-cli:test:repo` checks indexes and frontmatter deterministically. The semantic
[rules quality gate](workflows/rules-quality-gate.md) consumes that result rather than reproducing it, while the
[plan quality gate](workflows/plan-quality-gate.md) requires the same repository target before `PASS`.

## When an Index Reaches the Word Limit

Never make an index fit by dropping entries or their descriptions. An incomplete index hides work, the failure this
policy exists to prevent.

The [document word limit policy](conventions/document-word-limit-policy.md) states the limit, which documents it
governs, and how a document that has reached it is fixed; an index is fixed the same way.

## Maintenance

Create the README with a new directory. When adding, moving, renaming, or removing Markdown content or a child
directory, update every affected parent README in the same change. Keep each entry brief, and say when a reader should
open what it links to.

This policy implements [progressive disclosure](principles/progressive-disclosure.md): indexes make focused documents
discoverable without turning a parent README into a duplicate manual.
