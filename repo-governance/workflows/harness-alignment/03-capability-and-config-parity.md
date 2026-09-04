---
tldr:
  "Compares the shared subagents, skills, and commands, then each project config and directory index, and states what
  parity does not prove."
when_to_use: "Use after classifying derivative text, to check capability parity, harness configs, and READMEs."
---

# Capability and Config Parity

## Shared Capabilities

Compare the shared subagents, skills, and commands as the
[harness capability parity policy](../../conventions/harness-capability-parity-policy.md) requires. Each description
must take the two-part form the [agent harness support policy](../../conventions/agent-harness-support.md) requires.

## Configs and Indexes

Confirm each project config holds only documented settings, and that each directory is indexed as the
[documentation index policy](../../documentation-index-policy.md) requires, exemptions included.

## What Parity Does Not Prove

That proof is narrower than it looks. The parity check shows that the three harnesses expose equal shared capabilities,
but never that their instructions match canonical policy. The read-only [rules quality gate](../rules-quality-gate.md)
therefore still inspects affected entry points semantically.
