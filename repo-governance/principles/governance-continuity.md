---
tldr: "Preserves active governance state through compaction, handoff, and resumed work."
when_to_use: "Use when a governed workflow spans context compaction, agent handoff, or an interrupted session."
---

# Governance Continuity

Preserve the active workflow's authority, authorization, frozen inputs, decisions, task state, findings, evidence, and
pending verification through compaction or handoff. A shorter context may summarize those facts but must not silently
drop or reinterpret them.

On resume, reload the applicable canonical rule and verify the repository state before continuing. Do not restart a
bounded gate, reset its cycle counter, reopen a resolved finding, or infer new authorization because the harness context
changed.

A material input or repository-state change that invalidates a frozen snapshot ends the current gate with its defined
`BLOCKED_INPUT_CHANGED` result. Resume only when new input or authority starts a fresh invocation.

Task tracking records current execution state but never grants commit, push, deployment, or destructive authority.
