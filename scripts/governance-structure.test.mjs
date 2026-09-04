import assert from "node:assert/strict";
import test from "node:test";

import {
  validateDirectoryIndex,
  validateGovernanceFrontmatter,
} from "./governance-structure.mjs";

test("directory index requires every direct Markdown file and child directory", () => {
  const entries = [
    { isDirectory: false, name: "policy.md" },
    { isDirectory: true, name: "workflows" },
  ];
  assert.deepEqual(
    validateDirectoryIndex(
      "rules",
      entries,
      "[Policy](policy.md)\n[Workflows](workflows/README.md)",
    ),
    [],
  );
  assert.deepEqual(
    validateDirectoryIndex("rules", entries, "[Policy](policy.md)"),
    ["rules/README.md: missing direct entry workflows/README.md"],
  );
});

test("governance document requires routing frontmatter", () => {
  const valid = '---\ntldr: "Summary"\nwhen_to_use: "Route"\n---\n\n# Rule\n';
  assert.deepEqual(
    validateGovernanceFrontmatter("repo-governance/rule.md", valid),
    [],
  );
  assert.deepEqual(
    validateGovernanceFrontmatter("repo-governance/README.md", "# Index\n"),
    [],
  );
  assert.ok(
    validateGovernanceFrontmatter(
      "repo-governance/rule.md",
      "# Rule\n",
    )[0]?.includes("missing"),
  );
});
