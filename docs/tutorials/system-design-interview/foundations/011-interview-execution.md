---
tldr: "Provides a precise 45-minute interview script, communication patterns, and common failure corrections."
when_to_use: "Use for timed practice and immediately before a system-design interview."
---

# Interview Execution

Drive the conversation. State assumptions, invite correction, and keep a visible list of requirements and unresolved
questions. The interviewer should never wonder why a box appeared.

## Board layout

```text
+------------------+----------------------+------------------+
| FR / NFR / scope | architecture         | deep-dive notes  |
| estimates        | and main data flow   | failures / ops   |
+------------------+----------------------+------------------+
```

## Minute-by-minute script

### 0-5: clarify

Ask actors, top three behaviours, latency path, scale, data sensitivity, consistency, regions, and excluded features.
Repeat the prioritized scope: "I will optimize the synchronous decision and evidence path; reporting is asynchronous."

### 5-10: estimate

Calculate peak request rate, payload volume, concurrency, retention, and read/write ratio. Name the dominant pressure.
If numbers are absent, make round assumptions and label them.

### 10-22: broad design

Define public contracts and canonical entities, then draw the end-to-end path. Describe the write before secondary
views. Connect each major box to an `FR` or `NFR`.

### 22-35: deep dives

Choose two risks. Follow one request and one failure with sequence diagrams. Discuss idempotency, partition key,
consistency, data ownership, backpressure, and recovery.

### 35-42: operations and evolution

Cover SLO measurement, alerts, capacity, security, deployment, rollback, backup/restore, and the next 10x. State what
you deliberately did not distribute.

### 42-45: recap

Trace requirements to design choices, name the top trade-off and remaining risk, and explain one likely evolution.

## Communication phrases

- "I need to clarify this because it changes the acknowledgement point."
- "This store is a projection; the durable log remains the replay source."
- "I am preserving per-account order, not global order."
- "During dependency failure, the safer domain action is review, so availability degrades intentionally."
- "I would validate this assumption with a workload-specific benchmark."

## Common corrections

| Weak move                          | Correction                                                   |
| ---------------------------------- | ------------------------------------------------------------ |
| Naming products immediately        | State access pattern and guarantee first                     |
| Drawing only a happy path          | Trace timeout, duplicate, overload, and regional failure     |
| Claiming exactly-once              | Define effect deduplication and transaction boundary         |
| Saying "eventual consistency"      | Define stale read, convergence, conflict, and bounded impact |
| Splitting everything into services | Require an ownership, scaling, security, or release reason   |
| Ignoring the migration             | Give states, evidence gates, rollback, and decommissioning   |

## Practice rule

Record a timed answer. On review, count unsupported nouns: every database, queue, cache, service, region, and algorithm
must have a stated reason and a rejection case.
