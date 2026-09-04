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
mandatory rule. Run the narrowest applicable target and record evidence before completion; scheduled detection never
replaces local proof. Preserve applicable routes and evidence through compaction or handoff under
[governance continuity](../principles/governance-continuity.md).

## Enforcement Map

| Outcome                                    | Rule                                                   | Enforcement or evidence route                                                                         |
| ------------------------------------------ | ------------------------------------------------------ | ----------------------------------------------------------------------------------------------------- |
| Typed, lint-clean source                   | [Quality gates](quality-gates.md)                      | **Commit:** staged checks. **Required/push:** affected `test:quick`.                                  |
| Test-first behaviour delivery              | [TDD](tdd-policy.md)                                   | **Evidence:** task records RED, GREEN, and REFACTOR-green; automation proves final state.             |
| Unit behaviour and 99% coverage            | [Quality gates](quality-gates.md)                      | **Required/push:** unit and unit coverage through owner `test:quick`.                                 |
| Local-boundary behaviour and 99% coverage  | [Quality gates](quality-gates.md)                      | **Required:** integration coverage. **Scheduled:** twice daily.                                       |
| Public browser, process, and API journeys  | [E2E](end-to-end-testing.md)                           | **Required:** affected E2E. **Evidence:** affected APIs use `curl`. **Scheduled:** after integration. |
| Exact corpus, adapters, and exemptions     | [BDD](behaviour-driven-development-policy.md)          | **Required/push:** static behaviour coverage through `test:quick`.                                    |
| Substantive Gherkin implementation         | [BDD](behaviour-driven-development-policy.md)          | **Evidence:** [one-by-one review](../workflows/gherkin-implementation-review.md) and runtime gates.   |
| Synthetic isolated test state              | [Test data](test-data-isolation.md)                    | **Runtime:** fail-closed boundaries. **Required:** policy tests and cleanup.                          |
| Accessible and usable rendered UI          | [E2E](end-to-end-testing.md)                           | **Required:** affected E2E. **Evidence:** exact-origin and exploratory/usability review.              |
| Synchronized specs and project docs        | [Specs](specs-policy.md)                               | **Evidence:** semantic reconciliation. **Required/conditional push:** `test:repo`.                    |
| Executable formal plans                    | [Plan gate](../workflows/plan-quality-gate.md)         | **Evidence:** explicitly requested `PASS`. **Required:** repository checks.                           |
| Sufficient consistent rules                | [Rules gate](../workflows/rules-quality-gate.md)       | **Evidence:** automatic propagation; explicit audits hand off every non-pass.                         |
| Necessary reproducibly locked dependencies | [Dependency selection](dependency-selection-policy.md) | **Evidence:** selection review and affected gates.                                                    |

Nx projects, hooks, and CI implement this map but do not replace its rules. Project READMEs own resolved commands and
legitimate omissions.
