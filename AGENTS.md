# Repository Guidelines

Personal learning workspace. Write clear, simple English that non-native speakers can follow, including replies to the
user. Keep changes small.

## Layout

- `apps/` runnable applications: `wahidyankf-www` (Next.js site) and `wahidyankf-www-e2e` (Playwright suite). The
  `edutech-*` apps are standalone, see Projects and targets.
- `infra/` local infrastructure. `infra/edutech/compose.yaml` runs the local Postgres for `edutech-be`.
- Each app keeps its manual HTTP requests in `api.http`, next to its code, so the requests can be updated with the
  routes they call.
- `libs/` reusable packages. None yet.
- `docs/` human-facing documentation. Only the index READMEs and `reference/self-hosting.md` are kept.
- `specs/` Gherkin scenarios and C4 architecture for `wahidyankf-www`.
- `scripts/` repository-local automation. `scripts/check-repo.sh` runs the deterministic repository checks.
- `.husky/` git hooks: pre-commit (lint-staged), commit-msg (commitlint), and pre-push (Markdown checks, then Nx).
- `repo-config.yml` holds the Markdown policy that RHINO reads for its link and Mermaid checks. Nothing else reads it.
- `exercises/` and `local-tmp/` are personal, untracked work areas. Do not edit or commit them.

## Commands

Run from the repository root with `npm run <script>`:

- `npm run test:repo`: deterministic repository checks, including the Markdown checks and the spelling check.
- `npm run test:quick`: quick tests for every project. It does not select only affected projects.
- `npm run check:markdown-links`, `npm run check:markdown-line-length`: Markdown checks.
- `npm run validate:config`: validates the Markdown policy in `repo-config.yml`.
- `npm run format` and `npm run format:check`: Prettier.
- `npm run lint`, `npm run typecheck`, `npm run build`: Nx tasks for each project.

Nx is used only as a raw task runner. Do not add Nx plugins.

The `pre-push` hook runs the Markdown checks, then `nx affected -t test:quick` against `origin/main`. It runs
`npm run test:repo` only when the pushed range changes `specs`, `scripts`, `.github`, `.husky`, `package.json`,
`nx.json`, `repo-config.yml`, `rhino`, `rhino.lock`, or an app's `project.json`. It needs a local `origin/main`.

## Projects and targets

Run one target with `npx nx <target> <project>`. Every app is an Nx project with its own `project.json`.

- `wahidyankf-www` (`apps/wahidyankf-www`, the Next.js site):
  - `dev` and `start`: run the site on port 3201, or on `WAHIDYANKF_WWW_PORT` when it is set.
  - `build`: make a production build.
  - `typecheck` and `lint`: `tsc`, Biome, and ESLint.
  - `test:unit` and `test:integration`: Vitest. Integration tests use the real filesystem.
  - `test:coverage:unit`, `test:coverage:integration`, `test:coverage:behaviour`: coverage runs. The unit and
    integration runs enforce a 99% line floor.
  - `static-routes:validation`: build the site, then check its static routes.
  - `test:quick`: runs `static-routes:validation`, then typecheck, lint, unit tests, and coverage checks.
- `wahidyankf-www-e2e` (`apps/wahidyankf-www-e2e`, Playwright):
  - `install`: install Chromium for Playwright.
  - `test:e2e`: run the browser tests against the production build.
  - `test:coverage:behaviour:e2e`: check that the Gherkin scenarios match the browser tests.
  - `typecheck`, `lint`, and `test:quick`.
- `edutech-be` (`apps/edutech-be`, FastAPI on Python 3.13). Run `uv sync` in `apps/edutech-be` before its targets.
  - `dev` and `start`: run uvicorn on port 8100. `dev` reloads on change.
  - `lint`: ruff. `typecheck`: pyright. `test:unit`: pytest, without the `integration` tests.
  - `test:integration`: pytest for the Postgres tests. Run `db:up` first.
  - `db:up` and `db:down`: start and stop Postgres on port 5132, from `infra/edutech/compose.yaml`. `db:migrate` applies
    Alembic migrations. `db:reset` deletes all data, then starts and migrates again.
  - `test:quick`: typecheck, lint, then unit tests. It does not need Postgres.
- `edutech-web` (`apps/edutech-web`, Next.js):
  - `dev` and `start`: run the site on port 3100. `start` runs after `build`.
  - `build`: make a production build.
  - `typecheck`, `lint` (Biome, which skips `public/`), and `test:unit` (Vitest).
  - `test:quick`: typecheck, lint, then unit tests.
- `edutech-web-e2e` (`apps/edutech-web-e2e`, Playwright):
  - `test:e2e`: run the browser tests. Its Playwright config builds and starts `edutech-web` on port 3100, then waits
    for `/api/health`.
  - `lint`, `typecheck`, and `test:quick`.
- The `edutech-*` apps are standalone. They do not use the Gherkin corpus or the coverage targets above. Their required
  targets are not checked by any script, so review each `project.json` by hand when they change.

## Conventions

- Prettier is the formatter. Markdown lines stay within 120 columns.
- Commit messages follow Conventional Commits. `commitlint` checks them.
- Commit and push only when the owner asks. Landing on `main` is not deploying.
- Keep test data synthetic and isolated.
- Behaviour work starts from a Gherkin scenario in `specs/`, written test-first.
- Keep comments, commit messages, and notes brief. Explain why the code is needed, not what it does.

## Architecture

### Config and secrets

- Build config once from `process.env` at import time. Export the build function too, so tests can call it.
- In production, a missing or invalid required security variable makes the app fail at startup.
- `apps/wahidyankf-www/src/env.ts` is the validation seam. It declares no variables today.
- Read `docs/reference/self-hosting.md` before you change config or secrets. Do not add `.env.example` or
  `.env.template` files. The guide has no short summary on purpose, because the security stakes are high.

### Tests

- `wahidyankf-www` unit tests sit next to the source as `*.unit.test.ts` or `*.unit.test.tsx`. Vitest runs them in the
  `unit` project.
- `wahidyankf-www` config assertions that apply only in production mode go in `*.production.test.ts` files. Add a Vitest
  project for them before you write the first one.
- `wahidyankf-www` E2E tests use Playwright against a real running server. Playwright starts the server itself
  (`webServer` in `apps/wahidyankf-www-e2e/playwright.config.ts`) and waits for its URL before the tests run.
- The `edutech-*` apps put unit tests next to the source as `*_tests.py` (pytest) and `*_tests.ts` or `*_tests.tsx`
  (Vitest). Their production-only checks use `*_production_tests.*`, which the `test:unit` targets skip.
