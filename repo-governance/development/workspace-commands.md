---
tldr: "Lists the npm and Nx commands that build, test, and validate this workspace."
when_to_use: "Use when running, testing, or validating any part of the workspace."
---

# Workspace Commands

This document is the canonical command reference. `AGENTS.md` and `CLAUDE.md` link here rather than restating it, so the
list cannot drift between them.

## Setup

- `rtk ./hippo run --class transactional --disk-path . -- npm install` installs pinned dependencies and enables Husky
  hooks.

## Build and Test

Guard compute-bearing commands shown below with `rtk ./hippo run --class ephemeral --disk-path . -- <command>`.
Aggregate targets remain unguarded internally so one outer admission owns the whole run.

- `npm run build`, `npm run typecheck`, and `npm run lint` run the matching Nx targets.
- `npm run test:unit` runs deterministic unit suites.
- `npm run test:coverage` runs unit, local-integration, and behaviour-completeness coverage targets.
- `npm run test:behaviour` runs canonical corpus and adapter-completeness checks.
- `npm test` and `npm run test:quick` run the cacheable ordered quick gate: type-check, lint, unit test, unit coverage,
  then behaviour completeness.
- `npm run test:integration` runs uncached local-integration targets; `npm run test:e2e` runs dedicated public-process
  suites; pre-push skips both.
- `npm run test:scheduled` runs all four project quick gates, both owner integration-coverage gates, then both dedicated
  E2E suites in that operational order.

Narrower runs — prefix each with `rtk ./hippo run --class ephemeral --disk-path . --`, except the two marked
`# transactional`, which take `--class transactional`:

```sh
node --test scripts/rule-change.test.mjs
npm exec -- nx run wahidyankf-www:test:unit
npm exec -- nx run wahidyankf-www:test:integration
npm exec -- nx run wahidyankf-www:test:coverage:unit
npm exec -- nx run wahidyankf-www:test:coverage:integration
npm exec -- nx run wahidyankf-www:test:coverage:behaviour
npm exec -- nx run wahidyankf-www:test:coverage
npm exec -- nx run wahidyankf-www:test:quick
npm exec -- nx run wahidyankf-www:static-routes:validation
npm exec -- nx run wahidyankf-www:generate:cv-pdf  # transactional
npm exec -- nx run wahidyankf-www-e2e:test:coverage:behaviour:e2e
npm exec -- nx run wahidyankf-www-e2e:test:quick
npm exec -- nx run wahidyankf-www-e2e:install  # transactional
npm exec -- nx run wahidyankf-www-e2e:test:e2e
npm exec -- nx affected -t test:quick --base=origin/main --head=HEAD
```

Run `wahidyankf-www-e2e:install` once per machine before its E2E suite. The E2E target builds and starts the owner.

The [testing policy](testing-policy.md) owns the target contract and ordered `test:quick` sequence.

## Formatting

- `npm run format` and `npm run format:check` apply or verify Prettier, the formatting source of truth.

## Repository Checks

- `npm run check:hygiene` is `./rhino gate run --surface ci`, which runs every gate that surface declares.
- `npm run check:governance` enforces the [document word limit policy](../conventions/document-word-limit-policy.md).
- `npm run check:harness-parity` validates instructions, skills, agent adapters, and the digest.
- `npm run check:markdown-links` validates repository-local Markdown links. It reads Git-tracked files, so `git add -N`
  a new document before trusting a local run.
- `npm run check:project-contract` validates the deterministic owner/E2E descriptor contract for every declared pair.
- `npm run check:workflow-contract` validates stable authorization, terminal-result, convergence, and TDD-evidence
  tokens without attempting semantic review.
- `npm run test:repo` runs every deterministic repository mechanism, including directory maps and frontmatter.
- `rtk ./.github/scripts/test-hippo-bootstrap.sh` verifies the pinned consumer without network or the real installation
  cache; `rtk ./hippo version --json` verifies the installed release identity.
- `npm run check:rule-change` automatically triggers the [rules-propagation](../workflows/rules/rules-propagation.md)
  workflow for staged rule paths, and [harness-alignment](../workflows/harness-alignment.md) when a harness reads that
  path. It reports without blocking.
- `npm run check:workflows` validates GitHub Actions workflow syntax and schema with the owner-pinned Actionlint tool.
- `npm audit --audit-level=low` checks the locked dependency tree, and `npm run check:go-vulnerabilities` scans the Go
  module dependencies.

[RHINO](https://github.com/wahidyankf/rhino) implements the checks and the gate dispatch from `repo-config.yml`, through
the `./rhino` consumer pinned in `rhino.lock`; the scripts under `scripts/` implement the rest. When a check fails the
usual fix is the document, not the checker; read the [repository check policy](repository-check-policy.md) before
changing one.

## Hooks

Each hook dispatches a declared surface rather than listing checks, so what runs is read from `repo-config.yml`.
`public-safety` is first on every surface, because this repository publishes.

Commit-msg screens and lints the message. Pre-commit screens and formats the staged tree, then triggers applicable rule
workflows. Pre-push additionally requires `origin/main`, runs affected `test:quick` targets within HIPPO's fixed
allocation with Nx Cloud disabled, and conditionally runs guarded `test:repo`. See the
[commit hook policy](commit-hook-policy.md).
