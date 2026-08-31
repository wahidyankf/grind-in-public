# Specification Changes

This document records proposed specification work. `specs/` stays the canonical as-built truth until Phase 3 and Phase 5 land and verify.

## Durable Contracts Versus Plan-Only Outcomes

Fifty-three scenarios arrive as durable Gherkin: the forty that ship with the application, plus thirteen that ship with `ts-env-loader`, whose behavior this plan inlines into the application. They describe observable behavior and they survive the move. Two new scenarios join them, covering a local boundary that has no coverage today, for fifty-five in total.

The rest of what this plan must prove is **plan-only**. Each outcome below stays out of `specs/` because it describes the state of the repository rather than the behavior of the software, and Gherkin that asserts repository structure would bind a test to the plan rather than to the product.

| Plan-only outcome | Why it is not durable Gherkin | Delivery proof |
| --- | --- | --- |
| The repository holds exactly one CV record | Asserts repository layout, not site behavior | Phase 4 gate: `git ls-files cv` returns nothing and `rg` finds no surviving `cv/` reference |
| Unit and integration coverage reach 99% | A threshold is a gate, not an observable product behavior | Phase 3 gate: both coverage targets exit 0 |
| No `@open-sharia-enterprise/*` specifier remains | Dependency hygiene, invisible to a site visitor | Phase 3 gate: repository-wide search returns no match |
| No `dotnet`, `rhino-cli`, or `.fsproj` reference remains | Toolchain hygiene | Phase 3 gate: repository-wide search returns no match |
| The toolchain matches `tooling.md` or records its exception | Governance conformance | Phase 3 gate: the check passes or `tooling.md` carries the dated deviation |
| `vercel.json` parses to configuration identical to its source | Configuration fidelity, and the bytes are not what carries it — Vercel reads the parsed JSON | Phase 6, the two `vercel.json` items: the source SHA-256 is recorded into `evidence/vercel-json-digest.txt` while the copy is still byte-identical, then `npm run format` reformats it to this repository's width and a parsed-JSON equality check against the source blob exits 0. The delivered file does not checksum-match its source, and no gate asks it to |
| Every feature file is loaded by exactly one binding, `[AC-4]` | Asserts the completeness of the test harness, not a behavior a visitor can observe. The scenarios themselves are durable Gherkin; that each one is reached is a property of the adapters | Phase 3, `Bind the Corpus`: the per-feature binding count prints `1` for all eleven `tests/bdd/` files, then `npx nx run wahidyankf-www:test:coverage:behavior` exits 0 reporting 53 scenarios; `New Behavior` then proves `cv-export.feature` is loaded only by the integration test |
| The E2E suite runs against `next start` with no Docker, `[AC-5]` | Names the harness the scenarios run in, not what the site does. The same scenarios pass in `ose-public` inside a container, so this is a delivery-mechanism outcome | Phase 5 gate: `npx nx run wahidyankf-www-e2e:test:e2e` exits 0 and `rg -n 'docker' apps/wahidyankf-www-e2e` finds nothing |
| The four repository checks exit 0, `[AC-8]` | Repository hygiene — formatting, link resolution, document word limits, and workflow syntax — none of which the software does | Phase 7 gate: `npm run format:check`, `npm run check:markdown-links`, `npm run check:governance`, and `npm run check:workflows` all exit 0 |

## Why the Ported Scenarios Have No RED Cycle

The [TDD policy](../../../../repo-governance/development/tdd-policy.md) binds every behavior cycle to one scenario, and `plan-checker` reads that in both directions: every durable scenario should have a RED step. Fifty-three of the fifty-five here have none, and that is deliberate.

A RED step exists to prove the test fails before the code exists and fails for the stated reason. These fifty-three scenarios are already bound and already passing in `ose-public` against implementation code this plan copies rather than writes. Writing a RED step for them would mean deleting working code to watch a working test fail, which proves nothing about the port and risks losing the behavior in the round trip. What replaces the cycle is stricter, not looser: Phase 2 diffs the scenario titles against the source, Phase 3 proves each feature file is loaded by exactly one binding, and no phase gate closes while a scenario is unbound or coverage is below 99%.

The exemption stops there. The two CV export scenarios in `prd.md` are genuinely new, cover a boundary nothing asserts today, and run full RED, GREEN, and REFACTOR cycles in Phase 3.

## Gherkin: Copied and Preserved

All nine application feature files are copied from `$SRC/specs/apps/wahidyankf/behavior/wahidyankf-www/gherkin/<area>/` into `specs/apps/wahidyankf-www/behavior/`. Every scenario is preserved verbatim; none is renamed, split, merged, or deleted. Only the containing path changes.

Each is labelled `[N]`, not `[M]`. Nothing moves inside this repository: these are new files here, and the originals stay where they are, because `ose-public` is not touched. `[M]` would tell a reader to look for a path that disappeared, and none does.

`specs/apps/wahidyankf-www/behavior/accessibility.feature` `[N]` — Feature: Accessibility. Preserved: "Home page has zero axe-core WCAG 2.1 AA violations", "CV page has zero axe-core WCAG 2.1 AA violations", "Every page has exactly one H1", "Interactive controls expose accessible names".

`specs/apps/wahidyankf-www/behavior/responsive.feature` `[N]` — Feature: Responsive layout across viewports. Preserved: "Desktop viewport shows a fixed left sidebar", "Tablet viewport hides the sidebar and renders a bottom tab bar", "Mobile viewport hides the sidebar and renders a bottom tab bar", "The theme toggle is always reachable".

`specs/apps/wahidyankf-www/behavior/theme.feature` `[N]` — Feature: Theme toggle. Preserved: "Default theme is dark", "Clicking the toggle switches to light theme", "Theme persists across navigation", "Theme choice persists across reloads".

`specs/apps/wahidyankf-www/behavior/cv.feature` `[N]` — Feature: CV page. Preserved: "CV renders the Curriculum Vitae heading", "CV renders a search input", "CV renders the Highlights section header", "CV cross-linked via scrollTop query scrolls into the entries", "CV offers a downloadable PDF".

`specs/apps/wahidyankf-www/behavior/env-loader.feature` `[N]` — Feature: APP_ENV tier env-file loading. Preserved: "wahidyankf-www builds against the staging tier", "wahidyankf-www process env wins over the local tier file", "wahidyankf-www tolerates a missing tier file", and the Scenario Outline "wahidyankf-www fails loudly on a stray auto-loaded env file" with its Examples table intact.

`specs/apps/wahidyankf-www/behavior/home.feature` `[N]` — Feature: Home page. Preserved: "Home renders the welcome heading", "Home renders the About Me card", "Home renders the Skills & Expertise card with three subsections", "Home renders the Quick Links card with two internal links", "Home renders the Connect With Me card with five external links".

`specs/apps/wahidyankf-www/behavior/personal-projects.feature` `[N]` — Feature: Personal projects page. Preserved: "Personal projects page renders the heading", "Personal projects page renders a search input", "Personal projects page lists at least one project card", "Each project card exposes external links where applicable", "Each project card shows how long the project has been running", "Each project card exposes clickable skill tags", "Clicking a skill tag filters the project list".

`specs/apps/wahidyankf-www/behavior/search.feature` `[N]` — Feature: Search. Preserved: "Typing a term updates the URL query string", "Matching content is highlighted with a yellow mark", "Non-matching About Me shows a placeholder", "Clicking a skill pill navigates to the CV with scrollTop".

`specs/apps/wahidyankf-www/behavior/static-filterable-routes.feature` `[N]` — Feature: Static filtered portfolio routes. Preserved: "Search-filtered portfolio routes are static yet still filterable", "Public portfolio routes are available from the production server", "Crawlers receive discovery directives for every public route".

## Gherkin: Copied From the Inlined Library

Inlining `ts-env-loader` moves its behavior into this application, so its two feature files come with it. Without them the application would hold loader code whose only executable specification stayed in another repository.

`specs/apps/wahidyankf-www/behavior/port-resolver.feature` `[N]` — copied byte-identical from `$SRC/specs/libs/ts-env-loader/behavior/gherkin/port-resolver/port-resolver.feature`. Feature: Runtime listener port resolution. Preserved: "The CLI flag outranks every other source", "The prefixed variable outranks the fallback", "The fallback applies when nothing else supplies a port", "A bare PORT variable never moves the listener", the Scenario Outline "A blank value at a tier falls through to the next tier" with its three-row Examples table, the Scenario Outline "A present but malformed port fails loudly instead of falling through" with its ten-row Examples table, "A malformed prefixed variable names that variable in the error", and "An out-of-range compiled-in fallback is caught at startup". Eight scenarios, nothing changed.

`specs/apps/wahidyankf-www/behavior/tier-env-loading.feature` `[N]`, **changed** — copied from `$SRC/specs/libs/ts-env-loader/behavior/gherkin/env-loader/env-loader.feature`. This one is not a pure copy: its `Feature:` title changes from "APP_ENV tier env-file loading" to "Tier env-file loader module", because `behavior/env-loader.feature` in the same directory already carries the original title and two features cannot usefully share one in a flat corpus. Every scenario is preserved verbatim: "Loads the selected tier's file", "Process env always wins over the tier file", "Tolerates a missing tier file", the Scenario Outline "Fails loudly on a stray auto-loaded env file" with its three-row Examples table, and "Tolerates a stray file at the local tier". Five scenarios; nothing renamed, split, merged, or deleted. The narrative lines under the title are edited only to the extent that they name the module rather than the library.

The two files overlap in subject without duplicating: `env-loader.feature` states what the application does across tiers, and `tier-env-loading.feature` states what the loader module itself does. Both are kept because both were separately bound and separately passing in `ose-public`, and dropping either would silently reduce the specification this plan claims to preserve.

## Step Cardinality

Each copied file is checked against the [specs policy](../../../../repo-governance/development/specs-policy.md) step cardinality — exactly one primary `Given`, one `When`, and one `Then`, with `Background` and `Examples` exempt. Eight of the nine application files supply their `Given` through a `Background`, so a scenario in them legitimately shows no `Given` of its own. A scenario that arrives violating the rule is split in Phase 2, and the split is recorded in `learnings.md` because it means the source corpus disagreed with this repository's rule.

## Gherkin: New

`specs/apps/wahidyankf-www/behavior/cv-export.feature` `[N]` — Feature: CV export, holding two new scenarios. The local integration boundary is currently asserted by nothing at all, because `test:integration` is an `echo` no-op, so these are the first executable statements about it.

They get their own file rather than joining `cv.feature` for two reasons. `prd.md` writes them under `Feature: CV export`, and `cv.feature` is `Feature: CV page`. And `cv.feature` is loaded by `tests/bdd/cv.steps.ts`: `@amiceli/vitest-cucumber` throws `ScenarioNotCalledError` for any scenario in a loaded feature that the binding does not declare, so an integration-only scenario placed there would fail `test:coverage:behavior` for a reason that has nothing to do with the behavior. `cv-export.feature` is instead loaded by `tests/integration/cv-pdf.integration.test.ts` alone.

The corpus therefore ends at twelve feature files and 55 scenarios: eleven files and 53 scenarios bound from `tests/bdd/`, and this one bound from `tests/integration/`.

**"Generating the CV writes a PDF to the local filesystem."** The user is the repository owner running the CV export. The precondition is a CV record holding at least one entry and a writable output directory. The action is running the export. The expected outcome is a readable file at the configured output path whose first bytes are the PDF header.

**"Generating the CV reports an unwritable output location."** The user is the same. The precondition is a configured output directory that does not exist. The action is running the export. The expected outcome is a failure naming the output path, with no partial file left behind.

## Binding and Adapter Paths

| Layer | Adapter path | Change |
| --- | --- | --- |
| Unit and behavior | `apps/wahidyankf-www/tests/bdd/*.steps.ts` | `[N]` from `$SRC/apps/wahidyankf-www/test/unit/steps/`, nine files, retargeted at the flattened corpus and repointed off `@open-sharia-enterprise` |
| Unit and behavior | `apps/wahidyankf-www/tests/bdd/tier-env.unit.test.ts` | `[N]` from `$SRC/libs/ts-env-loader/src/env-loader.unit.test.ts`, bound to `tier-env-loading.feature` |
| Unit and behavior | `apps/wahidyankf-www/tests/bdd/port-resolver.unit.test.ts` | `[N]` from `$SRC/libs/ts-env-loader/src/port-resolver.unit.test.ts`, bound to `port-resolver.feature` |
| Contract | `apps/wahidyankf-www/tests/bdd/next-with-port-wrapper.unit.test.ts` | `[N]` from `$SRC/libs/ts-env-loader/src/next-with-port-wrapper.unit.test.ts`; not a Gherkin binding, but the only thing that executes the ported `scripts/next-with-port.mjs` |
| Integration | `apps/wahidyankf-www/tests/integration/cv-pdf.integration.test.ts` | `[N]`, the only loader of `cv-export.feature` and the only binding of its two scenarios |
| Process E2E | `apps/wahidyankf-www-e2e/steps/*.steps.ts` | `[N]` from `ose-public`, eight files, Docker runner removed, retargeted at the flattened corpus |

Every entry is `[N]`. Nothing moves within this repository; the sources stay in `ose-public`, which this plan does not touch.

### Specifically Incapable Layers

**The two new CV export scenarios have no unit binding and no E2E binding.** A unit test with a mocked filesystem would assert the mock, and the real boundary is the whole point. A browser cannot reach the export at all, because it is a build-time script rather than a route. Integration is the only capable layer. Proof runs through `npx nx run wahidyankf-www:test:coverage:integration`, whose focused journey is: construct a temporary directory, run the export against it, assert the header bytes, then repeat against a path that does not exist and assert the failure message and the absence of a partial file.

**Seventeen loader scenarios have no E2E binding: the four in `env-loader.feature`, the five in `tier-env-loading.feature`, and the eight in `port-resolver.feature`.** They describe how a Node process selects an env file and resolves a listener port before any server starts. A browser observes the result of those decisions, never the decisions, so a Playwright adapter has nothing to assert. The capable layer is the behavior project under Vitest, where all seventeen are bound.

That gap is not left implicit. `playwright-bdd` is configured with `missingSteps: "skip-scenario"`, which renders an unbound scenario as `test.fixme` and lets the suite exit 0 — so an accidentally unbound scenario looks exactly like a deliberately unbound one. `ose-public` closed that with `specs:e2e:coverage`, an `rhino-cli` subcommand that does not come here. Phase 5 replaces the one guarantee it gave: the nineteen deliberately unbound scenarios are named in `apps/wahidyankf-www-e2e/README.md`, the generated `test.fixme` count is recorded as a baseline beside them and, as a single integer, in the tracked `apps/wahidyankf-www-e2e/e2e-skip-baseline.json`, and a `specs:e2e:baseline` target fails when the regenerated count differs from that file.

## C4 Model

`specs/apps/wahidyankf-www/architecture.md` `[N]` is authored fresh rather than copied. `ose-public` splits its model across four directories — `product/overview.md`, `system-context/context.md`, `containers/container.md`, and `components/web/component-web.md` — while this repository keeps one `architecture.md` per corpus, as `specs/apps/badakmini-cli/architecture.md` does. The four source documents are consolidated into one.

The model records the final as-built state this plan delivers, not the source repository's state. The three differences below land in different phases — the component view in Phase 3, the system context in Phase 4, the container view in Phase 5 — while the model itself is authored in Phase 2, so a Phase 5 `delivery.md` item checks each one against the delivered tree and corrects the model where they disagree. Without that check the model would be a Phase 2 prediction wearing an as-built label, which is what the [architecture specification policy](../../../../repo-governance/development/architecture-specifications.md) forbids. Three things differ from the source and must be written that way:

- **Container view**: the runtime container set loses the Docker image. The application runs as a Next production server started by `next start`, and the E2E adapter drives that same process. No container node appears.
- **Component view**: the shared design-system component disappears. `web-ui`, `web-ui-token`, and `ts-env-loader` were external components with a dependency relationship into the application; after inlining, the four UI components and the environment loader are internal modules of the application itself, so the relationship is deleted rather than redrawn.
- **System context**: the CV evidence store moves inside the application boundary. `cv/` was a separate repository-level source of CV truth; after Phase 4 there is one CV record inside the application, so the context view shows one store where it previously showed two.

Every diagram is terminal-first ASCII in a fenced `text` block, per the [markdown style policy](../../../../repo-governance/conventions/markdown-style-policy.md) and the [architecture specification policy](../../../../repo-governance/development/architecture-specifications.md). The model links to `specs/apps/wahidyankf-www/behavior/`, and both `apps/wahidyankf-www/README.md` and `apps/wahidyankf-www-e2e/README.md` link back to it, because the implementation and the dedicated E2E project share one corpus and one model.
