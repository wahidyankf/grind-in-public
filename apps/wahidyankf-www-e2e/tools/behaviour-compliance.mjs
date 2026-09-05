import { readdir, readFile } from "node:fs/promises";
import path from "node:path";

const scenarioPattern = /^(?:Scenario(?: Outline| Template)?|Example):/iu;
const exemptionTags = new Set(["integration-exempt", "e2e-exempt"]);
const forbiddenTags = new Set([
  "unit",
  "integration",
  "e2e",
  "unit-exempt",
  "no-unit",
  "no-integration",
  "no-e2e",
]);
const exemptionCommentPattern =
  /^# Exemption\((integration|e2e)\): (.+); alternative-proof: (.+)$/u;
const invalidReasonPattern =
  /\b(?:hard|slow|flaky|cost(?:ly)?|expensive|not yet implemented|todo)\b/iu;

function startsStep(keyword, line) {
  return line.toLowerCase().startsWith(`${keyword.toLowerCase()} `);
}

function finishScenario(resourceName, state, errors) {
  if (state.currentScenario === undefined) return;
  if (!state.hasWhen) {
    errors.push(
      `${resourceName}: ${state.currentScenario} requires a When step.`,
    );
  }
  if (!state.hasThen) {
    errors.push(
      `${resourceName}: ${state.currentScenario} requires a Then step.`,
    );
  }
}

function inspectLine(resourceName, state, errors, line) {
  const trimmed = line.trim();
  if (trimmed.startsWith('"""') || trimmed.startsWith("```")) {
    state.insideDocString = !state.insideDocString;
    return;
  }
  if (state.insideDocString || trimmed.startsWith("#")) return;
  if (trimmed.toLowerCase().startsWith("feature:")) {
    state.hasFeature = true;
  } else if (scenarioPattern.test(trimmed)) {
    finishScenario(resourceName, state, errors);
    state.scenarioCount += 1;
    state.currentScenario = trimmed;
    state.hasWhen = false;
    state.hasThen = false;
  } else if (
    state.currentScenario !== undefined &&
    startsStep("When", trimmed)
  ) {
    state.hasWhen = true;
  } else if (
    state.currentScenario !== undefined &&
    startsStep("Then", trimmed)
  ) {
    state.hasThen = true;
  }
}

/** Validates structural Gherkin and canonical layer exemptions in one feature source. */
export function validateFeatureSource(resourceName, source) {
  const errors = [];
  const state = {
    hasFeature: false,
    hasThen: false,
    hasWhen: false,
    insideDocString: false,
    scenarioCount: 0,
  };
  for (const line of source
    .replaceAll("\r\n", "\n")
    .replaceAll("\r", "\n")
    .split("\n")) {
    inspectLine(resourceName, state, errors, line);
  }
  finishScenario(resourceName, state, errors);
  errors.push(...validateExemptionPolicy(resourceName, source));
  if (!state.hasFeature)
    errors.push(`${resourceName}: missing Feature: declaration.`);
  if (state.scenarioCount === 0) {
    errors.push(`${resourceName}: feature must contain at least one scenario.`);
  }
  return errors;
}

function validateExemptionPolicy(resourceName, source) {
  const errors = [];
  const lines = source
    .replaceAll("\r\n", "\n")
    .replaceAll("\r", "\n")
    .split("\n");
  let pending = [];
  lines.forEach((line, index) => {
    const trimmed = line.trim();
    const lineNumber = index + 1;
    if (trimmed.startsWith("@")) {
      pending.push(
        ...validateTagLine(resourceName, trimmed, lineNumber, errors),
      );
      return;
    }
    if (gherkinDeclaration(trimmed)) {
      errors.push(
        ...declarationErrors(resourceName, lines, pending, trimmed, lineNumber),
      );
      pending = [];
      return;
    }
    if (trimmed !== "" && !trimmed.startsWith("#") && pending.length > 0) {
      errors.push(
        `${resourceName}:${lineNumber}: tags must be followed by their Gherkin declaration.`,
      );
      pending = [];
    }
  });
  if (pending.length > 0)
    errors.push(
      `${resourceName}: dangling tags are not attached to a scenario.`,
    );
  return errors;
}

function validateTagLine(resourceName, line, lineNumber, errors) {
  return line
    .split(/\s+/u)
    .filter((part) => part.startsWith("@"))
    .map((part) => part.slice(1))
    .map((name) => {
      if (forbiddenTags.has(name)) {
        errors.push(
          `${resourceName}:${lineNumber}: @${name} is forbidden; use only canonical exemption tags.`,
        );
      }
      return { line: lineNumber, name };
    });
}

function declarationErrors(
  resourceName,
  lines,
  pending,
  declaration,
  lineNumber,
) {
  const exemptions = pending.filter(({ name }) => exemptionTags.has(name));
  if (exemptions.length === 0) return [];
  const errors = exemptions.flatMap((tag) =>
    documentedExemptionErrors(resourceName, lines, tag),
  );
  if (!scenarioPattern.test(declaration)) {
    errors.push(
      `${resourceName}:${lineNumber}: exemption tags may only annotate a Scenario or Scenario Outline.`,
    );
  }
  return errors;
}

function documentedExemptionErrors(resourceName, lines, exemption) {
  const match = exemptionCommentPattern.exec(
    lines[exemption.line - 2]?.trim() ?? "",
  );
  const layer = exemption.name.replace("-exempt", "");
  if (match === null || match[1] !== layer) {
    return [
      `${resourceName}:${exemption.line}: @${exemption.name} requires the immediately preceding canonical comment.`,
    ];
  }
  const errors = [];
  if (invalidReasonPattern.test(match[2] ?? "")) {
    errors.push(
      `${resourceName}:${exemption.line}: exemption reason cannot be difficulty, speed, cost, flakiness, or unfinished work.`,
    );
  }
  if (!/^[a-z0-9-]+:test(?::[a-z0-9-]+)*\s+\/\s+\S/iu.test(match[3] ?? "")) {
    errors.push(
      `${resourceName}:${exemption.line}: alternative proof must name an Nx test target and scenario after ' / '.`,
    );
  }
  return errors;
}

function gherkinDeclaration(line) {
  return /^(?:Feature|Rule|Background|Scenario(?: Outline| Template)?|Examples?|Example):/iu.test(
    line,
  );
}

/** Recursively returns predicate-matching files in deterministic order. */
export async function findFiles(root, predicate) {
  const entries = await readdir(root, { withFileTypes: true });
  const files = await Promise.all(
    entries.map((entry) => {
      const entryPath = path.join(root, entry.name);
      if (entry.isDirectory()) return findFiles(entryPath, predicate);
      return Promise.resolve(
        entry.isFile() && predicate(entry.name) ? [entryPath] : [],
      );
    }),
  );
  return files.flat().toSorted();
}

/** Validates a deterministic list of canonical feature files. */
export async function validateFeatureFiles(featureFiles) {
  const results = await Promise.all(
    featureFiles.map(async (featureFile) =>
      validateFeatureSource(featureFile, await readFile(featureFile, "utf8")),
    ),
  );
  return results.flat().toSorted();
}

/** Returns canonical scenario names and structured alternative-proof references. */
export function inspectFeatureReferences(resourceName, source) {
  const scenarios = [];
  const proofs = [];
  const lines = source
    .replaceAll("\r\n", "\n")
    .replaceAll("\r", "\n")
    .split("\n");
  lines.forEach((line, index) => {
    const declaration =
      /^(?:Scenario(?: Outline| Template)?|Example):\s*(.+)$/iu.exec(
        line.trim(),
      );
    if (declaration?.[1]) scenarios.push(declaration[1]);
    const match = exemptionCommentPattern.exec(line.trim());
    if (match?.[3]) {
      const [target, scenario] = match[3].split(/\s+\/\s+/u, 2);
      if (target && scenario)
        proofs.push({ line: index + 1, resourceName, scenario, target });
    }
  });
  return { proofs, scenarios };
}

/** Validates that every structured alternative proof names an existing target and scenario. */
export function validateAlternativeProofReferences(
  proofs,
  projectTargets,
  scenarioNames,
) {
  const findings = [];
  for (const proof of proofs) {
    if (!projectTargets.has(proof.target)) {
      findings.push(
        `${proof.resourceName}:${proof.line}: alternative-proof target ${proof.target} does not exist.`,
      );
    }
    if (!scenarioNames.has(proof.scenario)) {
      findings.push(
        `${proof.resourceName}:${proof.line}: alternative-proof scenario ${proof.scenario} does not exist.`,
      );
    }
  }
  return findings.toSorted();
}

/** Returns the owner-adapter rows required by each canonical scenario declaration. */
export function inspectOwnerCoverageRequirements(resourceName, source) {
  const requirements = [];
  const lines = source
    .replaceAll("\r\n", "\n")
    .replaceAll("\r", "\n")
    .split("\n");
  let pendingTags = [];
  for (const line of lines) {
    const trimmed = line.trim();
    if (trimmed.startsWith("@")) {
      pendingTags.push(
        ...trimmed
          .split(/\s+/u)
          .filter((part) => part.startsWith("@"))
          .map((part) => part.slice(1)),
      );
      continue;
    }
    const scenario =
      /^(?:Scenario(?: Outline| Template)?|Example):\s*(.+)$/iu.exec(trimmed);
    if (scenario?.[1]) {
      requirements.push({
        integrationExempt: pendingTags.includes("integration-exempt"),
        key: `${resourceName}:${scenario[1]}`,
      });
      pendingTags = [];
      continue;
    }
    if (gherkinDeclaration(trimmed)) pendingTags = [];
  }
  return requirements;
}

/** Returns every exact `@covers feature:scenario` marker in an adapter source. */
export function inspectCoverageMarkers(resourceName, source) {
  const markers = [];
  const pattern = /@covers\s+(specs\/apps\/[^:\s]+\.feature):([^\r\n]+)/gu;
  for (const match of source.matchAll(pattern)) {
    markers.push({ key: `${match[1]}:${match[2].trim()}`, resourceName });
  }
  return markers;
}

/** Proves exactly one Unit row and one non-exempt Integration row per scenario. */
export function validateOwnerAdapterCoverage(
  requirements,
  unitMarkers,
  integrationMarkers,
) {
  const findings = [];
  const requiredKeys = new Set(requirements.map(({ key }) => key));
  const count = (markers, key) =>
    markers.filter((marker) => marker.key === key).length;

  for (const requirement of requirements) {
    const unitCount = count(unitMarkers, requirement.key);
    if (unitCount !== 1) {
      findings.push(
        `${requirement.key}: expected exactly one Unit adapter marker, found ${unitCount}.`,
      );
    }
    const integrationCount = count(integrationMarkers, requirement.key);
    const expectedIntegrationCount = requirement.integrationExempt ? 0 : 1;
    if (integrationCount !== expectedIntegrationCount) {
      findings.push(
        `${requirement.key}: expected ${expectedIntegrationCount} Integration adapter marker(s), found ${integrationCount}.`,
      );
    }
  }

  for (const [layer, markers] of [
    ["Unit", unitMarkers],
    ["Integration", integrationMarkers],
  ]) {
    for (const marker of markers) {
      if (!requiredKeys.has(marker.key)) {
        findings.push(
          `${marker.resourceName}: ${layer} adapter marker references unknown scenario ${marker.key}.`,
        );
      }
    }
  }
  return findings.toSorted();
}
