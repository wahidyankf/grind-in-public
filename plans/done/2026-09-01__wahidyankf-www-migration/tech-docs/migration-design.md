# Migration Design

Eight transitions run inside this plan: a CV record consolidation, three dependency replacements, a linter replacement,
a configuration relocation, and two specification-corpus copies. None of them involves a credential or a private value,
so no value appears below — only names and locations, per
[plan document safety](../../../../repo-governance/conventions/plans-organization-policy/plan-document-safety.md).

## Inventory

| Source                  | Location or key                                            | Readers                                                                                                                                                                       | Writers                  | Accepted shape                                                 | Owner        | Destination                                                                                 | Compatibility                                                                   | Disposition proof                                                                          |
| ----------------------- | ---------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------ | -------------------------------------------------------------- | ------------ | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------ |
| CV record               | `apps/wahidyankf-www/src/features/cv/core/data.ts`         | CV route, PDF export, search index                                                                                                                                            | Owner, by hand           | `CVEntry[]`                                                    | Application  | Same path, unchanged                                                                        | Shape crosses unchanged                                                         | Phase 3: CV route and PDF tests pass                                                       |
| CV evidence base        | `cv/cv-raw.md`                                             | Owner, agents doing CV work                                                                                                                                                   | Owner, by hand           | Markdown prose                                                 | Repository   | `apps/wahidyankf-www/docs/cv-raw.md`                                                        | Copied byte-identical, then one broken link repaired in place (see below)       | Phase 4: checksum matches at copy time; `check:markdown-links` exits 0 after the repair    |
| LinkedIn profile draft  | `cv/cv-linkedin.md`                                        | Owner                                                                                                                                                                         | Owner                    | Markdown prose                                                 | Repository   | `apps/wahidyankf-www/docs/cv-linkedin.md`                                                   | Content byte-identical                                                          | Phase 4: checksum matches                                                                  |
| LinkedIn project drafts | `cv/linkedin-projects.md`                                  | Owner                                                                                                                                                                         | Owner                    | Markdown prose                                                 | Repository   | `apps/wahidyankf-www/docs/linkedin-projects.md`                                             | Content byte-identical                                                          | Phase 4: checksum matches                                                                  |
| ATS CV source           | `cv/cv-ats.md`                                             | `generate-cv-ats-pdf.py`                                                                                                                                                      | Owner                    | Markdown prose                                                 | Repository   | Deleted                                                                                     | Superseded by `data.ts`                                                         | Phase 4: the CV route renders every role the file listed                                   |
| ATS CV export           | `cv/cv-ats.pdf`                                            | Owner, recipients                                                                                                                                                             | `generate-cv-ats-pdf.py` | PDF                                                            | Repository   | Deleted                                                                                     | Superseded by `public/wahidyankf-kresna-fridayoka-cv.pdf`                       | Phase 4: the app-generated PDF opens and contains the same roles                           |
| ATS PDF generator       | `cv/generate-cv-ats-pdf.py`                                | Owner, via `uv`                                                                                                                                                               | —                        | Python, reportlab                                              | Repository   | Deleted                                                                                     | Superseded by `scripts/generate-cv-pdf.ts`                                      | Phase 4: `generate:cv-pdf` produces a readable PDF                                         |
| Shared UI components    | `@open-sharia-enterprise/web-ui`                           | 12 references in `src/` — 5 imports, 6 `vi.mock` targets, and 1 comment in `SearchSection.tsx`, which is removed rather than repointed, leaving 11 — plus 4 in the step files | `ose-public`             | React components                                               | `ose-public` | `src/features/ui/shell/`                                                                    | Same rendered output                                                            | Phase 3: ported component unit tests pass unchanged                                        |
| Design tokens           | `@open-sharia-enterprise/web-ui-token`                     | `src/app/globals.css`                                                                                                                                                         | `ose-public`             | Two CSS files                                                  | `ose-public` | `src/app/tokens.css` and `src/app/theme.css`                                                | Same custom properties                                                          | Phase 3: theme scenarios pass                                                              |
| Environment loader      | `@open-sharia-enterprise/ts-env-loader`                    | `src/env-loader.ts`, root port wrapper                                                                                                                                        | `ose-public`             | `loadTierEnv`, `resolveTier`, `tierEnvFilePath`, `resolvePort` | `ose-public` | `src/features/env/core/`                                                                    | Same five loader rules                                                          | Phase 3: all seventeen loader scenarios pass                                               |
| Loader Gherkin corpus   | `specs/libs/ts-env-loader/behaviours/gherkin/`             | The library's own unit bindings                                                                                                                                               | Owner                    | Two `.feature` files, 13 scenarios                             | `ose-public` | `specs/apps/wahidyankf-www/behaviours/tier-env-loading.feature` and `port-resolver.feature` | Scenario text byte-identical; one `Feature:` title changed to avoid a collision | Phase 2: 13 scenario titles match before and after                                         |
| Loader Gherkin bindings | `libs/ts-env-loader/src/*.unit.test.ts`                    | Vitest                                                                                                                                                                        | Owner                    | Three test files                                               | `ose-public` | `apps/wahidyankf-www/tests/bdd/`                                                            | Same assertions, imports retargeted                                             | Phase 3: `test:coverage:behaviour` exits 0                                                 |
| Env file reader         | `dotenv` package                                           | Inlined loader                                                                                                                                                                | npm                      | `config` with `processEnv` and `override: false`               | External     | Same package, now a direct pin in `apps/wahidyankf-www/package.json`                        | Unchanged call; only the manifest it is declared in moves                       | Phase 3: "process env wins" and "tolerates a missing tier file" scenarios pass             |
| Linter                  | `oxlint` and `apps/wahidyankf-www/oxlint.json`             | `lint` target                                                                                                                                                                 | Application              | oxlint config                                                  | Application  | Biome and `biome.json`                                                                      | Findings triaged, none suppressed broadly                                       | Phase 3: `lint` exits 0 with no blanket ignore                                             |
| App Gherkin corpus      | `specs/apps/wahidyankf/behaviours/wahidyankf-www/gherkin/` | Unit and E2E adapters                                                                                                                                                         | Owner                    | Nine `.feature` files, 40 scenarios                            | `ose-public` | `specs/apps/wahidyankf-www/behaviours/`                                                     | Scenario text byte-identical                                                    | Phase 2: forty scenario titles match before and after                                      |
| Deploy configuration    | `apps/wahidyankf-www/vercel.json`                          | Vercel                                                                                                                                                                        | Owner                    | JSON                                                           | Application  | Same path, parsed configuration identical, reformatted to this repository's Prettier width  | Branch gate preserved                                                           | Phase 6: checksum recorded before the reformat, then a parsed-JSON equality check after it |

`cv-raw.md` carries the one intentional divergence in this table. Its opening sentence links to `./cv-ats.md`, which
this plan deletes rather than moves, so the destination copy would name a file that never exists under
`apps/wahidyankf-www/docs/`. Badak Mini validates links in every tracked Markdown file, so leaving it would fail
`check:markdown-links` for the whole repository. Phase 4 therefore copies the file byte-identical, asserts the matching
digest, and only then replaces that one link with an inline-code reference to `src/features/cv/core/data.ts`, which is
what supersedes `cv-ats.md`. The digest assertion and the repair are two ordered items, not one, so the provenance of
the copy is still proved before anything is changed. `cv-linkedin.md` and `linkedin-projects.md` carry no relative link
at all and cross unchanged.

## Transition Order

The policy order is expand, migrate, verify, contract. Applied here it means nothing existing is deleted until its
replacement is proven, which matters most for `cv/`.

### 1. Expand

Phase 2 adds `specs/apps/wahidyankf-www/behaviours/` while the ose corpus continues to exist in its own repository.
Phase 3 adds the application, the inlined modules, and `apps/wahidyankf-www/docs/`, all alongside the untouched `cv/`
directory. At the end of Phase 3 the repository knowingly holds two CV records: the new authoritative one in `data.ts`
and the old one in `cv/`. That overlap is the expand step and it is deliberate.

### 2. Migrate

Phase 4 copies `cv-raw.md`, `cv-linkedin.md`, and `linkedin-projects.md` into `apps/wahidyankf-www/docs/`. The copy is
idempotent and identity-preserving: each destination file's SHA-256 is compared with its source before the source is
removed. Running the phase twice produces the same tree.

`cv-ats.md`, `cv-ats.pdf`, and `generate-cv-ats-pdf.py` are not copied anywhere. They are superseded, and their
disposition proof is behavioural rather than byte-comparative: the CV route must render every role that `cv-ats.md`
listed, and `generate:cv-pdf` must produce a readable PDF. That comparison happens before the deletion, not after.

### 3. Verify

Verification uses the normal product flow rather than a bespoke check. The CV page is rendered and its entries compared
against the roles the retired ATS source listed. `generate:cv-pdf` runs and its output is opened.

Restoration is rehearsed rather than assumed, and the rehearsal covers the whole directory rather than a sample. Before
the deletion, Phase 4 records the SHA-256 of all seven tracked files under `cv/` into `local-tmp/cv-digests-before.txt`.
Phase 4's gate then extracts `git archive HEAD cv` into `local-tmp/cv-recovery/`, confirms seven files return, and diffs
their digests against that record. Both halves are checklist items, so the rehearsal cannot be lost to a reading of this
paragraph.

Rehearsing only `cv-raw.md` would prove nothing that matters: that file is copied into `apps/wahidyankf-www/docs/` and
survives the phase regardless. The three files with no copy anywhere — `cv-ats.md`, `cv-ats.pdf`, and
`generate-cv-ats-pdf.py` — are the ones whose recovery source has to be real, and a binary PDF is exactly the case where
"it is in Git" is worth checking rather than assuming.

### 4. Contract

**This step deviates from the
[plan migrations](../../../../repo-governance/conventions/plans-organization-policy/plan-migrations.md) rule, and the
deviation is named here rather than left to be noticed.** That rule says to retain compatibility for a stated window and
schedule destructive deletion in a separately authorized later plan. Deleting `cv/` inside this plan does not do that.

The deviation is deliberate and its scope is exactly one directory. Consolidating the CV is the reason this plan exists;
a compatibility window here is a window in which two hand-maintained CVs coexist and the owner can publish the stale
one, which is the failure the plan is written to end. The stated window is therefore one commit: Phase 4 lands the
absorption and the deletion together. The recovery path is the Git history, which the Phase 4 gate rehearses across all
seven files before the deletion is committed, and the phase commit is the single revert point.

The plan follows the rule everywhere else. Removing the application from `ose-public` is left to a separate, separately
authorized plan, exactly as the `README.md` states.

`@open-sharia-enterprise/*` has no compatibility window at all: the specifiers are unresolvable in this repository from
the moment the application lands, so the inlining is not a migration with an overlap but a precondition of the
application building once.

## Reference Repair

Deleting `cv/` in Phase 4 breaks every reference to it held outside the directory. Those references are not all the same
thing, and the two kinds get opposite treatment by owner direction. Recording the distinction here is what keeps a later
reader from re-deriving it from the diff.

A **routing reference** tells an agent or a person where to read CV material. It is **repointed** at
`apps/wahidyankf-www/docs/`, because the material still exists and has simply moved.

A **scope enumeration** is a list of top-level directories that a rule, a workflow, or a subagent prompt applies to.
`cv/` is **struck from the list** and no `apps/` path replaces it. Repointing one would insert an `apps/` subdirectory
into a scope that deliberately excludes `apps/`: the
[documentation index policy](../../../../repo-governance/documentation-index-policy.md) does not cover `apps/` at all,
so an entry there would state a README rule the policy does not hold, and the `rules-checker` corpus lists bound what
that subagent reads, so an entry there would widen it. The CV material keeps the coverage it needs without any of that.
`readme-refresh.md`'s review list already names `apps/`, so `apps/wahidyankf-www/docs/README.md` is in that workflow's
scope through an entry that is already present.

Fifteen occurrences sit across nine files as this plan is authored, two routing and thirteen scope. Phase 4 labels each
one in `evidence/cv-references.txt` before editing anything, against this table.

| Location                                                              | The text carrying `cv/`                                                                                                                   | Kind    | Treatment                                                 |
| --------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------- | ------- | --------------------------------------------------------- |
| `AGENTS.md`                                                           | "For CV work, read `cv/README.md`"                                                                                                        | Routing | Repoint at `apps/wahidyankf-www/docs/README.md`           |
| `AGENTS.md`                                                           | "Every `docs/`, `repo-governance/`, `cv/`, `scripts/`, `plans/`, `specs/`, and harness directory requires an indexed README"              | Scope   | Strike `cv/`                                              |
| `CLAUDE.md`                                                           | "`cv/` holds career evidence; read `cv/README.md` before touching it"                                                                     | Routing | Repoint at `apps/wahidyankf-www/docs/README.md`           |
| `repo-governance/documentation-index-policy.md`                       | the `when_to_use` front-matter list                                                                                                       | Scope   | Strike `cv/`                                              |
| `repo-governance/documentation-index-policy.md`                       | "Every directory in ... must contain a `README.md`"                                                                                       | Scope   | Strike `cv/`                                              |
| `repo-governance/README.md`                                           | "Use it when adding, moving, or maintaining Markdown under ..."                                                                           | Scope   | Strike `cv/`                                              |
| `repo-governance/workflows/readme-refresh.md`                         | "Review the root `README.md` and every existing README below ..."                                                                         | Scope   | Strike `cv/`; `apps/` is already in this list             |
| `repo-governance/workflows/readme-refresh.md`                         | "Follow the documentation index policy everywhere it applies, which is ..."                                                               | Scope   | Strike `cv/`                                              |
| `repo-governance/workflows/rules-quality-gate/01-scope-and-corpus.md` | "`cv/`, `scripts/`, and the root `README.md` — each carries rule sentences, and `AGENTS.md` routes agents into `cv/README.md`"            | Scope   | Strike `cv/` and the now-false `AGENTS.md`-routing clause |
| `.claude/agents/rules-checker.md`                                     | "Read `docs/`, `cv/`, `scripts/`, and the root `README.md` narrowly"                                                                      | Scope   | Strike `cv/`                                              |
| `.claude/agents/rules-checker.md`                                     | "every directory README in `docs/`, `repo-governance/`, `cv/`, `scripts/`, `plans/`, `specs/`, and every harness directory registers ..." | Scope   | Strike `cv/`                                              |
| `.codex/agents/rules-checker.toml`                                    | the same corpus paragraph                                                                                                                 | Scope   | Strike `cv/`                                              |
| `.codex/agents/rules-checker.toml`                                    | the same README-registration sentence                                                                                                     | Scope   | Strike `cv/`                                              |
| `.opencode/agents/rules-checker.md`                                   | the same corpus paragraph                                                                                                                 | Scope   | Strike `cv/`                                              |
| `.opencode/agents/rules-checker.md`                                   | the same README-registration sentence                                                                                                     | Scope   | Strike `cv/`                                              |

The three `rules-checker` prompts are mirrored copies, so the same edit lands in all three and `check:harness-parity`
plus a three-way diff of the changed paragraphs proves it, as the
[harness capability parity policy](../../../../repo-governance/conventions/harness-capability-parity-policy.md)
requires.

## Mixed-Version Boundaries

There is one, and it is between repositories rather than inside this one. From Phase 3 until a later plan retires the
source, `ose-public` and this repository both hold a copy of the application, and both hold a CV record that can drift.
Nothing synchronizes them. The recorded source SHA in the plan README is the only anchor for diagnosing a later
divergence, which is why it is recorded rather than described as "current `main`".

## Retry and Rollback

Every migration step is idempotent, so a failed run is retried rather than repaired. A copy that already matches its
source is a no-op; a deletion of an already-absent path is a no-op.

Rollback is per-phase revert, as [technical design](README.md) states. The only irreversible-feeling step is the `cv/`
deletion, and it is reversible from the preceding commit. Phase 4's gate proves that specific recovery for all seven
files from the commit preceding the deletion — in execution that was `31aabe5^`, the deletion having been committed
before the rehearsal ran — and diffs their SHA-256 digests against the pre-copy record, so the rollback path is evidence
rather than an assumption.

## Malformed Input

If a ported `.feature` file fails the step-cardinality rule, or an inlined module fails to type-check, the file is
preserved as-is and the failure is reported into `learnings.md` with the exact path and message. Nothing is coerced,
silently reformatted, or dropped to make a phase gate pass, because a migration that appears successful by discarding
what it could not handle is worse than one that stops.

That sentence forbids a _silent_ reformat, and the Prettier reformat Phases 1, 3, 5, and 6 each run is neither silent
nor a workaround, so the two do not contradict. It is a checklist item with a verbatim command, `npm run format`, and
its own acceptance, `npm run format:check` exits 0; it is announced in
[technical design](README.md#toolchain-conformance-and-its-fallback) with the three configuration differences that make
it necessary; and it changes formatting alone — line breaks, quoting, and Tailwind class order — rather than content.
Conforming a ported file to the receiving repository's declared formatting source of truth is part of what porting it
means, not a way of making a gate pass with something the migration could not handle. The rule above still governs the
cases it names: a `.feature` file that fails the step-cardinality rule and a module that fails to type-check are
preserved as-is and reported into `learnings.md`, never reshaped into passing.
