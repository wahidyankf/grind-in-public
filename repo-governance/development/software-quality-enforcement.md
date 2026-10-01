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

- **Typed, lint-clean source**
  - Rule: [Quality gates](quality-gates.md)
  - Enforcement or evidence route: **Commit:** staged checks. **Required/push:** affected `test:quick`.
- **Test-first behaviour delivery**
  - Rule: [TDD](tdd-policy.md)
  - Enforcement or evidence route: **Evidence:** task records RED, GREEN, and REFACTOR-green; automation proves final
    state.
- **Unit behaviour and 99% coverage**
  - Rule: [Quality gates](quality-gates.md)
  - Enforcement or evidence route: **Required/push:** unit and unit coverage through owner `test:quick`.
- **Local-boundary behaviour and 99% coverage**
  - Rule: [Quality gates](quality-gates.md)
  - Enforcement or evidence route: **Required:** integration coverage. **Scheduled:** twice daily.
- **Public browser, process, and API journeys**
  - Rule: [E2E](end-to-end-testing.md)
  - Enforcement or evidence route: **Required:** affected E2E. **Evidence:** affected APIs use `curl`. **Scheduled:**
    after integration.
- **Exact corpus, adapters, and exemptions**
  - Rule: [BDD](behaviour-driven-development-policy.md)
  - Enforcement or evidence route: **Required/push:** static behaviour coverage through `test:quick`.
- **Substantive Gherkin implementation**
  - Rule: [BDD](behaviour-driven-development-policy.md)
  - Enforcement or evidence route: **Evidence:** [one-by-one review](../workflows/gherkin-implementation-review.md) and
    runtime gates.
- **Synthetic isolated test state**
  - Rule: [Test data](test-data-isolation.md)
  - Enforcement or evidence route: **Runtime:** fail-closed boundaries. **Required:** policy tests and cleanup.
- **Accessible and usable rendered UI**
  - Rule: [E2E](end-to-end-testing.md)
  - Enforcement or evidence route: **Required:** affected E2E. **Evidence:** exact-origin and exploratory/usability
    review.
- **Synchronized specs and project docs**
  - Rule: [Specs](specs-policy.md)
  - Enforcement or evidence route: **Evidence:** semantic reconciliation. **Required/conditional push:** `test:repo`.
- **Executable formal plans**
  - Rule: [Plan gate](../workflows/plan-quality-gate.md)
  - Enforcement or evidence route: **Evidence:** explicitly requested `PASS`. **Required:** repository checks.
- **Sufficient consistent rules**
  - Rule: [Rules gate](../workflows/rules-quality-gate.md)
  - Enforcement or evidence route: **Evidence:** automatic propagation; explicit audits hand off every non-pass.
- **Leak-free outbound history**
  - Rule: [Push leak review](../workflows/push-leak-review.md)
  - Enforcement or evidence route: **Evidence:** review of each outgoing commit. **Push:** commit-by-commit screen.
    **Detection:** hosted replay of each push to `main`.
- **Necessary reproducibly locked dependencies**
  - Rule: [Dependency selection](dependency-selection-policy.md)
  - Enforcement or evidence route: **Evidence:** selection review and affected gates.

Nx projects, hooks, and CI implement this map but do not replace its rules. Project READMEs own resolved commands and
legitimate omissions.
