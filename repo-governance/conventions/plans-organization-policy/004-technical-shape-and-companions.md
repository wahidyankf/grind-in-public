---
tldr: "Fixes the one technical shape a plan carries and how its ordered companions are named."
when_to_use: "Use when authoring tech-docs, adding a companion, or converting between the two shapes."
---

# Technical Shape and Ordered Companions

A plan carries exactly one technical shape:

| Shape       | Form                                          | Choose when                                                             |
| ----------- | --------------------------------------------- | ----------------------------------------------------------------------- |
| single file | `tech-docs.md`                                | the technical design fits in one document a reader will actually finish |
| directory   | `tech-docs/README.md` plus ordered companions | it does not                                                             |

Never both. A plan root containing `tech-docs.md` _and_ `tech-docs/` is a validation failure, not a transitional state,
because the two will diverge and nothing decides which is authoritative.

Growing out of the single-file shape is normal. Converting means deleting `tech-docs.md` in the same change that creates
`tech-docs/`, so the plan never holds both.

## What the Technical Shape Owns

Context, architecture, selected decisions, dependencies, risks, and — in the directory shape — reading order and links
to each companion. It owns no checklist; that is `delivery.md`.

Every non-archived plan states its file impact: every expected path marked exactly `[E]` edit, `[N]` new, `[M]` moved,
or `[D]` deleted. In the directory shape that is `tech-docs/file-impact.md`; in the single-file shape it is a section of
`tech-docs.md`. Add the companions named by the [specification-change](014-specification-changes.md),
[migration](015-plan-migrations.md), and [UI-design](016-plan-ui-design.md) rules when they apply; never create an empty
companion to satisfy a heading.

## Companion Naming

When the shape is a directory, `tech-docs/README.md` is the entrypoint and every other document is a companion named:

```text
NNN-descriptive-kebab-case.md
```

- `NNN` is exactly three digits, zero-padded: `001`, `002`, … `010`.
- Numbering starts at `001` and is contiguous. No gaps, no duplicates, no `000`.
- The remainder is lowercase, alphanumeric, hyphen-separated, and describes the content. `007-release-packaging.md`, not
  `007-part-7.md` and not `007-misc.md`.
- The numeric order **is** the reading order.

Three digits rather than two because a companion set that outgrows `99` would otherwise have to be renumbered wholesale,
and one that outgrows `9` under two digits sorts `10` before `2` in every tool that sorts lexically.

## The Entrypoint Declares the Order

`tech-docs/README.md` lists every companion as a link, in numeric order. That list is the contract: a companion on disk
but absent from the list, or listed but absent from disk, is a failure.

The list exists because the ordinal in a filename tells a reader the sequence but not the reason for it. The entrypoint
is where the sequence is explained.

## Renumbering Is Atomic

Inserting a companion between `003` and `004` renumbers everything from `004` upward, updates the entrypoint list, and
updates every inbound link — in one change. There is no `003a`, no `003.5`, and no reuse of a number freed by a
deletion.

This costs more than appending, and that cost is the point: it keeps the ordinals meaningful instead of turning them
into arbitrary identifiers that happen to be numeric.

## The Same Rule Elsewhere

Any ordered companion set this repository governs — not only plan technical documents — uses this naming. A reader who
learns it once should not have to learn a second variant. The module you are reading is one such set.

## Diagrams Are ASCII

A diagram in any plan document is a fenced `text` block, not Mermaid. This repository publishes to readers who may be in
a terminal, a plain-text feed, or a renderer that does not run Mermaid, and a diagram nobody can see is worse than the
sentence it replaced. See [Portability](010-portability.md) for the full reasoning and its consequences.
