# Delivery Checklist

## Execution Record

- 2026-09-01: Plan committed to `main` after the
  [plan-quality-gate](../../../repo-governance/workflows/plan-quality-gate.md) loop ended on its seven-cycle bound with
  every finding resolved. The plan [`README.md`](README.md) records that run: seven cycles, every finding resolved,
  status `settled`.
- 2026-09-01: Phase 0 source-cleanliness check passed vacuously on its first run — the pathspec named this repository's
  `apps/wahidyankf-www-e2e` rather than `ose-public`'s `apps/wahidyankf-www-fe-e2e`, and `git status --porcelain` prints
  nothing and exits 0 for a pathspec matching no tracked file. The item now pairs its emptiness assertion with
  `git ls-files --error-unmatch` over the same eight paths. Re-run: both halves clean.
- 2026-09-01: Phase 0 gate passed. `npm install`, `npm run test:quick`, `npm run format:check`, and
  `npm run check:markdown-links` all exited 0 before any change; `$SRC` confirmed at
  `e74818fc06c4c104725383384d2aa38305a503ef` with all eight source paths tracked and clean.
- 2026-09-01: Phase 1 gate failed at `npm audit --audit-level=low`. `tsx@4.21.0`, the pin inherited from `ose-public`,
  depends on `esbuild@0.27.7` and carries GHSA-g7r4-m6w7-qqqr. Fixed inside the phase by moving the pin to `4.23.13`,
  which resolves `esbuild@0.28.2`; not by an `overrides` entry and not by lowering the audit level. Full gate re-run:
  all five items green.
- 2026-09-01: Phase 1 complete. TypeScript `6.0.3`, Biome `2.5.11` as a linter only, ESLint `9.39.4`,
  `eslint-plugin-jsdoc` `64.3.2`, and `tsx` `4.23.13` pinned exactly; root `biome.json` authored with its formatter and
  assist disabled; `scripts/next-with-port.mjs` ported, repointed, and rewritten for this repository; six generated
  paths ignored at the root.
- 2026-09-01: Phase 2 found one ordering defect in this checklist. The behaviour `README.md` item precedes the
  `architecture.md` item but links to the file that one creates, so its `check:markdown-links` acceptance could not pass
  in checklist order. The two were executed in the opposite order; nothing else depends on the sequence.
- 2026-09-01: Phase 2 gate passed. Eleven feature files and 53 scenarios in `specs/apps/wahidyankf-www/behaviours/`,
  scenario titles diffed against both source corpora with no loss, all 53 conforming to the one-primary-step rule, and
  the as-built C4 model authored.
- 2026-09-01: Phase 3 found five defects that seven quality-gate cycles did not. `lint:commentary` could not parse one
  source file, because installing `eslint` and `eslint-plugin-jsdoc` installs no reader for TypeScript or JSX; the
  conditional item that would have dropped the check entirely was read against its wording and found not to have fired.
  The corpus and its binding kept `OSE_WWW_PORT`, a variable this repository does not have, which no planned sweep looks
  for. A `noArrayIndexKey` fix had changed a card's `id` alongside the `key` it reported, breaking two scenarios.
  `port-resolver.unit.test.ts` imported `EnvRecord` from a module that never exported it, caught by `next build` at push
  time rather than by the `typecheck` that had already passed before the file existed. And a new `CvContent` test made
  the unit suite flaky at Vitest's five-second default, which a per-project `testTimeout` fixed once it was learned that
  projects inherit no more of the root `test` block than they inherit of its plugins. One item was ticked late: `dotenv`
  was pinned when the manifest was written and its checkbox missed until a Close the Migration sweep printed the line.
- 2026-09-01: Phase 3's `rhino-cli|dotnet|\.fsproj|F#` sweep found two matches and both stay. They are `"F#"` in the
  `programmingLanguages` of two entries in the owner's project record — true statements about what those projects are
  written in, in the file that is now the repository's single authoritative CV. The three tokens `[AC-7]` actually names
  find nothing anywhere in scope, and `F#` finds nothing once that one data file is excluded. The sweep was narrowed
  rather than the data edited.
- 2026-09-01: Phase 3 gate passed. Unit coverage reached **100% lines** from a 97.99% start with no threshold lowered
  and no exclusion added; integration coverage is **100% lines** over exactly the two PDF modules; all 55 scenarios
  across twelve feature files run, 53 through `tests/bdd/` and two at the filesystem boundary. `test:quick`,
  `test:integration`, `npm run test:quick`, `format:check`, `check:markdown-links`, `npm audit --audit-level=low`, and
  `check:governance` all exit 0, `ls libs` prints only `README.md`, and all three components `tooling.md` names hold.
  `check:rule-change` named three edited rule paths and its Rules Propagation trigger was worked rather than dismissed,
  adding one reference in `code-style-policy.md` to the deviation register in `tooling.md`.

- 2026-09-01: Phase 4 complete. `cv/` is deleted and the application holds the repository's single CV record. Recovery
  was rehearsed rather than assumed: `git archive` from the commit preceding the deletion returns all seven files, and
  their SHA-256 digests diff clean against the ones taken before the copy, which is the only proof the recovery source
  is real for `cv-ats.md`, `cv-ats.pdf`, and `generate-cv-ats-pdf.py` — the three files deleted outright with no copy
  anywhere else in the tree. The archive was taken from `HEAD~1` rather than the `HEAD` the item names, because the
  deletion was committed first; that is the same commit the item means.
- 2026-09-01: Phase 4's Harness Alignment run found a defect no sweep in this plan looks for. Repointing the CV routing
  from `cv/README.md` to `apps/wahidyankf-www/docs/README.md` preserved a duplication that predates the plan:
  `AGENTS.md` and `CLAUDE.md` each carried the whole rule, which the
  [agent instruction alignment policy](../../../repo-governance/conventions/agent-instruction-alignment-policy.md)
  forbids a derivative and `CLAUDE.md`'s own preamble forbids itself. The sentence was removed from `CLAUDE.md` and
  recorded as an item added during execution. The Rules Propagation run alongside it added nothing:
  `apps/wahidyankf-www/docs/README.md` needs no index entry, because the documentation index policy's scope reaches no
  path under `apps/`, and `check:markdown-links` already protects the routing link.
- 2026-09-01: Phase 4 gate passed. `git ls-files cv` is silent and `test ! -e cv` succeeds; the seven-file recovery and
  its digest diff are clean; `check:markdown-links`, `check:governance`, `check:harness-parity`, `test:quick`, and
  `format:check` all exit 0, with 12 test files and 258 tests green. `check:rule-change` named both Rules Propagation
  and Harness Alignment for `CLAUDE.md` and both were worked; it names only that file at gate time because it reads
  staged paths and the phase's eight other rule-path edits were committed earlier in the phase as `f338a16`, where the
  same check named all nine.

- 2026-09-01: Phase 5 complete. `apps/wahidyankf-www-e2e` is delivered as a dedicated Playwright project driving a real
  `next start` with no container: 36 scenarios pass and 34 generated tests are skipped, from a cold `.next`, with the
  skip count held to a tracked baseline. The eight bound feature files reach 36 `@covers` notes, unchanged in count
  across the corpus repointing, and the four unbound ones are named with the layer that binds each.
- 2026-09-01: Phase 5 found four defects the plan did not anticipate, three of them in ported code that was green in
  `ose-public`. `tsconfig.json` carried `baseUrl`, which TypeScript 6 reports as an error once the local `5.8.3` pin is
  dropped. `@axe-core/playwright` declares an open `playwright-core` range no pinning item in this plan reaches, so npm
  installed two versions of one library and `AxeBuilder({ page })` took a structurally different `Page`; pinning
  `@playwright/test` up fixed the types but broke `playwright-bdd@8.5.1` at runtime, and moving that to `9.2.0` as well
  fixed both and cleared five moderate advisories `npm audit` had begun reporting. And `accessibility.steps.ts` asserted
  a navigation link named `Personal Projects` that neither this application nor `ose-public`'s has ever rendered — the
  sibling binding of the same scenario in the application already asserted the correct label, so two bindings of one
  scenario disagreed.
- 2026-09-01: Phase 5's `specs:e2e:baseline` command could not pass as the plan wrote it, for a reason no review would
  catch by reading. `node -p` renders through `util.inspect`, which colourises a number when stdout is a TTY, and Nx
  runs a target under a pty — so the comparison failed with the self-contradicting message
  `expected 34 fixme scenarios, found 34`, the expected value carrying ANSI codes the found value did not. `NO_COLOR=1`
  did not suppress it; reading the value with `node -e` and `process.stdout.write` did. Everything else in the command
  is verbatim.
- 2026-09-01: Phase 5 gate passed. `test:e2e` from a cold `.next`, `test:quick`, `typecheck`, `lint`, `check:workflows`,
  `format:check`, `check:markdown-links`, and `check:governance` all exit 0, and `npm audit --audit-level=low` reports
  no vulnerability. `check:rule-change` named Rules Propagation for the two `repo-governance/development/` edits and it
  was worked: its idempotency gate stopped a rule from being added, because `testing-policy.md` and the BDD policy role
  matrix already settle between them that a dedicated E2E project owns equivalent targets rather than the full set.
  Harness Alignment was correctly not triggered — no harness file changed.

- 2026-09-01: Phase 6 complete. `vercel.json` is ported unchanged and its provenance is proved rather than asserted: the
  SHA-256 taken before anything reformatted it matches the source byte for byte, and after `npm run format` the bytes
  differ while a deep-equality probe shows the parsed configuration identical. That is the whole point of taking the
  digest first, and it is the eighth and last evidence file. `repo-governance/development/deployment-policy.md` is
  authored at 379 words, `git branch --list 'prod-*'` prints nothing, and the configuration is therefore present and
  inert — the state the new policy names as expected for a project mid-migration.
- 2026-09-01: Phase 6's Rules Propagation run caught a duplication in the edit that triggered it. The sentence added to
  `AGENTS.md` began by restating that plans deliver to `main`, which `AGENTS.md` already says in its Planning section
  and which the plans organization policy owns; it was rewritten to carry only what is new. This is the second phase
  running where the workflow's value was in what it stopped rather than what it added — Phase 5's idempotency gate ruled
  out a rule, and this one ruled out a restatement.
- 2026-09-01: Phase 6 gate passed. `check:governance`, `check:markdown-links`, `format:check`, `check:harness-parity`,
  and `test:quick` all exit 0. Harness Alignment was triggered alongside Rules Propagation because `AGENTS.md` changed,
  and found nothing to change: no derivative mentions deployment, promotion, or a `prod-` branch at all.

- 2026-09-01: Phase 5's commit checkbox was ticked late. The seven gate checks above it were run and recorded in one
  pass and this one was missed in the same pass, which is the second time in this plan a finished item kept an unticked
  box; the first, in Phase 3, is recorded above. Both were found by reading the file rather than by any check, because
  nothing verifies that a tick matches the tree.
- 2026-09-01: Phase 7 reconciled all ten acceptance criteria against the delivered system in one pass, writing the
  result into a `## Reconciliation` section of `prd.md`. All ten hold. `[AC-2]` and `[AC-3]` were written against a 99%
  floor and the delivered system reaches 100% lines on both, with no threshold lowered and no exclusion added.
- 2026-09-01: Phase 7 gave all eight dormant recovery items a dated disposition — five as unticked checkboxes, two as
  clause notes inside ticked items, and one, the `@vitejs/plugin-react` removal, as evidence because it fired. Each of
  the five standalone bullets replaced an earlier note promising the disposition in future tense, so the plan no longer
  says the same thing twice in two tenses.
- 2026-09-01: Phase 7's learnings triage routed twelve entries into six documents and discarded two halves with a
  reason. It reached seven paths no phase had anticipated — `code-style-policy.md`, `dependency-selection-policy.md`,
  `behaviour-driven-development-policy.md`, `delivery-checklists.md`, and all three `plan-checker` prompts — because
  which documents a triage reaches cannot be known before the learnings exist. Those seven were added to
  `tech-docs/file-impact.md` as `[E]` entries rather than improvised past, which is what that document's own opening
  rule asks for.
- 2026-09-01: Phase 7's first strict `plan-checker` cycle reported eleven findings against a plan that had already
  passed seven strict cycles before execution. Two were high: the Execution Record stopped at Phase 6 while Phase 7 work
  already sat in the tree, and `file-impact.md` mapped none of the seven governance and harness paths the triage had
  just edited. The rest were the plan's own drift from execution, listed here by kind rather than one clause per
  finding, because several were filed as one finding over several sites — five finished Phase 7 items still unticked,
  three ticked items whose acceptance their note contradicted, two Phase 7 items with no observable acceptance, a Phase
  7 gate missing `check:harness-parity` while the phase edited three harness prompts, a rollback design still promising
  one commit per phase after four phases landed as commit series, and the duplicate dispositions. All eleven were
  resolved rather than argued down; the three contradicted acceptances were rewritten to the criterion actually met and
  marked amended, rather than the ticks being removed.

- 2026-09-01: Phase 1 rewrote an acceptance criterion after it failed. The `tsx` item accepted on `npx tsx --version`
  printing the pinned version; under this session's shell proxy that command's output is rewritten, so the criterion
  tested the proxy rather than the pin. It now reads the manifest and the resolved tree instead. An acceptance that
  reads a tool's own output is only as trustworthy as whatever sits between the tool and the terminal.
- 2026-09-01: Phase 3 added an item during execution. The ported `port-resolver.ts` carried a header naming
  `OSE_WWW_PORT`, a variable this repository does not have; the rewrite that fixed it reached only that file, and the
  corpus and its binding kept the old name until a later item caught them. No sweep in the plan looks for that token —
  the specifier sweep matches nothing here and the .NET sweep has no alternative for it — so it was found by reading
  rather than by a check.
- 2026-09-01: Phase 3 hit a `check:markdown-links` failure that forced out-of-order execution, the same class of
  ordering defect Phase 2 hit. Two dead links sat in the stale `apps/wahidyankf-www/README.md` the port carried in, and
  the item that replaces that README sits several items later in Close the Migration, so the failing acceptance could
  not pass in checklist order. The README item was executed early, as Phase 2's pair was.
- 2026-09-01: Phase 4 added an item during execution. The Harness Alignment run its gate requires found that `CLAUDE.md`
  restated the CV routing rule `AGENTS.md` owns — a duplication predating this plan that the repointing preserved — and
  it was removed at the derivative rather than the canonical source.
- 2026-09-01: Phase 5 amended an acceptance criterion. The CI browser-install item probed for the literal
  `playwright install` in `full-bdd.yml`, which would mean copying the install command into CI beside the copy in
  `project.json`; the step invokes the project's own `install` target instead and the criterion now names that target.
  Two of this plan's amended acceptances came from the same instinct — a probe written against an implementation rather
  than against the outcome.
- 2026-09-01: Phase 7's second strict `plan-checker` cycle verified all eleven cycle-1 fixes and reported eight new
  findings, every one of them a place the delivered plan had drifted from the executed one rather than a defect in the
  work: an Execution Record missing six events, a note still describing a criterion that had just been amended, a
  `## Quality Gate` section not in the shape the workflow specifies, dependency counts stale by one package and one
  removal, a scope bullet narrower than the file map beside it, a deviation count stated as three where the register
  holds two, a triage summary whose arithmetic did not match its own table, and a gate item ordered below checks it ran
  before. All eight were resolved.
- 2026-09-01: Phase 7's third strict `plan-checker` cycle verified all eight cycle-2 fixes and reported five findings,
  and unlike cycle 2's these were not all plan drift — two were defects in the delivered repository. The triage had
  added three clauses to `delivery-checklists.md` and mirrored only one into the three `plan-checker` prompts, which
  `check:harness-parity` cannot catch because it compares subagent presence rather than prompt content, and which two
  places in the plan already asserted was done; both remaining clauses were mirrored. `specs:e2e:baseline` was invoked
  by nothing automated, so it was wired into `test:scheduled` ahead of the suite and `workspace-commands.md` records the
  order. The triage's own additions had pushed `delivery-checklists.md` into the word limit policy's headroom band,
  which only relocation closes, so its Execution Record section moved to a new sibling,
  `repo-governance/conventions/plans-organization-policy/execution-record.md`, indexed in the parent policy and the
  child README with the one inbound anchor repointed. The last two were plan drift: a misattributed caret-range pin in
  `tech-docs/README.md` and a `## Quality Gate` section still carrying the per-cycle table that
  [plan quality gate](../../../repo-governance/workflows/plan-quality-gate.md) forbids. All five were resolved, and
  `file-impact.md` gained the four paths the relocation reached.
- 2026-09-01: Phase 7's fourth strict `plan-checker` cycle verified all five cycle-3 fixes and reported seven findings,
  one of which cycle 3 created. Wiring `specs:e2e:baseline` into `test:scheduled` broke the Phase 5 item that had
  accepted on that script printing its Phase 3 value "and nothing else changed" — a criterion that forbids every later
  edit to a shared script cannot survive one, so it was amended to name the delivered value. Four Phase 1 items still
  accepted on `npx <tool> --version`, which is the form the clause this plan's own triage added to
  `delivery-checklists.md` now forbids; each was amended to read the resolved tree, which also made true an Execution
  Record line above that had claimed the `tsx` criterion was already repaired. The rest were plan drift: the Execution
  Record had no line for cycle 3 at all, a `specification-changes.md` proof cell named a repository-wide sweep the plan
  does not run and could not pass, a scope bullet still said seven paths where the map says eleven, and a sentence
  asserted an archival gate record that is not written yet. All seven were resolved.
- 2026-09-01: Phase 7's fifth strict `plan-checker` cycle verified all seven cycle-4 fixes and reported seven more, with
  one observation worth more than the findings: twice running, the previous cycle had fixed the instance reported rather
  than the class it belonged to, and the next cycle found the same defect one item or one table row away. Two
  `npx tsc --version` acceptances survived in Phases 3 and 5 after the four in Phase 1 were amended; a second plan-only
  proof cell named a repository-wide sweep the tree does not satisfy, one row above the cell cycle 4 fixed; and a second
  acceptance required a `learnings.md` entry that was never written and should not have been, one phase away from the
  first. All three were fixed by sweeping the whole plan for the class rather than by editing the cited line. The
  remaining four were single-site: a README naming one certain `tooling.md` deviation where the register holds two, an
  `[AC-7]` note calling this plan's own README a file outside this plan's documents, an unmarked amendment, and a record
  line enumerating six findings behind a count of seven. All seven were resolved.
- 2026-09-01: Phase 7's sixth strict `plan-checker` cycle verified all seven cycle-5 fixes, confirmed the three classes
  cycle 5 swept were swept clean, and reported six more — one of them a defect in delivered code rather than in the
  plan. A comment in `apps/wahidyankf-www/tests/bdd/accessibility.steps.ts` routed a reader at
  `apps/wahidyankf-www-fe-e2e/steps/accessibility.steps.ts`, the source project's old name, which this repository does
  not have; the item that swept for that name was ticked and its note claimed nothing under `apps/` matched, so the
  acceptance was correct and the reading of it was not. Two more sweeps stated results the tree does not produce: the
  `cv/` sweep's `--hidden` reaches `.git/`, where this phase's own commit message sits in the reflog, and the
  generalizability sweep matches two lines that are correct and stay. The fourth is the one worth naming: cycles 3
  through 5 changed the repository — a script wiring, three prompt edits, a governance relocation — and recorded all of
  it in this record and in `file-impact.md`, but wrote no checkbox for any of it, which is this plan's own convention
  for work added during execution. Three items were added and marked. The last two were prose precision: this record
  still pointed at per-cycle counts the plan README no longer carries, and `tech-docs/README.md`'s root-additions table
  row listed five packages where the prose above it says six. All six were resolved.
- 2026-09-01: Phase 7's seventh strict `plan-checker` cycle verified all six cycle-6 fixes and reported seven findings,
  none above MEDIUM and none that would make an executor do the wrong thing. The two MEDIUMs sat in the three archival
  items that had not run yet, which were the least execution-grade text left in the checklist: the move named no command
  and no guard for the one destructive step remaining, and the index update's only acceptance was
  `check:markdown-links`, which cannot catch `plans/done/README.md` gaining no entry at all — the exact hazard this plan
  pairs with a positive `rg` twice elsewhere. Both now carry commands. The five LOWs were record and prose precision,
  three of them recurrences of classes an earlier cycle had closed one line away. All seven were resolved.
- 2026-09-01: the archival quality gate ended on the seven-cycle bound with every finding of every cycle resolved and
  none waived, which is the second ending [plan quality gate](../../../repo-governance/workflows/plan-quality-gate.md)
  provides for. Across five archival cycles it reported 32 findings: 3 HIGH, 11 MEDIUM, and 18 LOW. Three were defects
  in the delivered repository rather than in the plan — two governance clauses that never reached the `plan-checker`
  prompts the plan asserted they had, a `specs:e2e:baseline` target nothing automated called, and a comment in
  `apps/wahidyankf-www/tests/bdd/accessibility.steps.ts` routing a reader at a project this repository does not have.
  The status is `settled` rather than `pass`, because cycle 7's fixes were applied after the last check and no cycle has
  read them.
- 2026-09-01: Phase 7 gate passed. `test:quick`, `typecheck`, `lint`, `test:integration`, and `test:e2e` all exit 0 —
  258 unit and behaviour tests, 8 integration tests, and 36 E2E scenarios with the 34 skips the tracked baseline
  records. `format:check`, `check:markdown-links`, `check:governance`, `check:workflows`, `check:harness-parity`, and
  `npm audit --audit-level=low` all exit 0, the audit reporting no vulnerability. `git ls-files cv` is silent with
  `test ! -e cv` succeeding beside it, and `ls libs` prints only `README.md`. `typecheck` and `lint` were re-run with
  `--skip-nx-cache`, because a 100% cache hit proves a hash matched rather than that a command passed.
  `check:rule-change` named no workflow, correctly: the phase's rule-path edits were committed earlier today and
  `plans/` is not a rule path.
- 2026-09-01: plan complete and archived to `plans/done/2026-09-01__wahidyankf-www-migration/`. Seven phases, and the
  count of what execution changed about the plan is the honest summary of it: eight items added during execution, eleven
  acceptances amended and marked, three notes corrected after they were found to assert something the tree did not have,
  and five dormant recovery items that never fired and carry dated `Not triggered` dispositions instead of ticks.
  `ose-public` was never modified.

## How to Read This Checklist

Tags: `[AI]` means an agent completes the item, `[HUMAN]` means the owner completes it, and `[AI+HUMAN]` means an agent
prepares it for owner action.

Throughout this checklist, `$SRC` is the executor's own `ose-public` working copy at commit
`e74818fc06c4c104725383384d2aa38305a503ef`. Export it once before the first copy —
`export SRC="<path to your ose-public clone>"` — rather than writing a machine-specific path into this file, which is a
public record. Confirm that commit is checked out before copying anything from it; a copy taken from a different state
invalidates the provenance record in `README.md`.

Scratch output goes to `local-tmp/`, which `.gitignore` already excludes, so a verification artifact never reaches a
commit.

Six items accept on a pinned tool's version — four in Phase 1, and the two that delete a nested `typescript` pin in
Phases 3 and 5 — and each names a command that reads the resolved tree directly rather than `npx <tool> --version`. The
four in Phase 1 use `node node_modules/typescript/bin/tsc --version` or
`node -p "require('./node_modules/<pkg>/package.json').version"`, both rooted at the workspace. The two that delete a
nested `typescript` pin use a third form, `node -p "require('typescript/package.json').version"` run from inside the
project directory, and the difference is the point of those items: rooting the path at the workspace would prove the
root pin exists, where what they have to prove is which TypeScript the project itself resolves. All six were amended
during execution, the Phase 1 four at the fourth archival gate cycle and the other two at the fifth, because the
original form tested the wrong thing: this harness rewrites shell commands through a proxy, and it turns
`npx tsc --version` into a summary of a compile that never ran and never prints a version at all. The acceptance is
about the pinned version, not about which shell wrapper reported it. Phase 7's learnings triage carried the general rule
into [delivery checklists](../../../repo-governance/conventions/plans-organization-policy/delivery-checklists.md).

## Phase 0: Baseline

- [x] [AI] Run `npm install` — acceptance: locked workspace dependencies install and `git diff --stat package-lock.json`
      reports no change.
- [x] [AI] Run `npm run test:quick` — acceptance: the pre-change quick gate exits 0.
- [x] [AI] Run `npm run format:check` — acceptance: the pre-change format gate exits 0.
- [x] [AI] Run `npm run check:markdown-links` — acceptance: the pre-change Markdown-link gate exits 0.
- [x] [AI] Run `git -C "$SRC" rev-parse HEAD` — acceptance: the output is `e74818fc06c4c104725383384d2aa38305a503ef`,
      matching the provenance recorded in `plans/in-progress/wahidyankf-www-migration/README.md`.
- [x] [AI] Run
      `git -C "$SRC" status --porcelain apps/wahidyankf-www apps/wahidyankf-www-fe-e2e specs/apps/wahidyankf specs/libs/ts-env-loader libs/web-ui libs/web-ui-token libs/ts-env-loader scripts/next-with-port.mjs`
      — acceptance: the output is empty **and**
      `for p in apps/wahidyankf-www apps/wahidyankf-www-fe-e2e specs/apps/wahidyankf specs/libs/ts-env-loader libs/web-ui libs/web-ui-token libs/ts-env-loader scripts/next-with-port.mjs; do git -C "$SRC" ls-files --error-unmatch "$p" > /dev/null || echo "MISSING $p"; done`
      prints nothing, proving every source path is both clean and real, including the `ts-env-loader` corpus this plan
      also ports. The second half is not redundant: `git status --porcelain` prints nothing for a pathspec matching no
      tracked file and still exits 0, so a mistyped source path passes the emptiness check by describing nothing. The
      E2E project is `apps/wahidyankf-www-fe-e2e` in `ose-public` and `apps/wahidyankf-www-e2e` here, so only the source
      name belongs in a `$SRC` pathspec.
- [x] [AI] Record the six baseline commands, their exit statuses, and the two source-repository results in
      `plans/in-progress/wahidyankf-www-migration/evidence/phase-0-baseline.md`, and convert that file's entry in the
      `## Directory Map` of `evidence/README.md` to a `[phase-0-baseline.md](phase-0-baseline.md)` relative link in this
      same item — acceptance: each preceding command has one non-secret result entry, with any absolute path rewritten
      as `$SRC`, and the map entry is a link rather than plain text. Every item that writes an evidence file links it
      here, so the map never holds an unlinked entry for a file that already exists beside it. No separate command
      applies, because this action records evidence from the preceding shell output. It goes to `evidence/` rather than
      `learnings.md`: the
      [five-document structure](../../../repo-governance/conventions/plans-organization-policy/five-document-structure.md)
      names `evidence/` for command output, and
      [knowledge capture](../../../repo-governance/conventions/plans-organization-policy/knowledge-capture.md) defines a
      `learnings.md` entry as one paragraph of lesson, which Phase 7 must then route to a durable home. Six exit
      statuses have no durable home to reach.

### Phase 0 Gate

> Every check below passes before Phase 1 begins. If a baseline check fails, stop without repairing the repository and
> record the command, exit status, and observed failure in `learnings.md` for owner direction.

- [x] [AI] Run `git status --short` — acceptance: the output is empty, or every path in it is under
      `plans/in-progress/wahidyankf-www-migration/` or is `plans/in-progress/README.md`; no other path appears. Empty is
      the expected result whenever the [plan-planning](../../../repo-governance/workflows/plan-planning.md) workflow
      already committed the plan at authoring time, which the next-but-one item allows, so an acceptance that demanded
      modified plan paths would fail on the ordinary case. Either way the point is the same: nothing outside the plan
      has changed.
- [x] [AI] Run `npm run test:quick` — acceptance: the existing Badak Mini quick gate exits 0.
- [x] [AI] Commit and push the plan folder and its `plans/in-progress/README.md` entry to `main`, unless the
      [plan-planning](../../../repo-governance/workflows/plan-planning.md) workflow already did so at authoring time —
      acceptance: `git status --short` is empty and
      `git log --oneline origin/main -1 -- plans/in-progress/wahidyankf-www-migration` names a commit. This phase
      commits nothing but the plan itself, so this is the whole of its output.

> **Pause Safety**: The validated plan is the only repository change and it is on `main`, the existing Go application is
> green, and the source repository is confirmed clean at the recorded commit. Nothing has been copied. Safe to stop.
> Resume with `npm run test:quick`.

## Phase 1: TypeScript Workspace Foundation

This phase makes the workspace capable of holding a TypeScript project. No application code lands, so no coverage gate
applies yet.

- [x] [AI] Add `typescript` at an exact version to `devDependencies` in `package.json`, preferring `6` — acceptance,
      **amended during execution**: `node node_modules/typescript/bin/tsc --version` prints the pinned version and
      `rg -n '"typescript": "[\^~]' package.json` finds nothing. [AC-9]
  - Note: pinned `6.0.3`, the latest stable of the 6 line; `7.0.2` is current but
    [tooling](../../../repo-governance/development/testing-policy/tooling.md) names TypeScript 6. `ose-public` pins
    `5.8.3`, so the port raises the compiler two majors, which is what the Phase 3 fallback decision exists to absorb.
    The amended criterion prints `Version 6.0.3`.
- [x] [AI] Add `@biomejs/biome` at an exact version to `devDependencies` in `package.json` — acceptance, **amended
      during execution**: `node -p "require('./node_modules/@biomejs/biome/package.json').version"` prints the pinned
      version. [AC-9]
  - Note: pinned `2.5.11`.
- [x] [AI] Add `tsx` at an exact version to `devDependencies` in root `package.json` — `ose-public` pins `4.21.0`, so
      use that unless `npm install` reports it unresolvable — acceptance, **amended during execution**:
      `node -p "require('./node_modules/tsx/package.json').version"` prints the pinned version and
      `rg -n '"tsx": "[0-9]' package.json` finds an exact pin with no range prefix. The `generate:cv-pdf` target runs
      `npx tsx`, `tsx` appears in no manifest in this repository, and that target is the proof Phase 4 requires before
      `cv/cv-ats.pdf` is deleted, so without this the phase cannot close. If the Phase 3 attempt below shows Node runs
      the script unaided, remove this pin then and record why, as the same policy requires.
  - Note: pinned `4.23.13`, not `4.21.0`. `4.21.0` resolved and installed cleanly, so the item's own escape clause — use
    it "unless `npm install` reports it unresolvable" — never fired; the Phase 1 gate is what caught it. `tsx@4.21.0`
    depends on `esbuild@0.27.7`, which carries GHSA-g7r4-m6w7-qqqr, a low-severity arbitrary file read in esbuild's
    development server on Windows affecting `>=0.27.3 <0.28.1`, and `npm audit --audit-level=low` fails on it. `4.23.13`
    depends on `esbuild@~0.28.0` and resolves `0.28.2`, and the audit is clean. Fixed by moving the pin rather than by
    adding an `overrides` entry or lowering the audit level: an override would leave the vulnerable range reachable for
    anything else that asks for it, and the audit level is what
    [tooling](../../../repo-governance/development/testing-policy/tooling.md) requires. The advisory does not apply to
    this workspace, which is not Windows and does not run esbuild's dev server, but the gate is not scoped by
    exploitability and the fix cost one version bump.
- [x] [AI] Add `eslint` at an exact version to `devDependencies` in root `package.json` — `ose-public` pins `9.39.4` in
      two of its app manifests, so use that unless `npm install` reports it unresolvable — acceptance, **amended during
      execution**: `node -p "require('./node_modules/eslint/package.json').version"` prints the pinned version and
      `rg -n '"eslint": "[0-9]' package.json` finds an exact pin with no range prefix.
      [Testing tooling](../../../repo-governance/development/testing-policy/tooling.md) requires "project-local ESLint
      commentary checks" alongside TypeScript and Biome, and this repository has no ESLint anywhere today, so `[AC-9]`
      cannot be satisfied without it. [AC-9]
  - Note: `9.39.4` resolved and pinned as written.
- [x] [AI] Add `eslint-plugin-jsdoc` at the exact version `npm install` resolves to `devDependencies` in root
      `package.json` — acceptance: `rg -n '"eslint-plugin-jsdoc": "[0-9]' package.json` finds an exact pin with no range
      prefix. The [code commentary policy](../../../repo-governance/development/code-commentary-policy.md) names this
      plugin by name as what enforces complete-sentence summaries for named executable declarations; no version is
      written into this plan because no manifest in either repository pins it today, so the resolved version is recorded
      rather than guessed. [AC-9]
  - Note: resolved to `64.3.2`. Its version was read with
    `node -p "JSON.parse(require('fs').readFileSync('node_modules/eslint-plugin-jsdoc/package.json')).version"`, because
    the package does not expose `./package.json` through its `exports` map and the ordinary
    `require('<pkg>/package.json')` form throws `ERR_PACKAGE_PATH_NOT_EXPORTED`.
- [x] [AI] Create root `biome.json` configuring the linter for TypeScript and React with the JSX accessibility rules
      enabled, mirroring the intent of `$SRC/apps/wahidyankf-www/oxlint.json`, and set `formatter.enabled` and
      `assist.enabled` to `false` so `biome check` is a linter and nothing else — acceptance:
      `npx biome check --max-diagnostics=0 biome.json` exits 0, `rg -n 'jsx-a11y|a11y' biome.json` finds the
      accessibility configuration, and
      `node -p "const c=require('./biome.json'); [c.formatter.enabled, c.assist.enabled].join(',')"` prints
      `false,false`. Prettier stays the formatting source of truth, as the
      [code style policy](../../../repo-governance/development/code-style-policy.md) states, and Biome must not become a
      second one. Biome v2 defaults `formatter.indentStyle` to tab while Prettier here uses two spaces —
      `.prettierrc.json` sets only `proseWrap`, so Prettier's `useTabs: false` and `tabWidth: 2` defaults apply — so an
      enabled Biome formatter would report every ported file as unformatted and fail `lint`, and "fixing" that with
      `biome check --write` would retab those files and break `npm run format:check` repository-wide. Disabling the
      assist actions closes the same hole for import sorting, which Biome would otherwise apply on top of Prettier's
      ordering. [AC-9]
  - Note: written against Biome `2.5.11`. Its JSX accessibility configuration is the `a11y` rule group at
    `preset: "recommended"` plus the `react` and `next` linter domains, which is how Biome 2 spells what `oxlint.json`
    expressed as a `jsx-a11y` plugin entry; the three explicit rules `oxlint.json` names carry over as
    `correctness/noUnusedVariables`, `suspicious/noConsole`, and `suspicious/noDoubleEquals`. The group and preset
    fields are spelled `preset`, not `recommended`: Biome 2.5 accepts `recommended` but emits a deprecation diagnostic
    saying it is removed in the next major, so the first draft was rewritten rather than left to break on an upgrade.
    `--max-diagnostics=0` in the acceptance command is safe only because the file is clean — it reports "the number of
    diagnostics exceeds the limit" and withholds the text of any diagnostic at all, including an informational one, so a
    future reader debugging this file should drop the flag before reading the output.
- [x] [AI] Record in `evidence/phase-1-toolchain.md` whether TypeScript 6, Biome, ESLint, and `eslint-plugin-jsdoc` all
      installed and resolved, naming the exact version of each, and convert that file's entry in the `## Directory Map`
      of `evidence/README.md` to a relative link in this same item — acceptance: one dated entry names all four
      versions, and the map entry is a link that `npm run check:markdown-links` resolves at this phase's gate. This is
      the evidence the Phase 3 fallback decision will be judged against, so it is durable plan evidence rather than a
      `learnings.md` lesson. Add a `learnings.md` paragraph only if a version resolved differently than expected.
  - Note: all four resolved; no version differed from what the plan expected, so no `learnings.md` paragraph was added
    on that count. The evidence file also records how each version was read, because `npx <tool> --version` is
    unreliable under this harness.
- [x] [AI] Copy `$SRC/scripts/next-with-port.mjs` to `scripts/next-with-port.mjs` and change its `resolvePort` import
      from `../libs/ts-env-loader/src/port-resolver.ts` to
      `../apps/wahidyankf-www/src/features/env/core/port-resolver.ts` — acceptance:
      `rg -n 'from "\.\./libs/ts-env-loader' scripts/next-with-port.mjs` finds nothing and
      `rg -n 'apps/wahidyankf-www/src/features/env/core/port-resolver.ts' scripts/next-with-port.mjs` finds the import.
      The pattern is anchored to the `from` clause rather than sweeping the file, because the `CONTAINER REQUIREMENT`
      comment names the same old path in prose and the item below is what removes it; a whole-file sweep here would make
      this item unverifiable until that one ran.
  - Note: the acceptance was a whole-file `rg -n 'libs/ts-env-loader'` when this item was executed, and it failed on the
    comment at line 32 with the import already correct. Anchored to the `from` clause and re-run: clean. This is the
    second acceptance in this plan scoped wider than the action it verifies; the wrapper header item below already
    carries an `awk` range for exactly this reason.
- [x] [AI] Rewrite the header comment of `scripts/next-with-port.mjs` for this repository: drop "the F# backends", the
      "six container images" and "four of the six container images" counts, the `organiclever-www` comparison in the
      `PROCESS SHAPE` paragraph, and the container framing of the `CONTAINER REQUIREMENT` paragraph — its `COPY`
      instruction and the requirement that an image preserve the two files' relative layout — since this repository has
      no F# service, no container image, and one Next.js application. Retain, reworded without the container framing,
      the two sentences that paragraph carries which are still true here: that `port-resolver.ts` is deliberately
      dependency-free — no `dotenv`, no `node:fs` — so it needs no `node_modules`, and that Node strips its TypeScript
      types natively, so there is no build step. Those two are the reason a `.mjs` script can import a `.ts` module at
      all, and they exist nowhere else in this file; deleting the paragraph whole would delete them with it. The
      `CONTAINER REQUIREMENT:` label goes, because the sentences it now introduces are about the resolver rather than
      about an image. Repoint the three-line `Usage:` block in the same pass: its `dev` and `start` examples read
      `--env OSE_WWW_PORT --default 3100`, which is another repository's application, so give both this one's
      `--env WAHIDYANKF_WWW_PORT --default 3201`, matching the `dev` and `start` commands the Target Contract writes
      into `apps/wahidyankf-www/project.json`; and its third example names `--server apps/ose-www/server.js`, so give it
      `apps/wahidyankf-www/server.js` and say in the sentence below it that no application here sets
      `output: "standalone"`, so the `--server` form is supported and unused rather than exercised — acceptance:
      `awk '/^ \*\//{exit} {print}' scripts/next-with-port.mjs | rg -in 'F#|container|organiclever|six|ose-www|OSE_WWW_PORT|3100'`
      finds nothing, and the same header block still shows all three usage forms — the `-i` matters, because the label
      being removed is spelled `CONTAINER REQUIREMENT` in capitals and a case-sensitive pattern would step straight over
      it, and the surviving comment still states the precedence rule, why the wrapper exists at all, and why
      `port-resolver.ts` stays dependency-free, naming `dotenv` and `node:fs`. The `awk` range is the header block
      alone: it prints from the top of the file and stops at the first ` */` line, which is where that comment ends.
      Scoping matters here, because two later inline comments in the signal-forwarding block also say `container` — they
      explain why SIGTERM is forwarded and re-raised so an orchestrator does not read a stop as a crash — and this item
      does not authorize touching them. A whole-file grep could only be satisfied by deleting reasoning no item asked to
      remove.
  - Note: the `CONTAINER REQUIREMENT` label became `RESOLVER`, which is what the surviving sentences are about. Both
    retained sentences are present, now naming `apps/wahidyankf-www/src/features/env/core/port-resolver.ts` as the
    layout the wrapper depends on, and the header states outright that no application here sets `output: "standalone"`,
    so the `--server` form is supported and unused. The two `container` mentions in the signal-forwarding block are
    untouched, as the item requires; a whole-file grep for `container` still returns exactly those two.
    `node --check scripts/next-with-port.mjs` parses.
- [x] [AI] Delete `scripts/.gitkeep` — acceptance: `test ! -e scripts/.gitkeep` succeeds, because the directory now
      holds a real file.
- [x] [AI] Add a `## Directory Map` section to `scripts/README.md` listing `next-with-port.mjs` with a one-line
      description of the port contract — the file has no such section today, so it is created rather than appended to —
      acceptance: `rg -n '^## Directory Map' scripts/README.md` finds the heading and
      `rg -n 'next-with-port.mjs' scripts/README.md` finds the entry.
- [x] [AI] Delete the closing sentence of the second paragraph of `scripts/README.md`, which states that the `.gitkeep`
      file preserves the directory while it has no scripts — that becomes false once `next-with-port.mjs` lands and
      `.gitkeep` is removed — acceptance: `rg -n 'gitkeep' scripts/README.md` finds nothing.
- [x] [AI] Edit `.gitignore` to ignore six generated paths at the repository root — `*.tsbuildinfo`, `.next/`,
      `coverage/`, `test-results/`, `playwright-report/`, and `.features-gen/` — acceptance:
      `git check-ignore -v --no-index apps/wahidyankf-www/tsconfig.tsbuildinfo apps/wahidyankf-www/.next/build-manifest.json apps/wahidyankf-www/coverage/index.html apps/wahidyankf-www-e2e/test-results/failure.png apps/wahidyankf-www-e2e/playwright-report/index.html apps/wahidyankf-www-e2e/.features-gen/home.feature.spec.js`
      exits 0 and prints six lines, each naming a rule in the root `.gitignore`. `--no-index` makes this observable now,
      before any of those paths exists. All six rules go at the root rather than only the first, because two gates read
      this file by different rules. Git honours a nested `.gitignore`, so for `git status --short` the copied
      `apps/wahidyankf-www/.gitignore` and `apps/wahidyankf-www-e2e/.gitignore` would indeed be enough for the five
      directories. Prettier does not: `npm run format:check` is `prettier --check .` run from the workspace root, and
      Prettier resolves ignore patterns from the root `.gitignore` and a root `.prettierignore` alone, never from a
      nested `.gitignore`. A cycle-3 review removed these five rules on exactly the Git reasoning, and this is why they
      are back — that reasoning was true for the gate it named and silent about `format:check`. The invariant that has
      kept `format:check` green so far is that every generated path in this repository is already matched by the root
      `.gitignore`; this plan is the first to introduce generated output that only a nested file would cover, and
      without these rules the Phase 3 gate's `format:check` fails on the thousands of files `static-routes:validation`
      leaves in `apps/wahidyankf-www/.next`. `*.tsbuildinfo` is load-bearing for Git as well as for Prettier: the ported
      `tsconfig.json` sets both `incremental` and `noEmit`, so the first `typecheck` writes
      `apps/wahidyankf-www/tsconfig.tsbuildinfo`; this repository's root `.gitignore` has no such rule today and the
      ported app `.gitignore` does not add one, so without it every later "`git status --short` is empty" gate would
      fail on a build cache. The Phase 3 and Phase 5 gates, which require both `git status --short` to be empty and
      `format:check` to exit 0 after a build and after a Playwright run, are where all six are proved rather than
      assumed.
  - Note: the six rules carry a comment above them stating why they are at the root rather than only in each project, so
    a later reader does not move them down into the app and silently break `format:check`.
    `git check-ignore -v --no-index` printed six lines, each naming a rule in the root `.gitignore`, before any of those
    paths exists.
- [x] [AI] Reformat the copied wrapper to this repository's Prettier configuration — run `npm run format` — acceptance:
      `npm run format:check` exits 0. `$SRC/.prettierrc.json` sets `printWidth: 120`, sets `proseWrap: "preserve"`, and
      loads `prettier-plugin-tailwindcss`; this repository's `.prettierrc.json` sets `proseWrap: "never"` and nothing
      else, leaving Prettier's `printWidth: 80` default in force, so `scripts/next-with-port.mjs` fails
      `prettier --check` the moment it lands. Phase 0 proved `format:check` exits 0 before anything was copied, so this
      plan is what creates that failure and this is the named step that closes it. `npm run format` is
      `prettier --write .`, so it sweeps the whole tree; the copied wrapper is the only file in it that arrives
      pre-formatted to another repository's configuration, and the same run normalizes anything this phase authored by
      hand, such as the new root `biome.json` and the `scripts/README.md` edits. The item sits last in the phase
      deliberately: after every edit to the copied file, so nothing is reformatted twice, and before the gate, so no
      gate check depends on a step that has not run. The diff is a formatting change and not a content change — it moves
      line breaks and quoting and alters no identifier, no string, and no comment sentence — and
      [technical design](tech-docs/README.md#toolchain-conformance-and-its-fallback) records the three configuration
      differences so a later reader comparing a ported file against its `ose-public` original expects it. No ignore
      entry is added instead: [file impact](tech-docs/file-impact.md) rules out introducing a root `.prettierignore` at
      all, and exempting a ported path from this repository's formatting source of truth is the outcome this item exists
      to prevent.
  - Note: the prediction held exactly. Before the run, `format:check` named one file and only one —
    `scripts/next-with-port.mjs` — and nothing this phase authored by hand, because `biome.json`, the
    `scripts/README.md` edits, and the plan documents were each formatted as they were written. After the run
    `format:check` exits 0, the file still parses under `node --check`, and the header sweep is still clean.

### Phase 1 Gate

> Every check below passes before Phase 2 begins. A failure is fixed inside Phase 1.

- [x] [AI] Run `npm run test:quick` — acceptance: exits 0; adding root dev dependencies did not disturb the Go project.
- [x] [AI] Run `npm run format:check` — acceptance: exits 0.
- [x] [AI] Run `npm run check:markdown-links` — acceptance: exits 0, proving the new `scripts/README.md` entry resolves.
- [x] [AI] Run `npm audit --audit-level=low` — acceptance: exits 0 with the five new dev dependencies — `typescript`,
      `@biomejs/biome`, `tsx`, `eslint`, and `eslint-plugin-jsdoc` — in the tree.
- [x] [AI] Commit and push the phase to `main` — acceptance: `git status --short` is empty and `git log -1 --format=%s`
      names the workspace foundation.

> **Pause Safety**: The workspace can now hold a TypeScript project and owns the port wrapper, but no application exists
> and no target references it. `scripts/next-with-port.mjs` imports a path that does not exist yet, which is inert
> because nothing invokes it. Safe to stop. Resume with `npm run test:quick`.

## Phase 2: Canonical Corpus and Architecture Model

Gherkin lands before the code that implements it, as the
[specs policy](../../../repo-governance/development/specs-policy.md) requires.

Nothing in this phase runs a RED cycle. Every scenario here already exists and already passes in `ose-public` against
code this plan copies rather than writes, so there is no failing-first state to reach.
[Specification changes](tech-docs/specification-changes.md#why-the-ported-scenarios-have-no-red-cycle) states that
exemption and its boundary: the only scenarios in this plan that are genuinely new are the two CV export scenarios in
Phase 3, and those do run full RED, GREEN, and REFACTOR cycles.

- [x] [AI] Create `specs/apps/wahidyankf-www/behaviours/` and copy all nine `.feature` files from
      `$SRC/specs/apps/wahidyankf/behaviours/wahidyankf-www/gherkin/` into it, flattening the `app-shell/`, `cv/`,
      `env-loader/`, `home/`, `personal-projects/`, and `search/` subdirectories away — acceptance:
      `ls specs/apps/wahidyankf-www/behaviours/*.feature | wc -l` prints `9` and no subdirectory exists under
      `behaviours/`.
  - Note: the source directory also holds nine `.feature` files, so the flatten collapsed six subdirectories without a
    single basename collision. Counting both sides matters here: a collision would have overwritten one file and still
    left `ls | wc -l` short of nine, which the acceptance would have caught, but a collision plus a spare file would
    not.
- [x] [AI] Copy `$SRC/specs/libs/ts-env-loader/behaviours/gherkin/env-loader/env-loader.feature` to
      `specs/apps/wahidyankf-www/behaviours/tier-env-loading.feature`, carrying its five scenarios — "Loads the selected
      tier's file", "Process env always wins over the tier file", "Tolerates a missing tier file", the Scenario Outline
      "Fails loudly on a stray auto-loaded env file" with its Examples table, and "Tolerates a stray file at the local
      tier" — acceptance:
      `rg -c '^\s*(Scenario|Scenario Outline):' specs/apps/wahidyankf-www/behaviours/tier-env-loading.feature` prints
      `5`. Inlining `ts-env-loader` brings its behaviour into this application, so its corpus comes with it.
- [x] [AI] Change the `Feature:` line of `specs/apps/wahidyankf-www/behaviours/tier-env-loading.feature` from
      `APP_ENV tier env-file loading` to `Tier env-file loader module`, leaving every scenario byte-identical —
      acceptance:
      `rg -n '^Feature: Tier env-file loader module' specs/apps/wahidyankf-www/behaviours/tier-env-loading.feature`
      finds the new title and
      `rg -c '^Feature: APP_ENV tier env-file loading' specs/apps/wahidyankf-www/behaviours/*.feature` reports the count
      only for `env-loader.feature`. The rename is required because the app corpus already carries a feature with the
      original title; two features sharing a title in one directory is ambiguous to a reader and to a binding.
  - Note: the collision Decision A anticipated is real and was confirmed before the rename — `env-loader.feature` and
    the copied `tier-env-loading.feature` both opened `Feature: APP_ENV tier env-file loading`. After the rename all
    eleven files carry eleven distinct `Feature:` titles, and the scenario count of the renamed file is unchanged.
- [x] [AI] Copy `$SRC/specs/libs/ts-env-loader/behaviours/gherkin/port-resolver/port-resolver.feature` to
      `specs/apps/wahidyankf-www/behaviours/port-resolver.feature` unchanged, carrying its eight scenarios — "The CLI
      flag outranks every other source", "The prefixed variable outranks the fallback", "The fallback applies when
      nothing else supplies a port", "A bare PORT variable never moves the listener", the Scenario Outline "A blank
      value at a tier falls through to the next tier", the Scenario Outline "A present but malformed port fails loudly
      instead of falling through", "A malformed prefixed variable names that variable in the error", and "An
      out-of-range compiled-in fallback is caught at startup" — acceptance:
      `diff "$SRC/specs/libs/ts-env-loader/behaviours/gherkin/port-resolver/port-resolver.feature" specs/apps/wahidyankf-www/behaviours/port-resolver.feature`
      reports no difference.
- [x] [AI] Confirm the corpus is complete — acceptance: `ls specs/apps/wahidyankf-www/behaviours/*.feature | wc -l`
      prints `11` and
      `grep -rhE '^[[:space:]]*(Scenario|Scenario Outline):' specs/apps/wahidyankf-www/behaviours/*.feature | wc -l`
      prints `53`.
- [x] [AI] Verify the nine app feature files moved without loss by comparing scenario titles — acceptance:
      `diff <(grep -rhE '^\s*(Scenario|Scenario Outline):' "$SRC/specs/apps/wahidyankf/behaviours/wahidyankf-www/gherkin" | sort) <(grep -rhE '^\s*(Scenario|Scenario Outline):' specs/apps/wahidyankf-www/behaviours/accessibility.feature specs/apps/wahidyankf-www/behaviours/cv.feature specs/apps/wahidyankf-www/behaviours/env-loader.feature specs/apps/wahidyankf-www/behaviours/home.feature specs/apps/wahidyankf-www/behaviours/personal-projects.feature specs/apps/wahidyankf-www/behaviours/responsive.feature specs/apps/wahidyankf-www/behaviours/search.feature specs/apps/wahidyankf-www/behaviours/static-filterable-routes.feature specs/apps/wahidyankf-www/behaviours/theme.feature | sort)`
      reports no difference, and each side prints 40 lines. [AC-4]
- [x] [AI] Verify the two ported loader feature files moved without loss by comparing scenario titles — acceptance:
      `diff <(grep -rhE '^\s*(Scenario|Scenario Outline):' "$SRC/specs/libs/ts-env-loader/behaviours/gherkin" | sort) <(grep -rhE '^\s*(Scenario|Scenario Outline):' specs/apps/wahidyankf-www/behaviours/tier-env-loading.feature specs/apps/wahidyankf-www/behaviours/port-resolver.feature | sort)`
      reports no difference, and each side prints 13 lines. Only the `Feature:` title changed, so no scenario title may
      differ. [AC-4]
- [x] [AI] Produce the step-cardinality report over the copied corpus by running
      `mkdir -p local-tmp && awk 'FNR==1 && name!="" {printf "%s | Given=%d When=%d Then=%d\n", name, g, w, t; name=""} /^[[:space:]]*(Scenario|Scenario Outline):/ {if (name != "") printf "%s | Given=%d When=%d Then=%d\n", name, g, w, t; name=FILENAME": "$0; g=0; w=0; t=0} /^[[:space:]]*Given /{g++} /^[[:space:]]*When /{w++} /^[[:space:]]*Then /{t++} END {if (name != "") printf "%s | Given=%d When=%d Then=%d\n", name, g, w, t}' specs/apps/wahidyankf-www/behaviours/*.feature > local-tmp/step-cardinality.txt`
      — acceptance: `wc -l < local-tmp/step-cardinality.txt` prints `53` and
      `grep -cvE 'Given=[01] When=1 Then=1' local-tmp/step-cardinality.txt` prints `0`. The pattern names `Given`
      explicitly: `grep -cv 'When=1 Then=1'` matches a line reading `Given=2 When=1 Then=1` and so counts it as
      conforming, which would let a two-primary-`Given` scenario through the check the next item says flags it. No
      scenario in the source corpus has more than one `Given` today, so this is a latent gap rather than a live one —
      which is exactly why it has to be closed here, before a later feature edit makes it live. `Given=0` is allowed in
      the pattern because `Background` and `Examples` are exempt from the cardinality rule, per the
      [specs policy](../../../repo-governance/development/specs-policy.md); the next item is what proves a `Given=0`
      scenario really does sit in a file with a `Background`.
- [x] [AI] Confirm every `Given=0` line in the report belongs to a file that carries a `Background` supplying the Given
      — acceptance: the file list from `grep 'Given=0' local-tmp/step-cardinality.txt | sed 's/:.*//' | sort -u` is a
      subset of the file list from
      `rg -l '^[[:space:]]*Background:' specs/apps/wahidyankf-www/behaviours/*.feature | sort`, the comparison is
      recorded in `evidence/phase-2-background-coverage.md`, and that file's entry in the `## Directory Map` of
      `evidence/README.md` is converted to a relative link in this same item, which this phase's gate then resolves.
  - Note: the two lists are equal, not merely in a subset relation — the same eight files carry every `Given=0` scenario
    and every `Background`. `comm -23` over the sorted lists printed nothing.
- [ ] [AI] Split any scenario the two preceding checks flag — more than one primary `Given`, `When`, or `Then`, or a
      `Given=0` in a file with no `Background` — and record each split in `learnings.md`, because a split means the
      source corpus disagreed with this repository's rule — acceptance: rerunning the report leaves nothing flagged, and
      `learnings.md` names each split scenario and its file. Trigger: either preceding check reports a flagged line.
      Phase 7 gives this item a dated `Not triggered` disposition when nothing was flagged.
  - **Not triggered, 2026-09-01.** Neither preceding check flagged anything — `grep -cvE 'Given=[01] When=1 Then=1'`
    over all 53 report lines printed `0`, and the `Given=0` and `Background` file lists are equal, which is what
    [evidence](evidence/phase-2-background-coverage.md) records. The source corpus already agrees with this repository's
    cardinality rule, so no scenario was split and `learnings.md` names none. The box is deliberately left unticked: a
    tick would claim work that did not happen, and a dated disposition is what the
    [delivery checklist rules](../../../repo-governance/conventions/plans-organization-policy/delivery-checklists.md)
    require of a dormant item instead.
- [x] [AI] Create `specs/apps/wahidyankf-www/behaviours/README.md` with a `## Directory Map` linking all eleven feature
      files, then run `git add -N specs/apps/wahidyankf-www/behaviours/README.md` before checking it — acceptance:
      `npm run check:markdown-links` exits 0 and the map lists eleven entries. The check takes its document list from
      `git ls-files`, as
      [workspace commands](../../../repo-governance/development/workspace-commands.md#repository-checks) states, so
      without the intent-to-add it would skip this file entirely and exit 0 whether or not a single link resolved.
  - Note: executed **after** the `architecture.md` item below rather than before it. As authored, this item's acceptance
    was unreachable in checklist order: the behaviour README links to `../architecture.md`, which the next item creates,
    so `npm run check:markdown-links` failed with `"../architecture.md" targets a file that does not exist`. The eleven
    feature links resolved on the first run — the checker resolves against the filesystem, so the intent-add matters for
    the README being _read_, not for its targets being _found_. Swapping the two items is enough; nothing else depends
    on the order. A future edit of this plan should reorder them rather than repeat the discovery.
- [x] [AI] Author `specs/apps/wahidyankf-www/architecture.md` by consolidating
      `$SRC/specs/apps/wahidyankf/product/overview.md`, `system-context/context.md`, `containers/container.md`, and
      `components/web/component-web.md` into one model that describes the post-migration state: no Docker container
      node, no external design-system component, and one CV store — acceptance: the file contains a system-context view
      and a container view as fenced `text` ASCII blocks, contains no Mermaid, links to `behaviours/`, and
      `rg -n 'docker|web-ui|open-sharia' specs/apps/wahidyankf-www/architecture.md` finds nothing. The model is written
      here as an as-built description of a system Phases 3, 4, and 5 build, so authoring it is not the whole of it: the
      Phase 5 reconciliation item checks each of its three stated differences from the source against the delivered tree
      and corrects the model where they disagree.
  - Note: the source component document describes bounded contexts under `src/contexts/`, which is not where they live —
    the tree at the recorded commit has `src/features/`. The as-built model states `src/features/`, and adds the two
    components the migration inlines, `ui/shell` and `env/core`, which no source document mentions because they were
    external packages there. Three views authored as fenced `text` ASCII: system context, containers, and components.
    The container view draws the E2E project, because starting a server and driving it over a local port is a real
    process boundary and is what makes it a different toolchain from the in-process behaviour adapter.
- [x] [AI] Create `specs/apps/wahidyankf-www/README.md` naming the corpus path, the three required adapters, and the Nx
      targets, with a `## Directory Map` linking `architecture.md` and `behaviours/README.md`, then run
      `git add -N specs/apps/wahidyankf-www/README.md specs/apps/wahidyankf-www/architecture.md` before checking —
      acceptance: `npm run check:markdown-links` exits 0 and reports on both files. `architecture.md` is intent-added
      too, because it is new and carries links of its own.
- [x] [AI] Edit `specs/apps/README.md` to index the new corpus — acceptance:
      `rg -c 'wahidyankf-www' specs/apps/README.md` prints a non-zero count.
- [x] [AI] Edit `specs/README.md` to index the new corpus and repair the one sentence this corpus makes false. Its
      `## Current Specifications` section opens with a link to Badak Mini followed by "is the only subject so far";
      rewrite it to name both subjects — Badak Mini's five-feature corpus with its unit, local integration, and
      public-process E2E adapters, and `wahidyankf-www`'s eleven-feature corpus with its unit, behaviour, integration,
      and process E2E adapters — keeping the closing claim that the adapters fail when a feature, step, binding, or
      adapter drifts, which stays true of both — acceptance: `rg -c 'wahidyankf-www' specs/README.md` prints a non-zero
      count, `rg -n 'only subject' specs/README.md` finds nothing, and `npm run check:markdown-links` exits 0 on the
      links the rewritten sentence carries. The stale sentence is named here for the same reason the `scripts/README.md`
      and `port-resolver.unit.test.ts` items name theirs: a count-only acceptance passes on an added line while the
      contradicting line beside it survives. `specs/apps/README.md`, edited by the item above, needs no equivalent
      repair — it is a Directory Map with no claim about how many subjects exist.

### Phase 2 Gate

> Every check below passes before Phase 3 begins. A failure is fixed inside Phase 2.

- [x] [AI] Run `npm run check:markdown-links` — acceptance: exits 0.
- [x] [AI] Run `npm run format:check` — acceptance: exits 0.
- [x] [AI] Run `npm run test:quick` — acceptance: exits 0; no project consumes the new corpus yet, so the Go gate is
      unaffected.
- [x] [AI] Commit and push the phase to `main` — acceptance: `git status --short` is empty.

> **Pause Safety**: The canonical corpus and the as-built model exist and are indexed, but no adapter binds them, so 53
> scenarios are currently documentation rather than executable specification. That is the one state this repository
> normally forbids, and it is closed by the end of Phase 3. Safe to stop. Resume with `npm run check:markdown-links`.

## Phase 3: Port the Application to a Green 99% State

This phase is large by owner direction: no gate may close while coverage is below 99% or a required behaviour layer is
missing, so the application arrives complete or not at all. The checklist is ordered so each area reaches green before
the next begins.

### Scaffold

- [x] [AI] Copy
      `$SRC/apps/wahidyankf-www/{package.json,tsconfig.json,next.config.ts,postcss.config.mjs,vitest.config.ts,next-env.d.ts,LICENSE,README.md,.env.example,.gitignore}`
      to `apps/wahidyankf-www/` — acceptance: all ten files exist and `git status --short apps/wahidyankf-www` lists
      them as untracked.
- [x] [AI] Do not copy `$SRC/apps/wahidyankf-www/Dockerfile`, `.dockerignore`, or `oxlint.json` — acceptance:
      `test ! -e apps/wahidyankf-www/Dockerfile && test ! -e apps/wahidyankf-www/.dockerignore && test ! -e apps/wahidyankf-www/oxlint.json`
      succeeds.
- [x] [AI] Pin the five caret-ranged specifiers in `apps/wahidyankf-www/package.json` — `class-variance-authority`,
      `clsx`, `react-icons`, `tailwind-merge`, and `@vitejs/plugin-react` — to the exact versions `npm install` resolves
      — acceptance: `rg -n '"\^|"~' apps/wahidyankf-www/package.json` finds nothing.
  - Note: resolved to `class-variance-authority` `0.7.1`, `clsx` `2.1.1`, `react-icons` `5.7.0`, `tailwind-merge`
    `2.6.1`, and `@vitejs/plugin-react` `5.2.0`. Resolved with `npm view <name>@<range> version` rather than from a
    lockfile, because the app was not yet an npm workspace at this point in the phase. A sweep of the whole manifest
    afterwards finds no remaining `^` or `~` on any specifier, in either section.
- [x] [AI] Delete the `"typescript": "5.8.3"` entry from `apps/wahidyankf-www/package.json` `devDependencies`, so the
      root pin from Phase 1 is the only TypeScript in the tree — acceptance, **amended during execution**:
      `rg -n '"typescript"' apps/wahidyankf-www/package.json` finds nothing,
      `test ! -e apps/wahidyankf-www/node_modules/typescript` succeeds, and
      `node -p "require('typescript/package.json').version"` run from inside `apps/wahidyankf-www` prints the root pin.
      The criterion first read `npx tsc --version`, the form the four Phase 1 version items were amended away from and
      the form
      [delivery checklists](../../../repo-governance/conventions/plans-organization-policy/delivery-checklists.md) now
      forbids: this harness answers it with a compile summary carrying no version at all. Absence of a nested directory
      and the version Node actually resolves are the two facts the item is about, so the criterion names both. Under npm
      workspaces a nested pin resolves ahead of the root one, so leaving it would let `[AC-9]` pass while the
      application still compiled on 5.8.3. [AC-9]
  - Note, 2026-09-01: the entry is gone, no `node_modules/typescript` exists inside the project, and the amended read
    prints `6.0.3`, resolving to the root `node_modules/typescript/package.json`. [AC-9]
- [x] [AI] Remove the two `@open-sharia-enterprise/*` entries from `apps/wahidyankf-www/package.json` dependencies —
      `@open-sharia-enterprise/ts-env-loader` and `@open-sharia-enterprise/web-ui` — acceptance:
      `rg -n 'open-sharia-enterprise' apps/wahidyankf-www/package.json` finds nothing. `web-ui-token` is not a
      dependency entry; it reaches the app only through two CSS imports. [AC-6]
- [x] [AI] Remove the two `@open-sharia-enterprise/*` `paths` entries from `apps/wahidyankf-www/tsconfig.json` and point
      its `extends` at the workspace `tsconfig.base.json` — acceptance:
      `rg -n 'open-sharia-enterprise' apps/wahidyankf-www/tsconfig.json` finds nothing and
      `rg -n '"extends": "../../tsconfig.base.json"' apps/wahidyankf-www/tsconfig.json` finds the inheritance.
- [x] [AI] Set `compilerOptions.types` in `apps/wahidyankf-www/tsconfig.json` to the type packages this application
      needs, because `tsconfig.base.json` sets `"types": ["node"]` and an inherited `types` array replaces rather than
      extends it — acceptance: `rg -n '"types"' apps/wahidyankf-www/tsconfig.json` finds the array. The acceptance stops
      at the file's content on purpose: `npx nx run wahidyankf-www:typecheck` cannot pass here, because
      `apps/wahidyankf-www/project.json` is authored further down this section, `npm install` runs after that, and no
      application source lands until the Inline the Three Libraries section, so there is no target to invoke and nothing
      for it to check. The `typecheck` proof for this project belongs to the `npx nx run wahidyankf-www:typecheck` item
      later in this phase, which runs once all three exist and which the version-fallback item beside it backs up.
      `tsconfig.base.json` itself is deliberately not edited: every Next-specific option — `module`, `moduleResolution`,
      `target`, `lib`, `jsx`, `plugins`, and `types` — is overridden per project. If one of them proves impossible to
      override per project, stop and amend the plan rather than editing the shared base by improvisation.
  - Note: set to `["node", "vitest/globals", "@testing-library/jest-dom"]`. `node` is what the base contributes and has
    to be restated because the inherited array is replaced rather than merged; the other two are what the ported test
    suite needs. Nothing was edited in `tsconfig.base.json`, and no Next-specific option proved impossible to override
    per project, so the stop-and-amend branch of this item did not fire. `extends` is written as the first key so a
    reader meets the inheritance before the overrides.
- [x] [AI] Remove the two `@open-sharia-enterprise/*` entries from the `transpilePackages` array in
      `apps/wahidyankf-www/next.config.ts` — acceptance:
      `rg -n 'open-sharia-enterprise' apps/wahidyankf-www/next.config.ts` finds nothing.
- [x] [AI] Delete the three source exclusions `src/app/layout.tsx`, `src/app/head.tsx`, and
      `src/features/cv/core/data.ts` from `coverage.exclude` in `apps/wahidyankf-www/vitest.config.ts` — acceptance:
      `rg -n 'layout.tsx|head.tsx|core/data.ts' apps/wahidyankf-www/vitest.config.ts` finds nothing. All three are
      runtime code, and the [testing policy](../../../repo-governance/development/testing-policy.md) forbids omitting
      runtime code from the denominator. The remaining exclusions — `src/app/fonts/**`, `src/app/**/*.css`,
      `src/test/**`, `**/*.config.*`, `**/.next/**`, `**/dist/**`, `**/coverage/**` — are fixtures, assets, and build
      output and stay. [AC-2]
- [x] [AI] Delete the whole `coverage.thresholds` block from `apps/wahidyankf-www/vitest.config.ts` — the `lines: 80`,
      `functions: 80`, `branches: 75`, `statements: 80` metadata — acceptance:
      `rg -n 'thresholds' apps/wahidyankf-www/vitest.config.ts` finds nothing. The threshold this repository enforces is
      the `--coverage.thresholds.lines=99` flag on the `test:coverage:unit` and `test:coverage:integration` targets, and
      the testing policy forbids duplicating an executable threshold as metadata, where the two can silently disagree.
      [AC-2]
- [x] [AI] Replace the two ported vitest projects `unit-fe` and `integration` in `apps/wahidyankf-www/vitest.config.ts`
      with three named `unit`, `integration`, and `behaviour`, where `unit` includes only `src/**/*.unit.test.{ts,tsx}`
      under `jsdom`, `behaviour` includes only `tests/bdd/**` under `jsdom`, and `integration` includes only
      `tests/integration/**` under `node`, carrying forward the two attributes the ported `unit-fe` project holds beyond
      its name, environment, and include list: `plugins: sharedPlugins` goes on all three, and
      `setupFiles: ["./src/test/setup.ts"]` goes on `unit` and `behaviour` only — acceptance:
      `(cd apps/wahidyankf-www && npx vitest --config vitest.config.ts list --project unit --project integration --project behaviour)`
      runs without an unknown-project error, `rg -c 'plugins: sharedPlugins' apps/wahidyankf-www/vitest.config.ts`
      reports four — the root entry plus one per project — and `rg -c 'setupFiles' apps/wahidyankf-www/vitest.config.ts`
      reports two, on `unit` and `behaviour`. Neither attribute is optional. `sharedPlugins` is
      `[react(), tsconfigPaths()]`: `@vitejs/plugin-react` is what compiles the JSX every component test and step file
      renders, and `vite-tsconfig-paths` is what resolves the `@/features/…` specifiers they import, so a project
      without it fails to resolve its own subject. `integration` needs `sharedPlugins` for the same path resolution —
      the CV export test imports the export module by its `@/` path — but not the setup file, whose
      `@testing-library/jest-dom/vitest` matchers and `afterEach(cleanup)` have nothing to act on in a `node`
      environment. `unit` and `behaviour` do need it: the ported step files carry a header comment saying that
      `@amiceli/vitest-cucumber` registers every `Given`/`When`/`Then` as its own Vitest test and that
      `src/test/setup.ts` runs `cleanup()` after every test, so a `render()` done in a `When` step does not survive into
      the following `Then` — which is why each assertion step re-renders. Without the setup file those repeated renders
      accumulate in one document and the `screen` queries start matching several elements, and the `jest-dom` matchers
      those files call are not registered at all. Neither failure names the config that caused it. The ported config
      names the first project `unit-fe` and points it at `test/**/*.steps.{ts,tsx}`, neither of which matches the target
      contract's `--project unit` and `--project behaviour`. Every direct `npx vitest` invocation in this checklist runs
      from `apps/wahidyankf-www` in a subshell, with the config path written relative to that directory. Vitest's `root`
      defaults to `process.cwd()` rather than to the directory holding the config file, so the same command issued from
      the workspace root resolves `src/**`, `tests/bdd/**`, and `tests/integration/**` against the workspace root and
      matches nothing at all. The ported config sets `passWithNoTests: true` and no item here removes it, so a run that
      matched nothing still exits 0 — the miss is silent, which is why the working directory is written into each
      command rather than left to the reader. The `test:unit`, `test:integration`, `test:coverage:unit`,
      `test:coverage:integration`, and `test:coverage:behaviour` targets have the same requirement met for them by the
      `cwd: {projectRoot}` the Target Contract in [technical design](tech-docs/README.md#target-contract) gives every
      single-command target.
  - Note: the `rg -c 'setupFiles'` half of the acceptance is sensitive to prose. A first draft explained the integration
    project's omission with the words "No `setupFiles` here", which made the count three and failed a check that was
    measuring the configuration correctly. The comment now says "No setup file here" and the count is two. Worth knowing
    before writing another comment near a counted token: an acceptance that greps a source file counts what the comments
    say as readily as what the code does.
- [x] [AI] Set the coverage denominator explicitly in `apps/wahidyankf-www/vitest.config.ts` with
      `coverage.include: ["src/**"]` — acceptance:
      `rg -n 'include: \["src/\*\*"\]' apps/wahidyankf-www/vitest.config.ts` finds it, and a later
      `npx nx run wahidyankf-www:test:coverage:unit` run lists every file under `src/` in its report, including files no
      test imports. Vitest 4 reports only the files some test covered unless an include widens it, so without this an
      untested module is absent from the denominator instead of counting against it, which turns a 99% gate into a
      measure of what happened to be imported. No `coverage.all` accompanies it: that is not a Vitest 4 option —
      `BaseCoverageOptions` at the pinned `4.1.0` has no such key — and the application's `typecheck` covers
      `vitest.config.ts`, so the excess property would fail `tsc --noEmit` with TS2769. [AC-2]
- [x] [AI] Create `apps/wahidyankf-www/biome.json` extending the root configuration — acceptance:
      `npx biome check apps/wahidyankf-www/biome.json` exits 0.
  - Note: Biome 2 spells nested inheritance as `"root": false` with `"extends": "//"`, and both halves are required.
    `"extends": ["//"]` alone fails with `Could not resolve //: found directory without index`, and
    `"extends": ["../../biome.json"]` fails with
    `Found a nested root configuration, but there's already a root configuration` — a relative path does not tell Biome
    this config is subordinate. The file carries nothing but the inheritance: a `files.includes` restriction was tried
    and removed, because scoping to `src/**` and `tests/**` made the acceptance command find no file to process and exit
    1, and Biome already scopes a run by the path it is given.
- [x] [AI] Create `apps/wahidyankf-www/eslint.config.mjs` holding only the commentary check — `eslint-plugin-jsdoc`
      configured to require a complete-sentence summary on every named executable declaration under `src/`, with no rule
      that duplicates a Biome rule — acceptance:
      `(cd apps/wahidyankf-www && npx eslint --config eslint.config.mjs --print-config src/env.ts)` exits 0 and its
      output names a `jsdoc/` rule. The working directory is part of the check, not incidental: a flat config's `files`
      globs resolve against ESLint's base path, which is the current directory rather than the config file's directory,
      so the same command run from the workspace root exits 0 with an empty rule set — the `src/**/*.ts` pattern matches
      nothing from there. `{projectRoot}` is where the `lint:commentary` target runs it, so this is the invocation that
      reflects production.
  - Note: the acceptance originally ran from the workspace root and passed on exit status while printing no `jsdoc/`
    rule at all — the failure mode the reworded command now rules out. Run from `{projectRoot}` it names all three:
    `jsdoc/require-jsdoc`, `jsdoc/require-description`, and `jsdoc/require-description-complete-sentence`. Arrow
    functions and function expressions are deliberately exempt from `require-jsdoc`, because a callback's purpose
    belongs at the call site that reads it. The
    [code commentary policy](../../../repo-governance/development/code-commentary-policy.md) makes this check the
    TypeScript half of what golangci-lint's Revive does for `badakmini-cli`, and it is deliberately narrow: Biome owns
    style and correctness, this file owns commentary alone, so a finding is never reported twice.
- [x] [AI] Author `apps/wahidyankf-www/project.json` with the target contract in
      [technical design](tech-docs/README.md), using the `command` shorthand for every single-command target and
      `nx:run-commands` with `options.commands` and `parallel` set to `false` only for `lint`, `test:quick`, and
      `test:coverage`, and giving each target the `cwd` that table names: `{projectRoot}` on every single-command
      target, `{workspaceRoot}` on `static-routes:validation` alone, and none on the three aggregates — acceptance:
      `npx nx show project wahidyankf-www --json` lists every target in that table with the `cwd` the table states, and
      `rg -n 'rhino-cli|dotnet|\.fsproj|test:specs|specs:structure-validation|specs:behaviour:coverage' apps/wahidyankf-www/project.json`
      finds nothing. The two working directories are not interchangeable: `next build`, `tsc --noEmit`, `biome check`,
      and the `../../scripts/next-with-port.mjs` path only resolve from the project directory, while
      `static-routes:validation` names `apps/wahidyankf-www/scripts/validate-static-routes.mjs` from the root. The
      shorthand carries `options.cwd` — Nx merges `command` into whatever `options` the target already has — so
      declaring one does not push a single-command target into the aggregate form. [AC-7]
  - Note: sixteen targets, because the Target Contract's fifteen rows fold `dev` and `start` into one.
    `npx nx show project wahidyankf-www --json` resolves every one of them with the `cwd`, cache setting, and `outputs`
    the table states, and the forbidden-token sweep is clean. `static-routes:validation` shows an empty `cwd` in the
    resolved graph rather than a path, which is `{workspaceRoot}` interpolating to the root and is what the contract
    asks for.
- [x] [AI] Declare `outputs` on the four targets that write artifacts, exactly as the Target Contract table states —
      `{projectRoot}/.next` on `build`, `{projectRoot}/public/wahidyankf-kresna-fridayoka-cv.pdf` on `generate:cv-pdf`,
      and `{projectRoot}/coverage` on `test:coverage:unit` and `test:coverage:integration` — acceptance:
      `npx nx show project wahidyankf-www --json` shows all four `outputs` entries. The declaration cannot be exercised
      here, because no application source has landed yet; the Phase 5 item that runs
      `rm -rf apps/wahidyankf-www/.next && npx nx run wahidyankf-www-e2e:test:e2e` is where it is proved end to end.
      `build` is cached, so with no `outputs` Nx replays that run, restores nothing, and still reports success, and
      `next start` then fails on a `.next` directory the replay never wrote.
- [x] [AI] Define `lint` in `apps/wahidyankf-www/project.json` as an ordered aggregate of the two new single-command
      targets `lint:biome` and `lint:commentary`, matching the `npm exec nx -- run <project>:<target>` shape
      `badakmini-cli` uses for `test:quick` — acceptance: `npx nx show project wahidyankf-www --json` shows `lint` with
      `parallel` set to `false` and an `options.commands` list of exactly `npm exec nx -- run wahidyankf-www:lint:biome`
      then `npm exec nx -- run wahidyankf-www:lint:commentary`, and shows `lint:biome` and `lint:commentary` as
      `command`-shorthand targets. The
      [Nx workspace policy](../../../repo-governance/development/nx-workspace-policy.md) requires an aggregate to invoke
      existing target entry points rather than copy their commands, so the two underlying commands live only on the two
      child targets. Biome runs first because a syntax error there makes every ESLint finding noise. [AC-9]
- [x] [AI] Declare `"cache"` explicitly on every target in `apps/wahidyankf-www/project.json` that the Target Contract
      in [technical design](tech-docs/README.md) marks cached or uncached and that root `nx.json` `targetDefaults` does
      not already cover — `lint:biome`, `lint:commentary`, `test:coverage:unit`, `test:coverage:integration`,
      `test:coverage:behaviour`, `test:integration`, `static-routes:validation`, `generate:cv-pdf`, `dev`, and `start` —
      acceptance: `npx nx show project wahidyankf-www --json` reports each of those ten targets with the cache setting
      the table states. Root `targetDefaults` names six target names and no more — `build`, `typecheck`, `lint`,
      `test:unit`, `test:coverage`, and `test:quick` — and a default keyed on `lint` does not reach `lint:biome` or
      `lint:commentary`, both of which the table marks cached, so those two need the setting written out like the other
      eight. Root `nx.json` is deliberately not edited: adding a target name there would change caching for
      `badakmini-cli` too, which is outside this plan's scope.
- [x] [AI] Verify — do not re-set — that the `test:coverage:integration` command in `apps/wahidyankf-www/project.json`
      carries `--coverage.include='src/features/cv/**/pdf.ts'` before its `--coverage.thresholds.lines=99` flag, exactly
      as the Target Contract in [technical design](tech-docs/README.md#target-contract) writes that command and as the
      `project.json` authoring item above transcribed it — acceptance: `npx nx show project wahidyankf-www --json` shows
      that flag in that position on the target, and the `test:coverage:integration` run in Reach the Coverage Floor
      below reports exactly `src/features/cv/core/pdf.ts` and `src/features/cv/shell/pdf.ts` in its coverage table. If
      the flag is missing, the authoring item was not executed as written; add it here and say so, rather than treating
      this as the place the flag is first decided. The check exists because the flag is easy to drop and expensive to
      lose: in Vitest 4 `coverage` is a root-level option a project config cannot set — `ProjectConfig` omits it — so
      the `coverage.include: ["src/**"]` this phase writes into `vitest.config.ts` governs the `integration` project as
      well; the CLI flag replaces that configured include rather than adding to it, which is what narrows this run to
      the two filesystem modules [technical design](tech-docs/README.md#coverage-denominators) names. Without it the
      integration run measures all of `src/**` with two PDF tests executing and the 99% gate cannot pass. [AC-3]
- [x] [AI] Confirm `apps/wahidyankf-www/project.json` defines no `test:e2e` target — acceptance:
      `npx nx show project wahidyankf-www --json | rg -n '"test:e2e"'` finds nothing, because the testing policy names
      an echoing placeholder as the failure mode and absence as the correct signal.
- [x] [AI] Declare the `inputs` array of `test:unit`, `test:coverage:unit`, `test:coverage:behaviour`, `test:coverage`,
      and `test:quick` in `apps/wahidyankf-www/project.json` with `"default"` as its first entry and
      `{workspaceRoot}/specs/apps/wahidyankf-www/behaviours/**/*.feature` as its second — acceptance:
      `npx nx show project wahidyankf-www --json` shows both entries on all five targets, with `"default"` first, so a
      feature edit invalidates each and so does an edit to the project's own files. `"default"` is load-bearing rather
      than decorative: a target-level `inputs` array replaces Nx's default input set instead of extending it, and
      `project.json` is authored fresh earlier in this section, so there is no pre-existing array to append to. An array
      naming the corpus glob alone would hash the corpus and nothing else, and an edit under `src/**` would then leave
      `test:unit`, `test:coverage:unit`, and `test:quick` replaying a green cache over changed source. Both precedents
      put it first: every `inputs` array in `apps/badakmini-cli/project.json`, and `test:unit`, `test:quick`, and
      `static-routes:validation` in the source `apps/wahidyankf-www/project.json`.
- [x] [AI] Add `{workspaceRoot}/scripts/next-with-port.mjs` as a further entry in the `inputs` of `test:unit`,
      `test:coverage:unit`, `test:coverage:behaviour`, `test:coverage`, and `test:quick` in
      `apps/wahidyankf-www/project.json` — acceptance: `npx nx show project wahidyankf-www --json` shows the wrapper
      input on all five targets alongside the two the preceding item declared, with `"default"` still first, so an edit
      to that script invalidates each of them. `apps/wahidyankf-www/tests/bdd/next-with-port-wrapper.unit.test.ts`,
      copied later in this phase, spawns that script and is the only thing in the repository that executes it — but
      `scripts/` sits outside `apps/` and `libs/`, so an edit there marks no project affected: the pre-push affected run
      would skip this project entirely and the cached `test:coverage:unit` and `test:quick` would replay unchanged. The
      contract test would then silently stop running on exactly the change it exists to catch.
      `apps/badakmini-cli/project.json` sets the precedent, naming `{workspaceRoot}/apps/badakmini-cli/tests/e2e/**/*`
      on its own `test:coverage:behaviour` for the same reason.
- [x] [AI] Run `npm install` — acceptance: `apps/wahidyankf-www` resolves as an npm workspace and `npx nx show projects`
      lists `wahidyankf-www`.
  - Note: the install surfaced three **high**-severity advisories against the ported `next@16.2.6` pin, reaching `next`,
    `postcss`, and `sharp`. Bumped `next` and `@next/third-parties` from `16.2.6` to `16.3.3` — the fix npm named, and
    not a semver-major — after which the audit reports zero. Fixed here rather than at the phase gate, which is the
    lesson the Phase 1 `tsx` bump left in `learnings.md`: audit an inherited pin when it is written into the manifest,
    not four items after the code that depends on it lands. `@next/third-parties` moves with `next` because it tracks
    the framework version exactly.

### Inline the Three Libraries

- [x] [AI] Confirm `$SRC/libs/web-ui/src/utils/cn.ts` is not ported — acceptance:
      `test ! -e apps/wahidyankf-www/src/features/ui/core/cn.ts` succeeds. The application already owns the same helper
      at `src/features/app-shell/core/style.ts`, whose `cn` has an identical body and identical `clsx` and
      `tailwind-merge` imports and is already covered by `style.unit.test.ts`; none of the four components copied below
      imports `cn` at all. Porting it would add a second `cn` and a second test for one function, which the
      [file impact](tech-docs/file-impact.md) not-ported list records.
  - Note: Decision E holds under inspection, not just by argument. None of the four inlined components imports `cn` at
    all — the complete import set across the four `.tsx` files is `react`, `lucide-react`, and each other, with
    `@testing-library/react` and `vitest` in the tests. So the helper is not merely duplicated, it is unreferenced by
    everything this migration brings across.
- [x] [AI] Copy the four consumed components from `$SRC/libs/web-ui/src/components/` — `highlight-text`,
      `scroll-to-top`, `search-component`, and `theme-toggle` — into `apps/wahidyankf-www/src/features/ui/shell/`,
      taking each component and its `.unit.test.tsx` and leaving the `.stories.tsx` file behind — acceptance: eight
      files exist under `src/features/ui/shell/` and `ls apps/wahidyankf-www/src/features/ui/shell/*.stories.tsx` finds
      nothing. Each of the four directories holds exactly three files — the component, its story, and its unit test —
      and none holds a `.steps.tsx`, so the story is the only thing left behind.
- [x] [AI] Create `apps/wahidyankf-www/src/features/ui/shell/index.ts` re-exporting `ScrollToTop`, `ThemeToggle`,
      `SearchComponent`, and `HighlightText` — acceptance:
      `rg -n 'ScrollToTop|ThemeToggle|SearchComponent|HighlightText' apps/wahidyankf-www/src/features/ui/shell/index.ts`
      finds all four.
- [x] [AI] Copy `$SRC/libs/web-ui-token/src/tokens.css` to `apps/wahidyankf-www/src/app/tokens.css` and
      `$SRC/libs/web-ui-token/src/wahidyankf.css` to `apps/wahidyankf-www/src/app/theme.css` — acceptance: both files
      exist and neither imports a package specifier.
  - Note: `tokens.css` arrived carrying a header comment that named the old package and showed
    `@import "@open-sharia-enterprise/web-ui-token/src/tokens.css"` as the way to consume it — a specifier that resolves
    to nothing here, in a file the migration is meant to leave free of them. Rewritten to describe what the file holds
    and to say it is imported by relative path now. `theme.css` needed no edit: its one `@import` match is a comment
    recording that fonts are deliberately _not_ imported, which is still true.
- [x] [AI] Add `dotenv` at an exact version to `dependencies` in `apps/wahidyankf-www/package.json` —
      `$SRC/libs/ts-env-loader/package.json` pins `16.4.7`, so use that unless `npm install` reports it unresolvable —
      acceptance: `rg -n '"dotenv": "[0-9]' apps/wahidyankf-www/package.json` finds an exact pin with no range prefix.
      It becomes a direct dependency because inlining the loader inlines what the loader imports; it previously reached
      the application transitively through `@open-sharia-enterprise/ts-env-loader`, which this plan removes.
      [Technical design](tech-docs/README.md#selected-decisions) records the requirement, the rejected
      `process.loadEnvFile` alternative, and the evidence the
      [dependency selection policy](../../../repo-governance/development/dependency-selection-policy.md) asks for.
  - Note, 2026-09-01: `rg -n '"dotenv": "[0-9]' apps/wahidyankf-www/package.json` finds `"dotenv": "16.4.7"`, an exact
    pin with no caret or tilde, and the installed tree resolves 16.4.7. The work was done when the manifest was written;
    **this checkbox was missed at the time and is being ticked late**, which the `@open-sharia-enterprise/` sweep in
    Close the Migration is what surfaced — that sweep prints its matches with line numbers, and reading them showed an
    unticked line among items long since finished. Recorded plainly rather than quietly corrected: the phase's tick
    record was wrong for several subsections, and a plan whose checkboxes drift from its tree is worth less than one
    that admits where it drifted.
- [x] [AI] Port `$SRC/libs/ts-env-loader/src/index.ts` to `apps/wahidyankf-www/src/features/env/core/tier-env.ts`,
      keeping its `dotenv` call and all five loader rules and the stray-file guard over `.env`, `.env.production`, and
      `.env.local` — acceptance: `rg -n 'dotenv' apps/wahidyankf-www/src/features/env/core/tier-env.ts` finds the import
      and the `dotenv.config` call, and `rg -n 'loadEnvFile' apps/wahidyankf-www/src/features/env/core/tier-env.ts`
      finds nothing. Its `export { resolvePort } from "./port-resolver"` needs no repointing, because `port-resolver.ts`
      lands as a sibling in the same directory.
  - Note: three header sentences were repointed along with the code, because they described a shared library rather than
    an inlined module: the "every app in this repo that consumes this module" framing of the five rules, the
    composition-root paragraph that explained the no-auto-load choice by appealing to other consumers, and the closing
    re-export comment, which justified the dependency-free split by naming a container entrypoint. That last one now
    names `scripts/next-with-port.mjs`, which is the real consumer here. All five rules, the `dotenv.config` call with
    `override: false`, and the stray-file guard over `.env`, `.env.production`, and `.env.local` are unchanged;
    `loadEnvFile` appears nowhere, per Decision D.
- [x] [AI] Copy `$SRC/libs/ts-env-loader/src/port-resolver.ts` to
      `apps/wahidyankf-www/src/features/env/core/port-resolver.ts` — acceptance: the file exists and imports nothing,
      preserving the dependency-free contract that `scripts/next-with-port.mjs` documents.
- [x] [AI] Rewrite the header comment of `apps/wahidyankf-www/src/features/env/core/port-resolver.ts`, which describes a
      repo-wide contract mirrored by `libs/fsharp-env-loader`'s `PortResolver`, gives `OSE_WWW_PORT` as its worked
      example, points at two `ose-public` documents, refers to its sibling as `./index.ts`, and explains the
      dependency-free rule by appealing to a pruned `node_modules` inside a built container image — acceptance:
      `rg -in 'rhino-cli|dotnet|\.fsproj|F#|container|OSE_WWW_PORT|organiclever' apps/wahidyankf-www/src/features/env/core/port-resolver.ts`
      finds nothing, `rg -n './index' apps/wahidyankf-www/src/features/env/core/port-resolver.ts` finds nothing, the
      file still imports nothing at all, and the surviving comment still states the three-tier precedence, why a blank
      value falls through while a malformed one is a hard error, and why the module must stay dependency-free. **This
      item was added during execution.** No item covered this file's header, yet the Close the Migration sweep in this
      phase covers `apps/wahidyankf-www*` and would have failed on its five `F#` and `container` matches with nothing in
      the checklist saying what to change. The two sibling rewrites — the wrapper in Phase 1 and
      `next-with-port-wrapper.unit.test.ts` later in this phase — were both anticipated; this one was not.
  - Note: the dependency-free rationale is the part worth reading rather than deleting. It now says the wrapper imports
    this file by relative path before anything has resolved the application's dependencies, and that `./tier-env.ts`
    pulls `dotenv`, so reaching the resolver through the sibling would drag `node_modules` into a path that must not
    need it. That is the same reason the container framing gave, minus the container.

### Port the Source

- [x] [AI] Copy `$SRC/apps/wahidyankf-www/src/` into `apps/wahidyankf-www/src/` as a merge that overwrites same-named
      files and deletes nothing — acceptance:
      `ls apps/wahidyankf-www/src/features/env/core/tier-env.ts apps/wahidyankf-www/src/features/env/core/port-resolver.ts apps/wahidyankf-www/src/features/ui/shell/index.ts`
      still finds all three modules inlined above. The source tree has no `features/ui/` or `features/env/` directory,
      so a merge preserves them; a `rsync --delete` or a wipe-then-copy would destroy the inlining and is the wrong
      shape here. Rewriting `src/env-loader.ts` is a separate job and belongs to the `ts-env-loader` repointing item
      below, which names each of its three references; folding it in here would put one file's edit behind a checkbox
      for a whole-tree copy.
  - Note: done with `cp -R "$SRC/apps/wahidyankf-www/src/." apps/wahidyankf-www/src/`. The trailing `/.` is what makes
    it a merge into the existing directory rather than a copy of the directory into it. Verified afterwards that
    `features/env/core/` still holds both loader modules, `features/ui/shell/` still holds all nine inlined files, and
    `app/tokens.css` and `app/theme.css` survive — the source tree has none of those paths, so a merge preserves them
    and a `--delete` sync would have destroyed the whole inlining.
- [x] [AI] Edit `apps/wahidyankf-www/src/app/globals.css`: repoint the two `@import` lines at `./tokens.css` and
      `./theme.css`, and delete the `@source "../../../../libs/web-ui/src/**/*.{ts,tsx}"` directive — acceptance:
      `rg -n 'open-sharia-enterprise|libs/web-ui' apps/wahidyankf-www/src/app/globals.css` finds nothing. This edit sits
      after the copy on purpose: `globals.css` exists in the source tree, so editing it before the copy would be
      silently overwritten. [AC-6]
  - Note: the `@source` directive is deleted rather than repointed, and the comment above it goes with it. It existed to
    make Tailwind scan a sibling package's sources for class names, because a transpiled dependency sits outside the
    automatic content scan. The four components it was pointing at now live at `src/features/ui/shell/`, inside the tree
    Tailwind v4 already discovers from this stylesheet, so repointing it would restate the default. If a class used only
    by those four ever goes missing from the build, this is the deletion to reconsider first.
- [x] [AI] Search the whole ported and source trees for an importer of each of the four subjects this plan expects to be
      unused — run
      `rg -n 'lucide-react|react-icons|class-variance-authority|search/shell/SearchSection|search/shell/search-section|SEARCH_PLACEHOLDER' apps/wahidyankf-www/src "$SRC/apps/wahidyankf-www/src" "$SRC/apps/wahidyankf-www/test" | sed "s|$SRC|\$SRC|g" > plans/in-progress/wahidyankf-www-migration/evidence/unused-importers.txt`
      — acceptance: `evidence/unused-importers.txt` names every importing file for each subject, or records that a
      subject has none, `rg -n '^/' plans/in-progress/wahidyankf-www-migration/evidence/unused-importers.txt` finds
      nothing, every source-repository path having been rewritten to begin with the literal `$SRC`, and that file's
      entry in the `## Directory Map` of `evidence/README.md` is converted to a relative link in this same item. The
      `sed` filter is not cosmetic: `rg` prefixes each match with the search root it was given, and `"$SRC/..."` expands
      to the executor's own absolute machine path. `evidence/README.md` and
      [plan document safety](../../../repo-governance/conventions/plans-organization-policy/plan-document-safety.md)
      both require that path rewritten before the output is committed, and this file is committed. Inside double quotes
      `\$SRC` is the literal four characters, so the substitution replaces the expanded path with the placeholder this
      checklist uses everywhere else — the same rewrite the Phase 0 evidence item states in prose.
      `$SRC/apps/wahidyankf-www/test` is searched because the nine step files are copied later in this phase and could
      hold the only importer; searching the ported tree alone before they land would prove nothing. The redirect names
      the full `plans/in-progress/wahidyankf-www-migration/evidence/…` path rather than the bare `evidence/` this
      checklist uses in prose. Prose leaves the reader to resolve a relative path against the plan folder; a shell
      redirect does not. Its `rg` search roots are workspace-root-relative, so the executor runs this command from the
      workspace root and `> evidence/…` would create a top-level `evidence/` directory that
      [file impact](tech-docs/file-impact.md) does not list — an unmapped top-level path, which that document defines as
      a signal to stop and amend rather than improvise. The `rg -n '^/'` check beside it carries the full path for the
      same reason: it is a command rather than prose, so it reads the file the redirect wrote rather than one the reader
      resolves.
- [x] [AI] Remove each subject the preceding search found no importer for: the package entry from
      `apps/wahidyankf-www/package.json` for `react-icons`, `lucide-react`, or `class-variance-authority`, and the file
      `apps/wahidyankf-www/src/features/search/shell/SearchSection.tsx`, still PascalCase because the rename below has
      not run yet — acceptance:
      `rg -n 'lucide-react|react-icons|class-variance-authority' apps/wahidyankf-www/package.json` lists only what the
      search found in use, `test ! -e apps/wahidyankf-www/src/features/search/shell/SearchSection.tsx` succeeds unless
      the search found an importer, and `evidence/unused-importers.txt` records what was dropped and why. **Amended
      during execution**: the criterion first named a dated `learnings.md` entry. An importer search's verdict is
      command output rather than a lesson, which is the split [evidence](evidence/README.md) exists to keep, so the
      record belongs in the evidence file the search already writes; no entry was written and none should have been. At
      the recorded source commit `lucide-react` is imported by four application files and three of the four inlined
      components, while `react-icons`, `class-variance-authority`, and `SearchSection.tsx` have no importer anywhere, so
      this item is expected to remove three of the four subjects; a subject with an importer is kept and the entry says
      so instead. The [dependency selection policy](../../../repo-governance/development/dependency-selection-policy.md)
      makes porting an unused dependency a cost with no requirement behind it, and the same reasoning applies to a
      source file no route or test reaches.
  - Note: three of the four expected-unused subjects were confirmed unused and removed — `react-icons`,
    `class-variance-authority`, and `src/features/search/shell/SearchSection.tsx` with its `SEARCH_PLACEHOLDER`
    constant, whose only two matches in either tree were the file declaring it. The fourth, `lucide-react`, is **kept**:
    it has seven importers, four in the ported application and three in the components inlined from `web-ui`. Decision G
    paid for itself here — searching `$SRC/apps/wahidyankf-www/test` as well as both `src` trees is what makes a
    zero-match verdict mean something, since the step files land later in this phase and could have held the only
    importer. Removing `SearchSection.tsx` left `features/search/shell/` empty, so the directory was removed too; the
    `search` component now consists of `core/search.ts` and its unit test, which is what the C4 model's conformist
    relationship actually describes.
- [x] [AI] Rename the remaining PascalCase source files to lower-hyphenated names per the
      [code style policy](../../../repo-governance/development/code-style-policy.md):
      `features/app-shell/shell/Navigation.tsx` to `navigation.tsx`, `features/cv/shell/CvContent.tsx` to
      `cv-content.tsx`, `features/home/shell/HomeContent.tsx` to `home-content.tsx`, and
      `features/personal-projects/shell/PersonalProjectsContent.tsx` to `personal-projects-content.tsx`, plus
      `features/app-shell/shell/Navigation.unit.test.tsx` to `navigation.unit.test.tsx`, which is the only co-located
      test any of the four has — acceptance: `LC_ALL=C find apps/wahidyankf-www/src -name '*[A-Z]*.tsx'` finds nothing.
      `SearchSection.tsx` is not on this list because the preceding item removes it; if that item kept it, rename it
      here too. `LC_ALL=C` is required: under this environment's `en_US.UTF-8` collation, BSD `find` matches `[A-Z]`
      case-insensitively and the same command returns every `.tsx` file in the tree — 25 of them at this point, the
      source's eighteen less `SearchSection.tsx` removed by the preceding item, plus the eight component and unit-test
      files inlined into `src/features/ui/shell/` — whether or not the rename happened, so the acceptance could never be
      satisfied.
  - Note: all five renamed. The `LC_ALL=C` warning in this item is worth keeping — the file count it predicts as the
    false-positive result, 25 `.tsx` files, is exactly what the tree holds after the rename, so an executor who dropped
    the locale prefix would see a number that looks like a real finding rather than an obvious collation artifact. A
    matching sweep over `*[A-Z]*.ts` also returns nothing, so no non-JSX module needed renaming.
- [x] [AI] Repoint the eleven `@open-sharia-enterprise/web-ui` references in `apps/wahidyankf-www/src` at
      `@/features/ui/shell` — five `import` statements in `app/layout.tsx`, `features/cv/shell/markdown.tsx`,
      `features/cv/shell/cv-content.tsx`, `features/home/shell/home-content.tsx`, and
      `features/personal-projects/shell/personal-projects-content.tsx`; and six `vi.mock` targets in
      `features/cv/shell/markdown.unit.test.tsx`, `app/layout.unit.test.tsx`, `app/page.unit.test.tsx`,
      `app/cv/page.unit.test.tsx`, `app/personal-projects/page.unit.test.tsx`, and
      `app/static-route-content.unit.test.tsx` — acceptance:
      `rg -c 'open-sharia-enterprise/web-ui' apps/wahidyankf-www/src` finds nothing and
      `rg -c '@/features/ui/shell' apps/wahidyankf-www/src` reports eleven matching lines. A twelfth reference, a prose
      comment at the top of `SearchSection.tsx`, goes with the file the unused-importer item removes; if that item kept
      the file, repoint the comment too and expect twelve. [AC-6]
  - Note: eleven exactly, as the plan counted — five `import` statements and six `vi.mock` calls, across eleven distinct
    files with one reference each.
- [x] [AI] Repoint every reference to the renamed modules inside `apps/wahidyankf-www/src` — seventeen in all, none of
      them affected by the `SearchSection.tsx` removal, because nothing referenced that file. Eight name `Navigation`:
      three `@/features/app-shell/shell/Navigation` imports in `features/cv/shell/cv-content.tsx`,
      `features/home/shell/home-content.tsx`, and `features/personal-projects/shell/personal-projects-content.tsx`; one
      relative `./Navigation` import in `features/app-shell/shell/navigation.unit.test.tsx`; and four
      `vi.mock("@/features/app-shell/shell/Navigation")` targets in `app/page.unit.test.tsx`,
      `app/cv/page.unit.test.tsx`, `app/personal-projects/page.unit.test.tsx`, and
      `app/static-route-content.unit.test.tsx`. Nine name a content component: six imports, in `app/page.tsx` and
      `app/page.unit.test.tsx` for `HomeContent`, `app/cv/page.tsx` and `app/cv/page.unit.test.tsx` for `CvContent`, and
      `app/personal-projects/page.tsx` and `app/personal-projects/page.unit.test.tsx` for `PersonalProjectsContent`; and
      three `contentSource` string literals in `app/static-route-content.unit.test.tsx`, which that test hands to
      `readFileSync` — acceptance:
      `rg -n 'shell/Navigation|shell/CvContent|shell/HomeContent|shell/PersonalProjectsContent|\./Navigation' apps/wahidyankf-www/src`
      finds nothing. The owner's APFS volume is case-insensitive, so every one of these still resolves locally after the
      rename and breaks only on the Linux CI runner: an unresolved `vi.mock` mocks nothing and silently renders the real
      component, and an unresolved `contentSource` throws `ENOENT` from `readFileSync`.
  - Note: seventeen exactly, and the count only closes if the last three are found. Fourteen are module specifiers an
    import-path sweep reaches. The other three are **string literals**, not imports:
    `static-route-content.unit.test.tsx` records a `contentSource` for each route as a relative path with the `.tsx`
    extension, so a pattern anchored on `@/features/...` or on a `from` clause steps straight over them. They are not
    decoration — that test reads them to assert each static route renders from the file it claims to. Left stale, they
    would name three files that no longer exist while the suite still passed on everything it could resolve.
- [x] [AI] Repoint the three `@open-sharia-enterprise/ts-env-loader` references in
      `apps/wahidyankf-www/src/env-loader.ts` at `@/features/env/core/tier-env`: the `import { loadTierEnv }`, the
      `export { loadTierEnv, resolveTier, tierEnvFilePath }` re-export, and the prose comment that says the loader logic
      lives in a package "shared across every Next.js app in this repo", which is no longer where it lives or true here
      — acceptance: `rg -n 'open-sharia-enterprise' apps/wahidyankf-www/src` finds nothing at all, the two
      `web-ui-token` CSS imports having been handled by the `globals.css` item above, and
      `rg -n 'features/env/core/tier-env' apps/wahidyankf-www/src/env-loader.ts` finds both the import and the
      re-export. [AC-6]
  - Note: the prose reference is the one worth reading rather than pattern-replacing. It said the loader logic lives in
    a package "shared across every Next.js app in this repo" and explained the no-auto-load rule by appealing to what a
    shared library must not do — both false once the module is inlined into the only app that uses it. Rewritten to give
    the same reason on its own terms: which loader won would otherwise depend on import order.
    `rg -n 'open-sharia-enterprise' apps/wahidyankf-www/src` now finds nothing at all.
- [x] [AI] Copy `$SRC/apps/wahidyankf-www/public/` and
      `$SRC/apps/wahidyankf-www/scripts/{generate-cv-pdf.ts,validate-static-routes.mjs}` — acceptance: all five files
      exist and `node --check apps/wahidyankf-www/scripts/validate-static-routes.mjs` exits 0. `--check` parses without
      executing, which is the only smoke test available here: the script handles no flags and reads
      `.next/prerender-manifest.json` and `.next/routes-manifest.json` at module top level, so any run at all —
      including `--help` — exits on `ENOENT` before the build below produces `.next`. Its behaviour is proved later by
      `static-routes:validation`.
- [x] [AI] Try Node's own type stripping before relying on the `tsx` pin — run
      `node apps/wahidyankf-www/scripts/generate-cv-pdf.ts` and record the command, its exit status, and any error text
      in `evidence/node-type-stripping.md`, converting that file's entry in the `## Directory Map` of
      `evidence/README.md` to a relative link in this same item — acceptance: one dated entry holds that evidence, the
      map entry is a link, and the next item reads its verdict. The script imports three `src/` modules with
      extensionless specifiers and uses `__dirname`, so whether Node runs it unchanged is a question to answer by
      running it. The [dependency selection policy](../../../repo-governance/development/dependency-selection-policy.md)
      requires reaching for the standard library first, and this is that attempt.
  - Note: the attempt **failed**, so the `tsx` pin stays. Node stripped the types successfully and then failed at module
    resolution with `ERR_MODULE_NOT_FOUND` on the script's extensionless `.../core/data` import, which Node's ESM
    resolver will not extend to `.ts`. That is the precise capability `tsx` supplies, and
    [`evidence/node-type-stripping.md`](evidence/node-type-stripping.md) records the command, the exit status, and the
    full error.
- [ ] [AI] Remove the `tsx` pin from root `package.json` and change the `generate:cv-pdf` command to
      `node scripts/generate-cv-pdf.ts` — acceptance: `rg -n '"tsx"' package.json` finds nothing and
      `npx nx run wahidyankf-www:generate:cv-pdf` exits 0. Trigger: the preceding attempt succeeded. If it failed, this
      item takes a dated `Not triggered` disposition in Phase 7 and `evidence/node-type-stripping.md` carries the error
      that justifies keeping the dependency.
  - **Not triggered, 2026-09-01.** Its stated trigger is that the preceding attempt succeeded; it exited 1. `tsx`
    remains pinned at `4.23.13` and `generate:cv-pdf` keeps its `npx tsx` command.
    [evidence](evidence/node-type-stripping.md) carries the command, its exit status, and the error text that justifies
    keeping the dependency, so `rg -n '"tsx"' package.json` still finding the pin is the correct delivered state rather
    than an unfinished removal. Deliberately left unticked.
- [x] [AI] Run `npx nx run wahidyankf-www:typecheck` — acceptance: exits 0 on the pinned TypeScript version.
  - Note: **passes on TypeScript 6.0.3**, two majors above the `5.8.3` the source pins, so the fallback item below is
    not triggered. It took two `tsconfig.json` changes, both of which are the port meeting a newer compiler rather than
    the application being wrong. First, TS 6 errors on `baseUrl` as deprecated and removed in 7 (`TS5101`). It was
    deleted rather than silenced with `ignoreDeprecations: "6.0"`: the only reason it was there is the `@/*` path
    mapping, and since TypeScript 5 `paths` resolves relative to the config file's own directory when `baseUrl` is
    absent, so the option was already doing nothing. Silencing it would have bought a compiler major and no more.
    Second, `next.config.ts` imports `./src/env-loader.ts` and `./src/env.ts` with explicit extensions, which is
    `TS5097` unless `allowImportingTsExtensions` is on. That option is enabled, which is legal here because the project
    sets `noEmit: true` — the constraint the option carries. The alternative, stripping `.ts` from the two specifiers,
    would have edited ported source to satisfy a setting that exists precisely to permit it.
- [x] [AI] Run `npx nx run wahidyankf-www:build` — acceptance: exits 0 and `apps/wahidyankf-www/.next` is produced.
  - Note: passes, producing `.next` with all six routes prerendered as static content — `/`, `/_not-found`, `/cv`,
    `/personal-projects`, `/robots.txt`, `/sitemap.xml` — which is the fully-static claim the C4 model makes, now
    observed rather than asserted. One fix was needed and it is not obvious from reading either file.
    `src/env-loader.ts` had been repointed at `@/features/env/core/tier-env`, matching every other module; the build
    then failed with `Cannot find module` before it began. Next compiles `next.config.ts` through its own config
    transpiler, which does not read this project's `paths` mapping, and `env-loader.ts` is imported from that config —
    so it is the single module in this application that cannot use the `@/` alias. Its specifier is now relative, with a
    comment saying why, because the next reader to "fix the inconsistency" would reintroduce the failure.
- [ ] [AI] If `typecheck` or `build` fails on the TypeScript version, pin TypeScript at the highest version that passes,
      keep Biome, and record the exact error and the blocking version in `learnings.md` — acceptance:
      `npx nx run wahidyankf-www:typecheck` exits 0 and `learnings.md` names the version and the error. Trigger: a
      TypeScript-version-attributable failure in either command. [AC-9]
  - **Not triggered, 2026-09-01.** Both `typecheck` and `build` pass on TypeScript `6.0.3`, so no fallback pin was
    needed and none was taken; [evidence](evidence/phase-1-toolchain.md) records the resolved version. The two errors
    that did appear were a deprecated `baseUrl` and an unset `allowImportingTsExtensions` — configuration meeting a
    newer compiler, not a source construct the compiler rejects — and neither is a version block. `baseUrl` surfaced
    again in Phase 5 for the same reason and was fixed the same way, on its own item. Deliberately left unticked. [AC-9]
- [x] [AI] Add `@typescript-eslint/parser` at an exact version to root `devDependencies` and set it as
      `languageOptions.parser` in `apps/wahidyankf-www/eslint.config.mjs`, with `parserOptions.ecmaFeatures.jsx` enabled
      and no rule from that package turned on — acceptance: `npm audit --audit-level=low` exits 0 with the new pin
      installed, and
      `cd apps/wahidyankf-www && node ../../node_modules/eslint/bin/eslint.js --config eslint.config.mjs src` reports no
      `Parsing error`. **This item was added during execution.** Phase 1 installed `eslint` and `eslint-plugin-jsdoc`
      and nothing else, which is what `tooling.md` names; neither package carries a reader for TypeScript or JSX, so
      ESLint fell back to its own JavaScript parser.
  - Note, 2026-09-01: `@typescript-eslint/parser` `8.69.0`, installed with `--save-exact`; `npm audit --audit-level=low`
    exits 0 with it in the tree. The parser is loaded for its reader alone and no `@typescript-eslint` rule is enabled,
    because every rule that package offers for this project is one Biome already reports — the duplicate-finding split
    `eslint.config.mjs` already documents. The config comment says so at the `languageOptions` block, so the next reader
    does not add rules on the assumption the plugin was brought in for them.
- [x] [AI] Run `npx nx run wahidyankf-www:lint` — acceptance: exits 0 across both child targets, `lint:biome` then
      `lint:commentary`, with no blanket ignore and no rule disabled without a nearby reason.
  - Note, 2026-09-01: exits 0. `lint:biome` reports `Checked 60 files in 51ms. No fixes applied.` and `lint:commentary`
    reports nothing. Getting there took three distinct kinds of work. First, ESLint could not read the sources at all —
    39 `Parsing error: Unexpected token {` and `Unexpected token <` across 25 files — which is the added item above.
    Second, `jsdoc/require-description-complete-sentence` fired three times, and two were the rule's sentence splitter
    ending a sentence at the period inside `e.g.` and reading the lowercase word after it as a new sentence's first
    word; that is punctuation misread, not a finding, so the rule now carries an `abbreviations` list with a comment
    saying why rather than being switched off. The third was real — a `@throws` description genuinely starting lowercase
    — and was reworded. Third, `jsdoc/require-jsdoc` reported 14 undocumented declarations, each of which was given a
    written sentence saying what the declaration is for and, where there was one, the reason behind its shape. `--fix`
    was available for all 14 and was not used: it inserts an empty comment frame, which satisfies the rule and tells a
    reader nothing, and this phase already recorded once what an unreviewed bulk autofix costs. One of the 14 replaced a
    stale `// In the main CV component, update the type of topSkills` note left over from the source repository. No rule
    is disabled and no path is ignored.
- [ ] [AI] If Biome cannot lint the Next 16 and React 19 sources, retain `oxlint` with the ported `oxlint.json` as the
      `lint:biome` command, keep the TypeScript version and the commentary check, and record the exact failure in
      `learnings.md` — acceptance: `npx nx run wahidyankf-www:lint` exits 0 and `learnings.md` names the Biome failure.
      Trigger: a Biome-attributable lint failure that no narrow, explained suppression resolves. [AC-9]
  - **Not triggered, 2026-09-01.** Biome `2.5.11` lints all 60 files of this Next 16 and React 19 application and
    reports nothing, and `lint:biome` exits 0 for both TypeScript projects. It did report 58 findings and 23 CSS parse
    errors on first run; every one was resolved by fixing the source, by enabling `css.parser.tailwindDirectives` so
    Biome parses the Tailwind at-rules it was choking on, or by a narrow `biome-ignore` carrying its reason. `oxlint` is
    not retained and `oxlint.json` is not ported. Biome does run as a linter only, but that is a deliberate boundary
    recorded as a deviation in [tooling](../../../repo-governance/development/testing-policy/tooling.md) rather than
    this fallback firing. Deliberately left unticked. [AC-9]
- [ ] [AI] If ESLint or `eslint-plugin-jsdoc` cannot run against these sources, drop `lint:commentary` from the
      aggregate, keep TypeScript and Biome, and record the exact failure in `learnings.md` — acceptance:
      `npx nx run wahidyankf-www:lint` exits 0 and `learnings.md` names the ESLint failure and which of the two packages
      caused it. Trigger: an ESLint- or plugin-attributable failure that no narrow, explained suppression resolves. The
      fallback ladder covers this component the same way it covers TypeScript and Biome, because `tooling.md` names all
      three. [AC-9]
  - **Not triggered, 2026-09-01.** ESLint `9.39.4` and `eslint-plugin-jsdoc` `64.3.2` both run against these sources and
    `lint:commentary` exits 0 in both TypeScript projects. The failure that looked like this trigger was neither
    package's: ESLint had no parser able to read TypeScript or JSX, which adding `@typescript-eslint/parser` fixed, and
    two further findings were real rule behaviour that the rule's own `abbreviations` option resolved — exactly the
    narrow, explained suppression this trigger excludes. Dropping `lint:commentary` on that evidence would have removed
    a working check because its configuration was incomplete, so the trigger was read against its actual wording and
    found not to have fired. Deliberately left unticked. [AC-9]

### Bind the Corpus

- [x] [AI] Copy the nine step files from `$SRC/apps/wahidyankf-www/test/unit/steps/` to
      `apps/wahidyankf-www/tests/bdd/`, matching the `tests/bdd` layout `apps/badakmini-cli` already uses, and repoint
      each `loadFeature` path at `specs/apps/wahidyankf-www/behaviours/<name>.feature` — acceptance:
      `ls apps/wahidyankf-www/tests/bdd/*.steps.ts | wc -l` prints `9` and
      `rg -n 'gherkin/' apps/wahidyankf-www/tests/bdd` finds no stale nested path.
  - Note, 2026-09-01: nine files copied; `ls apps/wahidyankf-www/tests/bdd/*.steps.ts | wc -l` prints `9` and
    `rg -n 'gherkin/' apps/wahidyankf-www/tests/bdd` finds nothing. The repointing was wider than the `loadFeature` call
    in each file. Every step in every file carries a `@covers specs/.../gherkin/<group>/<name>.feature:<scenario name>`
    comment naming the scenario it implements, so the old nested path appeared 53 more times than the nine load sites.
    One substitution of the nested prefix `specs/apps/wahidyankf/behaviours/wahidyankf-www/gherkin/<group>/` for the
    flat `specs/apps/wahidyankf-www/behaviours/` covered both, and the `gherkin/` sweep in this item's own acceptance is
    what would have caught it had it not. Each of the nine now names a feature file that exists.
- [x] [AI] Repoint the fourteen references to the renamed PascalCase modules inside the copied step files. Six name
      `Navigation`: the `@/features/app-shell/shell/Navigation` import in `responsive.steps.ts`, and the
      `vi.mock("@/features/app-shell/shell/Navigation")` targets in `cv.steps.ts`, `home.steps.ts`,
      `personal-projects.steps.ts`, `search.steps.ts`, and `static-filterable-routes.steps.ts`. Eight name a content
      component: `HomeContent` in `home.steps.ts`, `search.steps.ts`, and `accessibility.steps.ts`; `CvContent` in
      `cv.steps.ts`, `static-filterable-routes.steps.ts`, and `accessibility.steps.ts`; and `PersonalProjectsContent` in
      `personal-projects.steps.ts` and `accessibility.steps.ts` — acceptance:
      `rg -n 'shell/Navigation|shell/CvContent|shell/HomeContent|shell/PersonalProjectsContent' apps/wahidyankf-www/tests`
      finds nothing. This item sits here rather than beside the Phase 3 rename because the files it edits do not exist
      until the copy above; the rename item cannot reach them, and on the owner's case-insensitive APFS volume every
      stale path still resolves locally, so nothing local reports the gap.
  - Note, 2026-09-01: fourteen references rewritten, and the count the plan predicted is the count found — six
    `Navigation` and eight content components, distributed across exactly the files named.
    `rg -n 'shell/Navigation|shell/CvContent|shell/HomeContent|shell/PersonalProjectsContent' apps/wahidyankf-www/tests`
    now finds nothing. Each new specifier was checked against the file that exists rather than derived by rule:
    `navigation.tsx`, `cv-content.tsx`, `home-content.tsx`, `personal-projects-content.tsx`. The item's own reasoning
    held — on this case-insensitive volume every one of the fourteen stale paths resolved fine locally, so nothing but
    this sweep would have reported them before CI.
- [x] [AI] Repoint the five `@open-sharia-enterprise` references inside the copied step files — the `ThemeToggle`
      imports in `theme.steps.ts`, `accessibility.steps.ts`, and `responsive.steps.ts`, and the
      `vi.mock("@open-sharia-enterprise/web-ui")` in `static-filterable-routes.steps.ts`, all at `@/features/ui/shell`;
      and the `loadTierEnv` import in `env-loader.steps.ts` at `@/features/env/core/tier-env` — acceptance:
      `rg -n 'open-sharia-enterprise' apps/wahidyankf-www/tests` finds nothing. Without this the behaviour run cannot
      resolve its imports at all, so the `test:coverage:behaviour` check below would never reach exit 0. [AC-6]
  - Note, 2026-09-01: five references rewritten and `rg -n 'open-sharia-enterprise' apps/wahidyankf-www/tests` finds
    nothing. Four `web-ui` specifiers now read `@/features/ui/shell`, which resolves to the barrel the inlining step
    wrote; that barrel re-exports `ThemeToggle` as a named symbol even though the component declares itself as a
    default, so the three `import { ThemeToggle }` sites needed no change beyond the path. The `vi.mock` in
    `static-filterable-routes.steps.ts` mocks the same barrel, so the mocked specifier and the real one still name one
    module and the mock will take effect. The fifth is `loadTierEnv`, now at `@/features/env/core/tier-env`.
- [x] [AI] Copy `$SRC/libs/ts-env-loader/src/env-loader.unit.test.ts` to
      `apps/wahidyankf-www/tests/bdd/tier-env.unit.test.ts`, repointing its `loadFeature` path at
      `specs/apps/wahidyankf-www/behaviours/tier-env-loading.feature` and its module import at
      `@/features/env/core/tier-env` — acceptance:
      `rg -n 'tier-env-loading.feature' apps/wahidyankf-www/tests/bdd/tier-env.unit.test.ts` finds the binding and
      `rg -n 'ts-env-loader|\./index' apps/wahidyankf-www/tests/bdd/tier-env.unit.test.ts` finds nothing. It lives under
      `tests/bdd/` rather than beside the source because it is a Gherkin binding, and that is where this repository puts
      those.
  - Note, 2026-09-01: both acceptances hold. The repointing reached four things, not two. The `loadFeature` path and the
    module import are the two the item names; beyond them, the file's header said it bound "ts-env-loader's own" feature
    and each of its five steps carried a `@covers` comment naming the old library path, and a temp-directory prefix
    string read `ts-env-loader-test-`. The `rg 'ts-env-loader'` half of the acceptance is what surfaced the last one,
    which no amount of reading the import list would have. The `loadFeature` base was also changed from `__dirname` to
    `process.cwd()`, matching the nine step files now sitting beside it, so all bindings in this directory resolve a
    feature the same way and none depends on `__dirname` being defined under whichever module format vitest gives them.
- [x] [AI] Copy `$SRC/libs/ts-env-loader/src/port-resolver.unit.test.ts` to
      `apps/wahidyankf-www/tests/bdd/port-resolver.unit.test.ts`, repointing its `loadFeature` path at
      `specs/apps/wahidyankf-www/behaviours/port-resolver.feature` and its module import at
      `@/features/env/core/port-resolver` — acceptance:
      `rg -n 'port-resolver.feature' apps/wahidyankf-www/tests/bdd/port-resolver.unit.test.ts` finds the binding and
      `rg -n 'ts-env-loader|\./index' apps/wahidyankf-www/tests/bdd/port-resolver.unit.test.ts` finds nothing.
  - Note, 2026-09-01: both acceptances hold. `port-resolver.feature` is named ten times — the `Covers:` header line, the
    `loadFeature` call, and the eight `@covers` comments, one per scenario — and no `ts-env-loader` or `./index`
    reference survives; the last of those was in the header sentence rather than in any import. The `loadFeature` base
    was moved from `__dirname` to `process.cwd()` for the same reason as its sibling above. The F# pairing paragraph
    this file also carries is left standing here, because the next item removes it and doing it now would tick that
    item's acceptance before reaching it.
- [x] [AI] Delete the paragraph in `apps/wahidyankf-www/tests/bdd/port-resolver.unit.test.ts` that pairs these scenarios
      one-for-one with `libs/fsharp-env-loader/tests/unit/Tests/PortResolverTests.fs` — acceptance:
      `rg -n 'fsharp|PortResolverTests|F#' apps/wahidyankf-www/tests/bdd/port-resolver.unit.test.ts` finds nothing. That
      twin does not exist here, so the comment would tell a future reader to keep a file in sync with nothing. [AC-7]
  - Note, 2026-09-01: the paragraph is gone and `rg -n 'fsharp|PortResolverTests|F#'` finds nothing in the file. Only
    that paragraph was removed; the sentence above it, saying every scenario drives `resolvePort()` with an isolated
    `env` record rather than the real `process.env`, is about this suite's own isolation and still holds. What the
    deleted text asserted was a cross-implementation contract that no longer has a second side here — this repository
    has one port resolver, not a TypeScript one and an F# one — so leaving it would have told a reader to keep a twin in
    sync with a file that does not exist.
- [x] [AI] Copy `$SRC/libs/ts-env-loader/src/next-with-port-wrapper.unit.test.ts` to
      `apps/wahidyankf-www/tests/bdd/next-with-port-wrapper.unit.test.ts`, repointing the wrapper path it spawns at the
      repository-root `scripts/next-with-port.mjs` — acceptance:
      `(cd apps/wahidyankf-www && npx vitest --config vitest.config.ts run --project behaviour tests/bdd/next-with-port-wrapper.unit.test.ts)`
      exits 0 **and** its summary reports one test file with at least one passing test, not `No test files found`. Both
      halves are needed: `passWithNoTests: true` in the ported config makes a run that collected nothing exit 0 too, so
      an exit status alone would not distinguish the contract test passing from the contract test never being reached.
      This plan copies the wrapper, so its contract test comes with it; nothing else in the repository executes that
      file.
  - Note, 2026-09-01: exits 0 reporting `Test Files 1 passed (1)` and `Tests 6 passed (6)`, so both halves of the
    acceptance are satisfied and the `passWithNoTests` hazard the item names did not apply. The repointing was a depth
    change rather than a path change: the file computes its repository root by walking up from its own location, and it
    moved from `libs/ts-env-loader/src` to `apps/wahidyankf-www/tests/bdd`, so `../../..` became `../../../..`. A
    comment now states which directory that count is walking up from, because the constant is silently wrong if the file
    is ever moved again and every one of its six tests would then fail on a spawn error rather than on the contract.
  - Note, 2026-09-01: the run emitted two advisory warnings that are not failures and are not acted on here. Vite warns
    that `vitest.config.ts` uses ESM syntax while being loaded as CommonJS under a `configLoader` that is planned to
    become the default, and that `vite-tsconfig-paths` now has a native equivalent. Neither affects this run or any
    target's exit status; both are future-version notices about the config written in Scaffold, and changing that config
    to chase them is outside every item in this phase.
- [x] [AI] Rewrite the header comment of `apps/wahidyankf-www/tests/bdd/next-with-port-wrapper.unit.test.ts` matching
      the wrapper rewrite in Phase 1. Drop the "four of the six container images" clause from the `--server` paragraph,
      which names images this repository does not build. Keep the sentence "The wrapper lives at the repository root,
      outside any Nx project, so nothing would otherwise execute it" exactly as it stands: it is still true here, and it
      is the reasoning the Scaffold item above depends on when it adds `{workspaceRoot}/scripts/next-with-port.mjs` to
      the `inputs` of five targets, because `scripts/` sits outside `apps/` and `libs/`. The stale sentence is the one
      after it — "it exists solely to apply this library's `resolvePort` to a Next.js server" — which names a library
      that no longer exists once the loader is inlined; repoint `this library's resolvePort` at
      `apps/wahidyankf-www/src/features/env/core/port-resolver.ts`, leaving the rest of that sentence intact. Also
      rename its test case `it("rejects the numeric-literal forms the F# resolver also rejects")` to name the rejected
      forms — `"0x10"`, `"1e3"`, `"+3100"`, and `"0b1010"` — rather than the absent F# twin — acceptance:
      `rg -n 'container|six|F#|this library' apps/wahidyankf-www/tests/bdd/next-with-port-wrapper.unit.test.ts` finds
      nothing, `rg -n 'outside any Nx project' apps/wahidyankf-www/tests/bdd/next-with-port-wrapper.unit.test.ts` still
      finds the retained sentence,
      `rg -n 'features/env/core/port-resolver' apps/wahidyankf-www/tests/bdd/next-with-port-wrapper.unit.test.ts` finds
      the repointed reference, and the surviving comment still says why only the `--server` form is exercised. The
      `this library` alternative is what makes the stale sentence match at all: it carries none of `container`, `six`,
      or `F#`, so the three-token pattern alone would report the file clean while the dangling reference survived, and
      the retained sentence carries none of them either, so the pattern never asked for its deletion. This is the same
      dangling reference the `port-resolver.unit.test.ts` item removes, and it is the one the
      `rhino-cli|dotnet|\.fsproj` sweep in Close the Migration would miss without its `F#` alternative. [AC-7]
  - Note, 2026-09-01: all four acceptance clauses hold — `rg -n 'container|six|F#|this library'` finds nothing,
    `outside any Nx project` is still present at line 4, `features/env/core/port-resolver` is present at line 6, and the
    surviving paragraph still gives the reason only the `--server` form is exercised. Six tests still pass after the
    edit. The item's warning about the `this library` alternative was well placed: the dangling sentence carried none of
    `container`, `six`, or `F#`, so a three-token sweep would have called the file clean while it still told a reader to
    look for a library this repository does not have. One thing was written and then withdrawn: the replacement
    `--server` paragraph first ended by saying the spawning form is what the end-to-end suite covers, which is a claim
    about a harness Phase 5 has not built yet. A comment asserting a fact about work not yet done is wrong at the moment
    it is written, so the sentence now states only what booting the spawning form would cost.
- [x] [AI] Confirm the `behaviour` project defined in Scaffold now discovers every binding — acceptance:
      `(cd apps/wahidyankf-www && npx vitest --config vitest.config.ts list --project behaviour)` lists exactly twelve
      files, the nine `.steps.ts` bindings plus `tier-env.unit.test.ts`, `port-resolver.unit.test.ts`, and
      `next-with-port-wrapper.unit.test.ts`, and no other file.
  - Note, 2026-09-01: exactly twelve files, and exactly the twelve the item names — the nine `.steps.ts` bindings plus
    `tier-env.unit.test.ts`, `port-resolver.unit.test.ts`, and `next-with-port-wrapper.unit.test.ts`, with no
    thirteenth. Reading the result took a detour worth recording: `vitest list` prints one line per step, not per file,
    and the harness's command proxy collapsed the whole listing to `PASS (0) FAIL (0)` — a test-run summary for a
    command that ran no tests. The listing was recovered with `rtk proxy` to bypass the rewriting, and the distinct
    files counted from it. This is the same hazard Phase 1 recorded for `npx tsc --version`, in a second shape: the
    proxy did not drop a flag here, it reformatted output as though a different command had run.
  - Note, 2026-09-01: the recovered listing showed every `port-resolver` scenario naming the environment variable
    `OSE_WWW_PORT`. That is followed up by the item added directly below.
- [x] [AI] Rename `OSE_WWW_PORT` to `WAHIDYANKF_WWW_PORT` in
      `specs/apps/wahidyankf-www/behaviours/port-resolver.feature` and
      `apps/wahidyankf-www/tests/bdd/port-resolver.unit.test.ts`, changing both in one edit so the Gherkin step text and
      the binding's step string stay identical — acceptance:
      `rg -l 'OSE_WWW_PORT' --hidden --glob '!node_modules' --glob '!plans/**' .` finds nothing,
      `grep -cE '^[[:space:]]*(Scenario|Scenario Outline):' specs/apps/wahidyankf-www/behaviours/port-resolver.feature`
      still prints `8`, and
      `(cd apps/wahidyankf-www && npx vitest --config vitest.config.ts run --project behaviour tests/bdd/port-resolver.unit.test.ts)`
      exits 0. **This item was added during execution.** The discovered header rewrite earlier in this phase stripped
      `OSE_WWW_PORT` from `port-resolver.ts` on the grounds that it names a variable this repository does not have, but
      it reached only that one file; the corpus and its binding kept the name, and no sweep in Close the Migration looks
      for it — `@open-sharia-enterprise/` is the specifier form and matches nothing here, and the
      `rhino-cli|dotnet|\.fsproj|F#` pattern has no alternative for it. The glob excluding `plans/**` is needed for the
      same reason the `@open-sharia-enterprise/` sweep needs one: these documents name the old variable in order to
      record removing it. [AC-6]
  - Note, 2026-09-01: 18 occurrences in the feature and 33 in the binding, all renamed together; `rg -l` outside this
    plan's documents now finds none, the eight scenario titles are untouched, and the binding passes all 77 of its
    tests. The lockstep matters and the test run is what proves it held: `@amiceli/vitest-cucumber` matches a step by
    its literal string, so renaming the variable in the feature alone would have left eight scenarios with steps no
    binding declares, and renaming it in the binding alone would have done the mirror image. Neither would have been a
    type error or a lint finding. This edit does not disturb any Phase 2 acceptance: those compare scenario **titles**
    between this corpus and the source, and the variable appears only in step lines beneath them.
- [x] [AI] Verify — do not re-set — that the `test:coverage:unit` command in `apps/wahidyankf-www/project.json` reads
      `vitest run --project unit --project behaviour --coverage --coverage.thresholds.lines=99`, exactly as the Target
      Contract in [technical design](tech-docs/README.md#target-contract) writes it and as the `project.json` authoring
      item in Scaffold transcribed it — acceptance: `npx nx show project wahidyankf-www --json` shows both project flags
      on that target. If only `--project unit` is there, the authoring item was not executed as written; add the second
      flag here and say so, rather than treating this as the place it is first decided. The check sits here, after the
      bindings land, because this is the first point at which the missing flag would have a visible cost:
      `src/features/env/core/tier-env.ts` and `port-resolver.ts` are exercised only by the `tests/bdd` bindings, so a
      unit-only run would report them at zero against a `src/**` denominator. `badakmini-cli` does the same thing for
      the same reason: its `test:coverage:unit` passes `./tests/bdd` alongside `./tests/unit`. [AC-2]
  - Note, 2026-09-01: verified rather than re-set, and it matches. `apps/wahidyankf-www/project.json` holds
    `vitest run --project unit --project behaviour --coverage --coverage.thresholds.lines=99`, character for character
    what the Target Contract's row states. The surrounding keys agree too — `cwd` is `{projectRoot}`, `cache` is true,
    `outputs` is `{projectRoot}/coverage`, and `inputs` begins with `"default"` before naming the corpus glob and the
    wrapper script. The point of verifying instead of setting is that the Scaffold item already wrote this target and
    asserted it against the same contract; re-setting it here would make the two items agree with each other rather than
    with the contract, so a drift introduced in Scaffold would be quietly repaired instead of reported.
- [x] [AI] Confirm each of the eleven feature files is named by exactly one binding file under `tests/bdd/` —
      acceptance:
      `for f in specs/apps/wahidyankf-www/behaviours/*.feature; do n=$(basename "$f"); printf '%s %s\n' "$n" "$(rg -l "behaviours/$n" apps/wahidyankf-www/tests/bdd | wc -l | tr -d ' ')"; done`
      prints eleven lines and every one ends in `1`. The runner cannot detect a feature file that no binding loads at
      all, so this count is what proves the mapping. [AC-4]
  - Note, 2026-09-01: eleven lines, every one ending in `1` — no feature is loaded twice and, more to the point, none is
    loaded zero times. This is the check the item says it is: the runner reports a scenario a loaded binding fails to
    declare, but it has nothing to say about a feature file no binding ever opens, so a corpus file could sit in the
    directory contributing to the count of eleven while specifying nothing that runs. The twelfth file,
    `cv-export.feature`, does not exist yet and is authored by the next subsection, which is why this count is eleven
    rather than twelve.
- [x] [AI] Run `npx nx run wahidyankf-www:test:coverage:behaviour` — acceptance: exits 0, and the run reports all 53
      scenarios across the eleven loaded features. `@amiceli/vitest-cucumber` fails the run on a missing scenario, a
      missing or mistyped step, and a missing Scenario Outline variable inside a loaded feature; it does not detect an
      unused binding. [AC-4]
  - Note, 2026-09-01: exits 0 with `Test Files 12 passed (12)` and `Tests 258 passed (258)`. The scenario count was
    checked rather than inferred from the pass: the distinct `Scenario` names in the run are 53, and
    `grep -rhE '^[[:space:]]*(Scenario|Scenario Outline):'` over the eleven feature files is also 53, so every specified
    scenario was reached and none was invented.
  - Note, 2026-09-01: the first run failed two scenarios, and the cause was a defect this phase introduced rather than
    anything the port carried. Both failures were `expected null not to be null` on
    `document.getElementById('project-<n>')` in `personal-projects.steps.ts`. Earlier in this phase Biome's
    `noArrayIndexKey` fired on `key={index}` in `personal-projects-content.tsx`; the fix replaced the index with
    `project.title`, and it changed the `id` on the same element at the same time —
    `id={`project-${index}`}` became `id={`project-${project.title}`}`. The `key` change was right and is kept: an index
    reused as identity makes React carry the wrong card's state across a search filter. The `id` change was wrong,
    because the behaviour corpus addresses cards positionally and an element id is observable DOM in a way a React key
    never is. The element now carries a positional `id` and a title-based `key`, with a comment saying why the two
    deliberately disagree, and all 53 scenarios pass. Both attributes were audited across the whole ported tree
    afterwards: this is the only dynamic `id` in `src/`, and every other `key` the lint pass rewrote changes nothing a
    test or a browser can observe.

- [x] [AI] Import `EnvRecord` into `apps/wahidyankf-www/tests/bdd/port-resolver.unit.test.ts` from
      `@/features/env/core/tier-env` as a type-only import, rather than from `@/features/env/core/port-resolver` —
      acceptance: `npx nx run wahidyankf-www:typecheck` and `npx nx run wahidyankf-www:build` both exit 0, and
      `rg -n '^import type' apps/wahidyankf-www/tests/bdd/port-resolver.unit.test.ts` finds the type-only form. **This
      item was added during execution.** The item that copied this test repointed its module import at the resolver,
      which is where `resolvePort` lives; `EnvRecord` was never there. In `ose-public` the test imported both from the
      library barrel, and the inlining split that barrel into two modules — `port-resolver.ts` declares the
      `process.env`-shaped type privately as `EnvLike` precisely so it can import nothing at all, and `tier-env.ts` is
      what exports it under the name the test uses.
  - Note, 2026-09-01: both targets pass. The import is `import type` rather than a plain import so it is erased at
    compile time and this test does not pull `tier-env.ts`, or the `dotenv` beneath it, into a suite whose entire point
    is a module that depends on nothing. The alternative — re-exporting the type from `port-resolver.ts` — was rejected
    because it would give one type two exported names inside one feature.
  - Note, 2026-09-01: this is the first defect in the phase that neither `typecheck` nor `lint` reported when it was
    introduced, and the reason is ordering rather than coverage. `typecheck` includes `tests/` and would have caught it,
    but it was last run at the end of Port the Source, before Bind the Corpus created the file. `next build` runs its
    own type check and is what surfaced it — at `git push`, inside the pre-push `test:quick`, four items after the
    mistake was made. The lesson is not that a check was missing; it is that a check already passed is not evidence
    about files created after it ran.

### New Behaviour: CV Export at the Filesystem Boundary

- [x] [AI] Create `specs/apps/wahidyankf-www/behaviours/cv-export.feature` with the title `Feature: CV export` and the
      first scenario "Generating the CV writes a PDF to the local filesystem" — acceptance:
      `rg -n '^Feature: CV export' specs/apps/wahidyankf-www/behaviours/cv-export.feature` finds the title and
      `rg -n 'Generating the CV writes a PDF' specs/apps/wahidyankf-www/behaviours/cv-export.feature` finds the
      scenario. These two scenarios get their own file rather than joining `cv.feature`, because `prd.md` writes them
      under `Feature: CV export`, and because `cv.feature` is loaded by `tests/bdd/cv.steps.ts` —
      `@amiceli/vitest-cucumber` throws `ScenarioNotCalledError` for any scenario in a loaded feature that the binding
      does not declare, so an integration-only scenario placed there would fail `test:coverage:behaviour`. [AC-3]
  - Note, 2026-09-01: both acceptances hold. The scenario's four step lines are `prd.md`'s verbatim, and the file
    carries the `As / I want / So that` narrative and a `@integration` tag to match the house style of the eleven files
    beside it. One structural choice was made and then reversed: the `Given` was first lifted into a `Background`, which
    is what most of this corpus does, and that was withdrawn before the file was final. The second scenario, added three
    items below, carries its own distinct `Given` about a missing output directory; with a `Background` in place that
    scenario would have had two primary `Given` steps, which is exactly what the Phase 2 step-cardinality rule flags.
    Each scenario keeping its own single `Given` leaves the file at `Given=1 When=1 Then=1` per scenario and needs no
    `Background` at all.
- [x] [AI] RED — **Gherkin (binds)** "Generating the CV writes a PDF to the local filesystem". Add the failing test to
      `apps/wahidyankf-www/tests/integration/cv-pdf.integration.test.ts`. Verify with
      `npx nx run wahidyankf-www:test:integration` — the new test fails because no integration binding exists yet.
      Scenario: [AC-3]
  - Note, 2026-09-01: RED reached, and reached for the stated reason. `test:integration` reports
    `Tests 1 failed | 3 passed (4)` with `Error: the CV export is not bound to the integration layer yet` raised from
    the `When` step. The three passing steps are the `Given`, which asserts the CV record is non-empty and is genuinely
    true today, and the two assertions after the `When`, which the runner still executes. Making the `When` throw an
    explicit message is what keeps this a real RED: the file could have been written to fail on a missing PDF instead,
    but that failure mode is indistinguishable from an export that ran and silently produced nothing, which is the very
    thing this scenario exists to detect.

```gherkin
Scenario: Generating the CV writes a PDF to the local filesystem
  Given the application CV record contains at least one entry
  When the CV export runs against a writable output directory
  Then a readable PDF file exists at the configured output path
  And the file begins with the PDF header bytes
```

- [x] [AI] GREEN — make the test pass by binding it to the real export in
      `apps/wahidyankf-www/src/features/cv/shell/pdf.ts` against a temporary directory. Verify with
      `npx nx run wahidyankf-www:test:integration` — exits 0. [AC-3]
  - Note, 2026-09-01: `test:integration` exits 0 with all four steps passing, and the RED sentinel string appears
    nowhere in the file. The binding calls the same three functions the `generate:cv-pdf` script calls —
    `buildCvPdfDocument`, `renderCvPdf`, then a pipe to a write stream — against a `mkdtemp` directory instead of
    `public/`, so it exercises the real export rather than a reimplementation of it. One detail decides whether this
    scenario means anything: `renderCvPdf` has already called `end()` on the document when it returns, so the bytes are
    still draining into the file when the `When` step would otherwise finish. The step awaits the write stream's
    `finish` event, and without that await the `Then` reads a partial file or an empty one, which would make the
    assertion flaky in exactly the direction that looks like a pass.
- [x] [AI] REFACTOR — extract the temporary-directory setup into a reusable fixture within the same test file. Verify
      with `npx nx run wahidyankf-www:test:integration` — still exits 0. [AC-3]
  - Note, 2026-09-01: still exits 0 with four passing steps, and no assertion changed. `outputFixture()` now creates the
    throwaway directory and returns the directory, the file path inside it, and the `cleanup` that removes it, so the
    scenario holds one variable rather than two and the creation and the removal sit next to each other. That pairing is
    the point of the extraction rather than the line saving: the two were previously eleven lines apart, in different
    steps, which is how a scenario ends up creating a directory it never deletes. The fixture is deliberately a plain
    function in this file rather than a `beforeEach`, because the second scenario added below needs a path in a
    directory that does **not** exist, and a hook that creates one unconditionally would work against it.
- [x] [AI] Add the scenario "Generating the CV reports an unwritable output location" to
      `specs/apps/wahidyankf-www/behaviours/cv-export.feature` — acceptance:
      `rg -n 'reports an unwritable output location' specs/apps/wahidyankf-www/behaviours/cv-export.feature` finds it.
      [AC-3]
  - Note, 2026-09-01: added verbatim from `prd.md`, tagged `@integration` like its sibling, and the file now holds two
    scenarios. Its `Given` names a directory that does not exist, which is the opposite precondition to the first
    scenario's — the reason the earlier item did not lift either one into a `Background`.
- [x] [AI] RED — **Gherkin (binds)** "Generating the CV reports an unwritable output location". Add the failing test to
      `apps/wahidyankf-www/tests/integration/cv-pdf.integration.test.ts`. Verify with
      `npx nx run wahidyankf-www:test:integration` — the new test fails. Scenario: [AC-3]
  - Note, 2026-09-01: RED reached — `Tests 2 failed | 6 passed (8)`, the two failures both in the new scenario and the
    first scenario's four steps still green. The `When` throws the sentinel, and the `Then` fails alongside it because
    nothing has been assigned to `failure`. Adding the scenario to the feature file had already turned the suite red
    before a line of binding was written, with
    `ScenarioNotCalledError: Scenario: Generating the CV reports an unwritable output location was not called`; that is
    the runner reporting an unbound scenario rather than a failing behaviour, so it was not treated as this item's RED.
    The `Given` is worth noting: it creates a temp directory through the fixture and then removes it, rather than naming
    a path that never existed, so the parent is a real location this process owns and the only thing absent is the
    directory the export was pointed at.

```gherkin
Scenario: Generating the CV reports an unwritable output location
  Given the configured CV output directory does not exist
  When the CV export runs
  Then the export fails with a message naming the output path
  And no partial file is left behind
```

- [x] [AI] GREEN — make the test pass, adding the output-path check to
      `apps/wahidyankf-www/src/features/cv/shell/pdf.ts` only if the export does not already fail with that message.
      Verify with `npx nx run wahidyankf-www:test:integration` — exits 0. [AC-3]
  - **Clause disposition, 2026-09-01: not triggered.** The conditional half of this step never fired. `renderCvPdf` and
    `buildCvPdfDocument` write through a stream whose failure the binding already surfaces, so the RED test failed for
    the reason the scenario states rather than for a missing output-path check, and GREEN was reached without adding
    one. The delivered `cv-export.feature` scenario about a failing export names the output path from the error the
    stream raises, not from a check this clause would have added.
  - Note, 2026-09-01: `test:integration` exits 0 with all eight steps passing, and `pdf.ts` was **not** changed, because
    the item's condition was tested rather than assumed. Writing into a directory that does not exist makes the write
    stream emit `ENOENT: no such file or directory, open '<path>'`, which already names the output path, so both `Then`
    steps are satisfied by the platform's own error and a hand-written check in `pdf.ts` would only restate it. Adding
    one would also have moved a filesystem concern into a module whose job is to render a document and which never opens
    a file. The `no partial file is left behind` assertion passes for the same reason the error is raised: the open
    failed, so nothing was ever created to leave behind. The step captures the rejection rather than letting it
    propagate, since `events.once` rejects when the stream emits `error` in place of the awaited event, and an uncaught
    rejection would fail the `When` before either `Then` could examine the message.
- [x] [AI] REFACTOR — with both tests passing, remove any duplication between the two cases. Verify with
      `npx nx run wahidyankf-www:test:integration` — still exits 0. [AC-3]
  - Note, 2026-09-01: still exits 0 with all eight steps passing, and `lint` is clean across all 73 files. The
    duplication was the export itself: both `When` steps built the document, rendered it, opened a write stream, piped,
    and awaited `finish`, differing only in the destination and in whether a rejection was expected.
    `exportCvPdfTo(file)` now holds those three calls and returns the error or `undefined`, so each scenario's `When` is
    one line. Returning the failure rather than throwing it is what lets both cases share the helper: the unwritable
    case needs the error as a value it can assert the message of, and a helper that threw would have forced the caller
    to reintroduce the `try` this refactor removed. The success case asserts the returned value is `undefined`, so a
    write that silently failed cannot pass as a write that succeeded.
- [x] [AI] Confirm `cv-export.feature` is bound exactly once, from the integration layer — acceptance:
      `rg -l 'behaviours/cv-export.feature' apps/wahidyankf-www` lists only
      `apps/wahidyankf-www/tests/integration/cv-pdf.integration.test.ts`, and
      `rg -n 'cv-export' apps/wahidyankf-www/tests/bdd` finds nothing. The corpus now holds twelve feature files and 55
      scenarios: eleven files and 53 scenarios bound from `tests/bdd/`, and this one bound from `tests/integration/`.
      [AC-4]
  - Note, 2026-09-01: `rg -l 'behaviours/cv-export.feature' apps/wahidyankf-www` lists exactly one file,
    `tests/integration/cv-pdf.integration.test.ts`, and `rg -n 'cv-export' apps/wahidyankf-www/tests/bdd` finds nothing.
    The corpus totals the item predicts also hold: twelve `.feature` files and 55 scenarios. The single-binding rule
    matters more here than for the other eleven, because a second binding under `tests/bdd/` would put these two
    scenarios in the `behaviour` project, where they would run under jsdom with no setup file and write real PDFs on
    every quick gate.
- [x] [AI] Add `cv-export.feature` to the `## Directory Map` in `specs/apps/wahidyankf-www/behaviours/README.md`, noting
      that it binds at the integration layer rather than under `tests/bdd/` — acceptance: `npm run check:markdown-links`
      exits 0 and the map lists twelve entries.
  - Note, 2026-09-01: the entry is added, noting the integration binding, and the map now lists twelve. Two counts in
    the prose above the map were false the moment the entry landed and were repaired with it — "Eleven feature files
    carry 53 scenarios" and "Nine of the eleven" — plus a paragraph saying why this one file binds outside `tests/bdd/`.
    The item's acceptance names only the map entry and the link check, so neither sentence was in its scope; leaving
    them would have made the document's own summary disagree with the list beneath it.
  - Note, 2026-09-01: `npm run check:markdown-links` failed on the first attempt, and not because of this edit. Both
    errors were at `apps/wahidyankf-www/README.md:26`, in the stale README the port carried in, pointing at
    `../../specs/apps/wahidyankf/behaviours/wahidyankf-www/README.md` and `../wahidyankf-www-fe-e2e/README.md` — the
    source repository's corpus path and the old E2E project name. The item that replaces that README sits in Close the
    Migration, several items later, so this acceptance could not pass in the order written. It was executed out of
    order, exactly as the Phase 2 ordering defect was; see the note on that item below. The check exits 0 now.

### Reach the Coverage Floor

- [x] [AI] Run `npx nx run wahidyankf-www:test:coverage:unit` and record the starting line percentage in
      `evidence/phase-3-measurements.md`, converting that file's entry in the `## Directory Map` of `evidence/README.md`
      to a relative link in this same item, since this is where the file first exists — acceptance: the command reports
      a percentage and the map entry is a link; it is expected to fail against the 99% threshold at this point, and that
      failure is the measurement.
  - Note, 2026-09-01: **97.99% lines**, 1.01 points under the floor, recorded in `evidence/phase-3-measurements.md`
    along with every uncovered line and why each one is uncovered. The `evidence/README.md` map entry is now a relative
    link and `npm run check:markdown-links` exits 0. Two things the measurement showed that the checklist did not
    predict. First, `robots.ts` and `sitemap.ts` are already at 100%, reached through the `static-filterable-routes`
    scenarios rather than by any test of their own; their tests are still written three items below, because a module
    covered only as a side effect of a crawler scenario has nothing that fails when its own output changes. Second,
    `src/app/favicon.ico` and `src/features/ui/shell/index.ts` appear in the table at `0 | 0 | 0 | 0` — the `src/**`
    denominator is a path glob, not a language filter — and neither carries a statement, so neither moves the
    percentage. That is recorded in the evidence so a later reader does not read them as uncovered code and try to test
    an icon.

The nine unit tests named in [file impact](tech-docs/file-impact.md) are authored one feature area at a time, matching
the mitigation this phase states for its own size: a failure then localizes to an area rather than to the phase. Each
item below is complete when its own files exist and the whole unit suite is still green, so a bisect lands on one area.

- [x] [AI] Author the three route-metadata unit tests — `apps/wahidyankf-www/src/app/head.unit.test.tsx`,
      `app/robots.unit.test.ts`, and `app/sitemap.unit.test.ts` — acceptance: all three files exist as co-located
      siblings and `npx nx run wahidyankf-www:test:unit` exits 0.
  - Note, 2026-09-01: all three files exist beside the modules they test and `test:unit` exits 0 at
    `Test Files 21 passed (21)`, `Tests 131 passed (131)`. `head.tsx` took two attempts. The first assumed the `<link>`
    elements would be in the render container; React 19 hoists them into `document.head` on its own, which is also why
    the component works at all with no Next.js around it, so the assertions read `document.head`. The second attempt
    then removed those links in an `afterEach` and failed with `Cannot read properties of null (reading 'removeChild')`
    — the nodes belong to React, the setup file already runs `afterEach(cleanup)` to unmount the tree that owns them,
    and deleting them by hand leaves React's own removal with nothing to remove. The hand cleanup is gone and a comment
    records why it must not come back. The `robots.ts` and `sitemap.ts` tests assert what the directives say rather than
    that the routes exist, which is the gap the coverage measurement showed: both modules were already at 100% through
    the crawler scenarios, with nothing that fails when the allow rule, the sitemap URL, or a route priority changes.
- [x] [AI] Author the two composition-root unit tests — `apps/wahidyankf-www/src/env.unit.test.ts` and
      `src/env-loader.unit.test.ts` — acceptance: both files exist as co-located siblings and
      `npx nx run wahidyankf-www:test:unit` exits 0. `env-loader.ts` calls `loadTierEnv()` at import time, so its test
      asserts the re-export surface and the single call rather than re-testing the loader, whose own scenarios are bound
      under `tests/bdd/`.
  - Note, 2026-09-01: both files exist beside their modules and `test:unit` exits 0 at 23 files and 135 tests.
    `env-loader.unit.test.ts` follows the item's instruction exactly and does not re-test the loader: it mocks
    `./features/env/core/tier-env` and asserts two things only — that importing the module calls `loadTierEnv()` exactly
    once, and that it re-exports the three names `next.config.ts` reaches for. The mock is not a convenience, it is what
    makes the first assertion possible at all, because the call happens at import time and there is nothing left to
    observe once the module has been evaluated. `env.unit.test.ts` asserts that `createEnv()` validates and that the
    resulting object declares no keys, which is the same claim `.env.example` makes in prose — so a variable added
    without a matching entry in that file fails here.
- [x] [AI] Author the CV shell unit test `apps/wahidyankf-www/src/features/cv/shell/cv-content.unit.test.tsx` —
      acceptance: the file exists as a co-located sibling and `npx nx run wahidyankf-www:test:unit` exits 0.
  - Note, 2026-09-01: the file exists beside `cv-content.tsx` and `test:unit` exits 0 at 24 files and 139 tests. The
    measurement said what to write: the module's only uncovered lines were `534`, the recent-only toggle's click
    handler, and `437-440`, the five-year window helper that handler reaches. No scenario in `cv.feature` clicks that
    control, so the four tests here drive it — off, on, off again, and the effect on what is rendered. The last one
    asserts that filtering shows no more entries than not filtering, rather than a specific count, because a count would
    pin the test to today's CV record and fail the next time a job is added; the direction of the change is what the
    filter actually promises.
- [x] [AI] Author the home shell unit test `apps/wahidyankf-www/src/features/home/shell/home-content.unit.test.tsx` —
      acceptance: the file exists as a co-located sibling and `npx nx run wahidyankf-www:test:unit` exits 0.
  - Note, 2026-09-01: the file exists beside `home-content.tsx` and `test:unit` exits 0 at 25 files and 142 tests. The
    single uncovered line was `173`, the click handler on an AI-related skill pill. `home.feature` asserts the skills
    card carries three subsections; the AI-related list is a fourth, rendered only when the CV record supplies AI
    skills, so no scenario reaches it. The three tests here render the subsection, click a pill, and assert the CV route
    it pushes. Two choices keep them independent of the record's contents: the pill is selected positionally from the
    element after the heading rather than by a skill name, and the encoding test round-trips the pushed URL through
    `URL` and compares the parsed `search` parameter against the pill's own label, so a skill containing a space or an
    ampersand is checked rather than assumed.
- [x] [AI] Author the two personal-projects unit tests —
      `apps/wahidyankf-www/src/features/personal-projects/core/projects.unit.test.ts` and
      `features/personal-projects/shell/personal-projects-content.unit.test.tsx` — acceptance: both files exist as
      co-located siblings and `npx nx run wahidyankf-www:test:unit` exits 0.
  - Note, 2026-09-01: both files exist beside their modules and `test:unit` exits 0 at 27 files and 159 tests. Neither
    module was short of line coverage — both were already at 100% — so the tests were written against what the
    measurement showed was actually untested: the `filterProjects` field list, asserted field by field so removing one
    from the wrapper fails here rather than quietly narrowing what the projects page can find, and the empty-term branch
    of the search URL, which no scenario reaches because none of them clears the box. That second one is worth stating:
    clearing the search must route to `/personal-projects` and not `/personal-projects?search=`, which are different
    URLs, and the second would leave the address bar advertising a filter the page is no longer applying.
- [x] [AI] Set `testTimeout` on each of the three projects in `apps/wahidyankf-www/vitest.config.ts` through one named
      constant — acceptance: `npx nx run wahidyankf-www:typecheck` exits 0,
      `rg -c 'TEST_TIMEOUT_MS' apps/wahidyankf-www/vitest.config.ts` reports the constant plus one use per project, and
      ten consecutive `npx vitest --config vitest.config.ts run --project unit` runs all report `159 passed` with no
      `timed out in` line. **This item was added during execution.** The `cv-content` test written three items above
      turned the unit suite intermittently red — green in isolation every time, red roughly one run in four under a full
      parallel run — with `Error: Test timed out in 5000ms` on the first `CvContent` render.
  - Note, 2026-09-01: ten consecutive runs pass and `typecheck` is clean. The cause is that Vitest measures its timeout
    from when a test starts rather than from when its worker is ready, so the first test in a file pays that file's
    jsdom setup and module import out of its own budget; `CvContent` renders the entire CV in one component, and under
    contention that crossed five seconds. The limit is raised rather than the test narrowed, because the render is the
    behaviour under test — a version that rendered less would be faster and would stop measuring the thing. The value is
    set on each project rather than once at the root, and the first attempt did exactly that and kept failing: this file
    already documents that Vitest projects do not inherit plugins from the root config, and the same is true of a
    project's other `test` options. That inheritance assumption cost two more failing runs before it was checked rather
    than assumed.
- [x] [AI] Run `npx nx run wahidyankf-www:test:coverage:unit` — acceptance: exits 0 at the 99% line threshold, with no
      threshold lowered and no broad exclusion added. [AC-2]
  - Note, 2026-09-01: exits 0 at **100% lines** — 99.57% statements, 92.3% branches, 100% functions — against the
    unchanged `--coverage.thresholds.lines=99` on the target. No threshold was lowered and nothing was added to
    `coverage.exclude`; the 97.99% starting point closed entirely through the seven test files the four items above
    wrote. The two zero-row entries noted in the starting measurement, `src/app/favicon.ico` and
    `src/features/ui/shell/index.ts`, are still in the table at `0 | 0 | 0 | 0` and still carry no statements, which is
    why 100% lines and a pair of visible zeroes are consistent rather than contradictory.
- [x] [AI] Run `npx nx run wahidyankf-www:test:coverage:integration` — acceptance: exits 0 at the 99% line threshold,
      and its coverage table lists exactly the two PDF modules named in [technical design](tech-docs/README.md) and no
      other file, proving the `--coverage.include` flag on the target reached the run. [AC-3]
  - Note, 2026-09-01: exits 0 at **100% lines**, and the table lists exactly two files — `src/features/cv/core/pdf.ts`
    and `src/features/cv/shell/pdf.ts` — with no third. That second half is the part worth having checked. `coverage` is
    a root-level option in Vitest 4 that a project config cannot set, so the `coverage.include: ["src/**"]` in
    `vitest.config.ts` governs the `integration` project too, and without the
    `--coverage.include='src/features/cv/**/pdf.ts'` flag on the target this run would measure all of `src/**` with two
    tests executing and could not reach 99% at all. The table proves the flag replaced the configured include rather
    than adding to it, which is what the technical design predicted and what a two-row table now confirms rather than
    assumes.
- [x] [AI] Run `npx nx run wahidyankf-www:static-routes:validation` and record its wall-clock duration in
      `evidence/phase-3-measurements.md` — acceptance: exits 0, and the recorded duration is the evidence a later
      decision about its place in `test:quick` would need. It is a measurement rather than a lesson, so it belongs where
      a later plan can read it back rather than in the log Phase 7 drains.
  - Note, 2026-09-01: exits 0 and reports
    `Verified static build output for /, /cv, /personal-projects, /robots.txt, /sitemap.xml.` Three runs with
    `--skip-nx-cache`, so none of them rode a cached result: **5.14s, 4.38s, 4.24s**. Recorded in
    `evidence/phase-3-measurements.md` with the breakdown — nearly all of it is the `next build` the target runs first,
    and the validator that reads the build output is the cheap half. It is recorded and not acted on, which is what the
    item asks: `test:quick` still declares this target in `dependsOn`, and no item in this plan moves it.

### Close the Migration

- [x] [AI] Edit the `test:scheduled` script in root `package.json` — and only that script — to
      `nx run badakmini-cli:test:quick && nx run wahidyankf-www:test:quick && nx run badakmini-cli:test:coverage:integration && nx run wahidyankf-www:test:coverage:integration && nx run badakmini-cli:test:e2e`,
      preserving the quick-then-integration-then-E2E order
      [workspace commands](../../../repo-governance/development/workspace-commands.md#build-and-test) states —
      acceptance: `node -e "console.log(require('./package.json').scripts['test:scheduled'])"` prints exactly that
      string. No other script is touched: `build`, `lint`, `typecheck`, `test:unit`, `test:quick`, `test:coverage`,
      `test:behaviour`, `test:integration`, and `test:e2e` are all `nx run-many -t <target>` and pick up a new project
      with no edit at all. `test:scheduled` is the only script that names projects, which is why it is the only one that
      has to change.
  - Note, 2026-09-01: the script prints exactly the specified string, and `git diff --stat package.json` reports one
    insertion and one deletion, so the "and only that script" half is measured rather than asserted. The edit was made
    by parsing the manifest, comparing every script against its previous value, and refusing to write unless the changed
    set was exactly `["test:scheduled"]` — which matters because a hand edit to a manifest is the kind of change where a
    neighbouring line is easy to disturb and nothing downstream reports it. The item's reasoning holds: every other
    script is `nx run-many -t <target>` and already covers the new project without an edit.
- [x] [AI] Author `apps/wahidyankf-www/README.md` naming its corpus path, its three adapters, its targets, and its
      coverage denominators, linking to `specs/apps/wahidyankf-www/architecture.md`, then run
      `git add -N apps/wahidyankf-www/README.md` before checking — acceptance: `npm run check:markdown-links` exits 0
      and `rg -n 'architecture.md' apps/wahidyankf-www/README.md` finds the backlink. Without the intent-to-add the
      check does not see this file at all.
  - Note, 2026-09-01: **executed out of order**, from the `cv-export.feature` indexing item in New Behaviour, because
    that item's acceptance is `npm run check:markdown-links` exiting 0 and this README was the only thing failing it.
    `rg -n 'architecture.md'` finds the backlink and the link check exits 0. The `git add -N` the item calls for was not
    needed: this file arrived with the ported tree and was already tracked by the time the item ran, so the check
    already saw it. The rewrite names the corpus path, the three adapters and what decides which one a feature reaches,
    the targets with `test:quick`'s ordering and why it stops short of `test:coverage`, and both coverage denominators
    including why `test:coverage:unit` runs the `behaviour` project too. What it drops is what the port made false: the
    OSE framing, a `test:specs` target this repository does not define, and the two dead links.
- [x] [AI] Edit the `## Current Applications` list in `apps/README.md` to index the new application — acceptance:
      `rg -n '\[`wahidyankf-www`\]' apps/README.md` finds the entry with a descriptive relative link to
      `wahidyankf-www/README.md`. The E2E project is not indexed here: it does not exist yet and gets its own item in
      Phase 5.
  - Note, 2026-09-01: the entry is present with a relative link to `wahidyankf-www/README.md`, and
    `npm run check:markdown-links` exits 0. The description names what distinguishes this application rather than
    restating its stack: it holds the repository's single authoritative CV record, which is the claim Phase 4 makes true
    by deleting `cv/`. The E2E project is not listed, as the item directs — it does not exist yet.
- [x] [AI] Amend `repo-governance/development/testing-policy/tooling.md` with the two deviations this phase makes
      certain. First, the language target: `apps/wahidyankf-www/tsconfig.json` sets `module` to `esnext`,
      `moduleResolution` to `bundler`, and `target` to `ES2017`, which is not the "CommonJS-compatible Node output"
      `code-style-policy.md` sets as the language target, and `strict` stays true. It reaches that state by extending
      `tsconfig.base.json` and overriding those three options on top of it. Second, the linter's boundary: Biome runs as
      a linter only, because the root `biome.json` sets `formatter.enabled` and `assist.enabled` to `false` and Prettier
      stays the formatting source of truth `code-style-policy.md` names. Biome v2 defaults `formatter.indentStyle` to
      tab while Prettier here uses two spaces, so an enabled Biome formatter would fail `lint` on every ported file and,
      if "fixed" with `biome check --write`, would retab them and break `npm run format:check` repository-wide —
      acceptance: `rg -n 'moduleResolution|bundler' repo-governance/development/testing-policy/tooling.md` finds the
      language-target deviation, `rg -n 'formatter' repo-governance/development/testing-policy/tooling.md` finds the
      Biome boundary, and `npm run check:governance` exits 0, the document being 111 words before this edit and governed
      by the 750-word cap. Both edits are certain rather than conditional: each is known before execution starts, and
      [technical design](tech-docs/README.md#toolchain-conformance-and-its-fallback) records why Next 16 leaves no
      alternative for the first and why Prettier and Biome cannot both format for the second. The E2E project carries
      the same language-target deviation but does not exist yet; Phase 5 extends the first sentence to it in the phase
      that creates it. A toolchain fallback that fires adds its own paragraph on top.
  - Note, 2026-09-01: both deviations are recorded under a new `## Recorded Deviations` heading,
    `rg -n 'moduleResolution|bundler'` and `rg -n 'formatter'` each find their paragraph, and
    `npm run check:governance`, `npm run check:markdown-links`, and `npm run format:check` all exit 0. The document went
    from 111 words to 235, well inside the 750-word cap the item names. Each paragraph states the deviation and the
    reason it is forced rather than chosen — Next 16 resolving its own imports as ESM through a bundler for the first,
    and Biome v2 defaulting `formatter.indentStyle` to tab against Prettier's two spaces for the second. That second
    reason is the one worth having written down: with the Biome formatter enabled, `biome check --write` would retab
    every ported file and break `npm run format:check` repository-wide, so the disabled formatter is load-bearing rather
    than a preference. The E2E project is not mentioned; it does not exist yet and Phase 5 extends the first paragraph
    to it. [AC-9]
- [x] [AI] Add the new project's narrower runs to the `Narrower runs` fenced block in
      `repo-governance/development/workspace-commands.md` — `npx nx run wahidyankf-www:test:unit`, `:test:integration`,
      `:test:coverage:unit`, `:test:coverage:integration`, `:test:coverage:behaviour`, `:test:coverage`, `:test:quick`,
      `:static-routes:validation`, and `:generate:cv-pdf` — acceptance:
      `rg -c 'wahidyankf-www' repo-governance/development/workspace-commands.md` reports at least nine matches and
      `npm run check:governance` exits 0, the document being 470 words before this edit and governed by the 750-word
      cap. That document says it is the canonical command reference and that a summary is what drifts, so a project's
      worth of new commands has to reach it.
  - Note, 2026-09-01: nine `wahidyankf-www` lines added to the fenced block and `rg -c 'wahidyankf-www'` reports exactly
    9; `npm run check:governance` and `npm run check:workflows` both exit 0. The nine are the ones the item names,
    listed in the same order the `badakmini-cli` lines above them use — the test targets first, then the two this
    project has that the CLI does not, `static-routes:validation` and `generate:cv-pdf`. They sit above the
    `nx affected` line rather than below it, keeping that line last where it reads as the summary of the block rather
    than as one more per-project command.
- [x] [AI] Edit root `README.md` to describe the workspace as holding both the Go CLI and the Next.js site — acceptance:
      `rg -n 'wahidyankf-www' README.md` finds the description.
  - Note, 2026-09-01: `rg -n 'wahidyankf-www' README.md` finds the description and `npm run check:markdown-links`
    exits 0. The sentence sits in the Nx Workspace section directly after the Badak Mini paragraph, where the reader has
    just been told what `apps/` is for, and it gives the reason the site lives here rather than only the fact that it
    does — the practice workspace and its public face under one set of rules, which is the user story `prd.md` opens
    with. It also states what the two applications have in common, a Gherkin corpus and the 99% floor, since a reader
    arriving at a Go CLI and a Next.js site in one workspace would otherwise have no reason to expect them held to the
    same gates.
- [x] [AI] Run `rg -n '@open-sharia-enterprise/' --hidden --glob '!node_modules' --glob '!package-lock.json' .` —
      acceptance: no match outside this plan's own documents, which name the specifiers in order to record what was
      removed. The pattern is the specifier form `[AC-6]` states rather than the bare organisation name, because one
      bare occurrence in this repository is owner CV prose that survives deliberately: `cv/linkedin-projects.md` records
      replacing an outdated OSE link with `ose-public` when those entries are ready to publish, and Phase 4 copies that
      file into `apps/wahidyankf-www/docs/` byte-identical and asserts its digest. A sweep on the bare token would
      therefore fail on prose no item is allowed to edit. [AC-6]
  - Note, 2026-09-01: every match is inside `plans/in-progress/wahidyankf-www-migration/`, which is what the acceptance
    allows — `prd.md`, the three tech-docs, `delivery.md`, and the plan README, all naming the specifiers in order to
    record removing them. No match in `apps/`, `specs/`, `scripts/`, or any manifest. The item's reasoning about the
    pattern held in practice: `cv/linkedin-projects.md` does carry the bare organisation name in owner prose, and a
    sweep on that token rather than on the specifier form would have failed on a file no item is permitted to edit.
    [AC-6]
- [x] [AI] Run
      `rg -n 'rhino-cli|dotnet|\.fsproj|F#' --hidden --glob '!node_modules' apps/wahidyankf-www* specs scripts .github package.json nx.json`
      — acceptance, **amended during execution**: the three tokens `[AC-7]` names — `rhino-cli`, `dotnet`, and `.fsproj`
      — find no match, and `F#` finds no match once `**/personal-projects/core/projects.ts` is excluded. The criterion
      first read "no match" for the whole pattern, which the delivered tree does not satisfy and must not: that one file
      lists `"F#"` among the languages two of the owner's projects are written in, and it is now the repository's single
      authoritative CV record. Editing true CV data to satisfy a sweep was the alternative; the criterion was changed
      instead. The sweep covers what this plan writes — the ported projects, the corpus, the wrapper script, the
      workflows, and the root manifests — and is deliberately not repository-wide, because two occurrences elsewhere are
      correct and stay: `apps/badakmini-cli/README.md` names `rhino-cli` in the provenance sentence explaining where
      Badak Mini's command grammar comes from, and `plans/done/2026-08-23__badakmini-layered-bdd/brd.md` names F# in a
      non-goal of an archived plan, which
      [folder structure](../../../repo-governance/conventions/plans-organization-policy/folder-structure.md) keeps as an
      accurate historical record rather than something a later plan rewrites. Neither is a toolchain reference, and
      deleting either to make a repository-wide sweep pass would destroy a true record. `F#` is in the pattern because
      the ported wrapper contract test names it in a test title, which the other three tokens do not match. [AC-7]
  - Note, 2026-09-01: the pattern as written finds **two matches**, and both are owner CV data rather than toolchain
    references: `apps/wahidyankf-www/src/features/personal-projects/core/projects.ts` lists `"F#"` in the
    `programmingLanguages` of the OSE and OrganicLever entries, which is a true statement about what those projects are
    written in. Removing it to satisfy a sweep would falsify the CV record this application is now the single
    authoritative copy of, so it stays. The check was therefore split rather than forced.
    `rg -n 'rhino-cli|dotnet|\.fsproj'` over the same paths — the three tokens `[AC-7]` actually names — finds nothing
    at all. `rg -n 'F#'` over the same paths with `--glob '!**/personal-projects/core/projects.ts'` also finds nothing,
    so every F# reference this plan set out to remove is gone and the only survivors are two entries in a list of
    languages. This is the same false-positive shape the sibling `@open-sharia-enterprise/` item anticipated and wrote
    its pattern around; this item's pattern did not get the same treatment, and the exclusion is recorded here rather
    than the data being edited to fit. [AC-7]
- [x] [AI] Run `rg --files apps specs scripts | rg '/[A-Z][^/]*\.tsx?$'` and
      `rg -n 'shell/Navigation|shell/CvContent|shell/HomeContent|shell/PersonalProjectsContent|shell/SearchSection' apps`
      — acceptance: neither finds anything. `rg --files` is used rather than `find` because it honours `.gitignore` and
      so never walks a workspace `node_modules`. This is the only check in the plan that would catch a renamed-module
      reference missed in either `src/` or `tests/bdd/`; the owner's volume is case-insensitive, so a stale path
      resolves locally and surfaces first on the Linux CI runner.
  - Note, 2026-09-01: neither finds anything. This is the check the item calls the only one that would catch a
    renamed-module reference missed in either `src/` or `tests/bdd/`, and it earns that description: the owner's volume
    is case-insensitive, so all fourteen stale PascalCase specifiers found earlier in this phase resolved perfectly well
    locally and would have failed only on a case-sensitive CI filesystem. `rg --files` also honours `.gitignore`, so
    neither sweep walked a workspace `node_modules`.
- [x] [AI] Reformat the ported application tree to this repository's Prettier configuration — run `npm run format` —
      acceptance: `npm run format:check` exits 0. Everything this phase copied arrives formatted to
      `$SRC/.prettierrc.json` — `printWidth: 120`, `proseWrap: "preserve"`, and `prettier-plugin-tailwindcss` class
      ordering — and none of those three is set here, so most of the manifests, configuration files, inlined modules,
      stylesheets, `src/` sources, scripts, and `tests/bdd/` bindings this phase landed fail `prettier --check`
      unchanged. [Technical design](tech-docs/README.md#toolchain-conformance-and-its-fallback) records the measured
      counts and why the resulting diff is a formatting change and not a content change. The item sits last in the phase
      for the same reason the Phase 1 one does: after every copy, rename, and repoint, so nothing is reformatted twice,
      and before the gate, so the gate's `format:check` does not depend on a step that has not run. `npm run format` is
      `prettier --write .`, so it sweeps the whole tree, but Prettier has no parser for `.feature` and skips those files
      in a directory run, which is what leaves `specs/apps/wahidyankf-www/behaviours/` byte-identical and keeps the
      Phase 2 `diff` against the source corpus and the scenario counts above valid. It sweeps `cv/` as well — those
      files are already formatted to this repository's configuration, `npx prettier --check cv/` passes today, and this
      run therefore leaves them byte-identical, which is what keeps the Phase 4 SHA-256 comparisons between `cv/` and
      `apps/wahidyankf-www/docs/` valid. Nothing is added to an ignore file: `apps/wahidyankf-www/**` is held to the
      repository's formatting standard like every other tracked path, and [file impact](tech-docs/file-impact.md) rules
      out introducing a root `.prettierignore` at all.
  - Note, 2026-09-01: `npm run format:check` exits 0. The reformat itself was a no-op, and that is a departure from what
    the item expected rather than a sign it was skipped: the item is written for a tree that arrives formatted to
    `$SRC`'s `printWidth: 120` and is reformatted once at the end of the phase, but every file this phase wrote or
    copied was run through `npx prettier --write` as it landed, so `npm run format` reports every file unchanged and has
    nothing left to do. The sweep was still run in full rather than declared unnecessary, because "I formatted as I
    went" is a claim about memory and `format:check` is a measurement. Two of the item's own predictions were verified
    against the result. The eleven ported `.feature` files are byte-identical across the run — SHA-256 of every file in
    `specs/apps/wahidyankf-www/behaviours/` taken before and after, with an empty `diff` — so the Phase 2 comparisons
    against the source corpus and the scenario counts still hold. And `npx prettier --check cv/` exits 0, leaving those
    files untouched, which is what keeps the Phase 4 digest comparisons between `cv/` and `apps/wahidyankf-www/docs/`
    valid. Nothing was added to any ignore file.

### Phase 3 Gate

> Every check below passes before Phase 4 begins. A failure is fixed inside Phase 3. This gate is the 99% floor: it does
> not close while any coverage target is below threshold or any required layer is missing.

- [x] [AI] Run `npx nx run wahidyankf-www:test:quick` — acceptance: exits 0, running `typecheck`, `lint`, `test:unit`,
      `test:coverage:unit`, then `test:coverage:behaviour` in that order, with `static-routes:validation` satisfied as
      its dependency. It does not run `test:coverage:integration`, because pre-push invokes `test:quick` and the
      integration layer touches the real filesystem. [AC-2] [AC-4]
  - Note, 2026-09-01: exits 0, running `static-routes:validation` first through `dependsOn`, then `typecheck`, `lint`,
    `test:unit`, `test:coverage:unit`, and `test:coverage:behaviour` in order.
- [x] [AI] Run `npx nx run wahidyankf-www:test:integration` — acceptance: exits 0 on both new scenarios. [AC-3]
  - Note, 2026-09-01: exits 0 on both CV export scenarios.
- [x] [AI] Run `npm run test:quick` — acceptance: exits 0 across both projects.
  - Note, 2026-09-01: exits 0 across `badakmini-cli` and `wahidyankf-www`.
- [x] [AI] Run `npm run format:check` — acceptance: exits 0.
  - Note, 2026-09-01: exits 0.
- [x] [AI] Run `npm run check:markdown-links` — acceptance: exits 0.
  - Note, 2026-09-01: exits 0.
- [x] [AI] Run `npm audit --audit-level=low` — acceptance: exits 0 with the application's dependency tree installed.
  - Note, 2026-09-01: exits 0 with the full application tree installed. Two inherited pins had to be raised earlier in
    the plan to reach this — `tsx` in Phase 1 and `next` in Phase 3 — both caught at the moment the pin was written
    rather than here, which is what the Phase 1 learning asked for.
- [x] [AI] Confirm `ls libs` lists only `README.md` — acceptance: no library directory was created. [AC-6]
  - Note, 2026-09-01: `ls libs` prints `README.md` and nothing else. The three inlined subjects went into
    `apps/wahidyankf-www/src/features/`, so no library directory was created. [AC-6]
- [x] [AI] Run `npm run check:governance` — acceptance: exits 0 on the two `repo-governance/development/` documents this
      phase edited, `workspace-commands.md` and `testing-policy/tooling.md`, both of which live under the 750-word cap.
      [AC-8]
  - Note, 2026-09-01: exits 0. Three `repo-governance/development/` documents were edited by the end, not two:
    `tooling.md` and `workspace-commands.md` as planned, plus `code-style-policy.md`, which the Rules Propagation run
    below added a reference to. All three are inside the word cap — `tooling.md` at 235 and `code-style-policy.md`
    at 358.
- [x] [AI] Run `git add -A` and then `npm run check:rule-change` — acceptance: it names both `repo-governance/` edits
      and the workflows they trigger, and those workflows are followed rather than dismissed. This is the first phase to
      change a rule path, and the check reads only staged paths, so it is staged first for the same reason the Phase 4
      and Phase 6 gates stage first.
  - Note, 2026-09-01: the check names all three edited rule paths and the workflow they trigger, Rules Propagation. It
    was followed rather than dismissed, which is the half of this acceptance a passing exit status does not establish.
    Working its steps found one real gap: `code-style-policy.md` states "Use strict TypeScript with CommonJS-compatible
    Node output" flatly, and `tooling.md` now records an application that does not meet it, so a reader of the first
    document alone would find a rule one of the two applications violates with nothing pointing at the exception. The
    canonical home stays `tooling.md` — the plan designated it as the deviation register and `[AC-9]` is written against
    it — and `code-style-policy.md` gains one sentence linking there and stating that an unrecorded deviation is a
    defect rather than an exception. Nothing was duplicated and neither rule's requirement changed. The idempotency gate
    did not stop the work: no existing rule named a deviation register at all, and Biome is mentioned in exactly one
    governance document, so there was no overlapping rule to merge. Harness Alignment was not triggered. The workflow's
    own verification, `npm run format:check` and `npm run check:governance`, exits 0.
- [x] [AI] Confirm the toolchain outcome is recorded for all three components `tooling.md` names — acceptance: for each
      of TypeScript 6, Biome, and the project-local ESLint commentary check, either it holds — the root pin is
      TypeScript 6, `lint:biome` runs Biome, and `lint:commentary` runs ESLint with `eslint-plugin-jsdoc` — or
      `learnings.md` names that component, its version, and the error that stopped it. A gate that named only two of the
      three would report conformance for a component nothing had checked. [AC-9]
  - Note, 2026-09-01: all three hold, so no `learnings.md` entry is owed for any of them. TypeScript is pinned at
    `6.0.3` in the root manifest and both `typecheck` and `build` pass on it. `lint:biome` runs `biome check` and
    reports zero findings across 73 files. `lint:commentary` runs `eslint --config eslint.config.mjs src` with
    `eslint-plugin-jsdoc` at `64.3.2` and reports nothing. The third component is the one this gate exists to catch: it
    very nearly became a fallback, because ESLint could not parse a single source file until a parser was added, and the
    conditional item that would have dropped it entirely was read against its actual wording instead. Confirming all
    three separately is what makes that visible rather than letting two green components stand in for the set. [AC-9]
- [x] [AI] Commit and push the phase to `main` — acceptance: `git status --short` is empty.
  - Note, 2026-09-01: pushed as `73b5c32` and `eddc02f`, `dd2118c..eddc02f`, with `git status --short` empty at the
    moment of the push. Split into two commits under the thematic commits policy: the application and governance work,
    then the plan record. This tick is written after that push and rides along with the next commit, which is
    unavoidable — a checkbox asserting a push cannot be inside the push it asserts.

> **Pause Safety**: The application builds, lints, type-checks, and passes unit, integration, and behaviour layers at
> 99%. `libs/` is still empty and no .NET toolchain exists. The repository deliberately holds two CV records at this
> boundary — `data.ts` and `cv/` — which is the expand step of the migration and is closed in Phase 4. No E2E adapter
> exists yet. Safe to stop. Resume with `npx nx run wahidyankf-www:test:quick`.

## Phase 4: Retire `cv/` Into the Application

The application is green and authoritative before anything is deleted. Nothing here is destructive until its replacement
is proven in the same phase.

- [x] [AI] Create `apps/wahidyankf-www/docs/` and copy `cv/cv-raw.md`, `cv/cv-linkedin.md`, and
      `cv/linkedin-projects.md` into it — acceptance: `shasum -a 256 cv/cv-raw.md apps/wahidyankf-www/docs/cv-raw.md`
      prints two identical digests, and the same holds for the other two files. [AC-1]
  - Note, 2026-09-01: all three pairs print identical digests — `cv-raw.md` at `80784d6e…`, `cv-linkedin.md` at
    `6330c2aa…`, and `linkedin-projects.md` at `4321e23c…`. The copies are byte-identical at this point, which is what
    the next item then deliberately breaks for one of them. `cv-ats.md`, `cv-ats.pdf`, `generate-cv-ats-pdf.py`, and
    `cv/README.md` are not copied: the first two are replaced by the application's own generated PDF, the third is the
    generator that produced them, and the fourth's working rules are rewritten into a new README two items below rather
    than carried across.
- [x] [AI] Repair the one broken link the deletion creates, in the destination copy only: the opening sentence of
      `apps/wahidyankf-www/docs/cv-raw.md` names this file as the evidence base for `cv-linkedin.md` and `cv-ats.md`,
      linking each with a relative `./` path. The first target moves into `docs/` alongside it and still resolves; the
      second is deleted outright later in this phase, so `apps/wahidyankf-www/docs/cv-ats.md` never exists. Replace the
      `cv-ats.md` link with a plain inline-code reference to the application CV record `src/features/cv/core/data.ts` —
      not a link, because the check validates link targets and this one sits outside `docs/` — so the sentence still
      names both downstream consumers of the evidence base. [Migration design](tech-docs/migration-design.md) already
      records `data.ts` as what supersedes `cv-ats.md`, and records this sentence as the single intentional divergence
      from the byte-identical copy. Do not touch `cv/cv-raw.md` in the source directory; it is deleted a few items below
      and editing it first would break the digest rehearsal in the Phase 4 gate — acceptance:
      `rg -n 'cv-ats' apps/wahidyankf-www/docs/cv-raw.md` finds nothing,
      `rg -c 'cv-linkedin.md' apps/wahidyankf-www/docs/cv-raw.md` still prints `2` for the two surviving links,
      `rg -n 'features/cv/core/data.ts' apps/wahidyankf-www/docs/cv-raw.md` finds the replacement, and
      `npm run check:markdown-links` exits 0. This item runs after the digest comparison above, which is what proves the
      copy arrived intact before it is amended. The `\bcv/` sweep later in this phase cannot reach this link, because it
      is written `./cv-ats.md` with no directory component; nothing else in the plan would catch it, and Badak Mini
      validates links in every tracked Markdown file. [AC-1]
  - Note, 2026-09-01: the sentence now names `cv-linkedin.md`, still as a relative link, and then the application CV
    record at `src/features/cv/core/data.ts` as what the published CV and its generated PDF are both rendered from (the
    link syntax is not reproduced here, because a relative target quoted inside this document resolves against this
    document and breaks the link check), and `rg -n 'cv-ats' apps/wahidyankf-www/docs/cv-raw.md` finds nothing.
    `data.ts` is inline code rather than a link, as the item directs, because the target sits outside `docs/` and the
    link checker validates what a document points at. The source `cv/cv-raw.md` still digests to `80784d6e…`, unchanged
    — only the destination copy was edited, which is what makes this the single intentional divergence between the two
    rather than an edit to a file the plan is not repairing. The sentence still names both downstream consumers of the
    evidence base; one of them is now a TypeScript module rather than a Markdown file.
- [x] [AI] Author `apps/wahidyankf-www/docs/README.md` carrying the working rules from `cv/README.md` — that `cv-raw.md`
      is the factual evidence base, that public claims stay consistent with it, and that material marked unsuitable for
      public use is not published without owner direction — plus a `## Directory Map` linking the three files, then run
      `git add -N apps/wahidyankf-www/docs` before checking — acceptance:
      `rg -n 'cv-raw.md' apps/wahidyankf-www/docs/README.md` finds the rule and `npm run check:markdown-links` exits 0.
      The three absorbed documents are intent-added alongside the README, because this map is the only thing that links
      to them.
  - Note, 2026-09-01: `rg -n 'cv-raw.md' apps/wahidyankf-www/docs/README.md` finds the rule and
    `npm run check:markdown-links` exits 0, with the three absorbed documents intent-added alongside the README so the
    checker sees them. The three working rules the item names are carried across intact. Two things are new rather than
    copied, both because the directory's situation changed. The README states that the authoritative CV is
    `src/features/cv/core/data.ts` and not any document here, which `cv/README.md` had no reason to say, and it gives
    the order for changing a fact — evidence into `cv-raw.md` first, then the record, then regenerate — because
    absorbing these documents into the application is exactly what creates the chance of the two disagreeing. What is
    dropped is the `generate-cv-ats-pdf.py` instruction, which named a script and a PDF this phase deletes.
  - Note, 2026-09-01: the link check failed once on this item, at a line in this very document. The note on the
    preceding item quoted the repaired sentence verbatim, including its relative Markdown link, and the checker resolved
    that target against `delivery.md` rather than against the file being described. The note now names the link without
    reproducing its syntax. Worth recording because it is a hazard of writing evidence into a checked document at all:
    prose about a link is prose, and a link inside that prose is a link.
- [x] [AI] Confirm no absorbed document is imported by any route — acceptance:
      `rg -n 'docs/cv-raw|docs/cv-linkedin|docs/linkedin-projects' apps/wahidyankf-www/src` finds nothing, so the
      evidence base is repository material rather than published content.
  - Note, 2026-09-01: the search finds nothing, so none of the three absorbed documents reaches a rendered page. That
    keeps them what they were under `cv/` — evidence a person reads — rather than turning them into an input the site
    depends on, which would give the repository two CV sources again in a new shape. The `docs/README.md` written above
    states the same thing in prose for a reader who arrives at the directory rather than at this checklist.
- [x] [AI] Compare the retired ATS source against the application CV record — acceptance: every employer and role listed
      in `cv/cv-ats.md` appears in `apps/wahidyankf-www/src/features/cv/core/data.ts`. Record any role present in the
      retired file and absent from `data.ts` in `learnings.md` and add it to `data.ts` before proceeding. [AC-1]
  - **Clause disposition, 2026-09-01: not triggered.** Every employer and role in the retired `cv/cv-ats.md` appears in
    `data.ts`, so nothing was absent and nothing was added. The clause instructing this item to record a missing role in
    `learnings.md` and add it to `data.ts` before proceeding therefore never fired, and `learnings.md` correctly names
    no such role. This mattered: it is the check that had to pass before `cv/` could be deleted at all, and a role found
    only in the retired file would have made the deletion lossy. [AC-1]
  - Note, 2026-09-01: every employer and role in `cv/cv-ats.md` appears in `data.ts`, so nothing is owed to
    `learnings.md` and nothing had to be added. The three employers are Hijra, GudangAda, and Ruangguru. Hijra's three
    roles map one-for-one — Head of Engineering Hijra Bank, Engineering Manager Hijra Bank, and Engineering Manager for
    Alami P2P Lending and Hijra Bank Financing — as does GudangAda's Engineering Manager. Ruangguru is the one that does
    not map one-for-one, and it maps in the safe direction: the ATS compresses it into a single line, "Frontend Engineer
    to Engineering Manager", naming the five roles inside the bullet, while `data.ts` carries all five as separate
    entries — Junior Frontend Engineer, Frontend Engineer, Senior Frontend Engineer, Technical Lead, and Engineering
    Manager. The record is finer than the file being retired, not coarser, so deleting `cv-ats.md` loses no role. The
    same check was run past the acceptance's wording on the sections it does not name: the Bachelor of Engineering, the
    honors entries, and both independent projects the ATS lists, OSE and BeaverNest, are all present too. [AC-1]
- [x] [AI] Run `npx nx run wahidyankf-www:generate:cv-pdf` — acceptance: exits 0 and
      `apps/wahidyankf-www/public/wahidyankf-kresna-fridayoka-cv.pdf` is regenerated and begins with the PDF header
      bytes.
  - Note, 2026-09-01: exits 0, reports the file it wrote on stdout, and the result begins with `%PDF-`. 20.6K against
    the retired `cv/cv-ats.pdf` at 8.5K, which is expected rather than alarming: the ATS export was a two-page
    condensation produced by a Python script from `cv-ats.md`, and this one is rendered from the full CV record. The
    stdout line is the `console.log` an unsafe autofix deleted earlier in Phase 3 and that was restored with a narrow
    suppression; without it this run would be indistinguishable from one that produced nothing, which is exactly what
    this item is checking.
- [x] [AI+HUMAN] Open the regenerated PDF and confirm it is a usable replacement for `cv/cv-ats.pdf` — acceptance: the
      owner confirms the roles, dates, and layout are acceptable. This is the last point at which the ATS export can be
      preserved instead of deleted.
  - Note, 2026-09-01: **the owner confirmed the regenerated PDF is a usable replacement** — roles, dates, and layout
    acceptable — and directed the deletion to proceed. Both files were sent for side-by-side inspection: the regenerated
    `public/wahidyankf-kresna-fridayoka-cv.pdf` and the retired `cv/cv-ats.pdf`. This is the one item in the phase no
    automated check can stand in for, which is why the plan marks it `[AI+HUMAN]`: every check up to here establishes
    that a PDF was produced and that no role was lost from the record, and none of them establishes that the document is
    one the owner would send to an employer.
- [x] [AI] Record the pre-deletion digest of every tracked file under `cv/` — run
      `mkdir -p local-tmp && git ls-files cv | xargs shasum -a 256 | sort > local-tmp/cv-digests-before.txt` —
      acceptance: the file holds seven lines, one per tracked file, and `local-tmp/` is gitignored so the artifact never
      reaches a commit.
  - Note, 2026-09-01: seven tracked files recorded to the gitignored `local-tmp/cv-digests-before.txt` — `README.md`,
    `cv-ats.md`, `cv-ats.pdf`, `cv-linkedin.md`, `cv-raw.md`, `generate-cv-ats-pdf.py`, and `linkedin-projects.md`.
    `git check-ignore local-tmp` confirms the file is not committed, which is what this plan's evidence rules ask for
    output nothing reads back after the phase. Two of the digests can already be checked against work done earlier in
    this phase: `cv-linkedin.md` at `6330c2aa…` and `linkedin-projects.md` at `4321e23c…` match their copies under
    `apps/wahidyankf-www/docs/` exactly, and `cv-raw.md` at `80784d6e…` matches the source while its copy deliberately
    differs by the one repaired sentence.
- [x] [AI] Find every reference to the repository-root `cv/` directory with the broad pattern, excluding the ported
      application's own tree — acceptance, **amended during execution**:
      `rg -n '\bcv/' --hidden --glob '!node_modules' --glob '!package-lock.json' --glob '!.git' --glob '!apps/wahidyankf-www/**' .`
      lists every file to repair, the list is recorded in `evidence/cv-references.txt` before any edit, and that file's
      entry in the `## Directory Map` of `evidence/README.md` is converted to a relative link in this same item, which
      leaves every entry in that map linked. The narrow filename pattern misses six of them, because most references
      name the bare directory rather than a file inside it. The `apps/wahidyankf-www/**` exclusion is what keeps the
      sweep on the directory this phase retires: inside the ported application the same pattern matches its own CV
      feature directory — `@/features/cv/core/data`, `./cv/page`, and 26 further lines at the recorded source commit —
      and none of those names this repository's root `cv/`, because `ose-public` has no root `cv/` for the ported tree
      to point at. The `cv/` directory itself is deliberately left in this discovery sweep, so the recorded evidence is
      complete: expect `cv/README.md` line 20 to appear, and treat it as deleted rather than repaired, because the
      next-but-one item removes the whole directory. The nine files outside `cv/` are the repair set. The `.git`
      exclusion was added at the archival gate, to this sweep and to the confirmation sweep later in this phase in one
      edit; the confirmation item records why for both.
  - Note, 2026-09-01: recorded in `evidence/cv-references.txt` before any edit, and that file's map entry in
    `evidence/README.md` is now a relative link. The nine files the item predicts are exactly the nine that carry a
    routing or scope reference. The item's warning about the narrow pattern held: of the fifteen references, only three
    name a file inside the directory and the other twelve name the bare directory in an enumeration. The
    `apps/wahidyankf-www/**` exclusion also earned itself — inside the application the same pattern matches its own CV
    feature, and repointing those would have broken the feature this phase depends on. One thing the item did not
    predict: `specs/apps/wahidyankf-www/architecture.md` sits outside that exclusion and matches three times, twice for
    the application's own `src/features/cv/` and once in a past-tense sentence about the retired directory. That is
    handled by an item added during execution below.
- [x] [AI] Label every line of the recorded reference list as either a **routing reference** or a **scope enumeration**
      before repairing any of it, writing the label beside each line in `evidence/cv-references.txt` — acceptance: every
      recorded line outside `cv/` itself carries exactly one label, and every line matches the label
      [Reference Repair](tech-docs/migration-design.md#reference-repair) gives it; a line the sweep finds that the table
      there does not name is labelled beside the others in `evidence/cv-references.txt` rather than repaired on
      improvisation. **Amended during execution**: the criterion first sent such a line to `learnings.md`. Two lines
      needed labels the table does not name — `SELF` for `cv/README.md` line 20 and `NARRATIVE` for the architecture
      document's past-tense sentence — and both were labelled in the evidence file and settled by the classification
      item added later in this phase, which is the better handling and is where the record belongs; no `learnings.md`
      entry was written and none should have been. The two kinds get opposite treatment, by owner direction. A routing
      reference — text telling an agent where to read CV material, such as `AGENTS.md`'s "For CV work, read
      `cv/README.md`" — is **repointed** at `apps/wahidyankf-www/docs/`. A scope enumeration — a list of top-level
      directories a rule applies to, such as the indexed-README list — has `cv/` **struck from the list**, with no
      `apps/` path inserted in its place. Repointing an enumeration would push an `apps/` subdirectory into a rule scope
      that deliberately excludes `apps/`, which the Phase 5 `apps/README.md` item states of the documentation index
      policy, and would widen the corpus the `rules-checker` prompts bound. At this plan's authoring the sweep finds
      fifteen occurrences across nine files outside `cv/` itself: two routing references and thirteen scope
      enumerations.
  - Note, 2026-09-01: every line is labelled in `evidence/cv-references.txt` before any of it was repaired, and the two
    labels the item names were not enough for the list as found. Two lines are ROUTING and thirteen are SCOPE, which is
    the split the item anticipates. Two further labels were needed: SELF for `cv/README.md` line 20, which is inside the
    directory and goes with it rather than being repaired, and NARRATIVE for the architecture document's past-tense
    sentence, which is neither a route into the directory nor a rule scoped over it. Labelling first is what made the
    difference in treatment obvious — a ROUTING line gets a new destination, a SCOPE line gets `cv/` struck with nothing
    put in its place, and the other two get neither.
- [x] [AI] Repoint the two routing references at `apps/wahidyankf-www/docs/`: `AGENTS.md`'s "For CV work, read
      `cv/README.md`" and `CLAUDE.md`'s "`cv/` holds career evidence; read `cv/README.md` before touching it", each of
      which is a Markdown link to that path in the file it lives in — acceptance: each file carries one sentence naming
      `apps/wahidyankf-www/docs/README.md` as where CV material is read, `npm run check:markdown-links` exits 0 on the
      two rewritten links, and `npm run check:governance` exits 0. Both sentences route a reader to the material rather
      than bounding a rule's scope, which is why these two are repointed while the enumeration in the same `AGENTS.md`
      is not.
  - Note, 2026-09-01: both repointed and `npm run check:markdown-links` exits 0. `AGENTS.md` now reads "For CV work,
    read apps/wahidyankf-www/docs/README.md" and `CLAUDE.md` reads "`apps/wahidyankf-www/docs/` holds career evidence;
    read its README before touching it". These are the two lines that would have sent an agent into a directory that no
    longer exists, which is the failure mode the rules-gate corpus document called out by name before this phase struck
    that sentence.
- [x] [AI] Strike `cv/` from the one scope enumeration in `AGENTS.md` — its indexed-README sentence, "Every `docs/`,
      `repo-governance/`, `cv/`, `scripts/`, `plans/`, `specs/`, and harness directory requires an indexed README" —
      leaving the other six entries in place and adding no `apps/` path — acceptance:
      `rg -n '\bcv/' AGENTS.md CLAUDE.md` finds nothing, the line `rg -n 'requires an indexed README' AGENTS.md` returns
      still names `docs/`, `repo-governance/`, `scripts/`, `plans/`, `specs/`, and the harness directories, and names no
      `apps/` path, and `npm run check:governance` exits 0. `CLAUDE.md` carries no enumeration of this kind, so it is
      finished by the repoint above. Adding `apps/wahidyankf-www/docs/` here would state a rule the documentation index
      policy does not hold, since that policy's own scope does not reach `apps/`.
  - Note, 2026-09-01: struck, with no `apps/` path added in its place. The sentence now reads "Every `docs/`,
    `repo-governance/`, `scripts/`, `plans/`, `specs/`, and harness directory requires an indexed README". Adding the
    new location would have been the wrong repair: the absorbed documents sit inside a project, and no project directory
    appears in this enumeration — `apps/badakmini-cli` does not — so naming one would extend the rule rather than
    preserve it.
- [x] [AI] Strike `cv/` from the five scope enumerations in three governance documents, adding no `apps/` path to any of
      them — `repo-governance/documentation-index-policy.md` twice, in its `when_to_use` front-matter line and in the
      "Every directory in ... must contain a `README.md`" list; `repo-governance/README.md` once, in the "Use it when
      adding, moving, or maintaining Markdown under ..." sentence of its documentation-index-policy entry, which
      restates that policy's scope and has to keep agreeing with it; and `repo-governance/workflows/readme-refresh.md`
      twice, in the "Review the root `README.md` and every existing README below ..." list and in the "Follow the
      documentation index policy everywhere it applies, which is ..." list — acceptance:
      `rg -n '\bcv/' repo-governance/documentation-index-policy.md repo-governance/README.md repo-governance/workflows/readme-refresh.md`
      finds nothing, each of the five lists keeps every entry other than `cv/`, none gains an `apps/` entry, and
      `npm run check:governance` exits 0. None of the five is a routing reference. The `readme-refresh.md` review list
      already names `apps/`, so `apps/wahidyankf-www/docs/README.md` stays in that workflow's scope through the entry
      that is already there; inserting it again would say the same thing twice.
  - Note, 2026-09-01: all five struck and `npm run check:governance` exits 0. Two are in
    `documentation-index-policy.md`, one in its `when_to_use` frontmatter and one in the body enumeration; one is in
    `repo-governance/README.md`'s entry for that policy; and two are in `readme-refresh.md`, its review scope and its
    restatement of the index policy's scope. The frontmatter one is the easiest to miss, because it is prose inside a
    metadata field rather than a sentence in the body, and a sweep that read only rendered text would not have found it.
    No `apps/` path was added to any of them.
- [x] [AI] Strike the `cv/` entry from the rules-gate corpus list in
      `repo-governance/workflows/rules-quality-gate/01-scope-and-corpus.md`, whose bullet reads "`cv/`, `scripts/`, and
      the root `README.md` — each carries rule sentences, and `AGENTS.md` routes agents into `cv/README.md`, so a stale
      reference there misdirects an agent": leave the bullet naming `scripts/` and the root `README.md` as surfaces that
      each carry rule sentences, and drop the trailing `AGENTS.md`-routing clause, which the repoint above makes false —
      acceptance: `rg -n '\bcv/' repo-governance` finds nothing across the whole directory, the bullet still names
      `scripts/` and the root `README.md`, and the `In Scope, Judged Narrowly` list gains no `apps/` entry. This is a
      corpus enumeration, not a routing reference: naming `apps/wahidyankf-www/docs/` here would put an application
      subdirectory into what the rules gate reads, which is the widening the classification item forbids.
  - Note, 2026-09-01: struck. The bullet read "`cv/`, `scripts/`, and the root `README.md` — each carries rule
    sentences, and `AGENTS.md` routes agents into `cv/README.md`, so a stale reference there misdirects an agent", and
    now reads "`scripts/` and the root `README.md` — each carries rule sentences, and a stale reference in one
    misdirects an agent". Both halves had to change rather than just the list: the justification named the `AGENTS.md`
    routing line specifically, and that line was repointed two items above, so leaving it would have described a route
    that no longer exists as the reason for a scope that no longer includes it.
- [x] [AI] Strike `cv/` from both scope enumerations in each of the three mirrored `rules-checker` prompts —
      `.claude/agents/rules-checker.md`, `.codex/agents/rules-checker.toml`, and `.opencode/agents/rules-checker.md` —
      the corpus paragraph's "Read `docs/`, `cv/`, `scripts/`, and the root `README.md` narrowly" and the
      README-registration sentence's "every directory README in `docs/`, `repo-governance/`, `cv/`, `scripts/`,
      `plans/`, `specs/`, and every harness directory registers its immediate documents and child directories" —
      acceptance: `rg -n '\bcv/' .claude .codex .opencode --hidden` finds nothing, neither sentence in any of the three
      gains an `apps/` entry, and the three files still say the same thing as each other, which
      `npm run check:harness-parity` and a three-way diff of the changed paragraphs confirm. The
      [harness capability parity policy](../../../repo-governance/conventions/harness-capability-parity-policy.md)
      requires the same edit in all three copies. Both occurrences are enumerations of what this subagent reads, so both
      are struck rather than repointed; repointing either would hand `rules-checker` an application subdirectory it is
      not meant to read, and the README-registration sentence would additionally assert an indexed-README rule the
      documentation index policy does not hold for `apps/`.
  - Note, 2026-09-01: six enumerations struck across the three prompts — the corpus sentence and the directory-README
    sentence in each — and `npm run check:harness-parity` exits 0, so no harness gained or lost a capability the others
    lack. The three were edited with one substitution list applied to all three files, each asserted to match exactly
    once, because these prompts are mirrors and an edit that reached two of them would be a parity defect that only the
    check would catch.
- [x] [AI] Classify the three `specs/apps/wahidyankf-www/architecture.md` matches the repair set does not cover, and
      repair none of them — acceptance: lines 78 and 113 name `src/features/cv/`, the application's own component, and
      line 115 names the retired directory only in past tense; `npm run check:markdown-links` and
      `npm run check:governance` both still exit 0. **This item was added during execution.** The discovery sweep
      excludes `apps/wahidyankf-www/**`, which keeps the application's own CV feature out of the results, but the C4
      model describing that application lives under `specs/` and is not excluded, so it reports three matches that the
      nine-file repair set does not and should not cover.
  - Note, 2026-09-01: none of the three is repaired, and each for a different reason. Lines 78 and 113 are the component
    table and the data-store table naming `src/features/cv/` and `src/features/cv/core/data.ts`; repointing them at
    `apps/wahidyankf-www/docs/` would break the model's description of the feature this phase depends on, which is the
    same hazard the `apps/wahidyankf-www/**` exclusion exists to prevent, reaching a file that exclusion does not cover.
    Line 115 is the one sentence that genuinely names the retired directory: "Before the migration this repository also
    kept a `cv/` directory at its root, and two records of the same career are a guarantee of eventually publishing the
    stale one. The migration deleted `cv/` and left this store as the single source." It is already written for the
    post-migration state, in the past tense, and it is the reason the model gives for why a single CV record is
    load-bearing rather than incidental. Deleting it to satisfy a pattern would remove the explanation and leave the
    claim. Rewording it to drop the trailing slash would satisfy the same pattern while changing nothing a reader
    understands, which is worse — it would make the sweep pass without making the statement more true. It stays, and the
    confirmation item below is read against it.
- [x] [AI] Confirm no reference survives — acceptance:
      `rg -n '\bcv/' --hidden --glob '!node_modules' --glob '!package-lock.json' --glob '!.git' --glob '!apps/wahidyankf-www/**' --glob '!cv/**' .`
      finds no match outside this plan's own documents **except in `specs/apps/wahidyankf-www/architecture.md`**, whose
      matches the item above classifies as correct — two naming the application's own `cv` component and one naming the
      retired directory in past tense; the nine files named in the five preceding repair items are the whole repair set.
      **Amended during execution**, twice. The criterion first read "no match outside this plan's own documents" and the
      delivered tree has three, each of them a match the plan intends to survive. The archival gate then added
      `--glob '!.git'` to both this sweep and the discovery sweep above: `--hidden` is in the command so it reaches
      `.claude/`, `.codex/`, and `.opencode/`, and it reaches `.git/` too, where this phase's own commit message —
      `feat: delete cv/ now that the application holds the CV record` — sits in the reflog. The criterion became
      unsatisfiable the moment the phase it belongs to landed, on any clone carrying that history. Three exclusions
      carry weight here. The `apps/wahidyankf-www/**` one carries over from the discovery item above: without it this
      confirmation would report the application's own `features/cv/` and `app/cv/` paths as unrepaired references, and
      repointing those at `apps/wahidyankf-www/docs/` would break the CV feature this phase depends on. The `cv/**` one
      is what makes the acceptance attainable at this point in the order: `cv/README.md` line 20 shows the regeneration
      command `uv run --with reportlab python cv/generate-cv-ats-pdf.py`, which is a reference to the directory by any
      pattern, and it is the only such line inside `cv/` today. That file is deleted by the very next item rather than
      repaired — repairing a file about to be removed would be work with no reader — so this sweep looks everywhere
      except the directory being retired. [AC-1]
  - Note, 2026-09-01: the sweep finds three matches outside this plan's documents, all three in
    `specs/apps/wahidyankf-www/architecture.md`, and all three classified by the item directly above as correct rather
    than unrepaired — two naming the application's own component and one naming the retired directory in past tense.
    Every routing reference and every scope enumeration is gone: nine files repaired, fifteen references, and no
    reference anywhere still sends a reader into `cv/` or scopes a rule over it. The acceptance as written says "finds
    no match", and it finds three; recorded as it happened rather than by narrowing the pattern until the sentence
    became true. The two exclusions the item does defend both earned themselves — without `apps/wahidyankf-www/**` the
    application's own CV feature would read as unrepaired, and without `cv/**` the regeneration command on line 20 of a
    file being deleted by the next item would have to be repaired for no reader. [AC-1]
- [x] [AI] Run `git rm -r cv` — acceptance: `git ls-files cv` returns nothing and `test ! -e cv` succeeds. [AC-1]
  - Note, 2026-09-01: `git ls-files cv` returns nothing and `test ! -e cv` succeeds. Seven tracked files removed — the
    three absorbed into `apps/wahidyankf-www/docs/`, the two ATS artefacts the application's own export replaces, the
    Python generator that produced them, and the directory README whose working rules were rewritten into the new one.
    Their digests were recorded before this ran and the recovery rehearsal in the gate below checks them back. [AC-1]
- [x] [AI] Run `npm run check:markdown-links` — acceptance: exits 0, proving no document still links into the deleted
      directory.
  - Note, 2026-09-01: exits 0, so no document anywhere links into the deleted directory. This is the check the reference
    repairs were made for, and it is a stronger statement than the `rg` sweep above: the sweep looks for a text pattern,
    and this resolves every link target in every tracked Markdown file, including the ones that named `cv/README.md`
    through a relative path rather than the literal string.

### Phase 4 Gate

> Every check below passes before Phase 5 begins. A failure is fixed inside Phase 4.

- [x] [AI] Run `git ls-files cv` — acceptance: no output. [AC-1]
  - Note, 2026-09-01: `git ls-files cv` prints nothing and exits 0, and `test ! -e cv` succeeds, so the directory is
    gone from the index and from the working tree rather than merely untracked. The pathspec form is the vacuous-pass
    hazard this plan warns about elsewhere — `git ls-files` on a path matching nothing is silent whether the path was
    deleted or never existed — so the working-tree test is run beside it, and the recovery rehearsal below is what
    proves the seven files were real.
- [x] [AI] Rehearse recovery of the whole deleted directory from the preceding commit — run
      `rm -rf local-tmp/cv-recovery && mkdir -p local-tmp/cv-recovery && git archive HEAD cv | tar -x -C local-tmp/cv-recovery`
      — acceptance: `find local-tmp/cv-recovery/cv -type f | wc -l` prints `7`, so `README.md`, `cv-raw.md`,
      `cv-linkedin.md`, `linkedin-projects.md`, `cv-ats.md`, `cv-ats.pdf`, and `generate-cv-ats-pdf.py` all come back.
      This runs before the deletion is committed, so `HEAD` is still the commit that holds them.
  - Note, 2026-09-01: seven files come back — `README.md`, `cv-raw.md`, `cv-linkedin.md`, `linkedin-projects.md`,
    `cv-ats.md`, `cv-ats.pdf`, and `generate-cv-ats-pdf.py` — and `find local-tmp/cv-recovery/cv -type f | wc -l` prints
    `7`. One deviation from the item as written: it says to archive from `HEAD` because it expects to run before the
    deletion is committed, but the deletion was committed first as `31aabe5`, so `HEAD` no longer holds `cv/` and the
    archive was taken from `HEAD~1`. That is the same commit the item means — the one immediately preceding the deletion
    — and the digest comparison below is what proves the source is the right one rather than an arbitrary earlier tree.
- [x] [AI] Confirm the rehearsed files are byte-identical to what was deleted — run
      `(cd local-tmp/cv-recovery && find cv -type f | sort | xargs shasum -a 256 | sort) | diff - local-tmp/cv-digests-before.txt`
      — acceptance: `diff` reports no difference across all seven digests. The three files that are deleted outright
      rather than moved — `cv-ats.md`, `cv-ats.pdf`, and `generate-cv-ats-pdf.py` — have no copy anywhere else in the
      tree, so this is the only proof their recovery source is real.
  - Note, 2026-09-01: `diff` reports no difference across all seven digests, so every recovered file is byte-identical
    to what was deleted. This matters most for `cv-ats.md`, `cv-ats.pdf`, and `generate-cv-ats-pdf.py`, which are
    deleted outright rather than moved and have no copy anywhere else in the tree: the digest match is the only evidence
    their recovery source is real.
- [x] [AI] Run `npm run check:markdown-links` — acceptance: exits 0.
  - Note, 2026-09-01: exits 0. Run after the `CLAUDE.md` edit recorded below, not before, so the check covers the
    removed link rather than a stale tree.
- [x] [AI] Run `npm run check:governance` — acceptance: exits 0, because `AGENTS.md`, `CLAUDE.md`, and four
      `repo-governance/` documents changed. [AC-8]
  - Note, 2026-09-01: exits 0. [AC-8]
- [x] [AI] Run `npm run check:harness-parity` — acceptance: exits 0, because the three `rules-checker` prompts changed
      and no harness may gain or lose a capability the others lack.
  - Note, 2026-09-01: exits 0. The six subagents are present in all three harness directories with the same names, so no
    harness gained or lost a capability when the three `rules-checker` prompts were edited.
- [x] [AI] Run `git add -A` and then `npm run check:rule-change` — acceptance: it names the `repo-governance/` and
      harness-directory changes and the workflows they trigger, and those workflows are followed rather than dismissed.
      Staging comes first because the check reads `git diff --cached --name-only`, as
      [workspace commands](../../../repo-governance/development/workspace-commands.md#repository-checks) states: run
      against an unstaged tree it sees no paths and reports nothing, which reads identically to a clean result.
  - Note, 2026-09-01: the check names `Rules Propagation automatically triggered by CLAUDE.md` and
    `Harness setup changed in CLAUDE.md. Run repo-governance/workflows/harness-alignment.md`, and both were worked
    rather than dismissed. It names only `CLAUDE.md` at gate time because it reads `git diff --cached --name-only` and
    the phase's other rule-path edits — `AGENTS.md`, four `repo-governance/` documents, and the three `rules-checker`
    prompts — were staged and committed earlier in the phase as `f338a16`, where the same check named all nine. **Rules
    Propagation**: the rule restated in one sentence is that career-evidence source documents live in
    `apps/wahidyankf-www/docs/` and an agent doing CV work reads that directory's README first. Its canonical home is
    `AGENTS.md`; the index-scope half stays in the documentation index policy. `rg -n '\bcv/'` across `AGENTS.md`,
    `CLAUDE.md`, `repo-governance/`, the three harness directories, and `opencode.json` finds nothing, and both routing
    pointers resolve, so the removal is idempotent. Nothing was added: `apps/wahidyankf-www/docs/README.md` needs no
    index entry, because the documentation index policy's scope names the top-level `docs/`, `repo-governance/`,
    `scripts/`, `plans/`, `specs/`, and harness trees and reaches no path under `apps/`; `check:markdown-links` already
    protects the routing link, which is enough under minimum sufficiency. **Harness Alignment**: the inventory found six
    subagents mirrored across `.claude/agents/`, `.codex/agents/`, and `.opencode/agents/`, and every command and path
    quoted in the edited derivatives resolves. Derivative comparison found one defect, recorded as its own item below.
- [x] [AI] Remove the sentence in `CLAUDE.md` that restates the CV routing rule `AGENTS.md` owns, keeping the
      `apps/badakmini-cli` clause it shared a line with — acceptance: `rg -n 'wahidyankf-www/docs' CLAUDE.md` finds
      nothing, `rg -n 'wahidyankf-www/docs' AGENTS.md` still finds the one routing pointer, and
      `npm run check:markdown-links`, `npm run check:governance`, and `npm run check:harness-parity` all exit 0. **This
      item was added during execution.** It was found by the Harness Alignment run the gate item above requires, at its
      derivative-comparison step. When the earlier item repointed both files from `cv/README.md` to
      `apps/wahidyankf-www/docs/README.md`, it preserved a duplication that predates this plan: `AGENTS.md` and
      `CLAUDE.md` each carried the whole rule, so either could drift. The
      [agent instruction alignment policy](../../../repo-governance/conventions/agent-instruction-alignment-policy.md)
      allows a derivative a short pointer that links to a rule's home but forbids one that silently duplicates canonical
      guidance, and its Verification section names exactly that; `CLAUDE.md`'s own preamble says the file must never
      restate canonical guidance. The fix is at the derivative rather than the canonical source, because `AGENTS.md` is
      where the rule belongs.
  - Note, 2026-09-01: the clause is gone and `AGENTS.md` line 9 is now the single home of the CV routing rule. **The
    item sits below the gate's `check:governance` and `check:harness-parity` lines but ran before both of them**,
    because it was found by the Harness Alignment run the `check:rule-change` item requires and fixed at once; those two
    checks, and `check:markdown-links`, `test:quick`, and `format:check`, were all run afterwards and all exit 0 against
    the edited file. The checklist order is what is misleading here, not the result — an item added during execution was
    appended where its trigger was recorded rather than inserted above the checks it would otherwise invalidate.
    `apps/badakmini-cli` owns repository-local checks stays, because `AGENTS.md` never states it and it is therefore
    Claude Code-specific detail rather than a restatement. Removing the sentence also bought `CLAUDE.md` headroom under
    the [document word limit policy](../../../repo-governance/conventions/document-word-limit-policy.md), which is a
    side effect and not the reason.
- [x] [AI] Run `npm run test:quick` — acceptance: exits 0 across both projects.
  - Note, 2026-09-01: exits 0 across both projects — 12 test files and 258 tests pass for `wahidyankf-www`, and
    `badakmini-cli` passes.
- [x] [AI] Run `npm run format:check` — acceptance: exits 0.
  - Note, 2026-09-01: exits 0. `prettier --check .` reports all matched files use Prettier code style.
- [x] [AI] Commit and push the phase to `main` — acceptance: `git status --short` is empty, and this commit is the
      single revert point that restores `cv/` intact.
  - Note, 2026-09-01: `git status --short` is empty after the push. The phase landed in three commits rather than one,
    under the [thematic commits policy](../../../repo-governance/conventions/thematic-commits-policy.md): `f338a16`
    absorbed the CV materials and rewired their references, `31aabe5` deleted `cv/`, and the record and the `CLAUDE.md`
    alignment fix follow. `31aabe5` is the single revert point the item names — reverting it restores all seven files
    intact, which the rehearsal and digest diff above prove.

> **Pause Safety**: The repository holds exactly one CV record. The evidence base lives in the application's `docs/` and
> is not imported by any route. The retired ATS source, export, and Python generator are gone, and their replacement was
> proven before they were deleted. Recovery from the preceding commit was rehearsed rather than assumed. Safe to stop.
> Resume with `npm run test:quick`.

## Phase 5: Process E2E Harness and Scheduled Verification

- [x] [AI] Copy `$SRC/apps/wahidyankf-www-fe-e2e/{package.json,playwright.config.ts,tsconfig.json,README.md,.gitignore}`
      and the eight files under `steps/` to `apps/wahidyankf-www-e2e/` — acceptance: all thirteen files exist.
  - Note, 2026-09-01: all thirteen exist — five root files and the eight under `steps/`. The source project at the
    recorded commit also holds `.features-gen/`, which is generated and ignored, so it is not among the thirteen.
- [x] [AI] Do not copy `$SRC/apps/wahidyankf-www-fe-e2e/scripts/run-docker-e2e.mjs` or `e2e-coverage-baseline.json` —
      acceptance:
      `test ! -e apps/wahidyankf-www-e2e/scripts && test ! -e apps/wahidyankf-www-e2e/e2e-coverage-baseline.json`
      succeeds, because the Docker runner is dropped and the baseline is consumed only by `rhino-cli`. [AC-5]
  - Note, 2026-09-01:
    `test ! -e apps/wahidyankf-www-e2e/scripts && test ! -e apps/wahidyankf-www-e2e/e2e-coverage-baseline.json`
    succeeds. [AC-5]
- [x] [AI] Point `apps/wahidyankf-www-e2e/tsconfig.json` at the workspace base by adding
      `"extends": "../../tsconfig.base.json"`, keeping its existing `module`, `moduleResolution`, and `target` values as
      overrides on top — acceptance:
      `rg -n '"extends": "../../tsconfig.base.json"' apps/wahidyankf-www-e2e/tsconfig.json` finds the inheritance. The
      acceptance stops at the file's content on purpose: `npx nx run wahidyankf-www-e2e:typecheck` cannot pass here,
      because `apps/wahidyankf-www-e2e/project.json` is authored further down this phase and `npm install` runs after
      that, so the target does not exist yet. This project's `typecheck` proof is the Phase 5 gate's `npm run typecheck`
      run, which the gate states is where the E2E project's own type-check is covered. The ported file carries no
      `extends` at the source commit. Leaving it standalone would give this repository two TypeScript projects that
      reach the shared base by different routes for no reason anyone could state later, and the deviation recorded in
      `tooling.md` is then about three options rather than about one project inheriting nothing.
  - Note, 2026-09-01: the `extends` is present and the three overrides stay. One thing the item did not anticipate:
    extending the base made `typecheck` fail with
    `TS5101: Option 'baseUrl' is deprecated and will stop functioning in TypeScript 7.0`, because the root pin is
    TypeScript 6 and the ported file carried `baseUrl: "."`. Removed it and kept `paths`, which TypeScript 5 and later
    resolve relative to the `tsconfig.json` itself. That is the same shape `apps/wahidyankf-www/tsconfig.json` already
    carries, so the two projects now agree; recorded as its own item below.
- [x] [AI] Remove `"baseUrl": "."` from `apps/wahidyankf-www-e2e/tsconfig.json`, keeping `paths` — acceptance:
      `npx nx run wahidyankf-www-e2e:typecheck` exits 0 and `rg -n 'baseUrl' apps/wahidyankf-www-e2e/tsconfig.json`
      finds nothing. **This item was added during execution.** The ported file carries `baseUrl` at the source commit,
      where TypeScript is pinned at `5.8.3`; the item above deletes that local pin so the root TypeScript 6 pin governs,
      and TypeScript 6 reports `TS5101: Option 'baseUrl' is deprecated and will stop functioning in TypeScript 7.0` as
      an error. `paths` needs no `baseUrl` in TypeScript 5 and later — it resolves relative to the `tsconfig.json`
      holding it — so removing the deprecated option alone is the whole fix, and it leaves this project agreeing with
      `apps/wahidyankf-www/tsconfig.json`, which Phase 3 landed in exactly that shape.
  - Note, 2026-09-01: `typecheck` exits 0. The `@/*` alias has no importer in this project today; `paths` is kept
    anyway, so a step file that reaches for it later resolves the way the application's does.
- [x] [AI] Change the `"name"` field in `apps/wahidyankf-www-e2e/package.json` from `wahidyankf-www-fe-e2e` to
      `wahidyankf-www-e2e` — acceptance: `rg -n '"name": "wahidyankf-www-e2e"' apps/wahidyankf-www-e2e/package.json`
      finds it and `rg -n 'fe-e2e' apps/wahidyankf-www-e2e/package.json` finds nothing. The project is renamed on
      arrival by owner direction, and an npm workspace whose manifest name disagrees with its directory and its Nx
      project name resolves inconsistently between the two tools.
  - Note, 2026-09-01: `rg -n '"name": "wahidyankf-www-e2e"'` finds it and `rg -n 'fe-e2e'` finds nothing in the
    manifest.
- [x] [AI] Rename the first heading of `apps/wahidyankf-www-e2e/README.md` from `# wahidyankf-www-fe-e2e` to
      `# wahidyankf-www-e2e` — acceptance: `rg -n '^# wahidyankf-www-e2e$' apps/wahidyankf-www-e2e/README.md` finds the
      heading and `rg -n '^# wahidyankf-www-fe-e2e$' apps/wahidyankf-www-e2e/README.md` finds nothing. The acceptance is
      the heading line alone, because six further `fe-e2e` occurrences are `npm exec nx -- run` command lines inside
      `Run the suite` and `Checks and specs`; the section rewrite later in this phase replaces those, and the
      repository-wide sweep at the end of the phase is what proves none survived. The copied file carries the old name
      in its heading at the source commit, and this plan's [README](README.md) names that heading alongside the manifest
      `"name"` field as repaired in this phase. The section rewrites later in this phase do not reach it: they replace
      `Run the suite` and `Checks and specs`, and their acceptance pattern does not match a title line. [AC-5]
  - Note, 2026-09-01: the heading reads `# wahidyankf-www-e2e`. The six `npm exec nx -- run` command lines the item
    predicted were replaced by the section rewrites later in this phase, and the repository-wide sweep confirms none
    survived.
- [x] [AI] Delete the `"typescript": "5.8.3"` entry from `apps/wahidyankf-www-e2e/package.json` `devDependencies`, so
      the root pin from Phase 1 governs this project too — acceptance, **amended during execution**:
      `rg -n '"typescript"' apps/wahidyankf-www-e2e/package.json` finds nothing,
      `test ! -e apps/wahidyankf-www-e2e/node_modules/typescript` succeeds, and
      `node -p "require('typescript/package.json').version"` run from inside `apps/wahidyankf-www-e2e` prints the root
      pin. The criterion was amended for the same reason as its Phase 3 twin above, and in the same edit. Under npm
      workspaces a nested pin resolves ahead of the root one, so leaving it would let `[AC-9]` pass while this project
      still compiled on 5.8.3. [AC-9]
  - Note, 2026-09-01: the entry is gone, no `node_modules/typescript` exists inside the project, and the amended read
    prints `6.0.3`, resolving to the root `node_modules/typescript/package.json`, so the root TypeScript 6 pin from
    Phase 1 governs this project.
- [x] [AI] Pin the caret-ranged `@vitejs/plugin-react` in `apps/wahidyankf-www-e2e/package.json` to the exact version
      `npm install` resolves, or remove it if no file in this project imports it — acceptance:
      `rg -n '"\^|"~' apps/wahidyankf-www-e2e/package.json` finds nothing. The Phase 3 pinning item covers only the
      application manifest, so this specifier would otherwise cross the move as a range. [AC-9]
  - Note, 2026-09-01: **the removal branch fired.** `rg -rn 'vitejs/plugin-react' apps/wahidyankf-www-e2e` matched the
    manifest and nothing else — no file in this project imports it — so it was deleted rather than pinned. It sat under
    `dependencies` rather than `devDependencies`, which is what a Playwright project has no use for at all.
    `rg -n '"\^|"~' apps/wahidyankf-www-e2e/package.json` now finds nothing. [AC-9]
- [x] [AI] Author `apps/wahidyankf-www-e2e/project.json` with `implicitDependencies` naming `wahidyankf-www`, and five
      targets carrying the verbatim commands the E2E Target Contract in
      [technical design](tech-docs/README.md#the-e2e-projects-target-contract) states, each with `cwd` set to
      `{projectRoot}` except the `lint` aggregate, which takes none: `install` running
      `npx playwright install --with-deps chromium`; `typecheck` running `tsc --noEmit`; `lint` as the ordered aggregate
      the item below defines; `test:e2e` running `npx bddgen && npx playwright test` behind the guard the item below
      prepends; and `specs:e2e:baseline` running the command the dedicated item later in this phase writes out. Give
      each an `inputs` array whose first entry is `"default"` and whose second is
      `{workspaceRoot}/specs/apps/wahidyankf-www/behaviours/**/*.feature` — acceptance:
      `npx nx show project wahidyankf-www-e2e --json` lists those five targets with the command and `cwd` that table
      states, each with `"default"` first and the corpus input second, and
      `rg -n 'rhino-cli|dotnet|bddgen.*rhino|run-docker-e2e' apps/wahidyankf-www-e2e/project.json` finds nothing.
      `"default"` is named because a target-level `inputs` array replaces Nx's default input set rather than extending
      it, and this file is authored fresh, so an array holding the corpus glob alone would leave the project's own
      `steps/`, `playwright.config.ts`, and manifest out of every hash — the same reason the application's `inputs`
      items in Phase 3 name it, and the same shape `apps/badakmini-cli/project.json` and the source
      `apps/wahidyankf-www/project.json` both use. `specs:e2e:baseline` is declared here as a named target; its verbatim
      command, `cwd`, and `inputs` are written by the dedicated item later in this phase, which is where the baseline
      file it reads is also created. Until then the target exists and is not yet runnable, which is why nothing before
      that item invokes it. `npx bddgen && npx playwright test` is a reduction of the source target rather than a new
      command: at the recorded source commit `test:e2e` ran the guard and then `node scripts/run-docker-e2e.mjs`, and
      the only two commands that runner issues once its Docker work is done are `npx bddgen` and then
      `npx playwright test`, both from the project directory. The Docker build, health wait, and published-port lookup
      are what this plan drops, replaced by the `webServer` block added further down this phase. Declaring `cache` on
      `install`, `test:e2e`, and `specs:e2e:baseline` is part of this item too, because root `nx.json` `targetDefaults`
      reaches only `typecheck` and `lint` of this project's seven targets, and that table states the setting for each.
      [AC-7]
  - Note, 2026-09-01: `npx nx show project wahidyankf-www-e2e --json` shows all five with the command and `cwd` the
    contract states, `"default"` first in each `inputs` array and the corpus glob second, and
    `rg -n 'rhino-cli|dotnet|run-docker-e2e'` finds nothing. `cache` is declared explicitly on `install`, `test:e2e`,
    and `specs:e2e:baseline`; `typecheck` and `lint` take theirs from root `targetDefaults`. [AC-7]
- [x] [AI] Carry the source project's unconditional-`test.skip` guard onto the new `test:e2e` command in
      `apps/wahidyankf-www-e2e/project.json`, prepending it ahead of that target's `npx bddgen && npx playwright test`
      and keeping `cwd` at `{projectRoot}`:
      `if grep -rn -E --include='*.ts' --exclude-dir=node_modules --exclude-dir=.features-gen --exclude-dir=test-results --exclude-dir=playwright-report '\$?test\.skip\([^,)]*\)' .; then echo 'ERROR: unconditional test.skip() found in test files above - use test.skip(condition, reason) for legitimate environment guards, or remove'; exit 1; fi && `
      — acceptance: `npx nx show project wahidyankf-www-e2e --json` shows `test:e2e` as the guard followed by
      `&& npx bddgen && npx playwright test`, which is what the E2E Target Contract in
      [technical design](tech-docs/README.md#the-e2e-projects-target-contract) states, and once the suite is runnable
      later in this phase, adding an unconditional `test.skip()` to one step file makes
      `npx nx run wahidyankf-www-e2e:test:e2e` exit non-zero with that message, and removing the line restores its
      previous result. The source carries this guard on the same target; this plan drops only the
      `node scripts/run-docker-e2e.mjs` half of that command, replacing it with the two commands that runner itself
      issued once its Docker work was done, so authoring the target fresh would otherwise lose the guard without anyone
      deciding to. Nothing else here catches an unconditional skip: `specs:e2e:baseline` counts `test.fixme` entries in
      `.features-gen`, which a hand-written `test.skip()` in a step file never becomes, and this project carries no
      coverage floor for a disabled scenario to drop below. As with `specs:e2e:baseline`, each `\.` is escaped once more
      when the command is written into JSON. [File impact](tech-docs/file-impact.md) records this guard as carried and
      the source's `test:unit` skip guard as deliberately dropped.
  - Note, 2026-09-01: the guard is verbatim and proved both ways rather than read. Appending `test.skip();` to
    `steps/theme.steps.ts` made `npx nx run wahidyankf-www-e2e:test:e2e` exit 1 printing
    `ERROR: unconditional test.skip() found in test files above`; restoring the file returned it to exit 0 with 36
    passed.
- [x] [AI] Give this project the same three-component lint shape the application got: create
      `apps/wahidyankf-www-e2e/eslint.config.mjs` with the same `eslint-plugin-jsdoc` commentary rules over `steps/`,
      add the two single-command children the E2E Target Contract in
      [technical design](tech-docs/README.md#the-e2e-projects-target-contract) states — `lint:biome` running
      `biome check` and `lint:commentary` running `eslint --config eslint.config.mjs steps`, both with `cwd` at
      `{projectRoot}` and both declaring `cache` explicitly, because root `nx.json` `targetDefaults` keys on `lint` and
      does not reach either child — and define `lint` as an ordered aggregate of `lint:biome` then `lint:commentary` —
      acceptance: `npx nx show project wahidyankf-www-e2e --json` shows `lint` with `parallel` set to `false` and an
      `options.commands` list of `npm exec nx -- run wahidyankf-www-e2e:lint:biome` then
      `npm exec nx -- run wahidyankf-www-e2e:lint:commentary`, shows both children as `command`-shorthand targets
      carrying those two commands, and `npx nx run wahidyankf-www-e2e:lint` exits 0. This is the second TypeScript
      project the plan creates, and `tooling.md`'s three components are stated of TypeScript projects without exception,
      so conforming one and not the other would let `[AC-9]` report a conformance half the new code does not have. The
      same per-component fallback applies: a component that cannot run is dropped from the aggregate and named in
      `learnings.md`. [AC-9]
  - Note, 2026-09-01: `lint` shows `parallel: false` with the two `npm exec nx -- run` entries in order, both children
    are `command`-shorthand targets with `cache: true`, and `npx nx run wahidyankf-www-e2e:lint` exits 0. No component
    was dropped, so no fallback fired. `eslint.config.mjs` reads `steps/**/*.ts` rather than `src/`, and enables no
    `jsx` parser feature — a Playwright step file has no JSX. Biome reports one warning, `useOptionalChain` at
    `steps/responsive.steps.ts:37`; it is a warning rather than an error, `lint` exits 0, and it was deliberately not
    fixed. Phase 3 recorded what an unsolicited Biome fix to ported test code costs: the `noArrayIndexKey` fix there
    changed an `id` alongside a `key` and broke two scenarios. [AC-9]
- [x] [AI] Give the `test:e2e` target `"dependsOn": ["wahidyankf-www:build"]` in `apps/wahidyankf-www-e2e/project.json`
      — acceptance: `npx nx show project wahidyankf-www-e2e --json` shows `"dependsOn": ["wahidyankf-www:build"]` on
      `test:e2e`. The acceptance stops at the declaration: the suite cannot run here, because `npm install` is still
      ahead, the `webServer` block is added after that, and `featuresRoot` still names the `ose-public` path until later
      in this phase. The cold-`.next` behaviour this declaration exists for is proved by the item below that runs
      `rm -rf apps/wahidyankf-www/.next` immediately before the suite. The `webServer` block runs `next start`, which
      needs a `.next` directory it does not build; `implicitDependencies` shapes the project graph and the affected
      calculation, not task ordering, so without this the suite fails on a cold checkout. `badakmini-cli` carries the
      same shape on its own `test:e2e`, whose `dependsOn` names `build` and `test:coverage:behaviour`. [AC-5]
  - Note, 2026-09-01: `"dependsOn": ["wahidyankf-www:build"]` is on `test:e2e`. Proved rather than declared: the
    cold-`.next` run below deletes `apps/wahidyankf-www/.next` first and the suite still passes, because the dependency
    builds it. [AC-5]
- [x] [AI] Confirm the E2E project defines no numeric coverage gate and no separate corpus — acceptance:
      `rg -n 'coverage.thresholds|\.feature' apps/wahidyankf-www-e2e/project.json` finds only the shared
      `specs/apps/wahidyankf-www/behaviour` input, as the BDD role matrix requires of a dedicated E2E project.
  - Note, 2026-09-01: `rg -n 'coverage.thresholds|\.feature' apps/wahidyankf-www-e2e/project.json` finds nothing. The
    corpus reaches this project through `featuresRoot` in `playwright.config.ts` and through the `inputs` globs, never
    as a corpus of its own.
- [x] [AI] Run `npm install` — acceptance: `apps/wahidyankf-www-e2e` resolves as an npm workspace and
      `npx nx show projects` lists `wahidyankf-www-e2e`.
  - Note, 2026-09-01: `npx nx show projects` prints `wahidyankf-www-e2e`, `wahidyankf-www`, and `badakmini-cli`.
- [x] [AI] Edit `apps/wahidyankf-www-e2e/playwright.config.ts` to start the application under test with a `webServer`
      block running `npx nx run wahidyankf-www:start` instead of relying on a container — acceptance:
      `rg -n 'webServer' apps/wahidyankf-www-e2e/playwright.config.ts` finds the block and
      `rg -n 'docker' apps/wahidyankf-www-e2e` finds nothing. [AC-5]
  - Note, 2026-09-01: the `webServer` block runs `npx nx run wahidyankf-www:start` from `../..` and waits on
    `http://localhost:3201`, and `rg -rni 'docker' apps/wahidyankf-www-e2e` finds nothing. The comment explaining what
    replaced the container says "container image" rather than naming the tool, so the sweep stays true without the
    reasoning being lost. [AC-5]
- [x] [AI] Repoint `featuresRoot` in `apps/wahidyankf-www-e2e/playwright.config.ts` at
      `../../specs/apps/wahidyankf-www/behaviour` — acceptance:
      `rg -n 'featuresRoot' apps/wahidyankf-www-e2e/playwright.config.ts` names the new path.
  - Note, 2026-09-01: `featuresRoot` reads `../../specs/apps/wahidyankf-www/behaviour`.
- [x] [AI] Repoint the `@covers` comments in the eight step files at
      `specs/apps/wahidyankf-www/behaviours/<name>.feature` — acceptance:
      `rg -n 'behaviours/wahidyankf-www/gherkin' apps/wahidyankf-www-e2e` finds no stale path, and
      `rg -c '@covers' apps/wahidyankf-www-e2e/steps` still reports the same number of lines as before the edit, so no
      coverage note was dropped rather than repointed.
  - Note, 2026-09-01: `rg -n 'behaviours/wahidyankf-www/gherkin' apps/wahidyankf-www-e2e` finds nothing, and the
    `@covers` count is 36 before the edit and 36 after — nothing dropped. The source nests the corpus by domain
    directory and this repository keeps it flat, so the rewrite removed one path segment as well as renaming the root.
- [x] [AI] Rewrite the `missingSteps` comment block inside the `defineBddConfig` call in
      `apps/wahidyankf-www-e2e/playwright.config.ts`, which cites the `specs:e2e:coverage` target and
      `ayokoding-www-fe-e2e` — neither of which exists in this repository — replacing it with the named inventory of
      scenarios this adapter deliberately does not bind — acceptance:
      `rg -n 'specs:e2e:coverage|ayokoding' apps/wahidyankf-www-e2e/playwright.config.ts` finds nothing, and the
      surviving comment still says why `missingSteps` is not left at its `fail-on-gen` default and why a
      `tags: "not @unit"` filter was tried and rejected. That rejection is the reason the setting is glob-wide rather
      than tag-scoped, and it is stated nowhere else.
  - Note, 2026-09-01: `rg -n 'specs:e2e:coverage|ayokoding' apps/wahidyankf-www-e2e/playwright.config.ts` finds nothing.
    The replacement names all four unbound features with their scenario counts, and keeps both surviving reasons: why
    `missingSteps` is not left at `fail-on-gen`, and why a `tags: "not @unit"` filter was tried and rejected.
- [x] [AI] Rewrite the separate three-line comment above `process.env.APP_ENV ??= "test";` at the top of
      `apps/wahidyankf-www-e2e/playwright.config.ts`, which cites
      `plans/in-progress/restrict-env-access-to-prod-and-stag` — a plan that does not exist in this repository —
      repointing it at the loader contract as it lands here, `apps/wahidyankf-www/src/features/env/core/tier-env.ts` and
      `specs/apps/wahidyankf-www/behaviours/tier-env-loading.feature` — acceptance:
      `rg -n 'restrict-env-access' apps/wahidyankf-www-e2e/playwright.config.ts` finds nothing,
      `rg -n 'tier-env' apps/wahidyankf-www-e2e/playwright.config.ts` finds the repointed citation, and the comment
      still states both halves of why the pin exists: that an unset `APP_ENV` falls back to `local`, and that `local`
      reads a developer's real `.env.local` instead of test fixtures. This is a second comment in a different part of
      the same file, which is why it is a separate item: the one above edits the `defineBddConfig` block and its sweep
      would not reach here, while a reader who deleted this pin because its citation is dead would hand the suite a
      developer's real environment files.
  - Note, 2026-09-01: `rg -n 'restrict-env-access'` finds nothing and `rg -n 'tier-env'` finds the repointed citation at
    three lines. Both halves of the reason survive: that an unset `APP_ENV` falls back to `local`, and that `local` is
    the one tier whose stray-file guard is skipped, so it would hand the suite a developer's real `.env.local`.
- [x] [AI] Record the unit-only and integration-only scenario inventory in `apps/wahidyankf-www-e2e/README.md`: the four
      scenarios of `env-loader.feature`, the five of `tier-env-loading.feature`, and the eight of
      `port-resolver.feature`, all Node-process environment concerns with no browser equivalent; and the two CV export
      scenarios in `cv-export.feature`, which bind at the integration layer because the export is a build-time script no
      browser reaches — acceptance: `rg -c 'tier-env-loading' apps/wahidyankf-www-e2e/README.md`,
      `rg -c 'port-resolver' ...`, `rg -c 'env-loader' ...`, and `rg -c 'CV export' ...` each report a non-zero count,
      and the same nineteen-scenario inventory appears in [specification changes](tech-docs/specification-changes.md).
      [AC-5]
  - Note, 2026-09-01: all four `rg -c` probes report non-zero. The inventory is a table naming each feature, its
    scenario count, and the layer that binds it: `env-loader` 4, `tier-env-loading` 5, `port-resolver` 8 at the unit
    layer, and `cv-export` 2 at the integration layer. The same nineteen appear in
    [specification changes](tech-docs/specification-changes.md) at its E2E-binding section. [AC-5]
- [x] [AI] Record the generated skip baseline in two places: run `npx bddgen` inside `apps/wahidyankf-www-e2e`, count
      the generated `test.fixme` entries, write that number and the feature files behind it into
      `apps/wahidyankf-www-e2e/README.md` for a reader, and write the number alone into a new tracked file
      `apps/wahidyankf-www-e2e/e2e-skip-baseline.json` as `{"skippedScenarios": <count>}` for the target below to read —
      acceptance: every `fixme` entry traces to one of the four features named in the recorded inventory above, no entry
      traces to any other feature,
      `node -p "require('./apps/wahidyankf-www-e2e/e2e-skip-baseline.json').skippedScenarios"` prints an integer, and
      that integer equals the count the README states. The number counts **generated tests, not scenarios**, and the two
      do not match here: `playwright-bdd` calls `forceFixme()` once per generated test and generates one test per
      `Examples` row, and three of the four unbound features carry Scenario Outlines — `env-loader.feature` with a 3-row
      table, `tier-env-loading.feature` with a 3-row table, and `port-resolver.feature` with a 3-row and a 10-row table.
      The nineteen-scenario inventory therefore produces roughly 34 entries, so a baseline recorded as `19` would fail
      on its first run. For the same reason the check is worded by feature rather than by title: an outline row's
      generated title carries its parameter values and will not appear verbatim in a scenario-title list. The field name
      `skippedScenarios` is kept as written because the target below reads that key; the README states beside the number
      what it counts. `missingSteps: "skip-scenario"` lets generation succeed with an unbound scenario rendered as
      `test.fixme`, so the suite exits 0 whether the gap is the intended one or a newly-broken binding; the recorded
      baseline is what makes the two distinguishable. The number needs a machine-readable home because prose in a README
      is not something a target can compare against, and this file is what `specs:e2e:baseline` reads. It is a new
      authored file, not the `e2e-coverage-baseline.json` this plan declines to port: that one holds an `allowedUnbound`
      list for `rhino-cli` and nothing here consumes its shape. [AC-5]
  - Note, 2026-09-01: **34**, which is what the item predicted as "roughly 34" and not the nineteen a scenario count
    would give. Every `fixme` traces to one of the four named features — `cv-export` 2, `env-loader` 6, `port-resolver`
    19, `tier-env-loading` 7 — and a per-file sweep over the eight bound features finds none, so no entry traces
    anywhere else. `node -p "require('./apps/wahidyankf-www-e2e/e2e-skip-baseline.json').skippedScenarios"` prints `34`,
    matching what the README states beside an explanation of what it counts. [AC-5]
- [x] [AI] Give the `specs:e2e:baseline` target in `apps/wahidyankf-www-e2e/project.json` this verbatim command, with
      `cwd` set to `{projectRoot}` like every other single-command target here:
      `npx bddgen && expected=$(node -p "require('./e2e-skip-baseline.json').skippedScenarios") && actual=$(grep -rho 'test\.fixme(' .features-gen | wc -l | tr -d ' ') && if [ "$actual" != "$expected" ]; then echo "e2e skip baseline moved: expected $expected fixme scenarios, found $actual"; exit 1; fi`
      — acceptance: `npx nx run wahidyankf-www-e2e:specs:e2e:baseline` exits 0, and exits non-zero with that message
      after temporarily renaming one step file, and is restored to 0 when the file is renamed back; and
      `npx nx show project wahidyankf-www-e2e --json` shows `"default"` as the first entry of this target's `inputs`.
      Two details are load-bearing. The `\.` in the pattern is escaped once more when it is written into JSON, so
      `project.json` carries `test\\.fixme(`; an unescaped dot would also match `testXfixme(`. And `grep -rho` is used
      without a filename filter, so the count does not depend on the extension `playwright-bdd` happens to give its
      generated files. The target reads `e2e-skip-baseline.json` from the project directory and writes nothing;
      `.features-gen/` is ignored by the copied `apps/wahidyankf-www-e2e/.gitignore` and by the root rule Phase 1 adds,
      so a run leaves `git status --short` empty. Declare `"default"` first in its `inputs`, then
      `{projectRoot}/e2e-skip-baseline.json`, `{projectRoot}/steps/**/*.ts`, `{projectRoot}/playwright.config.ts`, and
      the shared corpus glob, so a corpus or binding edit invalidates the cache and the project's own files stay in the
      hash — a target-level `inputs` array replaces the default input set rather than extending it, so omitting
      `"default"` would drop everything the four explicit entries do not name. This is the one guarantee the dropped
      `specs:e2e:coverage` validator provided that nothing else here provides: without it, deleting a step file would
      look like a passing run. [AC-5]
  - Note, 2026-09-01: **the verbatim command as written could not pass, and the deviation is one character of shell
    rather than a change of intent.** `node -p` renders its result through `util.inspect`, which colourises a number
    when stdout is a TTY — and Nx runs a target's command under a pty. `expected` therefore held `\033[33m34\033[39m`
    while `actual` held `34`, so the comparison failed with the self-contradicting message
    `expected 34 fixme scenarios, found 34`. `NO_COLOR=1` did not suppress it. The fix reads the value without
    `util.inspect` at all:
    `node -e "process.stdout.write(String(require('./e2e-skip-baseline.json').skippedScenarios))"`, verified
    byte-for-byte under a pty — `node -p` emits `033 [ 3 3 m 3 4 033 [ 3 9 m`, the `-e` form emits `3 4`. Everything
    else is verbatim, including the doubly-escaped `test\\.fixme(` and the filter-free `grep -rho`. All three states
    then hold: the target exits 0; renaming `steps/theme.steps.ts` away makes it exit 1 printing
    `e2e skip baseline moved: expected 34 fixme scenarios, found 38`; renaming it back returns it to 0. `"default"` is
    first in its `inputs`, followed by the four explicit entries. A run leaves `git status --short` empty —
    `git check-ignore` confirms `.features-gen/` is ignored by the copied `.gitignore`. [AC-5]
- [x] [AI] Record why this E2E adapter gets a dedicated project at all in `apps/wahidyankf-www-e2e/README.md` — the
      [BDD policy](../../../repo-governance/development/behaviour-driven-development-policy.md) role matrix permits one
      only as a different-toolchain exception, and this one qualifies because the application's behaviour adapter runs
      `@amiceli/vitest-cucumber` under Vitest and jsdom while this adapter runs `playwright-bdd` under the Playwright
      runner against a downloaded Chromium binary, with its own generated test directory — acceptance:
      `rg -n 'playwright-bdd' apps/wahidyankf-www-e2e/README.md` finds the justification and it names both toolchains.
      Left unrecorded, a later reader sees only a second project where the policy defaults to co-location.
  - Note, 2026-09-01: recorded under `## Why this is a separate project`, naming the different-toolchain exception as
    the only ground and saying what the two toolchains are: `@amiceli/vitest-cucumber` under Vitest in jsdom for the
    application, `playwright-bdd` under the Playwright runner against a downloaded Chromium for this project. It also
    states what follows from existing under that exception and nothing more — no separate corpus, no unit layer, no
    numeric coverage gate.
- [x] [AI] Rewrite the `Run the suite` and `Checks and specs` sections of `apps/wahidyankf-www-e2e/README.md`, which
      describe building an isolated local production container and health-checking it, and which name four targets the
      new `project.json` does not define — `test:e2e:ui`, `test:e2e:report`, `test:quick`, and `test:specs` —
      acceptance: `rg -n 'container|docker|test:e2e:ui|test:e2e:report|test:specs' apps/wahidyankf-www-e2e/README.md`
      finds nothing, and every command the file still shows resolves to a target
      `npx nx show project wahidyankf-www-e2e --json` lists.
  - Note, 2026-09-01: both sections are rewritten. `Run the suite` now shows `install` then `test:e2e` under the new
    project name and explains that there is no container, that `webServer` starts the application, and that `dependsOn`
    builds it. `Checks and specs` became `Checks`, dropping the `test:specs` target this project does not have.
- [x] [AI] Repoint the corpus link at the end of `apps/wahidyankf-www-e2e/README.md` from
      `../../specs/apps/wahidyankf/behaviours/wahidyankf-www/gherkin/README.md`, which does not exist in this
      repository, at `../../specs/apps/wahidyankf-www/behaviours/README.md`, and add a link to
      `../../specs/apps/wahidyankf-www/architecture.md`, then run `git add -N apps/wahidyankf-www-e2e/README.md` before
      checking — acceptance: `npm run check:markdown-links` exits 0 and
      `rg -n 'architecture.md' apps/wahidyankf-www-e2e/README.md` finds the backlink. This README is copied rather than
      authored, so it is untracked here and invisible to the check until it is intent-added.
  - Note, 2026-09-01: the link now points at `../../specs/apps/wahidyankf-www/behaviours/README.md`, which exists, and
    `npm run check:markdown-links` exits 0.
- [x] [AI] Run `npx nx run wahidyankf-www-e2e:install`, then
      `rm -rf apps/wahidyankf-www/.next && npx nx run wahidyankf-www-e2e:test:e2e` — acceptance: the suite runs against
      `next start` with no Docker involved and exits 0, and `apps/wahidyankf-www/.next` exists again afterwards. The
      `rm -rf` is what makes this the cold-checkout proof rather than a run that happened to find a `.next` left by an
      earlier build: it is the only place `"dependsOn": ["wahidyankf-www:build"]` and `build`'s `{projectRoot}/.next`
      `outputs` declaration are exercised together. Without the removal, a `build` cache hit that restores nothing still
      leaves `next start` a directory to serve, and the trap [technical design](tech-docs/README.md) describes — a
      cached target with no `outputs`, replayed, reporting success, restoring nothing — passes unnoticed here and fails
      on a machine that has never built. [AC-5]
  - Note, 2026-09-01: `install` exits 0, and after `rm -rf apps/wahidyankf-www/.next` the suite exits 0 with **36
    passed, 34 skipped, 0 failed** — the skip count matching the recorded baseline exactly. Two defects had to be fixed
    first, both recorded as their own items below: a Playwright type mismatch that also turned out to be a runtime
    incompatibility, and a step file asserting a navigation label the application has never rendered.
- [x] [AI] Align the Playwright toolchain so one `playwright-core` is installed: pin `@playwright/test` to `1.62.1` and
      `playwright-bdd` to `9.2.0` in `apps/wahidyankf-www-e2e/package.json` — acceptance: `npm ls playwright-core --all`
      shows a single deduped version, `npx nx run wahidyankf-www-e2e:typecheck` exits 0, the suite runs, and
      `npm audit --audit-level=low` reports no vulnerability. **This item was added during execution.**
      `@axe-core/playwright@4.10.1` declares `playwright-core: ">= 1.0.0"`, an open range this plan's pinning items do
      not reach because it belongs to a transitive dependency's own manifest. npm therefore hoisted
      `playwright-core@1.62.1` beside the `1.60.0` that `@playwright/test@1.60.0` pins nested, and
      `AxeBuilder({ page })` took a structurally different `Page` from the one the step file holds:
      `TS2739: Type 'Page' is missing the following properties from type 'Page': localStorage, sessionStorage`. Two
      versions of one library in one dependency graph is the defect; a cast at the call site would have hidden it.
  - Note, 2026-09-01: this took two attempts, and the first is worth recording because it looked correct. Pinning
    `@playwright/test` to `1.62.1` alone deduped the graph and fixed the type error, but `playwright-bdd@8.5.1` then
    failed at runtime with `TypeError: Cannot destructure property 'registerESMLoader'` — it reaches into Playwright
    internals that 1.62 moved, so the types agreed while the code did not. Moving `playwright-bdd` to `9.2.0` as well
    fixed that, and paid for itself twice over: the pairing also cleared five moderate advisories that
    `playwright-bdd@8.5.1` carried through `@cucumber/gherkin` and `@cucumber/messages` into `uuid`, so
    `npm audit --audit-level=low` went from `5 moderate severity vulnerabilities` to `found 0 vulnerabilities`. This is
    the same call Phase 1 made when `tsx@4.21.0` carried GHSA-g7r4-m6w7-qqqr: move the pin, do not lower the bar or add
    an `overrides` entry. The v9 API is unchanged for everything used here — `defineBddConfig`, `createBdd`, and
    `missingSteps` all behave the same, and the generated `test.fixme` count is 34 on both versions. [AC-9]
- [x] [AI] Reconcile `specs/apps/wahidyankf-www/architecture.md` against the delivered tree, checking each of the three
      differences [specification changes](tech-docs/specification-changes.md#c4-model) says the model must record, and
      correcting the model where the delivered system disagrees — acceptance: all three hold, each with its own
      evidence, and any correction is a commit to `architecture.md` in this phase rather than a note deferred to
      Phase 7. **Container view**:
      `test ! -e apps/wahidyankf-www/Dockerfile && test ! -e apps/wahidyankf-www-e2e/scripts` succeeds and
      `rg -n 'webServer' -A5 apps/wahidyankf-www-e2e/playwright.config.ts` shows the suite starting
      `wahidyankf-www:start`, so the model's single Next-process node with the E2E adapter driving it is what was built.
      **Component view**: `rg -n 'open-sharia-enterprise' apps package.json` finds nothing, and
      `apps/wahidyankf-www/src/features/ui/shell/` and `apps/wahidyankf-www/src/features/env/core/` both exist, so the
      three external components really are internal modules and the relationship the model deletes really is gone.
      **System context**: `git ls-files cv` returns nothing and `apps/wahidyankf-www/src/features/cv/core/data.ts` is
      the only CV record, so the model's one store is the delivered count. This item exists because the model is
      authored in Phase 2 as an as-built description of a system that Phases 3, 4, and 5 are what actually build. The
      [architecture specification policy](../../../repo-governance/development/architecture-specifications.md) says the
      model "describes only the current, as-built system", and
      [specification changes](../../../repo-governance/conventions/plans-organization-policy/specification-changes.md)
      says to update the C4 model "only with the final implemented boundary" — neither is satisfied by authoring alone,
      and the Phase 7 reconciliation covers `[AC-1]` through `[AC-10]`, which never name the model. Phase 2 keeps the
      authoring, because the corpus and the model belong together for a reader; this is the check that makes the Phase 2
      claim true.
  - Note, 2026-09-01: all three hold and the model needed no correction, which is the outcome rather than a skipped
    check. **Container**: `test ! -e apps/wahidyankf-www/Dockerfile && test ! -e apps/wahidyankf-www-e2e/scripts`
    succeeds and the `webServer` block starts `wahidyankf-www:start`, matching the model's single Next-process node with
    the E2E adapter driving it. **Component**:
    `rg -n 'web-ui|web-ui-token|ts-env-loader' specs/apps/wahidyankf-www/architecture.md` finds nothing, and
    `src/features/ui/shell/` and `src/features/env/core/` both exist, so the three external components really are
    internal modules. The one `open-sharia-enterprise` match in the tree is the owner CV prose in
    `apps/wahidyankf-www/docs/linkedin-projects.md` that Phase 3 recorded as the deliberate survivor, not a package
    specifier. **System context**: `git ls-files cv` returns nothing and `data.ts` is the only CV record, so the model's
    one store is the delivered count.
- [x] [AI] Correct the navigation labels asserted in `apps/wahidyankf-www-e2e/steps/accessibility.steps.ts` from
      `["Home", "CV", "Personal Projects"]` to `["Home", "CV", "Independent Projects"]` — acceptance:
      `npx nx run wahidyankf-www-e2e:test:e2e` exits 0 with `Interactive controls expose accessible names` passing.
      **This item was added during execution.** The step file asserted a link named `Personal Projects`, and
      `Navigation` renders `Independent Projects` for the `/personal-projects` route. This is not damage the port
      caused: `ose-public`'s own `navigation.tsx` renders `Independent Projects` at the recorded source commit too, so
      the source step file asserted a label its own application never rendered. The scenario names no label at all — it
      says only that every navigation link exposes link text or an aria-label — so the list belongs to the binding and
      has to track the component. This repository's sibling binding of the same scenario,
      `apps/wahidyankf-www/tests/bdd/accessibility.steps.ts:137`, already asserts the corrected three, so the two
      bindings of one scenario disagreed and the E2E one was wrong.
  - Note, 2026-09-01: the suite went from `1 failed, 35 passed` to `36 passed`. A comment at the call site now says why
    the list is three labels rather than three route names, and names the sibling binding. Nothing in the corpus
    changed.
- [x] [AI] Edit `.github/workflows/full-bdd.yml` to add a Playwright browser install step before the verification run —
      acceptance, **amended during execution**: `rg -n 'wahidyankf-www-e2e:install' .github/workflows/full-bdd.yml`
      finds the step and `npm run check:workflows` exits 0. The criterion first probed for the literal
      `playwright install`, which would mean copying the install command into CI beside the copy in `project.json`; the
      step invokes the project's own `install` target instead, so the command keeps one home and the criterion names the
      target rather than the command it runs.
  - Note, 2026-09-01: a browser-install step now precedes the verification run, and `npm run check:workflows` exits 0 on
    the edited file. The criterion this item now carries is the amended one, and it is what was met: the step runs
    `npx nx run wahidyankf-www-e2e:install`. The criterion it carried when execution reached it probed for the literal
    `playwright install`, which would have meant copying the install command into CI beside the copy in `project.json` —
    two homes that drift apart silently, the failure this repository's governance exists to prevent. The target is the
    same one `workspace-commands.md` tells a developer to run, so the command keeps one home and CI still installs
    browsers before the run.
- [x] [AI] Append ` && nx run wahidyankf-www-e2e:test:e2e` to the `test:scheduled` script in root `package.json`,
      keeping it last so the quick-then-integration-then-E2E order holds — acceptance, **amended during execution**:
      `node -e "console.log(require('./package.json').scripts['test:scheduled'])"` prints the Phase 3 value with that
      clause appended, the script ending in
      ` && nx run wahidyankf-www-e2e:specs:e2e:baseline && nx run wahidyankf-www-e2e:test:e2e`. The criterion first read
      "and nothing else changed", which held when this item ran and stopped holding in Phase 7: the archival quality
      gate found `specs:e2e:baseline` had no automated caller and wired it into this same script, immediately ahead of
      the suite. A criterion that forbids every later edit to a shared script cannot survive one, so it names the
      delivered value instead. `npm run test:e2e` needs no edit: it is `nx run-many -t test:e2e` and reaches the new
      project through the target name alone.
  - Note, 2026-09-01: `test:scheduled` printed the Phase 3 value with ` && nx run wahidyankf-www-e2e:test:e2e` appended
    and nothing else changed when this item ran, keeping the quick-then-integration-then-E2E order with the new suite
    last. Phase 7's archival gate then edited the same script again, inserting
    ` && nx run wahidyankf-www-e2e:specs:e2e:baseline` immediately before the suite, so the criterion above is the
    amended one and the delivered script carries both clauses. The suite still runs last.
- [x] [AI] Edit the `## Current Applications` list in `apps/README.md` to index `wahidyankf-www-e2e` as the dedicated
      Playwright suite for `wahidyankf-www` — acceptance: `rg -n '\[`wahidyankf-www-e2e`\]' apps/README.md` finds the
      entry with a descriptive relative link to `wahidyankf-www-e2e/README.md`, and the list now holds three entries.
      The [readme-refresh](../../../repo-governance/workflows/readme-refresh.md) workflow puts every README below
      `apps/` in scope and asks for the smallest affected set in the same commit, which makes this an explicit item
      rather than something the application's own indexing item covers. Nothing automated catches an omission here:
      `apps/` is outside the [documentation index policy](../../../repo-governance/documentation-index-policy.md)'s
      scope, and `check:markdown-links` validates the links that exist rather than the entries that are missing, so a
      project left out of this list stays invisible until a person notices.
  - Note, 2026-09-01: the entry is present with a relative link to `wahidyankf-www-e2e/README.md`, the list now holds
    three entries, and `npm run check:markdown-links` exits 0.
- [x] [AI] Extend the language-target deviation already recorded in
      `repo-governance/development/testing-policy/tooling.md` to the second project, now that it exists:
      `apps/wahidyankf-www-e2e/tsconfig.json` sets `module` to `esnext`, `moduleResolution` to `bundler`, and `target`
      to `ES2022`, with `strict` true, and reaches that state the same way the application does — by extending
      `tsconfig.base.json`, which the item earlier in this phase adds, and overriding those three options on top of it —
      acceptance: `rg -n 'wahidyankf-www-e2e' repo-governance/development/testing-policy/tooling.md` finds the second
      project named in that sentence, `rg -n 'ES2022' repo-governance/development/testing-policy/tooling.md` finds its
      target, and `npm run check:governance` exits 0 against the 750-word cap. This is the E2E half of the Phase 3
      amendment, held back to this phase because `apps/wahidyankf-www-e2e/tsconfig.json` and its `extends` do not exist
      until now; writing it in Phase 3 would have stated a fact about a file no one could read. [AC-9]
  - Note, 2026-09-01: the sentence now names both projects. `rg -c 'wahidyankf-www-e2e'` and `rg -c 'ES2022'` each
    report 1, and `npm run check:governance` exits 0 against the cap. The E2E half carries its own reason rather than
    borrowing the application's: Next 16 is what forces the application, while this project follows the runner it hosts,
    because `playwright-bdd` generates ES modules and imports them through the Playwright runner's loader.
- [x] [AI] Add the E2E project's commands to `repo-governance/development/workspace-commands.md` —
      `npx nx run wahidyankf-www-e2e:install`, `:test:e2e`, and `:specs:e2e:baseline` in the `Narrower runs` block, and
      one sentence noting that this project owns no `test:quick` — acceptance:
      `rg -c 'wahidyankf-www-e2e' repo-governance/development/workspace-commands.md` reports at least four matches and
      `npm run check:governance` exits 0.
  - Note, 2026-09-01: `rg -c 'wahidyankf-www-e2e' repo-governance/development/workspace-commands.md` reports 4 and
    `npm run check:governance` exits 0. The three commands are in the `Narrower runs` block and the sentence beside it
    states that this project owns no `test:quick` and links to the two policies that decide it, rather than asserting
    the rule freestanding — see the Rules Propagation note on the gate item below.
- [x] [AI] Run `npm run check:workflows` — acceptance: exits 0 on the edited workflow. [AC-8]
  - Note, 2026-09-01: exits 0. [AC-8]
- [x] [AI] Sweep the repository for the source project's old name — acceptance:
      `rg -n 'wahidyankf-www-fe-e2e' --hidden --glob '!node_modules' .` matches only lines inside this plan's own
      documents, which name the source path and the old manifest name in order to record the rename, and nothing under
      `apps/`, `specs/`, `repo-governance/`, `.github/`, or the root manifests. The acceptance is worded by location
      rather than by count on purpose: the number of plan lines carrying the old name changes whenever an item is split
      or reworded, and a count baked in here would go stale without anything breaking. The copied `package.json` and
      `README.md` both carry the old name at the source commit, so an unrepaired occurrence would leave the npm
      workspace name disagreeing with the Nx project name. [AC-5]
  - Note, 2026-09-01: every match is inside this plan's own documents — `README.md`, `delivery.md`,
    `tech-docs/README.md`, and `learnings.md` — where the old name records the source path and the rename. Nothing under
    `apps/`, `specs/`, `repo-governance/`, `.github/`, or the root manifests.
  - Correction, 2026-09-01: that note was wrong when it was written. The archival gate re-ran the sweep and found one
    match under `apps/` — a comment in `apps/wahidyankf-www/tests/bdd/accessibility.steps.ts` routing a reader at
    `apps/wahidyankf-www-fe-e2e/steps/accessibility.steps.ts`, a path this repository does not have. It was repointed at
    `apps/wahidyankf-www-e2e/steps/accessibility.steps.ts` and the sweep now passes as written. The acceptance was
    correct and the item was ticked against a reading of it that was not; this is the third time in this plan a finished
    item's record was found wrong by reading rather than by a check. [AC-5]
- [x] [AI] Reformat the ported E2E project to this repository's Prettier configuration — run `npm run format` —
      acceptance: `npm run format:check` exits 0. The thirteen files copied at the head of this phase arrive formatted
      to `$SRC/.prettierrc.json` rather than to this repository's, the same three differences the Phase 1 and Phase 3
      reformat items name and [technical design](tech-docs/README.md#toolchain-conformance-and-its-fallback) records. It
      sits last in the phase for the same reason those two do: after every copy, rename, and repoint here, so nothing is
      reformatted twice, and before the gate, so the gate's `format:check` does not depend on a step that has not run.
      The diff is a formatting change and not a content change. Nothing is added to an ignore file:
      `apps/wahidyankf-www-e2e/**` is held to the repository's formatting standard like every other tracked path, and
      [file impact](tech-docs/file-impact.md) rules out introducing a root `.prettierignore` at all.
  - Note, 2026-09-01: `npm run format:check` exits 0. Unlike the Phase 3 run this one was not a no-op: Prettier
    rewrapped the ported step files and the authored README table, because the source is formatted to `printWidth: 120`
    and this repository leaves Prettier's 80 default in force. The eleven ported `.feature` files were digested before
    and after and diff empty, so the Phase 2 comparisons still hold.

### Phase 5 Gate

> Every check below passes before Phase 6 begins. A failure is fixed inside Phase 5.

- [x] [AI] Run `npx nx run wahidyankf-www-e2e:test:e2e` — acceptance: exits 0. [AC-5]
  - Note, 2026-09-01: exits 0 — 36 passed, 34 skipped, 0 failed — run after `rm -rf apps/wahidyankf-www/.next`, so the
    `dependsOn` build is exercised rather than assumed. [AC-5]
- [x] [AI] Run `npm run test:quick` — acceptance: exits 0 for the two projects that define `test:quick`, `badakmini-cli`
      and `wahidyankf-www`. The dedicated E2E project defines none, because the
      [testing policy](../../../repo-governance/development/testing-policy.md) puts a process E2E target outside
      `test:quick` and gives such a project the equivalent `typecheck`, `lint`, and `test:e2e` targets instead; its
      `typecheck` and `lint` are reached by `npm run typecheck` and `npm run lint`.
  - Note, 2026-09-01: exits 0. Nx runs it for `badakmini-cli` and `wahidyankf-www` only; `wahidyankf-www-e2e` defines no
    `test:quick` and is correctly absent rather than skipped.
- [x] [AI] Run `npm run typecheck` and `npm run lint` — acceptance: both exit 0 across all three projects, which is
      where the E2E project's own type-check and lint are covered.
  - Note, 2026-09-01: both exit 0 across all three projects, which is where this project's own type-check is covered —
    it has no `test:quick` to carry one.
- [x] [AI] Run `npm run check:workflows` — acceptance: exits 0. [AC-8]
  - Note, 2026-09-01: exits 0 on the edited `full-bdd.yml`. [AC-8]
- [x] [AI] Run `npm run format:check` and `npm run check:markdown-links` — acceptance: both exit 0. [AC-8]
  - Note, 2026-09-01: both exit 0. [AC-8]
- [x] [AI] Run `npm run check:governance` — acceptance: exits 0 on the two `repo-governance/development/` documents this
      phase edited, `workspace-commands.md` and `testing-policy/tooling.md`, both of which stay under the 750-word cap.
      [AC-8]
  - Note, 2026-09-01: exits 0 on both edited documents, `workspace-commands.md` and `testing-policy/tooling.md`, each
    still under its word cap.
- [x] [AI] Run `git add -A` and then `npm run check:rule-change` — acceptance: it names both `repo-governance/` edits
      and the workflows they trigger, and those workflows are followed rather than dismissed. This phase changes rule
      paths under the [rule change trigger policy](../../../repo-governance/development/rule-change-trigger-policy.md),
      and the check reads only staged paths, so it is staged first for the same reason the Phase 3, Phase 4, and Phase 6
      gates stage first: run against an unstaged tree it sees no paths and reports nothing, which reads identically to a
      clean result.
  - Note, 2026-09-01: the check names
    `Rules Propagation automatically triggered by repo-governance/development/testing-policy/tooling.md, repo-governance/development/workspace-commands.md`,
    and it was worked rather than dismissed. Harness Alignment is correctly not named: no instruction file, harness
    directory, or `opencode.json` changed in this phase. **Rules Propagation, in one sentence:**
    `apps/wahidyankf-www-e2e` is a second TypeScript project whose language target deviates from the code style policy,
    and its three commands and its lack of a `test:quick` belong in the commands document. The inventory found three
    documents in scope — `tooling.md` as the deviation register Phase 3 established, `workspace-commands.md` as
    canonical for commands, and `testing-policy.md`, which owns the target contract. **The idempotency gate stopped a
    rule from being added.** The sentence first written into `workspace-commands.md` asserted freestanding that this
    project owns no `test:quick` and why; but `testing-policy.md` already says a permitted dedicated project owns
    _equivalent targets_ rather than the full set, and the
    [BDD policy](../../../repo-governance/development/behaviour-driven-development-policy.md) role matrix already says
    such a project has no separate corpus, unit layer, or numeric coverage gate. Read together those two settle it, so
    no new rule was warranted and the sentence was rewritten to link to both instead of restating them. That is also
    what rules out the shape the source used: `ose-public`'s `project.json` gives this project `test:unit`,
    `test:coverage`, and `test:quick` as `echo 'no-op'` placeholders, and an `echo` standing in for a test is the
    failure mode the testing policy names, so absence is the correct signal and this project declares none of the three.
- [x] [AI] Commit and push the phase to `main` — acceptance: `git status --short` is empty.
  - Note, 2026-09-01: `git status --short` is empty after the push. The phase landed in three commits under the
    [thematic commits policy](../../../repo-governance/conventions/thematic-commits-policy.md): the E2E project with its
    manifest and CI step, the two governance documents, and this record. **This checkbox was ticked late** — the seven
    gate checks above were run and recorded first and this one was missed in the same pass, which is the second time in
    this plan a finished item kept an unticked box. Recorded plainly rather than backdated.

> **Pause Safety**: All three behaviour layers exist and run: unit and behaviour in the application, integration against
> the real filesystem, and process E2E in a real browser against `next start`. Scheduled verification covers all three.
> Nothing is deployed. Safe to stop. Resume with `npm run test:quick`.

## Phase 6: Deployment Configuration and Its Governance Rule

Configuration lands dormant. No branch is created and no deploy is triggered.

- [x] [AI] Copy `$SRC/apps/wahidyankf-www/vercel.json` to `apps/wahidyankf-www/vercel.json` unchanged, record its digest
      with the verbatim command
      `shasum -a 256 apps/wahidyankf-www/vercel.json > plans/in-progress/wahidyankf-www-migration/evidence/vercel-json-digest.txt`
      before anything reformats it, and convert that file's entry in the `## Directory Map` of `evidence/README.md` to a
      relative link in this same item — acceptance: `shasum -a 256 "$SRC/apps/wahidyankf-www/vercel.json"` names the
      same hash the evidence file now holds, and the map entry is a link that `npm run check:markdown-links` resolves at
      this phase's gate. The digest is taken from the destination rather than the source so that the record names the
      path in this repository, and the two agree only while the copy is untouched, which is why the command runs before
      the item below. This is the provenance proof, and it is the eighth and last evidence file, so the conversion
      belongs here for the same reason every other evidence item carries it: [evidence](evidence/README.md) states that
      the item writing a file is the item that links it. [AC-10]
  - Note, 2026-09-01: `shasum -a 256` against the source names
    `8cbaad2e07a9c25f05fe7afda9807b348918541822277d2d4bbfb448058dee29`, and the evidence file holds that same hash
    against the destination path. The map entry in [evidence](evidence/README.md) is now a relative link, which makes
    eight linked entries and eight files. [AC-10]
- [x] [AI] Run `npm run format` over the copied `apps/wahidyankf-www/vercel.json` and confirm the configuration is
      unchanged — acceptance: `npm run format:check` exits 0, and
      `node -e 'const a=require("child_process").execSync("git -C '"$SRC"' show HEAD:apps/wahidyankf-www/vercel.json");const b=require("fs").readFileSync("apps/wahidyankf-www/vercel.json");if(JSON.stringify(JSON.parse(a))!==JSON.stringify(JSON.parse(b)))process.exit(1)'`
      exits 0, proving the parsed configuration is identical after reformatting. The source file is formatted at
      `ose-public`'s `printWidth: 120`, so its single-line `headers` objects exceed this repository's 80-column default
      and `prettier --check` fails on it; both the Phase 6 and Phase 7 gates run `format:check`. The owner's settled
      decision was to port the deploy configuration unchanged, which is about the configuration and not its whitespace:
      Vercel reads the parsed JSON, so `installCommand`, `buildCommand`, `ignoreCommand`, and every header survive
      reformatting exactly. This mirrors the `cv-raw.md` shape — copy, prove the digest, then apply one named amendment
      in the same phase. [AC-10]
  - Note, 2026-09-01: `npm run format:check` exits 0 and the deep-equality probe exits 0, so the parsed configuration is
    identical after reformatting. The bytes are not: the file now hashes
    `3076ab9b974dba664b915189b7f7b630cad8be3667ea0c92769f246f5f0c9965` against the `8cbaad2e...` recorded above, which
    is exactly why the digest was taken before this ran. Prettier broke the single-line `headers` objects across lines
    at this repository's 80-column default; `installCommand`, `buildCommand`, `ignoreCommand`, and all five headers
    survive unchanged. [AC-10]
- [x] [AI] Author `repo-governance/development/deployment-policy.md` with the standard `tldr` and `when_to_use` front
      matter, stating that `main` is the delivery target, that a `prod-<project>` branch is a promotion pointer no plan
      advances automatically, that Vercel gates each project's build on its own `prod-` branch, and that the domain
      cutover for this application is not authorized by this plan — acceptance: the file exists,
      `npm run check:governance` exits 0 on it, and
      `rg -n 'prod-wahidyankf-www' repo-governance/development/deployment-policy.md` finds the concrete example. [AC-10]
  - Note, 2026-09-01: the file exists at 379 words with the standard front matter, `npm run check:governance` exits 0 on
    it, and `rg -n 'prod-wahidyankf-www'` finds the concrete example. It states all four required things and one more
    the phase needs: that deploy configuration may be committed for a project whose promotion branch does not exist yet,
    which is the state this phase leaves the repository in. [AC-10]
- [x] [AI] Edit `repo-governance/development/README.md` to index the new policy — acceptance:
      `rg -n 'deployment-policy' repo-governance/development/README.md` finds the entry.
  - Note, 2026-09-01: the entry is indexed alphabetically among its peers and `rg -n 'deployment-policy'` finds it.
- [x] [AI] Edit `AGENTS.md` to reference the deployment policy in one sentence, without restating it — acceptance:
      `rg -n 'deployment-policy' AGENTS.md` finds the reference and `npm run check:governance` exits 0.
  - Note, 2026-09-01: one sentence in `Testing and Commits`, and `npm run check:governance` exits 0. The sentence went
    in restating that plans deliver to `main` and came out saying only what is new — that landing on `main` is not
    deploying — because Rules Propagation found the restatement; see the gate note below.
- [x] [AI] Confirm no `prod-wahidyankf-www` branch exists in this repository — acceptance: `git branch --list 'prod-*'`
      prints nothing, because this plan lands configuration only.
  - Note, 2026-09-01: `git branch --list 'prod-*'` prints nothing. The configuration is present and inert, which is the
    state the new policy names as expected for a project mid-migration.
- [x] [AI] Run `git add -A` and then `npm run check:rule-change` — acceptance: the trigger names the new
      `repo-governance/development/deployment-policy.md` along with the `README.md` and `AGENTS.md` edits, and the
      reported workflow is followed rather than dismissed. Staging comes first because the check reads only
      `git diff --cached --name-only`, so an unstaged run reports nothing and cannot be told apart from a run with no
      rule change in it.
  - Note, 2026-09-01: the check names
    `Rules Propagation automatically triggered by AGENTS.md, repo-governance/development/README.md, repo-governance/development/deployment-policy.md`
    and also `Run repo-governance/workflows/harness-alignment.md`, and both were worked. **Rules Propagation** found one
    defect, in the edit this phase had just made. The sentence added to `AGENTS.md` opened
    `Plans deliver to \`main\``— which`AGENTS.md`already says in its Planning section and which [plans organization policy](../../../repo-governance/conventions/plans-organization-policy.md) owns — so the file would have stated one rule twice and either copy could drift. It was rewritten to carry only what is new: that landing on`main`is not deploying. The inventory sweep for existing statements about deployment, promotion, or a`prod-`branch found nothing anywhere else in`AGENTS.md`, `CLAUDE.md`, or `repo-governance/`, so the new policy is the first and only home for those rules and no contradiction needed settling. The one boundary it touches is the delivery target, which it links to rather than restates. **Harness Alignment** found nothing to change: `rg
    -ni 'deploy|prod-|vercel'
    CLAUDE.md`matches nothing, so no derivative carried a stale or conflicting statement, and`npm run
    check:harness-parity` exits 0.

### Phase 6 Gate

> Every check below passes before Phase 7 begins. A failure is fixed inside Phase 6.

- [x] [AI] Run `npm run check:governance` — acceptance: exits 0. [AC-8]
  - Note, 2026-09-01: exits 0. [AC-8]
- [x] [AI] Run `npm run check:markdown-links` and `npm run format:check` — acceptance: both exit 0. [AC-8]
  - Note, 2026-09-01: both exit 0. [AC-8]
- [x] [AI] Run `npm run check:harness-parity` — acceptance: exits 0, confirming no harness gained a capability the
      others lack.
  - Note, 2026-09-01: exits 0. No harness gained or lost a capability — the six subagents remain mirrored across all
    three directories, and this phase edited no harness file.
- [x] [AI] Run `npm run test:quick` — acceptance: exits 0 for the two projects that define `test:quick`, `badakmini-cli`
      and `wahidyankf-www`.
  - Note, 2026-09-01: exits 0 for `badakmini-cli` and `wahidyankf-www`.
- [x] [AI] Commit and push the phase to `main` — acceptance: `git status --short` is empty.
  - Note, 2026-09-01: `git status --short` is empty after the push. The phase landed in three commits under the
    [thematic commits policy](../../../repo-governance/conventions/thematic-commits-policy.md): the deploy configuration
    with its digest, the governance rule that explains what the branch it names is for, and this record.

> **Pause Safety**: The deploy configuration is present and inert, the branch it names does not exist here, and
> governance now explains what that branch is for. `ose-public` is still the only repository serving the site. Safe to
> stop. Resume with `npm run check:governance`.

## Phase 7: Knowledge Capture and Archival

- [x] [AI] Reconcile every acceptance criterion `[AC-1]` through `[AC-10]` in `prd.md` against the delivered system,
      recording the proving command and its result for each — acceptance, **amended during execution to name the
      destination**: the reconciliation is written into a `## Reconciliation` section of `prd.md`, beside the criteria
      it reconciles, and all ten are marked satisfied with evidence, or an unsatisfied one is written up with the
      reason.
  - Note, 2026-09-01: all ten are satisfied, each with the command that proved it and its result, in a
    `## Reconciliation` section of `prd.md`. Two results exceed what was asked: `[AC-2]` and `[AC-3]` were written
    against a 99% floor and the delivered system reaches **100% lines** on both, with no threshold lowered and no
    exclusion added. `[AC-7]`'s pattern does match two tracked files outside this plan's documents —
    `apps/badakmini-cli/README.md`, whose opening paragraph explains that Badak Mini's command grammar follows a slice
    of `rhino-cli`, and `plans/done/2026-08-23__badakmini-layered-bdd/brd.md`, which names F# in a non-goal — which is
    why the sweep the item runs is scoped rather than repository-wide. Neither is a target, a script, or a workflow,
    which is what the criterion is about, and neither may be edited to make a wider sweep pass. The destination is
    `prd.md` rather than a new evidence file because the confirmation item below fixes the evidence directory at eight
    files and no ninth; a reconciliation of the acceptance criteria also reads best beside them.
- [x] [AI] Record a dated, evidence-backed `Not triggered` disposition for every dormant recovery item that did not fire
      — the TypeScript version fallback, the Biome fallback, the ESLint fallback, the Phase 2 scenario-split item, the
      Phase 3 `tsx` removal, the Phase 3 `output-path check` clause inside the second GREEN step, the Phase 4 item that
      adds a missing role to `data.ts`, and the Phase 5 `@vitejs/plugin-react` removal clause — acceptance: each of the
      eight carries a dated disposition rather than a completion mark, and each one that did fire carries the evidence
      instead. The Phase 3 unused-dependency removal is not on this list: `react-icons` and `class-variance-authority`
      have no importer at the recorded source commit, so that item is expected to fire and carries evidence rather than
      a disposition.
  - Note, 2026-09-01: all eight carry a dated disposition. Five are standalone checkboxes left deliberately unticked
    with a `**Not triggered, 2026-09-01.**` bullet — the Phase 2 scenario split, the Phase 3 `tsx` removal, and the
    TypeScript, Biome, and ESLint fallbacks. Two are clauses inside otherwise-ticked items and carry a
    `**Clause disposition, 2026-09-01: not triggered.**` note: the `output-path check` clause in the second GREEN step,
    and the Phase 4 clause that would have added a role missing from `data.ts`. The eighth **did** fire and carries
    evidence instead of a disposition: the Phase 5 `@vitejs/plugin-react` clause took its removal branch, because no
    file in that project imports it. Each of the five standalone bullets replaced an earlier note that had promised the
    disposition in future tense; carrying both would have said the same thing twice, once in a tense the delivered plan
    contradicts.
- [x] [AI] Confirm `repo-governance/development/testing-policy/tooling.md` records every deviation this plan produced —
      acceptance: `git log --oneline -- repo-governance/development/testing-policy/tooling.md` shows at least one commit
      from this plan, and the document names two certain deviations — the `module`, `moduleResolution`, and `target`
      departure from `code-style-policy.md`'s language target, stated for both `apps/wahidyankf-www` and
      `apps/wahidyankf-www-e2e` after Phase 3 and Phase 5 each wrote their half, and the Biome boundary, that
      `formatter.enabled` and `assist.enabled` are `false` so Prettier remains the formatting source of truth — plus
      each of the three toolchain components — TypeScript, Biome, ESLint — that could not conform. A component that did
      conform is named nowhere, which is the correct absence. [AC-9]
  - Note, 2026-09-01: `git log --oneline` on that file shows three commits, two of them this plan's — `73b5c32` for
    Phase 3 and `0c7f80c` for Phase 5 — and the register holds every deviation this plan produced: the language target
    for both TypeScript projects, and the Biome-as-linter-only boundary. A sentence there reading "that project" was
    disambiguated to "both TypeScript projects" in this pass, because the paragraph above it now names two. [AC-9]
- [x] [AI] Triage every entry in `learnings.md` to exactly one durable home — a rule in `repo-governance/`, a document
      in `docs/`, a subagent or skill instruction, code or a test, a new two-pager in `plans/ideas/`, or discarded with
      a one-line reason — acceptance, **amended during execution to make the routing verifiable**: no entry remains
      untriaged, or the plan records the explicit escape `No generalizable learnings — <reason>`; `learnings.md` carries
      a table naming each entry's home; every path the triage edits is added to `tech-docs/file-impact.md` as an `[E]`
      or `[N]` entry in the same change; and `git add -A && npm run check:rule-change` names the workflows the routed
      edits trigger, which are then followed. The item as first written named six categories of destination but no path
      and no verification, and the destinations cannot be known before the learnings exist.
  - Note, 2026-09-01: no entry remains untriaged and the escape was not needed. Twelve entries reach six documents
    rather than twelve — `delivery-checklists.md`, `dependency-selection-policy.md`, `code-style-policy.md`,
    `testing-policy/tooling.md`, `behaviour-driven-development-policy.md`, and the three mirrored `plan-checker` prompts
    — because several entries are one lesson met in different places, and a rule stated once holds better than the same
    rule stated three times. Two halves are discarded with a reason rather than deleted silently. `learnings.md` carries
    the routing table, all seven newly-reached paths are now `[E]` entries in [file impact](tech-docs/file-impact.md),
    and the staged `check:rule-change` at the gate names the workflows the edits trigger. The `delivery-checklists.md`
    rule is mirrored into all three `plan-checker` prompts in the same change, as that document's own last line requires
    — which matters here more than usual, because the vacuous-pass defect this plan hit in Phase 0 survived seven strict
    `plan-checker` cycles and the prompt is what would have caught it.
- [x] [AI] Confirm every surviving learning is free of secrets and generalizable beyond this one migration — acceptance,
      **amended during execution to make it observable**: `learnings.md` records the result of both checks over the
      whole set, naming which environment variables appear and confirming each appears as a name rather than a value;
      and no routed rule as written in its durable home names this migration, this application, or `ose-public`, which
      `rg -n 'wahidyankf|ose-public|migration' <routed files>` confirms — matching, outside the routed sections
      themselves, only two lines that are correct and stay: `testing-policy/tooling.md`'s `## Recorded Deviations`,
      which names `apps/wahidyankf-www` and `apps/wahidyankf-www-e2e` because recording which projects deviate is what a
      register is for, and `delivery-checklists.md`'s "matching `ose-public` so a migrated plan needs no translation", a
      sentence that predates this plan. The item as first written stated a bar with no command and no artifact.
  - Note, 2026-09-01: both checks pass over the whole set and `learnings.md` records the result. **Secrets:** no entry
    names a credential or token; the only environment variables any of them names are `APP_ENV`, `NO_COLOR`, and
    `WAHIDYANKF_WWW_PORT`, each as a name whose behaviour is the lesson and never with a value. **Generalizable:**
    `rg -n 'wahidyankf|ose-public|migration'` over the routed files matches no routed rule, and matches exactly the two
    lines the criterion names as correct survivors — the deviation register's project names and the `ose-public` clause
    that predates this plan. Neither is a rule this triage wrote. That is the mechanical form of the bar — it is also
    why the three pin entries collapse into one rule about inherited and transitive pins rather than three about three
    packages.
- [x] [AI] Confirm the `## Directory Map` of `evidence/README.md` is complete and fully linked — `phase-0-baseline.md`,
      `phase-1-toolchain.md`, `phase-2-background-coverage.md`, `unused-importers.txt`, `node-type-stripping.md`,
      `phase-3-measurements.md`, `cv-references.txt`, and `vercel-json-digest.txt` — acceptance: each of the eight
      entries is a `[name](name)` link, `ls plans/in-progress/wahidyankf-www-migration/evidence` lists exactly those
      eight files plus `README.md`, and `npm run check:markdown-links` exits 0. This is a confirmation, not the
      conversion: each link is made by the Phase 0 to Phase 6 item that writes its file, so no entry sits unlinked
      beside a file that already exists. The completeness half still needs stating here, because nothing else counts the
      directory against the map, and `check:markdown-links` validates the links that exist rather than the entries that
      are missing, so a missing or unlinked entry stays invisible to every gate. If one of the eight was never written,
      this item records which and why in `learnings.md` and leaves that entry unlinked, rather than linking a path that
      does not exist, which would fail that check.
  - Note, 2026-09-01: eight files on disk, eight linked entries, and the two sets match exactly in both directions — no
    unlinked file and no link without a file. The eighth, `vercel-json-digest.txt`, was written and linked by its own
    Phase 6 item, which is the pattern [evidence](evidence/README.md) states: the item that writes a file is the item
    that links it.
- [x] [AI] Append ` && nx run wahidyankf-www-e2e:specs:e2e:baseline` to the `test:scheduled` script in root
      `package.json`, immediately before the `wahidyankf-www-e2e:test:e2e` clause — acceptance:
      `node -e "console.log(require('./package.json').scripts['test:scheduled'])"` prints a script ending in
      ` && nx run wahidyankf-www-e2e:specs:e2e:baseline && nx run wahidyankf-www-e2e:test:e2e`. **This item was added
      during execution.** Phase 5 authored the target and Phase 5's gate ran it by hand, but nothing automated ever
      called it: `npm run test:e2e` is `nx run-many -t test:e2e` and reaches only that target name, and the nightly
      `full-bdd.yml` runs `npm run test:scheduled`. A check that guards against an unbound scenario is worthless if it
      runs only when someone remembers it, and the suite it guards exits 0 in exactly the case it exists to catch. It
      goes ahead of the suite because it is the cheaper of the two and its failure makes the suite's exit code
      untrustworthy. [AC-5]
  - Note, 2026-09-01: the script prints the two clauses in that order. Found by the third archival gate cycle, which
    asked what invokes the target and found the answer was nothing.
- [x] [AI] Record the new `test:scheduled` order and its reason in `repo-governance/development/workspace-commands.md`,
      whose `npm run test:scheduled` bullet states what that script runs — acceptance:
      `rg -n 'E2E skip baseline' repo-governance/development/workspace-commands.md` finds the bullet naming the baseline
      in the operational order and the sentence stating why it runs ahead of the suite, and `npm run check:governance`
      exits 0. **This item was added during execution.** That document is canonical for every command in this workspace,
      so a script whose order changed without it says two different things about the same command. This is a separate
      item from the edit above for the same reason Phases 3 and 5 each split their script edit from its documentation
      item. [AC-8]
  - Note, 2026-09-01: the bullet names the baseline between integration coverage and E2E, and gives the reason — the
    suite exits 0 on an unbound scenario and the baseline is what fails on one. `check:governance` exits 0 with the
    document at 602 words, clear of the headroom band.
- [x] [AI] Mirror the two `delivery-checklists.md` clauses the learnings triage added but did not carry across into
      `.claude/agents/plan-checker.md`, `.codex/agents/plan-checker.toml`, and `.opencode/agents/plan-checker.md` — the
      acceptance-criterion clause requiring that a criterion reading a tool's own output name a command no shell wrapper
      rewrites, and the recovery clause requiring a trigger to be read against its wording rather than against the
      presence of a failure — acceptance:
      `rg -c 'no shell wrapper rewrites|read against its wording' .claude/agents/plan-checker.md .codex/agents/plan-checker.toml .opencode/agents/plan-checker.md`
      prints `2` for each of the three files, and a three-way diff of the two changed sentences across the three prompts
      shows no difference — the same pair of proofs the Phase 4 `rules-checker` item uses, because
      `npm run check:harness-parity` compares which subagents exist rather than what their prompts say and cannot stand
      alone here. **This item was added during execution.** The triage item above added three clauses to the policy and
      mirrored one; `npm run check:harness-parity` cannot catch the other two, because it compares which subagents exist
      rather than what their prompts say, and all three prompts were equally incomplete. Two places in this plan already
      asserted the mirroring was complete, so nothing would have looked again. [AC-8]
  - Note, 2026-09-01: all three print `2` and the two added sentences are byte-identical across the harnesses. Found by
    the third archival gate cycle, by diffing the policy against the prompts rather than by any check.
- [x] [AI] Relocate the `## Execution Record` section out of
      `repo-governance/conventions/plans-organization-policy/delivery-checklists.md` into a new sibling,
      `repo-governance/conventions/plans-organization-policy/execution-record.md`, leaving a one-line pointer behind,
      indexing the new document in `repo-governance/conventions/plans-organization-policy.md` and
      `repo-governance/conventions/plans-organization-policy/README.md`, and repointing the one inbound anchor in
      `repo-governance/workflows/plan-execution/02-phase-loop.md` — acceptance:
      `rg -n 'execution record' repo-governance/conventions/plans-organization-policy/delivery-checklists.md` finds the
      pointer line and no `## Execution Record` heading remains in that file,
      `rg -n 'execution-record.md' repo-governance/conventions/plans-organization-policy.md repo-governance/conventions/plans-organization-policy/README.md repo-governance/workflows/plan-execution/02-phase-loop.md`
      returns three lines, and `npm run check:governance` and `npm run check:markdown-links` both exit 0. **This item
      was added during execution.** The triage's own additions pushed `delivery-checklists.md` to 737 words, inside the
      [document word limit policy](../../../repo-governance/conventions/document-word-limit-policy.md)'s 700-word
      headroom band, which only relocation closes. The Execution Record is a separate artifact from a checkbox — a dated
      log at the top of `delivery.md` rather than a rule about checkbox shape — so the split is one the document wanted
      on its own terms, not only for the word count. [AC-8]
  - Note, 2026-09-01: the pointer is at `delivery-checklists.md` line 10, the three index and anchor lines resolve, and
    both checks exit 0. The two resulting documents are 541 and 271 words, so neither sits in the band. Found by the
    third archival gate cycle.
- [x] [AI] Run the [plan-quality-gate](../../../repo-governance/workflows/plan-quality-gate.md) workflow at strict level
      — acceptance: two consecutive runs report zero findings, or seven cycles have run with every finding of every
      cycle resolved rather than waived, which is the other ending that workflow provides for; either way the result is
      appended to the Quality Gate section of this plan's `README.md` rather than replacing what the pre-commit run
      recorded there.
  - Note, 2026-09-01: seven cycles ran and every finding of every cycle was resolved rather than waived; the
    two-consecutive-clean ending was not reached. The result is appended to the `## Quality Gate` section of
    [`README.md`](README.md) below the pre-execution line, which stays as it was. Five of the seven cycles are archival
    cycles run against the executed plan; cycles 1 and 2 ran at the start of this phase. The record above carries what
    each cycle found.
- [x] [AI] Move `plans/in-progress/wahidyankf-www-migration/` to `plans/done/<YYYY-MM-DD>__wahidyankf-www-migration/`,
      where `<YYYY-MM-DD>` is the date the final commit lands, guarding first that the destination does not exist —
      `dest="plans/done/$(date +%F)__wahidyankf-www-migration" && test ! -e "$dest" && git mv plans/in-progress/wahidyankf-www-migration "$dest"`
      — acceptance, **amended during execution to name the command and its guard**: the guard is what refuses an
      existing destination, so the move cannot run at all if one is there; afterwards
      `test ! -e plans/in-progress/wahidyankf-www-migration` succeeds,
      `ls -d plans/done/*wahidyankf-www-migration | wc -l` prints `1`, and `ls plans/done` shows the folder named with
      the date and no invented suffix.
      [Lifecycle moves](../../../repo-governance/conventions/plans-organization-policy/lifecycle-moves.md) requires
      refusing an already-existing destination and never merging, overwriting, or inventing a suffix; the item as first
      written stated all three as outcomes with no mechanism behind any of them, which left the plan's last irreversible
      step to be improvised. `git mv` rather than `mv` keeps the rename tracked as one, so the archived history stays
      readable.
  - Note, 2026-09-01: the guard ran first and the destination did not exist, so `git mv` moved the folder to
    `plans/done/2026-09-01__wahidyankf-www-migration`. `test ! -e plans/in-progress/wahidyankf-www-migration` succeeds,
    `ls -d plans/done/*wahidyankf-www-migration | wc -l` prints `1`, and `ls plans/done` shows the folder named with the
    date and no suffix. Git recorded all eighteen files as `R100` renames, so the archived history reads as one move
    rather than eighteen deletions and eighteen additions.
- [x] [AI] Update `plans/in-progress/README.md` and `plans/done/README.md` and resolve every archived internal link
      directly — remove the plan from the in-progress `## Active Plans` list and its `## Directory Map`, and add a dated
      `## Directory Map` entry for the archived folder in `plans/done/README.md` — acceptance, **amended during
      execution to make the added entry observable**: `rg -n 'wahidyankf-www-migration' plans/done/README.md` finds the
      entry and its link resolves to the dated folder, `rg -n 'wahidyankf-www-migration' plans/in-progress/README.md`
      finds nothing while `rg -c 'Directory Map' plans/in-progress/README.md` still prints `1`, and
      `npm run check:markdown-links` exits 0. The link check alone was the whole criterion and could not carry it: it
      validates the links that exist rather than the entries that are missing, so `plans/done/README.md` gaining no
      entry at all exits 0 — the same hazard this plan pairs with a positive `rg` twice already. The negative half needs
      the pair for the same reason.
  - Note, 2026-09-01: `plans/in-progress/README.md` now reads `None right now.` under `## Active Plans` and
    `No plan folders right now.` under its `## Directory Map`, which is the shape `plans/backlog/README.md` already uses
    for an empty stage; `rg -n 'wahidyankf-www-migration'` finds nothing in it while `rg -c 'Directory Map'` still
    prints `1`. `plans/done/README.md` carries a dated entry beside the Badak Mini plan, and
    `npm run check:markdown-links` exits 0 — every relative link inside the plan still resolves, because the archived
    path sits at the same depth as the one it left.

### Phase 7 Gate

> Every check below passes before the plan is considered complete.

- [x] [AI] Run `npm run test:quick` — acceptance: exits 0 for the two projects that define `test:quick`, `badakmini-cli`
      and `wahidyankf-www`.
  - Note, 2026-09-01: exits 0. 258 tests green across 12 test files for `wahidyankf-www`, and `badakmini-cli` green from
    cache.
- [x] [AI] Run `npm run typecheck` and `npm run lint` — acceptance: both exit 0 across all three projects.
  - Note, 2026-09-01: both exit 0 across all three projects. Both were re-run with `--skip-nx-cache` rather than
    accepted from a 100% cache hit, because a gate that reports a cached result proves the hash matched rather than that
    the command passed — the same reasoning this plan applied to every other criterion that reads a tool's own
    reporting.
- [x] [AI] Run `npm run test:integration` — acceptance: exits 0. [AC-3]
  - Note, 2026-09-01: exits 0. 8 integration tests green across the two projects that define the target. [AC-3]
- [x] [AI] Run `npm run test:e2e` — acceptance: exits 0. [AC-5]
  - Note, 2026-09-01: exits 0 from a cold cache, 0 of 5 tasks read from it. 36 scenarios passed and 34 generated tests
    skipped, which is exactly the number `e2e-skip-baseline.json` records, so the deliberate gap is the whole gap.
    [AC-5]
- [x] [AI] Run `npm run format:check`, `npm run check:markdown-links`, `npm run check:governance`, and
      `npm run check:workflows`, plus `npm run check:harness-parity` and `npm audit --audit-level=low` — acceptance: all
      six exit 0. **The last two were added during execution.** Parity, because this phase's learnings triage edits all
      three `plan-checker` prompts and the Phase 4 and Phase 6 gates both run parity for smaller harness surfaces;
      leaving it out would make this the one phase that edits three harness files without checking they stayed equal.
      And the audit, because Phase 5 moved two dependency pins to clear five advisories and `[AC-9]` leans on the
      result. [AC-8] [AC-9]
  - Note, 2026-09-01: all six exit 0, and `npm audit --audit-level=low` reports `found 0 vulnerabilities`. [AC-8] [AC-9]
- [x] [AI] Confirm `git ls-files cv` returns nothing and `ls libs` lists only `README.md` — acceptance: both hold.
      [AC-1] [AC-6]
  - Note, 2026-09-01: `git ls-files cv` prints nothing and `test ! -e cv` succeeds, the second proving the first looked
    at a path rather than passing on an empty argument; the three absorbed documents are tracked under
    `apps/wahidyankf-www/docs/`. `ls libs` prints `README.md` and nothing else. [AC-1] [AC-6]
- [x] [AI] Run `git add -A && npm run check:rule-change` — acceptance: it exits 0, and if the learnings triage routed an
      entry into `repo-governance/` or a harness instruction it announces that path, in which case the workflow it names
      is followed rather than dismissed. Phase 7's rule-path edits are conditional on what the triage finds, unlike
      Phase 5's, but a conditional edit that reaches a rule path needs the same announcement as a certain one, and this
      is the last gate that could still make it.
  - Note, 2026-09-01: exits 0 and names no workflow. Nothing under `repo-governance/` or any harness directory is staged
    here: the triage's rule-path edits and the three the archival gate made were committed earlier today as `abcea7c`,
    `924e408`, and `d251d1a`, where the same check named Rules Propagation and Harness Alignment for them and both were
    worked. This gate stages the archival move alone, and `plans/` is not a rule path.
- [x] [AI] Commit and push the archival move with a message naming the plan — acceptance: `git status --short` is empty.
  - Note, 2026-09-01: committed and pushed to `origin main`. The gate's five cycles of fixes landed first as five
    thematic commits, `d6ef3e0` through `42bb856`; the archival move and the two index updates land here.

> **Pause Safety**: The application, its E2E harness, and its corpus are delivered and green; the CV is consolidated;
> governance explains the deployment branch; and the plan is archived with its learnings routed. `ose-public` remains
> untouched and still serves the site. The domain cutover is deliberately not done and belongs to a separate, separately
> authorized plan. Safe to stop. Resume with `npm run test:quick`.
