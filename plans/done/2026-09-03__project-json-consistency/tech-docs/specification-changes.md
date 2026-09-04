# Specification Changes

This document records what this plan proposes to change in `specs/`. `specs/` stays the canonical as-built truth until
Phase 2's gate passes; nothing here is written into it ahead of the implementation it describes.

## Durable Contracts Versus Plan-Only Outcomes

Every acceptance criterion in [prd.md](../prd.md) is a **plan-only operational outcome**. None becomes a Gherkin
scenario in `specs/`, and none is bound to a test.

**The reason.** All seven criteria describe build configuration, project layout, or the content of a governance
document. `AC-1` through `AC-5` assert properties of `project.json` and `nx.json`; `AC-6` asserts which Nx project owns
a target; `AC-7` asserts that a policy states a rule. None of them is behaviour of the software this repository ships.
The [specs policy](../../../../repo-governance/development/specs-policy.md) reserves `specs/` for what an application
should do and how it is built as a system, and a scenario there binds to a test that fails when the behaviour breaks.
There is no application behaviour here to break: after this plan, `wahidyankf-www` renders exactly what it renders today
and `badakmini-cli` reports exactly what it reports today.

**Delivery proof for each.** The policy requires a plan-only outcome to name its proof rather than merely assert it is
unprovable in `specs/`:

| Criterion | Proved at    | By                                            |
| --------- | ------------ | --------------------------------------------- |
| `[AC-1]`  | Phase 2 gate | `npx nx show project` for both projects       |
| `[AC-2]`  | Phases 1-3   | the cache and the outputs inspections         |
| `[AC-3]`  | Phase 1 gate | a `grep` count and a cache-invalidation probe |
| `[AC-4]`  | Phase 1 gate | a `grep` plus `test:quick` and `test:e2e`     |
| `[AC-5]`  | Phase 1 gate | a `grep` for a bare `nx run`                  |
| `[AC-6]`  | Phase 2 gate | `nx show projects`, `test:e2e`, the baseline  |
| `[AC-7]`  | Phase 3 gate | the policy-against-`project.json` review      |

`[AC-1]` requires both project inspections to list all ten contract targets. `[AC-2]` is proved by two commands, one per
half of its scenario. The cache inspection must report no target whose `cache` resolves to undefined, and must name a
non-zero target count so the check is proved to have read a populated set. The outputs inspection must report no
uncached target declaring `outputs` and no cached target writing a mapped artifact without declaring it, against an
artifact map written out per project because no command can infer which cached target writes something. Both inspections
run in the Phase 1 gate over the three pre-merge projects, again in the Phase 2 gate over the two post-merge ones with
the `wahidyankf-www` map extended to carry the `specs:e2e:baseline` target that phase adds, and once more in the Phase 3
gate against the rule that phase writes. The Phase 2 placement is deliberate: a failure there is repaired in the
`project.json` that phase is already changing, inside its own commit theme, where the same repair forced in Phase 3
would put configuration in a documentation commit. Phase 4 reconciles this criterion against that extended, post-merge
form. `[AC-6]` requires three things together: exactly two discovered projects, a passing browser suite, and a
skipped-scenario count still equal to 34.

`[AC-7]`'s proof is the Phase 3 gate item that reads `repo-governance/development/testing-policy.md` against
`apps/badakmini-cli/project.json` and `apps/wahidyankf-www/project.json` and writes down, rule by rule, which target in
which file it was checked against. `npm run check:governance` and `npm run check:markdown-links` also run in that gate,
but neither observes this criterion: the first counts words and the second resolves links, and both pass over a policy
that says nothing at all about the ten targets. They are kept as the constraints the edit must respect, not as its
proof.

`AC-3`'s proof is deliberately two commands rather than one. The `grep` count is the "declared once" half: it must print
`1` per `project.json`, the single occurrence being the `namedInputs` declaration. On its own a count is equally
satisfied by a name nothing references, so it is paired with a half that proves the reference reaches the corpus. That
half is a cache-invalidation probe, not a reading of `npx nx show project --json`: the `--json` output reports the
**declared** `inputs` array verbatim and expands neither `default` nor a `namedInputs` reference, so after this refactor
it prints `["default", "behaviourCorpus"]` and the corpus path is gone from it by construction — a comparison against
the pre-change output would report a difference on every affected target and could never distinguish a working reference
from a broken one. Nx hashes file content instead, so appending a Gherkin comment line to a corpus feature file must
turn a cache hit into a miss on every target naming `behaviourCorpus`, and restoring the file with `git checkout --`
must turn it back into a hit. That is the probe `delivery.md`'s Phase 1 gate runs, and every probe ends by asserting
`git diff --stat` no longer names the file it touched, so the corpus is provably byte-identical afterwards.

## Gherkin

**No `.feature` file is added, edited, moved, or deleted.** Both corpora are unchanged in content and in location:

- `specs/apps/badakmini-cli/behaviours/` — five feature files, all `[unchanged]`.
- `specs/apps/wahidyankf-www/behaviours/` — twelve feature files, all `[unchanged]`.

Every scenario in both corpora is preserved. None is changed, moved, deleted, or added.

Phase 1's cache probe is the one place a `.feature` file is written to at all. It appends a single `#` comment line to
`capability-parity.feature` and to `accessibility.feature`, then restores each with `git checkout --` inside the same
checklist item and asserts `git diff --stat` names neither file afterwards. `#` opens a Gherkin comment, so even the
transient state holds no scenario change, and no committed state differs from the baseline.

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
| Behaviour         | `tests/bdd/`                     | unchanged                              |
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
- `[E]` `specs/apps/wahidyankf-www/behaviours/README.md` — the sentence pointing at the deleted project's README.

## C4

`specs/apps/wahidyankf-www/architecture.md` is the as-built model and it carries the deleted project in four places: the
Scope sentence naming "one dedicated E2E project" and the shared-model provision it invokes, the project table row, the
Container View node, and the paragraph below the diagram that justifies drawing it. Two of the four hold the literal
name; the other two describe the project without naming it. It is updated in Phase 2, in the same commit as the merge,
so the model never describes a container that does not exist. Its two later references to "the E2E adapter", under
Process and under Environment, are not among the four: each names a role that survives the merge and each stays true, so
[file-impact.md](file-impact.md) records them as untouched.

**The view that changes: Container View.** The lower node is redrawn. The relationship it carries — starting the
application through `next start` and driving it over HTTP through a browser — is preserved exactly, arrow and label
included, because that relationship is what makes it a real process boundary and the merge does not change it. What
stops being true is only the node's ownership: it is no longer a distinct deployable-scope unit with its own project.
Both blocks are written out here rather than described, because a described diagram is one the executor has to invent.

Before, exactly as `specs/apps/wahidyankf-www/architecture.md` holds it today:

```text
   +-----------+                +----------------------------------+
   |  Visitor  | -- HTTPS ----> |               web                |
   | [Person]  |                |  [Container: Next.js 16]         |
   +-----------+                |  Statically rendered at build    |
                                +----------------------------------+
                                                 ^
                                                 | starts and drives
                                                 | over a local port
                                  +--------------------------------+
                                  |     wahidyankf-www-e2e         |
                                  |  [Container: Playwright]       |
                                  |  Process E2E adapter           |
                                  +--------------------------------+
```

After, and this is the block the Phase 2 edit produces character for character:

```text
   +-----------+                +----------------------------------+
   |  Visitor  | -- HTTPS ----> |               web                |
   | [Person]  |                |  [Container: Next.js 16]         |
   +-----------+                |  Statically rendered at build    |
                                +----------------------------------+
                                                 ^
                                                 | starts and drives
                                                 | over a local port
                                  +--------------------------------+
                                  |  apps/wahidyankf-www/tests/e2e |
                                  |  [Test-time process]           |
                                  |  Playwright                    |
                                  |  Process E2E adapter           |
                                  +--------------------------------+
```

**Why the lower node loses its `[Container: ...]` stereotype.** The Containers section opens "One container." and its
table lists one row, `web`. The diagram contradicts that today: it draws a second box labelled `[Container: Playwright]`
that the sentence and the table both deny. The merge is the occasion to resolve it, and it is resolved in the direction
the prose already takes rather than the direction the diagram does. The opening sentence and the single-row table are
**not** edited; the lower node drops the C4 container stereotype and gains `[Test-time process]` instead, so the count
the section states and the boxes the diagram draws finally agree. The alternative — keeping `[Container: ...]` and
rewriting the sentence to "Two containers" — would assert that a Playwright harness is a deployable unit of this system,
which it is not: it is never deployed, it exists only while a test runs, and nothing serves traffic from it. The node
keeps its box because it is still a real process boundary, and the paragraph below the diagram is what says why a
non-container box appears in a container view at all.

**The constraint that changes.** The prose currently reads that the adapter "is a different toolchain from the
in-process behaviour adapter", offered as the justification for a separate project. The toolchain difference is still
true and is kept; what is removed is the inference that it requires a separate project. The
[BDD policy](../../../../repo-governance/development/behaviour-driven-development-policy.md) role matrix makes a
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
