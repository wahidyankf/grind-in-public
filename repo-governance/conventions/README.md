---
tldr: "Indexes stable, cross-cutting standards for repository work."
when_to_use: "Use when creating, changing, or reviewing work covered by a shared convention."
---

# Governance Conventions

This directory contains stable, cross-cutting standards that make repository work consistent. Conventions implement the
foundational principles without replacing focused development policies or repeatable workflows.

## Available Conventions

- [Rule Definition Policy](rule-definition-policy.md) — what counts as a rule here and how strongly each wording binds.
  Use it when writing a rule or judging an existing one's strength.
- [Drill Practice Policy](drill-practice-policy.md) — owner-solved, single-skill drills and useful reasoning records.
- [Agent Harness Support Policy](agent-harness-support.md) — which harnesses are supported and where each reads its
  instructions and config. Use it when adding or configuring a harness; its per-harness tables live in
  [`agent-harness-support/`](agent-harness-support/README.md).
- [Harness Capability Parity Policy](harness-capability-parity-policy.md) — canonical skill and agent bundles, thin
  adapters, semantic permission parity, and deterministic proof.
- [Agent Vocabulary](agent-vocabulary.md) — what harness, agent, instruction file, and subagent mean here. Use it when
  writing or reviewing any text about agents.
- [Agent Instruction Alignment Policy](agent-instruction-alignment-policy.md) — `AGENTS.md` as the sole rule body and
  the exact Claude import that exposes it.
- [Markdown Style Policy](markdown-style-policy.md) — source formatting for every repository Markdown file. Use it when
  creating, editing, reviewing, or formatting Markdown.
- [Document Naming Policy](document-naming-policy.md) — how Markdown documents and their child directories are named.
  Use it when creating, renaming, or splitting a document.
- [Document Word Limit Policy](document-word-limit-policy.md) — the word limit every governed document lives under, and
  how a document that reaches it is fixed. Use it before shortening one.
- [Language Policy](language-policy.md) — English as the authored language, and what is exempt. Use it when writing
  anything that lands in a file.
- [Thematic Commits Policy](thematic-commits-policy.md) — where the boundary between two commits falls. Use it when
  staging a session that touched several concerns.
- [Task Tracking Policy](task-tracking-policy.md) — how granular a task list must be and when it must be updated. Use it
  before starting or reviewing any task.
- [Grilling-With-Options Policy](grilling-with-options-policy.md) — the structured form an agent must use to resolve an
  open decision with the owner. Use it before asking the owner to decide anything; its validation checklist lives in
  [`grilling-with-options-policy/`](grilling-with-options-policy/README.md).
- [Grilling Harness Binding](grilling-harness-binding.md) — the question tool each harness uses and the Markdown
  fallback. Use it when asking a decision question or adding a harness; its per-harness detail lives in
  [`grilling-harness-binding/`](grilling-harness-binding/README.md).
- [Plans Organization Policy](plans-organization-policy.md) — how a plan is staged, named, structured, and archived
  under `plans/`. Use it when creating, executing, or archiving a plan; its detail lives in
  [`plans-organization-policy/`](plans-organization-policy/README.md).

## Adding a Convention

Add a focused convention here when a standard applies broadly across the repository and is not a foundational principle,
executable-development policy, or repeatable procedure. Keep the canonical rule in one document, link to it from concise
entry points, and let the [Rules Propagation](../workflows/rules/rules-propagation.md) workflow trigger before changing
governance.
