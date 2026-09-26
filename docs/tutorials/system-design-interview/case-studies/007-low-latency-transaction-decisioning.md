---
tldr: "Design a synchronous decision path combining features, rules, models, policy, and durable evidence."
when_to_use: "Use to practise latency budgets, dependency failure, explainability, and correctness under retries."
---

# Case 007: Low-Latency Transaction Decisioning

Design an API that receives a proposed transaction and returns `allow`, `review`, or `reject` with reason codes before
the caller continues.

### Functional requirements

| ID   | Requirement                                                                              |
| ---- | ---------------------------------------------------------------------------------------- |
| FR-1 | Validate and idempotently evaluate a transaction against active tenant policy.           |
| FR-2 | Retrieve entity, velocity, list-match, rule, and model signals needed by that policy.    |
| FR-3 | Return action, reason codes, decision id, policy/model versions, and evidence reference. |
| FR-4 | Persist a reproducible decision record and publish it for downstream investigation.      |
| FR-5 | Support shadow policy/model evaluation without changing the returned action.             |

### Non-functional requirements

| ID    | Requirement                                                                                |
| ----- | ------------------------------------------------------------------------------------------ |
| NFR-1 | Sustain 8,000 requests/s with 150 ms p99 end-to-end in-region.                             |
| NFR-2 | Provide 99.99% monthly availability with an explicit safe fallback per missing dependency. |
| NFR-3 | A retry with the same idempotency key and payload returns the same executed decision.      |
| NFR-4 | Acknowledge no decision until its evidence is durably recoverable.                         |
| NFR-5 | Isolate tenant policies, features, models, evidence, and capacity.                         |

## Assumptions and exclusions

The caller supplies a deadline and authenticated tenant identity. Investigation UI and model training are separate
cases.

## Interview prompts

1. Allocate the 150 ms latency budget and identify parallel calls.
2. Which data is precomputed, cached, or fetched synchronously?
3. How does each dependency fail: fail-open, fail-closed, rule-only, or review?
4. Where is idempotency reserved and evidence committed relative to the response?
5. How are canary, shadow comparison, overload, and tenant isolation operated?

Solve before reading [the worked solution](../solutions/007-low-latency-transaction-decisioning.md).
