---
tldr:
  "Records how this repository adopted the shared quality-gate contract: families, subjects, checks, and deviations."
when_to_use:
  "Use when running, adopting, or changing a quality gate or propagation here, or when one names an absent owner."
---

# Quality Gate Adapter

The [Quality Gate Contract](quality-gate-contract.md), [Sole-Writer Propagation](sole-writer-propagation.md), and each
family's gate, propagation, checker, and fixer were copied from the shared `ose-rules` catalog at the same paths, so
their links resolve unchanged. A stronger local rule wins over an adopted one, and every difference is recorded here.

## Families

| Family      | Subject here                                                       |
| ----------- | ------------------------------------------------------------------ |
| `plan`      | one plan folder under `plans/`                                     |
| `docs`      | `README.md` files, `docs/`, project documents, and `specs/` prose  |
| `rules`     | `AGENTS.md`, `repo-governance/`, agents, and skills                |
| `harness`   | the Claude Code, Codex, and OpenCode bindings                      |
| `ci`        | `.github/workflows/`, `.husky/`, and the `repo-config.yml` gates   |
| `pr-review` | one local commit range on `main`, before it is pushed              |
| `specs`     | listed folders under `specs/`                                      |
| `ui-web`    | `apps/wahidyankf-www`, served locally, with its `specs/` scenarios |
| `api-http`  | `apps/forum-be-python`, served locally, with its OpenAPI document  |

Not adopted: the content, `pdf-to-md`, and tutorial families, because the repository publishes no such content.

## Adopter Decisions

- **Callers.** None. Every gate runs only on an explicit owner request, as `AGENTS.md` requires; planning, execution,
  review, and propagation never start one. Each gate's Entry says so.
- **Entry and exit check.** `npm run test:repo` for every family. The `ui-web` family adds the `test:quick` targets of
  `wahidyankf-www` and `wahidyankf-www-e2e`, the `api-http` family adds `forum-be-python`'s `test:quick` and `test:e2e`,
  and the `ci` family adds `npm run check:workflows`.
- **Deterministic Boundary.** Each gate's last column names the tools those commands run. A property no tool here owns,
  such as plan layout, most page accessibility, or the API contract, left the table and is judgeable.
- **Structure.** `scripts/check-workflow-contract.mjs`, run by `npm run test:repo`, requires the plan and rules gates'
  explicit entry, three-cycle bound, and advisory verdict, and rejects the retired results and inputs. RHINO's
  `governance quality-gates validate`, a `main` gate entry, checks each declared gate's layout, headings, verdicts, and
  cycle ceiling.
- **Review surface.** The [integration path](../../conventions/integration-path-policy.md) is direct to `main`, so a
  review pass reads a local commit range. Its pipeline is `npm run test:repo` on the pinned head. The writer commits
  each repair as a new commit on local `main` and never pushes; pushing stays a separate permission.
- **Review route.** No scout or lens checker is adopted, so every pass takes the trivial tier: `pr-review-checker`
  reviews the whole change alone and returns its report for the caller to record.
- **Running surfaces.** No interface tester agent is adopted. `ui-web-checker` drives the served site with a headless
  browser and `api-http-checker` sends real requests to the served API, each started through its `dev` target under
  `./hippo` and stopped after the audit.
- **Agent fields.** Catalog `capabilities` become `requires`, `read-only` becomes denied `repository-write` and
  `nested-agent` with `inline-result-only`, and `tier` and `skills` are dropped. No harness here declares `network`, so
  `docs-checker` and `harness-checker` return each outside-world research need to their caller.
- **Skills not adopted.** `authoring-documentation`, `propagating-rules`, `checking-harness-compatibility`,
  `applying-ci-standards`, `validating-specification-structure`, `resolving-review-threads`,
  `synthesizing-review-findings`, `producing-review-findings`, `validating-factual-accuracy`, `validating-links`,
  `validating-governance-rules`, `understanding-governance-architecture`, `exploratory-testing`,
  `design-fidelity-review`, `usability-heuristic-evaluation`, and `developing-frontend-ui`. Each executor works from its
  workflow and agent text.
- **Shape.** Adopted documents carry `tldr` and `when_to_use` front matter, as the
  [documentation index policy](../../documentation-index-policy.md) requires.

## Local Owners

An adopted document links the local owner wherever the catalog linked its own, for example precedence for governance
layers, the task tracking policy for temporary files, agent harness support for harness adapters, and the push leak
review for the pull-request leak review.

A catalog owner with no local counterpart stays named without a link: Bounded Convergence, Documentation Architecture,
Content Quality, README Quality, Factual Validation, Repository Documentation Files, Related Repositories, CI Workflow
File Naming, the review disciplines, Web Research Delegation, Deterministic and Judgement Validation, Preexisting Error
Resolution, CI Post-Push Verification, Pull Request Merge, Release Cut, Parity Planning, UX Review Fix Planning, and the
principles this tree does not hold. Each repeated operation here states its own bound.

## Local Rules Kept

- **Rules propagation** starts automatically for a rule-path change, announced by `npm run check:rule-change` and the
  harness pre-edit hooks, and keeps the local
  [Idempotency Gate](../../workflows/quality/rules-propagation/004-idempotency-gate.md) as its fourth module.
- **Harness parity verification** keeps its automatic trigger for every harness-read path and its repair step.
- **Docs propagation** starts automatically before a documentation-changing commit.
