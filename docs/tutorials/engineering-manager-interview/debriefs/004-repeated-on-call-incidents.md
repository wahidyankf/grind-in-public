---
tldr:
  "Stabilize retry and pool behaviour, shift priority through error-budget evidence, and distribute operating
  capability."
when_to_use: "Use after Scenario 004 to assess reliability and team-health leadership."
---

# Debrief 004: Repeated On-Call Incidents

Immediately cap retries, add jitter/deadlines, bound dependency concurrency, protect database pools, shed optional work,
and verify alerts/runbooks. If risk is active, freeze or slow feature rollout using error-budget and incident evidence.

Run one cross-incident review to find the recurring mechanism. Fund permanent work: propagated deadlines, retry budgets,
bulkheads, capacity tests, SLO burn alerts, and a dependency-failure exercise. Track pages, user-impact minutes,
database saturation, retry ratio, time to mitigate, and recurrence.

Distribute knowledge with paired on-call, shadow shifts, runbook drills, architecture walkthroughs, and fair rotation.
Give the two overloaded engineers recovery time and stop treating their heroics as the system.

Communicate the feature trade-off as protected customer availability and reduced interruption cost, with a clear exit
criterion. Weak answers add another dashboard, ask engineers to be careful, or rotate unprepared people into incidents.
