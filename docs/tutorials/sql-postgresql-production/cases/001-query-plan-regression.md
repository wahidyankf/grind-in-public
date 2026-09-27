---
tldr: "Diagnose a sudden PostgreSQL query-plan regression after routine deployment."
when_to_use: "Use to practise evidence-first performance incident response."
---

# Case 001: Query-Plan Regression

## Functional requirements

- FR-1: Keep tenant alert search semantically identical, including stable pagination.
- FR-2: Support filters by status, created interval, and optional score.
- FR-3: Preserve write availability while diagnosis proceeds.

## Non-functional requirements

- NFR-1: Restore p99 below 300 ms from 6 s within the incident.
- NFR-2: Sustain 800 searches/s and 3,000 alert writes/s.
- NFR-3: Make mitigation reversible without data loss.
- NFR-4: Preserve tenant isolation and audit evidence.

The regression began after an application release and auto-analyse. Pool wait and CPU rose. Explain triage,
`pg_stat_statements`, plan comparison, parameter skew, statistics, indexes, mitigation, validation, and follow-up.

[Debrief](../debriefs/001-query-plan-regression.md)
