---
tldr: "Connects Python, algorithms, PostgreSQL, system design, and Engineering Manager preparation into one programme."
when_to_use: "Use to choose prerequisites, sequence courses, and finish with cross-course interview capstones."
---

# Engineering Interview Programme

This programme turns five courses into one learning graph. It targets Engineering Manager interviews where technical
depth, architecture judgment, delivery, reliability, and communication are evaluated together.

```text
Python foundations -> algorithms --------+
       |                                  |
       +-> SQL/PostgreSQL ----------------+-> system design -> EM judgment
                        |                 |          |
                        +-> performance --+----------+-> capstones
```

## Learning path

1. [Python Production Foundations](../python-production-foundations/README.md) for typed application boundaries.
2. [Python Algorithms Interview](../python-algorithms-interview/README.md) for invariants and complexity.
3. [SQL and PostgreSQL Production](../sql-postgresql-production/README.md) for authoritative data and operations.
4. [System Design Interview](../system-design-interview/README.md) for requirements, distributed correctness,
   Kubernetes, and migration cases.
5. [Engineering Manager Interview](../engineering-manager-interview/README.md) for team, delivery, and incident
   judgment.

Do not wait to finish every page before practising. After each technical topic, explain the customer consequence,
production signal, owner, rollout, rollback, and rejected alternative.

## Capstones

Each prompt requires functional and non-functional requirements, an ASCII architecture, complete Python/SQL teaching
slices, failure analysis, migration/operation steps, and an Engineering Manager decision narrative.

The [capstone prompt index](capstones/README.md) and [debrief index](debriefs/README.md) keep the exercises separate
from their reference answers.

1. [Tenant-safe schema, queries, and indexes](capstones/001-tenant-safe-data-path.md)
2. [Idempotent Python transaction and outbox](capstones/002-idempotent-outbox.md)
3. [Diagnose a slow decision query](capstones/003-slow-query-incident.md)
4. [Alert and case evidence platform](capstones/004-alert-case-platform.md)
5. [Extract database ownership from a modular monolith](capstones/005-database-ownership-extraction.md)
6. [Scale and recover through tenfold growth](capstones/006-scale-and-recover.md)

Attempt each before using the [debriefs](debriefs/README.md). In a real interview, compress the same reasoning into the
available time; the long form here builds the mental model first.
