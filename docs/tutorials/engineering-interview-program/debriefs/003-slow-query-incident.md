---
tldr: "Separates incident command from diagnosis and fixes a skew-driven join regression safely."
when_to_use: "Use after attempting Capstone 003."
---

# Debrief 003: Slow-Query Incident

Assign incident command, database diagnosis, application mitigation, and stakeholder communication. Freeze unrelated
changes. Bound request concurrency and retry budgets, preserve operator connections, and communicate customer impact and
next update time.

The technical lead traces pool wait to a query id, then to a 100-times underestimate caused by tenant/status
correlation. That estimate selected nested loops whose repeated B-tree probes became expensive; extended statistics and
a partial tenant/status/time index restore a bounded path. Validate exact result sets, actual/estimated rows, buffers,
spills, lock waits, and p99 under skewed load. Roll out to a canary and retain the old route.

Follow-up owners cover plan-regression tests, skewed fixtures, overload policy, statistics monitoring, and runbook. The
review asks why safeguards failed, not who wrote the query.
