---
tldr: "Requires the smallest responsible change and ends work after sufficient verification."
when_to_use: "Use when scoping, implementing, reviewing, or closing any repository task."
---

# Minimum Sufficiency

Do the least work that fully satisfies the explicit outcome and applicable repository rules. Work is a means, not
evidence of progress: every added line, artifact, mechanism, and process step creates review, maintenance, security, and
future-removal cost.

## Understand Before Simplifying

Read the task, the relevant tests or specifications, and the code the change would touch. Trace the real flow end to end
before deciding where to intervene. A small diff chosen without understanding the responsible path is not sufficient.

After understanding the problem, stop at the first responsible option that reaches the outcome:

1. Make no change when the requested outcome already holds or is not needed.
2. Reuse an existing repository mechanism, helper, or established pattern.
3. Use the language's standard library.
4. Use a native platform capability.
5. Use a suitable dependency that is already installed.
6. Use one clear expression when it remains correct and readable.
7. Only then write the minimum new code.

Do not climb the list to introduce an unrequested abstraction, boilerplate, or speculative flexibility. Prefer deletion
over addition and familiar constructs over clever ones. When equally small alternatives exist, choose the one that
handles the real edge cases. Challenge a requested mechanism when a simpler existing facility reaches the underlying
outcome; resolve any owner decision that remains under the
[grilling-with-options policy](../conventions/grilling-with-options-policy.md).

## Requirements

- Start with the explicit outcome and choose the smallest responsible change that reaches it.
- Prefer changing existing content, configuration, or mechanisms over adding code, dependencies, abstractions,
  validators, automation, infrastructure, or documents.
- Add an artifact or mechanism only when the requested outcome, an applicable rule, correctness, safety, or a
  demonstrated risk requires it.
- Do not generalize a one-time change or add speculative enforcement without evidence that it is needed.
- When code is necessary, keep it small and clear; every added line earns its continuing ownership cost.
- Keep verification proportional to the change and its risk, while always running every applicable mandatory gate. Stop
  once the outcome is achieved and required verification passes.

Efficiency never discounts input validation at trust boundaries, error handling that prevents data loss, security,
accessibility, real-hardware calibration, or anything the outcome explicitly requires. Where no applicable testing
policy already prescribes stronger evidence, non-trivial logic must leave the smallest runnable check that would fail if
the logic broke. A trivial one-line change needs no dedicated test only when no other rule requires one.

This principle never permits skipping TDD, safety boundaries, required documentation, governance propagation, or another
applicable rule; those are necessary work. Among compliant options, choose the one with less lasting complexity. Apply
[Maintenance Value](maintenance-value.md) as well: it decides whether a maintained surface earns its recurring cost,
while this principle bounds the task's scope and stopping point.
