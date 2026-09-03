# Business Rationale

## Why This Work Exists

An audit of every `project.json` in this workspace found the files policy-compliant but not standardized. All three
satisfy the [Nx workspace policy](../../../repo-governance/development/nx-workspace-policy.md) and the
[testing policy](../../../repo-governance/development/testing-policy.md) target contract. What differs is _how_ the same
thing is expressed, and the divergence has a single cause: there is a written rule for how Nx may be used, and none for
what a `project.json` must look like. Each file therefore encodes the judgment of whoever wrote it.

`badakmini-cli` was authored first. `wahidyankf-www` and `wahidyankf-www-e2e` arrived later through
[the migration plan](../../done/2026-09-01__wahidyankf-www-migration/README.md), which worked out a deliberate Target
Contract — explicit `cache` on every target the root `targetDefaults` does not reach, `outputs` on every cached target
that writes an artifact, `{projectRoot}` as the working directory for every single-command target. That contract was
never written down anywhere a third project could read it, so it exists only as an artifact of one archived plan.

The concrete cost today is one behavioral defect and one structural oddity:

- `badakmini-cli:test:coverage:unit` resolves to `cache=undefined` while `wahidyankf-www:test:coverage:unit` resolves to
  `cache=true`. Two projects run the same-named target inside the same ordered quick gate with opposite caching. Its
  sibling targets are explicitly `cache: false` in both files, so this reads as an omission rather than a decision.
- `wahidyankf-www` is the only application in the workspace that owns no `test:e2e` target. Its browser suite lives in a
  separate project, so the application's own target set is incomplete, and `npm run test:quick` and the pre-push
  affected run reach that separate project not at all.

## Who It Affects

The repository owner, working alone, and every agent session that reads these files to decide how to add or change a
target. A third project added tomorrow has four naming styles and two working-directory strategies to choose from and no
rule to choose by. That is the failure this plan closes.

## What Success Means

Both remaining projects expose the identical ten-target contract, written in one style, and that contract is stated in
`testing-policy.md` where the next project will find it. No behavior changes. Every gate that passes today passes
afterwards, at the same or lower cost.

## Decisions Taken

Each was resolved with the owner before authoring, as the
[grilling-with-options policy](../../../repo-governance/conventions/grilling-with-options-policy.md) requires. The
chosen option and its reason are recorded here; the technical consequences are in
[tech-docs/README.md](tech-docs/README.md).

| Decision                      | Chosen                                 |
| ----------------------------- | -------------------------------------- |
| how much divergence to fix    | configuration consistency, not renames |
| how to hold the convention    | a written rule, no checker             |
| where the browser suite lives | co-located in `wahidyankf-www`         |
| where the rule lives          | `testing-policy.md`                    |

**Not renaming.** The owner first chose the full sweep including a `check:*` naming family, then narrowed it. Renames
ripple into `package.json`, `workspace-commands.md`, `docs/how-to/run-nx-workspace.md`, and three READMEs in exchange
for a readability gain, and the target list the owner supplied names only standard targets.

**No checker.** The owner first chose a Badak Mini validator, then cancelled it. A `projecttargets` validator already
existed at `0e213bc` and was deleted at `b94d85c` with no recorded reason, so reinstating it would have reversed an
undocumented decision in order to build a Go feature carrying Gherkin, three adapters, and a 99% coverage gate.

**Co-locating the suite.** This makes the application expose the same ten targets `badakmini-cli` exposes, which is the
shape the [BDD policy](../../../repo-governance/development/behavior-driven-development-policy.md) role matrix states as
the default for an application. A dedicated E2E project is permitted there, not required.

**Putting the rule in the testing policy.** It already owns a partial contract — "cacheable `typecheck`, `lint`,
`test:unit`, `test:coverage`, and `test:quick`". A second document stating a fuller version of the same contract is
exactly the contradiction `rules-checker` exists to find. At 403 words it has room for the rest.

## Non-Goals

- **No target renames.** `governance`, `markdown-links`, `capability-parity`, `rule-change`, `static-routes:validation`,
  `specs:e2e:baseline`, and `generate:cv-pdf` keep their names. The workspace keeps more than one naming style; that is
  an accepted outcome, not an oversight.
- **No Badak Mini change.** No new subcommand, no new Nx target on that project, no pre-push wiring. The rule is
  verified in review, like the [grilling](../../../repo-governance/conventions/grilling-with-options-policy.md) and
  [task tracking](../../../repo-governance/conventions/task-tracking-policy.md) policies are.
- **No change to `static-routes:validation`'s placement.** It stays in `test:quick`'s `dependsOn`.
  [The migration plan measured it](../../done/2026-09-01__wahidyankf-www-migration/evidence/phase-3-measurements.md) at
  4.2–5.1 seconds and recorded the placement as a decision it deliberately did not act on. This plan does not reverse an
  evidence-backed choice as a side effect of a consistency sweep; it writes down when `dependsOn` is the right mechanism
  instead.
- **No coverage threshold, denominator, or exclusion change**, and no change to any Gherkin scenario.

## Risks

**The quick gate widens.** Today `wahidyankf-www-e2e` owns no `test:quick`, so editing a step file runs nothing at
pre-push. After the merge those files belong to `wahidyankf-www`, so editing one marks that project affected and runs
its full quick gate, including the 4.2–5.1 second `static-routes:validation` build. This is a real cost increase and it
is accepted: the same edit today gets no type-check and no lint at push time, which is the weaker position of the two.

**Step files lose some type strictness.** `apps/wahidyankf-www-e2e/tsconfig.json` sets `noUncheckedIndexedAccess`,
`noUnusedLocals`, and `noUnusedParameters` to `true`; the application's `tsconfig.json` sets all three to `false`. After
the merge the steps type-check under the looser settings. Keeping the stricter ones would mean a second `tsconfig` and a
second `typecheck` target, which reinstates the split the merge removes. This is a judgment call, taken knowingly, and
Phase 2 records the exact diagnostics that disappear rather than assuming there are none.

**The C4 model draws a container that will not exist.** `specs/apps/wahidyankf-www/architecture.md` names
`apps/wahidyankf-www-e2e` as a container in its Container View and justifies drawing it. That is as-built truth today
and becomes wrong the moment the merge lands, so it is updated in the same phase rather than a later one.
