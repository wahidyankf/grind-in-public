---
tldr: "Makes state and intent atomic while accepting at-least-once delivery and idempotent consumers."
when_to_use: "Use after attempting Capstone 002."
---

# Debrief 002: Idempotent Outbox

The API transaction reserves `(tenant_id, idempotency_key, request_hash)`, creates the case, inserts an outbox row, and
stores the response. A different hash conflicts. Workers claim ordered batches with `FOR UPDATE SKIP LOCKED`, publish,
then mark success. A crash after publish causes replay, so each consumer transaction inserts a unique event id beside
its effect.

```text
API -> [idempotency + case + outbox] COMMIT
                              |
workers --claim batches-------+-> broker -> [consumer inbox + effect] COMMIT
       \-> retries -> quarantine -> repair/replay
```

Scale workers on oldest-event age, not raw count alone. Bound retries with jitter and separate tenant quotas. Tests
crash at every arrow: before/after commit, after publish, before consumer acknowledgement, and during replay. The team
owns API correctness and dashboards; platform owns broker/operator; both share an incident runbook and quarterly
recovery game.
