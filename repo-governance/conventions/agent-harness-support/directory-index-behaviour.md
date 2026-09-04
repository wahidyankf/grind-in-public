---
tldr: "Records what each registering harness directory does with a README.md index."
when_to_use: "Use when adding or moving a README inside a harness directory that registers files by name."
---

# Directory Index Behaviour

Verified behaviour for an index inside a registering directory:

| Directory                               | Index behaviour                                          |
| --------------------------------------- | -------------------------------------------------------- |
| `.claude/agents/`                       | Ignored without agent frontmatter                        |
| `.claude/commands/`                     | Becomes `/README`; no README, index from parent          |
| `.codex/agents/`                        | Ignored; only `*.toml` is read                           |
| `.opencode/agents/`                     | Becomes an agent unless frontmatter sets `disable: true` |
| Every `skills/` and `.opencode/plugin/` | Inert: discovery keys on `SKILL.md` or a code extension  |
