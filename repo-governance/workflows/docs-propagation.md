---
tldr: "Carries one change into every human-facing document it affects, removing obsolete ones."
when_to_use:
  "Use automatically before committing a change that alters what a document describes, when adding, moving, or deleting
  a document, or when the Docs Quality Gate hands over findings."
---

# Docs Propagation

## Purpose

Carry one change into every human-facing document it affects: stale facts corrected, obsolete documents removed, each
fact kept in its one home, and the result readable by a newcomer. This is the one writer for documentation upkeep; the
[Docs Quality Gate](docs-quality-gate.md) only finds.

The document set is the root `README.md`; every README and document under `apps/`, `libs/`, `docs/`, and `scripts/`; the
`specs/` tree; and a `backlog/` or `in-progress/` plan's documents where they describe the repository. Governance and
agent instructions — `repo-governance/`, instruction files, and harness directories — stay with
[Rules Propagation](rules/rules-propagation.md) and [Harness Alignment](harness-alignment.md).

## When to Use

Entry is automatic: whoever makes a change starts here as part of the work, without a separate request, when the change
alters what a reader of a document relies on — purpose, layout, behaviour, commands, setup, dependencies, navigation, or
contributor expectations — or adds, moves, or deletes a document. A handed-over gate ledger also starts a run. Edits
made inside one run start no second one.

Inputs: the change (a revision range or the working tree) and, optionally, a frozen gate ledger.

## Steps

1. **Freeze the inputs:** the change, any ledger, the revision, and uncommitted paths. A material change ends the run as
   `input-changed`; it never restarts.
2. **Find what went stale.** Search the whole document set for every name, path, command, flag, version, and interface
   the change removed, renamed, or redefined. Each ledger row is an item too.
3. **Remove what is obsolete.** Delete a document that describes something the repository no longer has, with every link
   and index entry pointing at it. Move unique meaning that is still true to its canonical home first.
4. **Keep each fact in its one home.** The root README orients; a project README follows the
   [project README policy](../conventions/project-readme-policy.md); an index follows the
   [documentation index policy](../documentation-index-policy.md); a `docs/` page serves one
   [Diátaxis](../../docs/README.md) mode. A summary links one level down to its detail, per
   [progressive disclosure](../principles/progressive-disclosure.md); a fact with a canonical home is linked, never
   copied.
5. **Write for a newcomer.** Each affected document says from its opening what it is and why it matters, shows the next
   step without assuming the layout, and leaves no undefined term or skipped prerequisite, judged by reading against the
   [Markdown style](../conventions/markdown-style-policy.md) and [language](../conventions/language-policy.md) policies,
   never by a readability score. A sparing emoji may mark meaning to aid scanning; decoration never does.
6. **Run what is safe to run.** Execute each command and example an affected document shows through its declared entry
   point, guarded per [resource-aware development](../development/resource-aware-development.md). Never run one that
   touches a production or shared system, promotes or publishes, spends money, needs a secret, or cannot be undone; the
   document then says plainly that it was not exercised.
7. **Treat specifications as canonical,** per the [specs policy](../development/specs-policy.md). Refresh their
   readability, navigation, and links; when one disagrees with the implementation, the partial outcome applies.
8. **Change only what is stale, missing, or obsolete.** Never rewrite accurate prose, invent behaviour, or fold in
   unrelated work.
9. **Verify once** with the checks below. Repair only failures this run caused, and only while their count strictly
   decreases; stop and report when it does not.
10. **Commit with the change it explains,** per the
    [thematic commits policy](../conventions/thematic-commits-policy.md). A handed-over ledger's repairs land as their
    own commit.

## Verification

```sh
rtk npm run format:check
rtk npm run test:repo
```

These existing checks own formatting, links, indexes, frontmatter, and word limits; this workflow runs them and adds
none.

## Outputs and Recovery

Outputs: `status` (`no-change`, `landed`, `partial`, or `input-changed`), the updated and removed documents, and each
command left unexecuted with its reason.

Partial outcome: when the code, a specification, or the audience is ambiguous or they disagree, that document stays
unchanged and the owner is asked under the
[grilling-with-options policy](../conventions/grilling-with-options-policy.md); the rest lands. Repair the source of
truth first when it is wrong. A rerun on unchanged inputs changes nothing.
