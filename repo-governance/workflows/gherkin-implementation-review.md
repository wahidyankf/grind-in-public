---
tldr: "Reviews every expanded Gherkin scenario and applicable adapter for substantive implementation."
when_to_use: "Use after materially changing a feature, behaviour adapter, exemption, or compliance mechanism."
---

# Gherkin Implementation Review

Run this semantic review before completing a material feature, adapter, exemption, or behaviour-compliance change.
Static binding coverage proves that a binding exists; it cannot prove that the binding implements the behaviour.

## Procedure

1. Recursively inventory canonical features and expand every Scenario Outline example.
2. Create one row per expanded scenario and applicable Unit, Integration, and E2E adapter. Record feature, scenario,
   adapter, binding/support locations, and `PASS`, `EXEMPT`, or `FAIL`.
3. Trace each non-exempt Given-When-Then path:

   ```text
   Given establishes boundary-valid state
           |
           v
   When invokes production subject or public boundary
           |
           v
   Then reads independent observable evidence
   ```

4. Mark `FAIL` when a step is empty or a no-op, returns a literal success sentinel, selects success from expected data,
   asserts unrelated generic text, copies expected data into the value asserted, or can touch non-isolated user data.
5. For an exemption, verify its scenario-level tag, immediately preceding canonical comment, genuine boundary mismatch,
   and substantive alternative target/scenario. Unit exemptions always fail. Remove no-op or success branches from an
   exempt adapter so accidental execution fails instead of reporting false proof.
6. Run affected Unit, Integration, E2E, behaviour-compliance, and repository targets. Any failed row, invalid exemption,
   missing row, or runtime failure blocks completion.

Store a requested audit under ignored `generated-reports/`. Include corpus totals, every review row, exemption
inventory, commands, findings, fixes, and results. Exempt rows remain explicit and count toward the expected row total.

If a row cannot pass, repair the production seam or adapter. Use an exemption only when the layer fundamentally cannot
express the scenario and another named layer proves the omitted concern; never use one to finish faster or quarantine a
flake.
