---
tldr: "Makes a plan name its proposed C4 and Gherkin changes without confusing them with as-built specifications."
when_to_use:
  "Use when a formal plan changes observable behavior, architecture, an interface, or an executable specification."
---

# Specification Changes

`tech-docs/specification-changes.md` owns planned specification work and is linked from `tech-docs/README.md`. It
records a proposal; `specs/` remains the canonical as-built truth until implementation and verification finish.

Before its file list, separate durable contracts from plan-only acceptance outcomes. A plan-only outcome states why it
does not become C4 or Gherkin and names its exact `delivery.md` proof. Do not copy a PRD mechanically into `specs/`.

For every affected C4 or Gherkin file, give its exact repository-relative path and exactly one label: `[E]` edit, `[N]`
new, `[M]` move, or `[D]` delete. For Gherkin, name every preserved, changed, moved, deleted, and new scenario; new
scenarios name their user, preconditions, action, and expected outcome. State the unit, integration, behavior, and E2E
binding or adapter path that changes, or explain the specifically incapable layer and why. Name the target and focused
journey that proves the result.

For C4, name the exact view, node, relationship, data store, or constraint that changes and why. A planned diagram stays
terminal-first ASCII. Update the C4 model only with the final implemented boundary, never with a speculative design.

`tech-docs/file-impact.md` lists every expected code, test, specification, documentation, configuration, and runtime
path exactly with `[E]`, `[N]`, `[M]`, or `[D]`. Do not substitute a directory, glob, ellipsis, or generic area. An
unknown necessary filename is a discovery prerequisite that blocks execution, not a blank to hide in the tree.
