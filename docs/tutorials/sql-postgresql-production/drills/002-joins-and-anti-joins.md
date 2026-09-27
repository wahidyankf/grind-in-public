---
tldr: "Finds open accounts with no transaction in a time interval."
when_to_use: "Use to practise composite joins and null-safe anti-joins."
---

# Drill 002: Joins and Anti-Joins

Return open accounts for tenant 1 that had no transaction during January 2026, including customer name. Preserve
accounts that never had any transaction. Explain why `NOT IN` can be unsafe.

[Reference solution](../solutions/002-joins-and-anti-joins.md)
