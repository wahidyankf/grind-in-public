---
tldr: "Design a durable, ordered, tenant-isolated ingestion platform with replay and backpressure."
when_to_use: "Use to practise APIs, logs, partitioning, idempotency, quotas, and projection freshness."
---

# Case 001: Multi-Tenant Event Ingestion

Design the intake layer for a high-assurance SaaS platform. Customers send transaction and account events over HTTPS and
bulk object uploads. Multiple internal consumers need an immutable, replayable stream.

## Actors and scope

- customer producer and tenant administrator;
- ingestion API and bulk importer;
- independent decision, search, and analytics consumers;
- platform operator replaying a bounded interval.

### Functional requirements

| ID   | Requirement                                                                                            |
| ---- | ------------------------------------------------------------------------------------------------------ |
| FR-1 | Accept a versioned event carrying tenant, source, event identity, entity key, event time, and payload. |
| FR-2 | Return the same acceptance result when the same idempotency key and payload are retried.               |
| FR-3 | Preserve order for one tenant/entity key; no global ordering is required.                              |
| FR-4 | Let authorized operators replay one tenant and time range without affecting other consumers.           |
| FR-5 | Quarantine invalid or repeatedly failing events with a visible repair workflow.                        |
| FR-6 | Expose accepted, rejected, duplicate, lag, and oldest-event-age status per tenant.                     |

### Non-functional requirements

| ID    | Requirement                                                                    |
| ----- | ------------------------------------------------------------------------------ |
| NFR-1 | Sustain 20,000 events/s peak, 2 KiB average, and a tenfold tenant-size skew.   |
| NFR-2 | Acknowledge an accepted online event within 120 ms at p99 in-region.           |
| NFR-3 | Lose no acknowledged event after a single-node or single-zone failure.         |
| NFR-4 | Provide 99.95% monthly ingestion availability and bounded overload responses.  |
| NFR-5 | Keep replayable data for 30 days and immutable evidence for seven years.       |
| NFR-6 | Isolate tenant authorization, quotas, encryption context, metrics, and replay. |

## Assumptions and exclusions

Payload semantics are validated by downstream domain consumers; intake validates envelope and registered schema. Cross-
region active-active intake, end-user search, and decision logic are out of scope.

## Interview prompts

1. Where is the durable acknowledgement point?
2. Which key partitions the log, and how is a hot tenant handled?
3. How are online and bulk paths normalized without destroying producer backpressure?
4. What happens after commit but before response, and after consumer effect but before acknowledgement?
5. How are quarantine, replay, schema evolution, and retention operated?

Solve before reading [the worked solution](../solutions/001-multi-tenant-event-ingestion.md).
