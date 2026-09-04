---
tldr: "Produces one read-only semantic verdict for a proposed or effective repository-rule state."
when_to_use: "Use before rules propagation and after its edits; this gate never writes rules."
---

# Rules Quality Gate

Produce one read-only semantic verdict for one proposed or effective repository-rule state. This workflow never edits
rules, invokes [rules propagation](rules/rules-propagation.md), or starts another gate run. Propagation is the sole
writer.

## Sufficiency and Ownership

A passing rule is sufficient for the stated need, scope, and known risk, not perfect, exhaustive, or future-proof. Do
not create findings for wording preference, speculative cases, optional explanation, or automation without a
demonstrated need. Apply [minimum sufficiency](../principles/minimum-sufficiency.md).

This gate owns semantic rule quality. Deterministic tooling owns machine-decidable checks, including links, indexes,
word budgets, formatting, coding-harness parity, and automated contracts. Do not manually reproduce, sample, or
second-guess those checks. For a proposed check not yet implemented, Proposal mode verifies only explicit ownership,
executable delivery, and proof obligations. Effective mode requires the canonical target to exist and pass; never
simulate future tooling.

## Modes, Snapshot, and Ledger

Run in exactly one mode:

- `PROPOSAL` compares the requested outcome with current effective rules before edits.
- `EFFECTIVE` evaluates the repository after propagation edits.

Freeze the mode, requested outcome and rationale, intended normative strength, scope and consumers, proposed move or
deletion, relevant canonical sources and hierarchy, enforcement route, Git revision, and dirty paths. A material
external-input change returns `BLOCKED_INPUT_CHANGED`; it never restarts the gate.

Audit without editing. Record a finite ledger containing `ID`, canonical source, material semantic gap, required
resolution, evidence, and `OPEN`, `RESOLVED`, `NOT_APPLICABLE`, or `BLOCKED`. Admit only a rule violation or a gap that
makes the requested outcome unsafe, contradictory, undiscoverable, or materially ambiguous. `NOT_APPLICABLE` requires
evidence. Preserve the snapshot, mode, ledger, evidence, and result under
[governance continuity](../principles/governance-continuity.md).

## Semantic Audit

Inspect only the affected rule, its point-of-use routes, relevant higher authority, and directly overlapping guidance.
Decide whether:

1. the need, intended outcome, and rationale are concrete enough to evaluate;
2. `must`, `should`, or `may` expresses the intended strength;
3. scope, trigger, action or prohibition, boundaries, and necessary exceptions are explicit;
4. the canonical governance level is correct and no lower rule conflicts with higher authority;
5. one canonical source owns the meaning while concise point-of-use links keep it discoverable without duplication;
6. every enforcement claim names its truthful class and route, with required evidence where automation cannot decide;
7. the instruction survives compaction and handoff at every applicable entry point;
8. a reasonable reader can act without inventing policy while the rule remains minimally sufficient; and
9. a move or deletion preserves unique intent and updates affected consumers.

Do not audit unrelated governance or rerun checks owned by deterministic tooling.

## Terminal Results

In `PROPOSAL` mode, run the canonical repository gate and return:

- `PASS_NO_CHANGE` when current effective meaning already satisfies the request;
- `PASS_READY` when a material change is necessary and fully specified; or
- `BLOCKED_INPUT`, `BLOCKED_CONFLICT`, `BLOCKED_TOOLING`, or `BLOCKED_INPUT_CHANGED` with the finite ledger, evidence,
  and required external decision.

In `EFFECTIVE` mode, run:

```sh
rtk npm exec -- nx run badakmini-cli:test:repo
```

Return `PASS_EFFECTIVE` only when the semantic ledger is clear and tooling passes. Otherwise return `BLOCKED_SEMANTIC`,
`BLOCKED_TOOLING`, or `BLOCKED_INPUT_CHANGED` with evidence. Every result ends the invocation. The gate never repairs,
invokes propagation, retries without bound, or authorizes commit/push.
