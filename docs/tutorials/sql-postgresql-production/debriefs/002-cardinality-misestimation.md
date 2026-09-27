---
tldr: "Uses representative data, extended statistics, plan sensitivity, and workload isolation."
when_to_use: "Use after attempting Case 002."
---

# Debrief 002: Cardinality Misestimation

Begin at the first node where actual rows diverge from estimated rows, not the final slow node. Create dependency/MCV
statistics across tenant, status, and relevant keys; raise targets selectively and analyse. Compare custom plans for
small and huge tenants against a generic prepared plan. Materialize a 60-second read model or route exports to an
isolated replica. If the tenant still dominates CPU, I/O, vacuum, and recovery, place it in a dedicated cell through the
tenant directory rather than maintaining bespoke SQL.
