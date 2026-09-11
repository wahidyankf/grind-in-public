import assert from "node:assert/strict";
import { execFileSync, spawnSync } from "node:child_process";
import { mkdtempSync, mkdirSync, rmSync, writeFileSync } from "node:fs";
import { tmpdir } from "node:os";
import path from "node:path";
import test from "node:test";
import { fileURLToPath } from "node:url";

// The entrypoint, driven as a process. `rule-change.test.mjs` proves what the
// detector selects; these cases prove what the command does with it, which is
// the other half of the three scenarios named in that file's header. Each of
// them says "the command succeeds" — with a notice, without one, or with both
// workflows — and an exit code is not a property a pure function has.
//
// It is also the property this design rests on. The notice must never block:
// a reporting check that can fail closed will one day stop an unrelated commit,
// and the fix will be to delete the notice. That was proven by hand once, when
// the successor was wired. A hand proof does not survive the next refactor.
//
// Every case builds its own repository under the system temporary directory and
// removes it afterwards, so no case can read the staged state of the repository
// it is running in.

const ENTRYPOINT = fileURLToPath(
  new URL("./check-rule-change.mjs", import.meta.url),
);
const PROPAGATION = "repo-governance/workflows/rules/rules-propagation.md";
const HARNESS = "repo-governance/workflows/harness-alignment.md";

function withRepository(body) {
  const root = mkdtempSync(path.join(tmpdir(), "rule-change-"));
  try {
    execFileSync("git", ["init", "-q"], { cwd: root });
    body(root);
  } finally {
    rmSync(root, { recursive: true, force: true });
  }
}

function stage(root, relative, contents) {
  const absolute = path.join(root, relative);
  mkdirSync(path.dirname(absolute), { recursive: true });
  writeFileSync(absolute, contents);
  execFileSync("git", ["add", "--", relative], { cwd: root });
}

function announce(root, argument, input) {
  const result = spawnSync(
    "node",
    argument ? [ENTRYPOINT, argument] : [ENTRYPOINT],
    {
      cwd: root,
      input: input ?? "",
      encoding: "utf8",
    },
  );
  return { code: result.status, stdout: result.stdout ?? "" };
}

test("the staged command succeeds and names the automatically triggered workflow", () => {
  withRepository((root) => {
    stage(root, "repo-governance/development/tdd-policy.md", "# rule\n");
    const { code, stdout } = announce(root);
    assert.equal(code, 0);
    assert.ok(stdout.includes("repo-governance/development/tdd-policy.md"));
    assert.ok(stdout.includes(PROPAGATION));
    assert.ok(!stdout.includes(HARNESS));
  });
});

test("the staged command succeeds without output on an ordinary staged path", () => {
  withRepository((root) => {
    stage(
      root,
      "apps/wahidyankf-www/src/app/page.tsx",
      "export default null;\n",
    );
    const { code, stdout } = announce(root);
    assert.equal(code, 0);
    assert.equal(stdout, "");
  });
});

test("the hook command succeeds with both workflow notices for a harness edit", () => {
  withRepository((root) => {
    const payload = JSON.stringify({ tool_input: { file_path: "CLAUDE.md" } });
    const { code, stdout } = announce(root, "hook", payload);
    assert.equal(code, 0);
    const event = JSON.parse(stdout);
    assert.equal(event.hookSpecificOutput.hookEventName, "PreToolUse");
    for (const text of [
      event.hookSpecificOutput.additionalContext,
      event.systemMessage,
    ]) {
      assert.ok(text.includes(PROPAGATION));
      assert.ok(text.includes(HARNESS));
    }
  });
});

test("the hook command stays silent on a file no harness rule covers", () => {
  withRepository((root) => {
    const payload = JSON.stringify({
      tool_input: { file_path: "package.json" },
    });
    const { code, stdout } = announce(root, "hook", payload);
    assert.equal(code, 0);
    assert.equal(stdout, "");
  });
});

// Outside a repository both invocations lose the one thing they need. Neither
// may turn that into a non-zero exit: the staged run would break a commit it
// only comments on, and the hook run would break the edit.
test("neither invocation blocks when there is no repository to read", () => {
  const root = mkdtempSync(path.join(tmpdir(), "rule-change-bare-"));
  try {
    assert.equal(announce(root).code, 0);
    const payload = JSON.stringify({ tool_input: { file_path: "CLAUDE.md" } });
    assert.equal(announce(root, "hook", payload).code, 0);
  } finally {
    rmSync(root, { recursive: true, force: true });
  }
});

test("the hook command does not block on a payload it cannot read", () => {
  withRepository((root) => {
    const { code, stdout } = announce(root, "hook", "not json at all");
    assert.equal(code, 0);
    assert.equal(stdout, "");
  });
});
