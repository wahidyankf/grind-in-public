---
tldr: "Uses a recursive CTE with an array path and depth guard."
when_to_use: "Use after Drill 008."
---

# Solution 008: Recursive Relationships

```sql
WITH RECURSIVE descendants AS (
    SELECT team_id, parent_team_id, name, 0 AS depth, ARRAY[team_id] AS path
    FROM teams WHERE team_id = 1
    UNION ALL
    SELECT c.team_id, c.parent_team_id, c.name, d.depth + 1, d.path || c.team_id
    FROM teams AS c JOIN descendants AS d ON c.parent_team_id = d.team_id
    WHERE c.team_id <> ALL (d.path) AND d.depth < 20
)
SELECT * FROM descendants ORDER BY path;
```

Index `teams(parent_team_id)` for the recursive child lookup.
