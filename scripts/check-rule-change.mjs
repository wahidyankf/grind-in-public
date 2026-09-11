// Announces a rule change to whoever is making it. Two invocations:
//
//   node scripts/check-rule-change.mjs        the staged tree, before a commit
//   node scripts/check-rule-change.mjs hook   a harness pre-edit payload
//
// Neither ever blocks. The workflow it names is the author's next step, not a
// gate a hook can decide has been satisfied — and a notice that could break an
// edit would be worse than no notice at all. Both invocations exit 0 on every
// path, including their own failures.

import { execFileSync } from "node:child_process";
import { readFileSync } from "node:fs";

import {
  hookPaths,
  notice,
  parseStagedPaths,
  rulePaths,
} from "./rule-change.mjs";

function repositoryRoot() {
  return execFileSync("git", ["rev-parse", "--show-toplevel"], {
    encoding: "utf8",
  }).trim();
}

function announceStaged() {
  let staged;
  try {
    const root = repositoryRoot();
    const output = execFileSync(
      "git",
      ["diff", "--cached", "--name-only", "-z"],
      { cwd: root, encoding: "utf8" },
    );
    staged = parseStagedPaths(output);
  } catch (error) {
    process.stderr.write(
      `ERROR: ${error instanceof Error ? error.message : error}\n`,
    );
    return;
  }

  const paths = rulePaths(staged);
  if (paths.length === 0) return;
  process.stdout.write(`${notice(paths)}\n`);
}

// Answers a harness pre-edit hook. The notice arrives as additional context so
// the agent loads the workflow before it edits, and as a system message so the
// human reads it too. Every other file stays silent, so ordinary work is never
// interrupted.
function announceHook() {
  let paths;
  try {
    const payload = readFileSync(0, "utf8");
    paths = rulePaths(hookPaths(payload, repositoryRoot()));
  } catch {
    // A hook that cannot read its payload or find its repository must not
    // block the edit.
    return;
  }
  if (paths.length === 0) return;

  const text = notice(paths);
  try {
    process.stdout.write(
      `${JSON.stringify({
        hookSpecificOutput: {
          hookEventName: "PreToolUse",
          additionalContext: text,
        },
        systemMessage: text,
      })}\n`,
    );
  } catch {
    // As above: neither failure may block the edit.
  }
}

if (process.argv[2] === "hook") {
  announceHook();
} else {
  announceStaged();
}
