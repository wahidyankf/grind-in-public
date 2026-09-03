# File Impact

Every path this plan expects to touch, with exactly one label: `[E]` edit, `[N]` new, `[M]` moved, `[D]` deleted. No
directory, glob, or ellipsis stands in for a filename.

## Phase 1 — Configuration Normalization

- `[E]` `apps/badakmini-cli/project.json` — declare `namedInputs.behaviorCorpus`; declare `cache: true` and
  `outputs: ["{workspaceRoot}/local-tmp/badakmini-unit.out"]` on `test:coverage:unit`; add
  `"options": {"cwd": "{projectRoot}"}` to all thirteen command targets and strip `-C apps/badakmini-cli` from each;
  rewrite the two `mkdir -p local-tmp` prefixes and the `BADAKMINI_BIN` assignment; replace every repeated corpus glob
  with `behaviorCorpus`.
- `[E]` `apps/wahidyankf-www/project.json` — declare `namedInputs.behaviorCorpus` and `namedInputs.workspaceScripts`;
  replace every repeated glob with those names; remove `outputs` from `test:coverage:integration` and from
  `generate:cv-pdf`, both of which are `cache: false`; change `static-routes:validation`'s command from `nx run …` to
  `npm exec nx -- run …` and its `cwd` from `{workspaceRoot}` to `{projectRoot}`, dropping the literal
  `apps/wahidyankf-www/` prefix from its `node` invocation.
- `[E]` `apps/wahidyankf-www-e2e/project.json` — declare `namedInputs.behaviorCorpus`, replace its five repeated globs,
  and declare `outputs` on `specs:e2e:baseline`, which is cached and writes `.features-gen`. This file is deleted in
  Phase 2; it is normalized here anyway so Phase 1 ends with all three files in one style and its gate can assert that
  property across the whole workspace rather than across two thirds of it.

## Phase 2 — Merge the E2E Project

Moved:

- `[M]` `apps/wahidyankf-www-e2e/steps/accessibility.steps.ts` →
  `apps/wahidyankf-www/tests/e2e/steps/accessibility.steps.ts`
- `[M]` `apps/wahidyankf-www-e2e/steps/cv.steps.ts` → `apps/wahidyankf-www/tests/e2e/steps/cv.steps.ts`
- `[M]` `apps/wahidyankf-www-e2e/steps/home.steps.ts` → `apps/wahidyankf-www/tests/e2e/steps/home.steps.ts`
- `[M]` `apps/wahidyankf-www-e2e/steps/personal-projects.steps.ts` →
  `apps/wahidyankf-www/tests/e2e/steps/personal-projects.steps.ts`
- `[M]` `apps/wahidyankf-www-e2e/steps/responsive.steps.ts` → `apps/wahidyankf-www/tests/e2e/steps/responsive.steps.ts`
- `[M]` `apps/wahidyankf-www-e2e/steps/search.steps.ts` → `apps/wahidyankf-www/tests/e2e/steps/search.steps.ts`
- `[M]` `apps/wahidyankf-www-e2e/steps/static-filterable-routes.steps.ts` →
  `apps/wahidyankf-www/tests/e2e/steps/static-filterable-routes.steps.ts`
- `[M]` `apps/wahidyankf-www-e2e/steps/theme.steps.ts` → `apps/wahidyankf-www/tests/e2e/steps/theme.steps.ts`
- `[M]` `apps/wahidyankf-www-e2e/e2e-skip-baseline.json` → `apps/wahidyankf-www/tests/e2e/e2e-skip-baseline.json`
- `[M]` `apps/wahidyankf-www-e2e/playwright.config.ts` → `apps/wahidyankf-www/playwright.config.ts`

Deleted:

- `[D]` `apps/wahidyankf-www-e2e/project.json`
- `[D]` `apps/wahidyankf-www-e2e/package.json`
- `[D]` `apps/wahidyankf-www-e2e/tsconfig.json`
- `[D]` `apps/wahidyankf-www-e2e/eslint.config.mjs`
- `[D]` `apps/wahidyankf-www-e2e/.gitignore`
- `[D]` `apps/wahidyankf-www-e2e/README.md`

Edited to receive the merge:

- `[E]` `apps/wahidyankf-www/playwright.config.ts` — after the move, change `steps` from `"./steps/**/*.ts"` to
  `"./tests/e2e/steps/**/*.ts"`, and correct the comment describing `test:e2e`'s `dependsOn` so it names the
  intra-project `build`. `featuresRoot` and `webServer.cwd` both stay `"../.."`-relative and are unchanged, because the
  file's depth below the workspace root is unchanged. One earlier edit lands on this file before the move, while it is
  still `apps/wahidyankf-www-e2e/playwright.config.ts`: the reconciliation of the two scenario counts its `missingSteps`
  comment states, nineteen and 34, against a measured `bddgen` run.
- `[E]` `apps/wahidyankf-www/project.json` — add `install`, `test:e2e` with its skip guard scoped to `tests/e2e`, and
  `specs:e2e:baseline` carrying its `.features-gen` `outputs`; point `test:e2e`'s `dependsOn` at the intra-project
  `build`; extend `lint:commentary` to read `tests/e2e/steps` alongside `src`; point `specs:e2e:baseline`'s inputs at
  the moved baseline file and step directory.
- `[E]` `apps/wahidyankf-www/package.json` — add `@axe-core/playwright@4.10.1`, `@playwright/test@1.62.1`, and
  `playwright-bdd@9.2.0` to `devDependencies` at the same exact pins.
- `[E]` `apps/wahidyankf-www/tsconfig.json` — add `.features-gen` to `exclude`.
- `[E]` `apps/wahidyankf-www/eslint.config.mjs` — add a second configuration block for `tests/e2e/steps/**/*.ts`
  carrying the same three `jsdoc` rules and no `jsx` parser feature.
- `[E]` `apps/wahidyankf-www/README.md` — replace the "sibling `wahidyankf-www-e2e` project" sentence, and fold in the
  deleted project's record of the browser layer whole: the four feature files the Playwright adapter deliberately does
  not bind, the recorded skip baseline of 34 stated as generated tests with the file and the target that hold it, and
  the standing "Raise the number only when a scenario is deliberately left unbound" rule.
  `apps/wahidyankf-www-e2e/README.md` is the only document in the repository that states those three, and this phase
  deletes it.
- `[E]` `package-lock.json` — regenerated by `npm install` after the workspace and dependency moves.

Edited because they name the deleted project:

- `[E]` `package.json` — `test:scheduled` replaces both `wahidyankf-www-e2e:` prefixes with `wahidyankf-www:`.
- `[E]` `.github/workflows/full-bdd.yml` — the Playwright browser install step targets `wahidyankf-www:install`.
- `[E]` `apps/README.md` — remove the third application entry.
- `[E]` `specs/apps/wahidyankf-www/README.md` — the two-projects sentence, the Process E2E adapter row, the
  skip-baseline sentence, and the two verification-command rows.
- `[E]` `specs/apps/wahidyankf-www/architecture.md` — the Scope sentence naming "one dedicated E2E project" and the
  shared-model provision it invokes, the project table row, the Container View diagram, and the paragraph below the
  diagram explaining why the adapter is drawn. The `Containers` section's opening "One container." sentence and its
  single-row container table are **not** edited, and the redrawn node drops its `[Container: ...]` stereotype so that
  the diagram stops contradicting them; [specification-changes.md](specification-changes.md) writes out the exact before
  and after blocks and states why the tension is resolved in that direction. The two later references to "the E2E
  adapter", in the Process boundary and in the Environment boundary, are deliberately **not** edited, and this entry was
  narrowed to say so rather than leaving the delivery checklist and this document disagreeing. Each names the adapter as
  a role rather than as a project, and each stays true after the merge: the adapter still starts a server process and
  speaks to it over a local port, and the moved `playwright.config.ts` still pins its tier with
  `process.env.APP_ENV ??= "test"`. Neither contains the retired project name, so the delivery item's
  `grep -c 'wahidyankf-www-e2e'` acceptance could not have observed them either way.
- `[E]` `specs/apps/wahidyankf-www/behavior/README.md` — the sentence pointing at `apps/wahidyankf-www-e2e/README.md`,
  repointed at `apps/wahidyankf-www/README.md`. Its acceptance reads the repointed target for the skip-baseline content
  the sentence claims it holds; `npm run check:markdown-links` cannot serve, because the reference is inline code rather
  than a Markdown link.
- `[E]` `repo-governance/development/workspace-commands.md` — the narrower-runs block's three `wahidyankf-www-e2e`
  lines, and the first two sentences of the paragraph below it, which explain why that project owns no `test:quick`. The
  paragraph's third sentence is kept and rewritten under the merged target name: it is the only statement in the
  repository of the once-per-machine browser install, and it stays true after the merge.
- `[E]` `repo-governance/development/testing-policy/tooling.md` — the Recorded Deviations entry naming
  `apps/wahidyankf-www-e2e`'s `target: ES2022` override and the sentence explaining it. Removed in Phase 2 rather than
  Phase 3, because Phase 2 deletes `apps/wahidyankf-www-e2e/tsconfig.json` and the deviation's subject stops existing at
  the merge; the `apps/wahidyankf-www` half of the same paragraph stays, with `Both` reworded to the surviving single
  project.
- `[E]` `specs/apps/README.md` — the Directory Map line describing the `wahidyankf-www` corpus as shared with the
  dedicated E2E project.
- `[E]` `apps/wahidyankf-www/tests/bdd/accessibility.steps.ts` — a comment citing
  `apps/wahidyankf-www-e2e/steps/accessibility.steps.ts` as the e2e-tier home of the full axe-core scan, repointed at
  `apps/wahidyankf-www/tests/e2e/steps/accessibility.steps.ts`. This is a comment inside a source file, so Phase 2
  touches one `.ts` file that is not one of the eight moved step files. No assertion, no `@covers` tag, and no step
  binding changes.

## Phase 3 — Write the Rule

- `[E]` `repo-governance/development/testing-policy.md` — state the ten-target contract, which targets are
  eligibility-dependent, that `cache` is declared explicitly wherever `targetDefaults` does not reach, that a
  single-command target declares `options.cwd`, and when `dependsOn` rather than `options.commands` expresses ordering.
- `[E]` `repo-governance/development/nx-workspace-policy.md` — one sentence linking to the contract, so a reader who
  arrives at the Nx policy to add a target is sent to the shape rule rather than inferring it.
- `[E]` `docs/how-to/run-nx-workspace.md` — the "Run the Repository Checks" trigger list already names `project.json`
  and `nx.json`; add the pointer to the contract so the how-to and the policy do not drift apart.

## The Plan's Own Documents

These two are the execution record rather than the work, and they carry a label here because Phase 4 reconciles every
labelled path against the delivery diff and would otherwise read them as paths the plan touched without saying so.
[The execution workflow](../../../../repo-governance/workflows/plan-execution/03-gates-and-pushes.md) stages a phase's
ticked checkboxes and any `learnings.md` entries with that phase's work, so both appear inside the Phase 0 to Phase 3
range that Phase 4 reconciles. The Phase 4 folder move below carries both to the archive, and that move is labelled
once, on the directory.

- `[E]` `plans/in-progress/project-json-consistency/delivery.md` — ticked checkboxes and dated Execution Record lines,
  including each phase's commit SHA. Present in every phase commit by design, Phase 0 through Phase 4.
- `[E]` `plans/in-progress/project-json-consistency/learnings.md` — entries written as they happen, then drained to a
  terminal state in Phase 4. Present in each phase commit whose items wrote an entry. Three items always write here:
  Phase 2's pre-merge typecheck record, Phase 2's measured scenario counts from `bddgen`, and Phase 3's rule-by-rule
  policy review. Three write here only if triggered: Phase 1's conditional `tests/e2e` input removal, Phase 1's
  bare-`nx run` grep control, and Phase 2's module-resolution branch on the first `wahidyankf-www:test:e2e` run. The
  stated assumption about compiler state is recorded before execution begins, and Phase 4 gives each of the three
  conditionals a dated disposition whether or not it fired.

No other plan document is expected in the delivery diff. `brd.md`, `prd.md`, `README.md`, and this technical set change
only if the plan is amended, which routes through
[plan-planning](../../../../repo-governance/workflows/plan-planning.md) and the quality gate rather than through a phase
commit.

## Phase 4 — Knowledge Capture

- `[E]` `plans/in-progress/README.md` — the active-plan index, updated at Phase 0 and again when the plan leaves.
- `[E]` `plans/done/README.md` — the archive index gains this plan.
- `[M]` `plans/in-progress/project-json-consistency/` → `plans/done/YYYY-MM-DD__project-json-consistency/` at the
  completion date.

## Not Touched

The evidence this plan writes under `local-tmp/` belongs here too, because `local-tmp/` is gitignored: none of these
files is ever staged, none appears in any delivery diff, and none carries an `[E]`, `[N]`, `[M]`, or `[D]` label. Phase
0 writes `local-tmp/projects-before.txt`, `local-tmp/targets-badakmini-before.json`,
`local-tmp/targets-www-before.json`, `local-tmp/targets-www-e2e-before.json`, `local-tmp/coverage-www-before.txt`, and
`local-tmp/coverage-badakmini-before.txt`. The Phase 1 gate writes `local-tmp/targets-badakmini-after.json`,
`local-tmp/targets-www-after.json`, and `local-tmp/targets-www-e2e-after.json`. Phase 2 writes
`local-tmp/pre-merge-sha.txt` and `local-tmp/eslint-commentary.json`. Two more, `local-tmp/badakmini-unit.out` and
`local-tmp/badakmini-integration.out`, are written by the two `badakmini-cli` coverage targets as they already are
today; Phase 1 changes only whether the first is declared as an `outputs` path, not that it is written.

Named because a reader may expect them and their absence is a decision, not an omission: `apps/badakmini-cli/**/*.go`,
every file under `specs/apps/badakmini-cli/`, every `.feature` file in either corpus,
`apps/wahidyankf-www/vitest.config.ts`, `apps/wahidyankf-www/biome.json`, the root `biome.json`, the root `.gitignore`,
and `.husky/pre-push`.

`nx.json` belongs here too, and it is the one a reader is most likely to expect on the edited side. Its six
`targetDefaults` entries are left exactly as they are, and the shared input globs are declared per project rather than
workspace-wide, for the reason recorded in [README.md](README.md). Phase 1 reads its resolved effect — every target
whose `cache` this plan leaves undeclared is reached by one of those six defaults — but reading is not editing, and no
byte of the file changes. Labelling it `[E]` would fail Phase 4's reconciliation of every labelled path against the diff
by construction, because it will not appear in any diff.
