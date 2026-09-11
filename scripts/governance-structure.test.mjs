import assert from "node:assert/strict";
import test from "node:test";

import {
  isVendoredDirectory,
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

test("a vendored directory is not held to this repository's index rule", () => {
  // The publication screen under scripts/public-safety is not written here. It
  // is one shared layer, kept byte-identical across every repository that
  // publishes, so a finding in one of them means the same thing in all of them.
  // Adding indexes this repository happens to want would fork that copy, and a
  // forked safety layer is worth less than a missing README.
  assert.equal(isVendoredDirectory("scripts/public-safety"), true);
  assert.equal(isVendoredDirectory("scripts/public-safety/tests"), true);
  assert.equal(isVendoredDirectory("scripts/public-safety/tests/cases"), true);

  // The exclusion is a named path, not a name anywhere. A directory that merely
  // ends with the same word is this repository's own and stays governed.
  assert.equal(isVendoredDirectory("scripts"), false);
  assert.equal(isVendoredDirectory("docs/public-safety"), false);
  assert.equal(isVendoredDirectory("scripts/public-safety-notes"), false);
});
