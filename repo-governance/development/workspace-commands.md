---
tldr: "Lists the npm and Nx commands that build, test, and validate this workspace."
when_to_use: "Use when running, testing, or validating any part of the workspace."
---

# Workspace Commands

This document is the canonical command reference. `AGENTS.md` and `CLAUDE.md` link here rather than restating it, so the list cannot drift between them.

## Setup

- `npm install` installs pinned dependencies and enables Husky hooks.

## Build and Test

- `npm run build`, `npm run typecheck`, and `npm run lint` run the matching Nx targets.
- `npm run test:unit` runs deterministic unit suites.
- `npm run test:coverage` runs unit, local-integration, and behavior-completeness coverage targets.
- `npm run test:behavior` runs canonical corpus and adapter-completeness checks.
- `npm test` and `npm run test:quick` run the cacheable ordered quick gate: type-check, lint, unit test, unit coverage, then behavior completeness.
- `npm run test:integration` runs uncached local-integration targets; `npm run test:e2e` runs dedicated public-process suites; pre-push skips both.
- `npm run test:scheduled` runs quick verification, integration coverage, then E2E in that operational order.

Narrower runs:

```sh
go -C apps/badakmini-cli test ./internal/governance -run TestName
npx nx run badakmini-cli:test:unit
npx nx run badakmini-cli:test:integration
npx nx run badakmini-cli:test:coverage:unit
npx nx run badakmini-cli:test:coverage:integration
npx nx run badakmini-cli:test:coverage:behavior
npx nx run badakmini-cli:test:coverage
npx nx run badakmini-cli-e2e:test:e2e
npx nx affected -t test:quick --base=origin/main --head=HEAD
```

The [testing policy](testing-policy.md) owns the target contract and ordered `test:quick` sequence.

## Formatting

- `npm run format` and `npm run format:check` apply or verify Prettier, the formatting source of truth.

## Repository Checks

- `npm run check:governance` enforces the [document word limit policy](../conventions/document-word-limit-policy.md), which sets the limit and names every document it governs.
- `npm run check:harness-parity` compares the subagents, skills, and commands each harness exposes.
- `npm run check:markdown-links` validates repository-local Markdown links. It reads Git-tracked files, so `git add -N` a new document before trusting a local run.
- `npm run check:rule-change` announces the [rules-propagation](../workflows/rules-propagation.md) workflow for staged rule paths, and [harness-alignment](../workflows/harness-alignment.md) when a harness reads that path. It reports without blocking.
- `npm run check:workflows` validates GitHub Actions workflow syntax and schema with the owner-pinned Actionlint tool.
- `npm audit --audit-level=low` checks the locked dependency tree, and `npm run check:go-vulnerabilities` scans the Go module dependencies.

[Badak Mini](../../apps/badakmini-cli/README.md) implements the repository `check:` commands. When a check fails, read the [Badak Mini policy](badakmini-cli-policy.md) before changing it; the usual fix is the document, not the checker.

## Hooks

Pre-commit formats staged files and announces the rule workflows. Pre-push requires `origin/main` and uses the Nx project graph to run `test:quick` only for affected projects under `apps/` and `libs/`, comparing each pushed local commit with that base. It also runs the governance check when the push changes an instruction file, `repo-governance/`, or a harness directory, compares harness capabilities when a harness directory changes, and always validates Markdown links. See the [commit hook policy](commit-hook-policy.md).
