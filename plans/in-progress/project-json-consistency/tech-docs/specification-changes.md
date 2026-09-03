# Specification Changes

This document records what this plan proposes to change in `specs/`. `specs/` stays the canonical as-built truth until
Phase 2's gate passes; nothing here is written into it ahead of the implementation it describes.

## Durable Contracts Versus Plan-Only Outcomes

Every acceptance criterion in [prd.md](../prd.md) is a **plan-only operational outcome**. None becomes a Gherkin
scenario in `specs/`, and none is bound to a test.

**The reason.** All seven criteria describe build configuration, project layout, or the content of a governance
document. `AC-1` through `AC-5` assert properties of `project.json` and `nx.json`; `AC-6` asserts which Nx project owns
a target; `AC-7` asserts that a policy states a rule. None of them is behavior of the software this repository ships.
The [specs policy](../../../../repo-governance/development/specs-policy.md) reserves `specs/` for what an application
should do and how it is built as a system, and a scenario there binds to a test that fails when the behavior breaks.
There is no application behavior here to break: after this plan, `wahidyankf-www` renders exactly what it renders today
and `badakmini-cli` reports exactly what it reports today.

**Delivery proof for each.** The policy requires a plan-only outcome to name its proof rather than merely assert it is
unprovable in `specs/`:

| Criterion | Proved at    | By                                            |
| --------- | ------------ | --------------------------------------------- |
| `[AC-1]`  | Phase 2 gate | `npx nx show project` for both projects       |
| `[AC-2]`  | Phase 1 gate | the resolved-cache inspection                 |
| `[AC-3]`  | Phase 1 gate | a `grep` count and a resolved-inputs diff     |
| `[AC-4]`  | Phase 1 gate | a `grep` plus `test:quick` and `test:e2e`     |
| `[AC-5]`  | Phase 1 gate | a `grep` for a bare `nx run`                  |
| `[AC-6]`  | Phase 2 gate | `nx show projects`, `test:e2e`, the baseline  |
| `[AC-7]`  | Phase 3 gate | `check:governance` and `check:markdown-links` |

`[AC-1]` requires both project inspections to list all ten contract targets. `[AC-2]` requires the inspection to report
no target whose `cache` resolves to undefined, and to name a non-zero target count so the check is proved to have read a
populated set. `[AC-6]` requires three things together: exactly two discovered projects, a passing browser suite, and a
skipped-scenario count still equal to 34.

`AC-3`'s proof is deliberately two commands rather than one. A `grep` returning zero is a criterion satisfied by finding
nothing, and on its own it is equally satisfied by a file that was deleted or a pattern that never matched. The
resolved-inputs comparison is what proves the command looked at something real: the corpus path must still appear in
`npx nx show project --json` output for the same targets after the change.

## Gherkin

**No `.feature` file is added, edited, moved, or deleted.** Both corpora are unchanged in content and in location:

- `specs/apps/badakmini-cli/behavior/` — five feature files, all `[unchanged]`.
- `specs/apps/wahidyankf-www/behavior/` — twelve feature files, all `[unchanged]`.

Every scenario in both corpora is preserved. None is changed, moved, deleted, or added.

**What changes is which project runs one adapter, not what any adapter asserts.** The Playwright binding set moves from
`apps/wahidyankf-www-e2e/steps/` to `apps/wahidyankf-www/tests/e2e/steps/`. The eight bound feature files stay bound to
the same steps asserting the same properties, and the four deliberately unbound files — `env-loader.feature`,
`port-resolver.feature`, `tier-env-loading.feature`, and `cv-export.feature` — stay unbound for the same recorded
reason. The skipped-scenario baseline of 34 is the number that holds this: it is unchanged by the move, and Phase 2's
gate asserts that exact value rather than asserting the suite merely passed.

**Adapter and binding paths that change:**

| Layer             | Before                           | After                                  |
| ----------------- | -------------------------------- | -------------------------------------- |
| Unit              | `src/**/*.unit.test.tsx`         | unchanged                              |
| Behavior          | `tests/bdd/`                     | unchanged                              |
| Local integration | `tests/integration/`             | unchanged                              |
| Process E2E       | `apps/wahidyankf-www-e2e/steps/` | `apps/wahidyankf-www/tests/e2e/steps/` |

No layer becomes incapable and none is dropped. The target that proves the result is
`npx nx run wahidyankf-www:test:e2e`, and the focused journey is the browser suite's existing pass over the eight bound
features against a real `next start`.

## Index Files

`specs/` requires an index entry for every immediate entry whatever its file type, and the
[documentation index policy](../../../../repo-governance/documentation-index-policy.md) owns that requirement. Two
indexes describe the E2E adapter's location in prose and are edited for accuracy, though neither gains or loses an
entry, because no file under `specs/` is added or removed:

- `[E]` `specs/apps/wahidyankf-www/README.md` — the sentence naming two projects sharing the corpus, the Process E2E
  row's adapter path, the skip-baseline sentence, and the two verification-command rows.
- `[E]` `specs/apps/wahidyankf-www/behavior/README.md` — the sentence pointing at the deleted project's README.

## C4

`specs/apps/wahidyankf-www/architecture.md` is the as-built model and it names the deleted project in four places. It is
updated in Phase 2, in the same commit as the merge, so the model never describes a container that does not exist.

**The view that changes: Container View.** The node `wahidyankf-www-e2e [Container: Playwright] Process E2E adapter` is
removed as a separate container and redrawn as a test-time process belonging to `wahidyankf-www`. The relationship it
carries — starting the application through `next start` and driving it over HTTP through a browser — is preserved
exactly, because that relationship is what makes it a real process boundary and the merge does not change it. What stops
being true is only the container's ownership: it is no longer a distinct deployable-scope unit with its own project.

**The constraint that changes.** The prose currently reads that the adapter "is a different toolchain from the
in-process behavior adapter", offered as the justification for a separate project. The toolchain difference is still
true and is kept; what is removed is the inference that it requires a separate project. The
[BDD policy](../../../../repo-governance/development/behavior-driven-development-policy.md) role matrix makes a
dedicated E2E project conditional — "Use a dedicated E2E project only when its adapter requires a different toolchain" —
and permissive rather than mandatory, so co-locating a different-toolchain adapter is the policy's stated default rather
than an exception to it. No governance change is required, and none is proposed.

**The table row that changes.** The project table listing `apps/wahidyankf-www` and `apps/wahidyankf-www-e2e` as two
projects sharing one model becomes a single-project table, and the sentence invoking the architecture specification
policy's shared-corpus provision is removed with it, because one project sharing a model with itself is not what that
provision describes.

All diagrams stay terminal-first ASCII in fenced `text` blocks. No Mermaid is introduced, and the redrawn Container View
is a modification of the existing ASCII diagram rather than a new format.

**What is not updated:** `specs/apps/badakmini-cli/architecture.md` is untouched. That project's boundaries, containers,
and components are unchanged by this plan; only its `project.json` command strings change, and a working directory is
not an architectural boundary.
