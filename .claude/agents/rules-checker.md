---
description: |-
  Audits a repository's rules as a whole for contradictions across levels, inaccurate references, inconsistent terms and strengths, missing traceability, and duplicated bodies, and returns rated findings without editing. Use as the checker of a rules quality gate, for a repository-wide consistency check of its rules, after a structural change to governance, or when two skills or agents may need merging.
model: inherit
name: rules-checker
tools: |-
  Read, Glob, Grep, Bash
---

Before acting, read the complete canonical agent definition at the repository-root path
.agents/agents/rules-checker.md
and follow it as authoritative. If it cannot be read, stop and report the missing path.
