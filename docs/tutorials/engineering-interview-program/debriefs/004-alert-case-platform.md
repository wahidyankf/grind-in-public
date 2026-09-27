---
tldr: "Separates durable ingestion, transactional workflow, search projection, and immutable evidence."
when_to_use: "Use after attempting Capstone 004."
---

# Debrief 004: Alert and Case Platform

```text
events -> durable log -> detection -> PostgreSQL alerts/cases + outbox
                         |                    |
policy/version store ----+                    +-> search projection
                                              +-> immutable object evidence
review UI -> workflow API --------------------+
```

PostgreSQL owns case state, assignments, exact version references, transitions, and outbox. A partitioned log absorbs
alert peaks and replay; a search projection serves flexible queues; object storage retains signed evidence packages.
Correlation uses tenant/entity hash grouping, time windows, and bounded graph expansion. Composite tenant keys, RLS,
partial queue indexes, optimistic case versions, and append-only audit events protect workflow.

Deploy stateless APIs/workers across zones with bounded pools, quotas, disruption budgets, and lag-based autoscaling.
Trace policy/model/data versions in every decision. Deliver in vertical slices: authoritative workflow, audit/export,
then correlation/search. Each slice has threat model, load test, restore/replay drill, on-call ownership, and rollback.
