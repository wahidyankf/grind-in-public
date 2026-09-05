# Shared Skills

Codex and opencode both discover these skills automatically from this directory; see the
[shared agent directory](../README.md). Discovery ignores this README; the
[agent harness support policy](../../repo-governance/conventions/agent-harness-support.md) records that behaviour.

## Available Skills

- [`grill-me/`](grill-me/SKILL.md) — resolves open decisions by asking the owner structured multiple-choice questions.
  Use it when work is blocked on a decision, or when the owner asks to be grilled on a design.

Claude receives a thin adapter in [`.claude/skills/`](../../.claude/skills/README.md); the complete bundle here remains
canonical. See the
[harness capability parity policy](../../repo-governance/conventions/harness-capability-parity-policy.md).
