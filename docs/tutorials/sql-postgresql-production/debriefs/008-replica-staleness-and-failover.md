---
tldr: "Uses a synchronous recovery candidate, LSN-aware reads, and fenced promotion."
when_to_use: "Use after attempting Case 008."
---

# Debrief 008: Replica Staleness and Failover

Use at least one synchronous in-region standby to meet zero-RPO acknowledgement, with an explicit availability trade-
off if no synchronous candidate remains. Return a commit LSN/session consistency token; route immediate reads to the
writer or a replica replayed through that LSN. Reject or reroute search when lag exceeds 10 seconds.

```text
detect -> prove/fence old primary -> choose most advanced eligible standby -> promote
      -> invalidate pools/DNS -> verify writes and jobs -> rebuild old members
```

Measure detection, fence, promotion, client reconnection, data verification, and catch-up during quarterly drills.
