---
tldr: "Explains why memory favours hash join and existing order favours merge join."
when_to_use: "Use after completing Lab 002."
---

# Solution 002: Hash and Merge Join

Scenario A favours building the 100,000-row account hash and probing 5 million rows: expected `O(n + m)`. Scenario B
favours a streaming merge, also `O(n + m)`, because order is already available and hash would batch/spill. Heavy tenant
skew can create hot hash buckets, uneven parallel work, and bad estimates; extended statistics and tenant-aware
partitioning may help.
