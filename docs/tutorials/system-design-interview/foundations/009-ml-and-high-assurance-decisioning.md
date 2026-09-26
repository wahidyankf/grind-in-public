---
tldr: "Separates model training, serving, policy, evidence, monitoring, and fallback in high-assurance decisions."
when_to_use: "Use when a system combines rules, statistical models, human review, and explainability."
---

# ML and High-Assurance Decisioning

A model score is an input to a product decision, not the decision itself. Production design must connect data and model
versions to an auditable policy, define failure behaviour, and monitor outcome quality after deployment.

## Offline and online planes

```text
historical data -> validation -> features -> training -> registry -> approval
                                                              |
                                                              v
request -> online features -> model serving -> policy -> action + evidence
                  ^                              |
                  +-------- parity checks -------+
```

Training optimizes and evaluates a candidate. Serving loads an approved immutable artifact and returns a score. Policy
combines the score with rules, thresholds, exceptions, and action semantics. Store the model version, feature snapshot
or references, rule version, threshold, decision, and explanation metadata.

## Online feature contract

Point-in-time correctness prevents training leakage: a training example may use only information available at its event
time. Online/offline parity means the same feature definition and transformation produce equivalent values.

```text
event time ---- feature cutoff ---- outcome becomes known
                    ^
                    +-- training may read only left of cutoff
```

Reject a feature store if there are few features, one model, and a simple database lookup meets latency. Adopt it when
reuse, freshness, parity, and ownership justify the operational surface.

## Synchronous serving choices

| Choice       | Fits                                                | Cost / reject when                                      |
| ------------ | --------------------------------------------------- | ------------------------------------------------------- |
| In-process   | Tiny stable model, lowest latency, release together | Memory per Pod; reject for frequent independent updates |
| Model server | Shared models, independent scaling, batching        | Network hop; reject when availability cannot match path |
| Precompute   | Slowly changing entity scores, heavy computation    | Staleness; reject for per-event dynamic features        |

Set a deadline shorter than the caller's deadline. On timeout, choose an explicit policy: safe rule-only fallback,
manual review, rejection, or fail-open. The correct choice follows business harm, not generic availability preference.

## Release process

```text
offline gate -> shadow -> small canary -> cohort expansion -> default
                   |          |                |
                   +----- compare outcome, drift, latency, cost
```

Shadow evaluation must not create the business effect. Compare candidate and incumbent with delayed labels and sliced
metrics. A global average can hide harm to a language, region, channel, or tenant.

## Monitoring

- input schema and missingness;
- feature freshness and distribution drift;
- prediction distribution and calibration;
- service latency, timeout, load failure, and resource use;
- decision/action rate by meaningful slice;
- delayed precision, recall, cost, and investigator feedback;
- policy overrides and fallback rate.

Drift is an investigation signal, not automatic proof the model is bad. Seasonality, product changes, attacks, and
pipeline bugs can all move distributions.

## Explainability and evidence

Provide reason codes tied to versioned policy and stable feature definitions. Do not promise that a feature-attribution
method is a causal explanation. Preserve enough evidence to reproduce the executed decision, while minimizing sensitive
data and respecting retention.

## Human review

Human review is a queueing system: define priority, skills, assignment, capacity, service time, escalation, and quality
sampling. Feed outcomes back only after label validation; investigator actions can be inconsistent or biased.
