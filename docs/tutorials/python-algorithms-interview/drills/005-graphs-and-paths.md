---
tldr: "Practises connected components and bounded weighted paths in relationship graphs."
when_to_use: "Use after graph traversal, connectivity, and shortest paths."
---

# Graphs and Paths

## Drill A: linked-entity component

Given undirected verified links, return the sorted component containing a requested entity. Include isolated known
entities. Reject links that reference unknown identifiers. Target `O(V + E)`.

## Drill B: lowest-cost explanation path

Given a directed graph with non-negative integer edge costs, return the cost and node path from source to target. Return
`None` when unreachable and reject negative weights.

Explain why online traversal needs a node/edge budget when the graph has supernodes.

[Solution](../solutions/005-graphs-and-paths.md)
