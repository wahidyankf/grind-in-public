// Each gate keeps the bounded, advisory contract: explicit entry, a cycle
// ceiling of 3, the four verdicts, and a caller that is never stopped. The
// rules writer keeps automatic entry and leaves delivery to its caller.
const gate = [
  "The gate starts only on an explicit owner request",
  "| `max-cycles` | integer | 1, 2, or 3 |",
  "PASS_WITH_FINDINGS",
  "No verdict stops the caller",
];

const required = {
  planGate: gate,
  propagation: [
    "This is the `rules` family's sole writer",
    "Entry is automatic",
    "The run never commits",
  ],
  rulesGate: gate,
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

// Results and inputs of the retired unbounded and recursive gate models.
const legacy = [
  "PASS_READY",
  "BLOCKED_SEMANTIC",
  "BLOCKED_NON_CONVERGENT",
  "NEEDS_PROPAGATION",
  "max-iterations",
  "max-audits",
  "min-iterations",
];

const forbidden = {
  planGate: legacy,
  propagation: legacy,
  rulesGate: legacy,
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
