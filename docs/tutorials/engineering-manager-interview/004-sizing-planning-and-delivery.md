---
tldr: "Plans uncertain work through outcomes, thin slices, dependency/risk discovery, ranges, and evidence gates."
when_to_use: "Use for estimation, missed deadlines, roadmaps, migrations, and cross-team delivery questions."
---

# Sizing, Planning, and Delivery

Estimation is decision support under uncertainty. It should expose scope, dependencies, risk, and confidence—not create
false certainty.

## Decompose by outcome

```text
desired outcome
  +-- walking skeleton: smallest end-to-end proof
  +-- capability increments with user/operational value
  +-- risk spikes: unknown performance, data, or integration
  `-- rollout, recovery, and decommission work
```

Avoid component-only plans such as "build API, build database, build UI" when none produces a usable or testable flow.

## Forecasting

Use historical throughput/cycle time when work is comparable, or bottom-up ranges when it is not. State assumptions and
give a confidence interval. Reforecast as evidence arrives.

For migrations, plan by evidence states rather than percent complete:

```text
seam ready -> shadow equivalent -> canary SLO -> authority moved -> legacy retired
```

"80% migrated" hides the riskiest 20%: data authority, rollback, and decommissioning.

## Risk register

| Field        | Example                                        |
| ------------ | ---------------------------------------------- |
| Risk         | Shared database cannot sustain shadow reads.   |
| Signal       | Read IOPS or replica lag crosses threshold.    |
| Mitigation   | Snapshot/object backfill and bounded CDC.      |
| Contingency  | Pause cohort expansion; keep legacy authority. |
| Owner/review | Named owner and decision checkpoint.           |

Dependencies need an owner, contract, due condition, and fallback. "Another team is working on it" is not a plan.

## When a commitment is at risk

Surface it early with current evidence. Offer choices and consequences: cut a lower-value requirement, sequence a
smaller release, add capacity to the actual bottleneck, change the date, or accept a named risk. Adding people late may
increase coordination and onboarding cost.

## Delivery health

Track outcome completion, cycle time, work age, blocked time, rework, scope change, forecast confidence, and operational
quality. Avoid using velocity to compare teams or evaluate individuals; it is local planning data and easily gamed.
