---
tldr: "Uses bounded time partitions, tenant-leading local indexes, and detach-based retention."
when_to_use: "Use after attempting Case 009."
---

# Debrief 009: Partition Lifecycle

Range-partition by event month or day based on measured partition size and maintenance time; pre-create future ranges
and keep a monitored default partition for late/misrouted rows. Local indexes lead with tenant and match time/id access.
Because global uniqueness outside the partition key is unavailable, include time bucket in identity or maintain a small
deduplication authority. Detach an expired partition, verify/archive it, then drop it—metadata/lifecycle work replaces
row-by-row deletes. Partitioning remains one database failure domain; move to tenant cells when write, recovery, or
blast-radius limits are exceeded.
