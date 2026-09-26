---
tldr: "Practises stable deduplication and grouped counting with ordered sequences and hash tables."
when_to_use: "Use after the arrays, strings, and hashing lesson."
---

# Sequences and Hashing

## Drill A: stable event deduplication

An ingestion batch is a sequence of `(event_id, payload)` pairs. Return the first pair for each identifier, preserving
encounter order. Reject a repeated identifier when its payload differs from the first payload.

- Empty input returns an empty list.
- Identical duplicates are ignored.
- Conflicting duplicates raise `ValueError` naming the identifier.
- Target: `O(n)` expected time and `O(k)` additional space for `k` identifiers.

```text
[(a,x), (b,y), (a,x)] -> [(a,x), (b,y)]
[(a,x), (a,z)]        -> ValueError for a
```

Explain why a database uniqueness constraint is still needed across batches.

## Drill B: first-seen grouped totals

Given `(tenant, amount)` pairs, return a dictionary of integer totals in each tenant's first-seen order. Reject negative
amounts. State the invariant and the difference between local aggregation and an authoritative billing total.

[Solution](../solutions/001-sequences-and-hashing.md)
