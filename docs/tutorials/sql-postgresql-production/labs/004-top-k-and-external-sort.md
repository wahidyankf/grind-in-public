---
tldr: "Derives top-K heap and external merge-sort choices under memory limits."
when_to_use: "Use after sort algorithms to reason about LIMIT and spills."
---

# Lab 004: Top-K and External Sort

You need the largest 100 of 50 million scores with 64 MiB work memory, then a complete ordered export. Choose an
algorithm for each, give time/space complexity, draw external runs and merging, and identify the plan evidence of a
spill.

Compare with the [solution](../lab-solutions/004-top-k-and-external-sort.md).
