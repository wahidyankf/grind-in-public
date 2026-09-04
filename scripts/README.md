# Scripts

This directory is reserved for small, repository-local automation scripts that do not belong to an Nx project or a Git
hook.

Keep scripts focused, portable, and well commented. Prefer adding repeatable development tasks as Nx `command` targets;
place hook orchestration in `.husky/`.

Before adding a script, check whether [Badak Mini](../apps/badakmini-cli/README.md) already provides the needed
repository validation.

## Directory Map

- [next-with-port.mjs](next-with-port.mjs) — resolves a Next.js app's listening port before starting it, so `--port`,
  the app's prefixed environment variable, and its compiled-in default rank in that order rather than in whichever order
  Next's CLI happens to apply.
- [project-contract.mjs](project-contract.mjs) and its [tests](project-contract.test.mjs) define the deterministic
  four-project owner/E2E contract without filesystem or subprocess access.
- [check-project-contract.mjs](check-project-contract.mjs) is the production adapter that loads the four descriptors and
  reports sorted contract findings.
- [governance-structure.mjs](governance-structure.mjs) and its [tests](governance-structure.test.mjs) validate recursive
  directory indexes and governance routing frontmatter deterministically.
- [check-governance-structure.mjs](check-governance-structure.mjs) is the filesystem adapter for that structural check.
- [workflow-contract.mjs](workflow-contract.mjs) and its [tests](workflow-contract.test.mjs) validate stable terminal,
  authorization, convergence, and TDD-evidence tokens without judging semantic quality.
- [check-workflow-contract.mjs](check-workflow-contract.mjs) loads the canonical workflow documents and reports sorted
  contract findings.
