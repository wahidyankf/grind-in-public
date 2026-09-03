---
tldr: "Makes canonical Gherkin behavior executable through the required test layers for every app and library."
when_to_use: "Use when adding, changing, or reviewing behavior, test targets, specs, or project roles."
---

# Behavior-Driven Development Policy

## Scope

Every application and library owns executable Gherkin behavior under its mirrored `specs/` path. A change to a feature,
scenario, step, binding, or adapter must rerun the applicable corpus and fail when the corpus, bindings, or adapters
disagree.

## Role Matrix

| Project role          | Required behavior adapters                                                                                                                                       |
| --------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Application           | Unit, local-only integration, and process E2E co-located in the owner project. Use a dedicated E2E project only when its adapter requires a different toolchain. |
| Library               | Unit; local-only integration only when the library owns a real local boundary. Never E2E.                                                                        |
| Dedicated E2E project | Only a different-toolchain exception; its owner's corpus through the public process boundary, with no separate corpus, unit layer, or numeric coverage gate.     |

Applications and libraries use `specs/apps/<name>/behavior/` and `specs/libs/<name>/behavior/`. Every adapter
recursively discovers the same canonical corpus, validates every expression exactly once, and uses uncached test
execution when the language cannot include external specs in its native test cache key. Nx inputs include that recursive
corpus and every E2E binding and configuration that influences behavior completeness.

## Required Layers

Unit tests use doubles for external collaborators. Integration tests use only owned local boundaries such as local
filesystems, databases, or subprocesses; they never use network services. An E2E adapter invokes only the owning
application's built public executable or API and observes public results.

Godog is the execution base for Go behavior tests. A Go subject's binding set and its separate E2E binding set each
register `Given`, `When`, and `Then` directly on `*godog.ScenarioContext` and execute through `godog.TestSuite`;
repository compliance tooling may inspect the corpus and bindings but must not replace Godog's runtime registration,
matching, lifecycle, or execution.

Each project README names its corpus, adapters, targets, and any justified inapplicable layer. See
[Specs](specs-policy.md), [TDD](tdd-policy.md), and [Testing](testing-policy.md) for the complementary requirements.

## Bindings and What They May Assume

A scenario states a property; any concrete list a binding checks it against belongs to the binding and must track the
code. When two adapters bind one scenario, they assert the same property against the same current behavior, so a
disagreement between them is a defect in one of them rather than a difference of layer.
