---
tldr: "Defend a deliberate modular-monolith boundary with evidence, internal contracts, and future extraction triggers."
when_to_use: "Use after Scenario 012 to assess strategy and executive communication."
---

# Debrief 012: Choosing What Remains in the Monolith

Present the outcome, not an ideological defence: extracted capabilities receive independent scaling/ownership where it
matters; case/profile modules retain simple transactions, local latency, and one team. Show measured scaling similarity,
change coupling, transaction boundaries, and projected network/operating cost of splitting.

Strengthen the retained target: explicit module APIs, prevented cross-module table access, owned schemas or access
layers, contract tests, observability, capacity, and independently understandable domain models.

Record future triggers: sustained independent load, distinct security/residency need, release contention, team ownership
split, or transaction boundary change. Review them periodically.

```text
completed target = justified services + intentional modular monolith + no ambiguous authority
```

Communicate that service count is an implementation metric, while deployment lead time, incidents, SLO, change coupling,
cost, and team autonomy are outcomes. Weak answers resist all evolution or promise eventual extraction merely to satisfy
the original slogan.
