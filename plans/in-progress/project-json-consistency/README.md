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

Plus `nx.json`, the root `package.json`, the scheduled CI workflow, four READMEs, two `specs/` documents including the
C4 model, and three governance documents.

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

## Directory Map

- [brd.md](brd.md) — business rationale, decisions, non-goals, risks.
- [prd.md](prd.md) — user stories and acceptance criteria.
- [delivery.md](delivery.md) — the phased delivery checklist and its gates.
- [learnings.md](learnings.md) — the transient running log.
- [tech-docs/](tech-docs/README.md) — the technical entry point and its companions.
