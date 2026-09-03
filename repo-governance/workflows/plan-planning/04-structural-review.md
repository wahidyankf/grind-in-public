---
tldr: "Lists the structural checks a plan must pass before the quality gate runs."
when_to_use: "Use after authoring a plan and before handing it to plan-quality-gate."
---

# Structural Review

These checks are mechanical. Running them before the [plan-quality-gate](../plan-quality-gate.md) workflow keeps the
gate focused on judgment rather than on missing headings.

## Checks

1. **Five core documents exist** — `README.md`, `brd.md`, `prd.md`, `tech-docs/README.md`, and `delivery.md`, in one
   folder named per the [plan naming rules](../../conventions/plans-organization-policy/plan-naming.md).
2. **Scope names projects by path** — `README.md` states which projects the plan touches, not a vague area.
3. **Every checkbox is one action** — no checkbox hides two verbs.
4. **Every checkbox names a path, a command, and an acceptance criterion**, or explains why one does not apply.
5. **Every checkbox carries an executor tag and relevant `[AC-…]` label**, and the file opens with the tag legend.
6. **Every phase ends with a gate** and a Pause Safety note, and every gate item states a command and its acceptance.
7. **Phase 0 records a baseline** and changes nothing else.
8. **Every behavior cycle binds one Gherkin scenario**, inlined verbatim.
9. **The final phase is Knowledge Capture**, and `learnings.md` exists or is created by it.
10. **`tech-docs/file-impact.md` covers every checklist path** with an `[E]`, `[N]`, `[M]`, or `[D]` label, and the
    technical map explains every companion's reader job.
11. **Required conditional companions exist** — specification changes for C4, Gherkin, or interface work; migration
    design for representation transitions; UI design and twelve assets for material UI work.
12. **Diagrams are ASCII**, and each covers one concern.
13. **No secret appears anywhere** — variables are named, values are not.
14. **Links resolve**, and every planning directory map lists its direct siblings.

## Outcome

Fix what fails, then hand the plan to the quality gate. A plan that fails several of these checks is usually a plan
whose delivery checklist was written before its technical set, and rewriting the checklist is cheaper than patching it.
