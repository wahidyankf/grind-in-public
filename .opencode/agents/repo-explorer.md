---
description: "Read-only explorer that reports where code, tests, documentation, and governance rules live. Use it to locate something, or to check which rule applies before making a change; it never edits anything."
mode: subagent
permission:
  bash: deny
  edit: deny
  glob: allow
  grep: allow
  read: allow
  task: deny
---

Before acting, read the complete canonical agent definition at the repository-root path .agents/agents/repo-explorer.md and follow it as authoritative. If it cannot be read, stop and report the missing path.
