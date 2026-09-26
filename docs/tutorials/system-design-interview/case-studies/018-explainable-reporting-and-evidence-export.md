---
tldr: "Design reproducible reports and large evidence exports without harming online workloads."
when_to_use: "Use to practise snapshots, asynchronous jobs, object storage, signatures, and authorization."
---

# Case 018: Explainable Reporting and Evidence Export

Design reporting and export for decisions, alerts, cases, and their evidence. Reports may be interactive summaries or
large regulator/auditor packages.

### Functional requirements

| ID   | Requirement                                                                                 |
| ---- | ------------------------------------------------------------------------------------------- |
| FR-1 | Query permitted aggregate metrics by tenant, time, policy, result, and workflow dimensions. |
| FR-2 | Create an asynchronous export from a declared snapshot and filter specification.            |
| FR-3 | Include records, lineage, reason codes, manifests, checksums, and schema documentation.     |
| FR-4 | Authorize request, generation, download, and sharing separately and audit each action.      |
| FR-5 | Cancel, expire, regenerate, and verify an export without exposing partial output.           |

### Non-functional requirements

| ID    | Requirement                                                                                 |
| ----- | ------------------------------------------------------------------------------------------- |
| NFR-1 | Interactive summaries return within 3 seconds p95 for a 90-day range.                       |
| NFR-2 | Export up to 1 TB within 12 hours with resumable work and bounded online impact.            |
| NFR-3 | A package is internally consistent to its snapshot and reproducible from recorded versions. |
| NFR-4 | Encrypt outputs, use short-lived downloads, and enforce retention/legal holds.              |
| NFR-5 | Isolate tenant data, quotas, worker pools, and object prefixes.                             |

## Assumptions and exclusions

Exports are machine-readable plus a generated summary; arbitrary business-intelligence authoring is out of scope.

## Interview prompts

1. Where does report data come from and how fresh is it?
2. How is a consistent snapshot identified across several projections?
3. How are jobs partitioned, checkpointed, assembled, verified, and published atomically?
4. How are authorization changes handled after an export was created?

Solve before reading [the worked solution](../solutions/018-explainable-reporting-and-evidence-export.md).
