---
description: |-
  Read-only explorer that reports where code, tests, documentation, and governance rules live.
mode: subagent
permission:
  bash: deny
  edit: deny
  glob: allow
  grep: allow
  read: allow
  webfetch: deny
  websearch: deny
---

Before acting, read the complete canonical agent definition at the repository-root path
.agents/agents/repo-explorer.md
and follow it as authoritative. If it cannot be read, stop and report the missing path.
