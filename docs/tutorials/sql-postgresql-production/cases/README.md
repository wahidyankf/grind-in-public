---
tldr: "Indexes ten production-shaped PostgreSQL cases with explicit functional and non-functional requirements."
when_to_use: "Use after the production lessons; timebox each case before reading its debrief."
---

# PostgreSQL Production Cases

Use 35 minutes: clarify, quantify, diagnose/design, handle failure, and summarize trade-offs.

1. [Query-plan regression](001-query-plan-regression.md)
2. [Cardinality misestimation](002-cardinality-misestimation.md)
3. [Pagination and indexes](003-pagination-and-indexes.md)
4. [Lock storm and deadlocks](004-lock-storm-and-deadlocks.md)
5. [Long transactions and bloat](005-long-transactions-and-bloat.md)
6. [Kubernetes connection storm](006-kubernetes-connection-storm.md)
7. [Billion-row migration](007-billion-row-migration.md)
8. [Replica staleness and failover](008-replica-staleness-and-failover.md)
9. [Partition lifecycle](009-partition-lifecycle.md)
10. [Tenfold growth](010-tenfold-growth.md)

The [debriefs](../debriefs/README.md) show one defensible response, not a unique answer.
