---
tldr:
  "An outbox closes the commit/event gap; snapshot plus watermark seeds an idempotent shadow consumer before authority
  moves."
when_to_use: "Use after attempting Case 023 to compare asynchronous extraction, backfill, reconciliation, and cutover."
---

# Solution 023: Extracting an Event-Driven Service

## Requirement traceability

| Requirements | Design response                                                                       |
| ------------ | ------------------------------------------------------------------------------------- |
| FR-1, NFR-2  | Source transaction inserts canonical outbox row beside transaction commit.            |
| FR-2, NFR-3  | Snapshot at watermark plus ordered live events feeds an inbox-idempotent consumer.    |
| FR-3, NFR-4  | Shadow writes isolated findings; reconciliation compares stable semantic identities.  |
| FR-4, NFR-5  | Writer lease/epoch and cutover watermark establish exactly one alert authority.       |
| FR-5         | Scoped replay uses run id and supersession lineage rather than destructive overwrite. |
| NFR-1        | Partitioned log/workers and lag-based autoscaling meet 30-second freshness.           |

## Source publication

```text
monolith DB transaction
  +-- INSERT/UPDATE transaction
  `-- INSERT outbox(event_id, aggregate_id, sequence, schema, payload)
commit
   |
relay claims rows -> durable partitioned log -> service inbox + detector + finding/outbox
```

The relay may publish twice after a crash; stable `event_id` and consumer inbox handle it. Partition by tenant/entity to
preserve the detector's required order. Monitor unpublished outbox age as a correctness signal.

## Snapshot and stream join

1. Record database snapshot and outbox watermark `W` under a consistent boundary.
2. Export source rows in deterministic chunks with ids and source versions.
3. Load the extracted service's shadow state idempotently.
4. Consume log events strictly after `W`.
5. Wait until shadow lag reaches steady state and reconcile counts/hashes/samples.

If the database cannot expose such a boundary, CDC can capture the log position around the snapshot, but convert storage
changes into a versioned domain event contract at the boundary.

## Finding equivalence and cutover

Define semantic finding id from tenant, detector semantic id, subject, and observation window. Compare existence,
action, reason set, evidence hash, and version. Classify mismatches; do not require byte equality for harmless ordering
metadata.

```text
legacy writer epoch 7 -----------|
                                 | fence at source watermark C
new writer shadow ---------------|----> authority epoch 8
```

At `C`, stop/fence legacy alert writes, let it finish through `C`, verify new service consumed through `C`, then grant
the new writer epoch. Downstream accepts only current epoch. Rollback after this point fences epoch 8 before restoring a
new legacy epoch; never activate both.

## Replay and retirement

Replay writes a separate run, links new findings to superseded ones, and triggers downstream effects idempotently. After
the safety window, remove legacy detector code, alert writes/tables, CDC, flags, permissions, dashboards, and
compatibility consumers. Archive required evidence and update ownership maps.

## Alternatives rejected

Application dual writes can commit one side only. Backfill without a watermark creates gaps or overlap ambiguity.
Deleting old findings during replay destroys lineage and can duplicate downstream cases.
