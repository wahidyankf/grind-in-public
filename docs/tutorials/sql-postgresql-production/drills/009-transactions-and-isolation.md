---
tldr: "Protects a cross-row invariant under concurrent reviewers."
when_to_use: "Use for transaction and isolation interview scenarios."
---

# Drill 009: Transactions and Isolation

Two reviewers must never both close the last two open cases for one account concurrently. Describe the read-committed
race, then give two safe designs, their retry/locking behaviour, and rejection conditions.

[Reference solution](../solutions/009-transactions-and-isolation.md)
