---
tldr: "Calculates selectivity and detects a broken independence assumption."
when_to_use: "Use after planner statistics to practise estimate-first diagnosis."
---

# Lab 005: Selectivity Estimation

A table has 10 million rows. Tenant 7 owns 40%; status `open` is 5% globally but 60% for tenant 7. Calculate the
independence estimate and actual rows for both predicates. Explain the plan risk and propose statistics or schema fixes.

Compare with the [solution](../lab-solutions/005-selectivity-estimation.md).
