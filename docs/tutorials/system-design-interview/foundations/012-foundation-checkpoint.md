---
tldr: "Tests readiness to move from foundational concepts into datastore, pattern, and case-study practice."
when_to_use: "Use after Lessons 001-011 and repeat until each answer is concrete and measurable."
---

# Foundation Checkpoint

Answer aloud without notes, then inspect the linked lesson for gaps.

## Requirements and estimates

1. Convert "fast and highly available" into two measurable NFRs.
2. Explain why an SLA and an SLO are not interchangeable.
3. Estimate peak requests/s from daily volume and a peak factor.
4. Use Little's Law to estimate in-flight work.
5. Explain how backlog age changes the scaling discussion.

## Data and correctness

1. Define an idempotency-key collision and the required response.
2. Explain why at-least-once delivery does not imply duplicate business effects.
3. Choose a partition key for per-account ordered evaluation and name its hot-key risk.
4. State when stale reads are acceptable and how staleness is bounded.
5. Describe additive schema evolution across mixed producers and consumers.

## Reliability, security, and Kubernetes

1. Separate readiness, liveness, and startup probes.
2. Explain why an autoscaler cannot replace capacity headroom.
3. Trace tenant identity through an asynchronous worker.
4. Distinguish an audit record from a diagnostic log.
5. Give an RPO, RTO, restore order, and reconciliation step.

## ML and migration

1. Separate model score, policy, action, and evidence.
2. Explain point-in-time feature correctness.
3. Name three evidence gates before a strangler cutover.
4. Explain why dual write is a transition, not a destination.
5. Give one reason to retain a modular monolith.

## Architecture sketch

In 15 minutes, design a tenant-aware event intake that acknowledges durably, preserves per-entity order, provides a
searchable projection, and survives consumer restarts. Include one normal sequence, one duplicate sequence, and one
overload response. You are ready for the cases when every box maps to a requirement and every failure has a bounded
outcome.
