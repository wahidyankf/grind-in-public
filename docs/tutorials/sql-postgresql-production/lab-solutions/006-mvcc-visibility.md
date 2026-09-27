---
tldr: "Shows old-snapshot visibility and isolation-level snapshot refresh."
when_to_use: "Use after completing Lab 006."
---

# Solution 006: MVCC Visibility

S sees v1: transaction 130 is outside its snapshot, so the deletion and replacement are invisible. Under repeatable
read, S continues seeing v1 after 130 commits. Under read committed, a later statement obtains a new snapshot and sees
v2. The writer creates a version rather than overwriting v1; vacuum retains v1 while any active snapshot may need it.
