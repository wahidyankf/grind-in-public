---
tldr: "Audits human-facing documents on explicit request and hands every finding to Docs Propagation without editing."
when_to_use: "Use only when the owner explicitly requests a documentation audit of one change or the whole repository."
---

# Docs Quality Gate

## Purpose

Return one read-only verdict with a finite ledger of stale, obsolete, misplaced, and unreadable documents.
[Docs Propagation](docs-propagation.md) is the sole writer and the mandatory continuation for every finding.

## When to Use

Run only when the owner explicitly names this gate or directs its audit. A change, a review request, or a propagation
run never authorizes it alone.

Inputs:

- `scope`: `change`, the documents one change affects, or `all`, the whole document set Docs Propagation defines.
- `change`: the revision range or working-tree change, required when `scope` is `change`.

## Steps

1. **Freeze the snapshot:** scope, revision, and uncommitted paths. A material change ends the run as `input-changed`,
   never restarting it.
2. **Bound the audit.** Under `change`, the documents the change touches and every document citing what it changed;
   under `all`, the whole document set.
3. **Audit without editing.** Decide for each document whether:
   1. every claim is true to the implementation, and every command shown was run or is marked not exercised;
   2. it still describes something the repository has; if not, it is obsolete and its resolution is removal;
   3. each fact has one home, a summary sits above its detail per
      [progressive disclosure](../principles/progressive-disclosure.md), and a `docs/` page serves one
      [Diátaxis](../../docs/README.md) mode;
   4. a newcomer learns from the opening what it is and why it matters, and finds the next step, judged by reading,
      never by a score;
   5. under `all`, or when setup changed, a reader with no prior context can follow the setup exactly as written from a
      clean checkout, each step marked smooth, frustrating, or blocking; and
   6. it agrees with its specification, which is canonical under the [specs policy](../development/specs-policy.md).
4. **Record a finite ledger** under `local-tmp/docs-quality-gate/`. Each row names the document, the gap, the required
   resolution — update, move, or remove — the evidence, and a status: open, resolved, not applicable with evidence, or
   blocked. Admit only a document that is wrong, obsolete, unreachable, or unusable by a newcomer; wording preference is
   not a finding, per [minimum sufficiency](../principles/minimum-sufficiency.md).
5. **Leave machine checks to their tools.** Formatting, links, indexes, frontmatter, and word limits belong to
   `rtk npm run format:check` and `rtk npm run test:repo`; consume their result instead of repeating them.
6. **Return the verdict.** `pass` when the ledger is clear and those checks pass; otherwise `needs-propagation` with the
   ledger. A finding only the owner can decide, such as a specification that disagrees with the implementation, is asked
   under the [grilling-with-options policy](../conventions/grilling-with-options-policy.md).

## Results and Handoff

The verdict is `pass`, `needs-propagation`, or `input-changed`. `needs-propagation` is a handoff, not a blocked result:
the caller runs Docs Propagation with the frozen ledger without another request. An input change ends the audit with its
ledger kept. A verdict authorizes no commit or push, and the gate never edits a document.

## After a Finding

| Option                  | What happens                                                                       |
| ----------------------- | ---------------------------------------------------------------------------------- |
| verdict only            | the caller reports propagation's result; the gate does not run again               |
| repair to zero findings | the gate re-audits after propagation while open findings strictly decrease, capped |

This repository uses **verdict only**, matching the [Rules Quality Gate](rules-quality-gate.md), which never reruns
itself. A second audit needs a second request.

## Why It Runs on Request

Judging whether a document is still true, needed, and readable is a reading task. Wired into every change, it produces
noise nobody reads or a pass nobody earned; propagation already refreshes each change.
