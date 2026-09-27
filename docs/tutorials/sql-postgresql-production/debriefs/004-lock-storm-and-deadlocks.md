---
tldr: "Breaks the queue, aligns lock order, and retries only whole transient transactions."
when_to_use: "Use after attempting Case 004."
---

# Debrief 004: Lock Storm and Deadlocks

Use `pg_blocking_pids`, wait events, transaction age, and query/application identity to find the root blocker. Shed new
work and cancel the offending transaction under incident policy; never kill every waiter first. Make both paths lock
`case` then `alert` in sorted identity order, keep the transaction free of network calls, and preserve unique/FK/version
constraints. Retry `40P01` with jitter around the whole idempotent operation. Concurrency tests deliberately interleave
both paths and assert atomic state, bounded attempts, and visible conflicts.
