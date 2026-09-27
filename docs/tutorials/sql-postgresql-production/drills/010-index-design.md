---
tldr: "Designs one index for a tenant-scoped open-alert queue."
when_to_use: "Use for index column-order and partial-index interviews."
---

# Drill 010: Index Design

The hot query reads 100 oldest open alerts for one tenant and returns score/transaction id. Writes are heavy and closed
alerts dominate. Propose one index, explain ordering, covering, partial predicate, write cost, and a query it cannot
serve.

[Reference solution](../solutions/010-index-design.md)
