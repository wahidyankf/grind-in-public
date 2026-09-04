---
tldr: "Maps maintained software-quality outcomes to truthful enforcement and evidence routes."
when_to_use: "Use before completing a change or changing a gate, hook, scheduled workflow, or review obligation."
---

# Software Quality Enforcement

Apply this map to every repository change. Linked standards own detail; this document owns enforcement classification
and routing.

## Enforcement Classes

- **Required gate** is an automated command that must pass before applicable work is complete.
- **Commit gate** runs automatically and blocks `git commit` on failure.
- **Push gate** runs automatically and blocks `git push` on failure.
- **Scheduled detection** finds regressions on its cadence but does not block an earlier push.
- **Required evidence** is a mandatory human or agent review that blocks completion.
- **Runtime guard** fails closed before unsafe behaviour starts.

A row applies when a change can alter its outcome, boundary, artifact, or mechanism. Missing automation never weakens a
mandatory rule. Preserve applicable routes and evidence through compaction or handoff under
[governance continuity](../principles/governance-continuity.md).

## Enforcement Map

| Outcome                                              | Canonical rule                                         | Route                                                                                                                 |
| ---------------------------------------------------- | ------------------------------------------------------ | --------------------------------------------------------------------------------------------------------------------- |
| Typed, lint-clean source                             | [Quality gates](quality-gates.md)                      | **Commit:** staged checks. **Required/push:** affected `test:quick`.                                                  |
| Unit behaviour and 99% coverage                      | [TDD](tdd-policy.md)                                   | **Required/push:** owner unit and unit-coverage slices through `test:quick`.                                          |
| Local-boundary behaviour and 99% coverage            | [Quality gates](quality-gates.md)                      | **Required:** integration coverage. **Scheduled:** twice daily.                                                       |
| Public browser or process journeys                   | [E2E](end-to-end-testing.md)                           | **Required:** affected `test:e2e`. **Scheduled:** after integration.                                                  |
| Exact Gherkin corpus, bindings, adapters, exemptions | [BDD](behaviour-driven-development-policy.md)          | **Required/push:** behaviour coverage through `test:quick`.                                                           |
| Substantive Gherkin implementation                   | [BDD](behaviour-driven-development-policy.md)          | **Required evidence:** [one-by-one review](../workflows/gherkin-implementation-review.md) and affected runtime gates. |
| Accessible rendered UI                               | [E2E](end-to-end-testing.md)                           | **Required:** affected E2E plus accessibility and interaction evidence.                                               |
| Synchronized specs and project docs                  | [Specs](specs-policy.md)                               | **Required evidence:** semantic reconciliation. **Required/conditional push:** `badakmini-cli:test:repo`.             |
| Executable formal plans                              | [Plan gate](../workflows/plan-quality-gate.md)         | **Required evidence:** `PASS`. **Required:** repository machine checks.                                               |
| Sufficient consistent rules                          | [Rules gate](../workflows/rules-quality-gate.md)       | **Required evidence:** semantic gate and propagation. **Required/conditional push:** repository checks.               |
| Necessary reproducibly locked dependencies           | [Dependency selection](dependency-selection-policy.md) | **Required evidence:** selection review and affected gates.                                                           |

Run the narrowest applicable targets before completion. Scheduled coverage never replaces local proof. Nx projects,
hooks, and CI implement this map but do not replace its rules.
