---
disable: true
description: Directory index for the opencode subagents. Not an agent.
---

# opencode Subagents

opencode turns every `*.md` filename in this directory into an agent, so this index sets `disable: true` to keep itself
out of the agent list.

## Available Agents

- [`drill-reviewer.md`](drill-reviewer.md) — reviews a finished practice drill for correctness, complexity, edge cases,
  and explanation quality. Use it after solving an exercise by hand, when you want feedback rather than an answer; it
  never writes the solution.
- [`repo-explorer.md`](repo-explorer.md) — read-only explorer that reports where code, tests, documentation, and
  governance rules live. Use it to locate things or check which rule applies before making a change; it never edits
  anything.

Each role is mirrored in [`.claude/agents/`](../../.claude/agents/README.md) and
[`.codex/agents/`](../../.codex/agents/README.md) and must stay at parity; see the
[harness capability parity policy](../../repo-governance/conventions/harness-capability-parity-policy.md).
