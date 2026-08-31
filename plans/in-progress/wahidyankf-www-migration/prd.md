# Product Requirements

## User Stories

As the repository owner, I want the personal site to live in this repository so that the practice workspace and its public face are maintained under one set of rules.

As the repository owner, I want exactly one CV record in the repository so that the published CV cannot silently disagree with the evidence behind it.

As the repository owner, I want the ported application held to 99% coverage with real integration and process-E2E layers so that moving it does not import a weaker standard along with the code.

As a future agent working here, I want the application to consume a canonical Gherkin corpus laid out like every other project's so that a behavior change revalidates the same way regardless of which project it touches.

As the repository owner, I want no .NET toolchain and no shared-library directory introduced by this move so that the workspace stays a Node and Go workspace with one runnable application per concern.

## In Scope

- Porting `apps/wahidyankf-www` and `apps/wahidyankf-www-e2e` into this repository.
- Inlining `web-ui`, `web-ui-token`, and `ts-env-loader` into the application.
- Flattening the ose corpus into `specs/apps/wahidyankf-www/behavior/` and authoring the as-built C4 model. The `ts-env-loader` corpus comes with the library that is inlined, for eleven ported feature files and 53 ported scenarios; a twelfth file, `cv-export.feature`, is authored here for the two new scenarios below.
- Raising coverage to 99% and replacing the `echo` no-op integration and E2E targets with real ones.
- Conforming the toolchain: TypeScript 6, Biome, the project-local ESLint commentary check, exact pins, `command`-shorthand Nx targets, lower-hyphenated filenames.
- Deleting `cv/` after the application absorbs it, and repairing every reference to it.
- Porting `vercel.json` unchanged and documenting the `prod-wahidyankf-www` branch in governance.

## Out of Scope

- Any change to `ose-public`.
- Promoting the `prod-wahidyankf-www` branch here, or moving the production domain.
- Redesigning any part of the user interface.
- Publishing `libs/` packages, or preserving `web-ui` as a reusable design system in this repository.
- Porting the app `Dockerfile` and `.dockerignore`. This repository deploys through Vercel and runs E2E against `next start`, so a container image has no consumer here.

## New Acceptance Scenarios

These are the behaviors this plan adds. The existing corpus is preserved rather than re-specified; [specification changes](tech-docs/specification-changes.md) names every preserved scenario, where it comes from, and why 53 of the 55 scenarios here have no RED cycle.

The local integration boundary has no coverage today, because `test:integration` is an `echo` no-op. These two scenarios give it real coverage against the filesystem the application genuinely owns.

```gherkin
Feature: CV export

  Scenario: Generating the CV writes a PDF to the local filesystem
    Given the application CV record contains at least one entry
    When the CV export runs against a writable output directory
    Then a readable PDF file exists at the configured output path
    And the file begins with the PDF header bytes

  Scenario: Generating the CV reports an unwritable output location
    Given the configured CV output directory does not exist
    When the CV export runs
    Then the export fails with a message naming the output path
    And no partial file is left behind
```

## Acceptance Criteria

- **[AC-1]** The repository contains exactly one CV record, and `cv/` no longer exists.
- **[AC-2]** `npx nx run wahidyankf-www:test:coverage:unit` enforces 99% and passes.
- **[AC-3]** `npx nx run wahidyankf-www:test:coverage:integration` enforces 99% and passes against the two scenarios above.
- **[AC-4]** Every one of the twelve feature files in `specs/apps/wahidyankf-www/behavior/` is loaded by exactly one binding — eleven from `apps/wahidyankf-www/tests/bdd/` and `cv-export.feature` from `apps/wahidyankf-www/tests/integration/` — proved by count, and `npx nx run wahidyankf-www:test:coverage:behavior` exits 0 across the 53 scenarios it reaches. Inside a loaded feature, `@amiceli/vitest-cucumber` fails the run on a scenario the binding does not implement, a binding declared for a scenario the file does not contain, a missing or mistyped step, and a missing Scenario Outline variable. Two things it does not do: it cannot see a feature file that no binding loads at all, which is why the count is asserted separately, and it does not report a step definition no scenario uses. That second gap is accepted and stated rather than closed; `ose-public` covered it with `rhino-cli`, which this plan deliberately does not bring.
- **[AC-5]** `npx nx run wahidyankf-www-e2e:test:e2e` runs the ported Playwright suite against `next start` with no Docker involved.
- **[AC-6]** No `@open-sharia-enterprise/*` specifier remains anywhere in the repository, and `libs/` contains only its `README.md`.
- **[AC-7]** No `dotnet`, `rhino-cli`, or `.fsproj` reference reaches any target, script, or workflow.
- **[AC-8]** `npm run format:check`, `npm run check:markdown-links`, `npm run check:governance`, and `npm run check:workflows` all exit 0.
- **[AC-9]** The toolchain matches `testing-policy/tooling.md`, or `tooling.md` records the specific component that could not conform and why.
- **[AC-10]** `apps/wahidyankf-www/vercel.json` parses to configuration identical to its source — every one of `installCommand`, `buildCommand`, `ignoreCommand`, and the header rules unchanged — with the pre-reformat source digest recorded in `evidence/vercel-json-digest.txt`, and a governance rule states what the `prod-wahidyankf-www` branch is for. The file is reformatted to this repository's Prettier width, which changes its bytes and not its meaning.
