import { fileURLToPath } from "node:url";

import { validateGovernanceStructure } from "./governance-structure.mjs";

const workspaceRoot = fileURLToPath(new URL("../", import.meta.url));

try {
  const findings = await validateGovernanceStructure(workspaceRoot);
  if (findings.length > 0) throw new Error(findings.join("\n"));
  process.stdout.write("Governance structure passed.\n");
} catch (error) {
  process.stderr.write(`${error instanceof Error ? error.message : error}\n`);
  process.exitCode = 1;
}
