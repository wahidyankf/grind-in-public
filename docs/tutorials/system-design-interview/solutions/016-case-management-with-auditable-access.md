---
tldr:
  "A relational workflow authority, policy enforcement, secured search projection, and separate evidence ledger manage
  cases."
when_to_use: "Use after attempting Case 016 to compare workflow invariants, authorization, search, and audit."
---

# Solution 016: Case Management with Auditable Access

## Requirement traceability

| Requirements      | Design response                                                                                                 |
| ----------------- | --------------------------------------------------------------------------------------------------------------- |
| FR-1, FR-6, NFR-3 | Relational case aggregate, valid-transition table, and optimistic version on every mutation.                    |
| FR-2, NFR-4       | Immutable evidence references and append-only conclusion revisions.                                             |
| FR-3, NFR-2       | Central policy definitions with local enforcement using actor/resource/context attributes; deny on uncertainty. |
| FR-4              | Outbox-fed tenant-secured search projection with declared 30-second lag.                                        |
| FR-5              | Separate append-only audit path for reads, writes, exports, assignments, and overrides.                         |
| NFR-1, NFR-5      | Indexed work queues, cursor pagination, tenant quotas, and isolated heavy exports.                              |

## Architecture

```text
user -> API gateway/authn -> case service -> relational case DB
                              |     |               |
                              |     +-> evidence object references
                              |     +-> audit ledger
                              `-> outbox -> secured search index
policy bundle ----------------^
```

## Mutation transaction

```text
begin
  read case WHERE tenant=? AND id=?
  authorize(actor, action, current resource, context)
  validate transition and separation-of-duty rule
  UPDATE ... WHERE version=:expected
  append history + outbox + audit intent
commit
```

Zero updated rows means a concurrent mutation; return conflict with current version rather than silently overwriting.
Audit delivery from the outbox is monitored; security-sensitive reads append audit synchronously or use a durable edge
buffer before returning sensitive content.

## Authorization and search

Signed policy bundles can be cached briefly, but inputs include current assignment, sensitivity, tenant, role, and team.
Policy outage uses last-known-good within a bounded validity; beyond it, deny privileged operations. Search documents
include tenant and coarse authorization labels, and the API re-authorizes returned resource ids before disclosure. Never
trust a client-supplied tenant filter as the isolation boundary.

## Evidence and exports

Evidence objects use tenant keys, immutable object versions, checksums, and short-lived authorized download URLs. Export
is an asynchronous job with a snapshot, separate authorization at request and download, quotas, expiry, and audit. Legal
holds overlay case/evidence retention.

## Operations and alternatives

Monitor workflow conflicts, stuck states, policy denial/fallback, audit lag, search lag/count drift, unauthorized
probes, queue age, and export saturation. Rebuild search from case/outbox history.

Reject using the search index as workflow authority because it is asynchronous. Reject broad static RBAC alone because
assignment, sensitivity, ownership, and separation of duty are contextual.
