---
tldr: "Block unsafe authorization and irreversible schema risk, then reduce scope or exposure with explicit ownership."
when_to_use: "Use after Scenario 003 to assess launch-risk judgment."
---

# Debrief 003: Quality versus Launch Date

Classify by consequence and reversibility. An unproved authorization boundary can disclose tenant data; it blocks launch
until a deterministic test and manual boundary verification pass. An unproved schema rollback can make the release
irreversible; use expand/contract or delay the mutation. Restore proof is required when the release changes canonical
state or recovery assumptions; otherwise document why existing restore evidence remains valid.

Options include reducing the feature, internal-only or one-tenant canary, read-only mode, or rescheduling. A feature
flag is useful only if it truly disables effects, is tested, observable, owned, and expires.

```text
impact x likelihood x detectability x reversibility -> decision authority
```

The manager makes trade-offs visible and ensures the appropriate owner accepts residual business risk; security/data
invariants are not silently waived. Weak answers either launch because the date exists or block everything without
offering a smaller safe outcome.
