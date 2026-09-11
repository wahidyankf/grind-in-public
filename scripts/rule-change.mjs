// Detects edits to this repository's rules so Rules Propagation starts
// automatically instead of being remembered.
//
// Everything here is pure: paths in, paths or a notice out. The entrypoint
// beside this file owns the Git process and the hook's standard streams, which
// is what makes each behaviour below testable without either.

import path from "node:path";

// The procedure every rule change must follow. The notice points at it rather
// than restating its steps, so the workflow stays the one source.
export const WORKFLOW = "repo-governance/workflows/rules/rules-propagation.md";

// Proves the harnesses stayed equal after a change to one of them. A harness
// change is also a rule change, so both workflows can apply at once.
export const HARNESS_WORKFLOW =
  "repo-governance/workflows/harness-alignment.md";

// The paths a harness reads: the agent instruction files, each tool's
// configuration, and the directories holding its subagents, skills, and
// commands. Changing one of them can leave the harnesses unequal, which is
// what the alignment workflow checks.
const harnessFiles = ["AGENTS.md", "CLAUDE.md", "RTK.md", "opencode.json"];
const harnessDirectories = [".claude", ".codex", ".opencode", ".agents"];

// The paths that carry rules without being harness surfaces: the shared
// governance, and the Git hooks that enforce it.
const ruleFiles = harnessFiles;
const ruleDirectories = ["repo-governance", ".husky", ...harnessDirectories];

/** Every given path that carries rules, sorted and without duplicates. */
export function rulePaths(paths) {
  return selectPaths(paths, ruleFiles, ruleDirectories);
}

/**
 * Every given path a harness reads. Each of these is also a rule path; the
 * narrower list exists because only these can leave the harnesses unequal.
 */
export function harnessPaths(paths) {
  return selectPaths(paths, harnessFiles, harnessDirectories);
}

function selectPaths(paths, files, directories) {
  const matches = new Set();
  for (const candidate of paths ?? []) {
    const normalized = normalize(candidate);
    if (normalized && matchesAny(normalized, files, directories)) {
      matches.add(normalized);
    }
  }
  return [...matches].sort();
}

function matchesAny(candidate, files, directories) {
  if (files.includes(candidate)) return true;
  return directories.some(
    (directory) =>
      candidate === directory || candidate.startsWith(`${directory}/`),
  );
}

// Turns a path into the repository-relative slash form the rule lists use, so
// a Windows separator or a leading ./ still matches.
function normalize(candidate) {
  if (typeof candidate !== "string") return "";
  const trimmed = candidate.trim().replaceAll("\\", "/");
  if (trimmed === "") return "";
  const cleaned = path.posix.normalize(trimmed);
  if (cleaned === "." || cleaned === "./") return "";
  return cleaned.replace(/^\.\//, "");
}

/** The NUL-delimited paths Git emits for the next commit. */
export function parseStagedPaths(output) {
  return String(output)
    .split("\0")
    .filter((entry) => entry !== "");
}

// The lines that introduce a file in an apply_patch payload. A patch names
// every file it touches on one of these, which is what makes the paths
// readable before the edit lands.
const patchHeaders = [
  "*** Add File: ",
  "*** Update File: ",
  "*** Delete File: ",
  "*** Move to: ",
];

/**
 * The repository-relative paths a harness pre-edit payload is about to edit.
 * Harnesses name the target differently: Claude Code sends the file path,
 * Codex sends the whole patch, so both shapes are read. An unreadable payload
 * yields no paths rather than an error, because a notice must never break the
 * edit it comments on.
 */
export function hookPaths(payload, root) {
  let event;
  try {
    event = JSON.parse(payload);
  } catch {
    return [];
  }
  const input = event?.tool_input ?? {};

  const found = [];
  for (const candidate of [input.file_path, input.notebook_path]) {
    if (typeof candidate === "string" && candidate !== "") {
      found.push(relativeTo(root, candidate));
    }
  }
  for (const candidate of patchPaths(input.command)) {
    found.push(relativeTo(root, candidate));
  }
  return found;
}

// The files an apply_patch payload touches. A shell command arrives in the
// same field and simply contains no patch header, so it yields nothing rather
// than a false match.
function patchPaths(command) {
  if (typeof command !== "string") return [];
  const found = [];
  for (const line of command.split("\n")) {
    const trimmed = line.trim();
    const header = patchHeaders.find((candidate) =>
      trimmed.startsWith(candidate),
    );
    if (!header) continue;
    const candidate = trimmed.slice(header.length).trim();
    if (candidate !== "") found.push(candidate);
  }
  return found;
}

// Converts an absolute hook path into a repository-relative one and leaves
// anything outside the repository untouched, where it will not match.
function relativeTo(root, candidate) {
  if (!path.isAbsolute(candidate)) return candidate;
  return path.relative(root, candidate);
}

/**
 * Describes the detected rule change and starts the applicable workflow. It
 * names a workflow only when its paths changed, so a notice that always listed
 * both would teach readers to ignore the second line.
 */
export function notice(paths) {
  const text =
    `Rules Propagation automatically triggered by ${paths.join(", ")}.\n` +
    `Follow ${WORKFLOW} before completing this change.`;

  const harness = harnessPaths(paths);
  if (harness.length === 0) return text;

  return (
    `${text}\n` +
    `Harness setup changed in ${harness.join(", ")}.\n` +
    `Run ${HARNESS_WORKFLOW} so every harness stays equal.`
  );
}
