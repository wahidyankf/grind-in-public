---
tldr: "Practises B-tree lookup, leaf splitting, and range scans."
when_to_use: "Use after the B-tree lesson to recover its invariants by hand."
---

# Lab 001: B-Tree Search and Split

Given order-4 leaves `[10 20 30] <-> [50 60]`, insert `40`, then search `[25, 55]`. Draw the split and parent separator.
State search and output complexity, why leaves link, and two production costs of random inserts. Then propose the index
for tenant-scoped newest-first cursor pagination and identify its usable prefixes.

Compare with the [solution](../lab-solutions/001-btree-search-and-split.md).
