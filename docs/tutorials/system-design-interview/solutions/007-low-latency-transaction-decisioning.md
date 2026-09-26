---
tldr:
  "A deadline-aware orchestrator pins versions, runs signals in parallel, commits evidence, and applies explicit
  fallback."
when_to_use: "Use after attempting Case 007 to compare synchronous latency, idempotency, evidence, and degradation."
---

# Solution 007: Low-Latency Transaction Decisioning

## Requirement traceability

| Requirements      | Design response                                                                                               |
| ----------------- | ------------------------------------------------------------------------------------------------------------- |
| FR-1, NFR-3       | Idempotency reservation stores request hash and final immutable decision.                                     |
| FR-2, NFR-1       | Orchestrator pins policy, launches independent feature/list/model calls in parallel, then evaluates rules.    |
| FR-3, FR-4, NFR-4 | Evidence transaction stores result, versions, inputs/references, reasons, and outbox before response.         |
| FR-5              | Shadow work is sampled, lower priority, side-effect-free, and compared asynchronously.                        |
| NFR-2, NFR-5      | Per-dependency fallback matrix, tenant cells/quotas/caches, overload admission, and last-known-good versions. |

## Latency budget

```text
150 ms p99
edge/auth 15 | validate/idempotency 15 | parallel signals 60 | policy 20 | evidence 25 | reserve 15
```

The budget is validated under load; percentiles are not simply additive. Dependencies receive remaining deadlines and
bounded concurrency.

## Architecture and sequence

```text
client -> decision API -> idempotency/evidence DB
                 |
                 +---- parallel ----> entity cache/store
                 +------------------> velocity service
                 +------------------> list-match service
                 `------------------> model service
                 |
              local pinned rules -> policy result
```

```text
client      API       signals       policy       evidence/outbox
  | request  |           |             |                |
  |--------->| reserve   |             |                |
  |          |======== parallel ======>|                |
  |          |<======= bounded results =|               |
  |          |------------------------>| decide         |
  |          |----------------------------------------->| commit
  |<---------| action + reasons + versions              |
```

The evidence transaction updates the idempotency row and inserts an outbox record. Relay failure delays downstream
investigation but does not lose the committed decision. A retry reads the exact result rather than re-evaluating under
new policy.

## Dependency policy

| Dependency              | Failure response                                |
| ----------------------- | ----------------------------------------------- |
| entity optional data    | use declared missing feature and reason code    |
| velocity required       | `review` or tenant-configured reject            |
| reference list required | never clear; return review/reject               |
| model                   | approved rule-only fallback when policy permits |
| evidence database       | do not return an executed decision              |

The domain owns these choices. Availability cannot silently override safety.

## Kubernetes and operations

Use stateless API Deployments spread across zones, warmed immutable rule/model metadata, HPA on concurrency plus CPU,
and connection budgets below database capacity. Readiness requires safe baseline policy, not every optional dependency.
Load shed shadow work first, then noncritical enrichment; preserve core decision capacity.

Monitor end-to-end SLI, component deadline use, fallback/action rate by tenant, version skew, idempotency states, outbox
age, and shadow disagreement. Canary policy/model/service separately so regressions are attributable.

## Alternatives rejected

A serial dependency chain exceeds tail latency. Re-evaluating retries can change a business outcome. Returning before
evidence commit violates reproducibility. Synchronously waiting for downstream case creation couples the decision SLO to
a noncritical consumer.
