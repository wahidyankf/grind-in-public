import assert from "node:assert/strict";
import test from "node:test";

import {
  inspectFeatureReferences,
  inspectCoverageMarkers,
  inspectOwnerCoverageRequirements,
  validateAlternativeProofReferences,
  validateFeatureSource,
  validateOwnerAdapterCoverage,
} from "./behaviour-compliance.mjs";

const validFeature = `Feature: Compliance example

  Scenario: Observable behaviour
    Given a precondition
    When an action occurs
    Then an outcome is observed`;

test("accepts a feature with explicit When and Then steps", () => {
  assert.deepEqual(validateFeatureSource("example.feature", validFeature), []);
});

test("rejects missing declarations and scenarios", () => {
  assert.ok(
    validateFeatureSource("example.feature", "Feature: Empty").some((error) =>
      error.includes("at least one"),
    ),
  );
  assert.ok(
    validateFeatureSource(
      "example.feature",
      validFeature.replace("Feature:", "Subject:"),
    ).some((error) => error.includes("missing Feature")),
  );
});

test("rejects missing When and Then steps", () => {
  assert.ok(
    validateFeatureSource(
      "example.feature",
      validFeature.replace("    When an action occurs\n", ""),
    ).some((error) => error.includes("requires a When")),
  );
  assert.ok(
    validateFeatureSource(
      "example.feature",
      validFeature.replace("    Then an outcome is observed", ""),
    ).some((error) => error.includes("requires a Then")),
  );
});

test("accepts a documented higher-layer exemption", () => {
  const source = validFeature.replace(
    "  Scenario: Observable behaviour",
    "  # Exemption(integration): browser geometry needs a layout engine; alternative-proof: example-e2e:test:e2e / Observable behaviour\n" +
      "  @integration-exempt\n" +
      "  Scenario: Observable behaviour",
  );
  assert.deepEqual(validateFeatureSource("example.feature", source), []);
});

test("rejects legacy, unit, undocumented, and broad exemptions", () => {
  for (const tag of [
    "@unit",
    "@integration",
    "@e2e",
    "@unit-exempt",
    "@no-e2e",
  ]) {
    const source = validFeature.replace(
      "  Scenario: Observable behaviour",
      `  ${tag}\n  Scenario: Observable behaviour`,
    );
    assert.ok(
      validateFeatureSource("example.feature", source).some((error) =>
        error.includes("forbidden"),
      ),
    );
  }
  const undocumented = validFeature.replace(
    "  Scenario: Observable behaviour",
    "  @integration-exempt\n  Scenario: Observable behaviour",
  );
  assert.ok(
    validateFeatureSource("example.feature", undocumented).some((error) =>
      error.includes("preceding"),
    ),
  );
  const broad = validFeature.replace(
    "Feature: Compliance example",
    "@e2e-exempt\nFeature: Compliance example",
  );
  assert.ok(
    validateFeatureSource("example.feature", broad).some((error) =>
      error.includes("only annotate"),
    ),
  );
});

test("accepts independently documented integration and E2E exemptions", () => {
  const source = validFeature.replace(
    "  Scenario: Observable behaviour",
    "  # Exemption(integration): the local boundary cannot inject the private invariant failure; alternative-proof: example:test:unit / Observable behaviour\n" +
      "  @integration-exempt\n" +
      "  # Exemption(e2e): no public trigger can inject the private invariant failure; alternative-proof: example:test:unit / Observable behaviour\n" +
      "  @e2e-exempt\n" +
      "  Scenario: Observable behaviour",
  );
  assert.deepEqual(validateFeatureSource("example.feature", source), []);
});

test("requires independent documentation for each exemption", () => {
  const source = validFeature.replace(
    "  Scenario: Observable behaviour",
    "  # Exemption(integration): the local boundary cannot inject the private invariant failure; alternative-proof: example:test:unit / Observable behaviour\n" +
      "  @integration-exempt @e2e-exempt\n" +
      "  Scenario: Observable behaviour",
  );
  assert.ok(
    validateFeatureSource("example.feature", source).some((error) =>
      error.includes("@e2e-exempt requires the immediately preceding"),
    ),
  );
});

test("rejects exemptions for unfinished, flaky, or costly work", () => {
  for (const reason of [
    "too flaky and not yet implemented",
    "too expensive to implement",
  ]) {
    const source = validFeature.replace(
      "  Scenario: Observable behaviour",
      `  # Exemption(e2e): ${reason}; alternative-proof: example:test:integration / Observable behaviour\n` +
        "  @e2e-exempt\n" +
        "  Scenario: Observable behaviour",
    );
    assert.ok(
      validateFeatureSource("example.feature", source).some((error) =>
        error.includes("cannot be"),
      ),
    );
  }
});

test("alternative proof must resolve to an existing target and scenario", () => {
  const tagged = validFeature.replace(
    "  Scenario: Observable behaviour",
    "  # Exemption(e2e): local process boundary; alternative-proof: example:test:integration / Observable behaviour\n" +
      "  @e2e-exempt\n  Scenario: Observable behaviour",
  );
  const references = inspectFeatureReferences("example.feature", tagged);
  assert.deepEqual(
    validateAlternativeProofReferences(
      references.proofs,
      new Set(["example:test:integration"]),
      new Set(references.scenarios),
    ),
    [],
  );
  assert.equal(
    validateAlternativeProofReferences(references.proofs, new Set(), new Set())
      .length,
    2,
  );
});

test("owner coverage requires Unit and non-exempt Integration markers", () => {
  const resourceName = "specs/apps/example/behaviours/example.feature";
  const requirements = inspectOwnerCoverageRequirements(
    resourceName,
    validFeature,
  );
  const markers = inspectCoverageMarkers(
    "tests/example.steps.ts",
    `// @covers ${resourceName}:Observable behaviour`,
  );
  assert.deepEqual(
    validateOwnerAdapterCoverage(requirements, markers, markers),
    [],
  );
  assert.equal(validateOwnerAdapterCoverage(requirements, [], []).length, 2);
});

test("integration exemptions forbid an Integration adapter marker", () => {
  const resourceName = "specs/apps/example/behaviours/example.feature";
  const tagged = validFeature.replace(
    "  Scenario: Observable behaviour",
    "  # Exemption(integration): browser geometry; alternative-proof: example-e2e:test:e2e / Observable behaviour\n" +
      "  @integration-exempt\n  Scenario: Observable behaviour",
  );
  const requirements = inspectOwnerCoverageRequirements(resourceName, tagged);
  const markers = inspectCoverageMarkers(
    "tests/example.steps.ts",
    `// @covers ${resourceName}:Observable behaviour`,
  );
  assert.deepEqual(validateOwnerAdapterCoverage(requirements, markers, []), []);
  assert.equal(
    validateOwnerAdapterCoverage(requirements, markers, markers).length,
    1,
  );
});
