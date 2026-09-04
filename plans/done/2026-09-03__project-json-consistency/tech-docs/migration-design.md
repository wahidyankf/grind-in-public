# Migration Design

This plan retires a configuration representation: the Nx project `wahidyankf-www-e2e`. The
[plan migrations rule](../../../../repo-governance/conventions/plans-organization-policy/plan-migrations.md) governs it,
and this document holds the inventory, the transition, and the reason the contraction step carries a zero-length
compatibility window.

No secret, credential, or private value appears in the inventory below; the only environment variables named are
`APP_ENV`, `BASE_URL`, and `BADAKMINI_BIN`, which are named as variables and never as values.

## Inventory of Every Reader

Each row is a place that names `wahidyankf-www-e2e`, found by searching the repository for the project name and for each
of its target names. Generated caches under `.nx/` are excluded: they are gitignored, regenerate from the project graph,
and hold no authored content.

| Reader             | Location                                    | Destination                             |
| ------------------ | ------------------------------------------- | --------------------------------------- |
| root npm script    | `package.json`, `test:scheduled`            | the same targets under `wahidyankf-www` |
| scheduled CI       | `.github/workflows/full-bdd.yml`            | `wahidyankf-www:install`                |
| npm workspace glob | `package.json`, `workspaces`                | unchanged; stops matching               |
| command reference  | `workspace-commands.md`                     | three lines and one sentence rewritten  |
| tooling deviation  | `testing-policy/tooling.md`                 | deviation entry removed                 |
| application index  | `apps/README.md`                            | entry deleted                           |
| application README | `apps/wahidyankf-www/README.md`             | names the co-located directory          |
| apps corpus index  | `specs/apps/README.md`                      | corpus owned by one project             |
| corpus index       | `specs/apps/wahidyankf-www/README.md`       | co-located adapter paths                |
| behaviour index    | `specs/.../behaviours/README.md`            | repointed at the application README     |
| C4 model           | `specs/apps/wahidyankf-www/architecture.md` | one project, co-located adapter         |
| config comment     | the moved `playwright.config.ts`            | intra-project `dependsOn` name          |
| step comment       | `tests/bdd/accessibility.steps.ts`          | co-located step-file path               |

**The three executable readers** are the ones a stale name would actually break. `package.json`'s `test:scheduled`
invokes `wahidyankf-www-e2e:specs:e2e:baseline` and `…:test:e2e`; both move to the `wahidyankf-www:` prefix, are proved
in Phase 2 by a `grep` naming each new invocation string exactly, and are proved end to end by the Phase 4 gate's
`npm run test:scheduled` completing with both steps running. The scheduled workflow installs browsers through
`wahidyankf-www-e2e:install`; it moves to `wahidyankf-www:install` and is proved by a `grep` for that exact string plus
the target appearing in `npx nx show project wahidyankf-www --json`. `npm run check:workflows` runs alongside and must
pass, but proves nothing about either name: it is `actionlint`, which validates YAML and embedded shell and does not
resolve Nx targets, so it passes over the retired name and a typo alike. The `workspaces` glob is `apps/*` and needs no
edit at all: it simply stops matching a directory that no longer exists, and `npm install` regenerates
`package-lock.json` without the workspace entry.

**The documentation readers** each name the project in prose. `workspace-commands.md` carries three narrower-run lines,
which keep their three targets under the merged project's prefix, and one three-sentence paragraph below them. Its first
two sentences explain why `wahidyankf-www-e2e` owns no `test:quick` and are deleted, because their subject ceases to
exist. Its third sentence is not: it tells a reader to install the browser once per machine before `test:e2e`, which
stays true under the merged target name, and this document is the only place in the repository that states it once
`apps/wahidyankf-www-e2e/README.md` is deleted. That sentence is rewritten as `wahidyankf-www:install` rather than
dropped. `apps/README.md` loses its third application entry. `apps/wahidyankf-www/README.md` gains the retired README's
record of the browser layer, and three of its clauses are load-bearing rather than descriptive because that README is
their only home in the repository: the list of the four feature files the Playwright adapter deliberately does not bind,
the recorded skip baseline of 34 stated as generated tests rather than scenarios, and the standing rule "Raise the
number only when a scenario is deliberately left unbound, and say here why." `specs/.../behaviours/README.md` is
repointed at that same README and asserts of it that it "names them and records the generated skip baseline", so the
repoint's acceptance reads the target for those clauses. It is not the link check: that reference is inline code and not
a Markdown link, so `check:markdown-links` never resolves it and a pointer at a document holding none of the content
would pass. `repo-governance/development/testing-policy/tooling.md` carries a Recorded Deviations entry stating that
`apps/wahidyankf-www-e2e` sets `module` and `moduleResolution` as the application does with `target` at `ES2022`, and a
sentence explaining that the E2E project follows the runner it hosts. That entry is removed in Phase 2, not in Phase 3
with the other governance edits, because Phase 2 deletes `apps/wahidyankf-www-e2e/tsconfig.json` and the deviation's
subject stops existing at the merge. The `apps/wahidyankf-www` half of the same paragraph is kept unchanged: that
project and its three overrides remain. Most of them are proved by `npm run check:markdown-links` together with a
`git grep` for the retired name returning nothing outside `plans/`. Two are not, and each names its own proof instead:
the `workspace-commands.md` once-per-machine sentence is proved by a `grep` for that phrase still counting `1`, and the
`apps/wahidyankf-www/README.md` fold-in by the greps the behaviour-index repoint runs against it.

**The `specs/` readers and the C4 model** are as-built truth and are the ones that would be wrong rather than merely
stale. `specs/apps/wahidyankf-www/README.md` carries the adapter path, two verification-command rows, and the
skip-baseline sentence; `behaviours/README.md` points at the retired project's README; `architecture.md` draws the
project as a container. `specs/apps/README.md` describes the `wahidyankf-www` behaviour corpus as "shared with the
dedicated E2E project `apps/wahidyankf-www-e2e`", a sentence that becomes false at the merge: after it the corpus is
owned and bound by one project. All of them are proved against `npx nx show projects` output, which must list exactly
two projects, and by the link check — except `behaviours/README.md`, whose sentence is an inline-code reference the link
check cannot see and whose proof is therefore a read of the document it now points at.

**The moved configuration's own comment** describes `test:e2e` as declaring `dependsOn` on `wahidyankf-www:build`. It
travels with the file and is corrected to the intra-project `build`, proved by reading that comment line against the
`dependsOn` the merged target declares. It is not proved by looking for a project name that no longer exists:
`wahidyankf-www` survives the merge, so that question is answered `no` before the correction as well as after. The
comment naming `specs:e2e:baseline` in the same file already uses the bare target name, carries no cross-project prefix,
and needs no edit.

**One comment outside the moved configuration** names the retired path too.
`apps/wahidyankf-www/tests/bdd/accessibility.steps.ts` explains that the full axe-core scan runs at the e2e tier and
cites `apps/wahidyankf-www-e2e/steps/accessibility.steps.ts` as its home; the citation is repointed to
`apps/wahidyankf-www/tests/e2e/steps/accessibility.steps.ts`. This is the only `.ts` file Phase 2 edits that is not one
of the eight moved step files, and the edit is confined to the comment: no assertion, no `@covers` tag, and no step
binding changes, so `npm run test:quick` proves the file still behaves as it did.

**One shape delta is not a reader, and it is inventoried here rather than asserted away.**
`apps/wahidyankf-www-e2e/package.json` declares `"type": "module"`. `apps/wahidyankf-www/package.json` declares no
`type` at all, and the root `package.json` declares `"type": "commonjs"`. Node resolves a `.js` file's module system
from the nearest enclosing `package.json`, which after the merge is the application's, so the absent `type` there means
the default, CommonJS, rather than the ESM the deleted manifest declared. What lands in that package is ESM: `bddgen`
writes `apps/wahidyankf-www/.features-gen/*.feature.spec.js`, and each generated file opens with
`import { test } from "playwright-bdd";`. Whether that combination fails is not claimed here in either direction. Phase
2's first `wahidyankf-www:test:e2e` run — the guard-is-green-at-rest item in the checklist body, before the gate — is
what determines it, and that item carries the note saying what happens if it fails this way. No item of this plan adds
`"type": "module"` to a Next.js application manifest, and no decision in [brd.md](../brd.md) sanctions one.

Accepted shape and version are otherwise unchanged for every reader above: no target's command, no dependency pin, and
no file's content changes as part of the move except where this document names it. The module-system declaration is the
one exception, it belongs to the deleted manifest rather than to any row in the table, and it is recorded above rather
than claimed away. The owner of every row is the repository owner.

Readers under `plans/` are excluded from the inventory and from the completeness grep in Phase 2's gate. The archived
migration plan under `plans/done/` names the project as history and is deliberately not edited, as the
[lifecycle moves rule](../../../../repo-governance/conventions/plans-organization-policy/lifecycle-moves.md) requires of
a done plan; this plan's own documents name it because they describe the change that retires it. Neither depends on the
name resolving to a project, so neither is a reader a stale name could break.

## Accepted Shape and Ownership

Every reader above is inside this repository and is owned by the repository owner. There is no published package, no
external consumer, no deployed artifact, and no persisted data keyed on the project name. The Nx project graph is the
only derived representation, it is regenerated on every invocation, and its cache lives in the gitignored `.nx/`
directory.

## The Transition

**Expand.** `apps/wahidyankf-www/project.json` gains `install`, `test:e2e`, and `specs:e2e:baseline`, and
`apps/wahidyankf-www/package.json` gains the three Playwright dependencies, before anything is deleted. At this point
both projects can express the browser suite.

**Migrate.** The eight step files, the skip baseline, and the Playwright configuration move by `git mv`, which preserves
history and makes the move reviewable as a rename rather than as a delete-plus-add. The identity being preserved is the
binding between a Gherkin scenario and the step that implements it; it is verified by the generated-scenario count
rather than by file presence, because a file that moved but stopped binding would still be present.

**Verify.** Verification uses the normal product flow: `npx nx run wahidyankf-www:test:e2e` builds the application,
starts `next start`, drives Chromium, and passes — the same path the scheduled workflow takes. Alongside it,
`specs:e2e:baseline` regenerates and asserts the skipped-scenario count is still exactly 34. That number is what proves
the migration preserved bindings: if a step file failed to move or failed to be discovered at its new path, `bddgen`
renders its scenarios as `test.fixme` and the count rises above 34, so the baseline fails where a passing suite would
not.

**Contract.** The deletion happens in the same commit as the move, with no compatibility window. This departs from the
rule's default of retaining compatibility for a stated window and scheduling deletion in a separately authorized later
plan, and the departure is a judgment call recorded as one. Its grounds: every reader is inventoried above and inside
this repository, none is out of band, the representation carries no data whose loss would be irreversible, and the whole
transition reverses with `git revert` of one commit. A retention window would leave two projects binding one corpus and
running two browser suites over the same scenarios, which costs real time on the scheduled run and creates a second
place for the skip baseline to drift — a worse position than the one it would be protecting against.

## Mixed-Version Boundaries, Retry, and Rollback

There is no mixed-version boundary: at no committed state do both projects exist with the merged one active, because the
expand and contract steps land in a single commit whose gate runs the suite. There is no retry behaviour to define,
because no step is partially applicable — `git mv` either moves the tree or fails, and the Nx graph either resolves the
project or reports it as unknown.

Rollback reader and writer behaviour is symmetric: reverting the commit restores the deleted project's six configuration
files and the eight step files at their original paths, restores the three `wahidyankf-www-e2e` invocations in
`package.json` and the workflow, and leaves `package-lock.json` to be regenerated by `npm install`. The recovery source
is Git history, and Phase 2's gate rehearses restoration by confirming the pre-merge commit is reachable and its tree
contains the eight step files before the phase is committed.

## Malformed and Unknown Input

The one place this transition could meet input it does not understand is `bddgen`'s discovery of step files at the new
path. A step file that fails to parse, or that parses but registers no matching step, is not coerced or discarded:
`playwright-bdd` is configured with `missingSteps: "skip-scenario"`, which renders the affected scenarios as
`test.fixme` and leaves them visible in the generated output. The skip baseline then reports the changed count as a
failure. That is the opaque-record-with-a-reported-outcome behaviour the rule requires, and it is carried over unchanged
rather than introduced here.
