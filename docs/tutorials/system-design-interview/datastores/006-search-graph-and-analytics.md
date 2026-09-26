---
tldr:
  "Treats search, graph, and analytical databases as purpose-built projections with freshness and rebuild contracts."
when_to_use: "Use for text/entity retrieval, relationship exploration, and aggregate analysis beyond an OLTP store."
---

# Search, Graph, and Analytics

These systems answer specialized queries efficiently, but they usually should not own the only canonical business
record.

## Search index

Use inverted indexes for token, phrase, fuzzy, filtered, and ranked retrieval.

```text
canonical change -> outbox/log -> indexer -> search index
                         |                       |
                         +-> checkpoint + lag <-+
```

Define analyzer and language per field, exact-key subfields, tenant filters, freshness SLO, and reindex strategy. Alias
swaps enable building a new version before cutover. Reject search for transaction invariants or unbounded joins.

## Graph

Graph storage fits variable-depth traversal over explicit relationships: accounts sharing devices, addresses, or
counterparties. Model edge meaning, direction, confidence, time validity, and provenance.

```text
[account A] --uses--> [device X] <--uses-- [account B]
     |                                        |
   pays                                     pays
     v                                        v
[merchant M] <----------- shared ---------- [address Z]
```

Bound traversal depth and fan-out. Precompute common neighbourhood features for synchronous decisions. Reject a graph
database if all queries are fixed one-hop joins that a relational index serves cheaply.

## Analytics

Columnar analytical storage scans selected columns, compresses repeated values, and supports aggregates across large
history. Feed it from durable change streams or snapshots, partition by common filters, and separate interactive serving
from batch workloads.

Lambda-style separate batch and speed implementations can diverge. A simpler architecture uses one durable log,
streaming projections for freshness, and periodic canonical reconciliation. Choose based on latency, replay cost, and
team capacity.

## Projection contract

Every derived store needs:

- source of truth and version;
- checkpoint and observed lag;
- idempotent updater;
- full rebuild procedure;
- reconciliation metric;
- query fallback or declared outage behaviour.

## References

- [OpenSearch documentation](https://docs.opensearch.org/latest/)
- [Neo4j documentation](https://neo4j.com/docs/)
