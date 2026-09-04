import assert from "node:assert/strict";
import test from "node:test";

import { validateWorkflowContract } from "./workflow-contract.mjs";

function validDocuments() {
  return {
    planGate:
      "Run only when the owner explicitly cycle `2` BLOCKED_NON_CONVERGENT authorizes neither execution nor commit/push",
    propagation:
      "sole writer, never invokes the quality gate strictly decreases PASS_NO_CHANGE PASS_CHANGED BLOCKED_INPUT BLOCKED_CONFLICT BLOCKED_TOOLING BLOCKED_INPUT_CHANGED",
    rulesGate:
      "Run only when the owner explicitly NEEDS_PROPAGATION never ends blocked, repairs rules, reruns itself",
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
    "Run only when the owner explicitly",
    "",
  );
  documents.taskTracking = "";
  const findings = validateWorkflowContract(documents);
  assert.ok(findings.some((finding) => finding.startsWith("planGate:")));
  assert.ok(findings.some((finding) => finding.startsWith("taskTracking:")));
});

test("rejects legacy recursive rules-gate results", () => {
  const documents = validDocuments();
  documents.propagation += " PASS_READY BLOCKED_NON_CONVERGENT";
  documents.rulesGate += " BLOCKED_SEMANTIC";
  const findings = validateWorkflowContract(documents);
  assert.ok(findings.some((finding) => finding.includes("PASS_READY")));
  assert.ok(findings.some((finding) => finding.includes("BLOCKED_SEMANTIC")));
  assert.ok(
    findings.some((finding) => finding.includes("BLOCKED_NON_CONVERGENT")),
  );
});

test("returns findings in deterministic order", () => {
  const findings = validateWorkflowContract({});
  assert.deepEqual(findings, [...findings].toSorted());
});
