---
description: |-
  Applies code checker findings after re-validating each against the current code and its cited standard, makes only high-confidence fixes that keep pinned behaviour, writes the failing test first where a fix needs one, and records every disposition. Use once a code checker has returned findings for the current revision of the projects in scope, rather than when a command is failing or new behaviour is needed.
mode: subagent
permission:
  bash: allow
  edit: allow
  glob: allow
  grep: allow
  read: allow
---

Before acting, read the complete canonical agent definition at the repository-root path .agents/agents/swe-code-fixer.md and follow it as authoritative. If it cannot be read, stop and report the missing path.
