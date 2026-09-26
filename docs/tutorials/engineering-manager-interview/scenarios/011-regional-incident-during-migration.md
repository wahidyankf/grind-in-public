---
tldr: "Practise incident command when old and new stacks disagree during regional failure."
when_to_use: "Use for incident leadership, migration safety, recovery, and communication preparation."
---

# Scenario 011: Regional Incident during Migration

A region fails while 20% of tenants use a new service and the rest use the monolith. Replication lag is unclear, the old
region may still accept writes, and dashboards disagree. Customer impact is growing.

Explain role assignment, first mitigations, writer fencing, evidence needed for recovery, cohort routing,
reconciliation, communication cadence, and post-incident actions. Do not assume failover is safe merely because standby
Pods are ready.
