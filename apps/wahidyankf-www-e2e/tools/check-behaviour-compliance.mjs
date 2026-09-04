import { spawnSync } from "node:child_process";
import { readFile } from "node:fs/promises";
import path from "node:path";

import {
  findFiles,
  inspectCoverageMarkers,
  inspectFeatureReferences,
  inspectOwnerCoverageRequirements,
  validateAlternativeProofReferences,
  validateFeatureFiles,
  validateOwnerAdapterCoverage,
} from "./behaviour-compliance.mjs";

const projectRoot = process.cwd();
const workspaceRoot = path.resolve(projectRoot, "../..");
const featureRoot = path.resolve(
  projectRoot,
  "../../specs/apps/wahidyankf-www/behaviours",
);
const testsRoot = path.join(projectRoot, "tests");
const directSpecPattern = /\.spec\.(?:[cm]?[jt]s)$/u;

function runBddgen(args) {
  const result = spawnSync("npm", ["exec", "--", "bddgen", ...args], {
    cwd: projectRoot,
    encoding: "utf8",
    env: { ...process.env, FORCE_COLOR: "0" },
  });
  const output = `${result.stdout}${result.stderr}`;
  if (result.status !== 0) throw new Error(output.trim());
  return output;
}

async function main() {
  const featureFiles = await findFiles(featureRoot, (fileName) =>
    fileName.endsWith(".feature"),
  );
  if (featureFiles.length === 0)
    throw new Error(`No feature files found below '${featureRoot}'.`);

  const structuralErrors = await validateFeatureFiles(featureFiles);
  if (structuralErrors.length > 0) throw new Error(structuralErrors.join("\n"));

  const references = await Promise.all(
    featureFiles.map(async (featureFile) =>
      inspectFeatureReferences(
        featureFile,
        await readFile(featureFile, "utf8"),
      ),
    ),
  );
  const projectTargets = new Set();
  for (const projectFile of await findFiles(
    path.resolve(projectRoot, "../"),
    (name) => name === "project.json",
  )) {
    const project = JSON.parse(await readFile(projectFile, "utf8"));
    for (const target of Object.keys(project.targets ?? {}))
      projectTargets.add(`${project.name}:${target}`);
  }
  const referenceErrors = validateAlternativeProofReferences(
    references.flatMap(({ proofs }) => proofs),
    projectTargets,
    new Set(references.flatMap(({ scenarios }) => scenarios)),
  );
  if (referenceErrors.length > 0) throw new Error(referenceErrors.join("\n"));

  const requirements = (
    await Promise.all(
      featureFiles.map(async (featureFile) =>
        inspectOwnerCoverageRequirements(
          path.relative(workspaceRoot, featureFile).split(path.sep).join("/"),
          await readFile(featureFile, "utf8"),
        ),
      ),
    )
  ).flat();
  const ownerRoot = path.resolve(projectRoot, "../wahidyankf-www");
  const unitAdapterFiles = [
    ...(await findFiles(
      path.join(ownerRoot, "tests/bdd"),
      (name) => /\.[cm]?[jt]sx?$/u.test(name) && name !== "test-layer.ts",
    )),
    path.join(ownerRoot, "tests/integration/cv-pdf.integration.test.ts"),
  ].toSorted();
  const sharedIntegrationAdapters = new Set([
    "env-loader.steps.ts",
    "port-resolver.behaviour.test.ts",
    "tier-env.behaviour.test.ts",
  ]);
  const integrationAdapterFiles = [
    ...(await findFiles(path.join(ownerRoot, "tests/integration"), (name) =>
      /\.[cm]?[jt]sx?$/u.test(name),
    )),
    ...unitAdapterFiles.filter((file) =>
      sharedIntegrationAdapters.has(path.basename(file)),
    ),
  ].toSorted();
  const readMarkers = async (files) =>
    (
      await Promise.all(
        files.map(async (file) =>
          inspectCoverageMarkers(
            path.relative(workspaceRoot, file).split(path.sep).join("/"),
            await readFile(file, "utf8"),
          ),
        ),
      )
    ).flat();
  const ownerCoverageErrors = validateOwnerAdapterCoverage(
    requirements,
    await readMarkers(unitAdapterFiles),
    await readMarkers(integrationAdapterFiles),
  );
  if (ownerCoverageErrors.length > 0)
    throw new Error(ownerCoverageErrors.join("\n"));

  const directSpecs = await findFiles(testsRoot, (fileName) =>
    directSpecPattern.test(fileName),
  );
  if (directSpecs.length > 0) {
    throw new Error(
      `Direct Playwright journey specs are forbidden:\n${directSpecs.join("\n")}`,
    );
  }

  runBddgen(["--config", "playwright.config.ts"]);
  const unusedOutput = runBddgen([
    "export",
    "--config",
    "playwright.config.ts",
    "--unused-steps",
  ]);
  const unusedCount = /Unused steps \((\d+)\):/u.exec(unusedOutput)?.[1];
  if (unusedCount === undefined)
    throw new Error(
      `Could not determine unused-step count:\n${unusedOutput.trim()}`,
    );
  if (unusedCount !== "0") throw new Error(unusedOutput.trim());
}

try {
  await main();
} catch (error) {
  process.stderr.write(`${error instanceof Error ? error.message : error}\n`);
  process.exitCode = 1;
}
