---
tldr: "Uses contiguous sequences for traversal and hash tables for keyed lookup, counting, and deduplication."
when_to_use: "Use when a problem is dominated by indexed traversal, membership, frequency, or lookup by identity."
---

# Arrays, Strings, and Hashing

Python lists provide ordered traversal and constant-time indexed access; strings are immutable sequences; dictionaries
and sets provide hash-based lookup. Choose from the operations, not the nouns in the prompt.

## Production motivation

An alert aggregator receives rule identifiers and must report their counts in first-seen order. Repeatedly calling
`list.count` rescans the batch. A dictionary performs one pass and, in modern Python, preserves insertion order.

```python
from collections.abc import Iterable


def count_alerts(rule_ids: Iterable[str]) -> dict[str, int]:
    """Count alerts while preserving each rule's first-seen order."""
    counts: dict[str, int] = {}
    for rule_id in rule_ids:
        counts[rule_id] = counts.get(rule_id, 0) + 1
    return counts


assert count_alerts(["velocity", "country", "velocity"]) == {
    "velocity": 2,
    "country": 1,
}
```

```text
input:   velocity -> country -> velocity
counts:  {velocity:1}
         {velocity:1, country:1}
         {velocity:2, country:1}
```

The invariant is that `counts[key]` equals the number of occurrences in the consumed prefix. Time is `O(n)` expected;
space is `O(k)` for `k` distinct keys.

## Strings are data, not identity

Text matching in production usually needs a declared normalization contract. Lowercasing alone does not resolve Unicode
equivalence, punctuation, transliteration, or locale. The following example solves only whitespace and case:

```python
def normalize_label(value: str) -> str:
    """Normalize case and runs of whitespace for an internal label."""
    return " ".join(value.casefold().split())


assert normalize_label("  High   Risk ") == "high risk"
```

Do not reuse it for legal names without domain rules and evaluation data. Production screening separates normalization,
candidate retrieval, scoring, and human review because a false match has a different cost from a missed match.

## When not to use hashing

Hash tables do not provide sorted range queries, durable cross-process uniqueness, or bounded worst-case memory. Use a
sorted sequence for binary search, a database index for shared durable lookup, or a probabilistic structure when a
bounded-memory false-positive trade-off is acceptable.

## Checkpoint

Given ten million identifiers that cannot fit in memory, explain how partitioning or external sorting changes the
solution and which correctness guarantee must survive the change.
