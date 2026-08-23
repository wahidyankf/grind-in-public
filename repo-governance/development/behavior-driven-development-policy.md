---
tldr: "Makes canonical Gherkin behavior executable through the required test layers for every app and library."
when_to_use: "Use when adding, changing, or reviewing behavior, test targets, specs, or project roles."
---

# Behavior-Driven Development Policy

## Scope

Every application and library owns executable Gherkin behavior under its mirrored `specs/` path. A change to a feature, scenario, step, binding, or adapter must rerun the applicable corpus and fail when the corpus, bindings, or adapters disagree.

## Role Matrix

| Project role | Required behavior adapters |
| --- | --- |
| Application | Unit, local-only integration, and process E2E in a dedicated E2E application. |
| Library | Unit; local-only integration only when the library owns a real local boundary. Never E2E. |
| Dedicated E2E application | Its owner's corpus through the public process boundary; no separate corpus, unit layer, or numeric coverage gate. |

Applications and libraries use `specs/apps/<name>/behavior/` and `specs/libs/<name>/behavior/`. Every adapter recursively discovers the same canonical corpus, validates every expression exactly once, and uses uncached test execution when the language cannot include external specs in its native test cache key. Nx inputs include that recursive corpus and every E2E binding and configuration that influences behavior completeness.

## Required Layers

Unit tests use doubles for external collaborators. Integration tests use only owned local boundaries such as local filesystems, databases, or subprocesses; they never use network services. A dedicated E2E app invokes only the owning application's built public executable or API and observes public results.

Each project README names its corpus, adapters, targets, and any justified inapplicable layer. See [Specs](specs-policy.md), [TDD](tdd-policy.md), and [Testing](testing-policy.md) for the complementary requirements.
