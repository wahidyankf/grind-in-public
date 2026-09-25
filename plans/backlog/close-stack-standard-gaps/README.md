# Close Stack-Standard Gaps

**Status: Backlog.** Filed 2026-09-26. No planning gate has run, no quality gate has run, and nothing here is authorized
for execution. The owner resolves the open decisions in [tech-docs.md](tech-docs.md#open-decisions) at the planning gate
before this plan moves to `in-progress/`.

Close the five compliance gaps that adopting the software-development stack packs (commit `10a3859`) recorded but
deliberately did not fix, so the
[repository adapter](../../../repo-governance/development/quality/stacks/repository-adapter.md) can drop its "Known
gaps" sentence and every adopted pack is enforced by a gate rather than by review alone.

## Context

The gaps were found during a cross-repository standards adoption. Adoption recorded them in the adapter's Deviations
section as "left for a separate plan rather than fixed by adoption"; this is that plan. Each was re-verified on `main`
at `82a5935` before filing — see [tech-docs.md](tech-docs.md#verified-state).

## Scope

Five delivery units, one per gap:

| Unit | Gap                                                          | Projects touched                                                                          |
| ---- | ------------------------------------------------------------ | ----------------------------------------------------------------------------------------- |
| U1   | TypeScript lint runs Biome without type-aware rules          | `wahidyankf-www`, `wahidyankf-www-e2e`                                                    |
| U2   | no JavaScript file is type-checked through JSDoc (`checkJs`) | `repo-scripts`, `opencode-plugin`, `wahidyankf-www`, `wahidyankf-www-e2e`, root configs   |
| U3   | the Python pilot has no formatter, linter, or coverage gate  | `forum-be-python`                                                                         |
| U4   | Nx project `tags` are empty, so no boundary can be enforced  | `wahidyankf-www`, `wahidyankf-www-e2e`, `forum-be-python`, `repo-scripts` (the validator) |
| U5   | the two `.github/scripts/` tests run under `sh`, not Bash    | `ci-scripts`                                                                              |

Each unit ends by editing the adapter to remove its gap. No application behaviour changes, and no Gherkin scenario
changes.

## Approach

Phase 0 records the baseline. Phases 1 to 5 deliver U1 to U5, one unit per phase; the units share no file except the
adapter, so they may be reordered at the planning gate without rework. Each phase ends at a gate and delivers to `main`.
Phase 6 captures learnings and archives. A unit that uncovers a finding backlog it cannot clear inside its phase stops
at its checkpoint and returns to the owner rather than waiving the finding.

## Documents

- [brd.md](brd.md) — why the gaps are worth closing, non-goals, and risks.
- [prd.md](prd.md) — user stories and acceptance criteria `[AC-1]` to `[AC-6]`.
- [tech-docs.md](tech-docs.md) — verified state, design per unit, open decisions, and file impact.
- [delivery.md](delivery.md) — the phased checklist.
- [learnings.md](learnings.md) — the running log, drained in Phase 6.

## Directory Map

- [brd.md](brd.md) — business rationale, non-goals, risks.
- [prd.md](prd.md) — user stories and acceptance criteria.
- [tech-docs.md](tech-docs.md) — the technical design and file impact.
- [delivery.md](delivery.md) — the phased delivery checklist and its gates.
- [learnings.md](learnings.md) — the transient running log.
