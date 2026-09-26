---
tldr: "Design extraction of asynchronous alert generation using outbox, backfill, shadowing, and reconciliation."
when_to_use: "Use to practise event contracts, CDC transitions, idempotent consumers, and cutover watermarks."
---

# Case 023: Extracting an Event-Driven Service

Extract alert generation from the modular monolith. Today a transaction commits and invokes an in-process detector;
alerts share the monolith database.

### Functional requirements

| ID   | Requirement                                                                                             |
| ---- | ------------------------------------------------------------------------------------------------------- |
| FR-1 | Publish a canonical transaction event for every committed source transaction.                           |
| FR-2 | Backfill historical events and continuously process new events without gaps or duplicate alert effects. |
| FR-3 | Run the extracted detector in shadow and reconcile findings with the legacy detector.                   |
| FR-4 | Transfer alert write authority at a declared watermark and preserve downstream compatibility.           |
| FR-5 | Replay a bounded range and supersede prior findings with full lineage.                                  |

### Non-functional requirements

| ID    | Requirement                                                                        |
| ----- | ---------------------------------------------------------------------------------- |
| NFR-1 | New alerts appear within 30 seconds p99 at 100,000 source events/s.                |
| NFR-2 | No committed source transaction is silently omitted from the durable event stream. |
| NFR-3 | Processing is at-least-once with idempotent business effects and per-entity order. |
| NFR-4 | Shadow processing cannot overload the source database or production broker.        |
| NFR-5 | Cutover and rollback preserve an unambiguous single alert writer.                  |

## Assumptions and exclusions

An outbox can be added to the monolith; CDC may bootstrap existing changes but is not the desired domain contract.

## Interview prompts

1. Place source update and outbox insert in one transaction.
2. How are snapshot/backfill and live stream joined at a consistent watermark?
3. Define alert identity, deduplication, ordering, and semantic comparison.
4. How are authority, rollback, replay, and old-table retirement controlled?

Solve before reading [the worked solution](../solutions/023-extracting-an-event-driven-service.md).
