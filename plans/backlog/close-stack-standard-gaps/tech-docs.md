# Technical Design

## Verified State

Each gap was re-read on `main` at `82a5935` on 2026-09-26. All five recorded gaps are open; none was dropped. The
re-reading found two unrecorded gaps, folded into U1 and U3 because their standards own them.

- **U1** — TypeScript Standards, Gates, Lint requires "a type-aware linter's recommended type-checked rules".
  `biome.json` enables the `next`, `react`, and `test` domains only; Biome 2.5.11 also offers a `types` domain ("rules
  that require type inference"), which is not enabled. Both `eslint.config.mjs` files load `@typescript-eslint/parser`
  for JSDoc commentary rules only and enable no `@typescript-eslint` rule.
- **U1, unrecorded** — TypeScript Standards, Strict Compiler Options require `strict`, `noUncheckedIndexedAccess`,
  `noImplicitReturns`, `noFallthroughCasesInSwitch`, `noUnusedLocals`, and `noUnusedParameters` in every project.
  `tsconfig.base.json` sets only `strict`. `apps/wahidyankf-www/tsconfig.json` sets `noUncheckedIndexedAccess`,
  `noUnusedLocals`, and `noUnusedParameters` to `false` and leaves the other two unset;
  `apps/wahidyankf-www-e2e/tsconfig.json` sets only `strict`. The adapter records neither as a deviation.
- **U2** — JavaScript Standards, Gates, Type check requires `checkJs` or `// @ts-check` under `strict` and
  `noUncheckedIndexedAccess`. There are 23 tracked `.js`, `.mjs`, and `.cjs` files; none carries `// @ts-check`, no
  `tsconfig*.json` sets `checkJs`, and `apps/wahidyankf-www/tsconfig.json` sets `allowJs` but includes only `*.ts` and
  `*.tsx`. The adapter's `check scope` row reads "none enforced yet, gap".
- **U3** — Python Standards, Gates requires a formatter and a linter, and Tests requires branch coverage. The dev group
  in `apps/forum-be-python/pyproject.toml` is `httpx2`, `pyright`, and `pytest`; `project.json` has `typecheck`,
  `test:unit`, `test:e2e`, and `test:quick`, and no lint, format, or coverage target.
- **U3, unrecorded** — Python Standards, Gates, Type check requires `reportUnnecessaryTypeIgnoreComment` enabled.
  `apps/forum-be-python/pyrightconfig.json` sets `typeCheckingMode: "strict"` and nothing for that rule, and the pinned
  Pyright 1.1.414 leaves it at `"none"` in every checking mode, strict included. One type-ignore or Pyright-ignore
  comment exists under `apps/forum-be-python/` today.
- **U4** — Nx Standards, Project Graph Boundaries requires tags from a recorded vocabulary and declared constraints.
  Both web projects set `tags: []`, `forum-be-python` has no `tags` key, and `checkCommon` in
  `scripts/project-contract.mjs` actively requires `tags must be []`. The adapter's `tag vocabulary` row reads "none
  yet, gap".
- **U5** — Shell Scripts, One Declared Interpreter, recorded in the adapter as Bash with `set -euo pipefail`.
  `.github/scripts/test-hippo-bootstrap.sh` and `.github/scripts/test-pre-push-contract.sh` start with `#!/bin/sh` and
  `set -eu`; `scripts/check-repo.sh` runs both.

## U1 TypeScript Type-Aware Lint and Strict Options

Enable type-aware rules for `wahidyankf-www` and `wahidyankf-www-e2e` inside the existing `lint` targets, keeping Biome
as the style and correctness linter and Prettier as the formatter. The standard's named rules to prove are floating
promises, unsafe `any`, and non-null assertions. Two routes, chosen at the planning gate (see Open Decisions):

- Biome's `types` domain in `biome.json` — no new package; several of its rules are still in Biome's nursery group, so
  the phase first lists which standard rules it covers.
- typescript-eslint's type-checked configuration in each project's existing `eslint.config.mjs` — the parser is already
  pinned; the plugin package would be new, and every rule overlapping Biome is disabled so no finding is reported twice.

The strict options land first, because they change what the type-aware rules see. Set the five missing options in
`tsconfig.base.json`, which both projects extend, and delete the three `false` overrides from
`apps/wahidyankf-www/tsconfig.json`, so no project can opt out by omission. Each project's `typecheck` target then
reports the fallout, which is fixed in code; no option is relaxed to make a file compile. U2's `tsconfig.scripts.json`
extends the same base, so it inherits them.

## U2 JavaScript Type Check

Record the adapter's `check scope` decision, then add one type-check command covering every authored JavaScript file
under `strict` and `noUncheckedIndexedAccess`. Project-wide `checkJs` is proposed: a per-file marker leaves an unmarked
file silently unchecked. Proposed layout: a root `tsconfig.scripts.json` extending `tsconfig.base.json` with `allowJs`,
`checkJs`, `noEmit`, and `noUncheckedIndexedAccess`, including `scripts/*.mjs`, `.opencode/plugin/*.js`,
`commitlint.config.cjs`, `apps/wahidyankf-www/scripts/*.mjs`, `apps/wahidyankf-www/*.config.mjs`,
`apps/wahidyankf-www-e2e/*.config.mjs`, and `apps/wahidyankf-www-e2e/tools/*.mjs`; run by a new `check:js-types` script
in `package.json` and a `js-types` gate entry in `repo-config.yml` on the pre-commit and main surfaces. Missing JSDoc on
exported functions is added as the checker requires; a file that cannot be checked is a written waiver with its reason.

## U3 Python Pilot Gates

Add a formatter, a linter with a committed rule selection, and branch coverage to `forum-be-python`, following the
standard's examples unless the owner picks otherwise: `ruff format --check`, `ruff check` with a committed `select`
including `I`, and coverage.py with `branch = true`. Add `lint` and `test:coverage:unit` targets and extend `test:quick`
to `typecheck`, `lint`, `test:unit`, then coverage. The pilot keeps its recorded exemption from the 99% floor; whether
to set any `fail_under` is an open decision. `testing-policy/tooling.md` lists the pilot's targets and is updated with
them.

Also set `"reportUnnecessaryTypeIgnoreComment": "error"` in `apps/forum-be-python/pyrightconfig.json`. The existing
waiver either still suppresses an error, in which case it gains its reason beside it as the standard requires, or it is
stale and is deleted.

## U4 Nx Tags and Boundaries

Record a tag vocabulary and dependency constraints in the adapter, tag every project, and replace the `tags must be []`
rule in `scripts/project-contract.mjs` with a check that every discovered project carries only vocabulary tags and that
each declared `implicitDependencies` edge is allowed by the constraints. The `@nx` module-boundary lint rule stays
excluded; the adapter records this validator as the boundary enforcement for every language. Proposed vocabulary:
`type:app`, `type:e2e`, and `lang:ts` or `lang:python`, with `type:e2e` the only tag allowed to depend on `type:app`.
The validator currently reads only the two web project files, so extending it to `forum-be-python` is part of the unit.

## U5 CI Test Scripts Under Bash

Change both scripts to `#!/usr/bin/env bash` and `set -euo pipefail`, then fix whatever `pipefail` or Bash parsing
exposes. The `hippo` wrapper they exercise stays POSIX `sh`, because the tests run it as its callers do; only the test
harness changes dialect. `.github/workflows/hippo-consumer-smoke.yml` invokes the bootstrap test by path and needs no
edit if the shebang is correct.

## Open Decisions

Resolved at the planning gate, one at a time, each with options and one recommendation:

1. U1 route: Biome `types` domain (recommended if it covers the three named rules) or typescript-eslint type-checked.
2. U2 scope: project-wide `checkJs` (recommended) or per-file `// @ts-check`.
3. U3 tools: Ruff plus coverage.py (recommended) or another formatter, linter, and instrument; and whether a
   `fail_under` floor is recorded.
4. U4 vocabulary and constraints: the proposal above (recommended) or another scheme.
5. Findings fallback for U1 (lint and compiler options) and U2: fix within the phase up to a declared count, else stop
   and return to the owner.

## Specification Changes

None durable. Every `[AC-…]` in [prd.md](prd.md) is plan-only: the plan changes gates and configuration, not behaviour a
user of any application observes, so nothing enters `specs/` and no C4 view changes. Each criterion's proof is named in
`prd.md` and repeated in the matching gate of [delivery.md](delivery.md).

## Dependencies and Risks

- U3 adds Python development dependencies, and U1 may add an npm package; each needs its justification recorded under
  the dependency selection policy.
- `scripts/project-contract.test.mjs` fixtures assume `tags: []`; U4 rewrites those fixtures test-first.
- Governance edits in U2, U3, and U4 trigger Rules Propagation, and the adapter and `tooling.md` are word-budgeted.

## File Impact

A path marked with a decision number depends on that open decision. A discovery prerequisite is a set of filenames that
only running the new gate can reveal; its phase names them before editing any, and execution does not start the phase
until they are named.

**U1**

- `[E]` `biome.json` (decision 1, Biome route)
- `[E]` `apps/wahidyankf-www/eslint.config.mjs`, `apps/wahidyankf-www-e2e/eslint.config.mjs` (decision 1, ESLint route)
- `[E]` `package.json`, `package-lock.json` (decision 1, ESLint route only)
- `[E]` `tsconfig.base.json` — the five strict options
- `[E]` `apps/wahidyankf-www/tsconfig.json` — delete the three `false` overrides
- Discovery prerequisite: the TypeScript files under `apps/wahidyankf-www/src/` and `apps/wahidyankf-www-e2e/tests/`
  that the strict options or the new rules report are named in `delivery.md` by Phase 1's inventory items before any is
  edited

**U2**

- `[N]` `tsconfig.scripts.json`
- `[E]` `package.json` — the `check:js-types` script
- `[E]` `repo-config.yml` — the `js-types` gate entry
- `[E]` each of the 23 tracked JavaScript files the checker reports, from this list:
  `.opencode/plugin/rule-change-notice.js`, `apps/wahidyankf-www-e2e/eslint.config.mjs`,
  `apps/wahidyankf-www-e2e/tools/behaviour-compliance.mjs`,
  `apps/wahidyankf-www-e2e/tools/behaviour-compliance.test.mjs`,
  `apps/wahidyankf-www-e2e/tools/check-behaviour-compliance.mjs`, `apps/wahidyankf-www/eslint.config.mjs`,
  `apps/wahidyankf-www/postcss.config.mjs`, `apps/wahidyankf-www/scripts/validate-static-routes.mjs`,
  `commitlint.config.cjs`, `scripts/check-governance-structure.mjs`, `scripts/check-project-contract.mjs`,
  `scripts/check-rule-change.mjs`, `scripts/check-rule-change.test.mjs`, `scripts/check-workflow-contract.mjs`,
  `scripts/governance-structure.mjs`, `scripts/governance-structure.test.mjs`, `scripts/next-with-port.mjs`,
  `scripts/project-contract.mjs`, `scripts/project-contract.test.mjs`, `scripts/rule-change.mjs`,
  `scripts/rule-change.test.mjs`, `scripts/workflow-contract.mjs`, `scripts/workflow-contract.test.mjs`

**U3**

- `[E]` `apps/forum-be-python/pyproject.toml`, `apps/forum-be-python/uv.lock`, `apps/forum-be-python/project.json`
- `[E]` `apps/forum-be-python/pyrightconfig.json` — `reportUnnecessaryTypeIgnoreComment`
- `[E]` `apps/forum-be-python/README.md` — its target table
- `[E]` `repo-governance/development/testing-policy/tooling.md` — the pilot's target list
- Discovery prerequisite: the Python files under `apps/forum-be-python/src/` that the formatter, linter, or the
  stale-waiver report flags are named in `delivery.md` by Phase 3's inventory item before any is edited

**U4**

- `[E]` `scripts/project-contract.mjs`, `scripts/project-contract.test.mjs`, `scripts/check-project-contract.mjs`
- `[E]` `apps/wahidyankf-www/project.json`, `apps/wahidyankf-www-e2e/project.json`, `apps/forum-be-python/project.json`

**U5**

- `[E]` `.github/scripts/test-hippo-bootstrap.sh`, `.github/scripts/test-pre-push-contract.sh`

**Every unit**

- `[E]` `repo-governance/development/quality/stacks/repository-adapter.md`
- `[E]` `plans/in-progress/close-stack-standard-gaps/delivery.md` and `learnings.md` as execution proceeds

**Not touched**

- `scripts/public-safety/` — vendored, byte-identical.
- `nx.json` — no plugin, executor, or generator is added.
