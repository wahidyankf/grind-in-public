---
tldr: "Indexes the adopted language-neutral quality standards, stack standards, and finding ratings."
when_to_use: "Use when writing, reviewing, or auditing project code under a stack pack or a language-neutral standard."
---

# Quality Standards

These standards were adopted with the stack packs. Each links the local policy that owns a shared concept here, such as
[quality gates](../quality-gates.md) or [TDD](../tdd-policy.md), rather than restating it.

- [`code/`](code/README.md) — type and boundary safety, and shell script mechanics. Use it when choosing a checker,
  adding a type escape, validating external input, or writing a script.
- [`testing/`](testing/README.md) — what a coverage number may measure. Use it when configuring coverage or deciding
  whether a layer carries a numeric target.
- [`evidence/`](evidence/README.md) — how a finding's criticality and confidence are rated. Use it when a checker rates
  a finding or a fixer orders its work.
- [`stacks/`](stacks/README.md) — the adopted stack standards and this repository's adapter. Use it when working in a
  project whose inventory entry lists a stack.
