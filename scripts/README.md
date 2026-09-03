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
