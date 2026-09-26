---
tldr: "Design continuously updated entity risk scores from events, relationships, and versioned policy."
when_to_use: "Use to practise materialized state, event-time updates, recomputation, and score explainability."
---

# Case 010: Continuous Customer Risk Scoring

Design a platform that maintains a current risk band and score for each customer as profile, transaction, list-match,
case, and relationship signals change.

### Functional requirements

| ID   | Requirement                                                                                |
| ---- | ------------------------------------------------------------------------------------------ |
| FR-1 | Consume ordered customer-affecting events and update derived features and score.           |
| FR-2 | Return current score, band, contributing factors, freshness, and version lineage.          |
| FR-3 | Trigger downstream review when a configured band transition occurs.                        |
| FR-4 | Recompute a tenant or cohort under a new policy/model without corrupting the active score. |
| FR-5 | Correct state after late, amended, or retracted source events.                             |

### Non-functional requirements

| ID    | Requirement                                                                       |
| ----- | --------------------------------------------------------------------------------- |
| NFR-1 | Update ordinary customers within 60 seconds of an accepted event.                 |
| NFR-2 | Serve 10,000 reads/s at 30 ms p99 for 100 million customers.                      |
| NFR-3 | Preserve per-customer deterministic ordering and idempotent transition effects.   |
| NFR-4 | Finish a full tenant recomputation within 24 hours without violating online SLOs. |
| NFR-5 | Bound cross-tenant resource contention and provide score lineage for seven years. |

## Assumptions and exclusions

The score combines rule and model outputs. Final investigative disposition is outside scope.

## Interview prompts

1. Choose event partitioning, state store, serving store, and recomputation path.
2. How do corrections and retractions affect incremental aggregates?
3. How is one active score version switched after shadow comparison?
4. How are duplicate band transitions prevented during replay?

Solve before reading [the worked solution](../solutions/010-continuous-customer-risk-scoring.md).
