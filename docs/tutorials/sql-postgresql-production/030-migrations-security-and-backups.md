---
tldr: "Combines expand-contract migrations, least privilege, encryption boundaries, backups, and restore drills."
when_to_use: "Use before changing a production schema or declaring a database recoverable."
---

# Migrations, Security, and Backups

Production migration is a compatibility protocol across old code, new code, schema, replicas, and background jobs.

```text
expand schema -> deploy dual-compatible code -> backfill in bounded chunks
      -> verify parity -> switch reads -> stop old writes -> contract later
```

For a large-table change:

1. set `lock_timeout` and observe blockers;
2. add nullable columns or `NOT VALID` constraints;
3. deploy code that tolerates both shapes;
4. backfill by stable key in rate-limited, resumable chunks;
5. validate counts, hashes, nulls, and application metrics;
6. validate constraints and add indexes concurrently when appropriate;
7. make the new path authoritative;
8. remove old state only after a rollback window.

`CREATE INDEX CONCURRENTLY` reduces blocking but takes longer, performs more work, cannot run inside a transaction, and
can leave an invalid index after failure. Inspect and repair explicitly.

Security layers include TLS in transit, encrypted storage/backups, secrets outside source control, separate owner,
migration, application, reporting, and replication roles, tenant predicates/RLS, statement auditing for privileged
actions, and tested revocation. Dynamic identifiers use `psycopg.sql.Identifier`; values use parameters.

Backups are a recoverability claim only after restore:

```text
base backup + continuous WAL archive -> restore target -> replay to time/LSN
                                          |
                                          +-> integrity + application checks
```

Define RPO/RTO, retention, encryption keys, offsite failure boundaries, restore ownership, and a recurring drill.
Replicas are not backups: operator error and logical corruption replicate too.
