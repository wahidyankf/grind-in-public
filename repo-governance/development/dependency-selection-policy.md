---
tldr: "Prefers the standard library and existing repository mechanisms over a new external dependency."
when_to_use: "Use before adding any runtime, development, build, or test dependency."
---

# Dependency Selection Policy

## Scope

This policy covers every external dependency this repository takes on, in any manifest: Go modules under `apps/` and
`libs/`, and npm packages at the workspace root. It governs the decision to add one. The
[Nx workspace policy](nx-workspace-policy.md) separately forbids Nx plugins, executors, and generators, which is a
narrower prohibition than this one and is not relaxed by it.

## Prefer What Is Already Here

Reach for the language's standard library and the mechanisms this repository already owns before reaching outside. An
external package is not free at the moment it is added: it carries supply-chain exposure, upgrade work, compatibility
risk, and an ownership cost that outlives whoever added it. This is
[maintenance value](../principles/maintenance-value.md) applied to code written by someone else.

This repository exists for learning, which sharpens the point. A dependency that removes the exact work a drill was
meant to teach has cost more than it saved.

## When One Is Justified

Add an external dependency only when all three hold:

- the requirement cannot be met reasonably with the standard library or an existing repository mechanism, without a
  disproportionate cost in correctness, security, or interoperability;
- the dependency is the established practice for that problem, rather than a convenience specific to this repository;
  and
- current evidence from the project itself shows active maintenance and compatibility with the supported toolchain.

Record the requirement, the built-in alternatives you rejected, and the evidence, in the plan or the commit message that
adds it. Lock the version through the ecosystem's normal mechanism — `go.mod` and `go.sum`, or `package-lock.json` — and
never add one to avoid learning a capable standard-library facility or to save a small amount of code this repository
would own clearly.

## Removal

When the conditions stop holding, assess replacement the next time the dependency is materially changed or creates a
concrete problem. This policy does not ask for churn on a dependency that is merely old.

## Verification

Verified in review, against the changed manifest and lockfile: the need is stated, the rejected built-in alternative is
named, and the repository's checks pass with the dependency in place.

## Inherited and Transitive Pins

A version copied from another repository carries that repository's vulnerabilities, and a successful install is not the
check that finds them. Audit after adding or copying a pin, and resolve a finding by moving the pin rather than lowering
the audit level or adding an `overrides` entry.

A pin binds only the manifest that declares it. A dependency's own open range can still put two versions of one library
in the tree, which surfaces as a type error naming the same type on both sides rather than as an install failure; fix it
by agreeing the versions, never by casting at the call site. A library and a plugin that reaches into its internals are
one decision, because the plugin is coupled to a version the library's range does not express.
