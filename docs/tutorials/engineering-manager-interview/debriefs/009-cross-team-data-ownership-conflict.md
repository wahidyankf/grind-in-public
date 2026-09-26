---
tldr: "Name one invariant and writer, treat replication as migration, and align team ownership with target authority."
when_to_use: "Use after Scenario 009 to assess data ownership and cross-team influence."
---

# Debrief 009: Cross-Team Data-Ownership Conflict

Define the invariant: one authoritative current risk status and one idempotent transition effect. Today the monolith is
writer; target authority may be the scoring service after evidence gates. The transition does not justify two writers.

```text
monolith writer -> outbox -> scoring shadow state -> compare
       |
       +-- authority until fenced cutover

after cutover: scoring writer -> compatibility event -> monolith read model
```

Backfill at a watermark, consume continuous changes, reconcile, shadow decisions, canary tenants, then fence the old
writer before granting a new epoch. Rollback before cutover returns routing; after cutover it requires fencing the new
writer before changing authority.

Align organizational ownership: the target team owns data contract, SLO, migration, on-call, and decommission; the
legacy team owns safe source publication until transfer. Set a transition expiry and executive tie-break only after
technical criteria are explicit. Weak answers accept permanent dual write or solve the org chart without solving data
authority.
