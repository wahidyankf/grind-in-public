---
tldr: "Uses a bounded heap for top-K and external merge sort for a complete order."
when_to_use: "Use after completing Lab 004."
---

# Solution 004: Top-K and External Sort

Top-K keeps a 100-item min-heap: `O(n log k)` time and `O(k)` memory. The full export creates memory-sized sorted runs,
writes them, then performs a k-way merge: `O(n log n)` comparisons plus temporary I/O.

```text
input -> [sorted run 1] --\
      -> [sorted run 2] ----> merge heads -> ordered output
      -> [sorted run 3] --/
```

`EXPLAIN ANALYZE` reports sort method and disk use; buffer/temp statistics expose temporary reads and writes.
