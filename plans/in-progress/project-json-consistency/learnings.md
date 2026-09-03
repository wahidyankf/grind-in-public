# Learnings

Written during execution, in the moment something is noticed — a surprise, a wrong assumption, a rule that failed to
prevent the failure it targets. Not reconstructed afterwards: a reconstructed entry records what the author already
believed rather than what happened.

Each entry is one short paragraph: what happened, and what a future reader should do differently. Phase 4 triages every
entry to exactly one durable home per the
[knowledge capture rules](../../../repo-governance/conventions/plans-organization-policy/knowledge-capture.md), and
archival is blocked until each has reached a terminal state.

## Entries

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
