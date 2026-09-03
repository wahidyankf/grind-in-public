# Learnings

Written during execution, in the moment something is noticed — a surprise, a wrong assumption, a rule that failed to
prevent the failure it targets. Not reconstructed afterwards: a reconstructed entry records what the author already
believed rather than what happened.

Each entry is one short paragraph: what happened, and what a future reader should do differently. Phase 4 triages every
entry to exactly one durable home per the
[knowledge capture rules](../../../repo-governance/conventions/plans-organization-policy/knowledge-capture.md), and
archival is blocked until each has reached a terminal state.

## Entries

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
