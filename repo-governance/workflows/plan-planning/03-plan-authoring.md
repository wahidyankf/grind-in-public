---
tldr: "Describes how to write the five plan documents in order."
when_to_use: "Use when writing a plan's README, brd, prd, tech-docs, and delivery files."
---

# Plan Authoring

Write the five core documents in dependency order, because each one constrains the next. The
[five-document structure](../../conventions/plans-organization-policy/five-document-structure.md) defines what each file
owns.

## Order

1. **`brd.md`** — why the work is worth doing, who it affects, what success means, and what is deliberately out of
   scope. Write the reasoning honestly; label a judgment call as one. An invented metric here corrupts every decision
   downstream.
2. **`prd.md`** — user stories and Gherkin acceptance criteria, each with a stable `[AC-…]` label, per the
   [specs policy](../../development/specs-policy.md). Mark any rollout or operational criterion that remains plan-only
   and give its reason and delivery proof; durable behavior scenarios land in `specs/`.
3. **`tech-docs/README.md`** — architecture, design decisions with rationale, dependencies, risks, rollback path, and
   technical reading order. Write `file-impact.md` with every exact path and `[E]`, `[N]`, `[M]`, or `[D]` label; add
   the applicable specification-change, migration, and UI-design companions before authoring delivery.
4. **`delivery.md`** — the phased checklist, written last because it derives from the three documents above. Follow the
   [delivery checklist rules](../../conventions/plans-organization-policy/delivery-checklists.md) and the
   [phase and gate rules](../../conventions/plans-organization-policy/phases-and-gates.md).
5. **`README.md`** — written last despite being read first: context, scope with projects named, the approach in a
   paragraph, and a map linking every plan artifact.

## Delivery Checklist Shape

Phase 0 records the clean baseline. Each later phase ends at a gate and a Pause Safety note, and delivers to `main` once
that gate passes. The final phase before archival is Knowledge Capture; see the
[knowledge capture rules](../../conventions/plans-organization-policy/knowledge-capture.md).

Every behavior change is a RED → GREEN → REFACTOR cycle bound to exactly one Gherkin scenario, per the
[TDD policy](../../development/tdd-policy.md).

## Diagrams

Plans use terminal-first ASCII diagrams, with no Mermaid exception. Draw one diagram per architectural concern the plan
touches rather than one crowded diagram covering all of them.
