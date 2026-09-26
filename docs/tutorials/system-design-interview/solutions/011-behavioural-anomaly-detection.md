---
tldr: "Versioned streaming features and an approved model server produce signals while policy owns the final action."
when_to_use: "Use after attempting Case 011 to compare feature parity, model lifecycle, monitoring, and fallback."
---

# Solution 011: Behavioural Anomaly Detection

## Requirement traceability

| Requirements | Design response                                                                                          |
| ------------ | -------------------------------------------------------------------------------------------------------- |
| FR-1, NFR-1  | Keyed stream processors update entity/cohort features; online store serves a point-in-time vector.       |
| FR-2         | Model service returns score, artifact id, feature-set id, values/references, and explanation metadata.   |
| FR-3, NFR-4  | Decision policy is separate and declares rule-only/no-model fallback.                                    |
| FR-4, NFR-3  | Shadow router mirrors sampled features; monitoring compares latency, score, drift, and delayed outcomes. |
| FR-5         | Historical snapshot and log replay write evaluation-only outputs under a run id.                         |
| NFR-2, NFR-5 | Shared feature definitions with parity tests; separate identities and approvals for train/serve.         |

## Architecture

```text
events -> stream features -> online feature store -> model server -> signal -> policy
   |             |                   ^                  |
   |             +-> offline table --+ parity          +-> evidence
   v
historical lake -> point-in-time dataset -> train/evaluate -> registry -> approval
```

Feature definitions declare source, event-time window, lateness, default, type, and owner. Training joins each example
to features as they existed at the decision cutoff, never to later outcomes.

## Serving sequence

```text
decision API     feature store       model server       evidence
     | read(vector,as_of) |                |                |
     |------------------->|                |                |
     |<-- values+freshness|                |                |
     | score(model,vector)---------------->|                |
     |<------ score+version+reasons -------|                |
     |----------------------------------------------------->| record
```

Set model deadline below the decision deadline. Pods load an immutable model during startup, verify its hash and feature
contract, warm representative inputs, then become ready. HPA uses concurrency/latency plus CPU; model size and startup
time require warm headroom.

## Monitoring and release

Before labels: schema, missingness, freshness, feature distribution, score distribution, latency, timeout, fallback, and
action rate by cohort. After labels: calibration, precision/recall, cost, and outcome delay. Shadow first, canary a
tenant cohort, expand only when technical and domain gates pass, and keep pointer rollback.

Drift opens investigation; it does not automatically retrain or promote. Training artifacts and datasets are immutable,
and promotion requires role-separated approval.

## Alternatives rejected

Embedding a frequently changing model in every decision image couples releases and multiplies memory. Letting the model
return the final business action hides policy and fallback. Joining training data to current feature tables leaks future
information.
