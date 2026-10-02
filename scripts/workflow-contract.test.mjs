import assert from "node:assert/strict";
import test from "node:test";

import { validateWorkflowContract } from "./workflow-contract.mjs";

function validDocuments() {
  return {
    planGate:
      "The gate starts only on an explicit owner request | `max-cycles` | integer | 1, 2, or 3 | PASS_WITH_FINDINGS No verdict stops the caller",
    propagation:
      "This is the `rules` family's sole writer Entry is automatic The run never commits",
    rulesGate:
      "The gate starts only on an explicit owner request | `max-cycles` | integer | 1, 2, or 3 | PASS_WITH_FINDINGS No verdict stops the caller",
    taskTracking:
      "separate RED, GREEN, and REFACTOR items expected behavioural RED reason REFACTOR-green results",
    tdd: "living documentation compilation, configuration, or infrastructure failure is not RED evidence characterization tests automation proves the final state, not the historical sequence",
  };
}

test("accepts the bounded plan, rules, and TDD workflow contracts", () => {
  assert.deepEqual(validateWorkflowContract(validDocuments()), []);
});

test("rejects a missing explicit authorization and TDD process evidence", () => {
  const documents = validDocuments();
  documents.planGate = documents.planGate.replace(
    "The gate starts only on an explicit owner request",
    "",
  );
  documents.taskTracking = "";
  const findings = validateWorkflowContract(documents);
  assert.ok(findings.some((finding) => finding.startsWith("planGate:")));
  assert.ok(findings.some((finding) => finding.startsWith("taskTracking:")));
});

test("rejects a gate that loses its three-cycle bound or advisory verdict", () => {
  const documents = validDocuments();
  documents.rulesGate = documents.rulesGate
    .replace("1, 2, or 3", "1 to 7")
    .replace("No verdict stops the caller", "");
  const findings = validateWorkflowContract(documents);
  assert.ok(findings.some((finding) => finding.includes("1, 2, or 3")));
  assert.ok(
    findings.some((finding) => finding.includes("No verdict stops the caller")),
  );
});

test("rejects legacy unbounded and recursive gate results", () => {
  const documents = validDocuments();
  documents.propagation += " PASS_READY BLOCKED_NON_CONVERGENT";
  documents.rulesGate += " NEEDS_PROPAGATION";
  documents.planGate += " max-iterations";
  const findings = validateWorkflowContract(documents);
  assert.ok(findings.some((finding) => finding.includes("PASS_READY")));
  assert.ok(findings.some((finding) => finding.includes("NEEDS_PROPAGATION")));
  assert.ok(
    findings.some((finding) => finding.includes("BLOCKED_NON_CONVERGENT")),
  );
  assert.ok(findings.some((finding) => finding.includes("max-iterations")));
});

test("returns findings in deterministic order", () => {
  const findings = validateWorkflowContract({});
  assert.deepEqual(findings, [...findings].toSorted());
});
