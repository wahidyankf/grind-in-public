---
tldr:
  "Migrates a Python modular monolith through measured seams, strangler routing, data ownership, and reconciliation."
when_to_use: "Use when independent scaling or ownership justifies extraction from an existing monolith."
---

# Modular Monolith to Microservices

The goal is not a service count. The goal is a safer operating model for a specific pressure: independent scaling,
release cadence, security isolation, fault containment, team ownership, or technology constraint. Preserve the modular
monolith when those pressures are absent.

## Establish the baseline

```text
+---------------- Python deployable ----------------+
| API -> application modules -> domain modules      |
|                    |                               |
|              shared database                      |
+---------------------------------------------------+
```

Before extraction, record p50/p95/p99 latency, errors, throughput, deployment frequency, recovery time, dependency
graph, change coupling, query ownership, and incident history. Strengthen in-process module APIs and prevent arbitrary
cross-module table access. A tangled monolith becomes a tangled distributed system if moved without this step.

## Choose a seam

Score candidate boundaries against:

- cohesive business capability and vocabulary;
- clear input/output contracts;
- independent load or availability needs;
- few synchronous dependencies;
- data that can acquire one writer;
- a team able to own build, run, on-call, and evolution.

Good first extractions are often asynchronous, read-heavy, or low-risk. Reject the highest-value synchronous core as a
pilot when rollback and equivalence are not yet proven.

## Migration states

```text
0. module call
   monolith --------------------> module + shared tables

1. observe
   monolith -> outbox -> new service builds shadow state

2. shadow
   request -> monolith result
           `-> new service result -> compare only

3. canary
   router -> small cohort -> new service
          `-> remainder   -> monolith

4. authority
   router -> new service -> owned data
                       `-> compatibility events for monolith

5. retire
   remove old code, writes, tables, flags, and reconciliation path
```

Every state has entry criteria, observable success, rollback, and a time limit. Permanent dual write is not a target
architecture.

## Data ownership

Two services must not independently update the same record. The transition tools are:

- transactional outbox: commit domain change and event in one database transaction;
- change-data capture: observe legacy changes when adding an outbox is initially impractical;
- backfill: copy historical data with resumable checkpoints;
- dual read: compare or fall back temporarily, with metrics;
- reconciliation: compare counts, hashes, invariants, and sampled records;
- expand/contract schema: support mixed versions during rollout.

```text
monolith transaction
  +-- update owned legacy row
  `-- insert outbox row
             |
             v
        relay -> durable log -> extracted service -> idempotent projection
```

CDC exposes storage changes, not necessarily domain meaning. Prefer an explicit outbox for the long-term contract.

## Synchronous extraction

Set end-to-end deadlines, propagate cancellation, retry only idempotent calls, bound concurrency, and define fallback.
Avoid long call chains:

```text
bad:  edge -> A -> B -> C -> D
better: edge -> orchestrator -> parallel B/C, or asynchronous workflow
```

Each network call multiplies availability risk and tail latency. Keep a capability inside the monolith if extraction
would create chatty calls around one transaction.

## Kubernetes rollout

Package the monolith first: deterministic image, health endpoints, graceful termination, resource requests, telemetry,
and safe schema migration. This separates containerization risk from service-boundary risk. Deploy extracted services
with independent service accounts, budgets, autoscaling signals, dashboards, and ownership.

## Cutover gates

Proceed only when:

- shadow outputs meet a defined equivalence threshold;
- backfill and continuous changes converge;
- new path meets latency and error SLO under peak load;
- rollback does not require data destruction;
- on-call runbooks and dashboards exist;
- ownership and cost are accepted.

## Decommissioning

An extraction is incomplete while old writes, tables, flags, queues, permissions, dashboards, and code remain active.
Observe a safety window, archive required evidence, remove compatibility paths, and update the system-of-record map.
