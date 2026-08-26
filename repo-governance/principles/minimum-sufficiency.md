---
tldr: "Requires the smallest responsible change and ends work after sufficient verification."
when_to_use: "Use when scoping, implementing, reviewing, or closing any repository task."
---

# Minimum Sufficiency

Do the least work that fully satisfies the explicit outcome and applicable repository rules. Work is a means, not evidence of progress: every added line, artifact, mechanism, and process step creates review, maintenance, security, and future-removal cost.

## Requirements

- Start with the explicit outcome and choose the smallest responsible change that reaches it.
- Prefer changing existing content, configuration, or mechanisms over adding code, dependencies, abstractions, validators, automation, infrastructure, or documents.
- Add an artifact or mechanism only when the requested outcome, an applicable rule, correctness, safety, or a demonstrated risk requires it.
- Do not generalize a one-time change or add speculative enforcement without evidence that it is needed.
- When code is necessary, keep it small and clear; every added line earns its continuing ownership cost.
- Keep verification proportional to the change and its risk, while always running every applicable mandatory gate. Stop once the outcome is achieved and required verification passes.

This principle never permits skipping TDD, safety boundaries, required documentation, governance propagation, or another applicable rule; those are necessary work. Among compliant options, choose the one with less lasting complexity. Apply [Maintenance Value](maintenance-value.md) as well: it decides whether a maintained surface earns its recurring cost, while this principle bounds the task's scope and stopping point.
