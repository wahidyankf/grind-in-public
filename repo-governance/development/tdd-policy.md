---
tldr: "Requires evidenced red-green-refactor cycles for application and library behaviour."
when_to_use: "Use when writing a delivery checklist or implementing any behaviour change."
---

# TDD Policy

Develop every new or changed application and library behaviour with TDD, using the
[red-green-refactor workflow](../workflows/quality/red-green-refactor.md). This includes the validation scripts under
`scripts/`, whose defects can silently disable repository gates.

Treat test declarations and implementations as living documentation. Before interpreting or changing production code,
read the relevant tests to establish intended behaviour, constraints, and boundaries. For projects governed by
[BDD](behaviour-driven-development-policy.md), begin with canonical Gherkin before reading or changing its adapters.

## The Cycle

Each behaviour starts Gherkin-first: add or update exactly one canonical scenario before its RED test, then deliver it
as a RED -> GREEN -> REFACTOR cycle:

1. **RED** — write the test that expresses the scenario and run only the check covering the current item (its narrowest
   Nx target). Do not run the block to prove a single item. It must fail for the stated behavioural reason; compilation,
   configuration, or infrastructure failure is not RED evidence.
2. **GREEN** — write only enough production code to pass the test. Demonstrate GREEN by running only that same covering
   check, not the block.
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

## Batched Cycles in Planned Delivery

When a plan already delivers work in batches (one pull request or one delivery boundary at a time), the cycle may run
per batch instead of per scenario, so planned work does not pay a full red-green ceremony for every scenario:

1. Write the batch's tests first, from its scenarios, before any production change in the batch.
2. Run one focused RED run for the batch. Every scenario's test must fail for its stated reason; group the reasons in
   one batch record, not one record per scenario. A test that passes is repaired before implementation starts.
3. Implement the batch, then run the identical selection to GREEN. Keep the selection frozen between RED and GREEN.
4. Refactor inside the batch with the selection green; a behaviour change noticed along the way joins the batch's tests
   or becomes its own cycle.
5. Run the full applicable regression and touched end-to-end flows only at the phase or push boundary, per the
   repository's focused-runs and regression-gate rules.
6. Do not repeat an accepted RED whose tests and production code are unchanged.

A batch is one cohesive behaviour slice sharing a boundary and fixtures. Unrelated failures in one RED run mean the
batch is too large; split it. Small single-scenario changes and single bug fixes keep the per-scenario cycle. This
section changes run and evidence granularity only: it never permits writing tests after the code, and the existing scope
exemptions are unchanged.

## Regression Tests

See [TDD Policy Details](tdd-policy/README.md).

Pure refactoring begins from a green baseline, preserves behaviour, and keeps relevant tests green throughout. Add
characterization tests before restructuring behaviour that lacks adequate coverage.

## Verification

`test:quick` and applicable broader targets run the resulting tests under the [testing policy](testing-policy.md). The
[plan quality gate](../workflows/quality/plan-quality-gate.md) checks traceability only when explicitly requested. Task
or plan records prove ordering by preserving the expected and observed RED plus final GREEN and REFACTOR-green results;
automation proves the final state, not the historical sequence.
