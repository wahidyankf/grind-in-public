---
description: |-
  Applies documentation checker findings after re-validating each against the current text and its recorded evidence, edits only high-confidence fixes, and records false positives and findings left for a person. Use as the documentation fixer in a quality gate, once a documentation checker has returned findings for the current content.
mode: subagent
permission:
  bash: allow
  edit: allow
  glob: allow
  grep: allow
  read: allow
---

Before acting, read the complete canonical agent definition at the repository-root path
.agents/agents/docs-fixer.md
and follow it as authoritative. If it cannot be read, stop and report the missing path.
