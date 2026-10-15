import assert from "node:assert/strict";
import test from "node:test";

import { validateProjectContract } from "./project-contract.mjs";

function target(command = "true") {
  return { command, inputs: [] };
}

function owner(name, e2eName) {
  const corpus = `{workspaceRoot}/specs/apps/${name}/behaviours/**/*.feature`;
  const targets = Object.fromEntries(
    [
      "build",
      "lint",
      "test:coverage",
      "test:coverage:integration",
      "test:coverage:unit",
      "test:integration",
      "test:quick",
      "test:unit",
      "typecheck",
    ].map((name) => [name, { ...target(), inputs: [corpus] }]),
  );
  targets["test:coverage:behaviour"] = {
    cache: true,
    inputs: [corpus],
    options: {
      commands: [
        `${name}:test:coverage:behaviour:${name}`,
        `${e2eName}:test:coverage:behaviour:e2e`,
      ],
      parallel: false,
    },
  };
  targets["test:integration"].cache = false;
  targets["test:coverage:integration"].cache = false;
  targets["test:quick"] = {
    inputs: [corpus],
    options: {
      commands: [
        `${name}:typecheck`,
        `${name}:lint`,
        `${name}:test:unit`,
        `${name}:test:coverage:unit`,
        `${name}:test:coverage:behaviour`,
      ],
      parallel: false,
    },
  };
  return {
    $schema: "../../node_modules/nx/schemas/project-schema.json",
    name,
    projectType: "application",
    root: `apps/${name}`,
    tags: [],
    targets,
  };
}

function e2e(name, ownerName) {
  const corpus = `{workspaceRoot}/specs/apps/${ownerName}/behaviours/**/*.feature`;
  return {
    $schema: "../../node_modules/nx/schemas/project-schema.json",
    implicitDependencies: [ownerName],
    name,
    projectType: "application",
    root: `apps/${name}`,
    tags: [],
    targets: {
      lint: { ...target(), inputs: [corpus] },
      "test:coverage:behaviour": {
        ...target(`${ownerName}:test:coverage:behaviour`),
        inputs: [corpus],
      },
      "test:coverage:behaviour:e2e": {
        ...target(),
        cache: true,
        inputs: [corpus],
      },
      "test:e2e": {
        cache: false,
        command: "run",
        dependsOn: ["test:coverage:behaviour"],
        inputs: [corpus],
      },
      "test:quick": {
        inputs: [corpus],
        options: {
          commands: [
            `${name}:typecheck`,
            `${name}:lint`,
            `${name}:test:coverage:behaviour:e2e`,
          ],
          parallel: false,
        },
      },
      typecheck: { ...target(), inputs: [corpus] },
    },
  };
}

function validWorkspace() {
  return [
    owner("wahidyankf-www", "wahidyankf-www-e2e"),
    e2e("wahidyankf-www-e2e", "wahidyankf-www"),
  ];
}

function standalone(name, targetNames) {
  return {
    $schema: "../../node_modules/nx/schemas/project-schema.json",
    name,
    projectType: "application",
    sourceRoot: `apps/${name}`,
    targets: Object.fromEntries(targetNames.map((name) => [name, target()])),
  };
}

function validStandalone() {
  const web = ["lint", "test:quick", "test:unit", "typecheck"];
  return [
    standalone("edutech-be", web),
    standalone("edutech-web", web),
    standalone("edutech-web-e2e", [
      "lint",
      "test:e2e",
      "test:quick",
      "typecheck",
    ]),
  ];
}

test("accepts the owner, E2E, and standalone project contract", () => {
  const projects = [...validWorkspace(), ...validStandalone()];
  assert.deepEqual(validateProjectContract(projects), []);
});

test("rejects a standalone app missing a required target", () => {
  const projects = [...validWorkspace(), ...validStandalone()];
  delete projects.find((project) => project.name === "edutech-be").targets[
    "test:quick"
  ];
  assert.deepEqual(validateProjectContract(projects), [
    "edutech-be: missing test:quick",
  ]);
});

test("still rejects other unknown projects", () => {
  const projects = [
    ...validWorkspace(),
    ...validStandalone(),
    standalone("other-app", ["lint"]),
  ];
  assert.deepEqual(validateProjectContract(projects), [
    "workspace: unexpected project other-app",
  ]);
});

test("rejects reverse ownership and E2E layer placeholders", () => {
  const projects = validWorkspace();
  projects[0].implicitDependencies = ["wahidyankf-www-e2e"];
  projects[1].targets["test:unit"] = target();
  const findings = validateProjectContract(projects);
  assert.ok(
    findings.some((finding) => finding.includes("owner must not depend")),
  );
  assert.ok(
    findings.some((finding) => finding.includes("must not expose test:unit")),
  );
});

test("rejects cache, corpus input, and quick-order drift", () => {
  const projects = validWorkspace();
  projects[0].targets["test:integration"].cache = true;
  projects[0].targets["test:unit"].inputs = [];
  projects[1].targets["test:quick"].options.commands.reverse();
  const findings = validateProjectContract(projects);
  assert.ok(findings.some((finding) => finding.includes("must be uncached")));
  assert.ok(findings.some((finding) => finding.includes("must include specs")));
  assert.ok(
    findings.some((finding) => finding.includes("commands must match")),
  );
});

test("returns findings in deterministic order", () => {
  const findings = validateProjectContract([]);
  assert.deepEqual(findings, [...findings].toSorted());
  // One missing-project finding per half of each pair, plus one per standalone app.
  assert.equal(findings.length, 5);
});
