---
tldr: "Defines public-boundary E2E ownership, isolation, accessibility, and execution requirements."
when_to_use: "Use when creating, changing, selecting, or reviewing browser or process E2E journeys."
---

# End-to-End Testing

Every application with a public browser or process boundary uses a dedicated Nx E2E project. The harness consumes the
owner's canonical Gherkin corpus and observes only its built public boundary; it owns no separate behaviour corpus.

## Required Behaviour

- Run the complete applicable, non-exempt corpus. Missing or ambiguous applicable steps fail generation or execution.
- Build the owner through its Nx target. Do not import private production modules to substitute for the public journey.
- Use isolated synthetic state and deterministic cleanup. A cleanup failure fails the run.
- Keep unconditional skips, focus markers, expected-failure baselines, and silent missing-step fallbacks out of the
  executable suite.
- Use exact local origins and fail when an unexpected existing process answers the test port in CI.
- Keep runtime E2E uncached and outside commit and push hooks. Run affected journeys before completion and the complete
  suite after integration coverage in scheduled CI.

Browser E2E covers representative desktop, tablet, and mobile viewports when layout changes with width. It verifies
semantic roles, keyboard reachability, focus behaviour, accessible names, and automated WCAG findings where applicable.
Assertions observe meaningful state, URL, response, file, or DOM evidence produced by the action; generic page text is
not proof of a specific journey.

Exploratory review separates spec-aware verification from usability review so the written scenario does not hide a
confusing interaction. Record requested reports under ignored `generated-reports/`; they are evidence, not authority.
