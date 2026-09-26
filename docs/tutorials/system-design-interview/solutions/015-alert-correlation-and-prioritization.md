---
tldr:
  "Deterministic and learned correlation produce bounded clusters, while versioned ranking feeds fair skilled queues."
when_to_use: "Use after attempting Case 015 to compare correlation algorithms, ranking, replay, and human capacity."
---

# Solution 015: Alert Correlation and Prioritization

## Requirement traceability

| Requirements | Design response                                                                         |
| ------------ | --------------------------------------------------------------------------------------- |
| FR-1, NFR-1  | Partitioned alert log, schema validation, and urgent fast lane.                         |
| FR-2, NFR-3  | Exact idempotency then bounded time-window correlation into versioned clusters.         |
| FR-3, NFR-2  | Deterministic ranking function/model with features, policy version, and explanation.    |
| FR-4         | Skill/tenant queues use deadlines, reserved capacity, and weighted fairness.            |
| FR-5         | Append-only cluster/rank history and governed raw-versus-validated feedback.            |
| NFR-4, NFR-5 | Stable identities, inbox/outbox effects, and tenant-scoped authorization/configuration. |

## Pipeline

```text
alerts -> exact dedup -> correlation candidates -> cluster state -> ranker -> fair work queues
                        |                         |              |
                   entity/time indexes       history        assignment service
```

Start with deterministic correlation keys: same subject/rule/window or same transaction. Add bounded joins through
shared entity, device, or counterparty indexes. Union-find works for a batch of additive links, but online clusters need
versioned merge records; split is a new cluster revision, not destructive mutation.

## Stable identities

`alert_id` derives from detector observation identity. `cluster_id` is stable across ordinary additions; merges create a
new canonical id plus aliases from parents. `case_id` is workflow-owned and may reference one or more cluster revisions.
This separation makes replay idempotent without pretending correlation never changes.

## Ranking and fairness

Compute score from declared versioned features. Maintain top candidates per queue in a heap, but persist authoritative
queue state because an in-memory heap is lost on restart. Weighted fair scheduling gives each tenant/skill class service
while urgent deadlines can preempt within a bounded reserve.

```text
queue A weight 3: A A A
queue B weight 1:       B
urgent reserve:         ^ may preempt, capped
```

Age increases priority to prevent starvation. Re-ranking writes a reasoned revision and never silently rewrites the
original score.

## Operations and alternatives

Monitor ingest lag, correlation candidate/cluster sizes, merge rate, rank distribution, queue age by tenant/skill,
deadline risk, assignment retries, feedback drift, and cases per alert. Cap cluster nodes/time span and route
exceptional superclusters to specialized handling.

Reject one case per alert because it overwhelms investigators. Reject a single global priority heap because it ignores
skills, authorization, tenant fairness, durability, and starvation.
