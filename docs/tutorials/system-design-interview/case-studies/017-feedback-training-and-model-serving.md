---
tldr: "Design governed feedback, point-in-time training data, model approval, serving, and rollback."
when_to_use: "Use to practise the complete ML lifecycle rather than only a prediction endpoint."
---

# Case 017: Feedback, Training, and Model Serving

Design a platform that turns validated investigation outcomes into versioned datasets, trains candidate models, and
serves approved artifacts to synchronous decision systems.

### Functional requirements

| ID   | Requirement                                                                                  |
| ---- | -------------------------------------------------------------------------------------------- |
| FR-1 | Capture raw feedback, reviewer identity, evidence, uncertainty, and later label corrections. |
| FR-2 | Build reproducible point-in-time datasets with data, feature, code, and policy lineage.      |
| FR-3 | Train and evaluate candidates by tenant-approved cohorts and cost-sensitive metrics.         |
| FR-4 | Require approval before promoting an immutable model version to shadow, canary, or active.   |
| FR-5 | Serve predictions with feature/model versions, explanations, deadlines, and fallback.        |
| FR-6 | Roll back the active pointer without losing decisions made under newer versions.             |

### Non-functional requirements

| ID    | Requirement                                                                             |
| ----- | --------------------------------------------------------------------------------------- |
| NFR-1 | Serving adds at most 40 ms p99 at 10,000 predictions/s.                                 |
| NFR-2 | Training jobs cannot starve online serving or production databases.                     |
| NFR-3 | Artifacts, datasets, approvals, and lineage are immutable and access-controlled.        |
| NFR-4 | Detect service, feature, drift, calibration, and delayed outcome regressions by cohort. |
| NFR-5 | Meet tenant residency and prohibit unapproved cross-tenant training.                    |

## Assumptions and exclusions

The exact model family is not prescribed. Feature computation details are covered in Cases 004 and 011.

## Interview prompts

1. Draw data/feature/training/registry/deployment/serving/monitoring boundaries.
2. How are noisy feedback and label corrections governed?
3. How is online/offline feature parity measured?
4. How do Kubernetes rollout, model warming, resource isolation, and rollback work?

Solve before reading [the worked solution](../solutions/017-feedback-training-and-model-serving.md).
