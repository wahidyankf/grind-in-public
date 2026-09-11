# Claude Code Subagents

Claude Code discovers these agents automatically from this directory. It ignores this README because the file carries no
agent frontmatter.

## Available Agents

- [`drill-reviewer.md`](drill-reviewer.md) — reviews a finished practice drill for correctness, complexity, edge cases,
  and explanation quality. Use it after solving an exercise by hand, when you want feedback rather than an answer; it
  never writes the solution.
- [`repo-explorer.md`](repo-explorer.md) — read-only explorer that reports where code, tests, documentation, and
  governance rules live. Use it to locate things or check which rule applies before making a change; it never edits
  anything.
- [`plan-maker.md`](plan-maker.md) — authors a formal plan end to end, runs both decision gates, and repairs its own
  draft within the declared budget. Use it when a plan is requested and no draft exists.
- [`plan-checker.md`](plan-checker.md) — audits a complete draft against the plan specification and returns findings
  with a terminal verdict. It modifies nothing.
- [`plan-execution-checker.md`](plan-execution-checker.md) — audits finished execution in fixed order and returns the
  verdict that permits or blocks archival.

Each file is a native adapter to the canonical role in [`.agents/agents/`](../../.agents/agents/README.md). Codex and
opencode provide their own adapters; see the
[harness capability parity policy](../../repo-governance/conventions/harness-capability-parity-policy.md).
