---
tldr: "Defines test boundaries and the required Nx quality-gate contracts."
when_to_use: "Use when classifying tests or adding, changing, running, or reviewing an Nx quality target."
---

# Quality Gates

Define project gates as documented Nx targets invoked through the workspace package manager. The
[enforcement map](software-quality-enforcement.md) classifies their blocking, scheduled, runtime, and evidence routes.

## Test Boundaries

- **Unit** tests run in-process and replace filesystem, database, environment, clock, randomness, child-process,
  network, and other OS-facing dependencies with injected doubles. Setup and assertions must not access those real
  resources.
- **Integration** tests may use isolated local filesystem, environment state, and same-machine child processes. They
  must never use network communication, including loopback, `localhost`, or a local server. Isolate and clean every
  resource deterministically.
- **End-to-end** tests exercise a public application boundary and may use OS resources, processes, browsers, and network
  communication required by the journey. External services need explicit authorization.

Classify a test by the strongest real boundary touched by setup, subject, or assertions. E2E is defined by
public-boundary observation, not merely by permission to use more resources. Keep executable unit, integration, and E2E
tests separate. Every application with a public browser or process boundary owns a dedicated Nx E2E project.

## Gate Contracts

Owner applications expose `build`, `typecheck`, `lint`, `test:unit`, `test:integration`, `test:coverage:unit`,
`test:coverage:integration`, `test:coverage:behaviour:<owner>`, `test:coverage:behaviour`, `test:coverage`, and
`test:quick`.

Dedicated E2E projects expose only `typecheck`, `lint`, `test:coverage:behaviour:e2e`, `test:coverage:behaviour`,
`test:e2e`, and `test:quick`, plus justified operational targets such as browser installation. They have no unit,
integration, numeric-coverage, corpus, or placeholder targets.

- `test:coverage:behaviour` statically proves corpus, adapter, step-binding, and exemption completeness. Semantic
  implementation remains a separate evidence gate.
- Owner `test:coverage:behaviour` composes owner and E2E slices. The E2E generic target delegates to that aggregate
  without creating a project-graph cycle.
- Owner `test:quick` runs `typecheck` -> `lint` -> `test:unit` -> unit coverage -> static behaviour coverage.
- E2E `test:quick` runs `typecheck` -> `lint` -> its static behaviour slice.
- Integration and E2E runtime targets are uncached and remain outside Git hooks. Numeric coverage stays at least 99%.
- Aggregate gates compose named Nx targets; do not duplicate their underlying tool commands.
- Do not invent an inapplicable target for symmetry. Explain legitimate omissions in the project README.

Scheduled CI runs every integration-coverage target before the owner's complete unfiltered E2E target. Fix failures at
their responsible cause; never weaken or bypass a gate to obtain a pass. A deliberate RED is temporary process evidence,
not a completed state. Run `test:quick` after the final [red-green-refactor](../workflows/red-green-refactor.md) cycle.
