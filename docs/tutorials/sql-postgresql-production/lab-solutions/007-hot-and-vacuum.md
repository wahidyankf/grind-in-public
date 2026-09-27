---
tldr: "Explains that HOT needs unchanged indexed values and space on the same heap page."
when_to_use: "Use after completing Lab 007."
---

# Solution 007: HOT and Vacuum

Changing only `assignee` is HOT-eligible if the same page has space. Changing indexed `status` is not. No-op assignment
still produces an update unless avoided, and a full page prevents HOT placement. A six-hour snapshot retains dead
versions and can grow heap/index work; identify the owner, cancel safely when policy permits, fix transaction scope, and
tune table vacuum/fillfactor from measured churn.
