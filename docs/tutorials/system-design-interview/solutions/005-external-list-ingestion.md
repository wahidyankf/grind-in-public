---
tldr: "A staged snapshot pipeline validates raw artifacts, builds immutable indexes, and activates them by pointer."
when_to_use: "Use after attempting Case 005 to compare ingestion, validation, publication, and rollback."
---

# Solution 005: External-List Ingestion

## Requirement traceability

| Requirements       | Design response                                                                                     |
| ------------------ | --------------------------------------------------------------------------------------------------- |
| FR-1, NFR-3        | Immutable raw object plus signed manifest preserves provider bytes, checksum, and receipt metadata. |
| FR-2, NFR-4        | Staged validation gates schema, signature, counts, uniqueness, references, and change thresholds.   |
| FR-3               | Versioned normalizer emits canonical records with source fields and transformation lineage.         |
| FR-4, NFR-1, NFR-2 | Build/index immutable candidate off-path; activate one signed version pointer atomically.           |
| FR-5, NFR-5        | Exact-key store, search projection, diff dataset, status API, and least-privilege subscriptions.    |

## Pipeline

```text
provider -> quarantine object -> signature/checksum -> parse/normalize -> canonical snapshot
                                                               |               |
                                                     validation report      +----+----+
                                                                            |         |
                                                                        key index  search index
                                                                            \         /
                                                                             verify
                                                                               |
                                                                        active pointer
```

Every run has an immutable `ingestion_id`. Workers write deterministic chunks keyed by provider record id. A manifest
lists chunk hashes, counts, schema/normalizer versions, provider metadata, and parent version for incrementals.

## Full and incremental feeds

Treat the latest verified full snapshot as a base. Apply incremental upserts/deletes into a new logical version; never
mutate the active version. Periodically reconcile the materialized result with the next full snapshot. Deletion records
remain tombstoned in provenance even if omitted from the serving index.

Validation compares total/add/change/delete counts with historical bands and hard policy limits. A feed that parses but
drops 80% of records fails activation and requires approval; syntax alone is insufficient.

## Publication sequence

```text
builder       validators       serving replicas        pointer store
   | complete     |                    |                     |
   |------------>| checks             |                     |
   |<------ signed approval -----------|                     |
   | announce version ---------------->| download/warm/hash  |
   |<---------------- readiness -------|                     |
   | activate ------------------------------------------------>|
   |                         replicas atomically swap pointer  |
```

The previous version remains available for rollback. Cache keys include list version. Queries declare the version used,
so in-flight work remains reproducible during activation.

## Scale and operations

Partition 50 million records into deterministic chunks and parallelize CPU-bound normalization and indexing in
Kubernetes Jobs with bounded concurrency. Avoid loading the full snapshot into memory. Alerts cover fetch lateness,
signature failure, validation deltas, chunk retries, readiness convergence, index count/hash mismatch, and active
version age.

## Alternatives rejected

Updating the live search index record by record exposes partial state and makes rollback ambiguous. Treating an
incremental provider stream as authoritative forever allows silent drift; full-snapshot reconciliation is required.
