---
tldr:
  "Prepares teams for incidents, separates command from diagnosis, communicates clearly, and turns learning into
  prevention."
when_to_use: "Use for incident, on-call, reliability prioritization, and high-pressure leadership questions."
---

# Reliability and Incident Leadership

During an incident, optimize for reducing harm while preserving enough evidence to understand the system. The most
senior engineer need not execute every action.

## Roles

```text
incident commander -- priorities, decisions, role clarity
      +-- operations lead -- mitigation execution
      +-- investigation leads -- hypotheses and evidence
      +-- communications -- stakeholder/customer cadence
      `-- scribe -- timeline, decisions, actions
```

The manager often commands or removes organizational impediments. Avoid debugging and commanding simultaneously when the
incident is complex.

## Response loop

```text
detect -> assess impact -> stabilize -> diagnose -> recover -> verify -> learn
             ^              |
             +-- communicate+
```

Prefer reversible mitigations: shed optional load, roll back, disable a cohort, fail over according to runbook, or
degrade to a safe policy. Record hypothesis, evidence, decision, owner, and result. Avoid stacking unobserved changes.

## Communication

An update states current impact, start/detection time, what changed, mitigation status, uncertainty, and next update
time. Do not promise recovery times without evidence. Internal technical detail and external impact communication have
different audiences but one factual timeline.

## Post-incident learning

Reconstruct contributing conditions across design, tests, rollout, capacity, observability, process, and organization.
Avoid stopping at "human error." Prioritize actions by recurrence/impact reduction and verify completion/effectiveness.

```text
weak:  engineer made mistake -> remind engineer
strong: unsafe action was possible and silent -> guardrail + review + detection + recovery exercise
```

## Reliability planning

Use error budgets and incident data to choose reliability work. Repeated SLO burn should change release or capacity
policy. Rotate on-call fairly, provide training and shadow shifts, keep runbooks executable, and compensate according to
the organization. Chronic paging is a system defect, not a rite of passage.
