# Technical Design

The entry point for this plan's technical set. Read this first, then [file-impact.md](file-impact.md) for every path,
[migration-design.md](migration-design.md) for the project-deletion transition, and
[specification-changes.md](specification-changes.md) for the C4 and plan-only-criteria decisions.

## Directory Map

- [file-impact.md](file-impact.md) — every expected path with an `[E]`, `[N]`, `[M]`, or `[D]` label.
- [migration-design.md](migration-design.md) — the inventory of every reader of `wahidyankf-www-e2e` and the disposition
  of each.
- [specification-changes.md](specification-changes.md) — the C4 container change, the unchanged Gherkin corpus, and the
  plan-only labelling of all seven acceptance criteria.

## What Was Inspected

| Read                                    | Showed                                     |
| --------------------------------------- | ------------------------------------------ |
| the three `project.json` files          | cache, inputs, and outputs per target      |
| `nx.json`                               | six `targetDefaults`, no `namedInputs`     |
| the archived migration plan             | a target contract never promoted out of it |
| its Phase 3 evidence                    | a measured, deliberate gate placement      |
| `git log` on `internal/projecttargets/` | a deleted validator, no recorded reason    |
| the E2E step file imports               | three external packages, no path alias     |
| `apps/wahidyankf-www/vitest.config.ts`  | three projects, none reaching `tests/e2e/` |
| the root `.gitignore`                   | the three generated paths already ignored  |

**The `project.json` files** were read directly and through `npx nx show project <name> --json`, which resolves what the
root `targetDefaults` contributes to each target. `badakmini-cli:test:coverage:unit` resolves to `cache=undefined` while
its `wahidyankf-www` counterpart resolves to `cache=true`: one ordered gate, two caching behaviours.

**`nx.json`** carries six `targetDefaults` entries and `analytics: false`, and declares no `namedInputs` at all. Nx
supplies a built-in `default`, which is why every `"default"` reference across the three projects resolves despite the
omission.

**[The archived migration plan](../../../done/2026-09-01__wahidyankf-www-migration/tech-docs/README.md)** worked out a
Target Contract for the two newer projects, including the reasoning behind `cwd`, `outputs`, and the `test:quick` shape.
It was never promoted out of the plan, which is why a third project would have nothing to read.

**[Its Phase 3 evidence](../../../done/2026-09-01__wahidyankf-www-migration/evidence/phase-3-measurements.md)** measured
`static-routes:validation` at 4.2–5.1 seconds and recorded the `dependsOn` placement as a decision that plan
deliberately did not act on.

**`git log` on `apps/badakmini-cli/internal/projecttargets/`** shows a `project.json` validator existed at `0e213bc` and
was deleted at `b94d85c`, whose one-line commit message records no reason.

**The E2E step files** import only `@playwright/test`, `playwright-bdd`, and `@axe-core/playwright`. No `@/*` alias
appears in any of the eight, so the merge does not have to reconcile two `paths` mappings.

**`apps/wahidyankf-www/vitest.config.ts`** defines three projects, including `src/**/*.unit.test.*`, `tests/bdd/**`, and
`tests/integration/**`. None of them would collect a file under `tests/e2e/`, and the coverage denominator is `src/**`.

**The root `.gitignore`** already ignores `.features-gen/`, `playwright-report/`, and `test-results/`, so the merged
project needs no new ignore rule of its own.

## Architecture After the Merge

```text
  before                                    after

  apps/badakmini-cli        (10 targets)    apps/badakmini-cli        (10 targets)
  apps/wahidyankf-www       ( 9 targets)    apps/wahidyankf-www       (10 targets)
  apps/wahidyankf-www-e2e   ( 7 targets)
          |                                         |
          | implicitDependencies                    | tests/e2e/ co-located,
          | dependsOn wahidyankf-www:build          | dependsOn build
          v                                         v
  two projects, one corpus                  one project, one corpus
  e2e outside every quick gate              e2e inside the affected calculation
```

The application's directory gains one subtree and keeps everything else:

```text
apps/wahidyankf-www/
+-- playwright.config.ts        moved from the deleted project
+-- tests/
|   +-- bdd/                    unchanged, vitest-cucumber
|   +-- integration/            unchanged, vitest node
|   +-- e2e/                    new home
|       +-- steps/              eight moved step files
|       +-- e2e-skip-baseline.json
+-- src/                        unchanged
```

This mirrors `badakmini-cli`, whose `tests/e2e` is a distinct Go package inside one Nx project. Co-locating the project
is not the same as sharing a runner: the merged project runs Vitest for three layers and the Playwright runner for the
fourth, exactly as `badakmini-cli` runs `go test` over four separate packages.

## Design Decisions

**Per-project `namedInputs`, not workspace-level.** Both projects need "the behavior corpus that belongs to me", and
those are different paths. A workspace-level entry would need two differently-named keys and each target would still
pick the right one by hand. A project-level `namedInputs` lets both projects declare the same name, `behaviorCorpus`, so
every affected target in both files reads `"inputs": ["default", "behaviorCorpus"]`. `wahidyankf-www` declares a second,
`workspaceScripts`, for `{workspaceRoot}/scripts/next-with-port.mjs`, because that path is a shared script rather than a
specification and folding it into `behaviorCorpus` would make the name lie.

**`{projectRoot}` and `cwd` for `badakmini-cli`, with three commands rewritten rather than trimmed.** Dropping
`-C apps/badakmini-cli` in favour of `options.cwd` is mechanical for ten of the thirteen commands. Three are not, and
each is a place a careless conversion silently breaks:

| Target                      | What the conversion must also change |
| --------------------------- | ------------------------------------ |
| `test:coverage:unit`        | the `mkdir` path                     |
| `test:coverage:integration` | the `mkdir` path                     |
| `test:e2e`                  | the `BADAKMINI_BIN` assignment       |

**Both coverage targets** begin `mkdir -p local-tmp && go -C apps/badakmini-cli test …`. Once `cwd` is declared, that
`mkdir` runs from the project directory, so an unrewritten one creates `apps/badakmini-cli/local-tmp` while the
`-coverprofile=../../local-tmp/…` argument still points at the workspace-root directory that was never created. Both
become `mkdir -p ../../local-tmp`.

**`test:e2e`** sets `BADAKMINI_BIN="$PWD/apps/badakmini-cli/dist/badak-mini"`. `$PWD` becomes the project directory once
`cwd` is declared, so leaving the literal path in place resolves to `apps/badakmini-cli/apps/badakmini-cli/dist/…`. It
becomes `BADAKMINI_BIN="$PWD/dist/badak-mini"`.

Each failure is silent in a different way: the coverage ones write a profile nobody reads and report success, and the
E2E one fails to find a binary that was built correctly.

The four `harness … validate` targets need no path handling at all: Badak Mini locates the repository with
`git rev-parse --show-toplevel`, so it behaves identically from either directory.

**`static-routes:validation` keeps `dependsOn`.** The
[migration plan's evidence](../../../done/2026-09-01__wahidyankf-www-migration/evidence/phase-3-measurements.md)
measured it and declined to move it. This plan does not overturn that as a side effect of a consistency sweep. What it
does instead is remove the ambiguity that made it look like an oversight, by writing the distinction into
`testing-policy.md`: `dependsOn` expresses a prerequisite that must precede the whole gate, and `options.commands`
expresses the ordered gate itself. Its bare `nx run` still becomes `npm exec nx -- run`, which is a separate finding and
a genuine defect — it is the one target in the workspace resolving Nx from the ambient `PATH`. Its `cwd` also moves from
`{workspaceRoot}` to `{projectRoot}`, so `node apps/wahidyankf-www/scripts/validate-static-routes.mjs` becomes
`node scripts/validate-static-routes.mjs` and the target stops naming its own project path. That is the same shape rule
applied to `badakmini-cli` above, and applying it here is what lets Phase 3 state the rule without a carve-out.

**The shape rule binds every target, so two more `wahidyankf-www` targets change.** `generate:cv-pdf` declares `outputs`
while `cache: false`, which is the identical inert declaration Phase 1 strips from `test:coverage:integration`; leaving
one and removing the other would make the rule false of the file it was derived from, and Phase 3's gate re-reads both
`project.json` files against the written rule. `static-routes:validation` is the other, above. The rule reaches in the
opposite direction too: `specs:e2e:baseline` is cached and runs `bddgen`, which writes `.features-gen/`, so it gains an
`outputs` declaration it does not carry today. `test:e2e` regenerates that directory anyway, which is why its absence
has never surfaced as a failure — but a cached target that restores nothing is the same defect whether or not something
downstream happens to paper over it.

**`badakmini-cli:test:coverage:unit` becomes cached and declares its output.** An undeclared `cache` means uncached, so
this is a real change to when the Go coverage gate runs rather than a notation fix. It writes
`local-tmp/badakmini-unit.out`, which sits outside `{projectRoot}`, so `outputs` names
`{workspaceRoot}/local-tmp/badakmini-unit.out`. Without that, a cache hit would replay the printed coverage line and
restore no profile — a target that reports success while producing nothing, which is the same silent-success class as
the `mkdir` rewrite above.

**The `test:e2e` skip guard is scoped rather than copied verbatim.** Its command opens with
`grep -rn --include='*.ts' … '\$?test\.skip\([^,)]*\)' .`, and `.` is the working directory. In the deleted project that
is one `steps/` directory. In the merged project it is the whole application: forty-three TypeScript files plus
`.next/`, which the command's `--exclude-dir` list does not name. The final `.` becomes `tests/e2e`, restoring the
pre-merge reach. This is the fourth command in this plan whose meaning changes with its working directory, and the only
one where a verbatim copy would have passed its gate today while quietly scanning a tree it was never written for.

**The looser `tsconfig` is accepted, and measured rather than assumed.** The deleted project set
`noUncheckedIndexedAccess`, `noUnusedLocals`, and `noUnusedParameters` to `true`; the application sets all three to
`false`. Preserving them would require a second `tsconfig.json` and a second `typecheck` target scoped to `tests/e2e/`,
which reinstates exactly the split this merge removes. Phase 2 therefore runs `tsc --noEmit` against the moved steps
under the old settings first and records what it reports, so the plan states the size of what is given up instead of
asserting it is nothing.

**`.features-gen/` must be excluded in one new place.** Generation moves to `apps/wahidyankf-www/.features-gen/`, inside
a project whose `tsconfig.json` includes `**/*.ts` and excludes only `node_modules`. Without a new exclude, `typecheck`
would compile generated test files. Vitest, Biome, and the coverage denominator need no change: no Vitest project's
`include` reaches `tests/e2e/`, Biome reads the root `.gitignore` through `vcs.useIgnoreFile`, and the denominator is
`src/**`.

## Dependencies

`@playwright/test@1.62.1`, `playwright-bdd@9.2.0`, and `@axe-core/playwright@4.10.1` move from the deleted project's
`devDependencies` into `apps/wahidyankf-www/package.json` at the same exact pins. No version changes, no additions, no
removals. The root `workspaces` glob is `apps/*`, so deleting the directory is the only change the root manifest needs
on that axis.

## Risks and What Absorbs Them

| Risk                                                 | Absorbed by  |
| ---------------------------------------------------- | ------------ |
| a rewritten command writes to the wrong path         | Phase 1 gate |
| `typecheck` compiles generated `.features-gen` files | Phase 2 gate |
| the skip baseline moves silently                     | Phase 2 gate |
| a `namedInputs` refactor changes a cache key         | Phase 1 gate |
| `nx affected` stops selecting the merged project     | Phase 2 gate |

**A wrong path in a rewritten command** would let a coverage gate measure nothing and still exit 0. Phase 1's gate
therefore runs `test:coverage` and `test:e2e` in full and asserts the coverage percentages still print and match the
baseline figures, rather than asserting the command exited 0.

**A `typecheck` reaching `.features-gen`** is only visible once that directory exists. Phase 2's gate runs `typecheck`
after `test:e2e` has populated it, so the new `tsconfig.json` exclude is proved against a populated tree rather than an
absent one.

**A silently moved skip baseline** would hide a step file that stopped binding at its new path. `specs:e2e:baseline` is
carried over unchanged and its gate asserts the recorded count of 34 specifically, not merely that the target passed.

**A changed cache key** could let Nx replay a stale result as a success. Phase 1 captures each target's resolved
`inputs` before and after the refactor and requires them to match.

**A graph change breaking affected selection** would quietly remove the merged project from pre-push. Phase 2's gate
runs `npx nx affected -t test:quick --base=origin/main --head=HEAD` and confirms `wahidyankf-www` is selected.

## Rollback

Every phase is one commit to `main`, and no phase writes outside the repository. Reverting a phase is `git revert` of
that commit, with no data, cache, or external state to reconcile — the Nx local cache re-keys itself from the reverted
configuration on the next run. The merge phase deletes a directory whose entire contents are recoverable from Git
history, which is why [migration-design.md](migration-design.md) records a zero-length compatibility window rather than
a staged contraction.

## Reading Order

1. [file-impact.md](file-impact.md) — what changes, exactly.
2. [migration-design.md](migration-design.md) — why the deletion is safe in one commit.
3. [specification-changes.md](specification-changes.md) — what happens to the C4 model and why no scenario changes.
4. [../delivery.md](../delivery.md) — the phased checklist.
