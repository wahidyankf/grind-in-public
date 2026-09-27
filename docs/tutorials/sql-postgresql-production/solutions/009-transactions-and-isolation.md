---
tldr: "Uses a locked invariant row or serializable transaction with whole-operation retry."
when_to_use: "Use after Drill 009."
---

# Solution 009: Transactions and Isolation

At read committed, each transaction counts two open rows, updates a different row, and both commit. One design locks an
account-level workflow row `FOR UPDATE`, rechecks, then closes: simple but serializes that account. Another uses
serializable isolation and retries SQLSTATE `40001`: more concurrency, but the transaction body must be repeatable and
retry rate monitored. A declarative constraint is preferable if the invariant can be remodelled into one row.
