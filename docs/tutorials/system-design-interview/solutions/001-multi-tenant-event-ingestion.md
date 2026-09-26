---
tldr:
  "A cell-based ingestion design acknowledges a replicated log, preserves keyed order, and isolates replay and quotas."
when_to_use: "Use after attempting Case 001 to compare durability, partitioning, replay, and overload choices."
---

# Solution 001: Multi-Tenant Event Ingestion

## Requirement traceability

| Requirements      | Design response                                                                                    |
| ----------------- | -------------------------------------------------------------------------------------------------- |
| FR-1, FR-2, NFR-2 | Gateway validates envelope; ingestion reserves idempotency and acknowledges only after log quorum. |
| FR-3, NFR-1       | Log key is `tenant_id + entity_id`; oversized tenants receive dedicated partition ranges.          |
| FR-4, FR-5, NFR-5 | Retained canonical log, per-consumer offsets, quarantine topic, and object evidence archive.       |
| FR-6, NFR-4       | Per-tenant acceptance, rejection, duplicate, quota, lag, and oldest-age telemetry.                 |
| NFR-3, NFR-6      | Multi-zone replication, tenant-scoped auth/keys/quotas, and cell-local blast radius.               |

## Architecture

```text
online producer -> gateway -> schema/idempotency -> partitioned durable log
                                     |                    |
bulk object -> importer --------------+               +----+-------+
                                                         |    |   |
                                                      decide search archive
                                                         |
invalid / poison <-------------------------------- quarantine

control plane: schema registry, tenant placement, quotas, replay authorization
```

The durable log is the accepted-event authority for 30 days. An immutable object archive stores compacted canonical
events and original bulk objects for seven years. Search and analytics are rebuildable projections.

## Write sequence

```text
producer        gateway       idempotency DB       log quorum
   | POST(key,payload) |              |                |
   |------------------>| validate     |                |
   |                   | reserve(hash)|                |
   |                   |------------->|                |
   |                   | append(key,event)------------>|
   |                   |<---------------- durable offset|
   |                   | complete(result,offset)       |
   |                   |------------->|                |
   |<--- 202 event id + offset --------|                |
```

If the response is lost, the same key and payload returns the stored result. The same key with a different payload
returns conflict. A crash between append and `complete` leaves `PROCESSING`; reconciliation searches by event id and
completes it. This avoids appending twice without pretending two stores share one transaction.

## Partitioning and capacity

At 20,000 events/s and 2 KiB, raw ingress is about 39 MiB/s before replication. Benchmark 2,000 events/s per partition,
then start near 20 partitions rather than the mathematical minimum of 10. Partition by tenant/entity so entity order is
preserved. A placement directory assigns a dedicated namespace to tenants whose measured rate would dominate a shared
partition.

Bulk importers stream bounded chunks through the same canonical producer library. They use a separate quota and worker
pool so a 1 TB upload cannot consume online admission capacity.

## Consumer correctness

Each consumer writes an inbox id and its projection/effect in one local transaction, then advances the offset. Poison
events go to quarantine only after finite retries; the record includes exception class, schema version, offset, and
repair status. Replay starts a new consumer group with a scope token enforcing tenant and time bounds.

## Overload and operations

- reject over-quota tenants with `429` and jittered retry guidance;
- reject global saturation with `503` before accepting bytes that cannot be durably appended;
- page on SLO burn or oldest-age growth, not raw queue count alone;
- alert on partition skew, idempotency reconciliation backlog, quarantine age, archive gaps, and schema failures;
- restore by replaying canonical log/archive into empty projections and comparing bucket counts/hashes.

## Alternatives rejected

Directly writing every consumer database couples acknowledgement to the slowest consumer and makes partial success
unavoidable. A random partition key improves balance but violates keyed order. One global tenant partition preserves too
much order and creates hot keys.
