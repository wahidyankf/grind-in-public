# Specifications

This directory describes what this repository's software should do and, for applications, its current architectural
boundaries. Specs state intent, observable behaviour, and as-built architecture; `apps/` and `libs/` contain the
implementation that delivers them.

The tree mirrors the workspace — `specs/apps/<name>/` and `specs/libs/<name>/` — and each subject carries a
`behaviours/` folder of Gherkin scenarios plus whatever `product/`, `system-context/`, `containers/`, or `components/`
detail it genuinely needs. Every non-drill application also has a root `architecture.md` C4 model of its current
as-built boundaries.

## Directory Map

- [Applications](apps/README.md) — the per-application C4 models and behaviour corpora.

## Current Specifications

No subject carries specifications at present. A directory appears under `specs/apps/` when the application it mirrors
gains a specification, not before.

Repository tooling under `scripts/` carries no Gherkin corpus. The corpus rule binds applications and libraries, and a
feature file no adapter consumes would claim coverage it does not have; those scripts are proved by the test file beside
each of them.
