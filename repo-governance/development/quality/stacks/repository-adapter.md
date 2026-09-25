---
tldr:
  "Records which stack packs this repository adopted, the decisions their standards leave open, and every deviation."
when_to_use:
  "Use when working in a project here, adopting or retiring a stack pack, or changing a recorded stack decision."
---

# Repository Adapter

This document owns the `extensions.software-development` inventory in [`repo-config.yml`](../../../../repo-config.yml),
as [Stack Packs](../../../conventions/structure/stack-packs.md) defines. A local policy that is stricter than an adopted
standard wins, and is recorded below as a deviation.

## Adopted Packs

| Pack         | Status  | Reason                              |
| ------------ | ------- | ----------------------------------- |
| `typescript` | adopted | as written                          |
| `javascript` | adopted | as written                          |
| `react`      | adopted | as written                          |
| `nextjs`     | adopted | as written                          |
| `python`     | adapted | one recorded pilot; see deviations  |
| `golang`     | adapted | no first-party Go source; see below |
| `shell`      | adopted | as written                          |
| `nx`         | adapted | no module-boundary lint; see below  |

## Adopter Decisions

| Source                          | Decision            | Choice                       | Reason            |
| ------------------------------- | ------------------- | ---------------------------- | ----------------- |
| `swe-code-maker`                | stack skill         | read on demand               | no edit per stack |
| `swe-code-checker`              | stack rules         | each listed stack's standard | shared choices    |
| coverage                        | floor               | 99% unit and integration     | local; see below  |
| task runner                     | runner              | Nx as a raw command runner   | local policy      |
| `nx-standards.md`               | tag vocabulary      | none yet                     | gap               |
| `javascript-standards.md`       | check scope         | none enforced yet            | gap               |
| `javascript-standards.md`       | test runner         | built-in `node --test`       | no dependency     |
| `python-standards.md`           | boundary validation | Pydantic through FastAPI     | already in use    |
| `python-standards.md`           | expected failures   | exceptions                   | FastAPI maps them |
| `python-standards.md`           | type checker        | Pyright in strict mode       | already in use    |
| `golang-standards.md`           | both decisions      | not applicable               | no Go tests       |
| `react-standards.md`            | store, query, forms | not applicable               | none in use       |
| `nextjs-standards.md`           | version line        | Next.js 16                   | current major     |
| `nextjs-standards.md`           | hosting             | managed, on Vercel           | already deployed  |
| `shell-scripts.md`              | interpreter         | Bash, `set -euo pipefail`    | the default       |
| `shell-standards.md`            | test tool           | Bash test scripts            | no dependency     |
| `generating-validation-reports` | report location     | `generated-reports/`         | already ignored   |

The Python pilot pins Pyright in `apps/forum-be-python/pyproject.toml` and sets strict mode in its `pyrightconfig.json`.
Report timestamps carry an explicit offset. Promotion and hosting follow the
[deployment policy](../../deployment-policy.md).

## Deviations

Each stronger local rule wins over the adopted standard it touches:

- [Quality gates](../../quality-gates.md) keep a 99% floor on unit and integration coverage, measured separately even
  over shared code, and forbid network in integration tests, loopback included, so a real listener is end-to-end.
- The [BDD policy](../../behaviour-driven-development-policy.md) requires a scenario corpus for every application except
  the recorded pilot.
- Python is confined to the pilot recorded in [testing tooling](../../testing-policy/tooling.md), which has no 99%
  floor.
- `tools/` pins executables only, and [code style](../../code-style-policy.md) excludes runtime Go.
- The module-boundary lint rule is an `@nx` package the [Nx workspace policy](../../nx-workspace-policy.md) excludes.
- `scripts/public-safety/` is a vendored copy kept byte-identical, so no stack rule edits it.
- Drills stay owner-solved under the [drill practice policy](../../../conventions/drill-practice-policy.md): no agent
  writes a drill's solution.

Known gaps, left for a separate plan rather than fixed by adoption: TypeScript lint runs Biome without type-aware rules;
no JavaScript file is type-checked through JSDoc; the Python pilot has no formatter, linter, or coverage gate; project
`tags` are empty; the two `.github/scripts/` tests run under `sh`.

## Project Applicability

- [wahidyankf-www](../../../../apps/wahidyankf-www/README.md)
- [wahidyankf-www-e2e](../../../../apps/wahidyankf-www-e2e/README.md)
- [forum-be-python](../../../../apps/forum-be-python/README.md)
- [repo-scripts](../../../../scripts/README.md), which also registers the vendored `public-safety/` screen
- [opencode-plugin](../../../../.opencode/plugin/README.md)
- [ci-scripts](../../../../.github/scripts/README.md)
- claude-hooks has no README: `.claude/hooks/require-hippo-boundary.sh` is the Claude pre-edit hook, with a Bash test
  beside it that no gate runs yet.
- go-tools has no README: `tools/go.mod` pins actionlint and govulncheck through `tool` directives, run by
  `npm run check:workflows` and `npm run check:go-vulnerabilities`; it has no tests, and no coverage applies.
- workspace: `nx.json` and `biome.json` are declarative configuration, outside numeric coverage.

## Version Sources

- `typescript`, `javascript`, `react`, `nextjs`, `nx`: `package.json`, `package-lock.json`, `apps/*/package.json`,
  `.nvmrc`
- `python`: `apps/forum-be-python/pyproject.toml`, `apps/forum-be-python/uv.lock`,
  `apps/forum-be-python/.python-version`
- `golang`: `tools/go.mod`, `tools/go.sum`
