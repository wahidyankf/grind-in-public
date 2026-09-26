---
tldr: "Applies outbox, inbox, saga, CQRS, event sourcing, and reconciliation to explicit consistency boundaries."
when_to_use: "Use when one business change crosses a database and asynchronous consumers."
---

# Messaging and Consistency

## Transactional outbox and inbox

```text
producer DB transaction                 consumer DB transaction
+ domain row                            + inbox(event_id unique)
+ outbox row                            + projection/effect
      |                                        |
      v                                        v
 relay -> broker -- at least once --> consumer ack
```

The outbox prevents "database committed, event lost." The inbox prevents a duplicate delivery from repeating the local
effect. Neither makes all services one transaction; downstream state is temporarily behind and must expose freshness.

## Saga

A saga coordinates local transactions with compensating actions:

```text
reserve funds -> reserve inventory -> finalize
      |                |
      |                `-- fail -> release funds
      `-- fail -> no later steps
```

Compensation is a new business action, not time travel. It can fail and needs idempotency and operator repair. Reject a
saga for a hard invariant that cannot tolerate intermediate states.

## CQRS

Separate a write model enforcing invariants from read projections optimized for named queries. This is useful when read
shapes differ greatly from writes or need independent scale. It adds lag, projection deployment, and rebuild work.
Reject it for ordinary CRUD whose relational queries already meet the SLO.

## Event sourcing

Event sourcing treats domain events as authoritative state and folds them into current state. It enables temporal
reconstruction and new projections but demands stable event semantics, upcasters, snapshots, privacy handling, and
careful side-effect separation.

```text
events: Opened -> ItemAdded -> Submitted -> Approved
fold(events) => current aggregate
```

Do not call an audit log event sourcing unless the application reconstructs authority from it. Reject event sourcing
when history is not a core requirement and the team cannot support event evolution.

## Reconciliation

Asynchronous systems need periodic proof:

- compare counts by tenant and time bucket;
- compare commutative checksums or canonical hashes;
- validate invariants such as one terminal result per accepted command;
- sample full records;
- produce repair work with a stable identifier and audit record.

Reconciliation is not an excuse for weak primary design. It bounds residual uncertainty and detects silent drift.
