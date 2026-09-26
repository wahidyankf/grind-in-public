---
tldr: "Design privacy-aware device identity, relationship signals, reputation updates, and low-latency lookup."
when_to_use: "Use to practise probabilistic identity, graph edges, TTLs, hot keys, and privacy controls."
---

# Case 012: Device Identity and Reputation

Design a service that maps privacy-approved device observations to pseudonymous device identities and produces
reputation signals from account relationships and historical outcomes.

### Functional requirements

| ID   | Requirement                                                                                |
| ---- | ------------------------------------------------------------------------------------------ |
| FR-1 | Accept bounded device observations and return a pseudonymous device id plus confidence.    |
| FR-2 | Record time-bounded device-account, device-network, and device-session relationships.      |
| FR-3 | Return reputation features and reason codes for an online decision.                        |
| FR-4 | Update reputation from validated outcomes and decay stale evidence.                        |
| FR-5 | Delete or unlink data under approved privacy workflows without breaking audit obligations. |

### Non-functional requirements

| ID    | Requirement                                                                              |
| ----- | ---------------------------------------------------------------------------------------- |
| NFR-1 | Resolve and read reputation within 40 ms p99 at 30,000 requests/s.                       |
| NFR-2 | Bound false merges and false splits with measured confidence thresholds.                 |
| NFR-3 | Retain raw observations for the minimum policy period and pseudonymize identifiers.      |
| NFR-4 | Resist enumeration, spoofed high-cardinality observations, and hot shared devices.       |
| NFR-5 | Keep tenant-specific reputation separate unless an explicit shared-data contract exists. |

## Assumptions and exclusions

Only approved signals are collected. The service does not attempt invasive fingerprinting or claim deterministic human
identity.

## Interview prompts

1. Which observations are exact keys versus probabilistic features?
2. How are identities merged, split, versioned, and corrected?
3. Which graph computations run online versus batch?
4. How do TTL, decay, deletion, encryption, abuse limits, and tenant boundaries work?

Solve before reading [the worked solution](../solutions/012-device-identity-and-reputation.md).
