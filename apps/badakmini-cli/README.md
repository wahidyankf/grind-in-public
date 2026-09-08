# Badak Mini

Badak Mini announces the workflows a rule change requires. Its name means “rhinoceros” in Indonesian, and its command
grammar follows the relevant slice of [rhino-cli](https://github.com/wahidyankf/ose-public/tree/main/apps/rhino-cli)
without porting Rhino's broader repository-management surface. It once also validated documentation hygiene; those three
checks now belong to [RHINO](https://github.com/wahidyankf/rhino), declared in `repo-config.yml` at the repository root
and listed under [repository checks](../../repo-governance/development/workspace-commands.md#repository-checks).

## Current Command

```sh
badak-mini --help
badak-mini harness --help
badak-mini harness rule-change validate
badak-mini harness rule-change hook
```

A name outside that list is an invalid invocation and exits `2` with nothing on standard output, so a hook still calling
a retired check fails loudly rather than reporting a clean run of a check that no longer exists.

The rule-change commands automatically trigger
[Rules Propagation](../../repo-governance/workflows/rules/rules-propagation.md) when a change touches a rule path, and
[Harness Alignment](../../repo-governance/workflows/harness-alignment.md) as well when that path is one a harness reads.
The `validate` form reads the staged paths during pre-commit; the `hook` form reads a harness pre-edit payload on stdin,
in either the file-path or the `apply_patch` shape, and places the triggered workflow in context. Both stay silent for
ordinary work and always exit zero, so neither can block an edit or a commit. See the
[rule change trigger policy](../../repo-governance/development/rule-change-trigger-policy.md).

## Run and Verify

From the repository root:

```sh
rtk ./hippo run --class ephemeral --disk-path . -- npm run check:hygiene
rtk ./hippo run --class ephemeral --disk-path . -- npm run check:workflows
rtk ./hippo run --class ephemeral --disk-path . -- npm run check:go-vulnerabilities
rtk ./hippo run --class ephemeral --disk-path . -- npm exec -- nx run -p badakmini-cli -t build
rtk ./hippo run --class ephemeral --disk-path . -- npm exec -- nx run -p badakmini-cli -t typecheck
rtk ./hippo run --class ephemeral --disk-path . -- npm exec -- nx run -p badakmini-cli -t lint
rtk ./hippo run --class ephemeral --disk-path . -- npm exec -- nx run -p badakmini-cli -t test:unit
rtk ./hippo run --class ephemeral --disk-path . -- npm exec -- nx run -p badakmini-cli -t test:integration
rtk ./hippo run --class ephemeral --disk-path . -- npm exec -- nx run -p badakmini-cli-e2e -t test:e2e
rtk ./hippo run --class ephemeral --disk-path . -- npm exec -- nx run -p badakmini-cli -t test:coverage:unit
rtk ./hippo run --class ephemeral --disk-path . -- npm exec -- nx run -p badakmini-cli -t test:coverage:integration
rtk ./hippo run --class ephemeral --disk-path . -- npm exec -- nx run -p badakmini-cli -t test:coverage:behaviour
rtk ./hippo run --class ephemeral --disk-path . -- npm exec -- nx run -p badakmini-cli -t test:coverage
rtk ./hippo run --class ephemeral --disk-path . -- npm exec -- nx run -p badakmini-cli -t test:quick
```

Badak Mini is pinned to Go 1.26.6. Its command tree uses exact-pinned Cobra as its only direct production library, while
its production validators remain standard-library-only; the dependency guard confines Cobra to the command adapter. Its
development module also pins GolangCI-Lint and govulncheck; run them through `go -C apps/badakmini-cli tool <name>`.
Lint starts from every applicable nondeprecated check, uses strict Go formatters, and treats every finding as blocking.
Unit coverage owns every `internal/...` package; integration coverage owns only `internal/cli`; both exact numeric gates
require at least 99% statements. The thin `cmd/badak-mini` process entrypoint is owned by the dedicated
[`badakmini-cli-e2e`](../badakmini-cli-e2e/README.md) project and is outside both numeric denominators. `tests/unit` is
the Gherkin unit adapter. Badak Mini is an intentional replacement for the former shell governance checker, not a
general Rhino CLI port.

Its canonical [C4 architecture model](../../specs/apps/badakmini-cli/architecture.md) describes the current as-built
boundaries. Its [canonical behaviour corpus](../../specs/apps/badakmini-cli/README.md) is executed by unit,
local-integration, and process E2E adapters. Unit runs every scenario; either or both higher layers may be independently
exempt only for documented boundary mismatches with substantive alternative proof. The owner module pins Actionlint and
owns `npm run check:workflows`, which validates the scheduled GitHub Actions workflow before it is relied on.
