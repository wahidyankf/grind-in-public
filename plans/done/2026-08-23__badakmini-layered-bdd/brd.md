# Business Requirements

## Why This Matters

Badak Mini currently has strong Go tests, but `specs/` is empty and no automated check proves that declared behaviour
stays executable across unit, integration, and public process boundaries. That leaves behaviour discussions detached
from the test suite and allows a feature-file change to escape affected-project selection.

The owner wants the testing discipline already proven in BeaverNest: one recursive Gherkin corpus, thin bindings,
interchangeable drivers, static completeness checks, fast local feedback, and slower real-boundary verification outside
hooks.

## Success

- Every Badak Mini acceptance scenario is executable through unit, local-only integration, and process E2E adapters.
- Editing the canonical corpus invalidates Nx caches and fails when structure, bindings, drivers, adapters, or behaviour
  no longer agree.
- Unit and integration coverage slices each enforce at least 99% aggregate Go statement coverage from their named
  profiles, the native line-equivalent metric used by `go tool cover`.
- Future applications and libraries receive the same role-based requirements from one canonical governance policy.
- Full integration and E2E verification runs daily at 06:00 WIB and on manual dispatch.

## Non-Goals

- Porting BeaverNest's F#, Elixir, TickSpec, or ExBdd implementation.
- Adding networked test resources, a hosted test service, or a general-purpose BDD framework.
- Changing Badak Mini commands, output, exit codes, or governance behaviour.
- Rejecting a compatible specification edit merely because its bytes changed.

## Risks

- A generic test framework could become more expensive than the small CLI it supports. The design therefore exposes only
  the bindings and driver operations needed by the canonical scenarios.
- Three adapters can drift. Static behaviour coverage compares all adapters before runtime E2E begins.
- High numeric coverage can reward exclusions. Each slice documents its instrumented boundary and uses no broad
  production-code exclusion.
