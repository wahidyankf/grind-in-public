---
tldr:
  "Audits human-facing documents on explicit request, hands every finding to Docs Propagation without editing, and
  audits again until two consecutive audits are clean."
when_to_use: "Use only when the owner explicitly requests a documentation audit of one change or the whole repository."
---

# Docs Quality Gate

## Purpose

Audit read-only for stale, obsolete, misplaced, and unreadable documents, and audit again after each repair until two
consecutive audits are clean. [Docs Propagation](docs-propagation.md) is the sole writer and the mandatory continuation
for every finding.

## When to Use

Run only when the owner explicitly names this gate or directs its audit. A change, a review request, or a propagation
run never authorizes it alone.

Inputs:

- `scope`: `change`, the documents one change affects, or `all`, the whole document set Docs Propagation defines.
- `change`: the revision range or working-tree change, required when `scope` is `change`.
- `max-iterations`: the ceiling on audits; default `7`.

## Steps

1. **Freeze the snapshot:** scope, revision, and uncommitted paths. A material change other than propagation's repairs
   ends the run as `input-changed`, never restarting it.
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
6. **Hand over or count a clean audit.** An audit is clean when the ledger is clear and those checks pass; otherwise the
   ledger goes to Docs Propagation. A finding only the owner can decide, such as a specification that disagrees with the
   implementation, is asked under the [grilling-with-options policy](../conventions/grilling-with-options-policy.md).
7. **Audit again** from step 1 with the same scope and a fresh snapshot. Two consecutive clean audits end the run
   `pass`. Continue only while open findings strictly decrease; when they stop, or after `max-iterations` audits, end
   `partial`, giving each remaining finding a durable owner: fixed, filed as a plan idea or backlog item, or asked the
   owner.

## Results and Handoff

The result is `pass`, `partial`, `input-changed`, or `fail`, with the number of audits run. A handoff is not a blocked
result: the caller runs Docs Propagation with the frozen ledger, then the next audit, without another request. An input
change ends the run with its ledger kept; an audit or propagation that cannot run ends it `fail`. A result authorizes no
commit or push, and the gate never edits a document.

## After a Finding

| Option                  | What happens                                                                       |
| ----------------------- | ---------------------------------------------------------------------------------- |
| verdict only            | the caller reports propagation's result; the gate does not run again               |
| repair to zero findings | the gate re-audits after propagation while open findings strictly decrease, capped |

This repository uses **repair to zero findings**, capped by `max-iterations`, which this gate declares because no shared
iteration ceiling exists here. The [Rules Quality Gate](rules-quality-gate.md) keeps verdict only and never reruns
itself.

## Why It Runs on Request

Judging whether a document is still true, needed, and readable is a reading task. Wired into every change, it produces
noise nobody reads or a pass nobody earned; propagation already refreshes each change.
