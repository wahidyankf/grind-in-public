# Shared Skills

Codex and opencode both discover these skills automatically from this directory; see the
[shared agent directory](../README.md). Discovery ignores this README; the
[agent harness support policy](../../repo-governance/conventions/agent-harness-support.md) records that behaviour.

## Available Skills

- [`grill-me/`](grill-me/SKILL.md) — resolves open decisions by asking the owner structured multiple-choice questions.
  Use it when work is blocked on a decision, or when the owner asks to be grilled on a design.
- [`plan-grooming-idea-briefs/`](plan-grooming-idea-briefs/SKILL.md) — judges whether a brief is worth promoting,
  keeping with a trigger, or retiring.
- [`plan-creating-project-plans/`](plan-creating-project-plans/SKILL.md) — writes the six documents so each answers its
  own question and the checklist is executable.
- [`plan-writing-gherkin-criteria/`](plan-writing-gherkin-criteria/SKILL.md) — writes acceptance scenarios that describe
  observable behaviour and can actually fail.
- [`plan-validating-quality/`](plan-validating-quality/SKILL.md) — judges a draft beyond what structural validation
  reaches.
- [`plan-verifying-execution/`](plan-verifying-execution/SKILL.md) — judges finished execution against the repository
  rather than against the checklist.
- [`developing-applications/`](developing-applications/SKILL.md) — places code by what it decides or does, and handles
  errors, input, and logging, for every project the code agents touch.
- [`programming-typescript/`](programming-typescript/SKILL.md),
  [`programming-javascript/`](programming-javascript/SKILL.md), [`programming-python/`](programming-python/SKILL.md),
  [`programming-golang/`](programming-golang/SKILL.md), and [`programming-shell/`](programming-shell/SKILL.md) — apply
  each language's stack standard while writing and testing code.
- [`framework-react/`](framework-react/SKILL.md) and [`framework-nextjs/`](framework-nextjs/SKILL.md) — apply the React
  and Next.js standards; [`tooling-nx/`](tooling-nx/SKILL.md) applies the Nx standard.
- [`assessing-criticality-confidence/`](assessing-criticality-confidence/SKILL.md),
  [`applying-maker-checker-fixer/`](applying-maker-checker-fixer/SKILL.md), and
  [`generating-validation-reports/`](generating-validation-reports/SKILL.md) — rate findings, run a make, check, and fix
  loop, and write its reports.

The stack skills are chosen per project by the
[repository adapter](../../repo-governance/development/quality/stacks/repository-adapter.md).

The five `plan-` skills carry the reusable judgement behind the planning lifecycle; the workflows that sequence them are
indexed in [`repo-governance/workflows/`](../../repo-governance/workflows/README.md).

Claude receives a thin adapter in `.claude/skills/`; the complete bundle here remains canonical. See the
[harness capability parity policy](../../repo-governance/conventions/harness-capability-parity-policy.md).
