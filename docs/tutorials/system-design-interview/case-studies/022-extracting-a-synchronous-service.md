---
tldr: "Design extraction of a synchronous capability from a modular monolith with shadow and canary evidence."
when_to_use: "Use to practise service seams, API contracts, latency budgets, data authority, and rollback."
---

# Case 022: Extracting a Synchronous Service

After containerization, extract policy evaluation because it needs independent scaling and frequent releases. The
monolith currently calls an in-process module and reads shared policy tables.

### Functional requirements

| ID   | Requirement                                                                                       |
| ---- | ------------------------------------------------------------------------------------------------- |
| FR-1 | Provide a versioned evaluation API equivalent to the existing module contract.                    |
| FR-2 | Build service-owned policy state from authoritative publication events and backfill.              |
| FR-3 | Shadow requests and compare result, reason codes, version, and latency without changing outcomes. |
| FR-4 | Canary selected tenants, roll back routing, and later make the service authoritative.             |
| FR-5 | Remove legacy reads and code after an observed safety window.                                     |

### Non-functional requirements

| ID    | Requirement                                                                                  |
| ----- | -------------------------------------------------------------------------------------------- |
| NFR-1 | Add no more than 20 ms p99 to the caller's latency budget.                                   |
| NFR-2 | Meet 99.99% availability or provide an approved last-known-good fallback.                    |
| NFR-3 | No request may combine policy data from two versions.                                        |
| NFR-4 | Rollback before authority transfer takes under five minutes and loses no policy publication. |
| NFR-5 | The new team owns runtime, data, SLO, capacity, incident response, and deployment.           |

## Assumptions and exclusions

The rule language does not change during extraction. Other modules remain in the monolith.

## Interview prompts

1. Which contract, data, and dependency seams must exist before network extraction?
2. How are backfill and continuous policy events reconciled?
3. How does shadow comparison handle nondeterminism and version skew?
4. What are the canary, authority-transfer, rollback, and decommission gates?

Solve before reading [the worked solution](../solutions/022-extracting-a-synchronous-service.md).
