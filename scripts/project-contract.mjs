const schema = "../../node_modules/nx/schemas/project-schema.json";
const ownerTargets = [
  "build",
  "lint",
  "test:coverage",
  "test:coverage:behaviour",
  "test:coverage:integration",
  "test:coverage:unit",
  "test:integration",
  "test:quick",
  "test:unit",
  "typecheck",
];
const e2eTargets = [
  "lint",
  "test:coverage:behaviour",
  "test:e2e",
  "test:quick",
  "typecheck",
];

function commandText(target) {
  if (typeof target?.command === "string") return target.command;
  return JSON.stringify(target?.options?.commands ?? []);
}

function checkCommon(project, findings) {
  const expectedRoot = `apps/${project.name}`;
  if (project.$schema !== schema)
    findings.push(`${project.name}: invalid Nx schema path`);
  if (project.projectType !== "application")
    findings.push(`${project.name}: projectType must be application`);
  if (project.root !== expectedRoot)
    findings.push(`${project.name}: root must be ${expectedRoot}`);
  if (!Array.isArray(project.tags) || project.tags.length !== 0)
    findings.push(`${project.name}: tags must be []`);
}

function checkTargetSet(project, required, findings) {
  for (const target of required) {
    if (project.targets?.[target] === undefined)
      findings.push(`${project.name}: missing ${target}`);
  }
}

function checkOwner(project, e2eName, findings) {
  checkTargetSet(project, ownerTargets, findings);
  if (project.targets?.["test:e2e"] !== undefined)
    findings.push(`${project.name}: owner must not expose test:e2e`);
  if ((project.implicitDependencies ?? []).includes(e2eName)) {
    findings.push(`${project.name}: owner must not depend on its E2E project`);
  }
  const aggregate = project.targets?.["test:coverage:behaviour"];
  if (
    !commandText(aggregate).includes(`${e2eName}:test:coverage:behaviour:e2e`)
  ) {
    findings.push(
      `${project.name}: behaviour aggregate must include the E2E static slice`,
    );
  }
  if (aggregate?.cache !== true)
    findings.push(`${project.name}: behaviour aggregate must be cached`);
}

function checkE2E(project, ownerName, findings) {
  checkTargetSet(project, e2eTargets, findings);
  for (const forbidden of [
    "test:unit",
    "test:integration",
    "test:coverage:unit",
    "test:coverage:integration",
  ]) {
    if (project.targets?.[forbidden] !== undefined)
      findings.push(
        `${project.name}: E2E project must not expose ${forbidden}`,
      );
  }
  if (
    JSON.stringify(project.implicitDependencies ?? []) !==
    JSON.stringify([ownerName])
  ) {
    findings.push(
      `${project.name}: implicitDependencies must be [${ownerName}]`,
    );
  }
  const delegate = project.targets?.["test:coverage:behaviour"];
  if (!commandText(delegate).includes(`${ownerName}:test:coverage:behaviour`)) {
    findings.push(
      `${project.name}: generic behaviour target must delegate to ${ownerName}`,
    );
  }
  const runtime = project.targets?.["test:e2e"];
  if (runtime?.cache !== false)
    findings.push(`${project.name}: test:e2e must be uncached`);
  if (
    !JSON.stringify(runtime?.dependsOn ?? []).includes(
      "test:coverage:behaviour",
    )
  ) {
    findings.push(
      `${project.name}: test:e2e must depend on static behaviour compliance`,
    );
  }
}

/** Validates the fixed four-project quality contract without I/O or subprocesses. */
export function validateProjectContract(projects) {
  const findings = [];
  const byName = new Map(projects.map((project) => [project.name, project]));
  const pairs = [
    ["badakmini-cli", "badakmini-cli-e2e"],
    ["wahidyankf-www", "wahidyankf-www-e2e"],
  ];
  for (const [ownerName, e2eName] of pairs) {
    const owner = byName.get(ownerName);
    const e2e = byName.get(e2eName);
    if (owner === undefined) findings.push(`workspace: missing ${ownerName}`);
    if (e2e === undefined) findings.push(`workspace: missing ${e2eName}`);
    if (owner !== undefined) {
      checkCommon(owner, findings);
      checkOwner(owner, e2eName, findings);
    }
    if (e2e !== undefined) {
      checkCommon(e2e, findings);
      checkE2E(e2e, ownerName, findings);
    }
  }
  for (const project of projects) {
    if (!pairs.flat().includes(project.name))
      findings.push(`workspace: unexpected project ${project.name}`);
  }
  return findings.toSorted();
}
