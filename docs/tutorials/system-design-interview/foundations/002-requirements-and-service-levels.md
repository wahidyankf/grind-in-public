---
tldr: "Turns an ambiguous prompt into prioritized, measurable functional and non-functional requirements."
when_to_use: "Use during the first five minutes of every system-design interview and before choosing components."
---

# Requirements and Service Levels

Requirements are the control surface for the whole answer. Number them so later decisions can cite them.

## Functional requirements

Write actor-visible behaviour, not implementation:

| ID   | Requirement                                                     | Priority |
| ---- | --------------------------------------------------------------- | -------- |
| FR-1 | A producer submits an event with a stable idempotency key.      | Must     |
| FR-2 | A policy owner publishes a versioned decision configuration.    | Must     |
| FR-3 | An investigator retrieves the decision and supporting evidence. | Must     |
| FR-4 | An administrator replays a bounded failed interval.             | Should   |

"Use a stream" is not an FR. The stream may be one solution for FR-1 and FR-4.

## Non-functional requirements

A useful NFR includes a measurement boundary, percentile or probability, time window, and conditions:

| ID    | Requirement                                                                            |
| ----- | -------------------------------------------------------------------------------------- |
| NFR-1 | Acknowledgement latency is at most 100 ms at p99 within a region at 10,000 requests/s. |
| NFR-2 | Accepted events have 99.999999999% annual durability.                                  |
| NFR-3 | Monthly API availability is 99.95%, excluding documented maintenance.                  |
| NFR-4 | Tenant data remains in its contracted region and is encrypted in transit and at rest.  |
| NFR-5 | A regional loss has an RPO of 5 minutes and RTO of 30 minutes.                         |

Ask which one wins when they conflict. Zero data loss across regions and single-digit-millisecond writes may be
incompatible during a partition.

## SLI, SLO, SLA, and error budget

```text
measurement       internal target       external promise
   SLI       --->      SLO        --->       SLA
                    99.95%/month
                         |
                         v
             0.05% error budget/month
```

- An SLI is the measured ratio, such as successful eligible requests divided by eligible requests.
- An SLO is the internal target for that SLI.
- An SLA is a contractual promise and may define credits or exclusions.
- The error budget is `1 - SLO`; it creates an explicit reliability-versus-change policy.

Do not call every `5xx` an availability failure. Exclude invalid client requests, health probes, and requests rejected
by an intentional overload policy. Include timeouts and semantically wrong successes when the user did not receive the
promised outcome.

## Latency budgets

Allocate an end-to-end target before optimizing a component:

```text
p99 budget: 200 ms

edge/auth      API queueing      decision       datastore      reserve
   20 ms   +      25 ms      +     80 ms    +      45 ms   +   30 ms
```

Percentiles do not add cleanly because calls are correlated and parallel work uses the maximum branch latency. The
budget is a design hypothesis; validate it with load tests and traces.

## Failure and recovery questions

Clarify these explicitly:

- Is it safer to reject, allow, or defer when a dependency is unavailable?
- May results be stale, and for how long?
- Is duplicate processing acceptable if the effect is idempotent?
- Which data loss is tolerable, and which records are legal evidence?
- Does recovery restore service first and reconcile later, or block until data is complete?

## Checkpoint

Before drawing architecture, the interviewer should be able to point to a numbered requirement for every important
constraint. If they cannot, ask rather than inventing precision.
