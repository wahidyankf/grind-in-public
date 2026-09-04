---
tldr: "Requires Gherkin behaviour and current application architecture in specs/, separate from implementation code."
when_to_use:
  "Use when adding or changing an app or library's behaviour or architecture, or when writing a plan's prd.md."
---

# Specs Policy

## Scope

This policy covers `specs/`, the repository's description of what its software should do and, for applications, its
current architectural boundaries. Specs state intent, observable behaviour, and as-built architecture; `apps/` and
`libs/` contain the implementation that delivers them. Keeping them apart means a behaviour or boundary discussion can
start from one source of truth instead of from implementation detail.

## Structure

Read [Specs Structure](specs-policy/structure.md) for the mirrored tree, the mandatory application C4 model, and when a
detail folder is warranted.

## Gherkin

Acceptance criteria are written as Gherkin scenarios in `.feature` files under `behaviours/`.

A scenario uses exactly one primary `Given`, one `When`, and one `Then`; further steps chain with `And` or `But`. A
`Background` block and a `Scenario Outline` `Examples` table are exempt. Two `When` steps in one scenario describe two
behaviours, so split them.

```gherkin
Feature: Greeting

  Scenario: The app greets the configured name
    Given the app is configured with the name "Wahidyan"
    When the app runs
    Then the output is "Hello, Wahidyan!"
```

## Binding to Tests

A scenario is not documentation: it binds to a test that fails when the behaviour breaks, as the
[TDD policy](tdd-policy.md) requires. A scenario with no test behind it is worse than no scenario, because it claims
coverage that does not exist.

## When Specs Are Required

A plan that adds or changes behaviour in `apps/` or `libs/` selects the affected durable Gherkin in
`tech-docs/specification-changes.md`, writes those scenarios into its `prd.md`, and lands them in `specs/` as part of
delivery. A plan may retain rollout or operational acceptance outcomes only when that document labels them plan-only,
gives a reason, and names the delivery proof. Architectural changes similarly name the affected C4 view and update
`specs/` only to the final as-built state. Every existing app and library has executable Gherkin under its mirrored
`behaviours/` directory; a missing corpus is a policy failure, not deferred retrofit work. Drills are exempt: a drill is
practice, not repository behaviour.

## Verification

Each subject's behaviour gate recursively discovers its corpus, rejects malformed features, and fails for undefined,
ambiguous, or unused bindings. It also proves every required adapter consumes the same catalog, so adding, editing,
renaming, nesting, or deleting a feature or binding cannot silently bypass verification. The
[plan quality gate](../workflows/plan-quality-gate.md) additionally verifies that a plan's Gherkin follows the
cardinality rule and every planned scenario has a RED step.
