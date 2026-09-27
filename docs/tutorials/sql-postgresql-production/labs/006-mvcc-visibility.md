---
tldr: "Practises snapshot visibility across insert and update transaction identifiers."
when_to_use: "Use after MVCC to reason about concurrent observations."
---

# Lab 006: MVCC Visibility

Snapshot S sees committed transaction IDs through 120. Tuple v1 has `xmin=90, xmax=130`; v2 has `xmin=130, xmax=0`.
Which version does S see before and after transaction 130 commits? How does the answer differ under read committed and
repeatable read? State why a reader may delay cleanup without blocking the writer.

Compare with the [solution](../lab-solutions/006-mvcc-visibility.md).
