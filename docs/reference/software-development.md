---
tldr: "Maps each project to the stack packs it follows and where their standards, skills, and decisions live."
when_to_use: "Use when looking up which language, framework, or tooling rules apply to a project here."
---

# Software Development

Each project follows the stack packs its entry in the `extensions.software-development` inventory of
[`repo-config.yml`](../../repo-config.yml) lists. A pack is a stack standard under
[`repo-governance/development/quality/stacks/`](../../repo-governance/development/quality/stacks/README.md) plus its
skill under [`.agents/skills/`](../../.agents/skills/README.md). The
[repository adapter](../../repo-governance/development/quality/stacks/repository-adapter.md) records this repository's
choices and deviations, and each project README holds its commands.

| Project              | Path                      | Packs                                  |
| -------------------- | ------------------------- | -------------------------------------- |
| `wahidyankf-www`     | `apps/wahidyankf-www`     | TypeScript, JavaScript, React, Next.js |
| `wahidyankf-www-e2e` | `apps/wahidyankf-www-e2e` | TypeScript, JavaScript                 |
| `forum-be-python`    | `apps/forum-be-python`    | Python                                 |
| `go-tools`           | `tools`                   | Go                                     |
| `repo-scripts`       | `scripts`                 | JavaScript, Shell                      |
| `public-safety`      | `scripts/public-safety`   | Shell                                  |
| `opencode-plugin`    | `.opencode/plugin`        | JavaScript                             |
| `claude-hooks`       | `.claude/hooks`           | Shell                                  |
| `ci-scripts`         | `.github/scripts`         | Shell                                  |
| `workspace`          | `.`                       | Nx                                     |

The `swe-code-maker`, `swe-code-checker`, and `swe-code-fixer` agents in
[`.agents/agents/`](../../.agents/agents/README.md) read that inventory and load each listed pack's local skill and
standard before they work in a project. The inventory is informational: it holds no command or version.
