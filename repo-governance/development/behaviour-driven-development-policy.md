---
tldr: "Makes canonical Gherkin executable through every applicable test layer with explicit boundary exemptions."
when_to_use:
  "Use when adding, changing, reviewing, or validating features, scenarios, bindings, adapters, or exemptions."
---

# Behaviour-Driven Development Policy

Every application and library owns one recursive canonical Gherkin corpus under its mirrored `specs/` path. Unit,
applicable local integration, and dedicated application E2E adapters consume that corpus; an E2E project owns no
separate features. A feature, scenario, step, binding, config, or adapter change must invalidate and rerun applicable
static compliance.

## Iron Rule

Read relevant features, scenarios, steps, and tests before changing production code. For every observable behaviour
change: update one canonical scenario -> bind its failing steps in every applicable adapter -> confirm the expected Nx
RED -> implement production code. Never skip or reorder this sequence. Refactors preserve Gherkin, begin from green or
characterization coverage, and remain green.

## Applicable Layers

- Unit proves behaviour with injected doubles and no real OS-facing dependency.
- Integration proves real local filesystem, environment, or same-machine process behaviour without network.
- E2E proves a built application's public browser or process boundary.

Every scenario has a substantive unit binding and either a substantive binding or a valid exemption for Integration and
E2E. When both upper layers appear inapplicable, split or redesign the scenario so one meaningful boundary supplies
alternative proof. The [quality-gates standard](quality-gates.md) owns exact boundaries.

Applications require Unit, local-only Integration, and dedicated-project E2E adapters. Libraries require Unit and add
Integration only when they own a real local resource boundary; libraries never own E2E. Dedicated E2E projects implement
their owner's corpus and never become behaviour owners.

## Exemptions

Only these scenario-level tags are valid:

```gherkin
# Exemption(integration): <boundary reason>; alternative-proof: <Nx target> / <scenario>
@integration-exempt
Scenario: ...

# Exemption(e2e): <boundary reason>; alternative-proof: <Nx target> / <scenario>
@e2e-exempt
Scenario: ...
```

The structured comment must be immediately before its tag. One scenario may not carry both tags. Feature-level,
`@unit-exempt`, malformed, and legacy layer-filter tags are forbidden. Slowness, implementation difficulty, flakiness,
and deferred work are not boundary reasons. The named target and scenario must exist and provide substantive proof.
Remove layer-specific no-op or success-sentinel branches for exempt scenarios.

## Deterministic Compliance

`test:coverage:behaviour` recursively discovers features, expands every Scenario Outline row, and checks every
applicable adapter. It must reject undefined, ambiguous, duplicate, and unused bindings; invalid exemptions; missing
alternative proof; adapter corpus drift; and direct non-Gherkin E2E specs. Discovery and diagnostics are stable and
sorted. Static compliance is read-only, network-free, and cannot replace runtime registration or the
[Gherkin implementation review](../workflows/gherkin-implementation-review.md).

Godog remains the runtime for Go behaviour. Each Go adapter registers Given, When, and Then directly on
`*godog.ScenarioContext`. TypeScript adapters use their project runner but obey the same exact-corpus, binding,
lifecycle, and semantic requirements.

Each project README names its corpus, applicable adapters, targets, exemptions, and justified omissions. Follow
[Specs](specs-policy.md), [TDD](tdd-policy.md), [test-data isolation](test-data-isolation.md), and
[E2E](end-to-end-testing.md) for complementary requirements.

Before completing observable behaviour, manually confirm every affected public browser boundary and every affected API
operation. Browser confirmation uses the exact served origin; API confirmation follows the
[API-testing standard](api-testing.md). Record `Public-boundary impact: none` when neither boundary is affected.
