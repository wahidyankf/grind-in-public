---
tldr:
  "Establish command, stop ambiguous writes, recover cohorts by authority map, and reconcile before normal operation."
when_to_use: "Use after Scenario 011 to assess incident command and migration-recovery reasoning."
---

# Debrief 011: Regional Incident during Migration

Assign incident commander, operations, old/new-stack investigation leads, communications, and scribe. Freeze migration
changes and stop routing new writes until authority is known. Retrieve the capability/tenant authority map and last
durable replication/outbox offsets.

Fence the old region with database credentials/leases and advance writer epochs before recovery promotion. Recover
critical cohorts separately: new-service tenants use its replicated state and log; monolith tenants use database
recovery. Do not mix a tenant across authorities because dashboards look healthy.

Route a small verified cohort, then expand. Reconcile accepted ids, sequence gaps, decisions/evidence, outbox/log
offsets, and projections. Communicate impact, uncertainty, mitigation, and next-update time on a fixed cadence.

Post-incident, improve authoritative dashboards, epoch enforcement, recovery automation, and a hybrid-state exercise.
Weak answers fail over immediately, debug without command roles, or trust Kubernetes readiness as data readiness.
