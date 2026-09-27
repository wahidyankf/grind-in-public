---
tldr: "Traverses bounded hierarchies and graphs with recursive CTEs and explicit cycle protection."
when_to_use: "Use for ownership trees, dependencies, and graph-shaped interview problems inside PostgreSQL."
---

# Recursive CTEs

A recursive CTE has an anchor, a recursive step, and a termination condition. This example uses a temporary hierarchy:

```sql
CREATE TEMP TABLE teams (
    team_id integer PRIMARY KEY,
    parent_team_id integer REFERENCES teams,
    name text NOT NULL
);

INSERT INTO teams VALUES
    (1, NULL, 'Platform'), (2, 1, 'Data'), (3, 1, 'Runtime'), (4, 2, 'Detection');

WITH RECURSIVE tree AS (
    SELECT team_id, parent_team_id, name, 0 AS depth, ARRAY[team_id] AS path
    FROM teams
    WHERE team_id = 1

    UNION ALL

    SELECT child.team_id, child.parent_team_id, child.name,
           tree.depth + 1, tree.path || child.team_id
    FROM teams AS child
    JOIN tree ON child.parent_team_id = tree.team_id
    WHERE child.team_id <> ALL (tree.path)
      AND tree.depth < 20
)
SELECT * FROM tree ORDER BY path;
```

The path prevents cycles; the depth limit bounds damage from bad data. Complexity is `O(V + E)` for the visited subgraph
conceptually, though SQL execution and indexes affect physical cost. Index the recursive join key. Use an application
graph engine when traversal is deep, highly connected, or central to the workload; keep authoritative workflow state
relational.

Related learning: [graph traversal](../python-algorithms-interview/009-graph-traversal-and-ordering.md).
