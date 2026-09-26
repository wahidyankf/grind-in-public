---
tldr: "Design deduplication, correlation, ranking, fairness, and operator feedback for high-volume alerts."
when_to_use: "Use to practise union-find, heaps, queues, ranking, and human-capacity constraints."
---

# Case 015: Alert Correlation and Prioritization

Many detectors emit overlapping alerts. Design a service that groups related alerts, prioritizes work, and prevents one
noisy rule or tenant from overwhelming investigators.

### Functional requirements

| ID   | Requirement                                                                                   |
| ---- | --------------------------------------------------------------------------------------------- |
| FR-1 | Ingest alerts with subject, event, rule, score, time, evidence, and version metadata.         |
| FR-2 | Deduplicate exact retries and correlate related alerts into an evolving cluster.              |
| FR-3 | Rank clusters using severity, confidence, value, recency, relationships, and tenant policy.   |
| FR-4 | Allocate prioritized clusters to skilled investigator queues with fairness and SLA deadlines. |
| FR-5 | Capture disposition feedback and explain ranking changes.                                     |

### Non-functional requirements

| ID    | Requirement                                                                    |
| ----- | ------------------------------------------------------------------------------ |
| NFR-1 | Ingest 50,000 alerts/s and make urgent work visible within 10 seconds.         |
| NFR-2 | Rank updates are deterministic for a declared feature and policy version.      |
| NFR-3 | Bound cluster growth, re-correlation cost, and queue starvation.               |
| NFR-4 | Survive replay without duplicate cases or repeated assignments.                |
| NFR-5 | Isolate tenant queues, configuration, metrics, and investigator authorization. |

## Assumptions and exclusions

Case workflow after assignment is Case 016. Feedback is not automatically accepted as a clean training label.

## Interview prompts

1. Define alert, cluster, case, and their stable identities.
2. Compare deterministic keys, time-window joins, union-find, and graph correlation.
3. How do top-k heaps and weighted fair queues avoid starvation?
4. How are late alerts, cluster merge/split, replay, and ranking rollout handled?

Solve before reading [the worked solution](../solutions/015-alert-correlation-and-prioritization.md).
