# Business Rationale

## Why This Work Exists

The stack packs were adopted as written for TypeScript, JavaScript, React, Next.js, and shell, and adapted for Python,
Go, and Nx. Adoption made the standards this repository's rules, but five of their gates have no enforcement here yet.
The adapter records them in one sentence under Deviations:

> Known gaps, left for a separate plan rather than fixed by adoption: TypeScript lint runs Biome without type-aware
> rules; no JavaScript file is type-checked through JSDoc; the Python pilot has no formatter, linter, or coverage gate;
> project `tags` are empty; the two `.github/scripts/` tests run under `sh`.

A rule without a gate is enforced only when a reviewer remembers it. Each of these five leaves a class of defect that
the adopted standard exists to catch and that nothing here catches today:

- a floating promise or an unsafe use of a library `any` in TypeScript passes `lint`;
- a wrong argument type in one of 23 authored JavaScript files passes every gate, including the scripts the gates
  themselves run;
- unformatted or unlinted Python lands, and nobody knows which Python branches the tests reach;
- a project can import from any other, because no boundary can be declared without tags; and
- the two CI test scripts run without `pipefail` in a dialect the shell standard does not declare here.

## Who It Affects

The owner, and every agent session that writes code under the adopted standards. The `swe-code-checker` agent reads the
standards and will report these gaps on every audit until a gate closes them, which turns a known, accepted gap into
recurring noise.

## What Success Means

The adapter's "Known gaps" sentence is gone, its two `gap` rows under Adopter Decisions are replaced by recorded
choices, and each gap has a named gate that fails on the defect it targets. Every gate green before the plan is green
after it. This is a judgement call about maintenance value, not a measured cost: the gaps have caused no recorded
incident.

## Non-Goals

- No application or script behaviour changes. Findings a new gate reports are fixed only as far as the gate needs.
- No new Nx plugin, executor, or generator; the
  [Nx workspace policy](../../../repo-governance/development/nx-workspace-policy.md) still excludes the `@nx`
  module-boundary lint rule, so boundaries are enforced by this repository's own validator.
- No change to the vendored `scripts/public-safety/` copy, which stays byte-identical.
- No 99% coverage floor for the Python pilot; its recorded exemption stands.
- No other stack gap. Anything else found while executing goes to `learnings.md` and is routed in Phase 6.

## Risks

- **Finding backlog.** Enabling type-aware lint or `checkJs` may surface more findings than one phase can fix. Each such
  phase carries a bounded checkpoint with a predeclared fallback rather than a waiver.
- **New dependencies.** U3 adds Python development tools and U1 may add a lint package; each is justified under the
  [dependency selection policy](../../../repo-governance/development/dependency-selection-policy.md) in the commit that
  adds it, and the owner may reject it at the planning gate.
- **Rule edits.** Every unit edits the adapter, and U3 and U4 edit further governance, so Rules Propagation runs on each
  of those phases.
