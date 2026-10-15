# Scripts

This directory is reserved for small, repository-local automation scripts that do not belong to an Nx project or a Git
hook.

Keep scripts focused, portable, and well commented. Prefer adding repeatable development tasks as Nx `command` targets;
place hook orchestration in `.husky/`.

Before adding a script, check whether an existing command or a declared validator already provides the needed repository
validation.

## Directory Map

- [next-with-port.mjs](next-with-port.mjs) — resolves a Next.js app's listening port before starting it, so `--port`,
  the app's prefixed environment variable, and its compiled-in default rank in that order rather than in whichever order
  Next's CLI happens to apply.
- [project-contract.mjs](project-contract.mjs) and its [tests](project-contract.test.mjs) define the deterministic
  owner/E2E contract for every declared project pair, without filesystem or subprocess access.
- [check-project-contract.mjs](check-project-contract.mjs) is the production adapter that loads those descriptors and
  reports sorted contract findings.
- [check-repo.sh](check-repo.sh) is what `npm run test:repo` runs: every deterministic repository mechanism in one
  ordered pass, cheapest first.
