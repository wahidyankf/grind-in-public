---
tldr:
  "Protect critical traffic immediately, make fairness explicit, and isolate bulk capacity while repairing the promise."
when_to_use: "Use after Scenario 008 to assess multi-tenant and stakeholder judgment."
---

# Debrief 008: Noisy Enterprise Tenant

Pause or throttle the replay before it causes broader failure, preserving its durable checkpoint so work resumes.
Protect synchronous decision and ingestion capacity using priority queues, independent consumer groups/pools, database
budgets, and per-tenant/global admission controls.

Explain impact to product and the tenant with measured throughput, backlog, and expected safe processing—not blame.
"Unlimited" is not operationally meaningful; align product/contract on a burst, sustained rate, completion objective,
and premium isolation option.

Longer term, isolate bulk workloads with separate Kubernetes Deployments, quotas/node capacity, queue partitions, and
oldest-age autoscaling. Very large tenants may receive a dedicated cell or partition, but only after cost and
operational evidence. Provide progress and resumption semantics.

Strong answers balance one strategic relationship with the service promise to all tenants. Weak answers let revenue
override system safety or reject the tenant with no recovery path or communication.
