# Phase 1 Toolchain

Resolved 2026-09-01. All four installed and resolved; none failed, so the Phase 3 fallback has nothing to weigh against a component that could not be obtained.

| Package | Pinned | Resolved | Notes |
| --- | --- | --- | --- |
| `typescript` | `6.0.3` | 6.0.3 | The latest stable of the 6 line, which [tooling](../../../../repo-governance/development/testing-policy/tooling.md) names. `7.0.2` is the current release and is deliberately not taken. `ose-public` pins `5.8.3`, so the port raises the compiler two majors — the exposure the owner's "fall back only on the component that breaks" decision covers. |
| `@biomejs/biome` | `2.5.11` | 2.5.11 | Linter only. `formatter.enabled` and `assist.enabled` are both `false` in the root `biome.json`, so Prettier stays the formatting source of truth. |
| `tsx` | `4.23.13` | 4.23.13 | Not `ose-public`'s `4.21.0`. That version installs cleanly but depends on `esbuild@0.27.7`, which carries GHSA-g7r4-m6w7-qqqr and fails the Phase 1 gate's `npm audit --audit-level=low`. `4.23.13` resolves `esbuild@0.28.2` and audits clean. |
| `eslint` | `9.39.4` | 9.39.4 | The version `ose-public` pins in two of its app manifests. This repository had no ESLint before this phase. |
| `eslint-plugin-jsdoc` | `64.3.2` | 64.3.2 | No manifest in either repository pinned it, so the resolved version is recorded rather than guessed. |

All four are exact pins in the root `package.json` `devDependencies`; none carries a caret or tilde.

## How These Versions Were Read

Not with `npx <tool> --version`. The executing harness proxies shell commands, and its `tsc` wrapper answered `npx tsc --version` with `TypeScript: No errors found` — a summary of a compile that never ran, with the version request dropped. Each version above was read from the installed package instead, with `node -p "require('<pkg>/package.json').version"`.

`eslint-plugin-jsdoc` needs a further step: it does not expose `./package.json` through its `exports` map, so that form throws `ERR_PACKAGE_PATH_NOT_EXPORTED`. Its version was read with `node -p "JSON.parse(require('fs').readFileSync('node_modules/eslint-plugin-jsdoc/package.json')).version"`.

[`learnings.md`](../learnings.md) carries the general lesson.
