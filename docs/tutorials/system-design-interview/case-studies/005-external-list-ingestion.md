---
tldr: "Design versioned ingestion and serving of large external entity lists with provenance and atomic publication."
when_to_use: "Use to practise batch validation, diffing, object storage, indexing, and rollback."
---

# Case 005: External-List Ingestion

Providers publish full and incremental files containing entities, aliases, identifiers, dates, and provenance. Design a
pipeline that validates, normalizes, versions, and serves those lists to matching systems.

### Functional requirements

| ID   | Requirement                                                                                |
| ---- | ------------------------------------------------------------------------------------------ |
| FR-1 | Fetch or receive signed provider artifacts and preserve original bytes and metadata.       |
| FR-2 | Validate schema, signature, counts, referential rules, and suspicious volume changes.      |
| FR-3 | Normalize records without losing source fields or provenance.                              |
| FR-4 | Publish one immutable version atomically and support rollback to a prior version.          |
| FR-5 | Expose exact identifier lookup, candidate text search, version diff, and ingestion status. |

### Non-functional requirements

| ID    | Requirement                                                                               |
| ----- | ----------------------------------------------------------------------------------------- |
| NFR-1 | Process 50 million records within two hours while serving the previous version.           |
| NFR-2 | Make a published version visible to all serving replicas within 60 seconds.               |
| NFR-3 | Never expose a partially built version; preserve provenance for seven years.              |
| NFR-4 | Detect missing, duplicate, malformed, and unexpectedly deleted records before activation. |
| NFR-5 | Limit provider credentials and tenant subscriptions by least privilege.                   |

## Assumptions and exclusions

Human policy decides which providers and lists apply to each tenant. Entity-match scoring is Case 009.

## Interview prompts

1. How are raw, canonical, and serving representations separated?
2. How is a full snapshot reconciled with incremental updates?
3. How is an index built, validated, activated, and rolled back without downtime?
4. Which checks prevent a syntactically valid but catastrophically incomplete feed from publishing?

Solve before reading [the worked solution](../solutions/005-external-list-ingestion.md).
