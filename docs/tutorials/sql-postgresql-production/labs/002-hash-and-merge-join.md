---
tldr: "Chooses between hash and merge joins from size, memory, skew, and ordering."
when_to_use: "Use after the join-algorithm lesson."
---

# Lab 002: Hash and Merge Join

Join 5 million transactions to 100,000 accounts. Scenario A has no useful order and the account hash fits memory.
Scenario B has both inputs ordered by `(tenant_id, account_id)` but memory allows only half the account hash. Predict
the algorithm, big-O work, spill risk, and effect of one tenant owning 70% of rows.

Compare with the [solution](../lab-solutions/002-hash-and-merge-join.md).
