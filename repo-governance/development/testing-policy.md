---
tldr: "Defines the ordered, cacheable quick gate, strict coverage, and explicit integration tests."
when_to_use: "Use when adding, changing, or running project test and quality targets."
---

# Testing Policy

## Scope

This policy applies to every Nx project in this repository.

## The Target Contract

Every project exposes the same ten targets, so a workspace-wide command reaches all of them without special-casing one:

`typecheck`, `lint`, `test:unit`, `test:integration`, `test:e2e`, `test:coverage:unit`, `test:coverage:integration`,
`test:coverage:behavior`, `test:coverage`, and `test:quick`.

Three are eligibility-dependent, and their absence is a signal rather than a gap. A library never owns `test:e2e`. A
project that owns no real local boundary defines no `test:integration`, and therefore no `test:coverage:integration`.
Everything else is required of every project.

[Target Shape](testing-policy/target-shape.md) owns what each target declares — `cache`, `outputs`, `options.cwd`, and
shared `namedInputs` — and binds every target a project declares, not only these ten.

## Quick Tests

`typecheck`, `lint`, `test:unit`, `test:coverage`, and `test:quick` are cacheable. An application's process E2E target
is uncached and outside `test:quick`. The [BDD policy](behavior-driven-development-policy.md) owns its placement: a
co-located package uses the owner's `typecheck`, `lint`, and `test:e2e` targets, while a permitted dedicated project
owns equivalent targets. `test:quick` is an ordered `nx:run-commands` aggregate with parallel execution disabled; it
invokes its required fast target entry points without copying their underlying commands.

Two mechanisms express ordering and they are not interchangeable. `options.commands` expresses the ordered gate itself,
the sequence a reader runs. `dependsOn` expresses a prerequisite that must precede the whole gate, whatever entry point
invoked it. A build a gate cannot run without belongs in `dependsOn`; a step of the gate belongs in `options.commands`.

```text
typecheck -> lint -> test:unit -> test:coverage
```

Coverage uses the language's native instrumentation and fails below the project role's documented threshold. Do not
duplicate an executable threshold as metadata, omit runtime code, lower the threshold, or add broad exclusions to make
the gate pass; an unavoidable generated-code exclusion requires the repository owner's explicit approval and a
documented reason.

Pre-push invokes Nx affected with `origin/main` as the base and each pushed local commit as the head. Nx uses the
project graph to run `test:quick` only for affected projects under `apps/` and `libs/`, so unrelated documentation
changes do not run project tests and shared changes still reach every project they affect. The hook intentionally uses
Nx's local task cache.

## Integration Tests

Use an uncached `test:integration` target for the local-integration layer the
[BDD policy](behavior-driven-development-policy.md) requires. It is not enforced by pre-push because it can be slow. A
library has this target only when its BDD role requires the layer; every application has it. Run applicable integration
suites explicitly:

```sh
npm run test:integration
```

A library that owns no local boundary defines no `test:integration` target at all. Its absence is the signal that it has
nothing to integration-test, so a placeholder that echoes and exits earns a passing run without testing anything and
hides the same absence it claims to report. The [BDD policy](behavior-driven-development-policy.md) is canonical for
corpus sharing, adapter roles, and process E2E.

## Tooling

See [Testing Policy Details](testing-policy/README.md).

## Verification

Run `npm run typecheck`, `npm run lint`, `npm run test:unit`, `npm run test:coverage`, `npm run test:quick`, and any
applicable `npm run test:integration` before handing off behavior changes. To inspect the pre-push selection without
pushing, run:

```sh
npm exec nx -- affected -t test:quick --base=origin/main --head=HEAD
```
