# opencode Skills

This directory is intentionally empty of shared skills. opencode loads the canonical `skills/*/SKILL.md` bundles from
`.agents/`; see the
[harness capability parity policy](../../repo-governance/conventions/harness-capability-parity-policy.md). A native copy
here would create a second prompt body and is therefore prohibited.

## Available Skills

None. The canonical shared skills live in [`.agents/skills/`](../../.agents/skills/README.md).

Do not add a skill here. The parity policy requires all three harnesses to expose each canonical bundle, and
`npm run check:harness-parity` rejects native copies and missing adapters. Every skill starts in `.agents/skills/`;
Claude's required thin adapter lives separately under `.claude/skills/`.
