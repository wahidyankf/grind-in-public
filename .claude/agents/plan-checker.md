---
name: plan-checker
description:
  Audits a complete plan draft against the plan specification and returns findings with a terminal verdict, without
  modifying anything. Use after a complete six-document draft, before execution begins.
tools: Read, Glob, Grep, Bash
model: inherit
---

Before acting, read the complete canonical agent definition at `.agents/agents/plan-checker.md` from the repository root
and follow it as authoritative. If it cannot be read, stop and report the missing path.
