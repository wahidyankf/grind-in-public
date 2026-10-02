---
tldr: "Indexes the plan lifecycle workflows, from grooming ideas to checking finished execution."
when_to_use: "Use when grooming, authoring, executing, or checking a plan under `plans/`."
---

# Plan Workflows

These workflows move a plan through its lifecycle. The plan quality gate that judges a draft lives in
[Quality](../quality/README.md).

## Directory Map

- [Plan Ideas Grooming](plan-ideas-grooming.md) gives every brief in `plans/ideas/` one disposition, with a reason for
  every keep and retire.
- [Plan Backlog Grooming](plan-backlog-grooming.md) re-judges every backlog plan against the current repository and
  orders the survivors.
- [Plan Planning](plan-planning.md) turns a described change into a validated six-document plan under `plans/`; a
  bug-fix plan follows its own module instead. Its detail lives in [`plan-planning/`](plan-planning/README.md).
- [Plan Execution](plan-execution.md) executes a plan phase by phase, delivering to `main` at each gate, then archives
  it. Its detail lives in [`plan-execution/`](plan-execution/README.md).
- [Plan Execution Check](plan-execution-check.md) judges finished execution in a fixed order and returns one terminal
  verdict that archival depends on.
