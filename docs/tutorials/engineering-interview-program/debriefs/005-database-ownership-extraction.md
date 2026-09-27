---
tldr: "Uses strangler routing, replicated history, shadow reads, and a reversible write fence."
when_to_use: "Use after attempting Capstone 005."
---

# Debrief 005: Database Ownership Extraction

Extract only after the module has stable ownership, contract tests, independent scaling/reliability need, and a team
ready to operate it. The new service exclusively owns new alert tables; no shared-table writes survive.

```text
monolith module -> outbox/change stream -> new alert database
       |                                  |
       +-> current reads       shadow <---+
                                      verify counts/hashes/semantics

cutover: prepare -> catch up -> fence old writes -> switch route -> verify -> observe
rollback: route back while compatibility and reverse reconciliation window remain
```

Backfill stable key ranges, replay changes idempotently, compare tenant/time buckets and semantic samples, then shadow
reads. A short old-write fence establishes one cutover position. Avoid dual-write as permanent architecture. Roll out
pods with bounded connections and versioned APIs/events. The EM defines shared success metrics, one incident commander,
decision owners, staffing/on-call readiness, and explicit abort thresholds for lag, mismatch, or p99.
