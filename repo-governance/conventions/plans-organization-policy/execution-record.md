---
tldr: "Requires delivery.md to open with a dated record of phases, gates, failures, and retries."
when_to_use: "Use when starting, executing, or archiving a plan, or when reading an archived one."
---

# Execution Record

`delivery.md` opens with a dated Execution Record, above the tag legend. A checkbox states what was intended and a tick
states that it eventually held; neither states what happened in between.

Add a line when a phase completes, when a gate passes or fails, when a retry proves something the first attempt did not,
and when execution changed the plan. Date each line and write it as the event happens: reconstructed at archival, the
record says what the author already believed.

```markdown
## Execution Record

- 2026-08-31: Phase 1 gate passed; `npx nx run badakmini-cli:test:quick` green.
- 2026-08-31: Phase 2 stopped at the link check — a renamed policy left two dead links in `docs/`. Fixed at the source,
  reran green.
```

A plan that has not started carries the heading and no lines. An archived plan is never rewritten to add a record it
never kept: `plans/done/` is history, so this rule binds a plan while it is being executed.

A record written during execution then stays for good. `learnings.md` is drained and may be deleted, and inline notes
scatter the sequence across the file, so this is the one place a reader of an archived plan can see the order, the
failures, and the retries.

Every `plan-checker` prompt states this record's shape in the imperative, because a subagent prompt has to stand alone.
Change them in the same edit, in all three harness copies.
