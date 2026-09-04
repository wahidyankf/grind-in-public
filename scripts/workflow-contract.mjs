const required = {
  planGate: [
    "Run only when the owner explicitly",
    "cycle `2`",
    "BLOCKED_NON_CONVERGENT",
    "authorizes neither execution nor commit/push",
  ],
  propagation: [
    "sole writer, never invokes the quality gate",
    "strictly decreases",
    "PASS_NO_CHANGE",
    "PASS_CHANGED",
    "BLOCKED_INPUT",
    "BLOCKED_CONFLICT",
    "BLOCKED_TOOLING",
    "BLOCKED_INPUT_CHANGED",
  ],
  rulesGate: [
    "Run only when the owner explicitly",
    "NEEDS_PROPAGATION",
    "never ends blocked",
    "never ends blocked, repairs rules, reruns itself",
  ],
  taskTracking: [
    "separate RED, GREEN, and REFACTOR items",
    "expected behavioural RED reason",
    "REFACTOR-green results",
  ],
  tdd: [
    "living documentation",
    "compilation, configuration, or infrastructure failure is not RED evidence",
    "characterization tests",
    "automation proves the final state, not the historical sequence",
  ],
};

const forbidden = {
  propagation: ["PASS_READY", "BLOCKED_SEMANTIC", "BLOCKED_NON_CONVERGENT"],
  rulesGate: ["PASS_READY", "BLOCKED_SEMANTIC", "BLOCKED_NON_CONVERGENT"],
};

/** Validates stable, machine-decidable tokens in semantic workflow contracts. */
export function validateWorkflowContract(documents) {
  const findings = [];
  for (const [name, fragments] of Object.entries(required)) {
    const source = (documents[name] ?? "").replaceAll(/\s+/g, " ");
    for (const fragment of fragments) {
      if (!source.includes(fragment)) {
        findings.push(`${name}: missing contract fragment ${fragment}`);
      }
    }
  }
  for (const [name, fragments] of Object.entries(forbidden)) {
    const source = (documents[name] ?? "").replaceAll(/\s+/g, " ");
    for (const fragment of fragments) {
      if (source.includes(fragment)) {
        findings.push(`${name}: forbidden legacy result ${fragment}`);
      }
    }
  }
  return findings.toSorted();
}
