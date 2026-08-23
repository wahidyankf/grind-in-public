# Technical Design

## Architecture

```text
specs/apps/badakmini-cli/behavior/**/*.feature
                         |
                 recursive catalog
                         |
            shared bindings + Driver contract
                   /          |          \
           unit driver   integration   process E2E
            doubles      local only      built CLI
```

Godog `v0.16.0` supplies Gherkin parsing and scenario execution. Repository-owned support records every binding as typed metadata before registering it with Godog. The same metadata lets `test:coverage:behavior` match every expanded step exactly once and reject unused bindings without relying on runtime output parsing.

The shared driver exposes scenario setup, command invocation, exit code, stdout, and stderr. Bindings translate English steps into that interface only. Unit setup programs injected CLI services and never touches the operating system. Integration setup creates isolated local Git and filesystem fixtures and calls the CLI in process. E2E setup creates equivalent fixtures and launches the built `badak-mini` binary with a 30-second timeout. The process driver requires `BADAKMINI_BIN` to name an absolute, existing, executable file and fails clearly when that contract is not met; it never discovers or builds the binary itself.

## Production Boundary

Move command orchestration into `apps/badakmini-cli/internal/cli`. `Run(context.Context, Runtime, []string) int` owns parsing, repository discovery, dispatch, streams, and exit behavior. `Runtime` supplies repository discovery, stdin/stdout/stderr, and command operations. The production constructor adapts the existing standard-library packages; `cmd/badak-mini/main.go` only constructs it and exits.

Governance, Markdown-link, parity, rule-change, and CLI tests are reclassified before the final unit target is accepted. Pure helpers and collaborator-double cases remain beside their owning production files. Cases requiring real files, directories, symlinks, Git repositories, staged indexes, subprocesses, or working-directory changes move into one responsibility-named file below `tests/integration`. Production gains narrow filesystem, Git/process, or working-directory collaborators only where needed to establish the boundary; no symbol is exported solely to make a test possible.

## Static Enforcement

`test:coverage:behavior` recursively discovers `.feature` files and fails when none exist. It validates the repository's exact primary Given/When/Then cardinality, expands outlines, and verifies every required adapter against the same normalized feature/scenario/step catalog. It rejects undefined, ambiguous, or unused bindings, missing driver methods, tags that omit a layer, direct E2E feature ownership, and adapter corpus mismatch.

The Nx dependency and input model follows the inspected BeaverNest Badak Mini pair, adapted to this repository's paths and American `behavior` spelling:

- `apps/badakmini-cli-e2e/project.json` declares `"implicitDependencies": ["badakmini-cli"]`.
- The owner app's `typecheck`, `test:unit`, `test:integration`, `test:coverage:unit`, and `test:coverage:integration` targets declare `"default"` plus `{workspaceRoot}/specs/apps/badakmini-cli/behavior/**/*.feature` as inputs.
- The owner app's `test:coverage:behavior`, aggregate `test:coverage`, and `test:quick` targets declare those inputs plus `{workspaceRoot}/apps/badakmini-cli-e2e/tests/**/*`, `{workspaceRoot}/apps/badakmini-cli-e2e/go.mod`, `{workspaceRoot}/apps/badakmini-cli-e2e/go.sum`, `{workspaceRoot}/apps/badakmini-cli-e2e/.golangci.yml`, and `{workspaceRoot}/apps/badakmini-cli-e2e/project.json`, so E2E binding or configuration changes invalidate the owning behavior gate.
- The E2E harness's `typecheck`, `lint`, `test:coverage:behavior`, `test:e2e`, and `test:quick` targets declare `"default"`, `"^default"`, and `{workspaceRoot}/specs/apps/badakmini-cli/behavior/**/*.feature` as inputs.
- The E2E `test:e2e` target declares `"dependsOn": ["^build", "test:coverage:behavior"]`; its command sets `BADAKMINI_BIN` to absolute `$PWD/apps/badakmini-cli/dist/badak-mini` before entering the E2E module, and its `test:coverage:behavior` target delegates to `badakmini-cli:test:coverage:behavior`.
- The E2E `test:quick` target depends on `badakmini-cli:test:quick`, then runs only E2E typecheck and lint. The owner quick gate runs unit scenarios, unit coverage, and static behavior coverage. Runtime integration and process E2E remain outside quick and hooks.

The recursive input makes an added, edited, renamed, nested, or deleted feature file invalidate every applicable target. The behavior gate then rejects a new or edited step until every required adapter resolves it exactly once, rejects a removed or renamed step while its old binding remains unused, and rejects an E2E binding or configuration edit that breaks adapter parity.

The owner `test:unit` command is exactly `go -C apps/badakmini-cli test ./cmd/... ./internal/... ./tests/bdd ./tests/unit`; it cannot select `tests/integration` or the separate E2E module. The E2E module requires Godog `v0.16.0` and the owner module at `v0.0.0`, replaces the owner module with `../badakmini-cli`, and carries the same GolangCI-Lint tool directive as the owner. Its targets use these exact commands: `go -C apps/badakmini-cli-e2e vet ./...`, `go -C apps/badakmini-cli-e2e tool golangci-lint run`, `npm exec nx -- run badakmini-cli:test:coverage:behavior`, and `BADAKMINI_BIN="$PWD/apps/badakmini-cli/dist/badak-mini" go -C apps/badakmini-cli-e2e test ./tests`.

The owner Go module also carries `github.com/rhysd/actionlint/cmd/actionlint` as a tool dependency with module version `v1.7.12`. The root `check:workflows` script invokes it as `go -C apps/badakmini-cli tool actionlint "$PWD"/.github/workflows/*.yml`; the absolute workspace-root glob remains valid after Go applies `-C`. Actionlint is the authoritative GitHub Actions syntax and schema check, while Prettier separately enforces YAML formatting.

## Coverage

`test:coverage:unit` runs current same-package tests plus `tests/bdd` and `tests/unit` with `-coverpkg=./internal/...`, writes `local-tmp/badakmini-unit.out`, and enforces the aggregate statement total reported by `go tool cover -func` at 99%. Focused branch cases live beside their owning production files; `tests/unit` remains only the Gherkin adapter. `cmd/badak-mini` is a thin process entrypoint owned by E2E and is outside the unit denominator.

`test:coverage:integration` runs only `tests/integration` with `-coverpkg=./internal/cli`, writes `local-tmp/badakmini-integration.out`, and enforces that profile's aggregate statement total at 99%. Integration owns CLI orchestration at the local filesystem and process boundary; governance, Markdown-link, parity, and rule-change validators remain unit-owned. The app README records these denominators and ownership boundaries.

## Compulsory Project Roles

| Project role | Required behavior adapters and gates |
| --- | --- |
| Application | Unit, local-only integration, and E2E in a dedicated Nx application, all consuming the application's recursive canonical corpus. |
| Library | Unit; add local-only integration only when the library owns a real filesystem, database, process, or similar local boundary. Never put E2E in a lib. |
| Dedicated E2E app | E2E for its owning application's corpus; it owns no separate corpus and may omit meaningless unit and numeric coverage targets. |

Each project README states its corpus path, required adapters, Nx targets, and any inapplicable library integration layer. An omitted library integration adapter means the library owns no local integration boundary; behavior needing public-system proof belongs to a consuming application's E2E corpus.

## File Impact

```text
.
+-- .github/workflows/full-bdd.yml
+-- AGENTS.md
+-- apps/badakmini-cli/
|   +-- .golangci.yml
|   +-- README.md
|   +-- cmd/badak-mini/main.go
|   +-- cmd/badak-mini/main_test.go
|   +-- go.mod
|   +-- go.sum
|   +-- internal/cli/*.go
|   +-- internal/governance/check.go
|   +-- internal/governance/check_test.go
|   +-- internal/markdownlinks/check.go
|   +-- internal/markdownlinks/check_test.go
|   +-- internal/parity/check.go
|   +-- internal/parity/check_test.go
|   +-- internal/rulechange/detect.go
|   +-- internal/rulechange/detect_test.go
|   +-- project.json
|   +-- tests/bdd/bindings.go
|   +-- tests/bdd/catalog.go
|   +-- tests/bdd/driver.go
|   +-- tests/bdd/state.go
|   +-- tests/bdd/suite.go
|   +-- tests/bdd/adapter_parity_test.go
|   +-- tests/bdd/binding_compliance_test.go
|   +-- tests/bdd/boundary_policy_test.go
|   +-- tests/bdd/driver_compliance_test.go
|   +-- tests/bdd/feature_compliance_test.go
|   +-- tests/integration/cli_test.go
|   +-- tests/integration/driver_test.go
|   +-- tests/integration/features_test.go
|   +-- tests/integration/governance_test.go
|   +-- tests/integration/markdownlinks_test.go
|   +-- tests/integration/parity_test.go
|   +-- tests/integration/rulechange_test.go
|   +-- tests/unit/driver_test.go
|   +-- tests/unit/features_test.go
+-- apps/badakmini-cli-e2e/
|   +-- .golangci.yml
|   +-- README.md
|   +-- go.mod
|   +-- go.sum
|   +-- project.json
|   +-- tests/driver_test.go
|   +-- tests/features_test.go
+-- package.json
+-- repo-governance/development/README.md
+-- repo-governance/development/badakmini-cli-policy.md
+-- repo-governance/development/behavior-driven-development-policy.md
+-- repo-governance/development/specs-policy.md
+-- repo-governance/development/tdd-policy.md
+-- repo-governance/development/testing-policy.md
+-- repo-governance/development/workspace-commands.md
+-- specs/README.md
+-- specs/apps/badakmini-cli/README.md
+-- specs/apps/badakmini-cli/behavior/capability-parity.feature
+-- specs/apps/badakmini-cli/behavior/cli-contract.feature
+-- specs/apps/badakmini-cli/behavior/instruction-size.feature
+-- specs/apps/badakmini-cli/behavior/markdown-links.feature
+-- specs/apps/badakmini-cli/behavior/rule-change.feature
+-- plans/backlog/README.md
+-- plans/in-progress/README.md
+-- plans/in-progress/badakmini-layered-bdd/{README.md,brd.md,prd.md,tech-docs.md,delivery.md,learnings.md}
+-- plans/done/README.md
+-- plans/done/YYYY-MM-DD__badakmini-layered-bdd/{README.md,brd.md,prd.md,tech-docs.md,delivery.md,learnings.md}
```

## Dependency and Compatibility Decisions

- Pin `github.com/cucumber/godog v0.16.0` exactly. Owner production excludes `$test` and `tests/bdd/**/*.go` from its standard-library-only Depguard rule; one test-support rule covers both and allows only the standard library, the owner module, and Godog. The E2E module separately allows only those same three dependency families.
- Pin `github.com/rhysd/actionlint v1.7.12` exactly in the owner module and expose `github.com/rhysd/actionlint/cmd/actionlint` through its Go tool block; `check:workflows` uses that module-owned tool rather than an ambient executable.
- Keep local `behavior/` spelling and `test:coverage:behavior` target naming.
- Keep all existing command text and exit codes compatible.
- Do not restore the removed central project-target validator; the canonical policy, project targets, Nx inputs, and quality gates follow BeaverNest's model.

## Scheduled Verification

`.github/workflows/full-bdd.yml` uses UTC cron `0 23 * * *`, which is daily 06:00 WIB, plus `workflow_dispatch`. Its `ubuntu-latest` job grants only `contents: read`, references no secrets, and uses `actions/checkout@v6`, `actions/setup-node@v6` with Node.js `24.15.0` plus npm caching, and `actions/setup-go@v6` with Go `1.26.6`, Go caching, and `cache-dependency-path: apps/*/go.sum`. It then runs `npm ci` followed by the canonical root scheduled command:

```sh
npm run test:scheduled
```

The `test:scheduled` script expands to `nx run badakmini-cli:test:quick && nx run badakmini-cli:test:coverage:integration && nx run badakmini-cli-e2e:test:e2e`. That exact order proves the owner quick gate, including unit 99% coverage and behavior coverage, before integration 99% coverage, while preserving integration before process E2E in both local and scheduled execution. `npm run check:workflows` validates the completed workflow structurally before Phase 3 and final-delivery gates accept it.

## Rollback

Each phase lands as a coherent commit. Revert the latest phase commit to return to the previous green boundary. The production orchestration refactor preserves behavior and can remain even if the BDD adapters are reverted. Removing the dependency also removes its `go.sum` entries and test-only allowlist.
