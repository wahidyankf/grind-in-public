# Repository Guidelines

Personal learning workspace. Write clear, simple English that non-native speakers can follow, including replies to the
user. Keep changes small.

## Layout

- `apps/` runnable applications: `wahidyankf-www` (Next.js site) and `wahidyankf-www-e2e` (Playwright suite).
- `libs/` reusable packages. None yet.
- `docs/` human-facing documentation. Only the index READMEs are kept.
- `specs/` Gherkin scenarios and C4 architecture for `wahidyankf-www`.
- `scripts/` repository-local automation. `scripts/check-repo.sh` runs the deterministic repository checks.
- `.husky/` git hooks: pre-commit, commit-msg, and pre-push.
- `repo-config.yml` declares the projects and the repository gates read by RHINO.
- `exercises/` and `local-tmp/` are personal, untracked work areas. Do not edit or commit them.

## Commands

Run from the repository root with `npm run <script>`:

- `npm run test:repo`: deterministic repository checks, including the RHINO `main` gate.
- `npm run test:quick`: quick tests for every affected project.
- `npm run check:markdown-links`, `npm run check:markdown-line-length`: Markdown checks.
- `npm run validate:config`: validates `repo-config.yml`.
- `npm run format` and `npm run format:check`: Prettier.
- `npm run lint`, `npm run typecheck`, `npm run build`: Nx tasks for each project.

Nx is used only as a raw task runner. Do not add Nx plugins.

## Conventions

- Prettier is the formatter. Markdown lines stay within 120 columns.
- Commit messages follow Conventional Commits. `commitlint` checks them.
- Commit and push only when the owner asks. Landing on `main` is not deploying.
- Keep test data synthetic and isolated.
- Behaviour work starts from a Gherkin scenario in `specs/`, written test-first.
