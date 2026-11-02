# Repository Guidelines

Personal learning workspace. Write clear, simple English that non-native speakers can follow, including replies to the
user. Keep changes small.

## Layout

- `apps/` runnable applications: the `edutech-*` and `knowledge-base-*` apps, see Projects and targets.
- `compose.edutech.yaml` and `compose.knowledge-base.yaml` run the local Postgres for each backend. Each file is its own
  Compose project, so resetting one database never touches the other.
- Each app keeps its manual HTTP requests in `api.http`, next to its code, so the requests can be updated with the
  routes they call.
- `libs/` reusable packages. None yet.
- `drill/` hands-on practice code. It is standalone, not an Nx project, and has its own `pyproject.toml`.
- `docs/` human-facing documentation. Only the index READMEs are kept.
- `specs/` Gherkin scenarios and C4 architecture for applications that carry specifications. None do at present.
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

## Toolchain and dependencies

- Volta pins Node and npm in the root `package.json` (`volta.node`, `volta.npm`). `engines.node` sets the allowed range.
  There is no `.nvmrc`. CI reads the same pin.
- One root `package-lock.json` serves every app. Run `npm install` from the repository root.
- Use caret ranges (`^`) in `package.json`, as in the `super-*` reference repository. A package used by more than one
  app uses the same range.
- Pick a release that is more than 7 days old, has no known advisory, and is the latest LTS when an LTS exists. Check
  with `npm view <package> time` and `npm audit`. Never run `npm audit fix --force`.
- Keep one copy of each shared package at the root. After a change, run `npm ls <package>` and look for stale copies
  under `apps/*/node_modules`. A second React copy causes "Invalid hook call".
- Biome is the only linter. The root `biome.json` is shared, and each app has a nested `biome.json` with `"root": false`
  and `"extends": "//"`. There is no ESLint, so no tool checks JSDoc comments.
- Each JS app has its own self-contained `tsconfig.json`. No app extends the root `tsconfig.base.json`.

Versions held back on purpose. Check the reason before you bump them:

- `@types/node` stays on 24 to match the Node 24 LTS runtime.
- `npm audit` reports findings in root tooling: `markdownlint-cli2` and its chain, `nx`, and `undici`. The offered fixes
  are downgrades or major upgrades, so they are not applied without a separate review.
- `node-pg-migrate` is not used. Its `glob` dependency has an open advisory, so `knowledge-base-be` runs its own small
  migration runner (`src/migrate.ts`).

## Projects and targets

Run one target with `npx nx <target> <project>`. Every app is an Nx project with its own `project.json`.

- `edutech-be` (`apps/edutech-be`, FastAPI on Python 3.13). Run `uv sync` in `apps/edutech-be` before its targets.
  - `dev` and `start`: run uvicorn on port 8000. `dev` reloads on change.
  - `lint`: ruff. `typecheck`: pyright. `test:unit`: pytest, without the `integration` tests.
  - `test:integration`: pytest for the Postgres tests. Run `db:up` first.
  - `db:up` and `db:down`: start and stop Postgres on port 5432, from `compose.edutech.yaml`. `db:migrate` applies
    Alembic migrations. `db:reset` deletes all data, then starts and migrates again.
  - `test:quick`: typecheck, lint, then unit tests. It does not need Postgres.
- `edutech-web` (`apps/edutech-web`, Next.js):
  - `dev` and `start`: run the site on port 3000. `start` runs after `build`.
  - `build`: make a production build.
  - `typecheck`, `lint` (Biome, which skips `public/`), and `test:unit` (Vitest).
  - `test:quick`: typecheck, lint, then unit tests.
- `edutech-web-e2e` (`apps/edutech-web-e2e`, Playwright):
  - `test:e2e`: run the browser tests. Its Playwright config builds and starts `edutech-web` on port 3000, then waits
    for `/api/health`.
  - `lint`, `typecheck`, and `test:quick`.
- `knowledge-base-be` (`apps/knowledge-base-be`, Fastify on Node 24, TypeScript). Its REST API lives under `/api`.
  - `dev` runs `tsx watch` on port 8200. `build` compiles to `dist/`, and `start` runs `dist/main.js` after `build`.
  - `lint` (Biome), `format` (Prettier), `typecheck` (`tsc`), and `test:unit` (Vitest, with the in-memory store).
  - `test:integration`: Vitest against Postgres. Run `db:up` and `db:migrate` first.
  - `db:up` and `db:down`: start and stop Postgres on port 5232, from `compose.knowledge-base.yaml`. `db:migrate`
    applies `migrations/*.sql` once each, in order. `db:reset` deletes all data, then starts and migrates again.
  - `test:quick`: typecheck, lint, then unit tests. It does not need Postgres.
- `knowledge-base-web` (`apps/knowledge-base-web`, TanStack Start in SPA mode, so pages render in the browser only, with
  TanStack Router and TanStack Query):
  - `dev` and `start` (Vite preview) run the site on port 3200. Both forward `/api` to the backend on port 8200.
  - `build`, `typecheck`, `lint` (Biome), `test:unit` (Vitest with jsdom), and `test:quick`.
- `knowledge-base-web-e2e` (`apps/knowledge-base-web-e2e`, Playwright):
  - `test:e2e`: run the browser tests. It depends on `knowledge-base-be:db:migrate`. Its Playwright config builds and
    starts the backend and the web app, then waits for `/api/health` on each.
  - `lint`, `typecheck`, and `test:quick`.
- The `edutech-*` and `knowledge-base-*` apps are standalone. They use no Gherkin corpus. Their required targets are not
  checked by any script, so review each `project.json` by hand when they change.

## Conventions

- Prettier is the formatter. Markdown lines stay within 120 columns.
- Commit messages follow Conventional Commits. `commitlint` checks them.
- Commit and push only when the owner asks. Landing on `main` is not deploying.
- Keep test data synthetic and isolated.
- Remove the test containers, volumes, and networks this work leaked only by exact name, found through their compose
  label `com.docker.compose.project.working_dir` under this checkout, and never with a blanket prune.
- Diagrams are Mermaid with `accTitle` and `accDescr`. Keep labels short and plain. Colours come only from the
  accessible palette in `repo-config.yml`, and each flowchart declares `classDef default`. Decisions are orange, start
  nodes blue, and end nodes teal. ASCII is kept only for file trees.
- Behaviour work starts from a Gherkin scenario in `specs/`, written test-first.
- Keep comments, commit messages, and notes brief. Explain why the code is needed, not what it does.

## Architecture

### Config and secrets

- Build config once from `process.env` at import time. Export the build function too, so tests can call it.
- In production, a missing or invalid required security variable makes the app fail at startup.
- Do not add `.env.example` or `.env.template` files.

### Tests

- The `edutech-*` apps put unit tests next to the source as `*_tests.py` (pytest) and `*_tests.ts` or `*_tests.tsx`
  (Vitest). Their production-only checks use `*_production_tests.*`, which the `test:unit` targets skip.
- The `knowledge-base-*` apps follow the same rule. Postgres tests are named `*_integration_tests.ts`, and only the
  `test:integration` target runs them.
