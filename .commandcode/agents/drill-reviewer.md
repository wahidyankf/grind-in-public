---
description: |-
  Reviews a finished practice drill for correctness, complexity, edge cases, and explanation quality.
disallowedTools: |-
  agent, agent_output, write_file, edit_file, write_file, edit_file
name: drill-reviewer
tools: |-
  read_file, read_directory, grep, glob
---

Before acting, read the complete canonical agent definition at the repository-root path
.agents/agents/drill-reviewer.md and follow it as authoritative.
If it cannot be read, stop and report the missing path.
