---
tldr: "Design temporal graph ingestion, bounded traversal, feature computation, and investigation queries."
when_to_use: "Use to practise graph modelling, connected components, centrality, projection, and abuse bounds."
---

# Case 014: Suspicious-Network Graph

Design a graph product connecting accounts, parties, devices, addresses, institutions, and transfers so investigators
and scoring services can identify suspicious networks.

### Functional requirements

| ID   | Requirement                                                                                      |
| ---- | ------------------------------------------------------------------------------------------------ |
| FR-1 | Ingest versioned nodes and time-bounded, provenance-rich edges idempotently.                     |
| FR-2 | Query a bounded neighbourhood with filters on time, edge type, confidence, and tenant.           |
| FR-3 | Compute reusable features such as degree, shared identifiers, components, and bounded path risk. |
| FR-4 | Explain every returned connection through source evidence.                                       |
| FR-5 | Apply corrections and edge retractions without erasing historical provenance.                    |

### Non-functional requirements

| ID    | Requirement                                                                           |
| ----- | ------------------------------------------------------------------------------------- |
| NFR-1 | Serve two-hop investigation queries within 2 seconds p95 on 5 billion edges.          |
| NFR-2 | Serve precomputed online graph features within 20 ms p99.                             |
| NFR-3 | Bound traversal depth, fan-out, result size, CPU, and memory per request.             |
| NFR-4 | New edges become queryable within five minutes and online features within one minute. |
| NFR-5 | Enforce tenant isolation and explicit policy for any shared reference nodes.          |

## Assumptions and exclusions

Graph signals inform a decision; they do not independently prove wrongdoing. Arbitrary unbounded graph queries are out
of scope.

## Interview prompts

1. Define node/edge identity, direction, time, confidence, and provenance.
2. Which queries belong in a graph projection versus precomputed features?
3. Compare BFS, union-find, shortest path, and batch centrality for the requirements.
4. How are supernodes, updates, retractions, and graph/index rebuilds handled?

Solve before reading [the worked solution](../solutions/014-suspicious-network-graph.md).
