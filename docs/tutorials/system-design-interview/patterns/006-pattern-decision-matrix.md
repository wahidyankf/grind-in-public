---
tldr: "Summarizes architecture-pattern triggers, evidence, costs, and safer starting alternatives."
when_to_use: "Use to review a design for unjustified complexity or to choose a deep dive."
---

# Pattern Decision Matrix

Each design pressure, with its candidate pattern, the evidence to require first, and a simpler starting point:

- **DB change plus reliable event**
  - Candidate: Transactional outbox
  - Evidence to require: measured event-loss consequence
  - Simpler starting point: in-process call in one transaction
- **Independent read shapes**
  - Candidate: CQRS projections
  - Evidence to require: query SLO/load mismatch
  - Simpler starting point: indexed relational query
- **Long-running multi-owner workflow**
  - Candidate: Orchestrated saga
  - Evidence to require: explicit intermediate states/compensation
  - Simpler starting point: one service transaction
- **Full temporal reconstruction**
  - Candidate: Event sourcing
  - Evidence to require: history is authority, not diagnostics
  - Simpler starting point: current state plus audit log
- **Repeated expensive reads**
  - Candidate: Cache-aside
  - Evidence to require: hit rate and latency/cost benefit
  - Simpler starting point: tune canonical query
- **One tenant or key dominates**
  - Candidate: isolation/cell/dedicated shard
  - Evidence to require: skew and blast-radius data
  - Simpler starting point: weighted quotas
- **Specialized text retrieval**
  - Candidate: Search projection
  - Evidence to require: analyzer/query requirements
  - Simpler starting point: database full-text/index
- **Variable-depth relationship query**
  - Candidate: Graph projection
  - Evidence to require: traversal dominates fixed joins
  - Simpler starting point: relational edges + recursive query
- **Independent scale/release/ownership**
  - Candidate: Microservice extraction
  - Evidence to require: baseline coupling and team ownership
  - Simpler starting point: modular monolith
- **Regional loss requirement**
  - Candidate: multi-region data strategy
  - Evidence to require: explicit RPO/RTO and conflict semantics
  - Simpler starting point: tested backup and single-region DR

## Review questions

For every selected pattern, record:

```text
requirement -> pattern -> guarantee -> new failure -> detection -> recovery -> owner
```

If the team cannot name detection and recovery, it is adopting a failure mode without an operating plan. If the
requirement is hypothetical, delay the pattern and preserve an evolution seam instead.
