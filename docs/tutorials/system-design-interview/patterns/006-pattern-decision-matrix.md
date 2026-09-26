---
tldr: "Summarizes architecture-pattern triggers, evidence, costs, and safer starting alternatives."
when_to_use: "Use to review a design for unjustified complexity or to choose a deep dive."
---

# Pattern Decision Matrix

| Pressure                            | Candidate                      | Evidence to require                       | Simpler starting point             |
| ----------------------------------- | ------------------------------ | ----------------------------------------- | ---------------------------------- |
| DB change plus reliable event       | Transactional outbox           | measured event-loss consequence           | in-process call in one transaction |
| Independent read shapes             | CQRS projections               | query SLO/load mismatch                   | indexed relational query           |
| Long-running multi-owner workflow   | Orchestrated saga              | explicit intermediate states/compensation | one service transaction            |
| Full temporal reconstruction        | Event sourcing                 | history is authority, not diagnostics     | current state plus audit log       |
| Repeated expensive reads            | Cache-aside                    | hit rate and latency/cost benefit         | tune canonical query               |
| One tenant or key dominates         | isolation/cell/dedicated shard | skew and blast-radius data                | weighted quotas                    |
| Specialized text retrieval          | Search projection              | analyzer/query requirements               | database full-text/index           |
| Variable-depth relationship query   | Graph projection               | traversal dominates fixed joins           | relational edges + recursive query |
| Independent scale/release/ownership | Microservice extraction        | baseline coupling and team ownership      | modular monolith                   |
| Regional loss requirement           | multi-region data strategy     | explicit RPO/RTO and conflict semantics   | tested backup and single-region DR |

## Review questions

For every selected pattern, record:

```text
requirement -> pattern -> guarantee -> new failure -> detection -> recovery -> owner
```

If the team cannot name detection and recovery, it is adopting a failure mode without an operating plan. If the
requirement is hypothetical, delay the pattern and preserve an evolution seam instead.
