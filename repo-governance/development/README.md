---
tldr: "Indexes development policies for code, testing, hooks, Nx, and validation."
when_to_use: "Use when changing executable code, development tooling, tests, or quality gates."
---

# Development Governance

This directory contains rules for building, testing, validating, and maintaining the repository's executable code. Read
only the policy that matches the work at hand:

- [Architecture Specifications](architecture-specifications.md) for each application's canonical as-built C4 model and
  its maintenance.
- [API Testing](api-testing.md) for automated contracts and manual public-boundary proof when an API changes.
- [Badak Mini](badakmini-cli-policy.md) for repository-local validation checks.
- [Behaviour-Driven Development](behaviour-driven-development-policy.md) for mandatory canonical corpus and adapter
  roles.
- [Code Commentary](code-commentary-policy.md) for learning-oriented comments.
- [Code Style](code-style-policy.md) for the language target, naming, indentation, and import style.
- [Commit Hooks](commit-hook-policy.md) for required Git-hook behaviour.
- [Dependency Selection](dependency-selection-policy.md) for when an external dependency may be added instead of
  standard-library or existing code.
- [Deployment](deployment-policy.md) for the delivery target, the `prod-` promotion branch, and domain cutover
  authorization.
- [End-to-End Testing](end-to-end-testing.md) for dedicated public-boundary harnesses and browser/process proof.
- [Harness Pre-Edit Triggers](harness-pre-edit-triggers.md) for what each harness wires before an edit, and how far it
  is verified.
- [Nx Workspace](nx-workspace-policy.md) for raw-Nx boundaries and verification.
- [Quality Gates](quality-gates.md) for test boundaries and owner/E2E target contracts.
- [Rule Change Triggers](rule-change-trigger-policy.md) for how a rule change announces the workflows that must follow
  it.
- [Resource-Aware Development](resource-aware-development.md) for checksum-pinned admission and recovery of
  compute-bearing Nx work.
- [Specs](specs-policy.md) for Gherkin acceptance criteria and the `specs/` tree; its
  [detail index](specs-policy/README.md) holds focused structural guidance.
- [Software Quality Enforcement](software-quality-enforcement.md) for truthful gate, hook, schedule, runtime, and
  evidence routing.
- [TDD](tdd-policy.md) for red-green-refactor cycles bound to scenarios.
- [TDD Policy Details](tdd-policy/README.md) for focused TDD requirements.
- [Testing](testing-policy.md) for quick and integration-test responsibilities.
- [Test Data Isolation](test-data-isolation.md) for synthetic identities, per-run boundaries, and cleanup.
- [Testing Policy Details](testing-policy/README.md) for focused testing requirements.
- [Workspace Commands](workspace-commands.md) for the canonical command, check, and hook reference.

Foundational principles remain in [`../principles/`](../principles/README.md), and repeatable procedures remain in
[`../workflows/`](../workflows/README.md).
