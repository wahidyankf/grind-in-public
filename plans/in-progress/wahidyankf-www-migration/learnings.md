# Learnings

Written during execution, in the moment something is noticed — a surprise, a wrong assumption, a rule that failed to prevent the failure it targets. Not reconstructed afterwards. Each entry is one short paragraph: what happened, and what a future reader should do differently.

Phase 7 triages every entry here to exactly one durable home before this plan may be archived.

## Entries

2026-09-01, Phase 0: a repository-wide rename in the plan reached a path it should not have. The owner renamed the destination E2E project from `wahidyankf-www-fe-e2e` to `wahidyankf-www-e2e`, and the sweep that applied it also rewrote the source-side pathspec in the Phase 0 cleanliness check, which addresses `ose-public` where the old name is still correct. The check then passed against a path that does not exist. Nothing caught it, in six strict `plan-checker` cycles or in execution, because `git status --porcelain` prints nothing for a pathspec matching no tracked file and exits 0 — an emptiness assertion is satisfied most easily by describing nothing at all. Two things to carry forward: a rename that spans two repositories is not a single sweep, because the same name is right on one side and wrong on the other; and an acceptance criterion of the form "the output is empty" needs a companion that proves the command looked at something, which here is `git ls-files --error-unmatch` over the same path list.
