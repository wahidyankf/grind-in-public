---
tldr:
  "Uses Scrum or flow practices as feedback mechanisms while making priorities, WIP, dependencies, and trade-offs
  visible."
when_to_use: "Use for process, stakeholder conflict, prioritization, and team-execution questions."
---

# Agile Flow and Stakeholders

Agile methods shorten feedback and adapt plans; ceremonies are not the objective. Choose Scrum when a stable cadence and
goal help a product team. Choose continuous flow when work arrival is unpredictable, such as operations or platform
support. Keep policies explicit in either model.

## Scrum responsibilities

- product owner orders value and clarifies outcomes;
- developers own the plan and quality of the increment;
- Scrum Master enables the framework and removes systemic impediments;
- Engineering Manager supplies people/technical context, develops capability, and handles organizational constraints.

Do not turn daily Scrum into status reporting to the manager. Use it for developers to coordinate toward the goal.

## Flow controls

```text
ready -> in progress (WIP 3) -> review (WIP 2) -> validate -> done
              ^                    |
              +-- finish before starting more
```

WIP limits expose bottlenecks and reduce context switching. Expedite has explicit entry criteria and capacity cost.
Operational work needs reserved or dynamically managed capacity rather than silently interrupting planned work.

## Stakeholder alignment

Start from shared business outcome, constraints, and decision authority. Present options:

| Option       | Outcome                           | Cost/risk                             | Evidence gate             |
| ------------ | --------------------------------- | ------------------------------------- | ------------------------- |
| Thin release | earlier feedback on core path     | defers reporting                      | adoption/SLO after cohort |
| Full scope   | complete workflow                 | later learning, more integration risk | end-to-end proof          |
| Pause        | protect incident/reliability work | delayed feature                       | error-budget recovery     |

Record the decision and changed assumptions. Avoid surprising stakeholders with invisible risk, but also avoid flooding
them with implementation detail that does not affect their choice.

## Impediments

Remove one-time blockers quickly, then fix recurring mechanisms: unclear ownership, slow environments, approval queues,
unstable tests, overloaded reviewers, or dependency contracts. The manager should not become a permanent human router.

## Retrospectives

Use data and psychological safety to find a small actionable experiment with an owner and review date. Repeated actions
without completion teach the team that reflection has no consequence.
