---
tldr: "Budgets connections globally, bounds pod pools, and separates readiness from database overload."
when_to_use: "Use after attempting Case 006."
---

# Debrief 006: Kubernetes Connection Storm

The naive maximum is `80 * 4 * 10 = 3,200`, far beyond the 540 application budget. Allocate budgets per workload and
derive a small per-process pool, or place a transaction-pooling proxy with its own hard backend limit.

```text
pods -> bounded local acquisition -> pool proxy -> <= 540 application sessions -> primary
                       |               |
                timeout/backpressure   +-> queue/wait metrics
```

Use startup jitter, exponential backoff, max surge limits, disruption budgets, and admission control. Readiness should
show whether a pod can serve, but aggressive probes must not turn a brief database pause into a restart storm. Run
migrations as a coordinated job with lock timeout. Alert on pool wait, acquisition timeouts, DB sessions, connection
rate, and overload rejections.
