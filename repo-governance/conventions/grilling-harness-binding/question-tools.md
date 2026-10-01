---
tldr: "Names the question tool each supported harness exposes and whether it is verified here."
when_to_use: "Use when asking a decision question in a harness, or recording the tool a new harness exposes."
---

# Question Tools

Each supported harness exposes one tool for a structured question:

- **Claude Code**
  - Tool: `AskUserQuestion`
  - Notes: Verified here. The harness appends a free-text entry, which serves as the Rule 8 write-in.
- **Codex**
  - Tool: `request_user_input`
  - Notes: Needs `default_mode_request_user_input` in `.codex/config.toml`; Codex marks the feature as in development,
    so its effect is unverified here.
- **opencode**
  - Tool: `question`
  - Notes: Documented by opencode; unverified here.

A tool that returns a structured choice removes the parsing step, which is why Rule 6 prefers it over prose. Only a
session with no such tool falls back to [Markdown](markdown-fallback.md).
