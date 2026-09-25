# Delivery

## Execution Record

<!--
One dated line per phase completion, gate pass or failure, retry that proved
something, and plan change, written as the event happens. Each phase-completion
line carries the commit SHA that phase pushed.
-->

## Executor Tags

`[AI]` an agent can fully perform it. `[HUMAN]` only the owner, for a credential, a physical action, an external
authority, or a decision whose information does not exist yet.

## Execution Checkout, Units, and Pause Safety

- **Checkout:** the primary checkout on local `main`, delivered by direct push to `origin/main` under the
  [integration path policy](../../../repo-governance/conventions/integration-path-policy.md). No branch or worktree.
- **Delivery units:** U1 to U5, one per phase, each with one outcome (its `[AC-…]`), one owner (the executor), and one
  rollback (revert that phase's commits). The adapter edit that retires a gap lands in the same phase as its gate.
- **Pause safety:** every phase ends at a green gate with its commits pushed and a dated Execution Record line, so a
  cold executor resumes at the first unticked item.

Nx target runs go through the `./hippo` guard verbatim as written. Never wrap an `npm run` script in a second guard;
those scripts carry their own.

## Phase 0: Planning Gate and Baseline

- [ ] [HUMAN] Resolve the five [open decisions](tech-docs.md#open-decisions) one at a time; the owner holds them.
      Acceptance: each decision's choice is recorded in `tech-docs.md` and every "(decision N)" path in its File Impact
      is fixed. [AC-1] [AC-2] [AC-3] [AC-4] [AC-7]
- [ ] [AI] Move `plans/backlog/close-stack-standard-gaps/` to `plans/in-progress/`, update both stage READMEs, and
      commit and push the move when separately authorized. Acceptance:
      `git ls-files plans/backlog/close-stack-standard-gaps` prints nothing and the in-progress index links the plan.
- [ ] [AI] Run `npm install`. Acceptance: exits 0 and `git status --short` shows no lockfile change.

### Phase 0 Gate

- [ ] [AI] `npm run test:quick`, `npm run format:check`, `npm run check:hygiene`, and `npm run test:repo` each exit 0.
      Acceptance: all four green before any change; record the result in the Execution Record.

> **Pause Safety**: nothing but the plan has changed. Safe to stop. Resume with `npm run test:quick`.

## Phase 1: U1 TypeScript Type-Aware Lint and Strict Options

- [ ] [AI] Add `noUncheckedIndexedAccess`, `noImplicitReturns`, `noFallthroughCasesInSwitch`, `noUnusedLocals`, and
      `noUnusedParameters`, each `true`, to `tsconfig.base.json`, and delete the three `false` overrides from
      `apps/wahidyankf-www/tsconfig.json`. Acceptance: `grep -c ': false' apps/wahidyankf-www/tsconfig.json` prints `0`
      and `grep -c noUnusedParameters tsconfig.base.json` prints `1`. [AC-7]
- [ ] [AI] Inventory the fallout: run
      `rtk ./hippo run --class ephemeral --resource-tier standard --disk-path . -- npm exec -- nx run -p wahidyankf-www -t typecheck`
      and the same for `wahidyankf-www-e2e`. Acceptance: every reporting file is named here with its error count; over
      the decision-5 ceiling, revert both files and stop for the owner. [AC-7]
- [ ] [AI] Fix each reported error in the files the inventory named, without relaxing an option or changing behaviour.
      Acceptance: both `typecheck` targets exit 0, and
      `rtk ./hippo run --class ephemeral --resource-tier light --disk-path . -- npm exec -- tsc --showConfig -p apps/wahidyankf-www-e2e/tsconfig.json`
      and the same for `apps/wahidyankf-www/tsconfig.json` print all six options `true`. [AC-7]
- [ ] [AI] Prove the options bite: add an unused local in a temporary, uncommitted file under
      `apps/wahidyankf-www/src/`, run the `typecheck` target, then delete the file. Acceptance: non-zero naming the
      unused local, then 0. [AC-7]
- [ ] [AI] Inventory: apply the decision-1 configuration locally and run
      `rtk ./hippo run --class ephemeral --resource-tier standard --disk-path . -- npm exec -- nx run -p wahidyankf-www -t lint`
      and the same for `wahidyankf-www-e2e`. Acceptance: every reported file is named here with its finding count; if
      the total exceeds the ceiling fixed at decision 5, revert the configuration and stop for the owner. [AC-1]
- [ ] [AI] Fix each reported finding in the files the inventory named, without changing behaviour. Acceptance: both lint
      targets exit 0 and both projects' `test:quick` exit 0. [AC-1]
- [ ] [AI] Prove the rule fires: add an unawaited promise-returning call in a temporary, uncommitted file under
      `apps/wahidyankf-www/src/`, run the lint target, then delete the file. Acceptance: lint exits non-zero naming the
      floating promise, then exits 0 after deletion. [AC-1]
- [ ] [AI] Edit `repo-governance/development/quality/stacks/repository-adapter.md`: remove the TypeScript clause from
      the known-gaps sentence and name the type-aware gate. Run Rules Propagation. Acceptance:
      `npm run check:governance` exits 0. [AC-6]

### Phase 1 Gate

- [ ] [AI] `npm run test:quick`, `npm run format:check`, and `npm run check:hygiene` exit 0. [AC-1] [AC-7]

> **Pause Safety**: both TypeScript projects compile under the strict options and lint is type-aware; nothing else
> changed. Safe to stop. Resume with `npm run test:quick`.

## Phase 2: U2 JavaScript Type Check

- [ ] [AI] Record the decision-2 check scope in the adapter's Adopter Decisions row. Acceptance: the row no longer reads
      "gap". [AC-2] [AC-6]
- [ ] [AI] Create `tsconfig.scripts.json` (or the decision-2 equivalent) and the `check:js-types` script in
      `package.json`. Inventory with `npm run check:js-types`. Acceptance: the checked-file count equals
      `git ls-files '*.js' '*.mjs' '*.cjs' | wc -l` minus written waivers, and every reporting file is named here; over
      the decision-5 ceiling, stop for the owner. [AC-2]
- [ ] [AI] Add JSDoc types and runtime guards the checker requires in the named files. Acceptance:
      `npm run check:js-types` exits 0 and `npm run test:repo` exits 0. [AC-2]
- [ ] [AI] Add the `js-types` gate entry to `repo-config.yml`. Acceptance: `npm run validate:config` exits 0 and
      `npm run check:hygiene` lists `js-types`. [AC-2]
- [ ] [AI] Prove the check fires with a temporary, uncommitted wrong-type argument in `scripts/rule-change.mjs`, then
      revert. Acceptance: non-zero, then 0. [AC-2]
- [ ] [AI] Remove the JavaScript clause from the adapter's known-gaps sentence; run Rules Propagation. Acceptance:
      `npm run check:governance` exits 0. [AC-6]

### Phase 2 Gate

- [ ] [AI] `npm run check:js-types`, `npm run test:quick`, `npm run test:repo`, and `npm run check:hygiene` exit 0.
      [AC-2]

> **Pause Safety**: every authored JavaScript file is type-checked. Safe to stop. Resume with `npm run check:js-types`.

## Phase 3: U3 Python Pilot Gates

- [ ] [AI] Add the decision-3 formatter, linter, and coverage instrument to the dev group in
      `apps/forum-be-python/pyproject.toml` with `uv lock`, recording the dependency-selection justification in the
      commit message. Acceptance: `uv.lock` pins each exactly. [AC-3]
- [ ] [AI] Commit the rule selection and `branch = true` in `pyproject.toml`; inventory reported files and name them
      here. Acceptance: the list is recorded before any source edit. [AC-3]
- [ ] [AI] Add `lint` and `test:coverage:unit` targets to `apps/forum-be-python/project.json` and extend `test:quick` to
      `typecheck`, `lint`, `test:unit`, `test:coverage:unit`. Fix the named findings. Acceptance:
      `rtk ./hippo run --class ephemeral --resource-tier standard --disk-path . -- npm exec -- nx run -p forum-be-python -t test:quick`
      exits 0. [AC-3]
- [ ] [AI] Prove the gate fires with a temporary unused import in a pilot source file, then revert. Acceptance:
      non-zero, then 0. [AC-3]
- [ ] [AI] Set `"reportUnnecessaryTypeIgnoreComment": "error"` in `apps/forum-be-python/pyrightconfig.json`, then run
      `rtk ./hippo run --class ephemeral --resource-tier standard --disk-path . -- npm exec -- nx run -p forum-be-python -t typecheck`.
      Give the one existing waiver its reason if it still suppresses an error, or delete it if Pyright reports it
      unnecessary. Acceptance: `typecheck` exits 0 and every remaining waiver carries a reason on its line. [AC-8]
- [ ] [AI] Prove the setting bites with a temporary type-ignore comment on a line with no type error, then revert.
      Acceptance: `typecheck` exits non-zero naming the unnecessary comment, then 0. [AC-8]
- [ ] [AI] Update the pilot's target list in `repo-governance/development/testing-policy/tooling.md` and the target
      table in `apps/forum-be-python/README.md`; remove the Python clause from the adapter's known-gaps sentence; run
      Rules Propagation. Acceptance: `npm run check:governance` and `npm run check:markdown-links` exit 0. [AC-6]

### Phase 3 Gate

- [ ] [AI]
      `rtk ./hippo run --class ephemeral --resource-tier standard --disk-path . -- npm exec -- nx run -p forum-be-python -t test:quick`,
      `npm run test:quick`, and `npm run check:hygiene` exit 0. [AC-3] [AC-8]

> **Pause Safety**: the pilot formats, lints, measures branches, and reports stale waivers. Safe to stop. Resume with
> the pilot's `test:quick`.

## Phase 4: U4 Nx Tags and Boundaries

- [ ] [AI] Record the decision-4 vocabulary and constraints in the adapter. Acceptance: the `tag vocabulary` row no
      longer reads "gap". [AC-4] [AC-6]
- [ ] [AI] RED: in `scripts/project-contract.test.mjs`, add failing cases for an untagged project, an unknown tag, and a
      forbidden dependency. Acceptance: `node --test scripts/project-contract.test.mjs` fails on exactly those cases.
      [AC-4]
- [ ] [AI] GREEN: replace the `tags must be []` rule in `scripts/project-contract.mjs` and extend
      `scripts/check-project-contract.mjs` to read `apps/forum-be-python/project.json`. Acceptance: the test suite
      passes. [AC-4]
- [ ] [AI] REFACTOR: simplify without changing results. Acceptance: the suite still passes. [AC-4]
- [ ] [AI] Tag `apps/wahidyankf-www/project.json`, `apps/wahidyankf-www-e2e/project.json`, and
      `apps/forum-be-python/project.json`. Acceptance: `node scripts/check-project-contract.mjs` exits 0. [AC-4]
- [ ] [AI] Remove the tags clause from the adapter's known-gaps sentence, naming the validator as boundary enforcement;
      run Rules Propagation. Acceptance: `npm run check:governance` exits 0. [AC-6]

### Phase 4 Gate

- [ ] [AI] `npm run check:project-contract`, `npm run test:quick`, and `npm run test:repo` exit 0. [AC-4]

> **Pause Safety**: every project is tagged and the validator enforces the constraints. Safe to stop. Resume with
> `npm run check:project-contract`.

## Phase 5: U5 CI Test Scripts Under Bash

- [ ] [AI] Change `.github/scripts/test-hippo-bootstrap.sh` and `.github/scripts/test-pre-push-contract.sh` to
      `#!/usr/bin/env bash` and `set -euo pipefail`, fixing what the change exposes. Acceptance: `head -n 2` on each
      prints both lines, and each exits 0 when run directly. [AC-5]
- [ ] [AI] Remove the last clause from the adapter's known-gaps sentence, deleting the sentence; run Rules Propagation.
      Acceptance: the [AC-6] proof in `prd.md` prints nothing and `1`. [AC-6]

### Phase 5 Gate

- [ ] [AI] `npm run test:repo`, `npm run check:workflows`, and `npm run check:hygiene` exit 0. [AC-5] [AC-6]

> **Pause Safety**: all five gaps are closed. Safe to stop. Resume with `npm run test:repo`.

## Phase 6: Knowledge Capture and Archival

- [ ] [AI] Route every `learnings.md` entry to one durable home under
      [knowledge capture](../../../repo-governance/conventions/plans-organization-policy/009-knowledge-capture-and-archival.md).
      Acceptance: each entry names its terminal state.
- [ ] [AI] Re-run every `[AC-…]` proof against `main`. Acceptance: all eight hold; results recorded in the Execution
      Record.
- [ ] [HUMAN] Direct a plan execution check; the owner holds that authority. Acceptance: a verdict that permits
      archival.
- [ ] [AI] Move the plan to `plans/done/YYYY-MM-DD__close-stack-standard-gaps/` with the completion date, update both
      stage READMEs, and commit and push when separately authorized. Acceptance: `npm run check:markdown-links` exits 0.

### Phase 6 Gate

- [ ] [AI] `npm run test:quick`, `npm run test:repo`, and `npm run check:markdown-links` exit 0.

> **Pause Safety**: the plan is archived. Nothing to resume.
