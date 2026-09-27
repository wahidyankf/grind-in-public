---
tldr: "Walks clock-sweep to the first zero-usage reusable buffer."
when_to_use: "Use after completing Lab 009."
---

# Solution 009: Buffer Replacement and I/O

Index 0 decrements `2 -> 1`; index 1 is zero and becomes the candidate, subject to pin/dirty handling. The scan ring
limits cache pollution. Rising shared reads/read time and storage latency indicate misses/I/O; `wait_event_type = Lock`
and blocking PIDs indicate contention even when buffer hits are high.
