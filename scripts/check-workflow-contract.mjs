import { readFile } from "node:fs/promises";

import { validateWorkflowContract } from "./workflow-contract.mjs";

const workspaceRoot = new URL("../", import.meta.url);
const paths = {
  planGate: "repo-governance/workflows/plan-quality-gate.md",
  propagation: "repo-governance/workflows/rules/rules-propagation.md",
  rulesGate: "repo-governance/workflows/rules-quality-gate.md",
  taskTracking: "repo-governance/conventions/task-tracking-policy.md",
  tdd: "repo-governance/development/tdd-policy.md",
};

try {
  const documents = Object.fromEntries(
    await Promise.all(
      Object.entries(paths).map(async ([name, path]) => [
        name,
        await readFile(new URL(path, workspaceRoot), "utf8"),
      ]),
    ),
  );
  const findings = validateWorkflowContract(documents);
  if (findings.length > 0) throw new Error(findings.join("\n"));
  process.stdout.write("Workflow contract passed.\n");
} catch (error) {
  process.stderr.write(`${error instanceof Error ? error.message : error}\n`);
  process.exitCode = 1;
}
