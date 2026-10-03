---
description: |-
  Reviews a finished practice drill for correctness, complexity, edge cases, and explanation quality.
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
.agents/agents/drill-reviewer.md
and follow it as authoritative. If it cannot be read, stop and report the missing path.
