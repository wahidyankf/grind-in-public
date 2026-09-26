---
tldr:
  "One versioned rule semantics feeds fast streaming findings and complete batch findings joined by stable identity."
when_to_use: "Use after attempting Case 013 to compare convergence, replay, and workload isolation."
---

# Solution 013: Batch and Stream Monitoring

## Requirement traceability

| Requirements       | Design response                                                                                       |
| ------------------ | ----------------------------------------------------------------------------------------------------- |
| FR-1, NFR-1        | Partitioned stream evaluates compiled rules and incremental features within 30 seconds.               |
| FR-2, NFR-2, NFR-3 | Batch engine evaluates finalized snapshots with identical versioned rule intermediate representation. |
| FR-3               | Stable finding identity and upsert/supersession prevent duplicate cases.                              |
| FR-4               | Findings record mode, run, source watermark, rule/data versions, and evidence.                        |
| FR-5, NFR-5        | Scoped rerun writes a new run and reconciliation classifies missing/extra/different.                  |
| NFR-4              | Separate queues, Kubernetes quotas/node pools, database pools, and admission priorities.              |

## Architecture

```text
                  rule authoring -> typed compiled rule IR
                                      /            \
events -> stream features -> live evaluator        batch evaluator <- finalized snapshot
                    |               \              /
                    |                finding sink
                    |                     |
                    +-------------- reconciliation -> alert/case dedup
```

Generate or interpret the same bounded rule intermediate representation in both engines. Test a golden corpus against
both implementations before publication; do not copy rule logic into unrelated codebases.

## Finding identity

Derive identity from tenant, rule semantic id, primary subject, normalized observation window, and correlation key—not
from run id. A result row contains multiple observations:

```text
finding F
  live observation:  run L17, provisional evidence
  batch observation: run B04, complete evidence
  current status: confirmed / superseded / live-only / batch-only
```

The case creator acts idempotently on `F`. Reprocessing adds an observation and may supersede status; it does not create
another case unless policy explicitly treats the period as a new episode.

## Event time and convergence

Live state has watermarks and allowed lateness; later records become corrections. Batch pins a finalized source snapshot
and policy version. Reconciliation joins by finding id and compares result plus canonical evidence hash. Live-only can
be valid approximation or stale data; batch-only may expose lateness or live defect. Each category has an owner and
repair path.

## Capacity and operations

Five billion records in six hours is about 231,500 records/s before retries and skew. Partition batch by tenant/time and
measure per-rule cost. Pause bulk partitions when online saturation or database pools cross thresholds. Checkpoint
immutable output partitions so retry does not restart the day.

Monitor live watermark/lag, batch completion forecast, rule cost, skew, reconciliation categories, case-dedup results,
and shared dependency saturation.

## Alternatives rejected

Two independently authored rule engines will drift. Making live processing wait for finalized data violates timeliness.
Treating live and batch findings as unrelated doubles investigator work and loses convergence evidence.
