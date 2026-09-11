# Claude Code Skills

Claude Code discovers these skills automatically from this directory. It ignores this README; the
[agent harness support policy](../../repo-governance/conventions/agent-harness-support.md) records that behaviour.

## Available Skills

- [`grill-me/`](grill-me/SKILL.md) — resolves open decisions by asking the owner structured multiple-choice questions.
  Use it when work is blocked on a decision, or when the owner asks to be grilled on a design.
- [`plan-grooming-idea-briefs/`](plan-grooming-idea-briefs/SKILL.md) — judges whether a brief is worth promoting,
  keeping with a trigger, or retiring.
- [`plan-creating-project-plans/`](plan-creating-project-plans/SKILL.md) — writes the six documents so each answers its
  own question and the checklist is executable.
- [`plan-writing-gherkin-criteria/`](plan-writing-gherkin-criteria/SKILL.md) — writes acceptance scenarios that describe
  observable behaviour and can actually fail.
- [`plan-validating-quality/`](plan-validating-quality/SKILL.md) — judges a draft beyond what structural validation
  reaches.
- [`plan-verifying-execution/`](plan-verifying-execution/SKILL.md) — judges finished execution against the repository
  rather than against the checklist.

Each entry is a thin adapter to the canonical bundle in [`.agents/skills/`](../../.agents/skills/README.md). opencode
reads the canonical directory directly, so it needs no third adapter; see the
[harness capability parity policy](../../repo-governance/conventions/harness-capability-parity-policy.md).
