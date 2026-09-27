---
tldr: "Uses a signed tenant-bound keyset cursor and matching B-tree."
when_to_use: "Use after attempting Case 003."
---

# Debrief 003: Pagination and Indexes

```sql
SELECT case_id, title, status, created_at
FROM cases
WHERE tenant_id = $1
  AND (created_at, case_id) < ($2, $3)
ORDER BY created_at DESC, case_id DESC
LIMIT $4;
```

Index `(tenant_id, created_at DESC, case_id DESC)`; consider a separate status-leading index only for measured filtered
traffic. The token signs tenant, last timestamp/id, filter hash, and version. Keyset traversal is stable relative to the
last key but does not provide a frozen snapshot across minutes; document that inserts appear before the cursor. Build
the index concurrently, verify validity, shadow results, then switch.
