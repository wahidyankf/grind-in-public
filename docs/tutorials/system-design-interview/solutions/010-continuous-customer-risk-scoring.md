---
tldr:
  "A keyed event processor maintains versioned customer state and atomically publishes serving scores and transitions."
when_to_use: "Use after attempting Case 010 to compare incremental scoring, recomputation, correction, and lineage."
---

# Solution 010: Continuous Customer Risk Scoring

## Requirement traceability

| Requirements       | Design response                                                                              |
| ------------------ | -------------------------------------------------------------------------------------------- |
| FR-1, NFR-1, NFR-3 | Customer-keyed log and idempotent state machine apply events in sequence.                    |
| FR-2, NFR-2        | Tenant/customer serving key holds active score, band, factors, freshness, and lineage.       |
| FR-3               | State transaction writes score plus uniquely identified band-transition outbox event.        |
| FR-4, NFR-4        | Recompute into a versioned shadow namespace, compare, then atomically switch active pointer. |
| FR-5               | Reversible feature contributions or bounded customer replay handle corrections/retractions.  |
| NFR-5              | Tenant partitions, quotas, separate bulk pool, and immutable historical score evidence.      |

## Architecture

```text
source events -> customer-keyed log -> stream scorer -> state/checkpoint
                                             |             |
                                             +-> score store + outbox -> review
historical objects -> recompute workers -> shadow score namespace -> compare/switch
```

The aggregate stores source version/vector, incremental features, policy/model version, last event sequence, score, and
band. Event ids deduplicate delivery. Per-customer partitioning preserves order; source-specific sequence or version
detects gaps and stale updates.

## Transition atomicity

```text
begin transaction
  verify event not processed
  update feature state and score
  if band changed: insert transition(id=customer+score_version+band+effective_time)
  mark event processed
commit; then acknowledge
```

The transition id prevents replay from creating another review. A later correction emits a superseding score record; it
does not rewrite the historic decision trail.

## Recompute

Pin source snapshot, feature code, policy, and model. Partition customers, checkpoint completion, and write a shadow
namespace. Continue live changes into both active and shadow after the snapshot watermark. Compare row counts, hashes,
score/band distributions, selected full records, and downstream transition simulation. Switch only a small tenant cohort
first; keep old namespace during rollback window.

Batch runs under separate Kubernetes Jobs, node pools/quotas, and database concurrency limits. Online lag and read SLOs
win over recompute completion.

## Failures and observability

Alert on partition lag, event gap, state-store checkpoint age, score freshness, transition-outbox age, shadow drift, hot
customers, and tenant recompute progress. A poison event is quarantined with the affected customer marked stale rather
than advancing silently.

## Alternatives rejected

Recomputing full history on every event violates latency and cost. Mutating the active score table during bulk recompute
exposes mixed versions. Partitioning only by tenant creates a severe hot partition for large customers.
