---
tldr: "Uses cell isolation, authoritative relational state, durable logs, evidence archives, and rehearsed recovery."
when_to_use: "Use after attempting Capstone 006."
---

# Debrief 006: Scale and Recover

Scale by evidence: tune access paths, primary and pools; separate asynchronous/analytical work; add replicas; partition
retention; then place tenant cohorts into cells. Each cell contains stateless Kubernetes workloads, PostgreSQL primary
and standbys, local log partitions, quotas, and SLOs. Global control maps tenant to region/cell and rolls schema/config;
immutable evidence and WAL archives cross the regional failure boundary within the five-minute RPO.

```text
tenant -> global placement -> region A / cell N -> online state + durable log
                     |                 |
                     |                 +-> encrypted archive/replication
                     v
              region B recovery capacity -> restore/promote -> replay -> verify -> route
```

Regional recovery freezes placement changes, fences old writers, restores/promotes eligible state, replays idempotent
work, verifies per-tenant positions/hashes, then shifts traffic gradually. Teams own cells end to end within common
platform guardrails. The EM funds capacity headroom, quarterly game days, restore evidence, cross-team interfaces, and a
roadmap triggered by measured saturation—not speculative technology adoption.
