import { readFile } from "node:fs/promises";

import { validateProjectContract } from "./project-contract.mjs";

const projectPaths = [
  "apps/badakmini-cli/project.json",
  "apps/badakmini-cli-e2e/project.json",
  "apps/wahidyankf-www/project.json",
  "apps/wahidyankf-www-e2e/project.json",
];
const workspaceRoot = new URL("../", import.meta.url);

try {
  const projects = await Promise.all(
    projectPaths.map(async (path) =>
      JSON.parse(await readFile(new URL(path, workspaceRoot), "utf8")),
    ),
  );
  const findings = validateProjectContract(projects);
  if (findings.length > 0) throw new Error(findings.join("\n"));
  process.stdout.write(
    `Project contract passed for ${projects.length} projects.\n`,
  );
} catch (error) {
  process.stderr.write(`${error instanceof Error ? error.message : error}\n`);
  process.exitCode = 1;
}
