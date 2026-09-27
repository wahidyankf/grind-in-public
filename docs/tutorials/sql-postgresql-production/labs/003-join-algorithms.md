---
tldr: "Implements hash and merge joins in complete Python to expose executor invariants."
when_to_use: "Use to connect Python data structures with database join operators."
---

# Lab 003: Join Algorithms in Python

Implement `hash_join(accounts, transactions)` and `merge_join(accounts, transactions)` for `(tenant_id, account_id)`.
Return transaction id plus account status. Include type hints, duplicate keys, deterministic tests, and complexity. The
hash version may use a dictionary; the merge version receives sorted inputs.

Compare with the [solution](../lab-solutions/003-join-algorithms.md).
