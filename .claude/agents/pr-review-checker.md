---
description: |-
  Checks one pinned change for the pr-review family by coordinating one review pass: after the lens checkers report, it deduplicates, re-categorizes, filters, verifies, and rates raw findings for criticality, then publishes the one review bound to the pinned head, editing no file. Use as the checker of a PR review quality gate cycle, or at the synthesis step of any review pass on a pull request or a local commit range, once the selected lens checkers have returned their findings, or alone when the tier runs none.
model: inherit
name: pr-review-checker
tools: |-
  Read, Glob, Grep, Bash
---

Before acting, read the complete canonical agent definition at the repository-root path
.agents/agents/pr-review-checker.md
and follow it as authoritative. If it cannot be read, stop and report the missing path.
