---
tldr: "Design event-time rolling counts and sums for low-latency decisions under duplicates and late arrivals."
when_to_use: "Use to practise deques, buckets, stream state, watermarks, and exact-versus-approximate trade-offs."
---

# Case 004: Sliding-Window Velocity

Design a service that answers rolling features such as transaction count and amount for an account, device, and
counterparty over 1 minute, 1 hour, 24 hours, and 30 days.

### Functional requirements

| ID   | Requirement                                                                                |
| ---- | ------------------------------------------------------------------------------------------ |
| FR-1 | Ingest versioned events and deduplicate by stable event identity.                          |
| FR-2 | Return count, sum, distinct counterparties, and maximum amount for configured windows.     |
| FR-3 | Use event time, accept events up to ten minutes late, and correct affected future answers. |
| FR-4 | Rebuild one key/time range from the durable event source.                                  |
| FR-5 | Expose feature freshness and whether a result is exact or approximate.                     |

### Non-functional requirements

| ID    | Requirement                                                                                    |
| ----- | ---------------------------------------------------------------------------------------------- |
| NFR-1 | Serve 50,000 feature queries/s at 15 ms p99 from 200,000 events/s.                             |
| NFR-2 | Preserve per-key ordering and survive worker restart with at most 30 seconds of recovery lag.  |
| NFR-3 | Bound state despite 100 million active keys and extreme hot-key skew.                          |
| NFR-4 | Exact count/sum is required for short windows; 1% error is allowed for 30-day distinct counts. |
| NFR-5 | Tenant data and replay operations remain isolated.                                             |

## Assumptions and exclusions

The service computes features, not the final decision. Events later than ten minutes flow to offline correction and do
not rewrite already executed actions.

## Interview prompts

1. Compare a timestamp deque, ring of time buckets, prefix aggregates, and stream-processor state.
2. How do watermarks and allowed lateness affect answers?
3. How are max and distinct computed without storing every event for 30 days?
4. How do checkpoint, replay, query routing, and hot-key splitting work?

Solve before reading [the worked solution](../solutions/004-sliding-window-velocity.md).
