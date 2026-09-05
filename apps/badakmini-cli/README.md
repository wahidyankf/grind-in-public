# Badak Mini

Badak Mini is a deliberately small Go governance CLI. Its name means “rhinoceros” in Indonesian, and its command grammar
follows the relevant slice of [rhino-cli](https://github.com/wahidyankf/ose-public/tree/main/apps/rhino-cli) without
porting Rhino's broader repository-management surface.

## Current Command

```sh
badak-mini --help
badak-mini harness --help
badak-mini harness instruction-size validate
badak-mini harness markdown-links validate
badak-mini harness rule-change validate
badak-mini harness rule-change hook
badak-mini harness capability-parity validate
```

The command finds the Git repository root and ensures that the root agent instruction files, `AGENTS.md` and
`CLAUDE.md`, every recursive Markdown file in `repo-governance/`, and every recursive `README.md` in `.agents/`,
`.claude/`, `.codex/`, and `.opencode/` contain at most 750 words. A missing instruction file fails the check, while an
absent harness directory is skipped. Agent and command definitions are prompts rather than indexes, so they are not
measured. Its `harness` command group names the family of harness-related checks, not the files they read. It ignores
non-Markdown files and reports each violation with a progressive-disclosure remediation.

The Markdown-link command scans every Git-tracked repository Markdown file. It validates local file targets and Markdown
heading fragments, including reference-style links. It does not check external URLs. The pre-push hook always runs this
command so deleting or moving a document cannot leave a dangling local link.

The rule-change commands automatically trigger
[Rules Propagation](../../repo-governance/workflows/rules/rules-propagation.md) when a change touches a rule path, and
[Harness Alignment](../../repo-governance/workflows/harness-alignment.md) as well when that path is one a harness reads.
The `validate` form reads the staged paths during pre-commit; the `hook` form reads a harness pre-edit payload on stdin,
in either the file-path or the `apply_patch` shape, and places the triggered workflow in context. Both stay silent for
ordinary work and always exit zero, so neither can block an edit or a commit. See the
[rule change trigger policy](../../repo-governance/development/rule-change-trigger-policy.md).

The capability-parity command validates the content-level harness contract: the exact Claude instruction import,
canonical skill bundles, canonical custom-agent prompts, thin native adapters, descriptions, routes, and supported
permission mappings. Success reports dynamic harness, skill, and agent counts plus a SHA-256 digest of normalized
canonical content. Findings are stable, path-specific, read-only, and network-free. See the
[harness capability parity policy](../../repo-governance/conventions/harness-capability-parity-policy.md).

## Run and Verify

From the repository root:

```sh
npm run check:governance
npm run check:harness-parity
npm run check:markdown-links
npm run check:workflows
npm run check:go-vulnerabilities
npx nx run badakmini-cli:build
npx nx run badakmini-cli:typecheck
npx nx run badakmini-cli:lint
npx nx run badakmini-cli:test:unit
npx nx run badakmini-cli:test:integration
npx nx run badakmini-cli-e2e:test:e2e
npx nx run badakmini-cli:test:coverage:unit
npx nx run badakmini-cli:test:coverage:integration
npx nx run badakmini-cli:test:coverage:behaviour
npx nx run badakmini-cli:test:coverage
npx nx run badakmini-cli:test:quick
```

Badak Mini is pinned to Go 1.26.6. Its command tree uses exact-pinned Cobra as its only direct production library, while
its production validators remain standard-library-only; the dependency guard confines Cobra to the command adapter. Its
development module also pins GolangCI-Lint and govulncheck; run them through `go -C apps/badakmini-cli tool <name>`.
Lint starts from every applicable nondeprecated check, uses strict Go formatters, and treats every finding as blocking.
Unit coverage owns every `internal/...` validator and orchestration package; integration coverage owns only
`internal/cli`; both exact numeric gates require at least 99% statements. The thin `cmd/badak-mini` process entrypoint
is owned by the dedicated [`badakmini-cli-e2e`](../badakmini-cli-e2e/README.md) project and is outside both numeric
denominators. `tests/unit` is the Gherkin unit adapter. Badak Mini is an intentional replacement for the former shell
governance checker, not a general Rhino CLI port.

Its canonical [C4 architecture model](../../specs/apps/badakmini-cli/architecture.md) describes the current as-built
boundaries. Its [canonical behaviour corpus](../../specs/apps/badakmini-cli/README.md) is executed by unit,
local-integration, and process E2E adapters. The owner module pins Actionlint and owns `npm run check:workflows`, which
validates the scheduled GitHub Actions workflow before it is relied on.
