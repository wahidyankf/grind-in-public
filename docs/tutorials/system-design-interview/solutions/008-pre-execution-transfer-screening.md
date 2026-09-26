---
tldr: "A durable screening state machine pins list and policy versions and issues a single-use clearance capability."
when_to_use: "Use after attempting Case 008 to compare workflow, matching, safe failure, and replay protection."
---

# Solution 008: Pre-Execution Transfer Screening

## Requirement traceability

| Requirements | Design response                                                                               |
| ------------ | --------------------------------------------------------------------------------------------- |
| FR-1, NFR-4  | Transfer-scoped idempotency and optimistic state version enforce one workflow.                |
| FR-2, NFR-1  | Parallel party normalization/candidate lookup; bounded scoring against pinned list snapshots. |
| FR-3, NFR-2  | Local policy combines matches and corridor rules; missing required state cannot return clear. |
| FR-4, NFR-3  | Review task references immutable evidence snapshot and exact versions.                        |
| FR-5, NFR-5  | Explicit rescreen transition; degraded mode returns pending/review rather than unsafe clear.  |

## State machine

```text
RECEIVED -> SCREENING -> CLEAR -----> CONSUMED
                |          \-------> EXPIRED
                +-> REVIEW -> CLEAR/BLOCK
                +-> BLOCK
                `-> PENDING_DEPENDENCY -> SCREENING

new list/policy while pending: REVIEW/PENDING -> RESCREENING -> ...
```

Every transition compares an expected state version in a relational transaction and appends evidence/outbox. A terminal
block is not overwritten; policy defines which states allow rescreening.

## Request path

```text
caller -> screening API -> pin policy/list versions
                             |
              +--------------+----------------+
              v              v                v
          party matcher  corridor rules  exception lookup
              +--------------+----------------+
                             v
                         policy result -> evidence DB -> clearance/review/block
```

Exact identifiers run first. Approximate text retrieval produces a capped candidate list, and field-aware scoring runs
only on candidates. Tenant thresholds and required fields are part of the pinned policy. Long inputs, tokens, aliases,
and candidate count are bounded to protect the 500 ms path.

## Clearance token

For `CLEAR`, issue a signed token containing transfer id, normalized request hash, tenant, result id, policy/list
versions, issued/expiry time, and nonce. Execution presents it once. The execution service atomically records nonce use
before transferring; changed amount/party data produces a different hash and cannot reuse clearance.

## Version activation and failure

Each request pins versions at `SCREENING` start. A list activation does not change it mid-flight. New policy may mark
unconsumed clearances for expiry/rescreening through an explicit event. Matcher/list/policy integrity failure can yield
`PENDING_DEPENDENCY` or `REVIEW`, never clear (`NFR-2`).

Monitor state age, dependency fallback, match candidate distribution, clearance expiry/replay attempts, review backlog,
and result rates by version. Reconcile one terminal screening record and at most one consumed clearance per transfer.

## Alternatives rejected

A stateless response token without server-side nonce consumption permits replay. Reading the "latest" list throughout
one request mixes versions. Treating similarity score alone as identity hides the policy and evidence boundary.
