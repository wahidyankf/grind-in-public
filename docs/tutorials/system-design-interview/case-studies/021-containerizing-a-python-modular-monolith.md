---
tldr: "Design the safe containerization and Kubernetes deployment of an existing Python modular monolith."
when_to_use: "Use to separate runtime modernization from service extraction."
---

# Case 021: Containerizing a Python Modular Monolith

An existing Python modular monolith runs on virtual machines, uses one relational database, processes HTTP and
background jobs, and performs database migrations during startup. Move it to Kubernetes without changing domain
boundaries yet.

### Functional requirements

| ID   | Requirement                                                                                          |
| ---- | ---------------------------------------------------------------------------------------------------- |
| FR-1 | Build one reproducible application image and run API, worker, and migration entry points explicitly. |
| FR-2 | Expose startup, readiness, liveness, metrics, and graceful-shutdown behaviour.                       |
| FR-3 | Run schema migration once per release with status visible before incompatible traffic.               |
| FR-4 | Roll out API and workers progressively and roll back application versions safely.                    |
| FR-5 | Preserve scheduled jobs, queue semantics, secrets, configuration, and audit behaviour.               |

### Non-functional requirements

| ID    | Requirement                                                                                     |
| ----- | ----------------------------------------------------------------------------------------------- |
| NFR-1 | No planned request loss and no duplicate non-idempotent job effects during rollout.             |
| NFR-2 | Meet the existing 99.9% availability and p99 latency SLO at peak load.                          |
| NFR-3 | Scale API and workers independently without exceeding database connection capacity.             |
| NFR-4 | Use least-privilege service accounts, immutable images, and external secret delivery.           |
| NFR-5 | Return to the prior runtime path during the migration window without destructive data rollback. |

## Assumptions and exclusions

The database remains managed outside Kubernetes. Splitting services and rewriting modules are out of scope.

## Interview prompts

1. Inventory hidden VM assumptions: filesystem, signals, cron, local cache, and process supervision.
2. Define images, commands, Deployments/Jobs/CronJobs, probes, and shutdown.
3. How do expand/contract migrations work across mixed versions?
4. How are capacity, connections, rollout, observation, and runtime rollback validated?

Solve before reading [the worked solution](../solutions/021-containerizing-a-python-modular-monolith.md).
