---
tldr: "Makes AGENTS.md the only repository rule body and requires exact harness routing to it."
when_to_use: "Use when creating, editing, or reviewing AGENTS.md, CLAUDE.md, or another instruction source."
---

# Agent Instruction Alignment Policy

`AGENTS.md` is the sole repository-owned agent instruction body. Claude Code receives it through a root `CLAUDE.md`
whose complete normalized content is exactly `@AGENTS.md`. Normalization permits a UTF-8 byte-order mark, CRLF or LF,
and surrounding blank lines only. Codex and opencode read `AGENTS.md` directly.

Do not add another always-on instruction source. Nested `AGENTS.md`, `AGENTS.override.md`, additional `CLAUDE.md`,
`.claude/rules/**/*.md`, and nonempty opencode `instructions` arrays are prohibited until this policy and deterministic
validation define an equivalent canonical route. User-global configuration, local overrides, caches, generated trees,
worktrees, links, and credentials remain outside the repository contract.

Tool configuration contains only documented operational settings, not rules. Detailed shared rules belong in
`repo-governance/` and are linked from `AGENTS.md`. A tool-specific need that changes agent behaviour is resolved at the
canonical source or recorded as native adapter metadata; it is never appended to the instruction body.

The [harness capability parity policy](harness-capability-parity-policy.md) owns canonical skills, agents, native
adapters, and deterministic contract verification. Follow [Harness Alignment](../workflows/harness-alignment.md) after
changing any instruction source or harness config.
