---
tldr: "Returns the three largest transactions per account with deterministic ties."
when_to_use: "Use to practise partitioned ranking."
---

# Drill 003: Ranking and Top-N

Return at most three transactions per account, ordered by amount descending and then transaction identity. Keep tied
amounts within the three-row limit rather than returning every tie.

[Reference solution](../solutions/003-ranking-and-top-n.md)
