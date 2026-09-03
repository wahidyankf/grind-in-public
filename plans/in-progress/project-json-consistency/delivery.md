# Delivery

## Execution Record

<!--
One dated line per phase completion, gate pass or failure, retry that proved
something, and plan change. Written as the event happens, never reconstructed
at archival. Each phase-completion line carries the commit SHA that phase
pushed, because Phase 4 reconciles file-impact.md over a SHA range read from
here.
-->

- 2026-09-03 — Phase 3 complete. The contract went into `testing-policy.md` (550 words, under the limit and outside the
  700-word headroom band) with the shape rules in a new `testing-policy/target-shape.md` companion, because stating all
  of it in one document would have crossed the limit. Gate passed: every rule the documents now state was read against
  both delivered `project.json` files and holds — cache declared on all 34 targets, no inert `outputs`, every cached
  artifact-writer naming its path, no command encoding its own project path, and shared inputs named once.
- 2026-09-03 — Phase 3 Rules Propagation found one real drift risk: `badakmini-cli-policy.md` restated eight of the ten
  contract targets. A duplicated list is what later disagrees with its source, so it now points at the contract.
  `workspace-commands.md` also gained the `badakmini-cli:test:quick` narrower run it had omitted. The BDD role matrix
  and the Nx workspace policy needed no change.
- 2026-09-03 — Phase 2 complete, delivered as `a953b9c`. Gate passed: exactly two projects, both exposing the same ten
  contract targets; the merged browser suite passes 36 and skips exactly 34, so every step file still binds at its new
  path; the skip guard proved green at rest and firing on an injected line; `typecheck` clean with `.features-gen`
  populated; cache and outputs inspections hold across both projects; `git grep` for the retired name is clean outside
  `plans/`, with a positive control confirming the pattern matches; governance, links, workflows, and harness parity all
  green.
- 2026-09-03 — Phase 2: `nx affected` selected `wahidyankf-www` after the commit and before the push, confirming the
  merged project is still reachable by the pre-push calculation. Run in the pre-commit position it had originally, it
  would have compared `origin/main` with an identical `HEAD` and reported nothing.
- 2026-09-03 — Phase 2 triggered Rules Propagation via `tooling.md` and `workspace-commands.md`. Ran it: the BDD role
  matrix and the testing policy both describe a dedicated E2E project as permitted rather than required, so both stay
  true with none in the workspace. No contradiction, no further edit.
- 2026-09-03 — Phase 1 complete, delivered as four thematic commits: `a093764` the bare-`nx run` defect fix, `639f124`
  shared inputs hoisted to `namedInputs`, `36e14e0` the `options.cwd` migration, `d8c91fa` cache and outputs. Gate
  passed: 38 targets across three projects all declare a cache state, no uncached target declares `outputs`, every
  cached artifact-writer names its path, all three cache probes missed on a content change and hit again on restore, and
  coverage matched baseline exactly — badakmini 99.2% unit and 99.3% integration, `wahidyankf-www` 99.57% statements and
  100% lines.
- 2026-09-03 — Phase 1 plan change: the `tests/e2e` input removal was reverted. The cache probe passed, but
  `TestE2EBindingInputRegression` in `apps/badakmini-cli/tests/bdd/adapter_parity_test.go` reads `project.json` and
  fails unless `test:coverage:behavior` declares that exact string. The input is test-enforced rather than redundant,
  and the plan's non-goal forbids editing Badak Mini's Go code. Recorded in `learnings.md`.
- 2026-09-03 — Phase 0 gate passed: `test:quick` for both projects, `test:integration`, `badakmini-cli:test:e2e`,
  `format:check`, and `check:markdown-links` all exit 0.
- 2026-09-03 — Phase 0 complete, pushed as `d57fc82`. The plan folder was already in history at `0a208aa` with a
  revision at `b3d74ea`, so this commit carried only what remained: the quality-gate fixes, the `plans/in-progress`
  index entry, and Phase 0's ticks, notes, and first learning. Baseline recorded — `wahidyankf-www` unit coverage 99.57%
  statements and 100% lines, `badakmini-cli` 99.2% statements, skip baseline 34, three projects resolved.
- 2026-09-03 — Plan quality gate, strict, four cycles, partial. 36 findings fixed; six MEDIUM accepted open by the owner
  and listed in [README.md](README.md#open-findings), which the executor reads before the phase each one touches.

## Executor Tags

`[AI]` an agent can fully perform it. `[HUMAN]` only the owner. `[AI+HUMAN]` an agent prepares and the owner approves or
performs the irreversible step.

## Phase 0: Baseline

Records a clean starting state so a later failure is attributable to the work rather than to the machine. It commits the
plan and nothing else.

- [x] [AI] Run `npm install` from the repository root — acceptance: exits 0 and `git status --short` reports no change
      to `package-lock.json`.
- [x] [AI] Run `mkdir -p local-tmp && npx nx show projects > local-tmp/projects-before.txt` — acceptance: the file
      contains exactly `wahidyankf-www-e2e`, `wahidyankf-www`, and `badakmini-cli`.
- [x] [AI] Capture the resolved target configuration of each project:
      `npx nx show project badakmini-cli --json > local-tmp/targets-badakmini-before.json`, then the same for
      `wahidyankf-www` and `wahidyankf-www-e2e` into `local-tmp/targets-www-before.json` and
      `local-tmp/targets-www-e2e-before.json` — acceptance: three files exist and each parses, verified with
      `node -e 'for (const f of ["badakmini","www","www-e2e"]) JSON.parse(require("fs").readFileSync("local-tmp/targets-"+f+"-before.json"))'`
      exiting 0.
- [x] [AI] Record the pre-change unit coverage figure by running
      `npx nx run wahidyankf-www:test:coverage:unit 2>&1 | tail -60 > local-tmp/coverage-www-before.txt` — acceptance:
      the file contains a `All files` row with a line percentage, and the command exited 0. - Recorded
      `All files | 99.57 | 92.3 | 100 | 100`. The planned `tail -20` truncated past that row, because Nx appends its run
      summary after the command's output; widened to `tail -60`.
- [x] [AI] Record the same for the Go side with
      `npx nx run badakmini-cli:test:coverage:unit 2>&1 | tail -15 > local-tmp/coverage-badakmini-before.txt` —
      acceptance: the file contains the `unit statement coverage:` line and the command exited 0. - Recorded
      `unit statement coverage: 99.2%`. The planned `tail -5` truncated past the line; widened to `tail -15`.
- [x] [AI] Record the current skipped-scenario count with `npx nx run wahidyankf-www-e2e:specs:e2e:baseline` —
      acceptance: exits 0, which is the assertion that the generated `test.fixme` count equals the 34 recorded in
      `apps/wahidyankf-www-e2e/e2e-skip-baseline.json`. - Exited 0, so the generated count still equals 34.
- [x] [AI] Update `plans/in-progress/README.md` so both of its lists name this plan: an entry under `## Active Plans`
      and an entry under `## Directory Map` — acceptance: neither list still reads `None right now.` or
      `No plan folders right now.`, both link `project-json-consistency/README.md`, and `npm run check:markdown-links`
      resolves them. - Already satisfied before execution began: both lists were written during the quality gate.
      Verified here, with `check:markdown-links` exiting 0.
- [x] [AI] Commit the plan folder and that index update with message `docs(plans): start project-json-consistency`, then
      push to `main` — acceptance: `git status --short` is empty and `git log origin/main -1 --oneline` names the
      commit. The plan folder, and possibly the index entry with it, may already be committed when execution begins,
      because planning commits them: run `git log --oneline -- plans/in-progress/project-json-consistency` first. If the
      folder is already in history, commit only what is still uncommitted, and tick this item with an implementation
      note naming the existing SHA rather than re-committing what is already there.
- [x] [AI] Record the Phase 0 commit SHA in the Execution Record above as one dated line, taken from
      `git rev-parse --short HEAD` after the push — or the SHA named in the implementation note above, if the plan
      folder was already committed — acceptance: the Execution Record names one SHA for Phase 0. Every later phase does
      the same at its own commit, and Phase 4's file-impact reconciliation reads its range ends from those recorded
      SHAs. It cannot read them from commit messages: `docs(plans): start project-json-consistency` already names a
      commit in this repository's history, so a `--grep` anchor is ambiguous and can select the wrong end.

### Phase 0 Gate

> Every check below passes before Phase 1 begins. A failure is fixed inside Phase 0.

- [x] [AI] `npm run test:quick` — acceptance: exits 0 with both `badakmini-cli` and `wahidyankf-www` reported as
      successful tasks.
- [x] [AI] `npm run test:integration` — acceptance: exits 0 for both projects.
- [x] [AI] `npx nx run badakmini-cli:test:e2e` — acceptance: exits 0.
- [x] [AI] `npm run format:check` — acceptance: exits 0 over the new plan documents.
- [x] [AI] `npm run check:markdown-links` — acceptance: exits 0. Run
      `git add -N plans/in-progress/project-json-consistency` first, because the check reads Git-tracked files and a new
      document is invisible to it otherwise.

> **Pause Safety**: the plan is committed and pushed, and every gate is green at the unmodified baseline recorded in
> `local-tmp/`. No workspace configuration has changed. Safe to stop. Resume with `npm run test:quick`.

## Phase 1: Configuration Normalization

Brings all three `project.json` files into one style. `wahidyankf-www-e2e` is normalized here even though Phase 2
deletes it, so this phase ends with the property true across the whole workspace and its gate can assert it
workspace-wide.

- [x] [AI] [AC-3] Edit `apps/badakmini-cli/project.json`: add a top-level
      `"namedInputs": {"behaviorCorpus": ["{workspaceRoot}/specs/apps/badakmini-cli/behavior/**/*.feature"]}` beside
      `"targets"` — acceptance: `npx nx show project badakmini-cli --json` exits 0, proving Nx still parses the file.
- [x] [AI] [AC-3] Edit `apps/badakmini-cli/project.json`: in every target whose `inputs` array contains the literal
      `{workspaceRoot}/specs/apps/badakmini-cli/behavior/**/*.feature`, replace that string with `"behaviorCorpus"` —
      acceptance: `grep -c 'specs/apps/badakmini-cli/behavior' apps/badakmini-cli/project.json` prints `1`, the single
      occurrence being the `namedInputs` declaration.
- [x] [AI] [AC-3] Edit `apps/badakmini-cli/project.json`: remove the `{workspaceRoot}/apps/badakmini-cli/tests/e2e/**/*`
      input from `test:coverage:behavior`, `test:coverage`, and `test:quick` — acceptance: the path lies inside
      `{projectRoot}` and is therefore already covered by the built-in `default` input, and a cache probe proves that
      coverage rather than assuming it. Run `npx nx run badakmini-cli:test:quick` twice and confirm the second run
      prints `Nx read the output from the cache`; then
      `printf '\n<!-- hash probe -->\n' >> apps/badakmini-cli/tests/e2e/README.md`, run it again, and confirm the cache
      line is absent; then `git checkout -- apps/badakmini-cli/tests/e2e/README.md` and run it once more, confirming the
      cache line returns and `git diff --stat -- apps/badakmini-cli/tests/e2e/README.md` prints nothing.
      `git checkout --` restores the path from the **index**, not from `HEAD`, so a probe using it is safe only on a
      path carrying no unstaged edit of its own. That precondition holds here without any staging step, and for a stated
      reason rather than by luck: this plan never edits `apps/badakmini-cli/tests/e2e/README.md`, so its index entry
      equals its `HEAD` entry throughout. Every probe in this plan carries the same precondition, and Phase 2's gate
      stages its whole phase before probing because there the precondition does not hold on its own. The removal is
      correct only if that third run missed the cache: a hit would mean `default` does not reach `tests/e2e/` and the
      explicit input was carrying it. **Conditional**: if the third run hits the cache, restore the input in all three
      targets and record the finding in `learnings.md` instead of forcing the removal. - **Conditional fired, on
      evidence this item did not anticipate.** The cache probe passed exactly as designed: the third run missed the
      cache, so `default` does reach `tests/e2e/`. The removal was reverted anyway, because
      `apps/badakmini-cli/tests/bdd/adapter_parity_test.go`'s `TestE2EBindingInputRegression` reads `project.json` and
      fails unless `test:coverage:behavior` declares that exact string. The input is test-enforced rather than
      redundant. Restored in all three targets, with `behaviorCorpus` still replacing the corpus glob beside it.
      Recorded in `learnings.md`; Phase 4's disposition for this conditional is "triggered", not "Not triggered".
- [x] [AI] [AC-2] Edit `apps/badakmini-cli/project.json`: add `"cache": true` to `test:coverage:unit` together with
      `"outputs": ["{workspaceRoot}/local-tmp/badakmini-unit.out"]` — acceptance:
      `npx nx show project badakmini-cli --json` reports `cache` as `true` for that target, matching its
      `wahidyankf-www` counterpart. The `outputs` declaration is not optional here: an undeclared `cache` means
      uncached, so this turns caching on for a gate that runs fresh today, and the profile it writes sits outside
      `{projectRoot}` where the built-in `default` output detection would not find it. Prove the pair together by
      running the target twice — `npx nx run badakmini-cli:test:coverage:unit`, then
      `rm -f local-tmp/badakmini-unit.out` and run it again — acceptance: the second run reports a cache hit and
      `ls local-tmp/badakmini-unit.out` succeeds afterwards, which is what proves the cached target restores its
      artifact rather than reporting a success that produced nothing.
- [x] [AI] [AC-4] Edit `apps/badakmini-cli/project.json`: add `"options": {"cwd": "{projectRoot}"}` to `build`,
      `governance`, `markdown-links`, `capability-parity`, `rule-change`, `lint`, `typecheck`, `test:unit`,
      `test:integration`, `test:coverage:behavior` — acceptance: each of those ten targets carries the `options` object
      and `npx nx show project badakmini-cli --json` still parses.
- [x] [AI] [AC-4] Edit the same ten targets in `apps/badakmini-cli/project.json` to strip the `-C apps/badakmini-cli`
      argument from their `go` invocations — acceptance: `npx nx run badakmini-cli:typecheck` and
      `npx nx run badakmini-cli:lint` both exit 0, proving the commands resolve from the project directory.
- [x] [AI] [AC-4] Edit `apps/badakmini-cli/project.json` target `test:coverage:unit`: add
      `"options": {"cwd": "{projectRoot}"}`, strip `-C apps/badakmini-cli`, and change the leading `mkdir -p local-tmp`
      to `mkdir -p ../../local-tmp` — acceptance: `npx nx run badakmini-cli:test:coverage:unit` exits 0, prints the
      `unit statement coverage:` line, and `ls local-tmp/badakmini-unit.out` succeeds from the repository root while
      `ls apps/badakmini-cli/local-tmp` fails.
- [x] [AI] [AC-4] Apply the identical three edits to `apps/badakmini-cli/project.json` target
      `test:coverage:integration` — acceptance: `npx nx run badakmini-cli:test:coverage:integration` exits 0, prints the
      `integration statement coverage:` line, and `ls local-tmp/badakmini-integration.out` succeeds from the repository
      root.
- [x] [AI] [AC-4] Edit `apps/badakmini-cli/project.json` target `test:e2e`: add `"options": {"cwd": "{projectRoot}"}`,
      strip `-C apps/badakmini-cli`, and change `BADAKMINI_BIN="$PWD/apps/badakmini-cli/dist/badak-mini"` to
      `BADAKMINI_BIN="$PWD/dist/badak-mini"` — acceptance: `npx nx run badakmini-cli:test:e2e` exits 0, which requires
      the built binary to be found at the rewritten path.
- [x] [AI] [AC-3] Edit `apps/wahidyankf-www/project.json`: add a top-level `"namedInputs"` declaring `behaviorCorpus` as
      `["{workspaceRoot}/specs/apps/wahidyankf-www/behavior/**/*.feature"]` and `workspaceScripts` as
      `["{workspaceRoot}/scripts/next-with-port.mjs"]` — acceptance: `npx nx show project wahidyankf-www --json`
      exits 0.
- [x] [AI] [AC-3] Edit `apps/wahidyankf-www/project.json`: in `test:unit`, `test:coverage:unit`,
      `test:coverage:behavior`, `test:coverage`, and `test:quick`, replace the two literal workspace paths with
      `"behaviorCorpus"` and `"workspaceScripts"` — acceptance:
      `grep -c 'specs/apps/wahidyankf-www/behavior' apps/wahidyankf-www/project.json` prints `1` and
      `grep -c 'scripts/next-with-port.mjs' apps/wahidyankf-www/project.json` prints `3`, the three being the
      `namedInputs` declaration and the two `dev` and `start` command strings that genuinely invoke the script.
- [x] [AI] [AC-2] Edit `apps/wahidyankf-www/project.json`: delete the `"outputs": ["{projectRoot}/coverage"]` line from
      `test:coverage:integration` — acceptance: that target is `cache: false`, so the declaration was inert;
      `npx nx run wahidyankf-www:test:coverage:integration` exits 0 and still writes `apps/wahidyankf-www/coverage`.
- [x] [AI] [AC-5] Before editing, run the Phase 1 gate's bare-`nx run` grep once against the unedited file:
      `grep -nE '"command": *"([^"]*[^-] )?nx run' apps/wahidyankf-www/project.json` — acceptance: it prints exactly one
      line, the `static-routes:validation` command (line 105 today), and exits 0. This run is what proves the gate's
      grep can see the defect; run after the edit alone it is silent whether the pattern works or not. Record the
      printed line in `learnings.md` if it does not match, because a silent run here means the gate pattern is wrong
      rather than the file being clean.
- [x] [AI] [AC-5] Edit `apps/wahidyankf-www/project.json` target `static-routes:validation`: change the command's
      leading `nx run wahidyankf-www:build --skip-nx-cache` to `npm exec nx -- run wahidyankf-www:build --skip-nx-cache`
      — acceptance: `npx nx run wahidyankf-www:static-routes:validation` exits 0 and prints
      `Verified static build output for /, /cv, /personal-projects, /robots.txt, /sitemap.xml.`
- [x] [AI] [AC-4] Edit the same target: change `"options": {"cwd": "{workspaceRoot}"}` to `{"cwd": "{projectRoot}"}` and
      shorten `node apps/wahidyankf-www/scripts/validate-static-routes.mjs` to `node scripts/validate-static-routes.mjs`
      — acceptance: the target prints the same `Verified static build output` line, and
      `grep -n 'apps/wahidyankf-www' apps/wahidyankf-www/project.json` returns no line inside a `"command"` string. Its
      `dependsOn` placement in `test:quick` is untouched; only the working directory moves.
- [x] [AI] [AC-2] Edit `apps/wahidyankf-www/project.json`: delete the
      `"outputs": ["{projectRoot}/public/wahidyankf-kresna-fridayoka-cv.pdf"]` line from `generate:cv-pdf` — acceptance:
      that target is `cache: false`, so the declaration is inert for exactly the reason `test:coverage:integration`'s
      is; `npx nx run wahidyankf-www:generate:cv-pdf` exits 0 and still writes the PDF to `apps/wahidyankf-www/public/`.
      Removing one and leaving the other would make the Phase 3 rule false of the file it was derived from.
- [x] [AI] [AC-3] Edit `apps/wahidyankf-www-e2e/project.json`: add a top-level `"namedInputs"` declaring
      `behaviorCorpus` as `["{workspaceRoot}/specs/apps/wahidyankf-www/behavior/**/*.feature"]`, and replace all five
      literal occurrences in `install`, `typecheck`, `lint`, `test:e2e`, and `specs:e2e:baseline` with
      `"behaviorCorpus"` — acceptance:
      `grep -c 'specs/apps/wahidyankf-www/behavior' apps/wahidyankf-www-e2e/project.json` prints `1`.
- [x] [AI] [AC-2] Edit `apps/wahidyankf-www-e2e/project.json` target `specs:e2e:baseline`: add
      `"outputs": ["{projectRoot}/.features-gen"]` — acceptance: the target is `cache: true` and its command runs
      `bddgen`, which writes that directory, so without the declaration a cache hit replays the baseline comparison and
      restores nothing. This is the same pairing the `badakmini-cli` coverage target gets above, and it is applied here
      even though Phase 2 deletes the file, because Phase 1's gate asserts the property across all three and Phase 2
      carries the declaration into the merged target.
- [x] [AI] Run `npm run format` — acceptance: exits 0, and `npm run format:check` afterwards also exits 0 over the three
      edited files.

### Phase 1 Gate

> Every check below passes before Phase 2 begins. A failure is fixed inside Phase 1.

- [x] [AI] [AC-2] Run this check for each of `badakmini-cli`, `wahidyankf-www`, and `wahidyankf-www-e2e`, substituting
      the project name:
      `npx nx show project badakmini-cli --json | node -e 'let s="";process.stdin.on("data",d=>s+=d).on("end",()=>{const t=JSON.parse(s).targets;const bad=Object.entries(t).filter(([,v])=>v.cache===undefined).map(([k])=>k);console.log(bad.length?"UNDECLARED: "+bad.join(", "):"all "+Object.keys(t).length+" targets declare cache");process.exit(bad.length?1:0)})'`
      — acceptance: each of the three runs exits 0 and prints the `all N targets declare cache` form, naming a non-zero
      N so the check is proved to have read a populated target set.
- [x] [AI] Capture the resolved configuration again into `local-tmp/targets-badakmini-after.json`,
      `local-tmp/targets-www-after.json`, and `local-tmp/targets-www-e2e-after.json` using the Phase 0 commands, then
      `diff` each against its `-before.json` — acceptance: every difference is one this phase's checklist named. This
      diff does not prove the named inputs still reach the corpus, and is not read as if it did:
      `npx nx show project --json` reports the **declared** `inputs` array verbatim and expands neither `default` nor a
      `namedInputs` reference, so after this phase every affected target prints `["default", "behaviorCorpus"]` and the
      corpus path is absent from the output by construction. The two probe items below carry that half.
- [x] [AI] [AC-3] Probe the `behaviorCorpus` reference through the cache key, which is what `--json` cannot show. Nx
      hashes file content, so a content change under a genuinely resolved input must miss the cache. For each pairing
      below: run the target twice and confirm the second run prints `Nx read the output from the cache`; then
      `printf '\n# hash probe\n' >> <feature file>` and run it a third time, confirming that line is now absent; then
      `git checkout -- <feature file>` and run it a fourth time — acceptance: the fourth run prints the cache line
      again, and `git diff --stat -- <feature file>` prints nothing, so the probe left the corpus byte-identical. Assert
      it per file rather than with a bare `git status --short`, which is not empty at gate time because this phase's
      `project.json` edits are not committed until the last gate item. `git checkout --` restores each feature file from
      the **index**, not from `HEAD`, so it is safe only on a path carrying no unstaged edit of its own; that holds here
      because this plan edits no `.feature` file at all, which
      [specification-changes.md](tech-docs/specification-changes.md) states and the probe's own `git diff --stat`
      assertion re-checks. `# ` opens a Gherkin comment, so the appended line is valid Gherkin and changes no scenario.
      The three pairings are `badakmini-cli:test:unit` with
      `specs/apps/badakmini-cli/behavior/capability-parity.feature`, `wahidyankf-www:test:unit` with
      `specs/apps/wahidyankf-www/behavior/accessibility.feature`, and `wahidyankf-www-e2e:typecheck` with that same
      `wahidyankf-www` feature file. All three targets resolve to `cache: true` through the root `targetDefaults` and
      all three declare the corpus input today.
- [x] [AI] [AC-3] Probe the `workspaceScripts` reference the same way, because it is the other glob this phase hoists
      and it has no other proof: run `npx nx run wahidyankf-www:test:unit` to a cache hit, then
      `printf '\n// hash probe\n' >> scripts/next-with-port.mjs`, run it again and confirm the cache line is absent,
      then `git checkout -- scripts/next-with-port.mjs` and run it once more — acceptance: the cache line returns and
      `git diff --stat -- scripts/next-with-port.mjs` prints nothing. `git checkout --` restores that path from the
      **index**, not from `HEAD`, so it is safe only on a path carrying no unstaged edit of its own; that holds here
      because this plan never edits `scripts/next-with-port.mjs` — [tech-docs/file-impact.md](tech-docs/file-impact.md)
      gives it no label — so its index entry equals its `HEAD` entry. A trailing `//` comment is valid in an ES module,
      so the probe cannot break the script it borrows.
- [x] [AI] [AC-4] `grep -n 'apps/badakmini-cli' apps/badakmini-cli/project.json` — acceptance: no line printed is inside
      a `"command"` string. That is the whole acceptance, and it is the form [prd.md](prd.md) states for `[AC-4]`. The
      grep still prints lines this plan does not remove, so it must not be read as required to print only input paths:
      lines 4 and 5 are `"root": "apps/badakmini-cli"` and `"sourceRoot": "apps/badakmini-cli"`, which declare where the
      project lives rather than what a command runs and which this plan leaves alone. The only other surviving line is
      the `namedInputs` declaration, which the earlier checklist item pins with
      `grep -c 'specs/apps/badakmini-cli/behavior' apps/badakmini-cli/project.json` printing `1`; every `inputs` entry
      references it by name and holds no path of its own — **unless** this phase's conditional fired. If the `tests/e2e`
      cache probe hit the cache and the checklist item restored `{workspaceRoot}/apps/badakmini-cli/tests/e2e/**/*` to
      `test:coverage:behavior`, `test:coverage`, and `test:quick`, three more lines print here. Those three are `inputs`
      entries rather than `"command"` strings, so they do not fail this acceptance, and `learnings.md` already carries
      the entry that item required. Read against that entry: three extra input lines with no matching `learnings.md`
      record means something else put them there.
- [x] [AI] [AC-5]
      `grep -nE '"command": *"([^"]*[^-] )?nx run' apps/badakmini-cli/project.json apps/wahidyankf-www/project.json apps/wahidyankf-www-e2e/project.json`
      — acceptance: prints nothing and exits non-zero. The group is optional because the workspace's one bare invocation
      opens its command string, with no character and no space before `nx run`; a pattern that requires them matches
      that line not at all and is silent before the edit as well as after. Two paired controls keep this from being a
      check that cannot fail: the same grep run before the edit, in the checklist item above, must print the one
      `static-routes:validation` line, and `grep -c 'npm exec nx -- run' apps/wahidyankf-www/project.json` must print a
      non-zero count here, proving the file searched really does invoke Nx.
- [x] [AI] `npm run test:quick` — acceptance: exits 0 for both projects.
- [x] [AI] `npm run test:integration` — acceptance: exits 0 for both projects.
- [x] [AI] `npx nx run badakmini-cli:test:e2e` — acceptance: exits 0.
- [x] [AI] `npx nx run badakmini-cli:test:coverage` — acceptance: exits 0 and prints all three coverage lines, with the
      unit figure matching `local-tmp/coverage-badakmini-before.txt`.
- [x] [AI] `npx nx run wahidyankf-www:test:coverage` — acceptance: exits 0 and the unit line percentage matches
      `local-tmp/coverage-www-before.txt`.
- [x] [AI] `npm run check:markdown-links` — acceptance: exits 0.
- [x] [AI] [AC-4] `npm run check:harness-parity` — acceptance: exits 0, with Nx reporting the `capability-parity` target
      for `badakmini-cli` as successfully run. This phase rewrites that target's command — it gains `options.cwd` and
      loses `-C apps/badakmini-cli` — and nothing else reaches it: `.husky/pre-push` runs `capability-parity` only when
      a commit touches `.claude`, `.codex`, `.opencode`, or `.agents`, and this plan touches none of them, so without
      this item the rewritten command would ship unexecuted. Its three siblings are already covered and are not repeated
      here: `rule-change` runs on every commit from `.husky/pre-commit`, `governance` runs in the Phase 2 and Phase 3
      gates, and `markdown-links` runs in this gate above.
- [x] [AI] [AC-2] `npx nx run wahidyankf-www-e2e:specs:e2e:baseline` — acceptance: exits 0, which is the same assertion
      Phase 0 recorded at baseline — the generated `test.fixme` count still equals the 34 in
      `apps/wahidyankf-www-e2e/e2e-skip-baseline.json` — and this run is now made against the rewritten target. This
      phase gives that target a new `"outputs": ["{projectRoot}/.features-gen"]` declaration and moves its `inputs` onto
      `behaviorCorpus`, and nothing else in this gate reaches it: the project declares no `test:quick` and no
      `test:integration`, so neither `npm run test:quick` nor `npm run test:integration` selects it through `run-many`,
      and the only other item here that executes one of its targets runs `typecheck`. The `cache` and `outputs`
      inspections above do name this project, but they read `npx nx show project --json` and execute nothing. Without
      this item a cached target that the nightly `npm run test:scheduled` workflow runs against whatever `main` holds
      would be edited, committed, and pushed unrun. That is the same reasoning the `capability-parity` item above
      states, applied to the project this phase normalizes and Phase 2 deletes.
- [x] [AI] [AC-3] `npx nx run wahidyankf-www-e2e:lint` — acceptance: exits 0, with both `lint:biome` and
      `lint:commentary` reported as run. It is unreached by this gate's aggregates for the same reason, and it is the
      one of this project's five rewritten targets that is an `nx:run-commands` aggregate rather than a single
      `command`, so it exercises the `behaviorCorpus` substitution on a differently shaped target. The remaining two,
      `install` and `test:e2e`, are deliberately not run here and the reason is stated rather than left as a gap: each
      receives only the shared `behaviorCorpus` substitution, which the `typecheck` cache probe above already proves
      resolves for this project, and each is minutes of browser download and browser driving that Phase 2 runs in full
      at the merged target names.
- [x] [AI] [AC-2] Check the outputs rule mechanically, in the same inspection style as the `cache` check above. Run this
      once per project, substituting the project name and its artifact map:
      `npx nx show project badakmini-cli --json | node -e 'let s="";process.stdin.on("data",d=>s+=d).on("end",()=>{const t=JSON.parse(s).targets;const writes={"build":"{projectRoot}/dist","test:coverage:unit":"{workspaceRoot}/local-tmp/badakmini-unit.out"};const bad=[];for(const [k,v] of Object.entries(t)){if(v.cache===false&&v.outputs!==undefined)bad.push(k+": uncached yet declares outputs");if(v.cache===true&&writes[k]&&!(v.outputs||[]).includes(writes[k]))bad.push(k+": cached and writes "+writes[k]+", undeclared")}console.log(bad.length?"VIOLATIONS: "+bad.join("; "):"outputs rule holds for all "+Object.keys(t).length+" targets");process.exit(bad.length?1:0)})'`
      — acceptance: each of the three runs exits 0 and prints the `outputs rule holds for all N targets` form with a
      non-zero N, so the check is proved to have read a populated target set. The artifact maps are
      `{"build":"{projectRoot}/dist","test:coverage:unit":"{workspaceRoot}/local-tmp/badakmini-unit.out","typecheck":null}`
      for `badakmini-cli`,
      `{"build":"{projectRoot}/.next","test:coverage:unit":"{projectRoot}/coverage","typecheck":null}` for
      `wahidyankf-www`, and `{"specs:e2e:baseline":"{projectRoot}/.features-gen"}` for `wahidyankf-www-e2e`. The `null`
      is a documented entry rather than a hole in the map: it records `typecheck` as cached and writing only compiler
      state — `apps/wahidyankf-www/tsconfig.tsbuildinfo` under `"incremental": true`, and nothing at all inside the
      workspace for `go vet` — which the Phase 3 rule defines as not an artifact, because nothing reads it but the
      compiler that wrote it. The check needs no change to honour it: `writes[k]` is falsy for a `null`, so the second
      clause skips the target, and writing the key down is what stops a later reader from adding it as a miss. The map
      is written out rather than inferred because no command can tell which cached target writes something: the first
      clause catches the two inert `outputs` this phase removes, and the map is what catches the opposite miss. The
      aggregates are deliberately absent from every map — `test:coverage` and `test:quick` are cached and write nothing
      themselves, because each artifact belongs to a sub-target that declares it. Phase 2's gate re-runs this check and
      its `cache` sibling over the post-merge project set, two projects rather than three and with the `wahidyankf-www`
      map extended to carry the `specs:e2e:baseline` target that phase adds, and Phase 3's gate re-runs that post-merge
      form once more against the written rule. The map written above is the pre-merge one and is deliberately not the
      map Phase 2, Phase 3, or Phase 4 uses: run as written after Phase 2 it would inspect a target set it predates and
      name a project that no longer exists. A failure here is still cheaper to fix than the same failure a phase later,
      which is why the check runs in all three places rather than only the last.
- [x] [AI] Commit Phase 1 as four commits rather than one, each staged and pushed in turn, because the
      [thematic commits policy](../../../repo-governance/conventions/thematic-commits-policy.md) defines a theme by
      intent and this phase carries four — acceptance: each commit's diff contains nothing its message does not name,
      `git status --short` is empty after the last, and all four SHAs are written into the Execution Record as Phase 0
      requires.
  - [x] [AI] [AC-5] `fix(wahidyankf-www): resolve nx through the workspace binary` — the bare `nx run` change alone. It
        is a defect fix, not a normalization, and it is the only Phase 1 change that alters which binary runs; it
        commits first so a bisect that lands on it names one suspect.
  - [x] [AI] [AC-3] `refactor(workspace): declare shared behavior inputs once per project` — the three `namedInputs`
        declarations, every glob they replace, and the redundant `tests/e2e` input removal.
  - [x] [AI] [AC-4] `refactor(workspace): resolve project commands through options.cwd` — the thirteen `badakmini-cli`
        `cwd` declarations with the two `mkdir` and one `BADAKMINI_BIN` rewrites, and `static-routes:validation`'s move
        to `{projectRoot}`.
  - [x] [AI] [AC-2] `refactor(workspace): declare every target's cache state and outputs` — the `badakmini-cli`
        `test:coverage:unit` `cache`/`outputs` pair, the new `outputs` on `wahidyankf-www-e2e:specs:e2e:baseline`, and
        the two inert `outputs` removals from `wahidyankf-www`'s `test:coverage:integration` and `generate:cv-pdf`.

> **Pause Safety**: all three `project.json` files are written in one style, every target declares its cache state, and
> every gate that passed at baseline passes now at unchanged coverage figures. The workspace still has three projects
> and the browser suite still lives where it did. Safe to stop. Resume with `npm run test:quick`.

## Phase 2: Merge the E2E Project into the Application

The expand, migrate, verify, and contract steps of [migration-design.md](tech-docs/migration-design.md), landing in one
commit because the compatibility window is zero.

- [x] [AI] [AC-6] Confirm the pre-merge commit is a reachable recovery source before deleting anything: run
      `git rev-parse HEAD > local-tmp/pre-merge-sha.txt && git ls-tree -r --name-only "$(cat local-tmp/pre-merge-sha.txt)" -- apps/wahidyankf-www-e2e/steps | wc -l`
      — acceptance: prints `8`, proving the eight step files are restorable from that commit.
- [x] [AI] [AC-6] Record what the stricter compiler settings currently report, so the accepted loss is measured rather
      than assumed: run `npx nx run wahidyankf-www-e2e:typecheck` — acceptance: exits 0, and the result is written into
      `learnings.md` as the pre-merge state of the step files under `noUncheckedIndexedAccess`, `noUnusedLocals`, and
      `noUnusedParameters` all set to `true`.
- [x] [AI] [AC-6] Resolve the two scenario counts the same comment block states, which cannot both describe the same
      quantity: it says the four unbound feature files hold nineteen scenarios and that `specs:e2e:baseline` "holds that
      count to nineteen", while `e2e-skip-baseline.json` records 34 and the target compares against that. Both readings
      are consistent if a `Scenario Outline` generates one `test.fixme` per example row, so measure rather than assume.
      From `apps/wahidyankf-www-e2e`, run `npx bddgen`, then `grep -rc 'test\.fixme(' .features-gen/*.spec.js` for the
      per-file split and `grep -rho 'test\.fixme(' .features-gen | wc -l` for the total — acceptance: the two recorded
      numbers are compared against that output, whichever of them the measurement contradicts is corrected in the
      comment, and the measured per-file counts for `env-loader`, `port-resolver`, `tier-env-loading`, and `cv-export`
      are written into `learnings.md` with the total. Correct the comment in
      `apps/wahidyankf-www-e2e/playwright.config.ts` here, before the `git mv` below carries the file, so the moved copy
      is already right and a later failure is attributable to the merge rather than to the prose. The baseline file's
      `34` is not edited unless the measured total differs from it: that number is the assertion the target makes, while
      the comment is prose about the same corpus and is the side free to be wrong.
- [x] [AI] [AC-6] Create the destination with `mkdir -p apps/wahidyankf-www/tests/e2e/steps` — acceptance: the directory
      exists and `apps/wahidyankf-www/tests/` now holds `bdd`, `e2e`, and `integration`.
- [x] [AI] [AC-6] Move the eight step files with
      `git mv apps/wahidyankf-www-e2e/steps/*.ts apps/wahidyankf-www/tests/e2e/steps/` — acceptance:
      `ls apps/wahidyankf-www/tests/e2e/steps | wc -l` prints `8` and `git status --short` shows eight `R` rename
      entries rather than delete-plus-add pairs.
- [x] [AI] [AC-6] Move the baseline with
      `git mv apps/wahidyankf-www-e2e/e2e-skip-baseline.json apps/wahidyankf-www/tests/e2e/e2e-skip-baseline.json` —
      acceptance: the moved file still contains `{ "skippedScenarios": 34 }`.
- [x] [AI] [AC-6] Move the Playwright configuration with
      `git mv apps/wahidyankf-www-e2e/playwright.config.ts apps/wahidyankf-www/playwright.config.ts` — acceptance: the
      file exists at the new path and `apps/wahidyankf-www-e2e/` no longer contains it.
- [x] [AI] [AC-6] Edit `apps/wahidyankf-www/playwright.config.ts`: change the `defineBddConfig` `steps` value from
      `"./steps/**/*.ts"` to `"./tests/e2e/steps/**/*.ts"` — acceptance: `featuresRoot` and `webServer.cwd` are left at
      their existing `"../.."`-relative values, because the file's depth below the workspace root is unchanged.
- [x] [AI] [AC-6] Edit `apps/wahidyankf-www/playwright.config.ts`: correct the comment sentence that reads "`test:e2e`
      declares `dependsOn` on `wahidyankf-www:build`" so it names the intra-project dependency, reading "`test:e2e`
      declares `dependsOn` on `build`" — acceptance: `grep -n 'dependsOn' apps/wahidyankf-www/playwright.config.ts`
      prints that one comment line and it no longer carries the `wahidyankf-www:` prefix, matching the `dependsOn` the
      target actually declares. The `specs:e2e:baseline` comment lower in the same block needs no edit: it already names
      the bare target with no project prefix. Neither correction can be observed by asking whether a comment names a
      project that no longer exists — `wahidyankf-www` survives, so that question is answered `no` before the edit as
      well as after.
- [x] [AI] [AC-6] Edit `apps/wahidyankf-www/package.json`: add `"@axe-core/playwright": "4.10.1"`,
      `"@playwright/test": "1.62.1"`, and `"playwright-bdd": "9.2.0"` to `devDependencies` — acceptance: the three pins
      are character-identical to those in `apps/wahidyankf-www-e2e/package.json` before deletion.
- [x] [AI] [AC-6] Run `npm install` from the repository root immediately after that edit — acceptance: exits 0,
      `npm ls playwright-bdd -w wahidyankf-www` exits 0 and prints `playwright-bdd@9.2.0` beneath the workspace line
      ending `-> ./apps/wahidyankf-www`, and `ls node_modules/.bin/bddgen` succeeds at the repository root. Both are
      asserted at the root because npm workspaces hoist here: `apps/wahidyankf-www-e2e/node_modules` does not exist
      today and the binary the deleted project runs is already the root one, so an acceptance naming a project-local
      `node_modules` would fail on a correct install. `npm ls` is the half that observes the new dependency specifically
      — the root binary is present before this edit too, because the sibling project depends on the same package. This
      install cannot wait for the one at the end of the phase: every `wahidyankf-www:test:e2e` run below invokes
      `npx bddgen` and `npx playwright test` from `apps/wahidyankf-www`, and until the workspace is linked those resolve
      to nothing. The later `npm install` is still required and is not redundant — it regenerates `package-lock.json`
      after the workspace directory is deleted, which is a different change from adding a dependency.
- [x] [AI] [AC-6] Edit `apps/wahidyankf-www/tsconfig.json`: add `".features-gen"` to the `exclude` array beside
      `"node_modules"` — acceptance: the array holds both entries. The reason is defence in depth, and it is stated
      plainly rather than dressed as a live defect, because a wrong reason is what makes a later reader delete the right
      line. `bddgen` emits only `*.feature.spec.js`, and this `tsconfig.json`'s `include` lists `next-env.d.ts`,
      `**/*.ts`, `**/*.tsx`, `.next/types/**/*.ts`, and `.next/dev/types/**/*.ts` — no `.js` pattern at all — so
      `typecheck` does not pick those files up as roots today. `allowJs` is `true`, so a `.js` file reached by an import
      from an included file would still be compiled, and nothing under `src/` or `tests/` imports `.features-gen`. The
      exclude therefore changes nothing observable now; it holds if `include` ever gains a `.js` pattern or `bddgen`
      ever emits `.ts`. No gate item asserts that removing it breaks anything, because none can.
- [x] [AI] [AC-6] Edit `apps/wahidyankf-www/eslint.config.mjs`: add a second configuration object for
      `files: ["tests/e2e/steps/**/*.ts"]` carrying the same three `jsdoc` rules as the existing block, with
      `languageOptions.parser` set to `tsParser` and no `ecmaFeatures.jsx` — acceptance: from the repository root,
      `node --input-type=module -e 'const c=(await import("./apps/wahidyankf-www/eslint.config.mjs")).default; console.log(JSON.stringify(c.map(b=>({files:b.files,jsx:b.languageOptions?.parserOptions?.ecmaFeatures?.jsx,rules:Object.keys(b.rules||{})}))))'`
      exits 0 and prints two objects: the first unchanged, with its `src` globs, `"jsx":true`, and its three `jsdoc/*`
      rule names; the second naming `tests/e2e/steps/**/*.ts`, listing the same three rule names in the same order, and
      printing no `jsx` key at all. The omitted `jsx` is the point rather than an oversight: a Playwright step file
      contains no JSX, so enabling the feature there would be configuration the code cannot exercise.
- [x] [AI] [AC-1] [AC-6] Edit `apps/wahidyankf-www/project.json`: add an `install` target running
      `npx playwright install --with-deps chromium` with `"cache": false`, `"options": {"cwd": "{projectRoot}"}` —
      acceptance: `npx nx run wahidyankf-www:install` exits 0, which is all this command observes — `playwright install`
      exits non-zero when it cannot place the Chromium build it was asked for, and reports nothing about a suite. That
      the suite can actually drive Chromium is observed by the first `wahidyankf-www:test:e2e` run below, not here.
- [x] [AI] [AC-1] [AC-6] Edit `apps/wahidyankf-www/project.json`: add a `test:e2e` target carrying the
      unconditional-`test.skip` guard and `npx bddgen && npx playwright test` from the deleted project's target, with
      `"cache": false`, `"options": {"cwd": "{projectRoot}"}`, `"dependsOn": ["build"]`, and
      `"inputs": ["default", "behaviorCorpus"]` — acceptance: `dependsOn` names the intra-project `build` rather than
      `wahidyankf-www:build`, matching how `badakmini-cli:test:e2e` names its own `build`.
- [x] [AI] [AC-6] Edit the guard's search path in that same command: its trailing `.` becomes `tests/e2e` — acceptance:
      the `grep -rn` argument reads `tests/e2e`, so the guard scans one directory as it did before the merge. Carried
      over unchanged it would scan the whole application, because `.` is the working directory and that is now
      `apps/wahidyankf-www`: forty-three `.ts` files plus `.next/`, which the command's `--exclude-dir` list does not
      name.
- [x] [AI] [AC-6] Prove the guard is green at rest: run `npx nx run wahidyankf-www:test:e2e` — acceptance: the run exits
      0, having passed the guard rather than printing its `ERROR: unconditional test.skip() found` message, reached
      `npx bddgen`, which is visible in the output, and driven Chromium over the eight bound feature files. Exit 0
      rather than merely reaching `bddgen` is the acceptance, because this is the first `test:e2e` run the merged
      project has ever had and it is where the module-system delta [migration-design.md](tech-docs/migration-design.md)
      inventories would first surface — after `bddgen` has written the generated ESM and `npx playwright test` loads it,
      which is downstream of everything a `reaches bddgen` acceptance observes. **If it fails on a module-resolution
      error**, that is that delta, and it is the one failure in this phase that is not fixed inside Phase 2. The deleted
      project's `package.json` declared `"type": "module"`; `apps/wahidyankf-www/package.json` declares no `type`, so
      the ESM that `bddgen` writes into `apps/wahidyankf-www/.features-gen/` now sits in a package Node resolves as
      CommonJS. Record the exact error in `learnings.md`, stop, and return to the owner. Do not add `"type": "module"`
      to the application manifest and do not add a second manifest under `tests/e2e/`: neither is sanctioned by any
      decision in [brd.md](brd.md), a `type` on a Next.js application manifest changes how every other `.js` in that
      package resolves, and improvising it here would widen the plan's scope. This run is also the only thing that
      determines the question — neither this plan nor its technical set claims the merge will or will not fail this way.
      The two `test:e2e` runs below and the one in this phase's gate carry the same disposition by reference and do not
      restate it.
- [x] [AI] [AC-6] Prove the guard still fires, because a guard that never fires is indistinguishable from one that scans
      nothing: append the throwaway comment line `// test.skip(1)` to
      `apps/wahidyankf-www/tests/e2e/steps/theme.steps.ts` and run `npx nx run wahidyankf-www:test:e2e` again —
      acceptance: the run exits non-zero, prints `ERROR: unconditional test.skip() found in test files above`, and never
      reaches `bddgen`. The injected line is a comment on purpose and is not to be "corrected" into real code: the guard
      is a `grep -rn` over `*.ts` text, and its pattern `'\$?test\.skip\([^,)]*\)'` matches the comment, while
      `test:e2e` carries `dependsOn: ["build"]` and the merged `tsconfig.json` includes `**/*.ts`, so a real
      `test.skip(1);` statement would fail the Next.js type check with TS2304 — `theme.steps.ts` imports only
      `createBdd` and `expect`, never `test` — and the run would die at `next build` before the guard ever ran. Three
      distinct non-zero exits are therefore possible, and only one observes the guard. A non-zero exit whose output
      names a TypeScript error rather than the guard message means the probe, not the guard, is wrong. A non-zero exit
      whose output names a module-resolution error means the guard did not fire at all: its `exit 1` precedes
      `npx bddgen`, so nothing downstream of the guard runs when the guard does fire, and reaching a resolution failure
      proves the run got past it. It is not a fresh sighting of the module-system delta, which the first `test:e2e` run
      above already settled by exiting 0, so it does not take that run's stop-and-return disposition. A zero exit
      carries the same verdict as that third case: the guard is scanning the wrong directory and the search path edit
      above is wrong.
- [x] [AI] [AC-6] Remove the throwaway line with `git checkout -- apps/wahidyankf-www/tests/e2e/steps/theme.steps.ts`
      and run `npx nx run wahidyankf-www:test:e2e` a third time — acceptance: `git status --short` shows that path only
      as the staged `R` rename it was before the injection, with no unstaged modification marker, and the run is green
      again, so the injection is proved undone rather than assumed undone. `git checkout --` restores the path from the
      **index**, not from `HEAD`, so it is safe only on a path carrying no unstaged edit of its own; that holds here
      because the `git mv` above staged this file and nothing in this phase edits a step file's content afterwards, so
      the index entry is the moved file's pre-injection content. It does not hold for
      `apps/wahidyankf-www/playwright.config.ts`, which this phase does edit after its move, and the Phase 2 gate stages
      the whole phase before probing that one. Leaving this confirming run to the Phase 2 gate would be enough to catch
      the injection, but not enough to attribute it.
- [x] [AI] [AC-1] [AC-6] Edit `apps/wahidyankf-www/project.json`: add a `specs:e2e:baseline` target carrying the deleted
      project's command with exactly one substitution, with `"cache": true`, `"options": {"cwd": "{projectRoot}"}`,
      `"outputs": ["{projectRoot}/.features-gen"]` carried over from Phase 1, and `inputs` naming only `default` and
      `behaviorCorpus`. The command is not copied verbatim: its `require('./e2e-skip-baseline.json')` is repointed to
      `require('./tests/e2e/e2e-skip-baseline.json')` to match the moved file. Its other two relative paths are left
      alone and the reason is stated so a reader does not "fix" them — `npx bddgen` and the
      `grep -rho 'test\.fixme(' .features-gen` that follows it are both resolved against `options.cwd`, which is
      `{projectRoot}` in both the old target and the new one, and `bddgen` writes `.features-gen` into whichever project
      root it runs in — acceptance: `grep -c 'e2e-skip-baseline' apps/wahidyankf-www/project.json` prints `1` and that
      occurrence names `./tests/e2e/`, while the `.features-gen` argument is unchanged from the deleted target. The
      three `{projectRoot}` paths the deleted project listed separately as `inputs` all sit inside `{projectRoot}` once
      the files move, so the built-in `default` input already covers them; listing them again is the same redundancy
      Phase 1 removes from `badakmini-cli`, and the Phase 2 gate's three cache probes are what prove `default` reaches
      each of them.
- [x] [AI] [AC-6] Edit `apps/wahidyankf-www/project.json` target `lint:commentary`: change the command to
      `eslint --config eslint.config.mjs src tests/e2e/steps` — acceptance: `npx nx run wahidyankf-www:lint:commentary`
      exits 0, and, because ESLint prints nothing at all on a clean run, the files it read are observed with a second
      command rather than inferred from that silence. From `apps/wahidyankf-www`, run
      `npx eslint --config eslint.config.mjs src tests/e2e/steps --format json > ../../local-tmp/eslint-commentary.json`,
      then from the repository root run
      `node -e 'const r=require("./local-tmp/eslint-commentary.json");const c=p=>r.filter(x=>x.filePath.includes(p)).length;console.log("src:"+c("/src/")+" steps:"+c("/tests/e2e/steps/"))'`
      — acceptance: the JSON formatter emits one entry per linted file whether or not it found anything, so this prints
      a non-zero count for each directory, with `steps:8` naming the eight moved step files. The redirect to a file is
      deliberate: it keeps the JSON out of a pipe that an output filter could rewrite.
- [x] [AI] [AC-6] Delete the six remaining files of the retired project:
      `git rm apps/wahidyankf-www-e2e/project.json apps/wahidyankf-www-e2e/package.json apps/wahidyankf-www-e2e/tsconfig.json apps/wahidyankf-www-e2e/eslint.config.mjs apps/wahidyankf-www-e2e/.gitignore apps/wahidyankf-www-e2e/README.md`
      — acceptance: the root `.gitignore` already covers `.features-gen/`, `playwright-report/`, and `test-results/`, so
      deleting the project-local ignore file leaves nothing newly visible to `git status --short`.
- [x] [AI] [AC-6] Remove the now-empty directory with `rm -rf apps/wahidyankf-www-e2e` and run `npm install` —
      acceptance: `npx nx show projects` returns exactly `wahidyankf-www` and `badakmini-cli`, and `package-lock.json`
      no longer contains a `wahidyankf-www-e2e` workspace entry.
- [x] [AI] Edit `package.json`: in `test:scheduled`, replace `wahidyankf-www-e2e:specs:e2e:baseline` with
      `wahidyankf-www:specs:e2e:baseline` and `wahidyankf-www-e2e:test:e2e` with `wahidyankf-www:test:e2e` — acceptance:
      `grep -o 'wahidyankf-www-e2e' package.json | wc -l` prints `0`, where it prints `2` today, **and** each new
      invocation is named exactly: `grep -o 'nx run wahidyankf-www:specs:e2e:baseline' package.json | wc -l` and
      `grep -o 'nx run wahidyankf-www:test:e2e' package.json | wc -l` each print `1`, where each prints `0` today. The
      absence grep alone would pass on a typo, because nothing in this phase runs `test:scheduled`; the Phase 4 gate
      does, two phases later, which is where a typo would otherwise first surface. Occurrence counts rather than
      `grep -c` line counts, because `test:scheduled` is one line and holds both invocations.
- [x] [AI] Edit `.github/workflows/full-bdd.yml`: change the browser install step to `npx nx run wahidyankf-www:install`
      — acceptance: `grep -c 'npx nx run wahidyankf-www:install' .github/workflows/full-bdd.yml` prints `1`, where it
      prints `0` today, and `grep -c 'wahidyankf-www-e2e' .github/workflows/full-bdd.yml` prints `0`; and the target
      that step names resolves, which `npx nx show project wahidyankf-www --json` shows by listing `install` among its
      targets, as it does not today. `npm run check:workflows` still runs in this phase's gate and still must exit 0,
      but it is a constraint on the edit rather than an observation of it: it is `actionlint`, which validates the YAML
      and the shell embedded in it and knows nothing about Nx target names, so it passes over `wahidyankf-www:install`,
      over the retired `wahidyankf-www-e2e:install`, and over a typo alike. Delivered on `check:workflows` alone, this
      edit would ship unobserved.
- [x] [AI] Edit `apps/README.md`: remove the `wahidyankf-www-e2e` bullet from `## Current Applications` — acceptance:
      two bullets remain and neither links a directory that no longer exists.
- [x] [AI] Edit `apps/wahidyankf-www/README.md`: replace the "sibling `wahidyankf-www-e2e` project" adapter bullet with
      one naming `tests/e2e/` — acceptance: the three adapter bullets still describe three layers, none names a deleted
      path, and the sentence above them still reads "Three adapters bind that one corpus", because this edit rehomes the
      browser adapter rather than adding one.
- [x] [AI] Edit `apps/wahidyankf-www/README.md` again: fold in the retired README's record of the browser layer whole,
      because this phase deletes the only document in the repository that holds it and repoints
      `specs/apps/wahidyankf-www/behavior/README.md` at this one for exactly that content. Three things move and none is
      a summary of the others. First, the list of the four feature files the Playwright adapter deliberately does not
      bind — `env-loader.feature`, `tier-env-loading.feature`, `port-resolver.feature`, and `cv-export.feature` — each
      with the layer that binds it instead. Second, the recorded skip baseline as it stands after the measurement item
      above, 34 unless that measurement corrected it, stated as generated tests rather than scenarios, and naming
      `apps/wahidyankf-www/tests/e2e/e2e-skip-baseline.json` as the file that records it and
      `npx nx run wahidyankf-www:specs:e2e:baseline` as the target that asserts it. Third, the standing rule, carried
      verbatim: "Raise the number only when a scenario is deliberately left unbound, and say here why." — acceptance:
      `grep -c 'Raise the number only when a scenario is deliberately left unbound' apps/wahidyankf-www/README.md`
      prints `1`; `grep -c 'env-loader.feature'`, `grep -c 'tier-env-loading.feature'`, and
      `grep -c 'port-resolver.feature'` over that file each print a non-zero count where each prints `0` today;
      `grep -c 'cv-export.feature'` prints at least `2`, because the Integration adapter bullet already names it once;
      and `grep -n '34'` prints at least one line, where it prints none today. This is the same class as
      `workspace-commands.md`'s "once per machine" sentence and is resolved the same way: a clause whose only home is a
      document this phase deletes moves before the deletion rather than being lost with it. This README is not a
      governed document — `npm run check:governance` measures `AGENTS.md`, `CLAUDE.md`, `repo-governance/`, and the
      harness READMEs — so the added words are under no word limit.
- [x] [AI] [AC-6] Edit `specs/apps/wahidyankf-www/README.md`: rewrite the two-projects sentence to name one project,
      change the Process E2E adapter row's path to `apps/wahidyankf-www/tests/e2e/steps/`, repoint the skip-baseline
      sentence at the merged project, and change the two verification-command rows to the `wahidyankf-www:` prefix —
      acceptance: `grep -c 'wahidyankf-www-e2e' specs/apps/wahidyankf-www/README.md` prints `0`, **and** each of the
      four edits is observed positively, because the absence count alone is satisfied by deleting a row or by writing a
      target name that does not exist. Over that same file, all counted today so each is proved capable of failing:
      `grep -c 'apps/wahidyankf-www/tests/e2e/steps/'` prints `1`, where it prints `0` today, naming the rehomed Process
      E2E adapter row; `grep -c 'nx run wahidyankf-www:test:e2e'` prints `1` and
      `grep -c 'nx run wahidyankf-www:specs:e2e:baseline'` prints `1`, each `0` today, naming the two rewritten
      verification rows; `grep -c 'npx nx run wahidyankf-www:'` prints `5`, where it prints `3` today, which is the row
      count itself and fails if either verification row is deleted rather than rewritten; and `grep -c 'skip baseline'`
      still prints `1`, so the skip-baseline sentence is repointed rather than dropped. Neither `wahidyankf-www:` grep
      matches `wahidyankf-www-e2e:`, because the character after `www` differs, so none of these counts can be satisfied
      by the unedited file.
- [x] [AI] [AC-6] Edit `specs/apps/wahidyankf-www/architecture.md`: remove the `apps/wahidyankf-www-e2e` project-table
      row and the shared-model sentence above it, and redraw the Container View to the exact block
      [specification-changes.md](tech-docs/specification-changes.md) writes out under "After, and this is the block the
      Phase 2 edit produces character for character". Nothing about the replacement is left to be chosen here: that
      document holds the before and the after side by side, and its `[Container: Playwright]` stereotype becomes
      `[Test-time process]` while the `Containers` section's opening "One container." sentence and its single-row table
      are deliberately left untouched, for the reason stated there — acceptance: the diagram is still one fenced `text`
      ASCII block and matches that after-block byte for byte, which is asserted rather than eyeballed with
      `diff <(sed -n '/^After, and this is the block the Phase 2 edit produces character for character:$/,/^```$/p' plans/in-progress/project-json-consistency/tech-docs/specification-changes.md | sed -n '/^```text$/,/^```$/p') <(sed -n '/^## Containers$/,/^## Components$/p' specs/apps/wahidyankf-www/architecture.md | sed -n '/^```text$/,/^```$/p')`
      printing nothing and exiting 0. Run the same command before the edit as the paired control: it must exit 1 and
      print the three-line-versus-two-line difference in the lower node, which is what proves the comparison reaches the
      diagram rather than two empty extractions. Both `sed` ranges are anchored on headings and fences this plan does
      not move. Also `grep -c 'wahidyankf-www-e2e' specs/apps/wahidyankf-www/architecture.md` prints `0`, where it
      prints `2` today. Those two lines are the table row and the old diagram node; the Scope sentence and the paragraph
      below the diagram describe the project without naming it, which is why the count is `2` rather than `4` and why it
      cannot serve as the acceptance for either of them.
- [x] [AI] [AC-6] Edit the paragraph below that diagram in `specs/apps/wahidyankf-www/architecture.md` to keep the
      toolchain-difference fact while dropping the separate-project inference, and to say why a box that is not a
      container is drawn in a container view — acceptance: it still states that the adapter starts the application
      through `next start`, drives it over HTTP through a browser, is a different toolchain from the in-process behavior
      adapter, and exists only at test time and is never deployed; and `npm run check:markdown-links` still exits 0 over
      the file. The `grep -c` in the item above is not this item's proof and must not be read as one: this paragraph
      holds the retired name in none of its sentences today, so the count reaches `0` whether the paragraph is
      rewritten, deleted, or left exactly as it is. The listed sentence content is the only thing that observes it.
- [x] [AI] Edit `specs/apps/wahidyankf-www/behavior/README.md`: repoint the sentence naming
      `apps/wahidyankf-www-e2e/README.md` at `apps/wahidyankf-www/README.md` — acceptance:
      `grep -c 'wahidyankf-www-e2e' specs/apps/wahidyankf-www/behavior/README.md` prints `0`, and the document the
      sentence now points at is read for the content the sentence claims it holds: the sentence says that document
      "names them and records the generated skip baseline that keeps the gap from widening unnoticed", so
      `apps/wahidyankf-www/README.md` must satisfy the four `grep` counts and the skip-baseline greps of the fold-in
      item above. Run those greps here as well, against the repointed target, rather than treating the fold-in item's
      tick as sufficient. `npm run check:markdown-links` is deliberately not this item's acceptance: the reference is
      inline code, `` `apps/wahidyankf-www-e2e/README.md` ``, and not a Markdown link, so the link check never resolves
      it and would pass over a pointer aimed at a document carrying none of the content.
- [x] [AI] Edit `repo-governance/development/workspace-commands.md`: change the three `wahidyankf-www-e2e` narrower-run
      lines to `wahidyankf-www:install`, `wahidyankf-www:test:e2e`, and `wahidyankf-www:specs:e2e:baseline`. Then edit
      the three-sentence paragraph below that block rather than deleting it: delete its first two sentences, the ones
      stating that `wahidyankf-www-e2e` owns no `test:quick` and that this is the shape the testing and BDD policies
      give a permitted dedicated project, and rewrite the third under the merged target name so it still tells a reader
      to run `wahidyankf-www:install` once per machine before `test:e2e`, which builds and starts `wahidyankf-www`
      itself — acceptance: `grep -c 'wahidyankf-www-e2e' repo-governance/development/workspace-commands.md` prints `0`,
      `grep -c 'once per machine' repo-governance/development/workspace-commands.md` still prints `1`, and
      `npm run check:governance` keeps the document under 750 words. Only the first two sentences lose their subject at
      the merge. The third is live guidance that stays true under the merged name, and this document is the only place
      in the repository that states it: the other file that shows the command, `apps/wahidyankf-www-e2e/README.md`, is
      deleted in this same phase, so deleting the whole paragraph would drop the browser-install instruction from the
      repository entirely.
- [x] [AI] [AC-6] Edit `repo-governance/development/testing-policy/tooling.md`: in Recorded Deviations, remove the
      `apps/wahidyankf-www-e2e` clause from the `target` sentence and the sentence beginning "The E2E project follows
      the runner it hosts", and reword the two `Both` references so they name only `apps/wahidyankf-www` — acceptance:
      `grep -c 'wahidyankf-www-e2e' repo-governance/development/testing-policy/tooling.md` prints `0`, the surviving
      paragraph still records `apps/wahidyankf-www`'s three overrides and the Next 16 reason for them, and
      `npm run check:governance` exits 0. This lands in Phase 2 rather than Phase 3 because Phase 2 deletes
      `apps/wahidyankf-www-e2e/tsconfig.json`, so leaving the entry until Phase 3 would leave one committed state where
      governance records a deviation of a file that does not exist.
- [x] [AI] [AC-6] Edit `specs/apps/README.md`: rewrite the Directory Map line for `wahidyankf-www` so the behavior
      corpus is described as belonging to that one project rather than "shared with the dedicated E2E project
      `apps/wahidyankf-www-e2e`" — acceptance: `grep -c 'wahidyankf-www-e2e' specs/apps/README.md` prints `0` and the
      line still states the feature count it states today.
- [x] [AI] [AC-6] Edit `apps/wahidyankf-www/tests/bdd/accessibility.steps.ts`: in the comment above the axe-core `Then`,
      repoint `apps/wahidyankf-www-e2e/steps/accessibility.steps.ts` at
      `apps/wahidyankf-www/tests/e2e/steps/accessibility.steps.ts` — acceptance: only the comment text changes, the
      `@covers` line below it is untouched,
      `grep -c 'wahidyankf-www-e2e' apps/wahidyankf-www/tests/bdd/accessibility.steps.ts` prints `0`, and
      `npm run test:quick` exits 0. This is the only `.ts` file Phase 2 edits that is not one of the eight moved step
      files.
- [x] [AI] Run `npm run format` — acceptance: exits 0 and `npm run format:check` afterwards also exits 0.

### Phase 2 Gate

> Every check below passes before Phase 3 begins. A failure is fixed inside Phase 2, with one named exception: a
> `wahidyankf-www:test:e2e` failure on a module-resolution error is recorded and returned to the owner, for the reason
> the checklist's first `wahidyankf-www:test:e2e` run states. That exception binds this phase's checklist as well as its
> gate, because the checklist runs `wahidyankf-www:test:e2e` three times before the gate does and the first of those is
> where the failure would surface.

- [x] [AI] Stage this phase's whole change before any probe below runs: `git add -A` — acceptance: `git status --short`
      lists every path this phase touched as a staged entry and nothing in its unstaged column, and `git diff --stat`
      with no pathspec prints nothing. Nothing is committed here, and this does not replace the staging the commit item
      below does: the checkboxes ticked in this document while the rest of the gate runs are written after this item and
      are staged with that commit. This is the same staging the Phase 3 gate performs before `check:rule-change`, but
      here it is a safety precondition rather than a visibility one. The three cache probes below restore their files
      with `git checkout -- <path>`, which restores from the **index**, not from `HEAD`, so it is safe only on a path
      that carries no unstaged edit. `apps/wahidyankf-www/playwright.config.ts` is the path that makes this
      load-bearing: `git mv` stages a rename carrying the file's `HEAD` content and leaves any working-tree modification
      unstaged, so this phase's three edits to that file — the scenario-count comment correction made before the move,
      the `steps` glob repoint, and the `dependsOn` comment correction — sit unstaged on top of a staged rename until
      this item runs. Probing before staging would discard all three, and the `specs:e2e:baseline` item below would then
      fail with a skip-baseline count misattributed to a broken step binding rather than to the discarded edits.
      `local-tmp/`, `.features-gen/`, `playwright-report/`, and `test-results/` are all named in the root `.gitignore`,
      so `-A` stages none of them.
- [x] [AI] [AC-6] `npx nx show projects` — acceptance: returns exactly two entries, `wahidyankf-www` and
      `badakmini-cli`.
- [x] [AI] [AC-1] `npx nx show project wahidyankf-www --json` — acceptance: the target list contains all ten contract
      targets: `typecheck`, `lint`, `test:unit`, `test:integration`, `test:e2e`, `test:coverage`, `test:coverage:unit`,
      `test:coverage:integration`, `test:coverage:behavior`, and `test:quick`.
- [x] [AI] [AC-1] `npx nx show project badakmini-cli --json` — acceptance: the same ten targets are present, so both
      projects expose one identical contract.
- [x] [AI] [AC-2] Re-run the Phase 1 gate's `cache` inspection over the post-merge project set, which is `badakmini-cli`
      and `wahidyankf-www` and no longer includes `wahidyankf-www-e2e`. The command is the Phase 1 one unchanged,
      substituting the project name — acceptance: both runs exit 0, each printing the `all N targets declare cache` form
      with a non-zero N rather than an `UNDECLARED:` line. Phase 1 ran this check before this phase added `install`,
      `test:e2e`, and `specs:e2e:baseline` to `wahidyankf-www` and before it deleted the third project, so the Phase 1
      run observed neither the three new targets nor the target set that survives the merge. It runs here rather than in
      Phase 3, which is where it would otherwise first see the post-merge set, because a failure is fixed by editing
      `apps/wahidyankf-www/project.json` — this phase's own subject, and inside this phase's commit theme. The same fix
      forced in Phase 3 would land in a commit messaged `docs(testing): state the project target contract` carrying a
      diff that message does not name, which the
      [thematic commits policy](../../../repo-governance/conventions/thematic-commits-policy.md) forbids and which the
      Phase 1 commit item asserts of every commit this plan makes.
- [x] [AI] [AC-2] Re-run the Phase 1 gate's `outputs` inspection over the same two projects, with the same one-liner,
      substituting the project name and its artifact map. `badakmini-cli`'s map is unchanged from Phase 1.
      `wahidyankf-www`'s is **extended** for the post-merge target set and is the map every later re-run uses:
      `{"build":"{projectRoot}/.next","test:coverage:unit":"{projectRoot}/coverage","typecheck":null,"specs:e2e:baseline":"{projectRoot}/.features-gen"}`.
      The `wahidyankf-www-e2e` map is retired with its project; its single entry is exactly what the new
      `specs:e2e:baseline` key carries forward, so nothing the Phase 1 map covered stops being covered. `install` and
      `test:e2e` are `cache: false` and appear in no map, which is the first clause's business rather than the second's
      — acceptance: both runs exit 0 and print `outputs rule holds for all N targets` with a non-zero N. This is the map
      [prd.md](prd.md)'s `[AC-2]` proof is reconciled against in Phase 4; the Phase 1 map predates `specs:e2e:baseline`
      on `wahidyankf-www` and would pass while missing it. A failure in either of these two items is a defect in what
      Phase 1 or this phase delivered, and it is fixed in the `project.json` that carries it before this gate passes,
      for the placement reason the item above states. Re-run `git add -A` after any such fix, so the staged tree the
      three cache probes below restore from still matches the working tree.
- [x] [AI] [AC-3] Probe that the built-in `default` reaches the three paths whose explicit `inputs` entries the merge
      drops, rather than reading them back out of `npx nx show project --json`, which prints only the declared array and
      expands neither `default` nor `behaviorCorpus`. `specs:e2e:baseline` is `cache: true`, so run it twice and confirm
      the second run prints `Nx read the output from the cache`, then repeat this three times, once per path: append a
      content change, re-run, confirm the cache line is absent, `git checkout --` the file, and re-run to a hit. The
      three probe files are `apps/wahidyankf-www/tests/e2e/e2e-skip-baseline.json` (append one blank line with
      `printf '\n' >>`, which changes the bytes and leaves the JSON parsing to the same `skippedScenarios` value of 34),
      `apps/wahidyankf-www/tests/e2e/steps/theme.steps.ts` (append `\n// hash probe\n`), and
      `apps/wahidyankf-www/playwright.config.ts` (append `\n// hash probe\n`) — acceptance: each of the three misses the
      cache and each restores to a hit, and `git diff --stat` names none of the three afterwards. Assert it per file
      rather than with a bare `git status --short`, which at gate time still holds the whole uncommitted merge. Every
      `git checkout --` here restores the path from the **index**, not from `HEAD`, which is why the staging item at the
      top of this gate is a precondition of running this one at all: all three paths arrive as staged `git mv` renames
      whose staged content is the pre-move `HEAD` content, and `apps/wahidyankf-www/playwright.config.ts` carries three
      unstaged edits on top of its rename until that staging runs. Run this probe before it and the restore silently
      reverts them. A hit on any of the three means `default` does not reach that path and the explicit entry must be
      restored rather than dropped. Only `specs:e2e:baseline` is probed, not `test:e2e`: `test:e2e` is `cache: false` in
      both the deleted project and the merged one, so Nx never hashes its inputs and no cache probe can observe them.
      Its `inputs` are declared for readability, and nothing about them can be proved or broken by this change.
- [x] [AI] [AC-6] `npx nx run wahidyankf-www:specs:e2e:baseline` — acceptance: exits 0, which asserts the generated
      `test.fixme` count is still exactly the 34 recorded in the moved baseline file. A count above 34 means a step file
      stopped binding at its new path.
- [x] [AI] [AC-6] `npx nx run wahidyankf-www:test:e2e` — acceptance: exits 0, having built the application, started
      `next start`, and driven Chromium over the eight bound feature files. A module-resolution failure here takes the
      disposition the checklist's first `test:e2e` run states — record the exact error in `learnings.md`, stop, and
      return to the owner, adding no `"type": "module"` and no second manifest — and is not fixed inside Phase 2. That
      item holds the reason and this one does not restate it. Reaching this gate at all means that run already exited 0,
      so a resolution failure appearing only here is a change since it, not the delta's first sighting.
- [x] [AI] `npx nx run wahidyankf-www:typecheck` — acceptance: exits 0, run after the preceding `test:e2e` has populated
      `apps/wahidyankf-www/.features-gen/`. What that observes is the outcome — a populated `.features-gen/` does not
      break the application's type check — and not the `tsconfig.json` exclude, which this item is not presented as
      proving. The run passes with the exclude absent too, because `bddgen` emits only `*.feature.spec.js` and the
      project's `include` carries no `.js` pattern. Ordering it after `test:e2e` is still worth doing: run against an
      absent directory it would observe nothing at all.
- [x] [AI] `npm run test:quick` — acceptance: exits 0 for both projects.
- [x] [AI] `npm run test:integration` — acceptance: exits 0 for both projects.
- [x] [AI] `npx nx run badakmini-cli:test:e2e` — acceptance: exits 0.
- [x] [AI] `git grep -n 'wahidyankf-www-e2e' -- ':!plans/'` — acceptance: prints nothing and exits non-zero, so every
      reader in [migration-design.md](tech-docs/migration-design.md)'s inventory has been rewritten. The whole `plans/`
      tree is excluded, not only `plans/done/`: the archived migration plan records the name as history and this plan's
      own documents record it because they describe the change that retires it. Neither depends on the name resolving to
      a project, and this plan's documents still hold the name in dozens of places while the gate runs, so a narrower
      pathspec could never pass. Pair it with `git grep -c 'wahidyankf-www-e2e' -- 'plans/' | wc -l` printing a non-zero
      count, which proves the pattern was searched against files that really do contain it.
- [x] [AI] `npm run check:markdown-links` and `npm run check:governance` — acceptance: both exit 0.
- [x] [AI] `npm run check:workflows` — acceptance: exits 0 over the edited scheduled workflow. This is a constraint on
      the edit, not the observation of it: `actionlint` knows nothing about Nx target names. The greps in that edit's
      own checklist item are what observe the new invocation.
- [x] [AI] Commit with message `refactor(wahidyankf-www): co-locate the browser E2E adapter`, without pushing yet —
      acceptance: `git status --short` is empty and `git log -1 --oneline` names the commit.
- [x] [AI] `npx nx affected -t test:quick --base=origin/main --head=HEAD` — acceptance: exits 0 and selects
      `wahidyankf-www`, confirming the merged project is still reachable by the pre-push calculation after the graph
      change. This item sits between the commit and the push, and the order is load-bearing rather than incidental:
      Phase 1 ended by pushing, so before the commit above `origin/main` and `HEAD` are the same ref and the phase's
      work is uncommitted, and `nx affected` reports `No tasks were run` — a green result that observes nothing. After
      the push `origin/main` moves to this commit and the diff empties again. The window in which the base is the Phase
      1 commit and the head is this one is exactly the window in which the diff is Phase 2's change.
- [x] [AI] Push to `main` — acceptance: `git status --short` is empty, `git log origin/main -1 --oneline` names the
      commit above, and its SHA is written into the Execution Record as Phase 0 requires.

> **Pause Safety**: the workspace holds two projects, each exposing the same ten targets and each passing the `cache`
> and `outputs` inspections over the post-merge target set, and the browser suite runs from inside the application it
> tests at an unchanged skip baseline of 34. The Gherkin corpus is byte-identical to the baseline and no application
> code changed. Safe to stop. Resume with `npx nx run wahidyankf-www:test:e2e`.

## Phase 3: Write the Target Contract Rule

Promotes the shape both projects now share out of this plan and into governance, so a third project has a rule to read
rather than two files to compare.

- [x] [AI] [AC-7] Edit `repo-governance/development/testing-policy.md`: in the Quick Tests section, replace the partial
      target list with the full contract naming `typecheck`, `lint`, `test:unit`, `test:integration`, `test:e2e`,
      `test:coverage:unit`, `test:coverage:integration`, `test:coverage:behavior`, `test:coverage`, and `test:quick`,
      and state which are eligibility-dependent — acceptance: the existing rule that a library defines no
      `test:integration` when it owns no local boundary, and that a library never owns `test:e2e`, is preserved rather
      than contradicted.
- [x] [AI] [AC-7] Edit `repo-governance/development/testing-policy.md`: state that every target declares `cache`
      explicitly wherever the root `nx.json` `targetDefaults` does not reach it, so no target resolves to an undeclared
      state — acceptance: the sentence names `targetDefaults` as the other source, so a reader does not add a redundant
      declaration to the six targets it already covers.
- [x] [AI] [AC-7] Edit `repo-governance/development/testing-policy.md`: state that a single-command target declares
      `options.cwd` rather than encoding its own project path in the command, and that a cached target that writes an
      artifact declares `outputs` while an uncached one declares none — acceptance: the `outputs` sentence gives the
      reason, that Nx replays a cache hit and restores nothing when a cached target declares no output path.
- [x] [AI] [AC-7] Edit `repo-governance/development/testing-policy.md`: in the same place, define what "writes an
      artifact" means — something a later target or a person consumes — and state that a compiler's own incremental
      state is not one, because nothing reads it but the compiler that wrote it and it is regenerated on demand —
      acceptance: the definition is written as a definition, not as an exception, and `wahidyankf-www:typecheck` is the
      case it settles: that target resolves to `cache: true` through the root `targetDefaults`, its `tsconfig.json` sets
      `"incremental": true`, and `apps/wahidyankf-www/tsconfig.tsbuildinfo` exists on disk, so without the definition
      the rule above would read as requiring an `outputs` declaration for a file no consumer has. The rule still binds
      every target with no carve-out, which is the item below; this says what the rule's own terms mean.
      `badakmini-cli:typecheck` is the same shape and needs no separate wording: `go vet` writes only into the Go build
      cache, outside the workspace entirely.
- [x] [AI] [AC-7] Edit `repo-governance/development/testing-policy.md`: state that the shape rules above bind every
      target a project declares, not only the ten in the contract — acceptance: the sentence is stated without an
      exception, which is what Phase 1's `generate:cv-pdf` and `static-routes:validation` edits make true; a rule that
      shipped with a carve-out would teach the next reader to add a second.
- [x] [AI] [AC-7] Edit `repo-governance/development/testing-policy.md`: state that `options.commands` expresses the
      ordered gate itself and `dependsOn` expresses a prerequisite that must precede the whole gate — acceptance: the
      distinction is stated as a rule with both halves, which is what makes `wahidyankf-www:test:quick`'s `dependsOn` on
      `static-routes:validation` a documented choice rather than an unexplained second ordering mechanism.
- [x] [AI] [AC-7] Edit `repo-governance/development/testing-policy.md`: state that a shared input path is declared once
      as a project-level `namedInputs` entry and referenced by name — acceptance: the sentence explains that the
      alternative, repeating the glob, is what let the three files drift into different input sets.
- [x] [AI] [AC-7] Run `npm run check:governance` — acceptance: exits 0, and
      `wc -w repo-governance/development/testing-policy.md` reports under 750. If it reports 700 or more, the document
      has entered the headroom band and the
      [document word limit policy](../../../repo-governance/conventions/document-word-limit-policy.md) governs how it is
      fixed; move the added detail into `repo-governance/development/testing-policy/` rather than dropping any of it.
- [x] [AI] [AC-7] Edit `repo-governance/development/nx-workspace-policy.md`: add one sentence in Required Approach
      linking to the testing policy's target contract — acceptance: a reader arriving to add a target is sent to the
      shape rule, and `npm run check:governance` still exits 0 for this document too.
- [x] [AI] [AC-7] Edit `docs/how-to/run-nx-workspace.md`: in the paragraph that already names `project.json` and
      `nx.json`, add the pointer to the target contract — acceptance: the how-to links the policy rather than restating
      it, so the two cannot drift.
- [x] [AI] Run `npm run format` — acceptance: exits 0 and `npm run format:check` afterwards also exits 0.

### Phase 3 Gate

> Every check below passes before Phase 4 begins. A failure is fixed inside Phase 3.

- [x] [AI] [AC-7] `npm run check:governance` — acceptance: exits 0, holding every edited governance document under the
      750-word limit.
- [x] [AI] `npm run check:markdown-links` — acceptance: exits 0 over the three new cross-document links.
- [x] [AI] [AC-2] Re-run both inspections from the Phase 2 gate unchanged, now against the rule this phase has written:
      the `cache` one and the `outputs` one, over `badakmini-cli` and `wahidyankf-www`, with the extended
      `wahidyankf-www` artifact map carrying `specs:e2e:baseline` — acceptance: all four runs exit 0 and print the
      `all N targets declare cache` and `outputs rule holds for all N targets` forms with a non-zero N, and the
      properties they test are the ones the sentences added to `repo-governance/development/testing-policy.md` state, so
      the written rule and its mechanical check are recorded as agreeing rather than assumed to. This is the form
      [prd.md](prd.md)'s `[AC-2]` proof is reconciled against in Phase 4. The re-run is read-only and cannot force a
      `project.json` edit: the same four commands passed in the Phase 2 gate against these same two files, and no item
      in this phase touches either. A failure here therefore means those files changed outside this plan — record it in
      the Execution Record, stop, and return to the owner rather than repairing configuration inside a documentation
      phase, where the repair would land in a commit messaged `docs(testing): state the project target contract`
      carrying a diff that message does not name.
- [x] [AI] [AC-7] Read `repo-governance/development/testing-policy.md` against `apps/badakmini-cli/project.json` and
      `apps/wahidyankf-www/project.json`. This stays a review rather than a command, because the rules are prose and no
      check parses them — but it writes down what it compared, so the record exists to be disputed: append to
      `learnings.md` one line per rule the policy now states, naming the rule, the target in each file it was checked
      against, and the verdict — acceptance: every rule the document states appears in that list with a verdict of holds
      for both files, and no rule is recorded as contradicted. A rule the delivered files violate is a defect in the
      rule or in the files, and is fixed inside this phase rather than noted. This item is `[AC-7]`'s proof;
      `check:governance` and `check:markdown-links` in this same gate are constraints on the edit, not observations of
      it, because a policy naming none of the ten targets would pass both.
- [x] [AI] Stage this phase's edits first with
      `git add repo-governance/development/testing-policy.md repo-governance/development/nx-workspace-policy.md docs/how-to/run-nx-workspace.md`,
      then run `npm run check:rule-change` — acceptance: `git diff --cached --name-only` lists exactly those three
      paths, with no `project.json` among them, because the inspection re-run above is read-only and this phase edits no
      executable configuration — and the check reports them and names
      [Rules Propagation](../../../repo-governance/workflows/rules/rules-propagation.md); the check reports without
      blocking, so the acceptance is that the notice appears and the named workflow is then run. The staging is
      load-bearing rather than housekeeping: the check reads `git diff --cached --name-only`, so an unstaged edit is
      invisible to it and the run would pass by seeing nothing. Phase 0's `git add -N` remedy does not transfer — that
      one makes a new, untracked document visible to `check:markdown-links`, while these three are already-tracked files
      whose modifications only a real `git add` puts in the index. Staging here also costs nothing later: the commit at
      the end of this gate stages the rest of the phase's work, the ticked checkboxes among it.
- [x] [AI] Run the [Rules Propagation](../../../repo-governance/workflows/rules/rules-propagation.md) workflow for the
      edited policies — acceptance: it reports no unresolved contradiction between the new contract and the
      [Nx workspace policy](../../../repo-governance/development/nx-workspace-policy.md), the
      [BDD policy](../../../repo-governance/development/behavior-driven-development-policy.md), or
      [workspace commands](../../../repo-governance/development/workspace-commands.md).
- [x] [AI] `npm run test:quick` — acceptance: exits 0 for both projects, confirming a documentation-only phase changed
      no executable behavior.
- [x] [AI] Commit with message `docs(testing): state the project target contract` and push to `main` — acceptance:
      `git status --short` is empty and the commit's SHA is written into the Execution Record as Phase 0 requires. This
      is the SHA Phase 4's file-impact reconciliation uses as the far end of its range.

> **Pause Safety**: the contract both projects satisfy is written in `testing-policy.md`, every governance check passes,
> and the `cache` and `outputs` inspections still pass over the post-merge target set they first ran against in the
> Phase 2 gate. No executable configuration changed in this phase. Safe to stop. Resume with `npm run check:governance`.

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
      system — acceptance: each criterion's proof named in
      [specification-changes.md](tech-docs/specification-changes.md) is re-run and passes, and any criterion that cannot
      be reconciled is recorded rather than ticked. Six of the seven proofs are commands and are simply re-run;
      `[AC-7]`'s is the Phase 3 gate review, so reconciling it means re-reading the recorded rule-by-rule comparison
      against the two `project.json` files as they stand now. `[AC-2]`'s two inspections are re-run in their
      **post-merge** form — the one the Phase 2 gate introduces and the Phase 3 gate re-runs — over `badakmini-cli` and
      `wahidyankf-www` and with the extended `wahidyankf-www` artifact map that carries `specs:e2e:baseline`. The Phase
      1 gate form is not the one to re-run here and re-running it would be a false reconciliation: it names three
      projects, one of which no longer exists, and its map predates the target Phase 2 added, so it would report the
      rule holding over a set it never inspected.
- [ ] [AI] Reconcile [tech-docs/file-impact.md](tech-docs/file-impact.md) against the diff of every delivery commit: the
      four Phase 1 commits, the single Phase 2 commit, and the single Phase 3 commit. Take the range end to end rather
      than by count — `git diff --stat <phase-0 SHA>..<phase-3 SHA>`, reading both SHAs from the Execution Record where
      Phase 0 required each phase to write them, so a later split of any phase into more commits cannot make this item
      wrong. Do not resolve either end by commit message: `docs(plans): start project-json-consistency` already names a
      commit in this repository's history, so a `--grep` anchor is ambiguous — acceptance: every path the document
      labels under Phase 1, Phase 2, or Phase 3 appears in the diff with the labelled operation, and every path in the
      diff appears in the document. This plan's own documents appear in the diff too, and that is expected rather than a
      failure: [the execution workflow](../../../repo-governance/workflows/plan-execution/03-gates-and-pushes.md) stages
      each phase's ticked checkboxes and any `learnings.md` entries with that phase's work, so `delivery.md` appears in
      every commit in the range and `learnings.md` in each phase whose items wrote an entry. Both are labelled in the
      document's own Plan Documents section, which states that expectation. The document's Phase 4 rows sit outside the
      range by design: `plans/done/README.md` and the folder move land in the archival commit at the end of this phase,
      and `plans/in-progress/README.md` is edited at the Phase 0 end — or was already committed before execution began,
      as Phase 0's first commit item allows — and again in that archival commit. Its Not Touched entries, `nx.json` and
      the gitignored `local-tmp/` evidence among them, must appear in neither.
- [ ] [AI] Give a dated, evidence-backed disposition to the first of this plan's two conditional items: the Phase 1
      checklist item that restores the `tests/e2e` input if the cache probe over
      `apps/badakmini-cli/tests/e2e/README.md` hits the cache instead of missing it — acceptance: it records either the
      restoration and its evidence, or `Not triggered` with the probe output showing the run that followed the content
      change reported no `Nx read the output from the cache` line.
- [ ] [AI] Give a dated, evidence-backed disposition to the second: the Phase 1 `[AC-5]` pre-edit control, which writes
      to `learnings.md` only if the bare-`nx run` grep run against the unedited `apps/wahidyankf-www/project.json` does
      not print exactly the one `static-routes:validation` line — acceptance: it records either that entry and what it
      concludes about the gate pattern, or `Not triggered` with the one line the pre-edit run printed.
- [ ] [AI] Give a dated, evidence-backed disposition to the third: Phase 2's module-resolution branch, which writes to
      `learnings.md` only if the checklist's first `npx nx run wahidyankf-www:test:e2e` fails on a module-resolution
      error — acceptance: it records either that entry, the exact error it captured, and what the owner decided when the
      plan stopped and returned to them, or `Not triggered` with that run's exit 0. These three are the conditionals
      `learnings.md`'s own header comment and [tech-docs/file-impact.md](tech-docs/file-impact.md) both name as writing
      here only if triggered, so reconciling some and not all would leave a named writer with no terminal state and
      archival would be blocked on an entry nobody was told to look for.
- [ ] [AI] Run the [plan-quality-gate](../../../repo-governance/workflows/plan-quality-gate.md) workflow at strict level
      — acceptance: `plan-checker` reports no findings.
- [ ] [AI+HUMAN] Confirm with the owner that the plan is complete before it is archived — acceptance: the owner agrees
      the delivered scope matches what they asked for; archival is a one-way move and the four scope decisions in
      [brd.md](brd.md) narrowed twice during planning.
- [ ] [AI] Move the folder to `plans/done/YYYY-MM-DD__project-json-consistency/` using the date the final commit landed
      — acceptance: the destination does not already exist; if it does, stop rather than merging, overwriting, or
      inventing a suffix.
- [ ] [AI] Update `plans/in-progress/README.md` and `plans/done/README.md` in the same change — acceptance: neither of
      the source index's two lists still names this plan, returning to their `None right now.` and
      `No plan folders right now.` placeholders if no other plan is active; the destination index gains a dated entry
      beside the two already there; and `npm run check:markdown-links` resolves every archived internal link.

### Phase 4 Gate

> Every check below passes before the plan is considered finished.

- [ ] [AI] `npm run check:markdown-links` — acceptance: exits 0 after the folder move, with no dead link left by the
      changed depth of the archived documents.
- [ ] [AI] `git grep -n 'plans/in-progress/project-json-consistency' -- ':!plans/done/'` — acceptance: prints nothing
      and exits non-zero, confirming no document outside the archive still points at the pre-archival path.
      `plans/done/` is excluded because the archived plan carries the literal string in prose it wrote about itself —
      `delivery.md` and `tech-docs/file-impact.md` both name their own in-progress path — and that prose is history, not
      a pointer to be rewritten. Before the move,
      `git grep -c 'plans/in-progress/project-json-consistency' -- ':!plans/done/'` names exactly those two files and
      nothing else, so the move alone is what makes this gate pass.
- [ ] [AI] `npm run test:quick` and `npm run test:integration` — acceptance: both exit 0 for both projects.
- [ ] [AI] `npm run test:scheduled` — acceptance: exits 0, running the full ordered verification including
      `wahidyankf-www:specs:e2e:baseline` and `wahidyankf-www:test:e2e` at their merged names. This is the one command
      that exercises every renamed invocation in `package.json` end to end.
- [ ] [AI] Commit the archival move with a message naming the plan and push to `main` — acceptance: `git status --short`
      is empty and the source path is absent from the working tree.

> **Pause Safety**: the plan is archived, every learning has a terminal disposition, and the full scheduled verification
> passes at the merged target names. Nothing is left in progress. Resume is not applicable; the work is complete.
