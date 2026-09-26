---
tldr: "Indexes datastore choices by access pattern, correctness need, scale, and operating cost."
when_to_use: "Use after requirements and estimates, before naming a database in an interview answer."
---

# Datastores

Datastore selection begins with access patterns and invariants. Read in order:

1. [Relational systems](001-relational-systems.md)
2. [Wide-column systems](002-wide-column-systems.md)
3. [Document systems](003-document-systems.md)
4. [Caches and ephemeral state](004-caches-and-ephemeral-state.md)
5. [Durable logs and messaging](005-durable-logs-and-messaging.md)
6. [Search, graph, and analytics](006-search-graph-and-analytics.md)

```text
canonical truth -> transaction-oriented store or durable log
derived access   -> cache / search / graph / analytical projection
```

Do not let a projection silently become the only source of truth. Define rebuild, reconciliation, and freshness.
