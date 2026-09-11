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
- [governance-structure.mjs](governance-structure.mjs) and its [tests](governance-structure.test.mjs) validate recursive
  directory indexes and governance routing frontmatter deterministically.
- [check-governance-structure.mjs](check-governance-structure.mjs) is the filesystem adapter for that structural check.
- [workflow-contract.mjs](workflow-contract.mjs) and its [tests](workflow-contract.test.mjs) validate stable terminal,
  authorization, convergence, and TDD-evidence tokens without judging semantic quality.
- [check-workflow-contract.mjs](check-workflow-contract.mjs) loads the canonical workflow documents and reports sorted
  contract findings.
- [rule-change.mjs](rule-change.mjs) and its [tests](rule-change.test.mjs) select the staged or about-to-be-edited paths
  that carry rules, and compose the notice naming the workflows a change to them starts.
- [check-rule-change.mjs](check-rule-change.mjs) is the entrypoint both the pre-commit hook and every harness pre-edit
  adapter run. Its [tests](check-rule-change.test.mjs) drive it as a process, because the property that matters most —
  that it never blocks — is an exit code rather than a return value.
- [check-repo.sh](check-repo.sh) is what `npm run test:repo` runs: every deterministic repository mechanism in one
  ordered pass, cheapest first.
- [public-safety/](public-safety/README.md) is the publication screen this repository runs first on every gate surface.
  It is not repository-local automation like the rest of this directory: the same layer, byte for byte, runs in the
  other repositories that publish, so a finding here means the same thing it means there.
