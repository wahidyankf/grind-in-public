---
tldr: "Reviews every expanded Gherkin scenario and applicable adapter for substantive, isolated implementation."
when_to_use: "Use after materially changing a feature, adapter, exemption, or behaviour-compliance mechanism."
---

# Gherkin Implementation Review

Run this semantic review before completing a material feature, adapter, exemption, or behaviour-compliance change.
Static binding coverage proves that a binding exists; it cannot prove that the binding implements the behaviour.

## Inputs

- the recursively discovered canonical `.feature` corpus;
- every required Unit, Integration, and E2E adapter;
- the [BDD policy](../development/behaviour-driven-development-policy.md),
  [test boundaries](../development/quality-gates.md#test-boundaries),
  [test-data iron rule](../development/test-data-isolation.md#iron-rule), and applicable Nx targets; and
- the changed scope, or the complete corpus when a full audit is requested.

## Procedure

Use an agent to perform the review. Never replace one-by-one inspection with counts, grep-only heuristics, or a green
test run.

1. Recursively inventory canonical features and expand every Scenario Outline example.
2. Create one row per expanded scenario and required Unit, Integration, and E2E adapter. Record feature, scenario,
   adapter, binding/support locations, and `PASS`, `EXEMPT`, or `FAIL`.
3. Trace each non-exempt Given-When-Then path:

   ```text
   Given establishes boundary-valid synthetic state
           |
           v
   When invokes the production subject or public boundary
           |
           v
   Then reads independent observable evidence
   ```

4. Mark `FAIL` when a step is empty or a no-op; returns or stores literal success; selects success from expected data;
   asserts unrelated generic text; copies expected data into the value asserted; simulates deployment against an
   unchanged server; or can touch production data. A helper action or embedded assertion does not make a later literal
   `true` independent evidence.
5. For each exemption, verify its scenario-level tag, immediately preceding canonical comment, genuine boundary
   mismatch, and substantive alternative target/scenario. Unit exemptions fail. Remove no-op or success branches from
   exempt adapters so accidental execution fails rather than reporting false proof.
6. Verify fixtures, identities, roots, processes, sessions, ports, browser contexts, and cleanup are synthetic,
   isolated, marked, and fail closed before the subject starts.
7. Run affected Unit, Integration, E2E, behaviour-compliance, and repository targets. Any failed row, invalid exemption,
   missing row, unresolved partial assessment, cleanup failure, or runtime failure blocks completion.

## Evidence and Recovery

Store a requested audit under ignored `generated-reports/`. Include corpus totals, every review row, exemption
inventory, commands, findings, fixes, and results. The row count equals expanded scenarios multiplied by required
adapters; exemptions remain explicit `EXEMPT` rows and are never subtracted.

If a row cannot pass, repair the production seam or adapter. Use an exemption only when the layer fundamentally cannot
express the scenario and another named layer proves the omitted concern; never use one to finish faster, quarantine a
flake, or defer implementation. Review the final diff for placeholder patterns after repairs.
