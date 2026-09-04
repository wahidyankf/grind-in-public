---
tldr: "Runs separate spec-aware exploratory and spec-blind usability passes for UI-affecting plans."
when_to_use: "Use after automation passes for every UI-affecting plan and before its completion checkpoint."
---

# Exploratory and Usability Testing

Run two independent passes over the application at its exact served origin and every supported viewport class. Use
isolated synthetic state under [test-data isolation](../development/test-data-isolation.md). Finish and record the
exploratory pass before beginning usability review so the lenses remain distinct.

## Exploratory Pass

Review the affected canonical Gherkin before testing. Probe edge cases, boundary conditions, URLs, state transitions,
and passive security signals beyond scripted E2E cases. Record route/state, category, result, and safe evidence under
`## Exploratory findings` in the plan's `learnings.md`.

A correct but unspecified behaviour becomes a proposed Gherkin scenario. Reconcile it through BDD and a new evidenced
red-green-refactor cycle; never add it directly to `specs/` without implementation proof.

## Usability Pass

Use a fresh agent context given only the exact origin, affected routes, and viewport classes. Withhold specifications,
source, and design assets so the pass is structurally spec-blind. If no fresh context is available, label the pass
spec-aware. Review first-time-user understanding, Nielsen's heuristics, cognitive walkthrough steps, empty/loading/error
states, keyboard and focus behaviour, and responsive usability. Record results separately under `## Usability findings`.

Cross-reference shared root causes instead of duplicating fixes. Reconcile accepted specification proposals through the
same BDD and TDD path. Before completion, confirm both sections exist or explicitly record that no findings were found.
Never record credentials, private values, or production data.
