---
tldr: "Lists the instruction file and project config each supported harness reads."
when_to_use: "Use when configuring a harness or checking which file it reads."
---

# Supported Harnesses

| Harness     | Effective instructions                    | Project config          |
| ----------- | ----------------------------------------- | ----------------------- |
| Claude Code | `CLAUDE.md` imports canonical `AGENTS.md` | `.claude/settings.json` |
| Codex       | `AGENTS.md`                               | `.codex/config.toml`    |
| opencode    | `AGENTS.md`                               | `opencode.json`         |

Where each harness loads canonical skills and native custom-agent adapters belongs to the
[harness capability parity policy](../harness-capability-parity-policy.md).
