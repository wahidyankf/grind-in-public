---
tldr:
  "Resolve architecture conflict by resetting behaviour, agreeing criteria, assigning decision ownership, and
  time-boxing evidence."
when_to_use: "Use after Scenario 001 to assess conflict and technical-leadership reasoning."
---

# Debrief 001: Senior-Engineer Architecture Disagreement

A strong answer first stops personal behaviour and meets both engineers separately, then together around a shared
problem statement. Clarify decision authority and a short deadline so debate does not become indefinite.

Require baseline evidence: change coupling, deployment/incidents, load skew, data transactions, ownership, latency, and
the exact pressure a split solves. Compare three options: strengthen modules, extract one measured seam, or broad split.
Score independent scaling, release, security, data ownership, team operation, and migration risk.

```text
shared requirements -> evidence spike -> written alternatives -> owner decides -> dissent recorded -> commit -> review
```

The manager may own the process while a senior engineer owns the technical decision. For a high-blast-radius data
migration, the manager ensures cross-team review and accountable risk acceptance. Preserve dignity, explicitly reject
dismissive conduct, and recognize useful contributions from both positions.

Set a revisit trigger: load threshold, deployment coupling, or team boundary. Weak signals are choosing by title,
compromising into an incoherent half-architecture, or personally rewriting the design to end conflict.
