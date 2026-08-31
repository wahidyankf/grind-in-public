---
tldr: "Sets TypeScript and Go quality tooling and dependency requirements."
when_to_use: "Use when configuring project tooling, linters, or dependency versions."
---

# Testing Tooling

TypeScript projects use TypeScript 6's `tsc --noEmit`, Biome, and project-local ESLint commentary checks. Go projects use `go vet` plus an exact-pinned GolangCI-Lint v2 tool configuration. That configuration starts from every linter, enables strict formatting, treats every finding as blocking, and disables only deprecated, duplicative, conflicting, or inapplicable checks with a nearby reason. A linter suppression must be narrow, specific, and explained; broad exclusions and issue caps are forbidden. Keep dependency versions exact; audit npm dependencies and scan Go module dependencies with the commands [workspace commands](../workspace-commands.md#repository-checks) lists.

## Recorded Deviations

`apps/wahidyankf-www` sets `module` to `esnext`, `moduleResolution` to `bundler`, and `target` to `ES2017`, overriding `tsconfig.base.json` on those three options while `strict` stays true. That is not the CommonJS-compatible Node output [code style](../code-style-policy.md) names as the language target. Next 16 leaves no alternative: it resolves its own imports as ESM through a bundler, and a CommonJS-compatible configuration fails to build.

Biome runs in that project as a linter only. The root `biome.json` sets `formatter.enabled` and `assist.enabled` to `false`, and Prettier remains the formatting source of truth. Biome v2 defaults `formatter.indentStyle` to tab where Prettier here uses two spaces, so an enabled Biome formatter would report every ported file, and `biome check --write` would retab them and break `npm run format:check` across the repository.
