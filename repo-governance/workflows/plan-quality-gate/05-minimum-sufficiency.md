---
tldr: "Makes plan quality review retain only artifacts needed to execute a safe, correct scope."
when_to_use: "Use when plan-checker or plan-fixer assesses redundant plan documents, assets, evidence, or enforcement."
---

# Plan Minimum Sufficiency

`plan-checker` applies [Minimum Sufficiency](../../principles/minimum-sufficiency.md): every plan document, asset, evidence item, and enforcement mechanism must be required for scope, safety, correctness, or execution. Redundant, orphaned, and needless fragments are findings.

`plan-fixer` may remove only a redundant or orphaned non-core plan artifact when checker evidence shows it has no unique required content. It does not delete a core document, a delivery step, or any requirement to make a finding disappear.

Every `plan-checker` and `plan-fixer` prompt states its applicable rule in the imperative. Change the three harness copies with this document so the workflow and prompts keep the same mandate.
