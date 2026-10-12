# Muse Harness

This directory holds the project-scoped Muse configuration for Grind in Public. Muse reads its repository rules from
[`AGENTS.md`](../AGENTS.md), so no Muse-specific instruction file exists; see the
[agent harness support policy](../repo-governance/conventions/agent-harness-support.md).

## Contents

- `hooks.json` — the `PreToolUse` hook that forwards every tool call to the repository policy endpoint with
  `--harness muse`. Muse honours exit statuses and ignores Claude-style stdout JSON decisions, so the adapter stays
  silent with exit 0 on allow and prints the denial to stderr with exit 2 on deny; see [Hooks](hooks/README.md).
- [Hooks](hooks/README.md) — the policy adapter and its six-case regression.

Muse skills live outside this directory, in [`.agents/skills/`](../.agents/skills/README.md), because that is where
shared skills are kept.
