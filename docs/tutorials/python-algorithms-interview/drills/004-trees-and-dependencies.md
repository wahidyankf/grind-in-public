---
tldr: "Practises iterative tree traversal, prefix retrieval, and deterministic dependency ordering."
when_to_use: "Use after trees, tries, and graph ordering."
---

# Trees and Dependencies

## Drill A: bounded reply traversal

Given `parent_id -> child_ids`, return breadth-first identifiers from a root up to a maximum depth. Detect repeated
identifiers so malformed cyclic input cannot loop forever.

## Drill B: deterministic feature order

Given every feature and its prerequisites, return a deterministic topological order. Reject missing prerequisites and
cycles. When multiple features are ready, choose lexicographically.

State why deterministic order helps reproducible rule evaluation even when independent features could run concurrently.

[Solution](../solutions/004-trees-and-dependencies.md)
