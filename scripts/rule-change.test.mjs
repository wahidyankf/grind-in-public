import assert from "node:assert/strict";
import test from "node:test";

// These cases inherit three scenarios from a Gherkin corpus that no longer
// exists. `specs/apps/badakmini-cli/behaviours/rule-change.feature` described
// the behaviour of an application that has been retired; the corpus rule binds
// applications and libraries, and this is neither. The scenarios survive here
// rather than as an orphan feature file no adapter consumes:
//
//   A staged rule path automatically triggers the workflow
//     -> selects a staged governance path as a rule change
//   An ordinary staged path stays silent
//     -> stays silent on an ordinary staged path
//   A harness edit automatically triggers both workflows
//     -> names both workflows when a harness surface changed
//
// Each case loads the module itself rather than importing it at the top. A
// top-level import that cannot resolve kills the whole file on the first line,
// which would collapse every behaviour below into one failure and prove only
// that a file is missing. Loading per case keeps each behaviour answerable on
// its own, before the successor exists and after.
async function load() {
  return import("./rule-change.mjs");
}

const PROPAGATION = "repo-governance/workflows/rules/rules-propagation.md";
const HARNESS = "repo-governance/workflows/harness-alignment.md";

test("selects a staged governance path as a rule change", async () => {
  const { rulePaths } = await load();
  assert.deepEqual(rulePaths(["repo-governance/development/tdd-policy.md"]), [
    "repo-governance/development/tdd-policy.md",
  ]);
});

test("stays silent on an ordinary staged path", async () => {
  const { rulePaths } = await load();
  assert.deepEqual(rulePaths(["apps/wahidyankf-www/src/app/page.tsx"]), []);
});

test("selects every rule-bearing file and directory, and nothing else", async () => {
  const { rulePaths } = await load();
  const selected = rulePaths([
    "AGENTS.md",
    "CLAUDE.md",
    "RTK.md",
    "opencode.json",
    "repo-governance/README.md",
    ".husky/pre-commit",
    ".claude/settings.json",
    ".codex/hooks.json",
    ".opencode/agents/plan-maker.md",
    ".agents/agents/plan-maker.md",
    "package.json",
    "README.md",
  ]);
  assert.ok(!selected.includes("package.json"));
  assert.ok(!selected.includes("README.md"));
  assert.equal(selected.length, 10);
});

test("normalises, de-duplicates, and sorts what it reports", async () => {
  const { rulePaths } = await load();
  assert.deepEqual(
    rulePaths([
      "./repo-governance/b.md",
      "repo-governance\\a.md",
      "repo-governance/b.md",
      "   ",
    ]),
    ["repo-governance/a.md", "repo-governance/b.md"],
  );
});

test("treats a directory's own name as inside it, and a prefix match as outside", async () => {
  const { rulePaths } = await load();
  assert.deepEqual(rulePaths([".husky"]), [".husky"]);
  assert.deepEqual(rulePaths([".husky-notes/x.md"]), []);
});

test("narrows harness paths to the surfaces that can leave harnesses unequal", async () => {
  const { harnessPaths } = await load();
  assert.deepEqual(
    harnessPaths(["repo-governance/development/tdd-policy.md"]),
    [],
  );
  assert.deepEqual(harnessPaths(["CLAUDE.md"]), ["CLAUDE.md"]);
  assert.deepEqual(harnessPaths([".husky/pre-commit"]), []);
});

test("names only the propagation workflow for a rule path no harness reads", async () => {
  const { notice } = await load();
  const text = notice(["repo-governance/development/tdd-policy.md"]);
  assert.ok(text.includes(PROPAGATION));
  assert.ok(!text.includes(HARNESS));
});

test("names both workflows when a harness surface changed", async () => {
  const { notice } = await load();
  const text = notice(["CLAUDE.md"]);
  assert.ok(text.includes(PROPAGATION));
  assert.ok(text.includes(HARNESS));
});

test("reads the edited file out of a Claude pre-edit payload", async () => {
  const { hookPaths } = await load();
  const payload = JSON.stringify({
    tool_input: { file_path: "repo-governance/README.md" },
  });
  assert.deepEqual(hookPaths(payload, "/repo"), ["repo-governance/README.md"]);
});

test("reads a notebook path out of the same payload shape", async () => {
  const { hookPaths } = await load();
  const payload = JSON.stringify({
    tool_input: { notebook_path: "repo-governance/notes.ipynb" },
  });
  assert.deepEqual(hookPaths(payload, "/repo"), [
    "repo-governance/notes.ipynb",
  ]);
});

test("reads every file an apply_patch payload names", async () => {
  const { hookPaths } = await load();
  const payload = JSON.stringify({
    tool_input: {
      command: [
        "*** Add File: repo-governance/a.md",
        "*** Update File: repo-governance/b.md",
        "*** Delete File: repo-governance/c.md",
        "*** Move to: repo-governance/d.md",
      ].join("\n"),
    },
  });
  assert.deepEqual(hookPaths(payload, "/repo"), [
    "repo-governance/a.md",
    "repo-governance/b.md",
    "repo-governance/c.md",
    "repo-governance/d.md",
  ]);
});

test("finds nothing in a shell command that carries no patch header", async () => {
  const { hookPaths } = await load();
  const payload = JSON.stringify({
    tool_input: { command: "rm -rf repo-governance" },
  });
  assert.deepEqual(hookPaths(payload, "/repo"), []);
});

test("makes an absolute hook path repository-relative", async () => {
  const { hookPaths } = await load();
  const payload = JSON.stringify({
    tool_input: { file_path: "/repo/repo-governance/README.md" },
  });
  assert.deepEqual(hookPaths(payload, "/repo"), ["repo-governance/README.md"]);
});

test("yields nothing rather than throwing on an unreadable payload", async () => {
  const { hookPaths } = await load();
  assert.deepEqual(hookPaths("not json at all", "/repo"), []);
  assert.deepEqual(hookPaths("", "/repo"), []);
});
