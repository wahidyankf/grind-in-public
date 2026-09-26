---
tldr:
  "Versioned analytical projections serve summaries, while snapshot-pinned resumable jobs publish signed export
  packages."
when_to_use:
  "Use after attempting Case 018 to compare reporting freshness, consistency, authorization, and bulk isolation."
---

# Solution 018: Explainable Reporting and Evidence Export

## Requirement traceability

| Requirements | Design response                                                                                   |
| ------------ | ------------------------------------------------------------------------------------------------- |
| FR-1, NFR-1  | Tenant-partitioned columnar projection and pre-aggregates serve bounded 90-day queries.           |
| FR-2, NFR-2  | Durable export job partitions a pinned snapshot, checkpoints chunks, and uses a bulk worker pool. |
| FR-3, NFR-3  | Manifest records query, snapshot/watermarks, schemas, versions, files, counts, and checksums.     |
| FR-4         | Authorization at create/run/download/share plus append-only audit.                                |
| FR-5         | Temporary prefix remains private; verify all chunks before atomic manifest publication.           |
| NFR-4, NFR-5 | Tenant encryption/prefixes, short-lived access, lifecycle retention, quotas, and legal holds.     |

## Architecture

```text
canonical events -> analytical projector -> columnar tables/pre-aggregates -> report API
       |                     |
       |                     `-> snapshot catalogue
       v                                |
evidence objects <- export workers <- durable job/partitions <- authorized request
       |
signed final manifest -> short-lived authorized download
```

## Snapshot semantics

The job pins a source snapshot id or a vector of projection watermarks plus rule/model/schema versions. If multiple
projections cannot provide one transactional snapshot, the manifest declares each watermark and the system waits until
they all cover the requested cutoff. This is reproducible bounded consistency, not a false global transaction.

## Job state

```text
REQUESTED -> AUTHORIZED -> RUNNING -> VERIFYING -> PUBLISHED -> EXPIRED
                              |           |
                              +-> FAILED <-+
                              `-> CANCELLED
```

Partition by tenant/time/hash into deterministic chunk ids. Each worker writes a temporary immutable object and records
row count/hash transactionally. Retry sees the completed chunk. The assembler verifies expected partitions, counts,
hashes, and schema, then writes and signs the final manifest. Only that manifest makes the package visible.

## Security and operations

The execution identity receives only the job's tenant/snapshot scope. Download authorization is rechecked because the
requester may have lost access. Object URLs are short-lived and every download is audited. Bulk pools have independent
CPU, I/O, database concurrency, and tenant quotas; online SLOs preempt them.

Monitor projection freshness, report scan bytes, job age/progress, retry hot spots, checksum failure, temporary-object
age, download authorization, and expiry deletion.

## Alternatives rejected

Running 1 TB exports against the OLTP database harms online transactions. Streaming partial output directly to the user
cannot resume or prove completeness. Treating authorization only at request time ignores later role changes.
