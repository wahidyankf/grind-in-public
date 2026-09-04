---
tldr: "Applies one bounded rule transaction and consumes any read-only rules-gate ledger."
when_to_use: "Use automatically whenever a repository rule is created, changed, moved, or deleted."
---

# Rules Propagation

Apply this workflow automatically whenever a repository [rule](../../conventions/rule-definition-policy.md) is created,
changed, moved, or deleted, or an explicitly requested [rules quality gate](../rules-quality-gate.md) emits
`NEEDS_PROPAGATION`. No separate owner instruction is required. Propagation is the sole writer, never invokes the
quality gate, and consumes its frozen ledger when supplied. Edits inside one transaction do not start another
transaction.

## Inputs and Transaction

- the proposed rule and rationale;
- its intended mandatory, expected, or permitted strength;
- the people, agents, files, or tasks in scope;
- known enforcement and evidence routes; and
- an optional frozen quality-gate ledger and evidence.

Freeze those inputs, the Git revision and dirty paths, affected entry points, relevant canonical sources, pending
verification, and authorization. Preserve them under [governance continuity](../../principles/governance-continuity.md).
A material external-input change returns `BLOCKED_INPUT_CHANGED` and never restarts the transaction.

## Finite Procedure

1. Build one finite ledger from the requested outcome and any supplied gate ledger. Inspect only the affected rule, its
   points of use, higher authority, and directly overlapping guidance. Record each material gap as `OPEN`, `RESOLVED`,
   `NOT_APPLICABLE`, or `BLOCKED`; do not add style preferences, speculative hardening, or machine-owned checks.
2. Before editing, return `BLOCKED_INPUT` for a missing decision or authority and `BLOCKED_CONFLICT` for an
   irreconcilable higher-authority conflict. Otherwise apply the minimum repair needed to close every `OPEN` row:
   - put concise action at each applicable instruction entry point;
   - keep one canonical source and replace copies with links;
   - place stable standards in conventions, engineering rules in development, and procedures in workflows;
   - resolve lower-level conflicts by `principles > conventions > development > workflows`;
   - never resolve a same-level contradiction without the owner;
   - apply progressive disclosure and truthful enforcement classification; and
   - add automation only for an explicit need or demonstrated risk.
3. Read the repaired surfaces once for semantic closure. Resolve only repair-caused conflicts, using hierarchy and
   [minimum sufficiency](../../principles/minimum-sufficiency.md); never broaden the ledger, reopen a settled
   preference, or seek perfection. Every semantic row must now be closed or have returned the specific blocker from
   step 2.
4. Run:

   ```sh
   rtk ./resource-guard run --class ephemeral --disk-path . -- npm exec -- nx run -p badakmini-cli -t test:repo
   ```

5. On success, return `PASS_NO_CHANGE` when no edit was necessary or `PASS_CHANGED` otherwise. For transaction-caused
   deterministic findings, freeze their exact set, repair mechanically, and rerun step 4 only while the ordered measure
   of failing checks and reported violations or overage strictly decreases and no new failure class appears. Because
   that nonnegative measure decreases, recovery terminates. Return `BLOCKED_TOOLING` if progress stops, a new or
   unrelated failure appears, or canonical recovery cannot obtain a verdict.

Canonical resource-guard recovery is infrastructure handling, not another propagation transaction.

## Terminal Contract

The only results are `PASS_NO_CHANGE`, `PASS_CHANGED`, `BLOCKED_INPUT`, `BLOCKED_CONFLICT`, `BLOCKED_TOOLING`, and
`BLOCKED_INPUT_CHANGED`. Propagation repairs every authorized semantic row; anything it cannot decide or verify maps to
a specific external blocker. Passing means sufficient, not perfect or future-proof, and authorizes neither commit nor
push. Deterministic recovery obeys the strictly decreasing measure above. With unchanged inputs and repository state,
another transaction produces no diff.
