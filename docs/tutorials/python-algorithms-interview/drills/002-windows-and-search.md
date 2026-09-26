---
tldr: "Practises rolling windows and monotonic boundary search under explicit ordering assumptions."
when_to_use: "Use after pointers, windows, prefix sums, and binary search."
---

# Windows and Search

## Drill A: rolling threshold breach

Given non-decreasing integer timestamps and a window size, return the first timestamp at which the number of events in
`(timestamp - window, timestamp]` reaches a threshold. Return `None` when it never does.

- Reject a non-positive window or threshold.
- Reject timestamps that move backwards.
- Target: `O(n)` time and `O(w)` space for the busiest active window.

## Drill B: version effective at time

Given rule versions sorted by unique `effective_from` timestamps, return the version active at a query timestamp. Reject
queries before the first version. Target `O(log n)` lookup after `O(n)` validation.

Explain why rebuilding the timestamp list on every lookup hides a linear cost.

[Solution](../solutions/002-windows-and-search.md)
