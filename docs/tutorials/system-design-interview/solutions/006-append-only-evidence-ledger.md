---
tldr:
  "Partitioned append segments, immutable archive, signed roots, and rebuildable indexes provide verifiable evidence."
when_to_use: "Use after attempting Case 006 to compare integrity, queryability, retention, and recovery."
---

# Solution 006: Append-Only Evidence Ledger

## Requirement traceability

| Requirements       | Design response                                                                             |
| ------------------ | ------------------------------------------------------------------------------------------- |
| FR-1, NFR-1, NFR-2 | API appends canonical record to a replicated tenant/time segment before acknowledgement.    |
| FR-2, NFR-3        | Hot 90-day query projection; immutable object segments hold long retention.                 |
| FR-3, FR-4         | Per-segment hash chain, Merkle root, signed manifest, and independent root anchor.          |
| FR-5               | Retention catalogue, legal-hold overlay, approved deletion workflow, and deletion evidence. |
| NFR-4, NFR-5       | Separate writer/reader/admin roles, tenant envelope keys, and export quotas.                |

## Architecture

```text
application -> ledger API -> replicated append log -> segment builder -> immutable object store
                     |               |                    |
                     |               +-> query index      +-> root signer/anchor
                     `-> receipt(offset, record hash)
```

Do not demand a single global chain at 30,000 events/s. Partition by tenant and bounded time/sequence segment. Each
record hashes canonical bytes plus the prior hash. The manifest commits segment id, first/last sequence, count, chain
head, object hash, schema, and signing key id. A separately administered anchor store receives signed Merkle roots.

```text
segment 42
H0 -> H(record 1 || H0) -> H(record 2 || H1) -> ... -> Hn
                                                       |
manifest ----------------------------------------------+ -> signed root
```

This detects record mutation, internal deletion, reordering, and missing closed segments. Sequence monitoring detects
gaps; anchoring detects silent suffix replacement. It still cannot prove upstream facts were truthful.

## Query and export

An indexer writes tenant-scoped fields and evidence object locations to a searchable projection. Index lag is visible;
the append receipt remains available even before indexing. Export pins a set of closed segment manifests, filters into
new immutable objects, and produces hashes plus the source inclusion proofs and a signed export manifest.

## Retention and recovery

The retention catalogue decides when a segment may expire. Legal holds override expiry. If deletion is legally required,
delete encrypted objects and eligible keys according to policy, then append a separate non-sensitive deletion
certificate. Avoid pretending a broken hash chain is normal: segments are independently rooted so one expired segment
does not invalidate retained ones.

Replicate open logs across zones; copy closed encrypted objects to the approved recovery region. Restore manifests and
objects first, verify roots, rebuild query indexes, and reconcile counts/sequence ranges before reopening export.

## Alternatives rejected

Ordinary mutable database audit tables give convenient queries but weak tamper evidence and costly long retention. A
single blockchain adds consensus cost without solving input truth, privacy, authorization, or deletion.
