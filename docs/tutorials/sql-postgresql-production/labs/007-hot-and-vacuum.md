---
tldr: "Predicts HOT eligibility, dead-tuple growth, and vacuum constraints."
when_to_use: "Use after MVCC and vacuum lessons."
---

# Lab 007: HOT and Vacuum

`cases` has indexes on its primary key and `status`; `assignee` is not indexed. For updates to `assignee`, `status`, and
an unchanged row on a full page, identify HOT eligibility. Then explain what a six-hour transaction does to vacuum and
the operational response.

Compare with the [solution](../lab-solutions/007-hot-and-vacuum.md).
