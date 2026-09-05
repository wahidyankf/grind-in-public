---
tldr: "Sets the language target, naming, indentation, and import style for source code."
when_to_use: "Use when writing or reviewing code in any project, or when adding a language to the workspace."
---

# Code Style Policy

## Scope

This policy covers how source code is written: the language target, naming, indentation, imports, and language-native
formatting. Prettier is the source of truth for the file types it supports, and the
[Markdown style policy](../conventions/markdown-style-policy.md) covers prose. What a comment must explain belongs to
the [code commentary policy](code-commentary-policy.md).

## Language Target

Use strict TypeScript with CommonJS-compatible Node output. A project that cannot reach that target records the
deviation and its reason under [testing tooling](testing-policy/tooling.md), which is the one register for them; a
deviation that is not recorded there is a defect rather than an exception. Badak Mini is the Go CLI for repository-local
validation; follow the [Badak Mini policy](badakmini-cli-policy.md) before extending it.

## Naming and Layout

TypeScript uses two-space indentation, `camelCase` variables and functions, `PascalCase` classes, and descriptive
lower-hyphenated file names.

Go uses `gofumpt` and `goimports` as enforced by the project lint target. Follow Go's idiomatic `MixedCaps` naming, keep
package names short and lowercase, and use lowercase file names with underscores only when they improve grouping. Do not
hand-format code in a style the configured formatters will replace.

A descriptive file name is one a reader can act on without opening the file. `parse-project-file.ts` says what it holds;
`utils.ts` says only that someone had nowhere to put it.

## Imports

Import internal TypeScript libraries by package name, not by relative cross-project paths. A relative path across a
project boundary compiles, so nothing stops it, and it hides the dependency from Nx — which then cannot tell that the
importing project is affected when the imported one changes.

Let `goimports` group and order Go imports. The [Badak Mini policy](badakmini-cli-policy.md) owns its production
dependency boundary; tool dependencies in its Go module do not authorize runtime imports.

## Verification

```sh
rtk ./hippo run --class ephemeral --disk-path . -- npm run lint
rtk ./hippo run --class ephemeral --disk-path . -- npm run typecheck
```

Linting and type checking catch the mechanical part. Naming is reviewed by a person, because a name is only wrong
relative to what the code does.

## Applying a Linter's Fixes

Review an autofix before keeping it, and never run one a tool labels unsafe across a tree unattended: a fix that removes
code a rule believes unreachable leaves every gate green over the loss, because the tests for what it deleted go with
it.

Fix the construct the rule named. An identical expression on the adjacent line is a separate decision and needs its own
reason — a rule reporting one attribute has said nothing about the one beside it, and the two may differ in whether
anything observable depends on them.
