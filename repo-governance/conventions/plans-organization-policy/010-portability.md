---
tldr: "Keeps the plan system self-contained and records this repository's deviations from it."
when_to_use: "Use when importing a plan rule from elsewhere or recording how this repository diverges."
---

# Portability

This specification is complete on its own. A repository that has read it can run the whole plan lifecycle without access
to any other repository, service, or index.

That is a constraint on how the specification may be written, not just a description of it.

## No External Dependencies

The plan system must not require:

- a coordinating or control-plane repository;
- a shared plan index, registry, or ledger held somewhere else;
- naming any particular repository, organization, group, or team;
- a maintainer's absolute path, machine, or account;
- an internal host, address, network, or topology; or
- a private artifact, convention, or process a reader cannot see.

A rule that cannot be stated without one of those is not portable, and belongs to the repository that needs it rather
than to this specification. This repository is public, so that is a publication rule as much as a design one.

Examples and placeholders are semantic: `<slug>`, `plans/in-progress/<slug>/delivery.md`, `YYYY-MM-DD`. They stand for a
shape, and a reader substitutes their own — never a real path from elsewhere with the identifying parts lightly changed.

## Delivery Mode Is the Adopter's

The plan system deliberately says nothing about how changes reach a repository's main line. Pull requests, direct
commits, and local-only commits in a repository with no remote at all are equally compatible, because `delivery.md`
declares its execution checkout and delivery units rather than assuming them.

This repository's route is direct commits on local `main` pushed to `origin/main`, with no task branch, worktree, or
pull request; the [integration path policy](../integration-path-policy.md) owns it. That artifact is this repository's
own. It applies alongside this convention and is not a variant of it.

## Deviation Is Allowed, Silence Is Not

An adopting repository may diverge. What it may not do is diverge invisibly.

For each block of this specification, an adopter records exactly one of:

| Status                      | Means                                                                 |
| --------------------------- | --------------------------------------------------------------------- |
| adopted                     | applied as written                                                    |
| adapted, with reason        | applied in a different form, and here is why                          |
| not applicable, with reason | this repository's characteristics make it irrelevant, and here is why |

A reason is required for the latter two, and "we do it differently" is not one — it restates the status. A future reader
has to be able to tell a considered divergence from an incomplete adoption, and those look identical from the outside.
An absent row is itself a finding: a repository that has adopted nine blocks and never mentions the tenth has not
adopted nine blocks; it has an unknown state.

## The Label That Was Dropped

This repository carried a third executor label, `[AI+HUMAN]`, for work an agent prepares and the owner approves. It is
gone. The shared validator emits `PLAN-DELIVERY-002` for it, so a live plan using it could not be validated at all — and
a label that turns a gate into something that cannot run is not a local flavour worth keeping. Its halves are now
written as what they are: `[AI]` preparation under a recorded authorization, and a `[HUMAN]` action.

Plans under `plans/done/` keep it. They are validated for lifecycle and slug form only, so it costs nothing there, and
rewriting it would falsify the record.

## This Repository's Deviations

| Block             | Status  | Reason                                                                                                                                                                                                                                        |
| ----------------- | ------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| diagrams          | adapted | Every diagram here is ASCII in a fenced `text` block, not Mermaid, per the [markdown style policy](../markdown-style-policy.md). This repository's output is read in terminals and plain-text feeds, where a Mermaid block renders as source. |
| idea layout       | adapted | Briefs are filed under quadrant directories rather than directly in `ideas/`, because an unsorted brief list stops being read. See [Lifecycle and Folders](001-lifecycle-and-folders.md).                                                     |
| executor labels   | adopted | The local third label was dropped for the canonical two; see the section above.                                                                                                                                                               |
| module numbering  | adapted | The canonical lifecycle block is split across modules `001` and `002`, because a single module would exceed this repository's 750-word document limit.                                                                                        |
| execution record  | adapted | `delivery.md` additionally opens with a dated execution record; see [Execution Record](012-execution-record.md). This is an addition, and it removes nothing.                                                                                 |
| every other block | adopted | Applied as written.                                                                                                                                                                                                                           |
