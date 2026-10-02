---
description: |-
  Re-validates each repository rules finding against current files, applies only high-confidence repairs through Rules Propagation with the higher governance level as authority, and hands every open judgement to the rule's owner. Use as the writer's executor in a rules quality gate cycle, or once a rules checker has returned findings, as the repair step of a rules consistency check.
mode: subagent
permission:
  bash: allow
  edit: allow
  glob: allow
  grep: allow
  read: allow
---

Before acting, read the complete canonical agent definition at the repository-root path
.agents/agents/rules-fixer.md
and follow it as authoritative. If it cannot be read, stop and report the missing path.
