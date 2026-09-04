import { readdir, readFile } from "node:fs/promises";
import path from "node:path";

const governedTrees = ["docs", "repo-governance", "scripts", "plans", "specs"];

function normalizedLinkTarget(target) {
  return target.replaceAll("\\", "/").replace(/\/$/u, "");
}

export function validateDirectoryIndex(directory, entries, readme) {
  const findings = [];
  for (const entry of entries.toSorted((left, right) =>
    left.name.localeCompare(right.name),
  )) {
    if (entry.name.startsWith(".") || entry.name === "README.md") continue;
    if (!entry.isDirectory && path.extname(entry.name) !== ".md") continue;
    const expected = entry.isDirectory ? `${entry.name}/README.md` : entry.name;
    const linkPattern = new RegExp(
      `\\([^)]*${expected.replace(/[.*+?^${}()|[\]\\]/gu, "\\$&")}[^)]*\\)`,
      "u",
    );
    if (!linkPattern.test(readme))
      findings.push(`${directory}/README.md: missing direct entry ${expected}`);
  }
  return findings;
}

export function validateGovernanceFrontmatter(resourceName, source) {
  if (path.basename(resourceName) === "README.md") return [];
  if (!source.startsWith("---\n"))
    return [`${resourceName}: missing YAML frontmatter`];
  const end = source.indexOf("\n---\n", 4);
  if (end < 0) return [`${resourceName}: unterminated YAML frontmatter`];
  const frontmatter = source.slice(4, end);
  const findings = [];
  for (const field of ["tldr:", "when_to_use:"]) {
    if (!frontmatter.split("\n").some((line) => line.startsWith(field))) {
      findings.push(
        `${resourceName}: frontmatter is missing ${field.slice(0, -1)}`,
      );
    }
  }
  return findings;
}

async function inspectTree(workspaceRoot, relativeDirectory, findings) {
  const absoluteDirectory = path.join(workspaceRoot, relativeDirectory);
  const entries = await readdir(absoluteDirectory, { withFileTypes: true });
  const readmePath = path.join(absoluteDirectory, "README.md");
  let readme;
  try {
    readme = await readFile(readmePath, "utf8");
  } catch {
    findings.push(`${relativeDirectory}/README.md: missing directory index`);
    readme = "";
  }
  findings.push(
    ...validateDirectoryIndex(
      relativeDirectory,
      entries.map((entry) => ({
        isDirectory: entry.isDirectory(),
        name: normalizedLinkTarget(entry.name),
      })),
      readme,
    ),
  );
  for (const entry of entries.toSorted((left, right) =>
    left.name.localeCompare(right.name),
  )) {
    const relativePath = path.join(relativeDirectory, entry.name);
    if (entry.isDirectory()) {
      await inspectTree(workspaceRoot, relativePath, findings);
    } else if (
      relativeDirectory.startsWith("repo-governance") &&
      entry.name.endsWith(".md")
    ) {
      findings.push(
        ...validateGovernanceFrontmatter(
          relativePath,
          await readFile(path.join(workspaceRoot, relativePath), "utf8"),
        ),
      );
    }
  }
}

/** Reads governed trees in lexical order and returns sorted structural findings. */
export async function validateGovernanceStructure(workspaceRoot) {
  const findings = [];
  for (const tree of governedTrees)
    await inspectTree(workspaceRoot, tree, findings);
  return findings.toSorted();
}
