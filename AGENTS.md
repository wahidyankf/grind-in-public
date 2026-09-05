# Repository Guidelines

## Purpose and Working Style

Grind in Public is a personal lifelong-learning workspace for software engineering. Follow the
[drill practice](repo-governance/conventions/drill-practice-policy.md) and
[task tracking](repo-governance/conventions/task-tracking-policy.md) policies; the second also states how to treat
another task's concurrent edits. Write everything that lands in a file in English; see the
[language policy](repo-governance/conventions/language-policy.md). Use the smallest responsible task change; maintained
surfaces must prove net recurring value. See [minimum sufficiency](repo-governance/principles/minimum-sufficiency.md)
and [maintenance value](repo-governance/principles/maintenance-value.md). Resolve an open decision by grilling with
options, not prose; see the [grilling-with-options policy](repo-governance/conventions/grilling-with-options-policy.md).

## Reference Repositories

Use [ose-public](https://github.com/wahidyankf/ose-public), [ose-primer](https://github.com/wahidyankf/ose-primer), and
[beaver-nest](https://github.com/wahidyankf/beaver-nest) as read-only references; local rules govern. For CV work, read
[apps/wahidyankf-www/docs/README.md](apps/wahidyankf-www/docs/README.md).

## Rule Changes and Audience

Rule-path automation triggers the [`rules-propagation` workflow](repo-governance/workflows/rules/rules-propagation.md)
before edits and at pre-commit; follow it without waiting for an owner reminder. A change to what a harness reads also
requires the [`harness-alignment` workflow](repo-governance/workflows/harness-alignment.md). Run Plan Quality Gate and
Rules Quality Gate only when the owner explicitly requests the semantic checkpoint; neither planning, execution, review,
nor propagation grants that authority. Governance is ordered — principles, then conventions, then development policies,
then workflows — so a cross-level conflict is settled by [precedence](repo-governance/principles/README.md#precedence)
and only a same-level one reaches the owner. What counts as a rule, and how strongly each wording binds, is the
[rule definition policy](repo-governance/conventions/rule-definition-policy.md). `README.md` and `docs/` serve people;
instruction files serve agents; `repo-governance/` serves both. `CLAUDE.md` must defer to this file; see the
[agent instruction alignment policy](repo-governance/conventions/agent-instruction-alignment-policy.md). Every `docs/`,
`repo-governance/`, `scripts/`, `plans/`, `specs/`, and harness directory requires an indexed README; see the
[documentation index policy](repo-governance/documentation-index-policy.md).

## Project Structure

- `apps/` runnable applications; `libs/` reusable packages.
- Every Nx project maintains its [project README](repo-governance/conventions/project-readme-policy.md).
- `docs/` human-facing Diátaxis documentation.
- `repo-governance/` shared policies and workflows; `plans/` delivery plans; `specs/` as-built application architecture
  and Gherkin behaviour.
- Root configs: `package.json`, `nx.json`, `tsconfig.base.json`.

Keep implementation and tests under `src/`; use lowercase-hyphenated project directories. Add assets only within the
project needing them.

## Commands

[Workspace commands](repo-governance/development/workspace-commands.md) is canonical for every command, check, and hook:
the common loop, the narrower runs, and what each Git hook does. Guard compute-bearing Nx work through the
checksum-pinned `./hippo` consumer under the
[resource-aware development policy](repo-governance/development/resource-aware-development.md); never bypass or
parallel-retry an admission deferral.

## Planning

Create plan documents only when the owner explicitly requests a plan. A requested plan has five core documents,
including `tech-docs/README.md`, promoted from a quadrant idea through `backlog/`, `in-progress/`, and `done/`; see the
[plans organization policy](repo-governance/conventions/plans-organization-policy.md). Author with
[plan-planning](repo-governance/workflows/plan-planning.md), validate with
[plan-quality-gate](repo-governance/workflows/plan-quality-gate.md) only at an explicitly requested checkpoint, and run
with [plan-execution](repo-governance/workflows/plan-execution.md). Plans deliver directly to `main`. Each `delivery.md`
opens with a dated Execution Record, written as phases complete and gates pass or fail.

## Nx and Coding Conventions

Prefer the standard library and mechanisms this repository already owns; the
[dependency selection policy](repo-governance/development/dependency-selection-policy.md) states when an external
dependency may be added. Use Nx only as a raw task runner with `command` targets. Add no plugin, executor, or generator
without explicit owner direction; see the [Nx workspace policy](repo-governance/development/nx-workspace-policy.md).

Prettier is the source of truth; this repository is read in a terminal, so Markdown wraps prose at 120 columns, keeps
table cells short enough to align, and draws terminal-first ASCII diagrams—see the
[Markdown style policy](repo-governance/conventions/markdown-style-policy.md). The language target, naming, indentation,
and import style follow the [code style policy](repo-governance/development/code-style-policy.md).

Comments must explain intent, flow, and non-obvious decisions without narrating syntax; see the
[code commentary policy](repo-governance/development/code-commentary-policy.md).

## Testing and Commits

Each Nx project follows the [testing policy](repo-governance/development/testing-policy.md) and the compulsory
[BDD policy](repo-governance/development/behaviour-driven-development-policy.md). Each non-drill application also
maintains its as-built C4 model in `specs/`; see the
[architecture specification policy](repo-governance/development/architecture-specifications.md). Behaviour is specified
as Gherkin in `specs/` and implemented test-first, one scenario per red-green-refactor cycle; see the
[specs policy](repo-governance/development/specs-policy.md) and the
[TDD policy](repo-governance/development/tdd-policy.md).

Represent every behaviour increment and bug fix as separate evidenced RED, GREEN, and REFACTOR task items. Use synthetic
isolated test data, keep integration network-free and E2E outside quick hooks, and run the one-by-one Gherkin
implementation review after material corpus or adapter changes. Manually confirm affected browser and API boundaries.

Use Conventional Commits, and give each commit one purpose under the
[thematic commits policy](repo-governance/conventions/thematic-commits-policy.md). Landing on `main` is not deploying:
promoting a commit and cutting a domain over are separate authorized acts under the
[deployment policy](repo-governance/development/deployment-policy.md). Commit and push are separate permissions, and
neither is implied by the work that produced the changes: the
[commit hook policy](repo-governance/development/commit-hook-policy.md) owns authorization, public-repository safety,
attribution, and hook requirements. Do not commit `node_modules/` or unreviewed dependency updates.
