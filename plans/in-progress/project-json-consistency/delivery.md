# Delivery

## Execution Record

<!--
One dated line per phase completion, gate pass or failure, retry that proved
something, and plan change. Written as the event happens, never reconstructed
at archival.
-->

## Executor Tags

`[AI]` an agent can fully perform it. `[HUMAN]` only the owner. `[AI+HUMAN]` an agent prepares and the owner approves or
performs the irreversible step.

## Phase 0: Baseline

Records a clean starting state so a later failure is attributable to the work rather than to the machine. It commits the
plan and nothing else.

- [ ] [AI] Run `npm install` from the repository root — acceptance: exits 0 and `git status --short` reports no change
      to `package-lock.json`.
- [ ] [AI] Run `mkdir -p local-tmp && npx nx show projects > local-tmp/projects-before.txt` — acceptance: the file
      contains exactly `wahidyankf-www-e2e`, `wahidyankf-www`, and `badakmini-cli`.
- [ ] [AI] Capture the resolved target configuration of each project:
      `npx nx show project badakmini-cli --json > local-tmp/targets-badakmini-before.json`, then the same for
      `wahidyankf-www` and `wahidyankf-www-e2e` into `local-tmp/targets-www-before.json` and
      `local-tmp/targets-www-e2e-before.json` — acceptance: three files exist and each parses, verified with
      `node -e 'for (const f of ["badakmini","www","www-e2e"]) JSON.parse(require("fs").readFileSync("local-tmp/targets-"+f+"-before.json"))'`
      exiting 0.
- [ ] [AI] Record the pre-change unit coverage figure by running
      `npx nx run wahidyankf-www:test:coverage:unit 2>&1 | tail -20 > local-tmp/coverage-www-before.txt` — acceptance:
      the file contains a `All files` row with a line percentage, and the command exited 0.
- [ ] [AI] Record the same for the Go side with
      `npx nx run badakmini-cli:test:coverage:unit 2>&1 | tail -5 > local-tmp/coverage-badakmini-before.txt` —
      acceptance: the file contains the `unit statement coverage:` line and the command exited 0.
- [ ] [AI] Record the current skipped-scenario count with `npx nx run wahidyankf-www-e2e:specs:e2e:baseline` —
      acceptance: exits 0, which is the assertion that the generated `test.fixme` count equals the 34 recorded in
      `apps/wahidyankf-www-e2e/e2e-skip-baseline.json`.
- [ ] [AI] Commit the plan folder and the index update with message `docs(plans): start project-json-consistency`, then
      push to `main` — acceptance: `git status --short` is empty and `git log origin/main -1 --oneline` names the
      commit.

### Phase 0 Gate

> Every check below passes before Phase 1 begins. A failure is fixed inside Phase 0.

- [ ] [AI] `npm run test:quick` — acceptance: exits 0 with both `badakmini-cli` and `wahidyankf-www` reported as
      successful tasks.
- [ ] [AI] `npm run test:integration` — acceptance: exits 0 for both projects.
- [ ] [AI] `npx nx run badakmini-cli:test:e2e` — acceptance: exits 0.
- [ ] [AI] `npm run format:check` — acceptance: exits 0 over the new plan documents.
- [ ] [AI] `npm run check:markdown-links` — acceptance: exits 0. Run
      `git add -N plans/in-progress/project-json-consistency` first, because the check reads Git-tracked files and a new
      document is invisible to it otherwise.

> **Pause Safety**: the plan is committed and pushed, and every gate is green at the unmodified baseline recorded in
> `local-tmp/`. No workspace configuration has changed. Safe to stop. Resume with `npm run test:quick`.

## Phase 1: Configuration Normalization

Brings all three `project.json` files into one style. `wahidyankf-www-e2e` is normalized here even though Phase 2
deletes it, so this phase ends with the property true across the whole workspace and its gate can assert it
workspace-wide.

- [ ] [AI] [AC-3] Edit `apps/badakmini-cli/project.json`: add a top-level
      `"namedInputs": {"behaviorCorpus": ["{workspaceRoot}/specs/apps/badakmini-cli/behavior/**/*.feature"]}` beside
      `"targets"` — acceptance: `npx nx show project badakmini-cli --json` exits 0, proving Nx still parses the file.
- [ ] [AI] [AC-3] Edit `apps/badakmini-cli/project.json`: in every target whose `inputs` array contains the literal
      `{workspaceRoot}/specs/apps/badakmini-cli/behavior/**/*.feature`, replace that string with `"behaviorCorpus"` —
      acceptance: `grep -c 'specs/apps/badakmini-cli/behavior' apps/badakmini-cli/project.json` prints `1`, the single
      occurrence being the `namedInputs` declaration.
- [ ] [AI] [AC-3] Edit `apps/badakmini-cli/project.json`: remove the `{workspaceRoot}/apps/badakmini-cli/tests/e2e/**/*`
      input from `test:coverage:behavior`, `test:coverage`, and `test:quick` — acceptance: the path lies inside
      `{projectRoot}` and is therefore already covered by the built-in `default` input, and the Phase 1 gate's
      before-and-after inputs comparison shows the resolved input set for those three targets is unchanged. If the
      comparison shows a difference, restore the input and record the finding in `learnings.md` instead of forcing the
      removal.
- [ ] [AI] [AC-2] Edit `apps/badakmini-cli/project.json`: add `"cache": true` to `test:coverage:unit` together with
      `"outputs": ["{workspaceRoot}/local-tmp/badakmini-unit.out"]` — acceptance:
      `npx nx show project badakmini-cli --json` reports `cache` as `true` for that target, matching its
      `wahidyankf-www` counterpart. The `outputs` declaration is not optional here: an undeclared `cache` means
      uncached, so this turns caching on for a gate that runs fresh today, and the profile it writes sits outside
      `{projectRoot}` where the built-in `default` output detection would not find it. Prove the pair together by
      running the target twice — `npx nx run badakmini-cli:test:coverage:unit`, then
      `rm -f local-tmp/badakmini-unit.out` and run it again — acceptance: the second run reports a cache hit and
      `ls local-tmp/badakmini-unit.out` succeeds afterwards, which is what proves the cached target restores its
      artifact rather than reporting a success that produced nothing.
- [ ] [AI] [AC-4] Edit `apps/badakmini-cli/project.json`: add `"options": {"cwd": "{projectRoot}"}` to `build`,
      `governance`, `markdown-links`, `capability-parity`, `rule-change`, `lint`, `typecheck`, `test:unit`,
      `test:integration`, `test:coverage:behavior` — acceptance: each of those ten targets carries the `options` object
      and `npx nx show project badakmini-cli --json` still parses.
- [ ] [AI] [AC-4] Edit the same ten targets in `apps/badakmini-cli/project.json` to strip the `-C apps/badakmini-cli`
      argument from their `go` invocations — acceptance: `npx nx run badakmini-cli:typecheck` and
      `npx nx run badakmini-cli:lint` both exit 0, proving the commands resolve from the project directory.
- [ ] [AI] [AC-4] Edit `apps/badakmini-cli/project.json` target `test:coverage:unit`: add
      `"options": {"cwd": "{projectRoot}"}`, strip `-C apps/badakmini-cli`, and change the leading `mkdir -p local-tmp`
      to `mkdir -p ../../local-tmp` — acceptance: `npx nx run badakmini-cli:test:coverage:unit` exits 0, prints the
      `unit statement coverage:` line, and `ls local-tmp/badakmini-unit.out` succeeds from the repository root while
      `ls apps/badakmini-cli/local-tmp` fails.
- [ ] [AI] [AC-4] Apply the identical three edits to `apps/badakmini-cli/project.json` target
      `test:coverage:integration` — acceptance: `npx nx run badakmini-cli:test:coverage:integration` exits 0, prints the
      `integration statement coverage:` line, and `ls local-tmp/badakmini-integration.out` succeeds from the repository
      root.
- [ ] [AI] [AC-4] Edit `apps/badakmini-cli/project.json` target `test:e2e`: add `"options": {"cwd": "{projectRoot}"}`,
      strip `-C apps/badakmini-cli`, and change `BADAKMINI_BIN="$PWD/apps/badakmini-cli/dist/badak-mini"` to
      `BADAKMINI_BIN="$PWD/dist/badak-mini"` — acceptance: `npx nx run badakmini-cli:test:e2e` exits 0, which requires
      the built binary to be found at the rewritten path.
- [ ] [AI] [AC-3] Edit `apps/wahidyankf-www/project.json`: add a top-level `"namedInputs"` declaring `behaviorCorpus` as
      `["{workspaceRoot}/specs/apps/wahidyankf-www/behavior/**/*.feature"]` and `workspaceScripts` as
      `["{workspaceRoot}/scripts/next-with-port.mjs"]` — acceptance: `npx nx show project wahidyankf-www --json`
      exits 0.
- [ ] [AI] [AC-3] Edit `apps/wahidyankf-www/project.json`: in `test:unit`, `test:coverage:unit`,
      `test:coverage:behavior`, `test:coverage`, and `test:quick`, replace the two literal workspace paths with
      `"behaviorCorpus"` and `"workspaceScripts"` — acceptance:
      `grep -c 'specs/apps/wahidyankf-www/behavior' apps/wahidyankf-www/project.json` prints `1` and
      `grep -c 'scripts/next-with-port.mjs' apps/wahidyankf-www/project.json` prints `3`, the three being the
      `namedInputs` declaration and the two `dev` and `start` command strings that genuinely invoke the script.
- [ ] [AI] [AC-2] Edit `apps/wahidyankf-www/project.json`: delete the `"outputs": ["{projectRoot}/coverage"]` line from
      `test:coverage:integration` — acceptance: that target is `cache: false`, so the declaration was inert;
      `npx nx run wahidyankf-www:test:coverage:integration` exits 0 and still writes `apps/wahidyankf-www/coverage`.
- [ ] [AI] [AC-5] Edit `apps/wahidyankf-www/project.json` target `static-routes:validation`: change the command's
      leading `nx run wahidyankf-www:build --skip-nx-cache` to `npm exec nx -- run wahidyankf-www:build --skip-nx-cache`
      — acceptance: `npx nx run wahidyankf-www:static-routes:validation` exits 0 and prints
      `Verified static build output for /, /cv, /personal-projects, /robots.txt, /sitemap.xml.`
- [ ] [AI] [AC-4] Edit the same target: change `"options": {"cwd": "{workspaceRoot}"}` to `{"cwd": "{projectRoot}"}` and
      shorten `node apps/wahidyankf-www/scripts/validate-static-routes.mjs` to `node scripts/validate-static-routes.mjs`
      — acceptance: the target prints the same `Verified static build output` line, and
      `grep -n 'apps/wahidyankf-www' apps/wahidyankf-www/project.json` returns no line inside a `"command"` string. Its
      `dependsOn` placement in `test:quick` is untouched; only the working directory moves.
- [ ] [AI] [AC-2] Edit `apps/wahidyankf-www/project.json`: delete the
      `"outputs": ["{projectRoot}/public/wahidyankf-kresna-fridayoka-cv.pdf"]` line from `generate:cv-pdf` — acceptance:
      that target is `cache: false`, so the declaration is inert for exactly the reason `test:coverage:integration`'s
      is; `npx nx run wahidyankf-www:generate:cv-pdf` exits 0 and still writes the PDF to `apps/wahidyankf-www/public/`.
      Removing one and leaving the other would make the Phase 3 rule false of the file it was derived from.
- [ ] [AI] [AC-3] Edit `apps/wahidyankf-www-e2e/project.json`: add a top-level `"namedInputs"` declaring
      `behaviorCorpus` as `["{workspaceRoot}/specs/apps/wahidyankf-www/behavior/**/*.feature"]`, and replace all five
      literal occurrences in `install`, `typecheck`, `lint`, `test:e2e`, and `specs:e2e:baseline` with
      `"behaviorCorpus"` — acceptance:
      `grep -c 'specs/apps/wahidyankf-www/behavior' apps/wahidyankf-www-e2e/project.json` prints `1`.
- [ ] [AI] [AC-2] Edit `apps/wahidyankf-www-e2e/project.json` target `specs:e2e:baseline`: add
      `"outputs": ["{projectRoot}/.features-gen"]` — acceptance: the target is `cache: true` and its command runs
      `bddgen`, which writes that directory, so without the declaration a cache hit replays the baseline comparison and
      restores nothing. This is the same pairing the `badakmini-cli` coverage target gets above, and it is applied here
      even though Phase 2 deletes the file, because Phase 1's gate asserts the property across all three and Phase 2
      carries the declaration into the merged target.
- [ ] [AI] Run `npm run format` — acceptance: exits 0, and `npm run format:check` afterwards also exits 0 over the three
      edited files.

### Phase 1 Gate

> Every check below passes before Phase 2 begins. A failure is fixed inside Phase 1.

- [ ] [AI] [AC-2] Run this check for each of `badakmini-cli`, `wahidyankf-www`, and `wahidyankf-www-e2e`, substituting
      the project name:
      `npx nx show project badakmini-cli --json | node -e 'let s="";process.stdin.on("data",d=>s+=d).on("end",()=>{const t=JSON.parse(s).targets;const bad=Object.entries(t).filter(([,v])=>v.cache===undefined).map(([k])=>k);console.log(bad.length?"UNDECLARED: "+bad.join(", "):"all "+Object.keys(t).length+" targets declare cache");process.exit(bad.length?1:0)})'`
      — acceptance: each of the three runs exits 0 and prints the `all N targets declare cache` form, naming a non-zero
      N so the check is proved to have read a populated target set.
- [ ] [AI] [AC-3] Capture the resolved configuration again into `local-tmp/targets-badakmini-after.json`,
      `local-tmp/targets-www-after.json`, and `local-tmp/targets-www-e2e-after.json` using the Phase 0 commands, then
      `diff` each against its `-before.json` — acceptance: every difference is one this phase's checklist named, and no
      target's resolved `inputs` array lost the behavior-corpus path or the `next-with-port.mjs` path it carried before.
- [ ] [AI] [AC-4] `grep -n 'apps/badakmini-cli' apps/badakmini-cli/project.json` — acceptance: every line printed is a
      `namedInputs` or `inputs` path, and no line is inside a `"command"` string.
- [ ] [AI] [AC-5]
      `grep -nE '"command": *"[^"]*[^-] nx run' apps/badakmini-cli/project.json apps/wahidyankf-www/project.json apps/wahidyankf-www-e2e/project.json`
      — acceptance: prints nothing and exits non-zero, and the paired
      `grep -c 'npm exec nx -- run' apps/wahidyankf-www/project.json` prints a non-zero count, proving the pattern was
      searched against a file that really does invoke Nx.
- [ ] [AI] `npm run test:quick` — acceptance: exits 0 for both projects.
- [ ] [AI] `npm run test:integration` — acceptance: exits 0 for both projects.
- [ ] [AI] `npx nx run badakmini-cli:test:e2e` — acceptance: exits 0.
- [ ] [AI] `npx nx run badakmini-cli:test:coverage` — acceptance: exits 0 and prints all three coverage lines, with the
      unit figure matching `local-tmp/coverage-badakmini-before.txt`.
- [ ] [AI] `npx nx run wahidyankf-www:test:coverage` — acceptance: exits 0 and the unit line percentage matches
      `local-tmp/coverage-www-before.txt`.
- [ ] [AI] `npm run check:markdown-links` — acceptance: exits 0.
- [ ] [AI] [AC-2] Read every target in all three edited `project.json` files against the outputs rule — acceptance: no
      target that resolves to `cache: false` declares `outputs`, and every target that resolves to `cache: true` and
      writes an artifact declares the path it writes. The aggregates are the easy miss in either direction:
      `test:coverage` and `test:quick` are cached and write nothing themselves, because each artifact belongs to a
      sub-target that declares it. This is the check Phase 3's gate re-runs against the written rule, so a failure here
      is cheaper to fix than the same failure two phases later.
- [ ] [AI] Commit Phase 1 as four commits rather than one, each staged and pushed in turn, because the
      [thematic commits policy](../../../repo-governance/conventions/thematic-commits-policy.md) defines a theme by
      intent and this phase carries four — acceptance: each commit's diff contains nothing its message does not name,
      and `git status --short` is empty after the last.
  - [ ] [AI] [AC-5] `fix(wahidyankf-www): resolve nx through the workspace binary` — the bare `nx run` change alone. It
        is a defect fix, not a normalization, and it is the only Phase 1 change that alters which binary runs; it
        commits first so a bisect that lands on it names one suspect.
  - [ ] [AI] [AC-3] `refactor(workspace): declare shared behavior inputs once per project` — the three `namedInputs`
        declarations, every glob they replace, and the redundant `tests/e2e` input removal.
  - [ ] [AI] [AC-4] `refactor(workspace): resolve project commands through options.cwd` — the thirteen `badakmini-cli`
        `cwd` declarations with the two `mkdir` and one `BADAKMINI_BIN` rewrites, and `static-routes:validation`'s move
        to `{projectRoot}`.
  - [ ] [AI] [AC-2] `refactor(workspace): declare every target's cache state and outputs` — the `badakmini-cli`
        `cache`/`outputs` pair and the two inert `outputs` removals.

> **Pause Safety**: all three `project.json` files are written in one style, every target declares its cache state, and
> every gate that passed at baseline passes now at unchanged coverage figures. The workspace still has three projects
> and the browser suite still lives where it did. Safe to stop. Resume with `npm run test:quick`.

## Phase 2: Merge the E2E Project into the Application

The expand, migrate, verify, and contract steps of [migration-design.md](tech-docs/migration-design.md), landing in one
commit because the compatibility window is zero.

- [ ] [AI] [AC-6] Confirm the pre-merge commit is a reachable recovery source before deleting anything: run
      `git rev-parse HEAD > local-tmp/pre-merge-sha.txt && git ls-tree -r --name-only "$(cat local-tmp/pre-merge-sha.txt)" -- apps/wahidyankf-www-e2e/steps | wc -l`
      — acceptance: prints `8`, proving the eight step files are restorable from that commit.
- [ ] [AI] [AC-6] Record what the stricter compiler settings currently report, so the accepted loss is measured rather
      than assumed: run `npx nx run wahidyankf-www-e2e:typecheck` — acceptance: exits 0, and the result is written into
      `learnings.md` as the pre-merge state of the step files under `noUncheckedIndexedAccess`, `noUnusedLocals`, and
      `noUnusedParameters` all set to `true`.
- [ ] [AI] [AC-6] Create the destination with `mkdir -p apps/wahidyankf-www/tests/e2e/steps` — acceptance: the directory
      exists and `apps/wahidyankf-www/tests/` now holds `bdd`, `e2e`, and `integration`.
- [ ] [AI] [AC-6] Move the eight step files with
      `git mv apps/wahidyankf-www-e2e/steps/*.ts apps/wahidyankf-www/tests/e2e/steps/` — acceptance:
      `ls apps/wahidyankf-www/tests/e2e/steps | wc -l` prints `8` and `git status --short` shows eight `R` rename
      entries rather than delete-plus-add pairs.
- [ ] [AI] [AC-6] Move the baseline with
      `git mv apps/wahidyankf-www-e2e/e2e-skip-baseline.json apps/wahidyankf-www/tests/e2e/e2e-skip-baseline.json` —
      acceptance: the moved file still contains `{ "skippedScenarios": 34 }`.
- [ ] [AI] [AC-6] Move the Playwright configuration with
      `git mv apps/wahidyankf-www-e2e/playwright.config.ts apps/wahidyankf-www/playwright.config.ts` — acceptance: the
      file exists at the new path and `apps/wahidyankf-www-e2e/` no longer contains it.
- [ ] [AI] [AC-6] Edit `apps/wahidyankf-www/playwright.config.ts`: change the `defineBddConfig` `steps` value from
      `"./steps/**/*.ts"` to `"./tests/e2e/steps/**/*.ts"` — acceptance: `featuresRoot` and `webServer.cwd` are left at
      their existing `"../.."`-relative values, because the file's depth below the workspace root is unchanged.
- [ ] [AI] [AC-6] Edit `apps/wahidyankf-www/playwright.config.ts`: correct the comment sentence that reads "`test:e2e`
      declares `dependsOn` on `wahidyankf-www:build`" to name the intra-project `build` dependency, and the comment
      naming `specs:e2e:baseline` to drop the cross-project prefix — acceptance: no comment in the file names a project
      that no longer exists.
- [ ] [AI] [AC-6] Edit `apps/wahidyankf-www/package.json`: add `"@axe-core/playwright": "4.10.1"`,
      `"@playwright/test": "1.62.1"`, and `"playwright-bdd": "9.2.0"` to `devDependencies` — acceptance: the three pins
      are character-identical to those in `apps/wahidyankf-www-e2e/package.json` before deletion.
- [ ] [AI] [AC-6] Edit `apps/wahidyankf-www/tsconfig.json`: add `".features-gen"` to the `exclude` array beside
      `"node_modules"` — acceptance: the array holds both entries; without this the project's `include` of `**/*.ts`
      would compile generated test files.
- [ ] [AI] [AC-6] Edit `apps/wahidyankf-www/eslint.config.mjs`: add a second configuration object for
      `files: ["tests/e2e/steps/**/*.ts"]` carrying the same three `jsdoc` rules as the existing block, with
      `languageOptions.parser` set to `tsParser` and no `ecmaFeatures.jsx` — acceptance: a Playwright step file contains
      no JSX, so enabling the feature there would be configuration the code cannot exercise.
- [ ] [AI] [AC-1] [AC-6] Edit `apps/wahidyankf-www/project.json`: add an `install` target running
      `npx playwright install --with-deps chromium` with `"cache": false`, `"options": {"cwd": "{projectRoot}"}` —
      acceptance: `npx nx run wahidyankf-www:install` exits 0 and Chromium is available to the suite.
- [ ] [AI] [AC-1] [AC-6] Edit `apps/wahidyankf-www/project.json`: add a `test:e2e` target carrying the
      unconditional-`test.skip` guard and `npx bddgen && npx playwright test` from the deleted project's target, with
      `"cache": false`, `"options": {"cwd": "{projectRoot}"}`, `"dependsOn": ["build"]`, and
      `"inputs": ["default", "behaviorCorpus"]` — acceptance: `dependsOn` names the intra-project `build` rather than
      `wahidyankf-www:build`, matching how `badakmini-cli:test:e2e` names its own `build`.
- [ ] [AI] [AC-6] Edit the guard's search path in that same command: its trailing `.` becomes `tests/e2e` — acceptance:
      the guard scans one directory, as it did before the merge. Copied verbatim it would scan the whole application,
      because `.` is the working directory and that is now `apps/wahidyankf-www`: forty-three TypeScript files plus
      `.next/`, which the command's `--exclude-dir` list does not name. Prove the scoping by running
      `npx nx run wahidyankf-www:test:e2e` and confirming it reaches `bddgen` rather than failing in the guard, and by
      adding a throwaway `test.skip(1)` line to a file under `tests/e2e/steps/`, confirming the guard fails, then
      removing it — a guard that never fires is indistinguishable from a guard that scans nothing.
- [ ] [AI] [AC-1] [AC-6] Edit `apps/wahidyankf-www/project.json`: add a `specs:e2e:baseline` target carrying the deleted
      project's command verbatim, with `"cache": true`, `"options": {"cwd": "{projectRoot}"}`,
      `"outputs": ["{projectRoot}/.features-gen"]` carried over from Phase 1, and `inputs` naming only `default` and
      `behaviorCorpus` — acceptance: the three `{projectRoot}` paths the deleted project listed separately all sit
      inside `{projectRoot}` once the files move, so the built-in `default` input already covers them; listing them
      again is the same redundancy Phase 1 removes from `badakmini-cli`, and the Phase 2 gate's resolved-inputs read
      confirms all three paths still appear. The command's `require('./e2e-skip-baseline.json')` is repointed to
      `./tests/e2e/e2e-skip-baseline.json` to match the moved file.
- [ ] [AI] [AC-6] Edit `apps/wahidyankf-www/project.json` target `lint:commentary`: change the command to
      `eslint --config eslint.config.mjs src tests/e2e/steps` — acceptance: `npx nx run wahidyankf-www:lint:commentary`
      exits 0 and its output shows it read files under both directories.
- [ ] [AI] [AC-6] Delete the six remaining files of the retired project:
      `git rm apps/wahidyankf-www-e2e/project.json apps/wahidyankf-www-e2e/package.json apps/wahidyankf-www-e2e/tsconfig.json apps/wahidyankf-www-e2e/eslint.config.mjs apps/wahidyankf-www-e2e/.gitignore apps/wahidyankf-www-e2e/README.md`
      — acceptance: the root `.gitignore` already covers `.features-gen/`, `playwright-report/`, and `test-results/`, so
      deleting the project-local ignore file leaves nothing newly visible to `git status --short`.
- [ ] [AI] [AC-6] Remove the now-empty directory with `rm -rf apps/wahidyankf-www-e2e` and run `npm install` —
      acceptance: `npx nx show projects` returns exactly `wahidyankf-www` and `badakmini-cli`, and `package-lock.json`
      no longer contains a `wahidyankf-www-e2e` workspace entry.
- [ ] [AI] Edit `package.json`: in `test:scheduled`, replace `wahidyankf-www-e2e:specs:e2e:baseline` with
      `wahidyankf-www:specs:e2e:baseline` and `wahidyankf-www-e2e:test:e2e` with `wahidyankf-www:test:e2e` — acceptance:
      `grep -c 'wahidyankf-www-e2e' package.json` prints `0`.
- [ ] [AI] Edit `.github/workflows/full-bdd.yml`: change the browser install step to `npx nx run wahidyankf-www:install`
      — acceptance: `npm run check:workflows` exits 0.
- [ ] [AI] Edit `apps/README.md`: remove the `wahidyankf-www-e2e` bullet from `## Current Applications` — acceptance:
      two bullets remain and neither links a directory that no longer exists.
- [ ] [AI] Edit `apps/wahidyankf-www/README.md`: replace the "sibling `wahidyankf-www-e2e` project" adapter bullet with
      one naming `tests/e2e/`, and fold in the retired README's description of the browser layer — acceptance: the four
      adapter bullets still describe four layers and none names a deleted path.
- [ ] [AI] [AC-6] Edit `specs/apps/wahidyankf-www/README.md`: rewrite the two-projects sentence to name one project,
      change the Process E2E adapter row's path to `apps/wahidyankf-www/tests/e2e/steps/`, repoint the skip-baseline
      sentence at the merged project, and change the two verification-command rows to the `wahidyankf-www:` prefix —
      acceptance: `grep -c 'wahidyankf-www-e2e' specs/apps/wahidyankf-www/README.md` prints `0`.
- [ ] [AI] [AC-6] Edit `specs/apps/wahidyankf-www/architecture.md`: remove the `apps/wahidyankf-www-e2e` project-table
      row and the shared-model sentence above it, redraw the Container View so the Playwright adapter is a test-time
      process belonging to `wahidyankf-www` rather than a separate container, and rewrite the paragraph below it to keep
      the toolchain-difference fact while dropping the separate-project inference — acceptance: the diagram stays a
      fenced `text` ASCII block, and `grep -c 'wahidyankf-www-e2e' specs/apps/wahidyankf-www/architecture.md` prints
      `0`.
- [ ] [AI] Edit `specs/apps/wahidyankf-www/behavior/README.md`: repoint the sentence naming
      `apps/wahidyankf-www-e2e/README.md` at `apps/wahidyankf-www/README.md` — acceptance:
      `npm run check:markdown-links` resolves the link.
- [ ] [AI] Edit `repo-governance/development/workspace-commands.md`: change the three `wahidyankf-www-e2e` narrower-run
      lines to `wahidyankf-www:install`, `wahidyankf-www:test:e2e`, and `wahidyankf-www:specs:e2e:baseline`, and delete
      the paragraph explaining why `wahidyankf-www-e2e` owns no `test:quick` — acceptance: the paragraph's subject no
      longer exists, `grep -c 'wahidyankf-www-e2e' repo-governance/development/workspace-commands.md` prints `0`, and
      `npm run check:governance` keeps the document under 750 words.
- [ ] [AI] Run `npm run format` — acceptance: exits 0 and `npm run format:check` afterwards also exits 0.

### Phase 2 Gate

> Every check below passes before Phase 3 begins. A failure is fixed inside Phase 2.

- [ ] [AI] [AC-6] `npx nx show projects` — acceptance: returns exactly two entries, `wahidyankf-www` and
      `badakmini-cli`.
- [ ] [AI] [AC-1] `npx nx show project wahidyankf-www --json` — acceptance: the target list contains all ten contract
      targets: `typecheck`, `lint`, `test:unit`, `test:integration`, `test:e2e`, `test:coverage`, `test:coverage:unit`,
      `test:coverage:integration`, `test:coverage:behavior`, and `test:quick`.
- [ ] [AI] [AC-1] `npx nx show project badakmini-cli --json` — acceptance: the same ten targets are present, so both
      projects expose one identical contract.
- [ ] [AI] [AC-3] `npx nx show project wahidyankf-www --json` for `specs:e2e:baseline` and `test:e2e` — acceptance: each
      target's resolved `inputs` still covers the moved baseline file, the moved step directory, and
      `playwright.config.ts`, proving the built-in `default` reaches them now that they sit inside `{projectRoot}` and
      that dropping the three explicit entries changed no cache key.
- [ ] [AI] [AC-6] `npx nx run wahidyankf-www:specs:e2e:baseline` — acceptance: exits 0, which asserts the generated
      `test.fixme` count is still exactly the 34 recorded in the moved baseline file. A count above 34 means a step file
      stopped binding at its new path.
- [ ] [AI] [AC-6] `npx nx run wahidyankf-www:test:e2e` — acceptance: exits 0, having built the application, started
      `next start`, and driven Chromium over the eight bound feature files.
- [ ] [AI] `npx nx run wahidyankf-www:typecheck` — acceptance: exits 0, run after the preceding `test:e2e` has populated
      `apps/wahidyankf-www/.features-gen/`, so the new `tsconfig.json` exclude is proved against a populated directory
      rather than an absent one.
- [ ] [AI] `npm run test:quick` — acceptance: exits 0 for both projects.
- [ ] [AI] `npm run test:integration` — acceptance: exits 0 for both projects.
- [ ] [AI] `npx nx run badakmini-cli:test:e2e` — acceptance: exits 0.
- [ ] [AI] `npx nx affected -t test:quick --base=origin/main --head=HEAD` — acceptance: exits 0 and selects
      `wahidyankf-www`, confirming the merged project is still reachable by the pre-push calculation after the graph
      change.
- [ ] [AI] `git grep -n 'wahidyankf-www-e2e' -- ':!plans/done'` — acceptance: prints nothing outside `plans/done/`,
      whose archived plan is history and is deliberately not edited.
- [ ] [AI] `npm run check:markdown-links` and `npm run check:governance` — acceptance: both exit 0.
- [ ] [AI] `npm run check:workflows` — acceptance: exits 0 over the edited scheduled workflow.
- [ ] [AI] Commit with message `refactor(wahidyankf-www): co-locate the browser E2E adapter` and push to `main` —
      acceptance: `git status --short` is empty.

> **Pause Safety**: the workspace holds two projects, each exposing the same ten targets, and the browser suite runs
> from inside the application it tests at an unchanged skip baseline of 34. The Gherkin corpus is byte-identical to the
> baseline and no application code changed. Safe to stop. Resume with `npx nx run wahidyankf-www:test:e2e`.

## Phase 3: Write the Target Contract Rule

Promotes the shape both projects now share out of this plan and into governance, so a third project has a rule to read
rather than two files to compare.

- [ ] [AI] [AC-7] Edit `repo-governance/development/testing-policy.md`: in the Quick Tests section, replace the partial
      target list with the full contract naming `typecheck`, `lint`, `test:unit`, `test:integration`, `test:e2e`,
      `test:coverage:unit`, `test:coverage:integration`, `test:coverage:behavior`, `test:coverage`, and `test:quick`,
      and state which are eligibility-dependent — acceptance: the existing rule that a library defines no
      `test:integration` when it owns no local boundary, and that a library never owns `test:e2e`, is preserved rather
      than contradicted.
- [ ] [AI] [AC-7] Edit `repo-governance/development/testing-policy.md`: state that every target declares `cache`
      explicitly wherever the root `nx.json` `targetDefaults` does not reach it, so no target resolves to an undeclared
      state — acceptance: the sentence names `targetDefaults` as the other source, so a reader does not add a redundant
      declaration to the six targets it already covers.
- [ ] [AI] [AC-7] Edit `repo-governance/development/testing-policy.md`: state that a single-command target declares
      `options.cwd` rather than encoding its own project path in the command, and that a cached target that writes an
      artifact declares `outputs` while an uncached one declares none — acceptance: the `outputs` sentence gives the
      reason, that Nx replays a cache hit and restores nothing when a cached target declares no output path.
- [ ] [AI] [AC-7] Edit `repo-governance/development/testing-policy.md`: state that the shape rules above bind every
      target a project declares, not only the ten in the contract — acceptance: the sentence is stated without an
      exception, which is what Phase 1's `generate:cv-pdf` and `static-routes:validation` edits make true; a rule that
      shipped with a carve-out would teach the next reader to add a second.
- [ ] [AI] [AC-7] Edit `repo-governance/development/testing-policy.md`: state that `options.commands` expresses the
      ordered gate itself and `dependsOn` expresses a prerequisite that must precede the whole gate — acceptance: the
      distinction is stated as a rule with both halves, which is what makes `wahidyankf-www:test:quick`'s `dependsOn` on
      `static-routes:validation` a documented choice rather than an unexplained second ordering mechanism.
- [ ] [AI] [AC-7] Edit `repo-governance/development/testing-policy.md`: state that a shared input path is declared once
      as a project-level `namedInputs` entry and referenced by name — acceptance: the sentence explains that the
      alternative, repeating the glob, is what let the three files drift into different input sets.
- [ ] [AI] [AC-7] Run `npm run check:governance` — acceptance: exits 0, and
      `wc -w repo-governance/development/testing-policy.md` reports under 750. If it reports 700 or more, the document
      has entered the headroom band and the
      [document word limit policy](../../../repo-governance/conventions/document-word-limit-policy.md) governs how it is
      fixed; move the added detail into `repo-governance/development/testing-policy/` rather than dropping any of it.
- [ ] [AI] [AC-7] Edit `repo-governance/development/nx-workspace-policy.md`: add one sentence in Required Approach
      linking to the testing policy's target contract — acceptance: a reader arriving to add a target is sent to the
      shape rule, and `npm run check:governance` still exits 0 for this document too.
- [ ] [AI] [AC-7] Edit `docs/how-to/run-nx-workspace.md`: in the paragraph that already names `project.json` and
      `nx.json`, add the pointer to the target contract — acceptance: the how-to links the policy rather than restating
      it, so the two cannot drift.
- [ ] [AI] Run `npm run format` — acceptance: exits 0 and `npm run format:check` afterwards also exits 0.

### Phase 3 Gate

> Every check below passes before Phase 4 begins. A failure is fixed inside Phase 3.

- [ ] [AI] [AC-7] `npm run check:governance` — acceptance: exits 0, holding every edited governance document under the
      750-word limit.
- [ ] [AI] `npm run check:markdown-links` — acceptance: exits 0 over the three new cross-document links.
- [ ] [AI] [AC-7] Read `repo-governance/development/testing-policy.md` against `apps/badakmini-cli/project.json` and
      `apps/wahidyankf-www/project.json` — acceptance: every rule the document now states is true of both files, and no
      rule it states is contradicted by either. A rule that the delivered files violate is a defect in the rule or the
      files, and is fixed inside this phase.
- [ ] [AI] `npm run check:rule-change` — acceptance: reports the staged governance paths and names
      [Rules Propagation](../../../repo-governance/workflows/rules/rules-propagation.md); the check reports without
      blocking, so the acceptance is that the notice appears and the named workflow is then run.
- [ ] [AI] Run the [Rules Propagation](../../../repo-governance/workflows/rules/rules-propagation.md) workflow for the
      edited policies — acceptance: it reports no unresolved contradiction between the new contract and the
      [Nx workspace policy](../../../repo-governance/development/nx-workspace-policy.md), the
      [BDD policy](../../../repo-governance/development/behavior-driven-development-policy.md), or
      [workspace commands](../../../repo-governance/development/workspace-commands.md).
- [ ] [AI] `npm run test:quick` — acceptance: exits 0 for both projects, confirming a documentation-only phase changed
      no executable behavior.
- [ ] [AI] Commit with message `docs(testing): state the project target contract` and push to `main` — acceptance:
      `git status --short` is empty.

> **Pause Safety**: the contract both projects satisfy is written in `testing-policy.md`, and every governance check
> passes. No executable configuration changed in this phase. Safe to stop. Resume with `npm run check:governance`.

## Phase 4: Knowledge Capture

Triages every `learnings.md` entry to a terminal state, then archives. Archival is blocked until each entry is routed or
discarded with a reason.

- [ ] [AI] Read every entry in `plans/in-progress/project-json-consistency/learnings.md` and route each to exactly one
      durable home per the
      [knowledge capture rules](../../../repo-governance/conventions/plans-organization-policy/knowledge-capture.md) —
      acceptance: each entry carries a one-line disposition naming its destination, or a one-line reason for discarding
      it. If the file holds no entry, record the explicit escape `No generalizable learnings — <reason>` instead of
      leaving it blank.
- [ ] [AI] Check every entry for secrets and for repository relevance before routing — acceptance: no entry names a
      credential value, and every surviving entry states a lesson that generalizes beyond this plan rather than
      describing one incident.
- [ ] [AI] Reconcile every acceptance criterion `[AC-1]` through `[AC-7]` in [prd.md](prd.md) against the delivered
      system — acceptance: each criterion's proof command in
      [specification-changes.md](tech-docs/specification-changes.md) is re-run and passes, and any criterion that cannot
      be reconciled is recorded rather than ticked.
- [ ] [AI] Reconcile [tech-docs/file-impact.md](tech-docs/file-impact.md) against `git diff --stat` for the three
      delivery commits — acceptance: every path the document labels appears in the diff with the labelled operation, and
      every path in the diff appears in the document.
- [ ] [AI] Give a dated, evidence-backed disposition to the one conditional item in this plan — the Phase 1 checklist
      item that restores the `tests/e2e` input if the resolved-inputs comparison shows a difference — acceptance: it
      records either the restoration and its evidence, or `Not triggered` with the comparison output that shows no
      difference.
- [ ] [AI] Run the [plan-quality-gate](../../../repo-governance/workflows/plan-quality-gate.md) workflow at strict level
      — acceptance: `plan-checker` reports no findings.
- [ ] [AI+HUMAN] Confirm with the owner that the plan is complete before it is archived — acceptance: the owner agrees
      the delivered scope matches what they asked for; archival is a one-way move and the four scope decisions in
      [brd.md](brd.md) narrowed twice during planning.
- [ ] [AI] Move the folder to `plans/done/YYYY-MM-DD__project-json-consistency/` using the date the final commit landed
      — acceptance: the destination does not already exist; if it does, stop rather than merging, overwriting, or
      inventing a suffix.
- [ ] [AI] Update `plans/in-progress/README.md` and `plans/done/README.md` in the same change — acceptance: the source
      index no longer lists the plan, the destination index does, and `npm run check:markdown-links` resolves every
      archived internal link.

### Phase 4 Gate

> Every check below passes before the plan is considered finished.

- [ ] [AI] `npm run check:markdown-links` — acceptance: exits 0 after the folder move, with no dead link left by the
      changed depth of the archived documents.
- [ ] [AI] `git grep -n 'plans/in-progress/project-json-consistency'` — acceptance: prints nothing, confirming no
      document still points at the pre-archival path.
- [ ] [AI] `npm run test:quick` and `npm run test:integration` — acceptance: both exit 0 for both projects.
- [ ] [AI] `npm run test:scheduled` — acceptance: exits 0, running the full ordered verification including
      `wahidyankf-www:specs:e2e:baseline` and `wahidyankf-www:test:e2e` at their merged names. This is the one command
      that exercises every renamed invocation in `package.json` end to end.
- [ ] [AI] Commit the archival move with a message naming the plan and push to `main` — acceptance: `git status --short`
      is empty and the source path is absent from the working tree.

> **Pause Safety**: the plan is archived, every learning has a terminal disposition, and the full scheduled verification
> passes at the merged target names. Nothing is left in progress. Resume is not applicable; the work is complete.
