---
tldr: "Provides complete typed Python models of hash and merge joins."
when_to_use: "Use after implementing Lab 003 yourself."
---

# Solution 003: Join Algorithms in Python

```python
from __future__ import annotations

from collections import defaultdict
from collections.abc import Iterable
from dataclasses import dataclass


@dataclass(frozen=True, order=True)
class Account:
    tenant_id: int
    account_id: int
    status: str


@dataclass(frozen=True, order=True)
class Transaction:
    tenant_id: int
    account_id: int
    transaction_id: int


def hash_join(accounts: Iterable[Account], transactions: Iterable[Transaction]) -> list[tuple[int, str]]:
    index: dict[tuple[int, int], list[Account]] = defaultdict(list)
    for account in accounts:
        index[(account.tenant_id, account.account_id)].append(account)
    return [
        (transaction.transaction_id, account.status)
        for transaction in transactions
        for account in index[(transaction.tenant_id, transaction.account_id)]
    ]


def merge_join(accounts: list[Account], transactions: list[Transaction]) -> list[tuple[int, str]]:
    result: list[tuple[int, str]] = []
    i = 0
    j = 0
    while i < len(accounts) and j < len(transactions):
        account = accounts[i]
        transaction = transactions[j]
        account_key = (account.tenant_id, account.account_id)
        transaction_key = (transaction.tenant_id, transaction.account_id)
        if account_key < transaction_key:
            i += 1
        elif account_key > transaction_key:
            j += 1
        else:
            result.append((transaction.transaction_id, account.status))
            j += 1
    return result


accounts = [Account(1, 1, "open"), Account(1, 2, "frozen")]
transactions = [Transaction(1, 1, 10), Transaction(1, 2, 20)]
assert sorted(hash_join(accounts, transactions)) == [(10, "open"), (20, "frozen")]
assert sorted(merge_join(accounts, transactions)) == [(10, "open"), (20, "frozen")]
```

Both are expected `O(n + m)` after ordering is available. The schema makes account keys unique; repeated transaction
keys remain beside each other and all produce an output. A many-to-many merge groups equal runs and emits their
Cartesian product.
