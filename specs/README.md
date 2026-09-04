# Specifications

This directory describes what this repository's software should do and, for applications, its current architectural
boundaries. Specs state intent, observable behaviour, and as-built architecture; `apps/` and `libs/` contain the
implementation that delivers them.

The tree mirrors the workspace — `specs/apps/<name>/` and `specs/libs/<name>/` — and each subject carries a
`behaviours/` folder of Gherkin scenarios plus whatever `product/`, `system-context/`, `containers/`, or `components/`
detail it genuinely needs. Every non-drill application also has a root `architecture.md` C4 model of its current
as-built boundaries.

For the structure, the Gherkin cardinality rule, and when specs are required, read the
[specs policy](../repo-governance/development/specs-policy.md). For how scenarios bind to tests, read the
[TDD policy](../repo-governance/development/tdd-policy.md) and the
[BDD policy](../repo-governance/development/behaviour-driven-development-policy.md).

## Directory Map

- [Applications](apps/README.md) — the per-application C4 models and behaviour corpora.

## Current Specifications

Two subjects carry specifications. [Badak Mini](apps/badakmini-cli/README.md) has a canonical
[C4 model](apps/badakmini-cli/architecture.md) and an executable five-feature corpus, consumed by unit, local
integration, and public-process E2E adapters. [wahidyankf-www](apps/wahidyankf-www/README.md) has its own
[C4 model](apps/wahidyankf-www/architecture.md) and a twelve-feature corpus, consumed by Unit, local Integration, and
browser E2E adapters. Both applications keep their public-boundary E2E harness in a dedicated project. Every adapter
fails when a feature, step, binding, exemption, or adapter drifts.
