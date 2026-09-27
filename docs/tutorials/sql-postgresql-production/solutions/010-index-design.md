---
tldr: "Uses a small partial covering B-tree for the open queue."
when_to_use: "Use after Drill 010."
---

# Solution 010: Index Design

```sql
CREATE INDEX alerts_open_queue_idx
ON alerts (tenant_id, created_at, alert_id)
INCLUDE (score, transaction_id)
WHERE status = 'open';
```

Equality tenant leads, then queue order and a deterministic tie-breaker. The partial predicate avoids closed-row write
and storage cost. It cannot serve historical closed-alert search or efficient queries without tenant identity.
