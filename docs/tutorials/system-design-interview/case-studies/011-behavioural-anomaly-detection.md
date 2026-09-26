---
tldr: "Design feature generation, anomaly scoring, cohort baselines, delayed labels, and safe model rollout."
when_to_use: "Use to practise streaming ML architecture and monitoring beyond service health."
---

# Case 011: Behavioural Anomaly Detection

Design a system that detects unusual transaction behaviour relative to an entity's history and peer cohort, then emits a
scored signal for policy evaluation.

### Functional requirements

| ID   | Requirement                                                                                          |
| ---- | ---------------------------------------------------------------------------------------------------- |
| FR-1 | Build event-time entity and cohort features from accepted events.                                    |
| FR-2 | Score an event with the approved model and return score, model version, and feature lineage.         |
| FR-3 | Emit a signal only; tenant policy decides the final action and threshold.                            |
| FR-4 | Shadow candidate models and compare them against the incumbent with delayed outcomes.                |
| FR-5 | Backfill features and rescore a historical interval for evaluation without duplicating live effects. |

### Non-functional requirements

| ID    | Requirement                                                                               |
| ----- | ----------------------------------------------------------------------------------------- |
| NFR-1 | Produce a live score within 100 ms p99 at 20,000 events/s.                                |
| NFR-2 | Keep online/offline feature discrepancy below defined per-feature tolerances.             |
| NFR-3 | Detect schema, feature freshness, distribution, calibration, and outcome drift by cohort. |
| NFR-4 | Fall back to a declared no-model or prior-score policy during serving failure.            |
| NFR-5 | Isolate training data, artifacts, approvals, and serving access.                          |

## Assumptions and exclusions

Labels arrive days or weeks later and can be noisy. Choosing a specific ML algorithm is secondary to lifecycle design.

## Interview prompts

1. Separate stream features, offline features, training, registry, serving, and policy.
2. How do point-in-time correct datasets avoid leakage?
3. What is monitored before labels arrive and after they arrive?
4. How are model loading, canary, shadow, rollback, and evidence handled on Kubernetes?

Solve before reading [the worked solution](../solutions/011-behavioural-anomaly-detection.md).
