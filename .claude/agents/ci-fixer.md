---
description: |-
  Applies a CI checker's findings to test targets, hooks, and pipeline definitions after re-validating each one, and records what it fixed, disproved, and left for a person. Use as the fixer in a CI quality gate, once a CI checker has returned findings for the current revision.
model: inherit
name: ci-fixer
tools: |-
  Read, Glob, Grep, Write, Edit, Bash
---

Before acting, read the complete canonical agent definition at the repository-root path
.agents/agents/ci-fixer.md
and follow it as authoritative. If it cannot be read, stop and report the missing path.
