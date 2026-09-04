import assert from "node:assert/strict";
import test from "node:test";

import { validateProjectContract } from "./project-contract.mjs";

function target(command = "true") {
  return { command };
}

function owner(name, e2eName) {
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
    ].map((name) => [name, target()]),
  );
  targets["test:coverage:behaviour"] = {
    cache: true,
    command: `${e2eName}:test:coverage:behaviour:e2e`,
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
  return {
    $schema: "../../node_modules/nx/schemas/project-schema.json",
    implicitDependencies: [ownerName],
    name,
    projectType: "application",
    root: `apps/${name}`,
    tags: [],
    targets: {
      lint: target(),
      "test:coverage:behaviour": target(`${ownerName}:test:coverage:behaviour`),
      "test:e2e": {
        cache: false,
        command: "run",
        dependsOn: ["test:coverage:behaviour"],
      },
      "test:quick": target(),
      typecheck: target(),
    },
  };
}

function validWorkspace() {
  return [
    owner("badakmini-cli", "badakmini-cli-e2e"),
    e2e("badakmini-cli-e2e", "badakmini-cli"),
    owner("wahidyankf-www", "wahidyankf-www-e2e"),
    e2e("wahidyankf-www-e2e", "wahidyankf-www"),
  ];
}

test("accepts the four-project owner and E2E contract", () => {
  assert.deepEqual(validateProjectContract(validWorkspace()), []);
});

test("rejects reverse ownership and E2E layer placeholders", () => {
  const projects = validWorkspace();
  projects[0].implicitDependencies = ["badakmini-cli-e2e"];
  projects[1].targets["test:unit"] = target();
  const findings = validateProjectContract(projects);
  assert.ok(
    findings.some((finding) => finding.includes("owner must not depend")),
  );
  assert.ok(
    findings.some((finding) => finding.includes("must not expose test:unit")),
  );
});

test("returns findings in deterministic order", () => {
  const findings = validateProjectContract([]);
  assert.deepEqual(findings, [...findings].toSorted());
  assert.equal(findings.length, 4);
});
