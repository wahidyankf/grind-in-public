---
tldr: "Selects durable logs and queues through retention, ordering, delivery, fan-out, backpressure, and replay needs."
when_to_use: "Use when producers and consumers need temporal decoupling or a replayable event history."
---

# Durable Logs and Messaging

Messaging decouples availability and pace only when the buffer is durable, bounded operationally, and consumed with
explicit effect semantics.

## Queue and log

```text
work queue                       partitioned durable log
producer -> [messages] -> worker producer -> P0: e1 e4 -> group A offset
                                  \-------> P1: e2 e3 -> group B offset
```

A work queue commonly assigns one message to one worker in a consumer set. A durable log retains ordered partition
records so independent consumer groups can replay. Real products overlap; decide from semantics, not label.

## Partition key

Key by the smallest entity needing order. Random keys maximize distribution but lose locality. A tenant key preserves
tenant order but lets one tenant dominate a partition. Monitor skew and provide an isolation path for outliers.

## Consumer contract

1. Validate schema and tenant context.
2. Start an idempotent effect transaction.
3. Apply the state change and write any outgoing event atomically when possible.
4. Commit the effect.
5. Acknowledge or advance the offset.

Retry transient failures with backoff. Send poison records to a quarantine stream only after preserving context and
alerting; a dead-letter queue without a replay owner becomes silent data loss.

## Retention and replay

Retention must cover outage recovery, consumer deployment rollback, and required backfills. Replaying into a current
consumer can produce different results if rules or reference data changed, so record versions and choose whether replay
means "rebuild historical truth" or "evaluate under current policy."

## Reject when

Reject asynchronous messaging when a caller needs a single atomic response and the queued workflow cannot express the
business contract. Reject a log as a query database; build named projections.

## References

- [Apache Kafka documentation](https://kafka.apache.org/documentation/)
