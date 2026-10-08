---
tldr: "Records the supported agent harnesses and where each one reads instructions, config, and subagents."
when_to_use:
  "Use when adding, configuring, or changing support for an agent harness such as Claude Code, Codex, or opencode."
---

# Agent Harness Support Policy

## Scope

This repository supports four harnesses, in the sense the [agent vocabulary](agent-vocabulary.md) defines: the tool that
runs the model. `AGENTS.md` is the sole repository-rule body. Codex, OpenCode, and Command Code read it directly; Claude
Code reaches it through the exact `@AGENTS.md` import in `CLAUDE.md`.

## Details

Read the table the task needs rather than the whole policy:

- [Supported Harnesses](agent-harness-support/supported-harnesses.md) — the instruction file and project config each
  harness reads.
- [Directory Index Behaviour](agent-harness-support/directory-index-behaviour.md) — what each registering harness
  directory does with a `README.md`.

## Rules

Do not create a `CODEX.md` or `OPENCODE.md`. Both tools read `AGENTS.md` already, and a third instruction file would add
a copy to keep aligned for no gain. Add an instruction file only when its harness cannot read `AGENTS.md`, and then
follow the [agent instruction alignment policy](agent-instruction-alignment-policy.md).

The [harness capability parity policy](harness-capability-parity-policy.md) owns the capability directories, the parity
requirement, and the check that enforces it.

Keep each project config to settings that the tool documents and that the repository actually needs. A config file is
not a place for rules; rules belong in `AGENTS.md` or `repo-governance/`. Do not re-list an auto-discovered file as an
extra instruction. `opencode.json` declares repository permissions and keeps model selection inherited.

Write every agent description as two things: a short statement of what the agent does, then a short sentence saying when
it should be used. Harnesses route work by description, so an agent that omits its trigger will be picked at the wrong
time or never. Keep the wording equivalent across formats.

Index a harness directory with a `README.md`, as the [documentation index policy](../documentation-index-policy.md)
requires, and update it in the same change that adds, renames, or removes a file there. Which directories register a
file by name, and what each does with an index, is recorded in
[directory index behaviour](agent-harness-support/directory-index-behaviour.md).

A permission control that one harness lacks is recorded in the canonical definition, the native mapping, and the
supported-harness table. OpenCode Serena registration is deferred; Claude Code and Command Code retain semantic
integration. Codex 0.161 retains its existing guards; Serena and new destination routing are deferred. Its [release role
loader][role] omits MCP overrides, and native shell hooks omit the requested working directory. No role-local MCP
capability enforcement is claimed. Shell controls remain separate from semantic MCP access.

[role]: https://github.com/openai/codex/blob/979011409de0a60b52f179721948e65531d26144/codex-rs/core/src/agent/role.rs

## Verification

Run the [Harness Parity Verification](../workflows/quality/harness-parity-verification.md) workflow after changing any
instruction file, harness config, or subagent. Test a new tool's discovery behaviour yourself before relying on it,
because support changes between releases.
