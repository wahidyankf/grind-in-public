---
tldr: "Designs identity, authorization, tenant isolation, evidence integrity, retention, and privacy into every path."
when_to_use: "Use for multi-tenant or high-assurance systems where a technically correct result is not sufficient."
---

# Security, Tenancy, and Audit

Security is a set of invariants across identity, data, runtime, and operations. Adding an API gateway does not protect a
worker that consumes an event without re-establishing tenant context.

## Trust boundaries

```text
untrusted            controlled edge              trusted services

client --TLS--> [WAF/API/authn] --identity--> [policy enforcement]
                                                  |       |
                                                  v       v
                                             [data]   [audit]
```

Authenticate the calling principal, authorize the action against resource and tenant, validate the request, and carry a
minimal signed identity context. Services should not trust a client-supplied `tenant_id` merely because it is a UUID.

## Authorization models

| Model | Production fit                                      | Cost and rejection case                                   |
| ----- | --------------------------------------------------- | --------------------------------------------------------- |
| RBAC  | Stable job roles and broad permissions              | Role explosion; reject for many contextual conditions     |
| ABAC  | Tenant, region, ownership, risk, and time policies  | Harder policy testing; reject if attributes are untrusted |
| ReBAC | Resource-sharing graphs and delegated relationships | Traversal/cache cost; reject for simple static roles      |

Many systems combine coarse RBAC with contextual ABAC. Centralize policy definition but enforce near each protected
resource so an internal call cannot bypass authorization.

## Tenant isolation levels

```text
shared table                 schema/database per tenant          dedicated stack
+----------------+          +---------+ +---------+             +-----------+
| tenant_id, ... |          | tenant A| | tenant B|             | tenant X  |
+----------------+          +---------+ +---------+             +-----------+
 low cost / high density       more isolation and operations       highest isolation
```

Shared storage requires every key, index, cache entry, queue message, log query, and metric label to preserve tenant
scope. Database row-level policies are defense in depth, not a substitute for application tests. Dedicated placement may
be justified by residency, noisy-neighbour risk, encryption keys, or contracted isolation. Reject it for every small
tenant when the fleet cannot be patched and observed consistently.

## Audit versus application logs

An audit record answers who did what, to which resource, when, under which authority, with which version, and what the
outcome was. It should be append-only, tamper-evident, access-controlled, and retained by policy. Application logs are
diagnostic and may be sampled or short-lived.

```text
record N-1 hash ----+
                    v
event N fields -> canonical encoding -> hash -> record N hash
```

A hash chain detects alteration but does not prevent deletion of the final suffix. Periodically anchor signed batch
roots in a separately administered store, monitor sequence gaps, and control clock and key rotation.

## Encryption and secrets

- TLS protects transport; mutual TLS may identify workloads.
- Envelope encryption uses a data-encryption key wrapped by a managed key-encryption key.
- Per-tenant keys improve isolation and deletion control but increase key lifecycle work.
- Secrets belong in a secret manager and short-lived workload identity, not images or plain configuration.

Encryption does not fix overbroad authorization, leaked plaintext logs, or an application that can read every tenant.

## Retention and deletion

Classify canonical data, evidence, derived indexes, caches, logs, and backups. A deletion workflow must find all copies,
honour legal holds, produce evidence, and tolerate retries. Crypto-shredding helps when data is encrypted with a unique
retirable key, but only if no plaintext copy survives elsewhere.

## Threat-driven review

For each flow, ask about spoofing, tampering, repudiation, information disclosure, denial of service, and privilege
escalation. Then name the control and its verification. Reject controls that cannot be observed or tested.
