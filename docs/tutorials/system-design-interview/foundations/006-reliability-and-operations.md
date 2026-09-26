---
tldr: "Connects SLOs to observability, overload control, deployment safety, incident response, backup, and recovery."
when_to_use: "Use to turn a happy-path architecture into an operable production system."
---

# Reliability and Operations

Reliability is a runtime property produced by capacity, safe change, observability, recovery practice, and ownership.
Three replicas do not make a service reliable if they share one failure domain or depend on one exhausted database.

## Failure-domain map

```text
region
+----------------------------------------------------------+
| zone A                 zone B                 zone C      |
| +------+ +------+      +------+ +------+      +------+    |
| | pod  | | pod  |      | pod  | | pod  |      | pod  |    |
| +------+ +------+      +------+ +------+      +------+    |
+----------------------------------------------------------+
        \________________ shared database? ______________/
```

List process, node, zone, region, dependency, credential, configuration, and human-operator failures. Spread replicas
only across domains the design can actually survive.

## Golden signals plus domain signals

- latency: percentile by endpoint and result, not a global average;
- traffic: requests, events, bytes, and active tenants;
- errors: timeout, rejection, validation, dependency, and semantic error;
- saturation: CPU, memory, queue age, pool utilization, disk, quota;
- domain: decisions delayed, evidence gaps, unmatched entities, replay drift.

Metrics show trends, traces show a request path, logs explain discrete events, and profiles explain resource cost. Use
correlation identifiers, but never put secrets or unrestricted personal data in telemetry.

## Alert on symptoms and exhaustion

```text
page:   SLO burn threatens users now or soon
ticket: capacity/risk needs working-hours action
record: diagnostic signal for later analysis
```

Multi-window burn-rate alerts catch fast outages and slow leaks without paging on every transient error. Pair them with
runbooks that name the service owner, dashboards, safe mitigation, rollback, and escalation.

## Load shedding and graceful degradation

Define priority before an incident:

```text
P0 synchronous decision -> protected capacity
P1 evidence persistence  -> durable buffer
P2 dashboard refresh     -> stale cache allowed
P3 bulk export           -> pause and resume
```

Admission control at the edge prevents work the system cannot finish. A bounded queue protects memory but needs an
explicit full policy. Reject silent dropping for auditable events; reject blocking forever because it converts
saturation into global timeout.

## Safe change

Use backward-compatible schema changes, feature flags with owners and expiry, canaries, automated health criteria, and
fast rollback. For stateful changes, rollback may not reverse data. Use expand/migrate/contract:

```text
old readers/writers
       |
add compatible field -> dual/read-compatible period -> backfill -> stop old writes -> remove old field
```

## Recovery

Backups are unproven until a restore succeeds. Define:

- RPO: maximum lost committed data measured in time;
- RTO: maximum time to restore the service outcome;
- restore order: identity, configuration, primary data, indexes, derived projections;
- reconciliation: how restored state is compared with durable inputs;
- exercise cadence and evidence.

For derived search indexes, rebuild from canonical storage instead of treating the index as the only backup. For legal
evidence, use immutable retention, integrity verification, and controlled deletion rather than ordinary snapshots alone.

## Operational review

Before calling a design complete, answer:

1. What pages the on-call engineer?
2. What can they safely do in ten minutes?
3. How does the system behave at twice tested capacity?
4. How is a bad release stopped and reverted?
5. When was restore and regional failover last exercised?
