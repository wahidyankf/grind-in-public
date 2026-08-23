---
tldr: "Sets TypeScript and Go quality tooling and dependency requirements."
when_to_use: "Use when configuring project tooling, linters, or dependency versions."
---

# Testing Tooling

TypeScript projects use TypeScript 6's `tsc --noEmit`, Biome, and project-local ESLint commentary checks. Go projects use `go vet` plus an exact-pinned GolangCI-Lint v2 tool configuration. That configuration starts from every linter, enables strict formatting, treats every finding as blocking, and disables only deprecated, duplicative, conflicting, or inapplicable checks with a nearby reason. A linter suppression must be narrow, specific, and explained; broad exclusions and issue caps are forbidden. Keep dependency versions exact; audit npm dependencies and scan Go module dependencies with the commands [workspace commands](../workspace-commands.md#repository-checks) lists.
