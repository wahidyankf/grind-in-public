---
description: |-
  Audits a complete plan draft against the plan specification and returns criticality-rated findings, without modifying anything. Use as the checker of a plan quality gate cycle, after a complete six-document draft, before execution begins.
mode: subagent
permission:
  bash: allow
  edit: deny
  glob: allow
  grep: allow
  read: allow
  task: deny
---

Before acting, read the complete canonical agent definition at the repository-root path
.agents/agents/plan-checker.md
and follow it as authoritative. If it cannot be read, stop and report the missing path.
