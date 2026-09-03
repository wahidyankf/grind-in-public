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

| Reader             | Location                                    | Destination                                |
| ------------------ | ------------------------------------------- | ------------------------------------------ |
| root npm script    | `package.json`, `test:scheduled`            | the same targets under `wahidyankf-www`    |
| scheduled CI       | `.github/workflows/full-bdd.yml`            | `wahidyankf-www:install`                   |
| npm workspace glob | `package.json`, `workspaces`                | unchanged; stops matching                  |
| command reference  | `workspace-commands.md`                     | two lines rewritten, one paragraph deleted |
| application index  | `apps/README.md`                            | entry deleted                              |
| application README | `apps/wahidyankf-www/README.md`             | names the co-located directory             |
| corpus index       | `specs/apps/wahidyankf-www/README.md`       | co-located adapter paths                   |
| behavior index     | `specs/.../behavior/README.md`              | repointed at the application README        |
| C4 model           | `specs/apps/wahidyankf-www/architecture.md` | one project, co-located adapter            |
| config comments    | the moved `playwright.config.ts`            | intra-project target names                 |

**The three executable readers** are the ones a stale name would actually break. `package.json`'s `test:scheduled`
invokes `wahidyankf-www-e2e:specs:e2e:baseline` and `…:test:e2e`; both move to the `wahidyankf-www:` prefix and are
proved by `npm run test:scheduled` completing with both steps running. The scheduled workflow installs browsers through
`wahidyankf-www-e2e:install`; it moves to `wahidyankf-www:install` and is proved by `npm run check:workflows` plus the
target appearing in `npx nx show project wahidyankf-www`. The `workspaces` glob is `apps/*` and needs no edit at all: it
simply stops matching a directory that no longer exists, and `npm install` regenerates `package-lock.json` without the
workspace entry.

**The five documentation readers** each name the project in prose. `workspace-commands.md` carries three narrower-run
lines, which become two under the merged project, and one paragraph explaining why `wahidyankf-www-e2e` owns no
`test:quick`; that paragraph is deleted rather than rewritten, because its subject ceases to exist. `apps/README.md`
loses its third application entry, and `apps/wahidyankf-www/README.md` gains the retired README's description of the
browser layer. All five are proved by `npm run check:markdown-links` together with a `git grep` for the retired name
returning nothing outside `plans/done/`.

**The two `specs/` readers and the C4 model** are as-built truth and are the ones that would be wrong rather than merely
stale. `specs/apps/wahidyankf-www/README.md` carries the adapter path, two verification-command rows, and the
skip-baseline sentence; `behavior/README.md` points at the retired project's README; `architecture.md` draws the project
as a container. All three are proved against `npx nx show projects` output, which must list exactly two projects, and by
the link check.

**The moved configuration's own comments** name `specs:e2e:baseline` and `test:e2e`'s `dependsOn` on
`wahidyankf-www:build`. They travel with the file and are corrected to the intra-project forms, proved by reading the
moved file for any project name that no longer exists.

Accepted shape and version are unchanged for every reader above: no target's command, no dependency pin, and no file's
content changes as part of the move except where this document names it. The owner of every row is the repository owner.

Two readers named the project only through the archived migration plan under `plans/done/`. Those are history and are
deliberately not edited, as the
[lifecycle moves rule](../../../../repo-governance/conventions/plans-organization-policy/lifecycle-moves.md) requires of
a done plan.

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
expand and contract steps land in a single commit whose gate runs the suite. There is no retry behavior to define,
because no step is partially applicable — `git mv` either moves the tree or fails, and the Nx graph either resolves the
project or reports it as unknown.

Rollback reader and writer behavior is symmetric: reverting the commit restores the deleted project's six configuration
files and the eight step files at their original paths, restores the three `wahidyankf-www-e2e` invocations in
`package.json` and the workflow, and leaves `package-lock.json` to be regenerated by `npm install`. The recovery source
is Git history, and Phase 2's gate rehearses restoration by confirming the pre-merge commit is reachable and its tree
contains the eight step files before the phase is committed.

## Malformed and Unknown Input

The one place this transition could meet input it does not understand is `bddgen`'s discovery of step files at the new
path. A step file that fails to parse, or that parses but registers no matching step, is not coerced or discarded:
`playwright-bdd` is configured with `missingSteps: "skip-scenario"`, which renders the affected scenarios as
`test.fixme` and leaves them visible in the generated output. The skip baseline then reports the changed count as a
failure. That is the opaque-record-with-a-reported-outcome behavior the rule requires, and it is carried over unchanged
rather than introduced here.
