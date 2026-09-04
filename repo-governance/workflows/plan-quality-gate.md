---
tldr: "Produces one bounded semantic verdict for a formal plan's execution readiness."
when_to_use: "Use only for a checkpoint the owner explicitly requests."
---

# Plan Quality Gate

Run only when the owner explicitly names this gate or unambiguously directs its semantic audit. Do not infer
authorization from creating, editing, reviewing, or executing a plan, Plan mode, or another workflow. An instruction may
authorize multiple named checkpoints; otherwise it authorizes one run.

Produce exactly one terminal result: `PASS` or one `BLOCKED_*` variant for one formal plan. When authorized, run at the
directed pre-execution, post-material-change, or completion checkpoint. Never recurse or automatically start another
run.

## Sufficiency and Ownership

`PASS` means sufficient for the authorized scope, known risks, and applicable rules, not perfect, exhaustive, or
future-proof. Do not block on stylistic preference, speculative hardening, optional detail, or an improvement that can
wait without making execution unsafe or ambiguous. Apply [minimum sufficiency](../principles/minimum-sufficiency.md).

This workflow evaluates meaning, consistency, safety, executability, and proof. Deterministic tooling owns
machine-decidable checks, including links, indexes, word budgets, formatting, harness parity, and automated contracts.
Do not manually reproduce, sample, or second-guess those checks. Run canonical tooling only during verification and
consume its findings. For a check delivered by the plan, verify before execution that `delivery.md` contains an
executable implementation and proof task; at completion, require the target to exist and pass.

## Snapshot and Ledger

Freeze the plan path and stage, Git revision and dirty paths, scope, relevant specifications and governance, unresolved
decisions, authorization, and cycle `1`. A material external-input change ends the run as `BLOCKED_INPUT_CHANGED`; it
never causes an automatic restart. Repairs recorded by this run do not trigger another quality-gate run.

Audit before editing. Freeze a finite ledger containing `ID`, canonical rule, location, material gap, required repair,
proof, and `OPEN`, `FIXED`, `NOT_APPLICABLE`, or `BLOCKED`. Admit only rule violations or gaps making scoped execution
unsafe, ambiguous, or unprovable. Mandatory findings cannot be waived; `NOT_APPLICABLE` requires evidence. Preserve the
snapshot, cycle, ledger, pending verification, and authorization under
[governance continuity](../principles/governance-continuity.md).

## Bounded Procedure

1. Recursively inventory and read the plan, assets, relevant implementation, specifications, governance, and active plan
   conflicts. Do not validate machine-owned concerns.
2. Complete one semantic audit without edits. Check:
   - lifecycle truth, one stage, required documents, one technical shape, and truthful status under
     [plan organization](../conventions/plans-organization-policy.md);
   - coherent purpose, decision, scope, risks, acceptance, and a junior-readable BRD/PRD-to-delivery route;
   - necessary non-placeholder artifacts with distinct reader jobs;
   - synchronized architecture, Gherkin, file impact, dependencies, and applicable
     [software-quality routes](../development/software-quality-enforcement.md);
   - executable ownership, acceptance traceability, RED/GREEN/REFACTOR tasks, checkpoints, evidence, cleanup, recovery,
     and rollback; and
   - applicable migration, UI, isolation, and live-service contracts, plus conflicts with current rules or plans.
3. Freeze the initial ledger. Repair only its findings in dependency and safety order. Each repair closes one `OPEN` row
   without expanding product scope. Missing decisions, authority, or irreconcilable rules become `BLOCKED`; never invent
   answers.
4. Verify semantically in read-only mode, reviewing only repaired meaning and cross-document effects. Then run:

   ```sh
   rtk ./resource-guard run --class ephemeral --disk-path . -- npm exec -- nx run -p badakmini-cli -t test:repo
   ```

5. Return `PASS` when no row is `OPEN` or `BLOCKED`, tooling passes, no new material semantic gap appears, and the
   snapshot changed only through recorded repairs.
6. Otherwise allow exactly one stabilization cycle. Add only repair-caused semantic gaps and deterministic-tool
   findings, set cycle `2`, repair them once, and repeat step 4. A fixed finding cannot reopen without changed input.
7. After cycle `2`, return `PASS` if step 5 holds. Otherwise return `BLOCKED_NON_CONVERGENT` with remaining ledger and
   evidence. Do not repair, restart, or invoke this workflow again automatically.

Canonical resource-guard recovery is infrastructure handling, not another quality-gate cycle.

If canonical tooling cannot obtain a deterministic verdict, return `BLOCKED_TOOLING` with its evidence. Never simulate
the check or retry without bound. `PASS` authorizes neither execution nor commit/push. Every blocker names the reason,
remaining rows, and required external change. Resume only after new input and explicit owner direction authorize a fresh
gate run; [plan execution](plan-execution.md) consumes its result but never starts it.
