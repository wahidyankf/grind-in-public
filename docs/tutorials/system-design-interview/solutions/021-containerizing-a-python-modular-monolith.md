---
tldr:
  "Containerize the unchanged monolith first with explicit processes, safe migrations, probes, draining, and rollback."
when_to_use: "Use after attempting Case 021 to compare runtime migration controls before service extraction."
---

# Solution 021: Containerizing a Python Modular Monolith

## Requirement traceability

| Requirements       | Design response                                                                                    |
| ------------------ | -------------------------------------------------------------------------------------------------- |
| FR-1               | One pinned multi-stage image; explicit API, worker, migration, and scheduled-job commands.         |
| FR-2, NFR-1        | Semantic probes, SIGTERM draining, bounded job acknowledgement, and termination grace.             |
| FR-3               | Release migration Job uses expand/contract schemas and a database advisory lease.                  |
| FR-4, NFR-2, NFR-5 | Canary runtime traffic, SLO comparison, compatible schema, and reversible routing to old runtime.  |
| FR-5, NFR-4        | Kubernetes CronJobs, external secrets, service accounts, immutable image digest, and audit parity. |
| NFR-3              | Separate API/worker Deployments and autoscaling under a shared database connection budget.         |

## Target runtime

```text
                    one immutable application image
                       /          |          \
                 API command  worker command  migration command
                    |             |                 |
               Deployment    Deployment             Job
                    |             |
                  Service       queue          CronJobs use same image
                     \           /
                     managed relational database
```

First inventory local filesystem writes, embedded scheduler, signal handling, process manager, environment config,
hostnames, caches, static files, credentials, and startup migrations. Move durable files to object storage, scheduling
to CronJobs, and secrets to workload identity/secret references.

## Probe and shutdown contract

- startup: process initialized, required configuration validated, database schema compatible;
- readiness: accepts new work and has required pools; transient optional dependency failure does not restart the Pod;
- liveness: event loop/process is irrecoverably stuck, not merely unable to reach the database;
- termination: become unready, stop accepting/fetching, finish or release current units, close pools, exit before grace.

Workers acknowledge only after durable effect. If SIGTERM arrives mid-job, finish within a bounded lease or let the
message be redelivered to an idempotent handler.

## Schema and rollout sequence

```text
1 expand: add nullable/new structures compatible with old and new code
2 run migration Job once; verify schema marker
3 canary new API/worker Pods while old runtime remains compatible
4 expand traffic; observe latency, errors, jobs, DB pools, domain outcomes
5 after safety window, contract unused schema in a later release
```

Do not run migration independently in every Pod startup. Rollback changes routing/image; it never depends on reversing a
destructive data migration.

## Capacity and Kubernetes

Set requests from measured usage, memory limits from safe process bounds, and HPA on API concurrency/latency and worker
queue age. A central connection budget allocates pools across maximum replicas so scaling Pods cannot exhaust the
database. Spread replicas across zones and use disruption budgets for voluntary work.

## Verification and alternatives

Mirror or canary low-risk traffic, compare response/effect, perform load and graceful-termination tests, and run
restore/ rollback exercises. Monitor Pod starts, probe failures, SIGTERM duration, job redelivery, schema version,
connection saturation, and old/new SLOs.

Splitting services during this move is rejected because it combines runtime and domain-risk variables. Running a
database inside a StatefulSet is rejected without a demonstrated database-operations capability.
