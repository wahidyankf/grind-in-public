# Project Configuration Consistency

Make every `project.json` in this workspace express the same contract in the same style, co-locate the browser E2E
adapter with the application it tests, and write the resulting contract into governance so a third project has a rule to
read instead of two files to compare.

## Context

An audit found all three `project.json` files policy-compliant but not standardized: same requirements, different
expressions. The cause is that this repository has a rule for how Nx may be used and none for what a `project.json` must
look like. A worked target contract does exist, in
[the archived migration plan](../../done/2026-09-01__wahidyankf-www-migration/tech-docs/README.md), but it never left
that plan.

Two consequences are concrete rather than cosmetic. `badakmini-cli:test:coverage:unit` resolves to an undeclared cache
state while its `wahidyankf-www` counterpart is cached, so one ordered quick gate runs the same-named target two ways.
And `wahidyankf-www` is the only application here that owns no `test:e2e`, because its browser suite is a separate
project that `npm run test:quick` and the pre-push affected run never reach.

## Scope

Three projects, by path:

- `apps/badakmini-cli` — configuration normalized; **no Go code, command, or Nx target added or changed**.
- `apps/wahidyankf-www` — configuration normalized, and it receives the browser E2E adapter.
- `apps/wahidyankf-www-e2e` — normalized in Phase 1, then **deleted** in Phase 2; its eight step files, Playwright
  configuration, and skip baseline move into `apps/wahidyankf-www/tests/e2e/`.

Plus the root `package.json`, the scheduled CI workflow, `apps/README.md`, `apps/wahidyankf-www/README.md`, three
`specs/` documents — `specs/apps/README.md`, `specs/apps/wahidyankf-www/README.md`, and its `behavior/README.md` — the
C4 model at `specs/apps/wahidyankf-www/architecture.md`, one comment in
`apps/wahidyankf-www/tests/bdd/accessibility.steps.ts`, and five documents under `repo-governance/` and `docs/`:
`workspace-commands.md`, `testing-policy.md`, `testing-policy/tooling.md`, `nx-workspace-policy.md`, and
`docs/how-to/run-nx-workspace.md`. `nx.json` is read but not edited;
[tech-docs/file-impact.md](tech-docs/file-impact.md) lists it under Not Touched and says why.

## Approach

Four phases, each ending at a passing gate and delivering to `main`. Phase 0 records a clean baseline and the figures a
later failure is measured against. Phase 1 normalizes all three `project.json` files — declaring the missing cache
state, moving `badakmini-cli` off its hardcoded project paths onto `options.cwd`, hoisting the repeated input globs into
per-project `namedInputs`, and replacing the workspace's one bare `nx run`. Phase 2 merges the E2E project into the
application in a single commit and updates every document that names the retired project, including the C4 model. Phase
3 writes the ten-target contract into `testing-policy.md`. Phase 4 captures learnings and archives.

Three commitments run through all of it: no application or CLI behavior changes, no Gherkin scenario changes, and no
target is renamed. The last was considered and deliberately dropped — [brd.md](brd.md) records that decision with the
three others the owner settled during planning.

## Documents

- [brd.md](brd.md) — why the work is worth doing, the four decisions taken and their reasons, the non-goals, and the two
  accepted risks.
- [prd.md](prd.md) — three user stories and seven acceptance criteria, `[AC-1]` through `[AC-7]`, each stated in Gherkin
  and each plan-only.
- [tech-docs/](tech-docs/README.md) — the technical set: what was inspected, the design decisions, the file impact, the
  migration inventory, and the specification changes.
- [delivery.md](delivery.md) — the phased checklist execution reads and ticks.
- [learnings.md](learnings.md) — the running log, drained in Phase 4.

## Quality Gate

- 2026-09-03 — strict — 4 cycles — partial (6 MEDIUM open, listed below; owner accepted and directed execution)

Cycles 1 to 3 found and fixed 36 findings, including five that would have failed at execution: a resolved-inputs proof
that cannot work because `nx show project --json` never expands a `namedInputs` reference, a completeness grep that
matched the plan's own documents, an `nx affected` check running before the commit it measures, an AC-5 grep blind to
the one defect it exists to catch, and a cache probe that would have discarded three uncommitted edits. Cycle 4 reported
no CRITICAL and no HIGH.

### Open Findings

Accepted by the owner on 2026-09-03, who ended the loop and directed execution. Each is read before the phase it
touches; none blocks a gate.

- MEDIUM — `delivery.md` Phase 3 gate. The cache and outputs inspections sit in Phase 3, so a failure there forces a
  `project.json` fix into a `docs(testing):` commit. They belong in the Phase 2 gate, where that file is still
  uncommitted. Execute the inspections in Phase 2 and re-run them read-only in Phase 3.
- MEDIUM — `delivery.md` Phase 2. The module-resolution stop-and-record note is attached to the gate's `test:e2e`, but
  the first of three earlier runs is where a `"type": "module"` failure would surface. Apply the note from the first run
  onward.
- MEDIUM — `delivery.md` Phase 4 and `learnings.md`. The module-resolution branch is a third conditional `learnings.md`
  writer with no `Not triggered` disposition, and `learnings.md` says four writers while listing five.
- MEDIUM — `tech-docs/specification-changes.md`. The C4 Container View delta is described but its replacement ASCII is
  never specified, and the section's "One container." sentence is left unresolved against drawing a Playwright box.
- MEDIUM — `delivery.md` Phase 2. The `specs/apps/wahidyankf-www/README.md` item makes four edits and observes none; its
  `grep -c` for the retired name passes even if a row is deleted or a target name is wrong.
- MEDIUM — `delivery.md` Phase 1 gate. Three `wahidyankf-www-e2e` targets are edited but never run before the push,
  because that project owns no `test:quick` and `run-many` never reaches it.

## Directory Map

- [brd.md](brd.md) — business rationale, decisions, non-goals, risks.
- [prd.md](prd.md) — user stories and acceptance criteria.
- [delivery.md](delivery.md) — the phased delivery checklist and its gates.
- [learnings.md](learnings.md) — the transient running log.
- [tech-docs/](tech-docs/README.md) — the technical entry point and its companions.
