---
tldr:
  "Governed labels feed reproducible point-in-time datasets, approved artifacts, warmed serving, and cohort monitoring."
when_to_use: "Use after attempting Case 017 to compare the complete governed ML lifecycle."
---

# Solution 017: Feedback, Training, and Model Serving

## Requirement traceability

| Requirements        | Design response                                                                                             |
| ------------------- | ----------------------------------------------------------------------------------------------------------- |
| FR-1                | Raw feedback is immutable; validated label revisions cite reviewer, evidence, policy, and superseded label. |
| FR-2, NFR-3         | Dataset manifest pins source snapshots, cutoff, feature/code versions, cohort query, and hashes.            |
| FR-3                | Isolated jobs evaluate cost, calibration, and cohort metrics against an approved baseline.                  |
| FR-4, FR-6          | Role-separated registry stages and immutable active pointer support shadow/canary/rollback.                 |
| FR-5, NFR-1         | Warm model-serving Pods return artifact/feature versions under a 40 ms budget.                              |
| NFR-2, NFR-4, NFR-5 | Separate compute quotas, multi-layer monitoring, and residency-scoped data/artifact stores.                 |

## Lifecycle

```text
investigator action -> raw feedback -> label review -> validated label versions
                                                   |
source snapshots + point-in-time features ---------+-> dataset manifest
                                                         |
                                                   train/evaluate
                                                         |
                                               registry + approval
                                                  /      |      \
                                               shadow  canary  active
```

A reviewer disposition is an observation, not automatically ground truth. Conflicts, uncertainty, and later outcomes
remain represented. Dataset construction joins features strictly before the prediction time and holds out entities/time
to avoid leakage.

## Artifact and serving contract

The artifact bundle contains model binary, input schema, feature-set id, preprocessing, runtime constraints, evaluation
report, training manifest, owner, signature, and checksum. Serving accepts a typed vector and deadline, then returns
score, artifact id, reason metadata, and latency. Policy remains outside the model service.

```text
Pod startup: fetch -> verify signature/hash -> load -> warm -> self-test -> ready
termination: unready -> drain -> release large model memory -> exit
```

Use dedicated node pools/requests when model memory or accelerators demand them. Scale on concurrency/latency and keep
warm replicas because node and model startup exceed request deadlines.

## Release and rollback

Shadow captures no business effect. Canary by stable tenant/entity cohort to avoid one entity alternating models.
Compare technical, score, action-simulation, fairness/cohort, and delayed outcome gates. Rollback switches active
pointer; historic evidence continues referencing the newer artifact.

## Monitoring and failure

Feature/service failures follow an approved fallback contract. Monitor artifact load failures, timeout, feature parity,
freshness, missingness, drift, calibration, outcome metrics, fallback, and slice sample size. Training cannot directly
read production databases; curated snapshots protect online workloads.

## Alternatives rejected

Automatically retraining and promoting from raw feedback creates a self-reinforcing loop. Mutable artifacts destroy
reproducibility. One global model ignores residency and tenant contract constraints.
