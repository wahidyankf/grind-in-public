---
tldr: "Builds tenant isolation into keys, queries, indexes, row security, and operational boundaries."
when_to_use: "Use for shared-database SaaS designs and tenant-isolation interviews."
---

# Multi-Tenant Data Modelling

There are three common placement models:

```text
shared tables        schema per tenant       database/cell per tenant group
lowest overhead ---> more isolation ------> strongest blast-radius boundary
```

Shared tables fit many small tenants, provided every key, foreign key, unique constraint, index prefix, query, metric,
and background job carries `tenant_id`. Row-level security adds defence in depth:

```sql
ALTER TABLE cases ENABLE ROW LEVEL SECURITY;

CREATE POLICY cases_tenant_policy ON cases
USING (tenant_id = current_setting('app.tenant_id')::bigint)
WITH CHECK (tenant_id = current_setting('app.tenant_id')::bigint);
```

In a transaction, the application can use `SET LOCAL app.tenant_id = '1'`. Test owner/bypass roles carefully; RLS is not
a substitute for least privilege or application authorization. Administrative cross-tenant jobs need an explicit,
audited role.

Graduate a noisy or regulated tenant to a dedicated cell based on measured load, residency, recovery, or contractual
isolation. Maintain a placement directory and design identifiers so moves do not collide. “Shard by tenant” is only a
partition key; the hard work is routing, resharding, schema rollout, cross-shard reporting, and recovery.
