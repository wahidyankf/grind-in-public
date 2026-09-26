---
tldr:
  "A capability authority map, priority capacity plan, fenced recovery, and explicit retirement gates finish migration
  safely."
when_to_use: "Use after attempting Case 024 as a capstone for architecture and engineering leadership."
---

# Solution 024: Hybrid Cutover, Recovery, and Decommissioning

## Requirement traceability

| Requirements | Design response                                                                                                 |
| ------------ | --------------------------------------------------------------------------------------------------------------- |
| FR-1, NFR-3  | Versioned authority map records capability/data/event owner, migration state, SLO, team, and exit gate.         |
| FR-2, NFR-4  | Stable tenant cohorts, shadow/canary routing, outbox/inbox, reconciliation, and reversible pre-authority steps. |
| FR-3, NFR-1  | Measured bottleneck plan, workload priority, cell scaling, load tests, and admission control.                   |
| FR-4, NFR-2  | Tenant writer epochs, warm recovery, canonical replay, projection rebuild, and invariant checks.                |
| FR-5, NFR-5  | Retirement checklist deletes obsolete runtime surfaces after evidence window.                                   |
| FR-6         | Capabilities without scaling/ownership/security/release pressure remain modules in the monolith.                |

## Authority map

```text
capability       writer authority       read projection       state       owner
policy author    monolith DB            service artifacts     migrating   team A
evaluation       evaluation service     local artifact cache  extracted   team A
alerting         alert service          case compatibility    canary      team B
case workflow    monolith module        search index          retained    team C
profiles         monolith module        feature projections   retained    team C
```

This map is more valuable than a target diagram with ambiguous arrows. A migration state must have entry evidence, exit
evidence, rollback, owner, and review date.

## Tenfold growth plan

Measure before splitting. Likely first constraints are database connections/write IOPS, hot log partitions, and
synchronous feature/model latency. Respond in this order:

1. remove waste and tune queries/indexes/pools;
2. add admission control and priority: decision, durable ingestion, evidence, projections, bulk;
3. partition tenant/entity workloads and isolate outliers;
4. scale stateless Pods and log consumers within downstream budgets;
5. add read projections/caches for proven shapes;
6. extract only capabilities whose independent pressure is measured.

```text
saturation -> pause shadow/bulk -> slow projections -> reject over-quota tenants
           -> preserve decision + durable evidence capacity
```

Load tests include skew, dependency slowdown, Pod/node loss, autoscaler delay, and backlog recovery—not only uniform
happy-path throughput.

## Regional recovery

Declare incident, freeze routing, advance an independently controlled writer epoch, promote recovery stores at recorded
watermarks, start critical workloads first, route a small cohort, replay durable events, rebuild derived stores, and
compare accepted ids, sequence gaps, counts, hashes, decisions/evidence, and open workflows. Restore lower-priority
reporting only after critical invariants pass. Failback is another fenced migration.

## Decommission gate

```text
no traffic/writes -> retention window observed -> reconciliation clean -> owner approval
       |
       v
delete old code -> tables -> events/CDC -> flags -> credentials -> infrastructure -> alerts/runbooks
```

Record cost, incident surface, deployment lead time, SLO, recovery, and team ownership before/after. Migration value is
improved outcomes, not number of services.

## Leadership decision

Keep case and profile capabilities in the modular monolith while their transactional cohesion and common ownership
outweigh extraction pressure. Invest in internal module boundaries and test contracts. This is an intentional target,
not unfinished work.

## Alternatives rejected

An all-at-once rewrite removes incremental evidence and rollback. Permanent dual stacks double operating surface.
Extracting retained modules solely for architectural uniformity creates network and ownership cost without a
requirement.
