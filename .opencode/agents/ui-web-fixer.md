---
description: |-
  Executes UI Web Propagation on a frozen ui-web ledger, reproducing each row, repairing the interface test-first, specifying correct but unspecified behaviour, and leaving design and behaviour choices to their owner. Use as the writer's executor in a UI web quality gate cycle, once the checker's findings are frozen in a ledger, or when someone explicitly names rows of one to repair.
mode: subagent
permission:
  bash: allow
  edit: allow
  glob: allow
  grep: allow
  read: allow
---

Before acting, read the complete canonical agent definition at the repository-root path
.agents/agents/ui-web-fixer.md
and follow it as authoritative. If it cannot be read, stop and report the missing path.
