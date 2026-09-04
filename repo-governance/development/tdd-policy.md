---
tldr: "Requires evidenced red-green-refactor cycles for application and library behaviour."
when_to_use: "Use when writing a delivery checklist or implementing any behaviour change."
---

# TDD Policy

Develop every new or changed application and library behaviour with TDD, using the
[red-green-refactor workflow](../workflows/red-green-refactor.md). This includes Badak Mini, whose defects can silently
disable repository gates.

Treat test declarations and implementations as living documentation. Before interpreting or changing production code,
read the relevant tests to establish intended behaviour, constraints, and boundaries. For projects governed by
[BDD](behaviour-driven-development-policy.md), begin with canonical Gherkin before reading or changing its adapters.

## The Cycle

Each behaviour starts Gherkin-first: add or update exactly one canonical scenario before its RED test, then deliver it
as a RED -> GREEN -> REFACTOR cycle:

1. **RED** — write the test that expresses the scenario and run its narrowest Nx target. It must fail for the stated
   behavioural reason; compilation, configuration, or infrastructure failure is not RED evidence.
2. **GREEN** — write only enough production code to pass the test and its surrounding suite.
3. **REFACTOR** — improve the design without adding behaviour and keep the target green.

## One Scenario per Cycle

Every BDD cycle targets exactly one Gherkin scenario from the [specs policy](specs-policy.md). Its RED item records the
canonical feature path, scenario name, test path, Nx target, and expected behavioural failure. Never duplicate the
scenario body into task or plan records; the corpus remains its canonical source.

```text
- [ ] [AI] RED — `specs/apps/example/behaviours/greeting.feature` / "The app greets the configured name".
      Add the failing test to `apps/example/src/greeting.unit.test.ts`. Run
      `npm exec -- nx run -p example -t test:unit`; expect the configured-name assertion to fail.
```

Bundling several scenarios into one cycle hides which behaviour a failure belongs to. Long checklists are the expected
outcome and are not a reason to merge cycles.

Pure data or calculation tests that underpin several scenarios use `**Gherkin (underpins)**` and may name a list.

## Regression Tests

See [TDD Policy Details](tdd-policy/README.md).

Pure refactoring begins from a green baseline, preserves behaviour, and keeps relevant tests green throughout. Add
characterization tests before restructuring behaviour that lacks adequate coverage.

## Verification

`test:quick` and applicable broader targets run the resulting tests under the [testing policy](testing-policy.md). The
[plan quality gate](../workflows/plan-quality-gate.md) checks traceability only when explicitly requested. Task or plan
records prove ordering by preserving the expected and observed RED plus final GREEN and REFACTOR-green results;
automation proves the final state, not the historical sequence.
