---
tldr:
  "A versioned probabilistic identity graph produces bounded reputation features while preserving privacy and
  correction."
when_to_use: "Use after attempting Case 012 to compare identity resolution, graph projection, TTL, and abuse defence."
---

# Solution 012: Device Identity and Reputation

## Requirement traceability

| Requirements       | Design response                                                                                         |
| ------------------ | ------------------------------------------------------------------------------------------------------- |
| FR-1, NFR-1, NFR-2 | Exact approved keys plus bounded similarity model resolve to versioned device clusters with confidence. |
| FR-2               | Append time-bounded, provenance-rich observations; graph projection derives relationships.              |
| FR-3               | Low-latency reputation row stores precomputed features, freshness, version, and reasons.                |
| FR-4               | Validated outcomes update weighted evidence; scheduled decay expires old contribution.                  |
| FR-5, NFR-3        | Data catalogue, short raw TTL, pseudonymous ids, unlink/delete workflow, and minimal audit record.      |
| NFR-4, NFR-5       | Cardinality/rate limits, supernode handling, tenant namespaces, and explicit shared-data contracts.     |

## Data flow

```text
observation -> validate/pseudonymize -> identity resolver -> device cluster version
                         |                    |
                         v                    v
                    append store -------> graph edges
                                              |
validated outcomes -> reputation updater -> feature store -> online query
```

Exact stable device keys, when policy permits them, map directly. Probabilistic observations produce weighted edges and
a confidence score. Merge creates a new cluster version with parent lineage; split reassigns future authority and emits
correction events. Historic decisions keep the cluster version they used.

## Online versus batch

Online queries read precomputed features: accounts seen in 24 hours, negative-outcome weight, age, geography changes,
and shared-device category. Batch jobs compute components, supernode classification, and longer graph features. A shared
public terminal or carrier network is a supernode and must not create millions of naive suspicious edges.

Reputation is a time-decayed weighted sum whose terms retain provenance. Decay jobs can use time buckets rather than
touching every device continuously.

## Privacy and abuse

HMAC approved identifiers with a rotating tenant-scoped key; encryption alone would allow broad plaintext recovery. Do
not log raw observations. Bound fields and novel observations per caller/device; otherwise an attacker creates unbounded
cardinality. Authorization separates raw, relationship, aggregate, and deletion access.

## Operations and alternatives

Monitor merge/split rate, confidence distribution, raw TTL compliance, high-degree nodes, new-key cardinality, lookup
latency, reputation freshness, and outcome drift. Reject a real-time graph traversal for the 40 ms path; precompute
features. Reject one irreversible device id because probabilistic identity needs correction lineage.
