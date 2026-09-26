---
tldr: "Extract policy evaluation through a hardened seam, replicated owned state, shadow comparison, and cohort canary."
when_to_use: "Use after attempting Case 022 to compare synchronous extraction and authority transfer."
---

# Solution 022: Extracting a Synchronous Service

## Requirement traceability

| Requirements | Design response                                                                                   |
| ------------ | ------------------------------------------------------------------------------------------------- |
| FR-1, NFR-1  | Existing in-process interface becomes a typed network contract with deadline and bounded payload. |
| FR-2, NFR-3  | Immutable policy publication events build service-owned artifacts; requests pin one version.      |
| FR-3         | Adapter invokes old authority and sampled new shadow, then records semantic/latency comparison.   |
| FR-4, NFR-4  | Stable tenant cohorts route through a reversible switch before authority transfer.                |
| FR-5         | Safety-window evidence gates removal of legacy reads, code, permissions, and tables.              |
| NFR-2, NFR-5 | Last-known-good fallback plus full team ownership of service/data/SLO/on-call.                    |

## Migration architecture

```text
                         +-> in-process evaluator (authority)
monolith decision adapter|
                         `-> network evaluator (shadow, no effect)
                                        ^
policy DB transaction -> outbox -> log -> owned artifact builder
                   historical snapshot -> backfill/reconcile
```

Before adding the network, make the module interface explicit and stop callers from reading its tables. Record a golden
contract corpus with inputs, version, result, reasons, and failure cases. The service contract carries request id,
tenant, policy version or active-pointer token, context, and deadline.

## Data convergence

Take a policy snapshot at outbox watermark `W`, load it into the service, then consume publication events after `W`.
Compare tenant/version counts, artifact hashes, active pointers, and sampled evaluations. Only the publication path may
write policy authority; the service owns compiled serving artifacts, not authoring truth.

## Shadow and canary

```text
shadow: return old; compare new result/reasons/version/latency asynchronously
canary: selected tenant -> new authority; old may shadow for comparison
expand: cohort gates -> more tenants
final: new authority -> remove old path after safety window
```

Normalize nondeterministic metadata before comparison and always pin identical policy/reference versions. Gates include
zero unexplained action mismatch on required corpus, bounded reason-code differences, p99 latency budget, fallback rate,
and capacity at peak.

## Failure and rollback

Give the remote call at most 20 ms added p99 through in-region placement, persistent connections, compact contracts, and
in-memory artifacts. The caller propagates its deadline. During service failure, use an explicitly approved local
last-known-good evaluator only while version compatibility is valid; otherwise return safe review/reject.

Before authority transfer, rollback is routing to the in-process module. After transfer, rollback means service version
or active artifact, not two policy writers. Preserve compatibility events to the monolith until all readers migrate.

## Alternatives rejected

Giving both systems policy-write authority creates conflicts. Calling the old database from the new service preserves
shared ownership. Extracting a chatty function rather than the evaluation capability would multiply network latency.
