# Learnings

Written during execution, in the moment something is noticed — a surprise, a wrong assumption, a rule that failed to
prevent the failure it targets. Not reconstructed afterwards: a reconstructed entry records what the author already
believed rather than what happened.

Each entry is one short paragraph: what happened, and what a future reader should do differently. Phase 4 triages every
entry to exactly one durable home per the
[knowledge capture rules](../../../repo-governance/conventions/plans-organization-policy/knowledge-capture.md), and
archival is blocked until each has reached a terminal state.

## Entries

**2026-09-03 — Phase 2 — the module-system delta produced a warning, not a failure.** The retired project declared
`"type": "module"` and the application declares none, with the root at `"commonjs"`, so `bddgen`'s ESM output now lands
in a package resolved as CommonJS. The plan inventoried this and required a stop-and-record if `test:e2e` failed on it.
It did not fail: the suite passes 36 and skips 34, and Node emits `MODULE_TYPELESS_PACKAGE_JSON`, reparsing as ESM with
a stated performance overhead. Left as is, because adding `"type": "module"` to a Next.js application manifest is a
change this plan does not sanction and the warning costs a parse rather than correctness. Worth revisiting only if the
E2E run's wall time becomes a complaint.

**2026-09-03 — Phase 2 — nineteen and thirty-four were both right.** `playwright.config.ts` said the four unbound
features hold nineteen scenarios and that `specs:e2e:baseline` "holds that count to nineteen", while the baseline file
recorded 34. Measured rather than guessed: the four generate 6, 19, 7, and 2 `test.fixme` entries, totalling 34, because
`playwright-bdd` emits one test per `Examples` row and three of the four are Scenario Outlines. The scenario count and
the generated-test count are different quantities and the comment conflated them. The deleted project's README had this
right and the config comment did not; the corrected sentence now names both numbers and says which is which. When two
numbers in one comment disagree, the likely fault is that they measure different things.

**2026-09-03 — Phase 1 — "redundant" was a property of the cache, not of the repository.** The plan removed
`{workspaceRoot}/apps/badakmini-cli/tests/e2e/**/*` from three targets on the reasoning that the path lies inside
`{projectRoot}` and is therefore already covered by Nx's built-in `default` input. The cache probe the plan designed for
it agreed: with the explicit input gone, changing a file under `tests/e2e/` still missed the cache, so `default` really
does reach it. The removal was still wrong. `apps/badakmini-cli/tests/bdd/adapter_parity_test.go` carries
`TestE2EBindingInputRegression`, which reads `project.json` and fails unless `test:coverage:behavior` declares that
exact string — a test whose name says it was written after this broke once before. The input is not redundant; it is a
deliberate declaration that the behavior target invalidates on E2E binding changes whatever `default` happens to cover
today. Restored in all three targets, and the plan's non-goal forbids editing Badak Mini's Go code, so the test is the
authority rather than a thing to adjust.

The lesson generalizes past this input: a probe can only show what the system does now, and "covered by a broader rule"
is not the same claim as "safe to delete". Before removing a declaration because something else subsumes it, search for
a test that names it. A behavioural probe and a grep for the literal string answer different questions, and this plan
only asked the first.

**2026-09-03 — Phase 0 — a `tail` window sized by guess truncated the evidence it was capturing.** Two Phase 0 items
piped a coverage run through `tail -20` and `tail -5` to record a baseline figure. Both windows landed past the line
they were meant to capture: Nx appends a four-to-six line run summary after the command's own output, so the `All files`
row and the `unit statement coverage:` line were already scrolled out. Both commands exited 0 and both files were
written, so the failure was silent — the acceptance criterion is what caught it, because it named the content the file
had to contain rather than merely that the file existed. Widened to `tail -60` and `tail -15`. A future reader writing a
capture step should either grep for the wanted line or capture the whole stream; a fixed line count is a guess about
output length that nothing verifies.

<!--
One dated paragraph per entry. Six checklist items write here directly. Three
always write: Phase 2's pre-merge typecheck result under the stricter compiler
settings, Phase 2's measured scenario counts from bddgen, and Phase 3's gate
review of the written rule against both project.json files. Three write only if
triggered: Phase 1's conditional input removal, Phase 1's bare-nx-run grep
control, and Phase 2's module-resolution branch, which records the exact error
when the first wahidyankf-www:test:e2e run fails that way. Phase 4 gives each of
those three a dated disposition, Not triggered included. The first entry below
was written during planning and is a stated assumption rather than an
observation.
-->

**2026-09-03 — stated assumption, recorded before execution.** "Writes an artifact", in the `outputs` rule this plan
writes into `testing-policy.md`, means producing something a later target or a person consumes. A compiler's own
incremental state is not one. `wahidyankf-www:typecheck` is the case that forces the question: it resolves to
`cache: true` through the root `targetDefaults`, its `tsconfig.json` sets `"incremental": true`, and
`apps/wahidyankf-www/tsconfig.tsbuildinfo` exists on disk — yet nothing reads that file but the `tsc` invocation that
wrote it, and it is regenerated on demand, so the target declares no `outputs` and is not a carve-out from the rule that
binds every target. `badakmini-cli:typecheck` is the same shape: `go vet` writes only into the Go build cache, outside
the workspace. This is a definition the plan asserts, not a measurement it took, which is why it is recorded here rather
than left implicit in the rule. Phase 4 routes it: if it survives the Phase 3 gate review it belongs beside the rule in
`testing-policy.md`, and if it does not, both the rule and the Phase 1 artifact map need the other answer.

## Phase 3 Rule Review — [AC-7]

One line per rule the delivered documents state, the target each was checked against in both `project.json` files, and
the verdict. Re-read in Phase 4 against the files as they stand, and every verdict below is that second reading.

| Rule                                                             | `badakmini-cli`                                            | `wahidyankf-www`                                                 | Verdict |
| ---------------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------------- | ------- |
| Every project exposes the same ten targets                       | all ten present, 15 targets total                          | all ten present, 19 targets total                                | Holds   |
| The three eligibility-dependent targets are present or justified | owns a real local boundary and a CLI process, so all three | owns both, so all three                                          | Holds   |
| Every target declares `cache` explicitly                         | `all 15 targets declare cache`                             | `all 19 targets declare cache`                                   | Holds   |
| An uncached target declares no `outputs`                         | `outputs rule holds for all 15`                            | `outputs rule holds for all 19`                                  | Holds   |
| A cached artifact-writer names its path                          | `build`, `test:coverage:unit`                              | `build`, `test:coverage:unit`, `specs:e2e:baseline`              | Holds   |
| A compiler's own incremental state is not an artifact            | `typecheck` runs `go vet`, no `outputs`                    | `typecheck` writes `tsconfig.tsbuildinfo`, no `outputs`          | Holds   |
| No command encodes its own project path                          | no `apps/badakmini-cli` in a command                       | no `apps/wahidyankf-www` in a command                            | Holds   |
| A single-command target declares `options.cwd`                   | all 13 command targets                                     | all command targets                                              | Holds   |
| A shared input glob is named once per project                    | `behaviorCorpus`, raw glob count 1                         | `behaviorCorpus` and `workspaceScripts`, raw glob count 1        | Holds   |
| `options.commands` states the gate, `dependsOn` a prerequisite   | `test:quick` orders in `commands`                          | `test:e2e` depends on `build`; `test:quick` orders in `commands` | Holds   |

No rule is recorded as contradicted. This list was written into the Execution Record during the Phase 3 gate rather than
here, which the gate item's own acceptance names as the destination; it is placed here in Phase 4 so `[AC-7]`'s proof
sits where the criterion says to look for it.

## Phase 4 Dispositions

Every entry above reaches exactly one terminal state, per the
[knowledge capture rules](../../../repo-governance/conventions/plans-organization-policy/knowledge-capture.md). No entry
names a credential, a private identifier, or a runtime payload; each was re-read for that before routing.

- **Module-system delta** → routed to `repo-governance/development/testing-policy/tooling.md`, Recorded Deviations,
  beside the `wahidyankf-www` compiler-settings deviation it continues. That section already exists to hold "this is the
  state, and here is why it is not fixed", which is exactly the shape of this entry.
- **Nineteen and thirty-four** → routed to code, the strongest form: the corrected `missingSteps` comment in
  `apps/wahidyankf-www/playwright.config.ts` and the skip-baseline record in `apps/wahidyankf-www/README.md` now name
  both numbers and say which is which. The general half — that two disagreeing numbers in one comment usually measure
  different things — is discarded: it is an observation about reading, not a rule a repository can hold anyone to.
- **"Redundant" was a property of the cache** → routed to `repo-governance/development/testing-policy/target-shape.md`,
  Shared Inputs, as a rule: search for a test that names a declaration before deleting it, because a behavioural probe
  and a grep for the literal string answer different questions. `TestE2EBindingInputRegression` already exists and needs
  no companion, so the executable capture is in place and the rule is what was missing.
- **A `tail` window sized by guess** → discarded. The generalizable half is already a standing rule: an acceptance
  criterion names the content a step must produce rather than that it ran, and that rule is precisely what caught this.
  Restating it would duplicate a rule rather than add one, and a duplicated rule is what later disagrees with its
  source.
- **Stated assumption, recorded before execution** → routed to
  `repo-governance/development/testing-policy/target-shape.md`, Outputs, where "artifact" is defined as something a
  later target or a person consumes and `tsconfig.tsbuildinfo` is named as the case that is not one. It survived the
  Phase 3 gate review, which is the condition the entry itself set for routing it there.

### Conditional Items

Three checklist items write here only if triggered. Each gets a dated disposition whether or not it fired, so no named
writer is left without a terminal state.

- **2026-09-03 — Phase 1, the `tests/e2e` input removal — Not triggered as written, but the removal was reverted for a
  different reason.** The item's trigger is the cache probe over `apps/badakmini-cli/tests/e2e/README.md` reporting a
  hit after the content change. It reported a miss: with the explicit input gone, `default` still invalidated the
  target, so the conditional's own criterion never fired. The input was restored anyway, because
  `TestE2EBindingInputRegression` fails without the literal string — a reason the probe was not built to detect. Both
  facts are recorded because either alone misreads the phase: the probe was right about the cache and wrong about
  whether the removal was safe.
- **2026-09-03 — Phase 1, the `[AC-5]` pre-edit control — Not triggered.** Run against the unedited
  `apps/wahidyankf-www/project.json`, the bare-`nx run` grep printed exactly the one expected line, the
  `static-routes:validation` command reading `"command": "nx run wahidyankf-www:build --skip-nx-cache && node ..."`. The
  gate pattern was therefore proved able to see the defect it exists to catch, and nothing was written. Re-run in Phase
  4 against the delivered files, the same pattern prints nothing.
- **2026-09-03 — Phase 2, the module-resolution branch — Not triggered.** The first `npx nx run wahidyankf-www:test:e2e`
  after the merge exited 0, passing 36 and skipping 34. The plan required a stop and a return to the owner only on a
  module-resolution _failure_; what appeared was a `MODULE_TYPELESS_PACKAGE_JSON` warning, which is recorded as its own
  entry above and routed to `tooling.md`. A warning is not the branch's trigger, and treating it as one would have
  stopped the phase on a passing suite.
