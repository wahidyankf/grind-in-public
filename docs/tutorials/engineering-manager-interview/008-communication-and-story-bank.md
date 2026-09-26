---
tldr: "Builds concise leadership stories with context, personal decisions, trade-offs, outcomes, and learning."
when_to_use: "Use to prepare behavioural answers and communicate sophisticated systems briefly."
---

# Communication and Story Bank

Use enough context for the decision to make sense, then spend most time on your reasoning and outcome. A useful shape is
`Context -> Goal -> Actions -> Result -> Learning`, with trade-offs inside Actions.

## Two-minute structure

```text
0:00-0:20  context, scale, stakes
0:20-0:35  your responsibility and constraint
0:35-1:25  two or three decisions, options, people mechanism
1:25-1:50  measured outcome and remaining cost
1:50-2:00  learning applied later
```

Pause for deeper questions. Do not narrate a ten-minute chronology before stating the result.

## Story inventory

Prepare distinct examples for:

- architecture or migration with a rejected alternative;
- production incident and systemic prevention;
- missed commitment or changed plan;
- disagreement with a senior engineer or stakeholder;
- difficult feedback and sustained improvement;
- underperformance handled fairly;
- hiring or team-shape decision;
- quality bar raised through a mechanism;
- delegation that grew an engineer;
- decision you got wrong and corrected;
- ambiguous strategy turned into execution;
- resource constraint or priority trade-off.

## Evidence sheet

For each story, record:

| Field       | Question                                                |
| ----------- | ------------------------------------------------------- |
| Baseline    | What was happening and how do you know?                 |
| Stake       | Who was harmed or what outcome was blocked?             |
| Choice      | Which alternatives existed and why this one?            |
| Your action | What did you personally decide, communicate, or change? |
| Team action | What did others own and how were they enabled?          |
| Result      | Which metric, behaviour, or risk changed?               |
| Cost        | What trade-off or unfinished risk remained?             |
| Learning    | What changed in later practice?                         |

## Explaining technical complexity

Lead with user or business impact, use one simple model, define one or two critical terms, and name the trade-off. For
example: "We moved alert generation out of the monolith to scale it independently. The hard part was not creating a
service; it was transferring data-write authority without losing or duplicating alerts. We used an outbox, shadow
comparison, and a fenced cutover watermark, accepting temporary dual operating cost."

## Final review

Remove confidential names, blame, invented numbers, and jargon that does no explanatory work. Make sure each answer
reveals judgment, collaboration, and learning—not only a successful ending.
