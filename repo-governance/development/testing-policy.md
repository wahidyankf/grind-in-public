---
tldr: "Defines ordered, cacheable quick gates, strict coverage, and explicit integration and E2E ownership."
when_to_use: "Use when adding, changing, or running project test and quality targets."
---

# Testing Policy

The [quality-gates standard](quality-gates.md) owns test boundaries and the target contract. This policy owns how those
targets operate in this Nx workspace; [target shape](testing-policy/target-shape.md) owns their declarations.

## Owner Applications

Every application exposes `build`, `typecheck`, `lint`, `test:unit`, `test:integration`, `test:coverage:unit`,
`test:coverage:integration`, `test:coverage:behaviour:<owner>`, `test:coverage:behaviour`, `test:coverage`, and
`test:quick`.

`test:quick` is a cacheable ordered aggregate with parallel execution disabled:

```text
typecheck -> lint -> test:unit -> test:coverage:unit -> test:coverage:behaviour
```

`test:coverage` composes every owner numeric and behaviour slice. Numeric line coverage stays at least 99%. Do not lower
a threshold, omit runtime code, or add broad exclusions to make a gate pass.

## Dedicated E2E Projects

Every application public browser or process boundary has a dedicated E2E project. It exposes `typecheck`, `lint`,
`test:coverage:behaviour:e2e`, `test:coverage:behaviour`, `test:e2e`, and `test:quick`; operational install targets are
allowed when the harness needs them. It exposes no unit, integration, or numeric-coverage placeholder.

Its `test:quick` runs only `typecheck`, `lint`, and `test:coverage:behaviour:e2e`. Its generic behaviour target
delegates to the owner's aggregate. The owner aggregate composes `test:coverage:behaviour:<owner>` with the E2E slice.
Declare inputs and project dependencies so a feature, binding, config, or owner change invalidates the correct cache
without an Nx cycle.

## Runtime and Selection

Integration and E2E targets are uncached and stay outside `test:quick` and Git hooks. Run affected suites before
completion. Scheduled CI runs complete integration coverage before complete E2E.

Tests create only synthetic state inside boundaries owned by their run and fail closed instead of reusing an unverified
process, identity, filesystem root, browser context, or data store. Follow [test data isolation](test-data-isolation.md)
for boundary validation, marking, and cleanup.

Pre-push invokes affected `test:quick` against `origin/main`; independent selected projects may overlap within HIPPO's
fixed allocation. It conditionally runs repository validation when governance, harness, workflow, project configuration,
or behaviour-compliance machinery changes. It intentionally uses the Nx cache. Invoke compute-bearing outer commands
through the pinned consumer described by [resource-aware development](resource-aware-development.md); aggregate targets
remain unaware of host admission.

Aggregate targets invoke named Nx targets and never duplicate their tool commands. `options.commands` expresses an
ordered gate; `dependsOn` expresses prerequisites that precede the whole target. Use the standard library and existing
tooling first; dependency additions remain governed by the
[dependency selection policy](dependency-selection-policy.md).

## Verification

Run the applicable commands from [workspace commands](workspace-commands.md), including owner quick, integration
coverage, E2E quick, E2E runtime, and `badakmini-cli:test:repo` for repository-mechanism changes.
