---
tldr: "Practise resolving competing data-authority claims during service extraction."
when_to_use: "Use for cross-team influence, migration, data governance, and decision-right preparation."
---

# Scenario 009: Cross-Team Data-Ownership Conflict

Two teams both plan to update customer risk status. One owns the current monolith table; another owns a new scoring
service. Each argues its path must remain available during migration, creating proposed permanent dual writes.

Explain how you establish the invariant, choose temporary and target authority, coordinate outbox/backfill/
reconciliation, define cutover and rollback, resolve organizational ownership, and prevent the transition from becoming
permanent.
