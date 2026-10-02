---
tldr: "Indexes the quality-gate and sole-writer contracts every gate and propagation follows, and their local adapter."
when_to_use: "Use when running, writing, or changing a quality gate or a propagation."
---

# Workflow Contracts

These development standards govern every quality gate and propagation under
[`../../workflows/quality/`](../../workflows/quality/README.md).

## Directory Map

- [Quality Gate Contract](quality-gate-contract.md) fixes the bounded, advisory cycle every gate follows; its modules
  live in [`quality-gate-contract/`](quality-gate-contract/README.md).
- [Sole-Writer Propagation](sole-writer-propagation.md) gives each gate family exactly one writer.
- [Quality Gate Adapter](quality-gate-adapter.md) records this repository's families, subjects, checks, and deviations.
