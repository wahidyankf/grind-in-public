---
tldr: "Applies one bounded rule transaction while the read-only rules gate judges proposal and effective states."
when_to_use: "Use automatically whenever a repository rule is created, changed, moved, or deleted."
---

# Rules Propagation

Apply this workflow automatically whenever a repository [rule](../../conventions/rule-definition-policy.md) changes. It
is the sole writer for rule propagation and composes the read-only [rules quality gate](../rules-quality-gate.md).
Neither workflow invokes itself; edits inside one transaction do not start another transaction.

## Inputs and Transaction

Freeze the proposed outcome and rationale, normative strength, scope and consumers, known enforcement/evidence routes,
Git revision and dirty paths, affected entry points, relevant canonical sources, authorization, and cycle `1`. Preserve
the transaction under [governance continuity](../../principles/governance-continuity.md). A material external-input
change returns `BLOCKED_INPUT_CHANGED` and never restarts the transaction.

## Bounded Procedure

1. Run the [rules quality gate](../rules-quality-gate.md) in `PROPOSAL` mode.
   - `PASS_NO_CHANGE`: make no edits and return `PASS_NO_CHANGE`.
   - `PASS_READY`: continue with its finite ledger.
   - Any `BLOCKED_*`: preserve evidence and stop.
2. Apply one propagation cycle, changing only what the accepted outcome and ledger require:
   - put concise action at each applicable instruction entry point;
   - keep one canonical source and replace copies with links;
   - place stable standards in conventions, engineering rules in development, and procedures in workflows;
   - resolve lower-level conflicts by `principles > conventions > development > workflows`;
   - never resolve a same-level contradiction without the owner;
   - apply progressive disclosure and truthful enforcement classification; and
   - add automation only for an explicit need or demonstrated risk.
3. Run the rules quality gate in `EFFECTIVE` mode.
   - `PASS_EFFECTIVE`: return `PASS_CHANGED`.
   - `BLOCKED_TOOLING` or `BLOCKED_INPUT_CHANGED`: stop with evidence.
   - `BLOCKED_SEMANTIC`: permit one stabilization cycle only when every row is within the accepted outcome or caused by
     first-cycle edits.
4. Freeze the combined ledger, set cycle `2`, and repair only that set. Do not expand scope or invent authority.
5. Run `EFFECTIVE` once more. Return `PASS_CHANGED` on success; otherwise return `BLOCKED_NON_CONVERGENT` with remaining
   rows and evidence. Never repair, restart, or retry automatically.

`PASS_NO_CHANGE` and `PASS_CHANGED` mean sufficient, not perfect. They authorize neither commit nor push. Every
`BLOCKED_*` result names the remaining rows and external change required. One transaction invokes the rules gate at most
three times; unchanged inputs and repository state produce no new diff.
