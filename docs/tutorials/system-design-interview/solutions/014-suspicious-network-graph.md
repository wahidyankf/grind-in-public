---
tldr: "An immutable edge stream builds a temporal graph projection and bounded precomputed features with provenance."
when_to_use: "Use after attempting Case 014 to compare graph modelling, algorithms, and serving paths."
---

# Solution 014: Suspicious-Network Graph

## Requirement traceability

| Requirements       | Design response                                                                              |
| ------------------ | -------------------------------------------------------------------------------------------- |
| FR-1, FR-5         | Versioned node/edge facts with source id, valid interval, confidence, and supersession.      |
| FR-2, NFR-1, NFR-3 | Tenant-scoped graph projection with depth/fan-out/time/result/compute budgets.               |
| FR-3, NFR-2        | Stream and batch jobs precompute bounded features into an online key-value store.            |
| FR-4               | Every edge and derived feature retains contributing source references and algorithm version. |
| NFR-4              | Log-to-graph lag SLO plus faster stream feature updater.                                     |
| NFR-5              | Tenant partition/property filters and approved shared-node namespace.                        |

## Architecture

```text
domain events -> canonical edge log -> graph projector -> graph store -> investigation API
                         |                   |
                         +-> stream features +-> batch graph jobs
                                      \         /
                                       feature store -> decision API
```

Model an edge as `(tenant, from, type, to, valid_from, valid_to, confidence, source, version)`. Do not overwrite a
retracted edge; close its validity and append supersession. Shared reference nodes live in a policy-controlled namespace
and never make tenant-private nodes mutually visible.

## Algorithm placement

- BFS: bounded two-hop investigation neighbourhood; stop at depth, fan-out, and result limits.
- Union-find: efficient batch connected components for mostly additive snapshots; recompute after deletions because
  basic union-find cannot split.
- Dijkstra: minimum cumulative-risk paths when non-negative weighted edges have meaning; ordinary hop count uses BFS.
- degree/shared identifiers: incremental counters with supernode categories.
- centrality/community: batch only, versioned, and never presented as proof.

## Supernode defence

```text
ordinary device: expand <= 100 neighbours
shared terminal: return summary node, do not fan out by default
high-degree institution: traverse only typed/time-filtered edges
```

Estimate fan-out before execution, require filters for expensive shapes, and terminate server-side when the budget is
spent. Cache only authorization-scoped, versioned query results.

## Rebuild and operations

Projectors checkpoint log offsets and apply edge versions idempotently. Rebuild a parallel graph, verify node/edge
counts by type/time, sample neighbourhood hashes, then switch an alias. Monitor projection lag, orphan edges, supernode
growth, query budget exhaustion, feature freshness, and tenant-filter enforcement.

## Alternatives rejected

Real-time arbitrary traversal cannot offer a stable latency or abuse boundary. A graph database adds little for fixed
one-hop joins; retain relational edges until variable-depth traversal dominates. Union-find alone cannot answer temporal
paths or correct deletions.
