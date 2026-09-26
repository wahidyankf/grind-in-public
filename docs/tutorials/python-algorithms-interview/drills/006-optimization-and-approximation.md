---
tldr: "Practises optimal scheduling and a bounded-memory membership prefilter."
when_to_use: "Use after greedy, dynamic programming, and production algorithms."
---

# Optimization and Approximation

## Drill A: weighted review scheduling

Each review has `(start, end, value)`. Select non-overlapping reviews with maximum total value. Return both the value
and selected reviews. Reject invalid intervals. State the DP state, recurrence, and reconstruction method.

## Drill B: membership prefilter

Implement a small Bloom filter for strings using a byte array and deterministic standard-library hashes. Support `add`
and `might_contain`. Demonstrate that inserted values always return `True`; explain why non-inserted values may also do
so and why an authoritative lookup must follow a positive result.

[Solution](../solutions/006-optimization-and-approximation.md)
