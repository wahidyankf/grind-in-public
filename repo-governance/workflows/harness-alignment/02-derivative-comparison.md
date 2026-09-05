---
tldr: "Checks that every harness reaches AGENTS.md directly or through the one permitted exact import."
when_to_use: "Use when checking instruction routing, harness config, or authority against AGENTS.md."
---

# Instruction Routing

## Sole Rule Body

Read `AGENTS.md` first. Confirm root `CLAUDE.md` contains only the exact import the
[instruction alignment policy](../../conventions/agent-instruction-alignment-policy.md) permits, and that no competing
instruction source exists. Resolve a contradiction at the canonical source through
[Rules Propagation](../rules/rules-propagation.md); never add a harness-specific rule body.

## Commands and Paths

Verify every command quoted in an instruction file exists in `package.json` or a `project.json` target, and that every
referenced path exists.

## Authority

Confirm every harness reaches `AGENTS.md` directly or through the exact import, with no overlay.
