---
tldr: "Sets TypeScript and Go quality tooling and dependency requirements."
when_to_use: "Use when configuring project tooling, linters, or dependency versions."
---

# Testing Tooling

TypeScript projects use TypeScript 6's `tsc --noEmit`, Biome, and project-local ESLint commentary checks. Go projects
use `go vet` plus an exact-pinned GolangCI-Lint v2 tool configuration. That configuration starts from every linter,
enables strict formatting, treats every finding as blocking, and disables only deprecated, duplicative, conflicting, or
inapplicable checks with a nearby reason. A linter suppression must be narrow, specific, and explained; broad exclusions
and issue caps are forbidden. Keep dependency versions exact; audit npm dependencies and scan Go module dependencies
with the commands [workspace commands](../workspace-commands.md#repository-checks) lists.

## Configuration Semantics

When a configuration format has once replaced a default set rather than extending it, treat that as the format's rule
and not a fact about the one key that surprised you: Nx `inputs` arrays and Vitest `projects` entries both replace, so a
project inherits neither the root's plugins nor the rest of its `test` options.

Read a JSON value into a shell variable with a command that writes the value, not one that pretty-prints it. `node -p`
renders through `util.inspect`, which colourises a number when stdout is a TTY — and a task runner supplies one — so the
value carries ANSI codes and compares unequal to an identical-looking number.

## Recorded Deviations

`apps/wahidyankf-www` sets `module` to `esnext`, `moduleResolution` to `bundler`, and `target` to `ES2017`. It overrides
`tsconfig.base.json` on those three options while `strict` stays true, and reaches that state by extending the base and
overriding on top of it. That is not the CommonJS-compatible Node output [code style](../code-style-policy.md) names as
the language target. Next 16 leaves the application no alternative: it resolves its own imports as ESM through a
bundler, and a CommonJS-compatible configuration fails to build. The browser adapter now lives inside that project and
compiles under the same settings.

Biome runs in both TypeScript projects as a linter only. The root `biome.json` sets `formatter.enabled` and
`assist.enabled` to `false`, and Prettier remains the formatting source of truth. Biome v2 defaults
`formatter.indentStyle` to tab where Prettier here uses two spaces, so an enabled Biome formatter would report every
ported file, and `biome check --write` would retab them and break `npm run format:check` across the repository.

The root `biome.json` also sets `css.parser.tailwindDirectives` to `true`. Tailwind v4 states its configuration in CSS
through at-rules the CSS grammar does not define — `@theme`, `@apply`, `@utility`, `@variant`, and their siblings — and
without that option Biome's parser rejects them, so `noUnknownAtRules` reports every one and the CSS lint carries no
signal. With it on, Biome parses them and still reports a genuine unknown at-rule, which is what makes the check worth
running.

Biome is therefore this repository's CSS at-rule check. A generic CSS language server run by an editor reads none of
this configuration and reports the same directives as unknown; that is an editor setting, not a repository defect, and
nothing here is weakened to quiet it.
