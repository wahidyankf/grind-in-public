---
tldr: "Design one monitoring product across low-latency streaming and complete batch reconciliation."
when_to_use: "Use to practise dual paths, event time, replay, rule versions, and result convergence."
---

# Case 013: Batch and Stream Monitoring

Design transaction monitoring that produces prompt alerts from a live stream and later runs complete daily evaluations
over corrected source data.

### Functional requirements

| ID   | Requirement                                                                           |
| ---- | ------------------------------------------------------------------------------------- |
| FR-1 | Evaluate live transactions against versioned rules and features.                      |
| FR-2 | Run a daily complete evaluation using finalized data and the declared policy version. |
| FR-3 | Correlate equivalent live and batch findings without duplicate cases.                 |
| FR-4 | Explain whether a finding came from live, batch, correction, or replay.               |
| FR-5 | Reprocess a tenant/date after data or rule correction with auditable supersession.    |

### Non-functional requirements

| ID    | Requirement                                                                                      |
| ----- | ------------------------------------------------------------------------------------------------ |
| NFR-1 | Live findings appear within 30 seconds for 100,000 events/s.                                     |
| NFR-2 | Daily processing of 5 billion records finishes within six hours.                                 |
| NFR-3 | Live path may use bounded approximation; batch path must produce complete deterministic results. |
| NFR-4 | Online workloads retain their SLO while batch consumes spare or isolated capacity.               |
| NFR-5 | Reconciliation exposes missing, extra, and semantically different findings.                      |

## Assumptions and exclusions

The source event log and finalized analytical snapshot are both available. Case workflow is Case 016.

## Interview prompts

1. Share rule semantics without maintaining two divergent implementations.
2. Define finding identity across live, batch, and replay.
3. How do event-time lateness and source corrections converge?
4. How are compute isolation, checkpointing, partial rerun, and reconciliation operated?

Solve before reading [the worked solution](../solutions/013-batch-and-stream-monitoring.md).
