# Scripts

This directory is reserved for small, repository-local automation scripts that do not belong to an Nx project or a Git
hook.

Keep scripts focused, portable, and well commented. Prefer adding repeatable development tasks as Nx `command` targets;
place hook orchestration in `.husky/`.

Before adding a script, check whether an existing command or a declared validator already provides the needed repository
validation.

## Directory Map

- [check-repo.sh](check-repo.sh) is what `npm run test:repo` runs: every deterministic repository mechanism in one
  ordered pass, cheapest first.
