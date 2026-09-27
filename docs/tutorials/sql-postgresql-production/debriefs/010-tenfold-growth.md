---
tldr: "Separates authoritative workflow, immutable history, and tenant cells with measurable evolution triggers."
when_to_use: "Use after attempting Case 010."
---

# Debrief 010: Tenfold Growth

```text
global tenant directory + schema/capacity control plane
                |
       +--------+---------+
       v                  v
cell A: API + PostgreSQL  cell B: API + PostgreSQL  ... <= 5% tenants each
       |                  |
       +-> outbox/log ----+----> object evidence + analytical/query systems
```

First repair queries/statistics/indexes, scale primary storage/compute, bound pools, isolate exports, and add read
replicas. Keep transactional workflow/configuration relational; stream immutable evidence to object/analytical stores.
Partition for lifecycle, then introduce tenant cells when write/recovery/blast-radius thresholds demand it. Moves use a
placement state machine, copy/catch-up, verification, brief write fence, routing switch, and rollback window. Kubernetes
owns stateless elasticity and controlled disruption; database automation still requires fencing, backups, restore
drills, capacity, and on-call ownership. Regional recovery consumes archived WAL and tested routing, matching stated
RPO.
