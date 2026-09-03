# Delivery Checklist

Tags: `[AI]` means an agent completes the item, `[HUMAN]` means the owner completes it, and `[AI+HUMAN]` means an agent
prepares it for owner action.

## Phase 0: Baseline

- [x] [AI] Run `npm install` — acceptance: locked workspace dependencies install successfully without changing
      `package-lock.json`.
- [x] [AI] Run `npm run test:quick` — acceptance: the pre-change quick gate exits 0.
- [x] [AI] Run `npm run format:check` — acceptance: the pre-change format gate exits 0.
- [x] [AI] Run `npm run check:markdown-links` — acceptance: the pre-change Markdown-link gate exits 0.
- [x] [AI] Record the four baseline commands, exit statuses, and concise results in
      `plans/in-progress/badakmini-layered-bdd/learnings.md` — acceptance: each preceding command has one non-secret
      result entry. No separate command applies because this action records evidence from the preceding shell output.

### Phase 0 Gate

> Every check below passes before Phase 1 begins. If any baseline check fails, stop without repairing the repository and
> record the command, exit status, and observed failure in `learnings.md` for owner direction.

- [x] [AI] Run `git status --short` — acceptance: only the plan and its stage index are modified.
- [x] [AI] Run `npm run test:quick` — acceptance: the existing Badak quick gate exits 0.

> **Pause Safety**: The validated plan is the only repository change and the existing application remains green. Safe to
> stop. Resume with `npm run test:quick`.

## Phase 1: Executable Test Foundation

- [x] [AI] Pin `github.com/cucumber/godog v0.16.0` in `apps/badakmini-cli/go.mod` — acceptance:
      `go -C apps/badakmini-cli list -m github.com/cucumber/godog` prints `github.com/cucumber/godog v0.16.0`.
- [x] [AI] Add `github.com/rhysd/actionlint/cmd/actionlint` to the `tool` block in `apps/badakmini-cli/go.mod` and pin
      its module to `v1.7.12` — acceptance: `go -C apps/badakmini-cli list -m github.com/rhysd/actionlint` prints
      `github.com/rhysd/actionlint v1.7.12`, and
      `rg -n '^\s*github\.com/rhysd/actionlint/cmd/actionlint$' apps/badakmini-cli/go.mod` finds the tool directive.
- [x] [AI] Change the owner `standard-library-only` Depguard rule in `apps/badakmini-cli/.golangci.yml` to files `$all`,
      `!$test`, and `!**/tests/bdd/**/*.go` while preserving its standard-library and owner-module allowlist —
      acceptance:
      `node -e 'const s=require("fs").readFileSync("apps/badakmini-cli/.golangci.yml","utf8");const r=/standard-library-only:[\s\S]*?files:[\s\S]*?- \$all[\s\S]*?- "!\$test"[\s\S]*?- "!\*\*\/tests\/bdd\/\*\*\/\*\.go"[\s\S]*?allow:[\s\S]*?- \$gostd[\s\S]*?- github\.com\/wahidyankf\/grind-in-public\/apps\/badakmini-cli/;if(!r.test(s))process.exit(1)'`
      exits 0.
- [x] [AI] Add a `test-support` Depguard rule to `apps/badakmini-cli/.golangci.yml` covering `$test` plus
      `**/tests/bdd/**/*.go` whose allowlist is exactly `$gostd`,
      `github.com/wahidyankf/grind-in-public/apps/badakmini-cli`, and `github.com/cucumber/godog` — acceptance:
      `node -e 'const s=require("fs").readFileSync("apps/badakmini-cli/.golangci.yml","utf8");const r=/test-support:[\s\S]*?files:[\s\S]*?- \$test[\s\S]*?- "\*\*\/tests\/bdd\/\*\*\/\*\.go"[\s\S]*?allow:[\s\S]*?- \$gostd[\s\S]*?- github\.com\/wahidyankf\/grind-in-public\/apps\/badakmini-cli[\s\S]*?- github\.com\/cucumber\/godog/;if(!r.test(s))process.exit(1)'`
      exits 0.
- [x] [AI] Implement the recursive feature catalog in `apps/badakmini-cli/tests/bdd/catalog.go` — acceptance: source
      review confirms recursive discovery and its fixtures are present. No Go command runs yet because the
      Godog-importing shared BDD sources must exist before the single module tidy action.
- [x] [AI] Implement scenario-scoped state in `apps/badakmini-cli/tests/bdd/state.go` — acceptance: source review
      confirms scenario-isolation behavior and its fixtures are present. No Go command runs yet because the
      Godog-importing shared BDD sources must exist before the single module tidy action.
- [x] [AI] Implement the shared driver contract in `apps/badakmini-cli/tests/bdd/driver.go` — acceptance: source review
      confirms complete and incomplete driver fixtures are present. No Go command runs yet because the Godog-importing
      shared BDD sources must exist before the single module tidy action.
- [x] [AI] Implement the typed binding registry in `apps/badakmini-cli/tests/bdd/bindings.go` — acceptance: source
      review confirms typed binding-registry fixtures are present. No Go command runs yet because the Godog-importing
      shared BDD sources must exist before the single module tidy action.
- [x] [AI] Implement the Godog suite initializer in `apps/badakmini-cli/tests/bdd/suite.go` and import
      `github.com/cucumber/godog` there — acceptance: source review confirms the initializer executes a fixture feature
      through Godog. No Go command runs yet because the immediately following module tidy action resolves the
      now-complete shared BDD imports.
- [x] [AI] Refresh `apps/badakmini-cli/go.sum` for the now-imported Godog test dependency and actionlint tool dependency
      by running `go -C apps/badakmini-cli mod tidy` — acceptance: repeating `go -C apps/badakmini-cli mod tidy` leaves
      both module files unchanged, and
      `node -e 'const s=require("fs").readFileSync("apps/badakmini-cli/go.mod","utf8");if(!/github\.com\/cucumber\/godog v0\.16\.0/.test(s)||!/github\.com\/rhysd\/actionlint v1\.7\.12/.test(s)||!/^\s*github\.com\/rhysd\/actionlint\/cmd\/actionlint$/m.test(s))process.exit(1)'`
      exits 0, proving tidy retained both exact module pins and the actionlint tool directive.
- [x] [AI] Run `go -C apps/badakmini-cli test ./tests/bdd` after module resolution — acceptance: recursive discovery,
      scenario isolation, complete and incomplete driver, typed binding-registry, and Godog fixture-execution checks all
      exit 0.
- [x] [AI] Move command orchestration behind `Run(context.Context, Runtime, []string) int` in the new
      `apps/badakmini-cli/internal/cli/` package; choose Go filenames by responsibility because the package is new —
      acceptance: `go -C apps/badakmini-cli test ./internal/cli` exits 0 with existing command text and exit-code
      assertions.
- [x] [AI] Reduce `apps/badakmini-cli/cmd/badak-mini/main.go` to production runtime construction and process exit —
      acceptance: `go -C apps/badakmini-cli test ./cmd/badak-mini` exits 0 with existing entrypoint assertions.
- [x] [AI] Add malformed-feature cases to `apps/badakmini-cli/tests/bdd/feature_compliance_test.go` — acceptance:
      `go -C apps/badakmini-cli test ./tests/bdd` exits 0 after proving malformed Gherkin is rejected.
- [x] [AI] Add empty-feature cases to `apps/badakmini-cli/tests/bdd/feature_compliance_test.go` — acceptance:
      `go -C apps/badakmini-cli test ./tests/bdd` exits 0 after proving an empty feature is rejected.
- [x] [AI] Add primary Given/When/Then cardinality cases to `apps/badakmini-cli/tests/bdd/feature_compliance_test.go` —
      acceptance: `go -C apps/badakmini-cli test ./tests/bdd` exits 0 after proving missing or duplicated primary steps
      are rejected.
- [x] [AI] Add expanded-scenario-count cases to `apps/badakmini-cli/tests/bdd/feature_compliance_test.go` — acceptance:
      `go -C apps/badakmini-cli test ./tests/bdd` exits 0 after proving outline-expansion drift is rejected.
- [x] [AI] Add exact recursive-corpus cases to `apps/badakmini-cli/tests/bdd/adapter_parity_test.go` — acceptance:
      `go -C apps/badakmini-cli test ./tests/bdd` exits 0 after proving corpus drift is rejected.
- [x] [AI] Add undefined-binding cases to `apps/badakmini-cli/tests/bdd/binding_compliance_test.go` — acceptance:
      `go -C apps/badakmini-cli test ./tests/bdd` exits 0 after proving an undefined step is rejected.
- [x] [AI] Add ambiguous-binding cases to `apps/badakmini-cli/tests/bdd/binding_compliance_test.go` — acceptance:
      `go -C apps/badakmini-cli test ./tests/bdd` exits 0 after proving an ambiguous step is rejected.
- [x] [AI] Add unused-binding cases to `apps/badakmini-cli/tests/bdd/binding_compliance_test.go` — acceptance:
      `go -C apps/badakmini-cli test ./tests/bdd` exits 0 after proving a stale binding is rejected.
- [x] [AI] Add incomplete-driver cases to `apps/badakmini-cli/tests/bdd/driver_compliance_test.go` — acceptance:
      `go -C apps/badakmini-cli test ./tests/bdd` exits 0 after proving a missing driver operation is rejected.
- [x] [AI] Add adapter-parity cases to `apps/badakmini-cli/tests/bdd/adapter_parity_test.go` — acceptance:
      `go -C apps/badakmini-cli test ./tests/bdd` exits 0 after proving corpus mismatch is rejected.
- [x] [AI] Add layer-filter cases to `apps/badakmini-cli/tests/bdd/adapter_parity_test.go` — acceptance:
      `go -C apps/badakmini-cli test ./tests/bdd` exits 0 after proving adapter-specific scenario filtering is rejected.
- [x] [AI] Refactor `apps/badakmini-cli/internal/governance/check.go` only as needed to inject its filesystem
      collaborator through the production boundary, without exporting internals solely for tests — acceptance:
      `go -C apps/badakmini-cli test ./internal/governance` exits 0.
- [x] [AI] Keep governance pure-helper and double-based cases beside production in
      `apps/badakmini-cli/internal/governance/check_test.go` — acceptance:
      `go -C apps/badakmini-cli test ./internal/governance` exits 0 without creating a real temporary filesystem.
- [x] [AI] Move governance cases that require real files, directories, or symlinks into
      `apps/badakmini-cli/tests/integration/governance_test.go` under the required `TestIntegrationGovernance...` prefix
      — acceptance:
      `set -o pipefail; go -C apps/badakmini-cli test -json ./tests/integration -run '^TestIntegrationGovernance' | rg '"Action":"pass".*"Test":"TestIntegrationGovernance[^" ]*"'`
      proves at least one named migrated test passed.
- [x] [AI] Refactor `apps/badakmini-cli/internal/markdownlinks/check.go` only as needed to inject filesystem and Git
      collaborators through the production boundary, without exporting internals solely for tests — acceptance:
      `go -C apps/badakmini-cli test ./internal/markdownlinks` exits 0.
- [x] [AI] Keep Markdown parsing, anchor, path, and double-based cases beside production in
      `apps/badakmini-cli/internal/markdownlinks/check_test.go` — acceptance:
      `go -C apps/badakmini-cli test ./internal/markdownlinks` exits 0 without creating a real filesystem or Git
      process.
- [x] [AI] Move Markdown-link cases that require real files, symlinks, or Git index state into
      `apps/badakmini-cli/tests/integration/markdownlinks_test.go` under the required `TestIntegrationMarkdownLinks...`
      prefix — acceptance:
      `set -o pipefail; go -C apps/badakmini-cli test -json ./tests/integration -run '^TestIntegrationMarkdownLinks' | rg '"Action":"pass".*"Test":"TestIntegrationMarkdownLinks[^" ]*"'`
      proves at least one named migrated test passed.
- [x] [AI] Refactor `apps/badakmini-cli/internal/parity/check.go` only as needed to inject its filesystem collaborator
      through the production boundary, without exporting internals solely for tests — acceptance:
      `go -C apps/badakmini-cli test ./internal/parity` exits 0.
- [x] [AI] Keep parity comparison, normalization, and double-based cases beside production in
      `apps/badakmini-cli/internal/parity/check_test.go` — acceptance: `go -C apps/badakmini-cli test ./internal/parity`
      exits 0 without creating a real temporary filesystem.
- [x] [AI] Move parity cases that require real harness directories, files, or symlinks into
      `apps/badakmini-cli/tests/integration/parity_test.go` under the required `TestIntegrationParity...` prefix —
      acceptance:
      `set -o pipefail; go -C apps/badakmini-cli test -json ./tests/integration -run '^TestIntegrationParity' | rg '"Action":"pass".*"Test":"TestIntegrationParity[^" ]*"'`
      proves at least one named migrated test passed.
- [x] [AI] Refactor `apps/badakmini-cli/internal/rulechange/detect.go` only as needed to inject filesystem and
      Git-process collaborators through the production boundary, without exporting internals solely for tests —
      acceptance: `go -C apps/badakmini-cli test ./internal/rulechange` exits 0.
- [x] [AI] Keep rule-path selection, payload parsing, normalization, and double-based cases beside production in
      `apps/badakmini-cli/internal/rulechange/detect_test.go` — acceptance:
      `go -C apps/badakmini-cli test ./internal/rulechange` exits 0 without creating a real filesystem or Git process.
- [x] [AI] Move rule-change cases that require a real Git repository, staged paths, or filesystem payload fixtures into
      `apps/badakmini-cli/tests/integration/rulechange_test.go` under the required `TestIntegrationRuleChange...` prefix
      — acceptance:
      `set -o pipefail; go -C apps/badakmini-cli test -json ./tests/integration -run '^TestIntegrationRuleChange' | rg '"Action":"pass".*"Test":"TestIntegrationRuleChange[^" ]*"'`
      proves at least one named migrated test passed.
- [x] [AI] Refactor CLI orchestration into `apps/badakmini-cli/internal/cli/` using the existing production `Runtime`
      ports for filesystem, Git/process, working-directory, and stream collaborators, without exporting internals solely
      for tests — acceptance: `go -C apps/badakmini-cli test ./internal/cli` exits 0.
- [x] [AI] Keep argument parsing, dispatch, exit-code, stream-failure, and double-based cases beside orchestration in
      `apps/badakmini-cli/internal/cli/run_test.go` — acceptance: `go -C apps/badakmini-cli test ./internal/cli` exits 0
      without creating a real filesystem or process.
- [x] [AI] Move CLI cases from `apps/badakmini-cli/cmd/badak-mini/main_test.go` that require a real repository, working
      directory, filesystem, or Git subprocess into `apps/badakmini-cli/tests/integration/cli_test.go` under the
      required `TestIntegrationCLI...` prefix — acceptance:
      `set -o pipefail; go -C apps/badakmini-cli test -json ./tests/integration -run '^TestIntegrationCLI' | rg '"Action":"pass".*"Test":"TestIntegrationCLI[^" ]*"'`
      proves at least one named migrated test passed.
- [x] [AI] Preserve only thin production-construction assertions beside `apps/badakmini-cli/cmd/badak-mini/main.go` in
      `apps/badakmini-cli/cmd/badak-mini/main_test.go` — acceptance: `go -C apps/badakmini-cli test ./cmd/badak-mini`
      exits 0 without creating a real repository, filesystem fixture, or subprocess.
- [x] [AI] Create `apps/badakmini-cli/tests/unit/driver_test.go` with the unit driver backed only by injected doubles —
      acceptance: `go -C apps/badakmini-cli test ./tests/unit -run '^$'` compiles the unit adapter without executing
      canonical scenarios.
- [x] [AI] Replace `badakmini-cli:test:unit` in `apps/badakmini-cli/project.json` with the unit-only command
      `go -C apps/badakmini-cli test ./cmd/... ./internal/... ./tests/bdd ./tests/unit` — acceptance:
      `node -e 'const p=require("./apps/badakmini-cli/project.json"); if(p.targets["test:unit"].command!=="go -C apps/badakmini-cli test ./cmd/... ./internal/... ./tests/bdd ./tests/unit")process.exit(1)'`
      exits 0 and the command contains neither `./tests/integration` nor the separate E2E module.
- [x] [AI] Add the exact `TestUnitBoundaryPolicy` check to `apps/badakmini-cli/tests/bdd/boundary_policy_test.go` —
      acceptance:
      `set -o pipefail; go -C apps/badakmini-cli test -json ./tests/bdd -run '^TestUnitBoundaryPolicy$' | rg '"Action":"pass".*"Test":"TestUnitBoundaryPolicy"'`
      exits 0 after proving unit tests use only doubles.
- [x] [AI] Add the exact `TestIntegrationBoundaryPolicy` check to `apps/badakmini-cli/tests/bdd/boundary_policy_test.go`
      — acceptance:
      `set -o pipefail; go -C apps/badakmini-cli test -json ./tests/bdd -run '^TestIntegrationBoundaryPolicy$' | rg '"Action":"pass".*"Test":"TestIntegrationBoundaryPolicy"'`
      exits 0 after proving integration uses isolated local resources without network access.
- [x] [AI] Run
      `set -o pipefail; go -C apps/badakmini-cli test -json ./tests/bdd -run '^TestUnitBoundaryPolicy$' | rg '"Action":"pass".*"Test":"TestUnitBoundaryPolicy"'`
      — acceptance: the required static check passed after finding no real filesystem setup, `os/exec` use, Git
      subprocess, or direct system collaborator in any test selected by `test:unit`.
- [x] [AI] Run
      `set -o pipefail; go -C apps/badakmini-cli test -json ./tests/bdd -run '^TestIntegrationBoundaryPolicy$' | rg '"Action":"pass".*"Test":"TestIntegrationBoundaryPolicy"'`
      — acceptance: the required static check passed after confirming the five exact integration files own the migrated
      real filesystem and Git/process cases and use no network, including loopback.
- [x] [AI] Run `npm exec nx -- run badakmini-cli:test:unit` — acceptance: the unit-only target passes after the static
      boundary policy and launches no real filesystem, Git, or subprocess fixture.
- [x] [AI] Run `go -C apps/badakmini-cli test ./tests/integration` — acceptance: all migrated real local-boundary cases
      pass in their integration-owned files without network access.
- [x] [AI] Add `test:coverage:unit` to `apps/badakmini-cli/project.json` with exact command
      `mkdir -p local-tmp && go -C apps/badakmini-cli test -coverpkg=./internal/... -coverprofile=../../local-tmp/badakmini-unit.out ./cmd/... ./internal/... ./tests/bdd ./tests/unit && go -C apps/badakmini-cli tool cover -func=../../local-tmp/badakmini-unit.out | awk '/^total:/ { coverage=$3; sub(/%$/, "", coverage); printf "unit statement coverage: %.1f%%\n", coverage; found=1; if (coverage + 0 < 99) exit 1 } END { if (!found) exit 1 }'`
      — acceptance: `npm exec nx -- run badakmini-cli:test:coverage:unit` aggregates the profile's total statement
      percentage across all `internal/...` runtime code and exits 0 only at or above 99%.
- [x] [AI] Add the owner `default` plus recursive corpus inputs to `badakmini-cli:test:coverage:unit` in
      `apps/badakmini-cli/project.json` — acceptance: `npm exec nx -- show project badakmini-cli --json` shows both
      inputs on `test:coverage:unit`; the recursive input becomes populated when Phase 2 creates the corpus.
- [x] [AI] Refactor `apps/badakmini-cli/internal/cli/run.go` so `Runtime` owns repository discovery, streams,
      governance-check, Markdown-check, staged-path-listing, and parity-check collaborators; make `Run` execute its
      production command handlers through those low-level fields instead of bypassable top-level handler doubles —
      acceptance: `go -C apps/badakmini-cli test ./internal/cli` exits 0 with every production handler branch driven by
      injected results and no real filesystem, Git process, subprocess, or network fixture.
- [x] [AI] Replace the host-root `Check(string)` boundary in `apps/badakmini-cli/internal/governance/check.go` with
      production-used `CheckFS(fs.FS)` and update same-package tests to call it — acceptance:
      `go -C apps/badakmini-cli test ./internal/governance` exits 0 and the package imports no `os` solely to construct
      a host filesystem.
- [x] [AI] Replace the host-root `Check(string)` boundary in `apps/badakmini-cli/internal/parity/check.go` with
      production-used `CheckFS(fs.FS)` and update same-package tests to call it — acceptance:
      `go -C apps/badakmini-cli test ./internal/parity` exits 0 and the package imports no `os` solely to construct a
      host filesystem.
- [x] [AI] Export the production-used injected runtime accepted by `apps/badakmini-cli/internal/markdownlinks/check.go`,
      make `Check` accept that runtime, and export tracked-output parsing while removing `os` and `os/exec` adapter
      construction from the package — acceptance: `go -C apps/badakmini-cli test ./internal/markdownlinks` exits 0 with
      file reads, stats, symlink resolution, and tracked paths supplied only through doubles.
- [x] [AI] Replace the Git-running `StagedPaths(string)` boundary in `apps/badakmini-cli/internal/rulechange/detect.go`
      with production-used staged-output parsing and keep path selection pure — acceptance:
      `go -C apps/badakmini-cli test ./internal/rulechange` exits 0 and the package imports no `os/exec`.
- [x] [AI] Move repository-root discovery, Git tracked/staged output, and host filesystem adapter construction into
      `apps/badakmini-cli/cmd/badak-mini/main.go`; add named
      `productionRuntime(io.Reader, io.Writer, io.Writer) cli.Runtime` with the production-used governance, Markdown,
      parity, and rule-change seams, and require `execute` to call that constructor — acceptance:
      `go -C apps/badakmini-cli test ./cmd/badak-mini` exits 0 with help still available before any host adapter
      executes.
- [x] [AI] Add exact `TestProductionRuntimeBindsAllAdapters` construction coverage to
      `apps/badakmini-cli/cmd/badak-mini/main_test.go`; call `productionRuntime` and compare every repository-discovery,
      governance, Markdown, staged-path, and parity function field with its intended named production adapter without
      invoking it — acceptance:
      `set -o pipefail; go -C apps/badakmini-cli test -json ./cmd/badak-mini -run '^TestProductionRuntimeBindsAllAdapters$' | rg '"Action":"pass".*"Test":"TestProductionRuntimeBindsAllAdapters"'`
      exits 0.
- [x] [AI] Run `if rg -n '"os"|"os/exec"' apps/badakmini-cli/internal -g '*.go'; then exit 1; fi` — acceptance: the
      structural assertion exits 0 after proving no unit-owned internal package imports a host adapter.
- [x] [AI] Update `apps/badakmini-cli/tests/integration/governance_test.go` to exercise real files and symlinks through
      `governance.CheckFS(os.DirFS(root))` — acceptance:
      `set -o pipefail; go -C apps/badakmini-cli test -json ./tests/integration -run '^TestIntegrationGovernance' | rg '"Action":"pass".*"Test":"TestIntegrationGovernance[^" ]*"'`
      proves at least one named test passed.
- [x] [AI] Update `apps/badakmini-cli/tests/integration/parity_test.go` to exercise real harness directories and
      symlinks through `parity.CheckFS(os.DirFS(root))` — acceptance:
      `set -o pipefail; go -C apps/badakmini-cli test -json ./tests/integration -run '^TestIntegrationParity' | rg '"Action":"pass".*"Test":"TestIntegrationParity[^" ]*"'`
      proves at least one named test passed.
- [x] [AI] Update `apps/badakmini-cli/tests/integration/markdownlinks_test.go` to construct the exported Markdown
      runtime from real file, stat, symlink, and Git tracked-output collaborators — acceptance:
      `set -o pipefail; go -C apps/badakmini-cli test -json ./tests/integration -run '^TestIntegrationMarkdownLinks' | rg '"Action":"pass".*"Test":"TestIntegrationMarkdownLinks[^" ]*"'`
      proves at least one named test passed.
- [x] [AI] Update `apps/badakmini-cli/tests/integration/rulechange_test.go` to pass real Git staged output through the
      exported rule-change parser — acceptance:
      `set -o pipefail; go -C apps/badakmini-cli test -json ./tests/integration -run '^TestIntegrationRuleChange' | rg '"Action":"pass".*"Test":"TestIntegrationRuleChange[^" ]*"'`
      proves at least one named test passed.
- [x] [AI] Update `apps/badakmini-cli/tests/integration/cli_test.go` to construct `cli.Runtime` with the same real
      exported seams as the production entrypoint — acceptance:
      `set -o pipefail; go -C apps/badakmini-cli test -json ./tests/integration -run '^TestIntegrationCLI' | rg '"Action":"pass".*"Test":"TestIntegrationCLI[^" ]*"'`
      proves at least one named test passed.
- [x] [AI] Add each focused uncovered-branch unit case beside its owning production file under
      `apps/badakmini-cli/internal/<package>/`, using the existing `<production-basename>_test.go` same-package naming
      rule; keep `apps/badakmini-cli/tests/unit/` exclusively as the Gherkin adapter — acceptance:
      `npm exec nx -- run badakmini-cli:test:coverage:unit` exits 0 at or above 99% without broad exclusions or a real
      filesystem, Git process, subprocess, or network fixture.
- [x] [AI] Replace the temporary legacy `badakmini-cli:test:coverage` target in `apps/badakmini-cli/project.json` with
      exact command `npm exec nx -- run badakmini-cli:test:coverage:unit` — acceptance:
      `npm exec nx -- run badakmini-cli:test:coverage` exits 0 through the deterministic all-runtime unit slice; Phase 3
      replaces this temporary alias with the reviewed unit, integration, and behavior coverage composition.

### Phase 1 Gate

> Every check below passes before Phase 2 begins. A failure is fixed inside Phase 1.

- [x] [AI] Run `npm exec nx -- run badakmini-cli:test:unit` — acceptance: all foundation, same-package double-based,
      compliance, and unit-adapter tests exit 0 without selecting integration tests.
- [x] [AI] Run `npm exec nx -- run badakmini-cli:typecheck` — acceptance: Go vet exits 0.
- [x] [AI] Run `npm exec nx -- run badakmini-cli:lint` — acceptance: strict lint exits 0.
- [x] [AI] Run `npm exec nx -- run badakmini-cli:test:coverage:unit` — acceptance: deterministic unit coverage exits 0
      at or above 99% across every `internal/...` runtime package.
- [x] [AI] Run `npm exec nx -- run badakmini-cli:test:quick` — acceptance: the cacheable pre-push target exits 0 with
      typecheck, lint, unit tests, and the unit-only all-runtime coverage slice.

> **Pause Safety**: Production behavior is preserved, unit tests use only doubles, aggregate statement coverage across
> all `internal/...` runtime packages is at or above 99%, and migrated real-boundary cases pass from
> `tests/integration`; the behavior target is added only after its canonical corpus exists in Phase 2. Safe to stop.
> Resume with `npm exec nx -- run badakmini-cli:test:quick`.

## Phase 2: Canonical Behavior and Required Adapters

Each scenario below follows its RED checkbox, GREEN checkbox, then REFACTOR checkbox before the next scenario begins.

- [x] [AI] Create `apps/badakmini-cli/tests/unit/features_test.go` with runnable `TestFeatures` discovery of
      `specs/apps/badakmini-cli/behavior/**/*.feature` through the shared Godog initializer — acceptance:
      `go -C apps/badakmini-cli test ./tests/unit -run '^$'` compiles the canonical unit runner before the first RED
      action without executing the still-empty corpus.

### Help is available outside a repository

- [x] [AI] Add the scenario to `specs/apps/badakmini-cli/behavior/cli-contract.feature` — acceptance:
      `go -C apps/badakmini-cli test -count=1 ./tests/unit -run TestFeatures` exits nonzero because the unit adapter
      lacks the new scenario support.

```gherkin
Scenario: Help is available outside a repository
  Given repository discovery would fail
  When Badak Mini runs with "--help"
  Then the command succeeds and prints usage
```

- [x] [AI] Register the help scenario's missing step expressions in `apps/badakmini-cli/tests/bdd/bindings.go` —
      acceptance: `go -C apps/badakmini-cli test -count=1 ./tests/bdd` exits 0 with each expression registered exactly
      once.
- [x] [AI] Implement the help scenario behavior in `apps/badakmini-cli/tests/unit/driver_test.go` — acceptance:
      `go -C apps/badakmini-cli test -count=1 ./tests/unit -run TestFeatures` exits 0.
- [x] [AI] Refactor the help scenario support in `apps/badakmini-cli/tests/unit/driver_test.go`; keep the shared
      responsibility-based driver filename and create no scenario-named file — acceptance:
      `go -C apps/badakmini-cli test -count=1 ./tests/unit -run TestFeatures` exits 0.

### Governance documents fit the word limit

- [x] [AI] Add the scenario to `specs/apps/badakmini-cli/behavior/instruction-size.feature` — acceptance:
      `go -C apps/badakmini-cli test -count=1 ./tests/unit -run TestFeatures` exits nonzero because the unit adapter
      lacks the fitting-document fixture.

```gherkin
Scenario: Governance documents fit the word limit
  Given a repository whose governance documents fit the word limit
  When Badak Mini runs instruction-size validation
  Then the command succeeds with the word-limit confirmation
```

- [x] [AI] Register the fitting-document scenario's missing step expressions in
      `apps/badakmini-cli/tests/bdd/bindings.go` — acceptance: `go -C apps/badakmini-cli test -count=1 ./tests/bdd`
      exits 0 with each expression registered exactly once.
- [x] [AI] Implement the fitting-document fixture in `apps/badakmini-cli/tests/unit/driver_test.go` — acceptance:
      `go -C apps/badakmini-cli test -count=1 ./tests/unit -run TestFeatures` exits 0.
- [x] [AI] Refactor the word-limit success support in `apps/badakmini-cli/tests/unit/driver_test.go`; keep the shared
      responsibility-based driver filename and create no scenario-named file — acceptance:
      `go -C apps/badakmini-cli test -count=1 ./tests/unit -run TestFeatures` exits 0.

### A governance document exceeds the word limit

- [x] [AI] Add the scenario to `specs/apps/badakmini-cli/behavior/instruction-size.feature` — acceptance:
      `go -C apps/badakmini-cli test -count=1 ./tests/unit -run TestFeatures` exits nonzero because the unit adapter
      lacks the oversized-document fixture.

```gherkin
Scenario: A governance document exceeds the word limit
  Given a repository with an oversized agent instruction file
  When Badak Mini runs instruction-size validation
  Then the command fails with the oversized document diagnostic
```

- [x] [AI] Register the oversized-document scenario's missing step expressions in
      `apps/badakmini-cli/tests/bdd/bindings.go` — acceptance: `go -C apps/badakmini-cli test -count=1 ./tests/bdd`
      exits 0 with each expression registered exactly once.
- [x] [AI] Implement the oversized-document fixture in `apps/badakmini-cli/tests/unit/driver_test.go` — acceptance:
      `go -C apps/badakmini-cli test -count=1 ./tests/unit -run TestFeatures` exits 0.
- [x] [AI] Refactor the word-limit failure support in `apps/badakmini-cli/tests/unit/driver_test.go`; keep the shared
      responsibility-based driver filename and create no scenario-named file — acceptance:
      `go -C apps/badakmini-cli test -count=1 ./tests/unit -run TestFeatures` exits 0.

### Repository Markdown links resolve

- [x] [AI] Add the scenario to `specs/apps/badakmini-cli/behavior/markdown-links.feature` — acceptance:
      `go -C apps/badakmini-cli test -count=1 ./tests/unit -run TestFeatures` exits nonzero because the unit adapter
      lacks the valid-link fixture.

```gherkin
Scenario: Repository Markdown links resolve
  Given a repository whose tracked Markdown links resolve
  When Badak Mini runs Markdown-link validation
  Then the command succeeds with the link confirmation
```

- [x] [AI] Register the valid-link scenario's missing step expressions in `apps/badakmini-cli/tests/bdd/bindings.go` —
      acceptance: `go -C apps/badakmini-cli test -count=1 ./tests/bdd` exits 0 with each expression registered exactly
      once.
- [x] [AI] Implement the valid-link fixture in `apps/badakmini-cli/tests/unit/driver_test.go` — acceptance:
      `go -C apps/badakmini-cli test -count=1 ./tests/unit -run TestFeatures` exits 0.
- [x] [AI] Refactor the valid-link support in `apps/badakmini-cli/tests/unit/driver_test.go`; keep the shared
      responsibility-based driver filename and create no scenario-named file — acceptance:
      `go -C apps/badakmini-cli test -count=1 ./tests/unit -run TestFeatures` exits 0.

### A tracked Markdown link is broken

- [x] [AI] Add the scenario to `specs/apps/badakmini-cli/behavior/markdown-links.feature` — acceptance:
      `go -C apps/badakmini-cli test -count=1 ./tests/unit -run TestFeatures` exits nonzero because the unit adapter
      lacks the broken-link fixture.

```gherkin
Scenario: A tracked Markdown link is broken
  Given a repository with a broken tracked Markdown link
  When Badak Mini runs Markdown-link validation
  Then the command fails with the missing-target diagnostic
```

- [x] [AI] Register the broken-link scenario's missing step expressions in `apps/badakmini-cli/tests/bdd/bindings.go` —
      acceptance: `go -C apps/badakmini-cli test -count=1 ./tests/bdd` exits 0 with each expression registered exactly
      once.
- [x] [AI] Implement the broken-link fixture in `apps/badakmini-cli/tests/unit/driver_test.go` — acceptance:
      `go -C apps/badakmini-cli test -count=1 ./tests/unit -run TestFeatures` exits 0.
- [x] [AI] Refactor the broken-link support in `apps/badakmini-cli/tests/unit/driver_test.go`; keep the shared
      responsibility-based driver filename and create no scenario-named file — acceptance:
      `go -C apps/badakmini-cli test -count=1 ./tests/unit -run TestFeatures` exits 0.

### Harness capabilities match

- [x] [AI] Add the scenario to `specs/apps/badakmini-cli/behavior/capability-parity.feature` — acceptance:
      `go -C apps/badakmini-cli test -count=1 ./tests/unit -run TestFeatures` exits nonzero because the unit adapter
      lacks the matching-harness fixture.

```gherkin
Scenario: Harness capabilities match
  Given a repository whose harness capabilities match
  When Badak Mini runs capability-parity validation
  Then the command succeeds with the parity confirmation
```

- [x] [AI] Register the matching-harness scenario's missing step expressions in
      `apps/badakmini-cli/tests/bdd/bindings.go` — acceptance: `go -C apps/badakmini-cli test -count=1 ./tests/bdd`
      exits 0 with each expression registered exactly once.
- [x] [AI] Implement the matching-harness fixture in `apps/badakmini-cli/tests/unit/driver_test.go` — acceptance:
      `go -C apps/badakmini-cli test -count=1 ./tests/unit -run TestFeatures` exits 0.
- [x] [AI] Refactor the matching-harness support in `apps/badakmini-cli/tests/unit/driver_test.go`; keep the shared
      responsibility-based driver filename and create no scenario-named file — acceptance:
      `go -C apps/badakmini-cli test -count=1 ./tests/unit -run TestFeatures` exits 0.

### A harness capability is missing

- [x] [AI] Add the scenario to `specs/apps/badakmini-cli/behavior/capability-parity.feature` — acceptance:
      `go -C apps/badakmini-cli test -count=1 ./tests/unit -run TestFeatures` exits nonzero because the unit adapter
      lacks the mismatched-harness fixture.

```gherkin
Scenario: A harness capability is missing
  Given a repository with a harness missing a shared subagent
  When Badak Mini runs capability-parity validation
  Then the command fails with the parity diagnostic
```

- [x] [AI] Register the mismatched-harness scenario's missing step expressions in
      `apps/badakmini-cli/tests/bdd/bindings.go` — acceptance: `go -C apps/badakmini-cli test -count=1 ./tests/bdd`
      exits 0 with each expression registered exactly once.
- [x] [AI] Implement the mismatched-harness fixture in `apps/badakmini-cli/tests/unit/driver_test.go` — acceptance:
      `go -C apps/badakmini-cli test -count=1 ./tests/unit -run TestFeatures` exits 0.
- [x] [AI] Refactor the mismatch support in `apps/badakmini-cli/tests/unit/driver_test.go`; keep the shared
      responsibility-based driver filename and create no scenario-named file — acceptance:
      `go -C apps/badakmini-cli test -count=1 ./tests/unit -run TestFeatures` exits 0.

### A staged rule path announces the workflow

- [x] [AI] Add the scenario to `specs/apps/badakmini-cli/behavior/rule-change.feature` — acceptance:
      `go -C apps/badakmini-cli test -count=1 ./tests/unit -run TestFeatures` exits nonzero because the unit adapter
      lacks the staged-rule fixture.

```gherkin
Scenario: A staged rule path announces the workflow
  Given a repository with a staged rule-bearing file
  When Badak Mini runs staged rule-change detection
  Then the command succeeds with the rules-propagation notice
```

- [x] [AI] Register the staged-rule scenario's missing step expressions in `apps/badakmini-cli/tests/bdd/bindings.go` —
      acceptance: `go -C apps/badakmini-cli test -count=1 ./tests/bdd` exits 0 with each expression registered exactly
      once.
- [x] [AI] Implement the staged-rule fixture in `apps/badakmini-cli/tests/unit/driver_test.go` — acceptance:
      `go -C apps/badakmini-cli test -count=1 ./tests/unit -run TestFeatures` exits 0.
- [x] [AI] Refactor the staged-rule support in `apps/badakmini-cli/tests/unit/driver_test.go`; keep the shared
      responsibility-based driver filename and create no scenario-named file — acceptance:
      `go -C apps/badakmini-cli test -count=1 ./tests/unit -run TestFeatures` exits 0.

### An ordinary staged path stays silent

- [x] [AI] Add the scenario to `specs/apps/badakmini-cli/behavior/rule-change.feature` — acceptance:
      `go -C apps/badakmini-cli test -count=1 ./tests/unit -run TestFeatures` exits nonzero because the unit adapter
      lacks the ordinary-path fixture.

```gherkin
Scenario: An ordinary staged path stays silent
  Given a repository with only an ordinary staged file
  When Badak Mini runs staged rule-change detection
  Then the command succeeds without output
```

- [x] [AI] Register the ordinary-path scenario's missing step expressions in `apps/badakmini-cli/tests/bdd/bindings.go`
      — acceptance: `go -C apps/badakmini-cli test -count=1 ./tests/bdd` exits 0 with each expression registered exactly
      once.
- [x] [AI] Implement the ordinary-path fixture in `apps/badakmini-cli/tests/unit/driver_test.go` — acceptance:
      `go -C apps/badakmini-cli test -count=1 ./tests/unit -run TestFeatures` exits 0.
- [x] [AI] Refactor the ordinary-path support in `apps/badakmini-cli/tests/unit/driver_test.go`; keep the shared
      responsibility-based driver filename and create no scenario-named file — acceptance:
      `go -C apps/badakmini-cli test -count=1 ./tests/unit -run TestFeatures` exits 0.

### A harness edit announces both workflows

- [x] [AI] Add the scenario to `specs/apps/badakmini-cli/behavior/rule-change.feature` — acceptance:
      `go -C apps/badakmini-cli test -count=1 ./tests/unit -run TestFeatures` exits nonzero because the unit adapter
      lacks the hook-payload fixture.

```gherkin
Scenario: A harness edit announces both workflows
  Given a pre-edit payload for a harness instruction file
  When Badak Mini runs hook rule-change detection
  Then the command succeeds with both workflow notices
```

- [x] [AI] Register the hook-payload scenario's missing step expressions in `apps/badakmini-cli/tests/bdd/bindings.go` —
      acceptance: `go -C apps/badakmini-cli test -count=1 ./tests/bdd` exits 0 with each expression registered exactly
      once.
- [x] [AI] Implement the hook-payload fixture in `apps/badakmini-cli/tests/unit/driver_test.go` — acceptance:
      `go -C apps/badakmini-cli test -count=1 ./tests/unit -run TestFeatures` exits 0.
- [x] [AI] Refactor the hook-payload support in `apps/badakmini-cli/tests/unit/driver_test.go`; keep the shared
      responsibility-based driver filename and create no scenario-named file — acceptance:
      `go -C apps/badakmini-cli test -count=1 ./tests/unit -run TestFeatures` exits 0.

- [x] [AI] Add `-count=1` to the corpus-consuming `test:unit` command in `apps/badakmini-cli/project.json` — acceptance:
      `node -e 'const p=require("./apps/badakmini-cli/project.json");if(p.targets["test:unit"].command!=="go -C apps/badakmini-cli test -count=1 ./cmd/... ./internal/... ./tests/bdd ./tests/unit")process.exit(1)'`
      exits 0 and a newly added feature cannot be hidden by Go's test-result cache.
- [x] [AI] Replace `badakmini-cli:test:coverage:unit` in `apps/badakmini-cli/project.json` with the exact command
      `mkdir -p local-tmp && go -C apps/badakmini-cli test -count=1 -coverpkg=./internal/... -coverprofile=../../local-tmp/badakmini-unit.out ./cmd/... ./internal/... ./tests/bdd ./tests/unit && go -C apps/badakmini-cli tool cover -func=../../local-tmp/badakmini-unit.out | awk '/^total:/ { coverage=$3; sub(/%$/, "", coverage); printf "unit statement coverage: %.1f%%\n", coverage; found=1; if (coverage + 0 < 99) exit 1 } END { if (!found) exit 1 }'`
      — acceptance:
      `node -e 'const c=require("./apps/badakmini-cli/project.json").targets["test:coverage:unit"].command,e=Buffer.from("bWtkaXIgLXAgbG9jYWwtdG1wICYmIGdvIC1DIGFwcHMvYmFkYWttaW5pLWNsaSB0ZXN0IC1jb3VudD0xIC1jb3ZlcnBrZz0uL2ludGVybmFsLy4uLiAtY292ZXJwcm9maWxlPS4uLy4uL2xvY2FsLXRtcC9iYWRha21pbmktdW5pdC5vdXQgLi9jbWQvLi4uIC4vaW50ZXJuYWwvLi4uIC4vdGVzdHMvYmRkIC4vdGVzdHMvdW5pdCAmJiBnbyAtQyBhcHBzL2JhZGFrbWluaS1jbGkgdG9vbCBjb3ZlciAtZnVuYz0uLi8uLi9sb2NhbC10bXAvYmFkYWttaW5pLXVuaXQub3V0IHwgYXdrICcvXnRvdGFsOi8geyBjb3ZlcmFnZT0kMzsgc3ViKC8lJC8sICIiLCBjb3ZlcmFnZSk7IHByaW50ZiAidW5pdCBzdGF0ZW1lbnQgY292ZXJhZ2U6ICUuMWYlJVxuIiwgY292ZXJhZ2U7IGZvdW5kPTE7IGlmIChjb3ZlcmFnZSArIDAgPCA5OSkgZXhpdCAxIH0gRU5EIHsgaWYgKCFmb3VuZCkgZXhpdCAxIH0n","base64").toString();if(c!==e)process.exit(1)'`
      exits 0, structurally proving the complete command retains `-count=1`, the `./internal/...` denominator, the unit
      profile, all four package selections, and the 99% `awk` threshold.
- [x] [AI] Run `npm exec nx -- run badakmini-cli:test:coverage:unit` after the exact command is installed — acceptance:
      the uncached corpus-consuming unit suite executes and exits 0 at or above 99%.
- [x] [AI] Add an uncached `test:integration` target to `apps/badakmini-cli/project.json` with command
      `go -C apps/badakmini-cli test -count=1 ./tests/integration` and inputs `default` plus
      `{workspaceRoot}/specs/apps/badakmini-cli/behavior/**/*.feature` — acceptance:
      `node -e 'const p=require("./apps/badakmini-cli/project.json");const t=p.targets["test:integration"];const i=["default","{workspaceRoot}/specs/apps/badakmini-cli/behavior/**/*.feature"];if(!t||t.cache!==false||t.command!=="go -C apps/badakmini-cli test -count=1 ./tests/integration"||JSON.stringify(t.inputs)!==JSON.stringify(i))process.exit(1)'`
      exits 0 before the target is invoked elsewhere.
- [x] [AI] Add `default` plus `{workspaceRoot}/specs/apps/badakmini-cli/behavior/**/*.feature` inputs to
      `badakmini-cli:typecheck` in `apps/badakmini-cli/project.json` — acceptance:
      `npm exec nx -- show project badakmini-cli --json` shows both inputs on `typecheck`.
- [x] [AI] Add `default` plus `{workspaceRoot}/specs/apps/badakmini-cli/behavior/**/*.feature` inputs to
      `badakmini-cli:test:unit` in `apps/badakmini-cli/project.json` — acceptance:
      `npm exec nx -- show project badakmini-cli --json` shows both inputs on `test:unit`.
- [x] [AI] Implement the canonical-scenario integration driver in `apps/badakmini-cli/tests/integration/driver_test.go`
      using isolated local filesystem and Git fixtures with no network — acceptance:
      `go -C apps/badakmini-cli test -count=1 ./tests/integration -run '^$'` compiles the driver.
- [x] [AI] Implement the recursive canonical runner in `apps/badakmini-cli/tests/integration/features_test.go` —
      acceptance: `npm exec nx -- run badakmini-cli:test:integration` executes every canonical scenario through the
      integration driver and exits 0.
- [x] [AI] Create `apps/badakmini-cli-e2e/go.mod` with module path
      `github.com/wahidyankf/grind-in-public/apps/badakmini-cli-e2e`, Go `1.26.6`, direct requirements on Godog
      `v0.16.0` and the owner module at `v0.0.0`,
      `replace github.com/wahidyankf/grind-in-public/apps/badakmini-cli => ../badakmini-cli`, and the same GolangCI-Lint
      tool directive as the owner — acceptance:
      `node -e 'const s=require("fs").readFileSync("apps/badakmini-cli-e2e/go.mod","utf8");for(const x of ["github.com/cucumber/godog v0.16.0","github.com/wahidyankf/grind-in-public/apps/badakmini-cli v0.0.0","replace github.com/wahidyankf/grind-in-public/apps/badakmini-cli => ../badakmini-cli","github.com/golangci/golangci-lint/v2/cmd/golangci-lint"])if(!s.includes(x))process.exit(1)'`
      exits 0.
- [x] [AI] Implement the process driver in `apps/badakmini-cli-e2e/tests/driver_test.go`; require `BADAKMINI_BIN`,
      reject it with a clear error when it is unset, relative, missing, or non-executable, launch only that absolute
      executable path, observe only arguments, exit status, stdout, stderr, and filesystem effects, and enforce a
      30-second timeout — acceptance: source review confirms those contract branches exist. No Go command runs yet
      because the E2E module remains dependency-incomplete until `go mod tidy` after all imports exist.
- [x] [AI] Add the exact `TestProcessDriverBinaryContract` test to `apps/badakmini-cli-e2e/tests/driver_test.go` for the
      required `BADAKMINI_BIN` validation and absolute-path execution behavior — acceptance:
      `rg -n '^func TestProcessDriverBinaryContract\(t \*testing\.T\)' apps/badakmini-cli-e2e/tests/driver_test.go`
      finds the exact declaration. No filtered Go command runs yet because the E2E module remains dependency-incomplete
      until `go mod tidy` after all imports exist.
- [x] [AI] Add the exact `TestE2EBoundaryPolicy` check to `apps/badakmini-cli/tests/bdd/boundary_policy_test.go` after
      the process driver exists — acceptance:
      `set -o pipefail; go -C apps/badakmini-cli test -count=1 -json ./tests/bdd -run '^TestE2EBoundaryPolicy$' | rg '"Action":"pass".*"Test":"TestE2EBoundaryPolicy"'`
      exits 0 after proving E2E observes only the public process boundary.
- [x] [AI] Implement the canonical E2E Godog runner in `apps/badakmini-cli-e2e/tests/features_test.go` using the owner
      BDD support through the local module requirement — acceptance: the source imports resolve to the requirements
      declared in `go.mod`; no Go command runs until `mod tidy` resolves those imports after all sources exist.
- [x] [AI] Create `apps/badakmini-cli-e2e/.golangci.yml` from the owner's strict lint configuration with an
      `e2e-test-support` Depguard rule over `$all` whose allowlist is exactly `$gostd`,
      `github.com/wahidyankf/grind-in-public/apps/badakmini-cli`, and `github.com/cucumber/godog` — acceptance:
      `node -e 'const s=require("fs").readFileSync("apps/badakmini-cli-e2e/.golangci.yml","utf8");const r=/e2e-test-support:[\s\S]*?files:[\s\S]*?- \$all[\s\S]*?allow:[\s\S]*?- \$gostd[\s\S]*?- github\.com\/wahidyankf\/grind-in-public\/apps\/badakmini-cli[\s\S]*?- github\.com\/cucumber\/godog/;if(!r.test(s))process.exit(1)'`
      exits 0; no lint command runs until module tooling is resolved by the next action.
- [x] [AI] Generate `apps/badakmini-cli-e2e/go.sum` after all E2E imports exist by running
      `go -C apps/badakmini-cli-e2e mod tidy` — acceptance: repeating `go -C apps/badakmini-cli-e2e mod tidy` leaves
      `go.mod` and `go.sum` unchanged.
- [x] [AI] Run the process-driver binary-contract test after module resolution — acceptance:
      `set -o pipefail; go -C apps/badakmini-cli-e2e test -count=1 -json ./tests -run '^TestProcessDriverBinaryContract$' | rg '"Action":"pass".*"Test":"TestProcessDriverBinaryContract"'`
      exits 0, so the filter cannot pass when the required test is absent.
- [x] [AI] Run `go -C apps/badakmini-cli-e2e test -count=1 ./tests -run '^$'` — acceptance: the dependency-complete E2E
      adapter compiles without launching the CLI.
- [x] [AI] Run `go -C apps/badakmini-cli-e2e tool golangci-lint run` — acceptance: the E2E module's exact test tooling
      exits 0.
- [x] [AI] Create `apps/badakmini-cli-e2e/README.md` describing its owner corpus, public-process boundary, targets, and
      omitted unit and numeric coverage gates — acceptance: `npm run check:markdown-links` exits 0.
- [x] [AI] Add `test:coverage:behavior` to `apps/badakmini-cli/project.json` with command
      `go -C apps/badakmini-cli test -count=1 ./tests/bdd`, `default`, the recursive corpus input, and every E2E
      test/configuration input from `tech-docs/README.md` — acceptance:
      `node -e 'const p=require("./apps/badakmini-cli/project.json");const t=p.targets["test:coverage:behavior"];const required=["default","{workspaceRoot}/specs/apps/badakmini-cli/behavior/**/*.feature","{workspaceRoot}/apps/badakmini-cli-e2e/tests/**/*","{workspaceRoot}/apps/badakmini-cli-e2e/go.mod","{workspaceRoot}/apps/badakmini-cli-e2e/go.sum","{workspaceRoot}/apps/badakmini-cli-e2e/.golangci.yml","{workspaceRoot}/apps/badakmini-cli-e2e/project.json"];if(!t||t.command!=="go -C apps/badakmini-cli test -count=1 ./tests/bdd"||required.some(x=>!t.inputs.includes(x)))process.exit(1)'`
      exits 0 before any E2E target delegates to it.
- [x] [AI] Create the Nx application scaffold in `apps/badakmini-cli-e2e/project.json` with name and root
      `badakmini-cli-e2e`, `projectType: application`, and `implicitDependencies: ["badakmini-cli"]` — acceptance:
      `node -e 'const p=require("./apps/badakmini-cli-e2e/project.json");if(p.name!=="badakmini-cli-e2e"||p.root!=="apps/badakmini-cli-e2e"||p.projectType!=="application"||JSON.stringify(p.implicitDependencies)!==JSON.stringify(["badakmini-cli"]))process.exit(1)'`
      exits 0.
- [x] [AI] Add `badakmini-cli-e2e:typecheck` with command `go -C apps/badakmini-cli-e2e vet ./...` to
      `apps/badakmini-cli-e2e/project.json` — acceptance:
      `node -e 'const p=require("./apps/badakmini-cli-e2e/project.json");if(p.targets.typecheck.command!=="go -C apps/badakmini-cli-e2e vet ./...")process.exit(1)'`
      exits 0.
- [x] [AI] Add `badakmini-cli-e2e:lint` with command `go -C apps/badakmini-cli-e2e tool golangci-lint run` to
      `apps/badakmini-cli-e2e/project.json` — acceptance:
      `node -e 'const p=require("./apps/badakmini-cli-e2e/project.json");if(p.targets.lint.command!=="go -C apps/badakmini-cli-e2e tool golangci-lint run")process.exit(1)'`
      exits 0.
- [x] [AI] Add `badakmini-cli-e2e:test:coverage:behavior` with command
      `npm exec nx -- run badakmini-cli:test:coverage:behavior` to `apps/badakmini-cli-e2e/project.json` — acceptance:
      `node -e 'const p=require("./apps/badakmini-cli-e2e/project.json");if(p.targets["test:coverage:behavior"].command!=="npm exec nx -- run badakmini-cli:test:coverage:behavior")process.exit(1)'`
      exits 0.
- [x] [AI] Add `badakmini-cli-e2e:test:e2e` with command
      `BADAKMINI_BIN="$PWD/apps/badakmini-cli/dist/badak-mini" go -C apps/badakmini-cli-e2e test -count=1 ./tests` and
      `dependsOn: ["^build", "test:coverage:behavior"]` to `apps/badakmini-cli-e2e/project.json` — acceptance:
      `node -e 'const p=require("./apps/badakmini-cli-e2e/project.json"),t=p.targets["test:e2e"],c="BADAKMINI_BIN=\"$PWD/apps/badakmini-cli/dist/badak-mini\" go -C apps/badakmini-cli-e2e test -count=1 ./tests";if(t.command!==c||JSON.stringify(t.dependsOn)!==JSON.stringify(["^build","test:coverage:behavior"]))process.exit(1)'`
      exits 0.
- [x] [AI] Verify the E2E owner-build and binary-path contract after the target exists — acceptance:
      `npm exec nx -- run badakmini-cli:build && test -x "$PWD/apps/badakmini-cli/dist/badak-mini" && npm exec nx -- run badakmini-cli-e2e:test:e2e`
      exits 0, proving the owner build creates the executable at the absolute path that the E2E target passes through
      `BADAKMINI_BIN` and the driver consumes.
- [x] [AI] Add ordered `badakmini-cli-e2e:test:quick` to `apps/badakmini-cli-e2e/project.json` with owner
      `badakmini-cli:test:quick` dependency, commands `npm exec nx -- run badakmini-cli-e2e:typecheck` then
      `npm exec nx -- run badakmini-cli-e2e:lint`, and `parallel: false` — acceptance:
      `node -e 'const p=require("./apps/badakmini-cli-e2e/project.json"),q=p.targets["test:quick"],c=["npm exec nx -- run badakmini-cli-e2e:typecheck","npm exec nx -- run badakmini-cli-e2e:lint"];if(JSON.stringify(q.options.commands)!==JSON.stringify(c)||q.options.parallel!==false||q.dependsOn[0].projects[0]!=="badakmini-cli"||q.dependsOn[0].target!=="test:quick")process.exit(1)'`
      exits 0.
- [x] [AI] Give `badakmini-cli-e2e:typecheck` the `default`, `^default`, and recursive corpus inputs in
      `apps/badakmini-cli-e2e/project.json` — acceptance: `npm exec nx -- show project badakmini-cli-e2e --json` shows
      all three inputs on `typecheck`.
- [x] [AI] Give `badakmini-cli-e2e:lint` the `default`, `^default`, and recursive corpus inputs in
      `apps/badakmini-cli-e2e/project.json` — acceptance: `npm exec nx -- show project badakmini-cli-e2e --json` shows
      all three inputs on `lint`.
- [x] [AI] Give `badakmini-cli-e2e:test:coverage:behavior` the `default`, `^default`, and recursive corpus inputs in
      `apps/badakmini-cli-e2e/project.json` — acceptance: `npm exec nx -- show project badakmini-cli-e2e --json` shows
      all three inputs on `test:coverage:behavior`.
- [x] [AI] Give `badakmini-cli-e2e:test:e2e` the `default`, `^default`, and recursive corpus inputs in
      `apps/badakmini-cli-e2e/project.json` — acceptance: `npm exec nx -- show project badakmini-cli-e2e --json` shows
      all three inputs on `test:e2e`.
- [x] [AI] Give `badakmini-cli-e2e:test:quick` the `default`, `^default`, and recursive corpus inputs in
      `apps/badakmini-cli-e2e/project.json` — acceptance: `npm exec nx -- show project badakmini-cli-e2e --json` shows
      all three inputs on `test:quick`.
- [x] [AI] Complete the static unit, integration, and E2E adapter comparison in
      `apps/badakmini-cli/tests/bdd/adapter_parity_test.go` — acceptance:
      `npm exec nx -- run badakmini-cli:test:coverage:behavior` reports matching feature, expanded-scenario, step, and
      binding counts for all three adapters.
- [x] [AI] Run `npm exec nx -- run badakmini-cli-e2e:test:coverage:behavior` — acceptance: the fully defined E2E target
      delegates to the already-green owner behavior gate and exits 0.
- [x] [AI] Run `npm exec nx -- run badakmini-cli-e2e:test:e2e` — acceptance: the fully defined E2E target builds Badak
      Mini, passes behavior coverage first, executes every canonical scenario through the binary, and exits 0.
- [x] [AI] Add an added-feature fixture to `apps/badakmini-cli/tests/bdd/adapter_parity_test.go` — acceptance:
      `go -C apps/badakmini-cli test -count=1 ./tests/bdd` exits 0 after proving a new feature invalidates the catalog
      and fails until every adapter resolves it.
- [x] [AI] Add an edited-step fixture to `apps/badakmini-cli/tests/bdd/adapter_parity_test.go` — acceptance:
      `go -C apps/badakmini-cli test -count=1 ./tests/bdd` exits 0 after proving an edited step fails until every
      adapter resolves it.
- [x] [AI] Add a renamed-feature fixture to `apps/badakmini-cli/tests/bdd/adapter_parity_test.go` — acceptance:
      `go -C apps/badakmini-cli test -count=1 ./tests/bdd` exits 0 after proving a renamed feature remains recursively
      discovered.
- [x] [AI] Add a nested-feature fixture to `apps/badakmini-cli/tests/bdd/adapter_parity_test.go` — acceptance:
      `go -C apps/badakmini-cli test -count=1 ./tests/bdd` exits 0 after proving nesting needs no manual registration.
- [x] [AI] Add a deleted-feature fixture to `apps/badakmini-cli/tests/bdd/adapter_parity_test.go` — acceptance:
      `go -C apps/badakmini-cli test -count=1 ./tests/bdd` exits 0 after proving stale bindings fail as unused.
- [x] [AI] Add an added-binding fixture to `apps/badakmini-cli/tests/bdd/binding_compliance_test.go` — acceptance:
      `go -C apps/badakmini-cli test -count=1 ./tests/bdd` exits 0 after proving the binding is rechecked and fails when
      unused.
- [x] [AI] Add an edited-binding fixture to `apps/badakmini-cli/tests/bdd/binding_compliance_test.go` — acceptance:
      `go -C apps/badakmini-cli test -count=1 ./tests/bdd` exits 0 after proving the changed expression is rechecked
      against every step.
- [x] [AI] Add a renamed-binding fixture to `apps/badakmini-cli/tests/bdd/binding_compliance_test.go` — acceptance:
      `go -C apps/badakmini-cli test -count=1 ./tests/bdd` exits 0 after proving the old expression becomes undefined
      unless every adapter changes.
- [x] [AI] Add a deleted-binding fixture to `apps/badakmini-cli/tests/bdd/binding_compliance_test.go` — acceptance:
      `go -C apps/badakmini-cli test -count=1 ./tests/bdd` exits 0 after proving its former step becomes undefined.
- [x] [AI] Add an E2E-binding input regression to `apps/badakmini-cli/tests/bdd/adapter_parity_test.go` — acceptance:
      `go -C apps/badakmini-cli test -count=1 ./tests/bdd` exits 0 after proving an E2E binding change invalidates
      `badakmini-cli:test:coverage:behavior`.
- [x] [AI] Add an E2E-configuration input regression to `apps/badakmini-cli/tests/bdd/adapter_parity_test.go` —
      acceptance: `go -C apps/badakmini-cli test -count=1 ./tests/bdd` exits 0 after proving an E2E module, lint, or
      project configuration change invalidates `badakmini-cli:test:coverage:behavior`.

### Phase 2 Gate

> Every check below passes before Phase 3 begins. A failure is fixed inside Phase 2.

- [x] [AI] Run `npm exec nx -- run badakmini-cli:test:unit` — acceptance: every unit scenario and focused unit test
      exits 0.
- [x] [AI] Run `npm exec nx -- run badakmini-cli:test:integration` — acceptance: every local-only integration scenario
      exits 0.
- [x] [AI] Run `npm exec nx -- run badakmini-cli-e2e:test:e2e` — acceptance: every process E2E scenario exits 0.
- [x] [AI] Run `npm exec nx -- run badakmini-cli:test:coverage:behavior` — acceptance: exact corpus, binding, driver,
      and adapter checks exit 0.

> **Pause Safety**: All canonical scenarios execute through all three required adapters, and static behavior
> completeness is green. Safe to stop. Resume with `npm exec nx -- run badakmini-cli:test:coverage:behavior`.

## Phase 3: Coverage, Automation, and Project Gates

- [x] [AI] Add `test:coverage:integration` to `apps/badakmini-cli/project.json` with exact command
      `mkdir -p local-tmp && go -C apps/badakmini-cli test -count=1 -coverpkg=./internal/cli -coverprofile=../../local-tmp/badakmini-integration.out ./tests/integration && go -C apps/badakmini-cli tool cover -func=../../local-tmp/badakmini-integration.out | awk '/^total:/ { coverage=$3; sub(/%$/, "", coverage); printf "integration statement coverage: %.1f%%\n", coverage; found=1; if (coverage + 0 < 99) exit 1 } END { if (!found) exit 1 }'`
      — acceptance:
      `node -e 'const c=require("./apps/badakmini-cli/project.json").targets["test:coverage:integration"].command,e=Buffer.from("bWtkaXIgLXAgbG9jYWwtdG1wICYmIGdvIC1DIGFwcHMvYmFkYWttaW5pLWNsaSB0ZXN0IC1jb3VudD0xIC1jb3ZlcnBrZz0uL2ludGVybmFsL2NsaSAtY292ZXJwcm9maWxlPS4uLy4uL2xvY2FsLXRtcC9iYWRha21pbmktaW50ZWdyYXRpb24ub3V0IC4vdGVzdHMvaW50ZWdyYXRpb24gJiYgZ28gLUMgYXBwcy9iYWRha21pbmktY2xpIHRvb2wgY292ZXIgLWZ1bmM9Li4vLi4vbG9jYWwtdG1wL2JhZGFrbWluaS1pbnRlZ3JhdGlvbi5vdXQgfCBhd2sgJy9edG90YWw6LyB7IGNvdmVyYWdlPSQzOyBzdWIoLyUkLywgIiIsIGNvdmVyYWdlKTsgcHJpbnRmICJpbnRlZ3JhdGlvbiBzdGF0ZW1lbnQgY292ZXJhZ2U6ICUuMWYlJVxuIiwgY292ZXJhZ2U7IGZvdW5kPTE7IGlmIChjb3ZlcmFnZSArIDAgPCA5OSkgZXhpdCAxIH0gRU5EIHsgaWYgKCFmb3VuZCkgZXhpdCAxIH0n","base64").toString();if(c!==e)process.exit(1)'`
      exits 0, structurally proving the complete uncached integration command, `internal/cli` denominator, profile,
      package selection, and 99% threshold.
- [x] [AI] Run `npm exec nx -- run badakmini-cli:test:coverage:integration` after the exact command is installed —
      acceptance: the uncached corpus-consuming integration suite aggregates the profile's `internal/cli` statement
      total and exits 0 only at or above 99%.
- [x] [AI] Add the owner `default` plus recursive corpus inputs to `badakmini-cli:test:coverage:integration` in
      `apps/badakmini-cli/project.json` — acceptance: `npm exec nx -- show project badakmini-cli --json` shows both
      inputs on `test:coverage:integration`.
- [x] [AI] Extend integration fixture behavior in `apps/badakmini-cli/tests/integration/driver_test.go` for uncovered
      `internal/cli` application-boundary branches — acceptance:
      `npm exec nx -- run badakmini-cli:test:coverage:integration` exits 0 at or above 99% without broad exclusions.
- [x] [AI] Document coverage ownership in `apps/badakmini-cli/README.md`: unit owns every `internal/...` validator and
      orchestration package; integration owns only `internal/cli`; `cmd/badak-mini` is a thin process entrypoint owned
      by E2E and is outside both numeric denominators; `tests/unit` is the Gherkin unit adapter — acceptance:
      `npm run check:markdown-links` exits 0 and the README names both exact 99% targets.
- [x] [AI] Compose the owner app's aggregate `test:coverage` target in `apps/badakmini-cli/project.json` — acceptance:
      `npm exec nx -- run badakmini-cli:test:coverage` runs unit, integration, then behavior coverage and exits 0.
- [x] [AI] Add the owner corpus plus E2E binding/configuration inputs specified in `tech-docs/README.md` to
      `badakmini-cli:test:coverage` — acceptance: `npm exec nx -- show project badakmini-cli --json` shows every
      specified input on `test:coverage`.
- [x] [AI] Compose the owner app's ordered `test:quick` target in `apps/badakmini-cli/project.json` — acceptance:
      `npm exec nx -- run badakmini-cli:test:quick` runs typecheck, lint, unit, unit coverage, then behavior coverage
      without integration or process E2E.
- [x] [AI] Add the owner corpus plus E2E binding/configuration inputs specified in `tech-docs/README.md` to
      `badakmini-cli:test:quick` — acceptance: `npm exec nx -- show project badakmini-cli --json` shows every specified
      input on `test:quick`.
- [x] [AI] Compose the E2E harness's `test:quick` target in `apps/badakmini-cli-e2e/project.json` — acceptance:
      `npm exec nx -- run badakmini-cli-e2e:test:quick` depends on the owner quick gate, then runs only E2E typecheck
      and lint.
- [x] [AI] Add `"test:behavior": "nx run-many -t test:coverage:behavior"` to `package.json` — acceptance:
      `node -e 'const p=require("./package.json");if(p.scripts["test:behavior"]!=="nx run-many -t test:coverage:behavior")process.exit(1)'`
      exits 0.
- [x] [AI] Add `"test:e2e": "nx run-many -t test:e2e"` to `package.json` — acceptance:
      `node -e 'const p=require("./package.json");if(p.scripts["test:e2e"]!=="nx run-many -t test:e2e")process.exit(1)'`
      exits 0.
- [x] [AI] Add the canonical
      `"test:scheduled": "nx run badakmini-cli:test:quick && nx run badakmini-cli:test:coverage:integration && nx run badakmini-cli-e2e:test:e2e"`
      script to `package.json` — acceptance:
      `node -e 'const p=require("./package.json");if(p.scripts["test:scheduled"]!=="nx run badakmini-cli:test:quick && nx run badakmini-cli:test:coverage:integration && nx run badakmini-cli-e2e:test:e2e")process.exit(1)'`
      exits 0 and fixes quick verification first, then integration coverage, then E2E in one reusable command.
- [x] [AI] Add `"check:workflows": "go -C apps/badakmini-cli tool actionlint \"$PWD\"/.github/workflows/*.yml"` to
      `package.json` — acceptance:
      `node -e 'const p=require("./package.json");if(p.scripts["check:workflows"]!=="go -C apps/badakmini-cli tool actionlint \"$PWD\"/.github/workflows/*.yml")process.exit(1)'`
      exits 0; the absolute glob is expanded from the workspace root before Go changes into the owner module.
- [x] [AI] Create `.github/workflows/full-bdd.yml` with `schedule` cron `0 23 * * *` for 06:00 WIB — acceptance:
      `node -e 'const s=require("fs").readFileSync(".github/workflows/full-bdd.yml","utf8");if(!s.includes("cron: \"0 23 * * *\""))process.exit(1)'`
      exits 0.
- [x] [AI] Add `workflow_dispatch` to `.github/workflows/full-bdd.yml` — acceptance:
      `node -e 'const s=require("fs").readFileSync(".github/workflows/full-bdd.yml","utf8");if(!/^  workflow_dispatch:/m.test(s))process.exit(1)'`
      exits 0.
- [x] [AI] Configure `.github/workflows/full-bdd.yml` with `contents: read` — acceptance:
      `node -e 'const s=require("fs").readFileSync(".github/workflows/full-bdd.yml","utf8");if(!/^permissions:\s*\n  contents: read$/m.test(s))process.exit(1)'`
      exits 0.
- [x] [AI] Keep `.github/workflows/full-bdd.yml` free of secret references — acceptance:
      `node -e 'const s=require("fs").readFileSync(".github/workflows/full-bdd.yml","utf8");if(s.includes("secrets."))process.exit(1)'`
      exits 0.
- [x] [AI] Add the verification job to `.github/workflows/full-bdd.yml` with `runs-on: ubuntu-latest` — acceptance:
      `node -e 'const s=require("fs").readFileSync(".github/workflows/full-bdd.yml","utf8");if(!/^jobs:\s*\n  [a-z0-9-]+:[\s\S]*?^    runs-on: ubuntu-latest$/m.test(s))process.exit(1)'`
      exits 0.
- [x] [AI] Add the checkout step with `uses: actions/checkout@v6` to `.github/workflows/full-bdd.yml` — acceptance:
      `node -e 'const s=require("fs").readFileSync(".github/workflows/full-bdd.yml","utf8");if(!/- name: Check out repository\s*\n\s*uses: actions\/checkout@v6/.test(s))process.exit(1)'`
      exits 0.
- [x] [AI] Add the Node setup step with `uses: actions/setup-node@v6`, `node-version: "24.15.0"`, and `cache: npm` to
      `.github/workflows/full-bdd.yml` — acceptance:
      `node -e 'const s=require("fs").readFileSync(".github/workflows/full-bdd.yml","utf8");if(!/- name: Set up Node\.js\s*\n\s*uses: actions\/setup-node@v6\s*\n\s*with:\s*\n\s*node-version: "24\.15\.0"\s*\n\s*cache: npm/.test(s))process.exit(1)'`
      exits 0.
- [x] [AI] Add the Go setup step with `uses: actions/setup-go@v6`, `go-version: "1.26.6"`, `cache: true`, and
      `cache-dependency-path: apps/*/go.sum` to `.github/workflows/full-bdd.yml` — acceptance:
      `node -e 'const s=require("fs").readFileSync(".github/workflows/full-bdd.yml","utf8");if(!/- name: Set up Go\s*\n\s*uses: actions\/setup-go@v6\s*\n\s*with:\s*\n\s*go-version: "1\.26\.6"\s*\n\s*cache: true\s*\n\s*cache-dependency-path: apps\/\*\/go\.sum/.test(s))process.exit(1)'`
      exits 0.
- [x] [AI] Add the locked-dependency step `npm ci` to `.github/workflows/full-bdd.yml` — acceptance:
      `node -e 'const s=require("fs").readFileSync(".github/workflows/full-bdd.yml","utf8");if(!/- name: Install npm dependencies\s*\n\s*run: npm ci/.test(s))process.exit(1)'`
      exits 0.
- [x] [AI] Add the canonical scheduled verification step `npm run test:scheduled` to `.github/workflows/full-bdd.yml`
      after `npm ci` — acceptance:
      `node -e 'const s=require("fs").readFileSync(".github/workflows/full-bdd.yml","utf8"),a=s.search(/- name: Install npm dependencies\s*\n\s*run: npm ci/),b=s.search(/- name: Run canonical scheduled verification\s*\n\s*run: npm run test:scheduled/);if(a<0||b<=a)process.exit(1)'`
      exits 0; the reused root script runs quick first while preserving integration coverage before E2E.

### Phase 3 Gate

> Every check below passes before Phase 4 begins. A failure is fixed inside Phase 3.

- [x] [AI] Run `npm run typecheck` — acceptance: every applicable project typecheck exits 0.
- [x] [AI] Run `npm run lint` — acceptance: every applicable project lint exits 0.
- [x] [AI] Run `npm run test:quick` — acceptance: all fast project gates exit 0.
- [x] [AI] Run `npm run test:integration` — acceptance: all applicable local integration suites exit 0.
- [x] [AI] Run `npm exec nx -- run badakmini-cli:test:coverage:integration` — acceptance: the scheduled local-boundary
      coverage slice exits 0 at or above 99%.
- [x] [AI] Run `npm run test:e2e` — acceptance: all dedicated E2E suites exit 0.
- [x] [AI] Run `npm run test:scheduled` — acceptance: the canonical operational suite runs quick verification, including
      unit 99% coverage and behavior coverage, then integration 99% coverage, then process E2E in that exact order and
      exits 0.
- [x] [AI] Run `npm exec -- prettier --check .github/workflows/full-bdd.yml` — acceptance: Prettier parses the workflow
      as YAML and reports it formatted.
- [x] [AI] Run `npm run check:workflows` — acceptance: exact-pinned actionlint `v1.7.12` reports no GitHub Actions
      syntax, schema, expression, event, job, or step error in `.github/workflows/full-bdd.yml`.

> **Pause Safety**: Layered tests, both 99% coverage slices, affected inputs, workflow schema, behavior coverage, and
> process E2E are operational. Safe to stop. Resume with `npm run check:workflows && npm run test:scheduled` to reverify
> the authoritative workflow check followed by quick verification, integration coverage, and E2E.

## Phase 4: Compulsory Repository Rules

- [x] [AI] Run the inventory step in `repo-governance/workflows/rules/rules-propagation.md` for the proposed role-based
      BDD rule — acceptance: current authority, audience, enforcement, harness, and index surfaces are identified before
      any rule edit. No shell command applies because this is the workflow's required evidence review.
- [x] [AI] Run the idempotency gate in `repo-governance/workflows/rules/rules-propagation.md` — acceptance: at least one
      objective clarity criterion fails before propagation proceeds. No shell command applies because this is the
      workflow's semantic decision gate.
- [x] [AI] Record the failed idempotency criteria and cited sources in
      `plans/in-progress/badakmini-layered-bdd/learnings.md` — acceptance: the entry names the missing automated
      verification. No separate command applies because the entry records the preceding workflow evidence.
- [x] [AI] Add `repo-governance/development/behavior-driven-development-policy.md` as the canonical role-based BDD
      policy — acceptance: it requires apps to use unit plus local-only integration plus dedicated-app E2E, libraries to
      use unit plus integration only for an owned local boundary and never E2E, and dedicated E2E apps to consume the
      owner's corpus while omitting meaningless unit and numeric coverage targets. No standalone semantic command
      applies; the strict rules quality gate verifies the new rule and its propagation.
- [x] [AI] Update `repo-governance/development/specs-policy.md` to require executable Gherkin for every app and library
      — acceptance:
      `rg -n 'Existing projects gain specs|no automated check' repo-governance/development/specs-policy.md` finds no
      obsolete deferred-retrofit or no-automation statement.
- [x] [AI] Update `repo-governance/development/testing-policy.md` to reference the canonical role matrix — acceptance:
      `npm run check:markdown-links` exits 0.
- [x] [AI] Update `repo-governance/development/tdd-policy.md` to require Gherkin-first behavior cycles — acceptance:
      `npm run check:markdown-links` exits 0.
- [x] [AI] Update `repo-governance/development/badakmini-cli-policy.md` with the implemented Badak targets — acceptance:
      `npm run check:markdown-links` exits 0.
- [x] [AI] Document exact-pinned actionlint ownership and the root `npm run check:workflows` command in
      `apps/badakmini-cli/README.md` — acceptance: `npm run check:markdown-links` exits 0 and
      `rg -n 'actionlint|check:workflows' apps/badakmini-cli/README.md` finds both command ownership terms.
- [x] [AI] Update `repo-governance/development/workspace-commands.md` with the implemented root scripts and project
      targets — acceptance:
      `node -e 'const p=require("./package.json"),o=require("./apps/badakmini-cli/project.json"),e=require("./apps/badakmini-cli-e2e/project.json");const rs=["check:workflows","test:behavior","test:e2e","test:integration","test:quick","test:scheduled"],ot=["test:unit","test:integration","test:coverage:unit","test:coverage:integration","test:coverage:behavior","test:coverage","test:quick"],et=["typecheck","lint","test:coverage:behavior","test:e2e","test:quick"];if(rs.some(x=>!p.scripts[x])||ot.some(x=>!o.targets[x])||et.some(x=>!e.targets[x]))process.exit(1)'`
      exits 0, proving root commands against `package.json` and project commands against their `project.json` sources.
- [x] [AI] Link the canonical BDD policy from `AGENTS.md` — acceptance: `npm run check:markdown-links` exits 0.
- [x] [AI] Index the canonical BDD policy in `repo-governance/development/README.md` — acceptance:
      `npm run check:markdown-links` exits 0.
- [x] [AI] Update `specs/README.md` with the now-populated corpus and automated enforcement — acceptance:
      `npm run check:markdown-links` exits 0.
- [x] [AI] Update `specs/apps/badakmini-cli/README.md` with the corpus index — acceptance:
      `npm run check:markdown-links` exits 0.
- [x] [AI] Update `apps/badakmini-cli/README.md` with its corpus, adapters, and commands — acceptance:
      `npm run check:markdown-links` exits 0.
- [x] [AI] Update `apps/badakmini-cli-e2e/README.md` with its owner relationship and applicable commands — acceptance:
      `npm run check:markdown-links` exits 0.
- [x] [AI] Run `repo-governance/workflows/harness-alignment.md` over every changed instruction, capability,
      configuration, path, and index — acceptance: derivative instructions defer to `AGENTS.md`, quoted paths exist, and
      quoted commands are defined. No single shell command replaces the workflow's inventory and propagation review.
- [x] [AI] Run `npm run check:harness-parity` after Harness Alignment — acceptance: supported harness capabilities
      match.
- [x] [AI] Run `repo-governance/workflows/rules-quality-gate.md` in strict mode with checker/fixer separation —
      acceptance: two consecutive rules-checker runs report no CRITICAL, HIGH, or MEDIUM findings and the untracked gate
      history records a pass. No shell command applies because the gate invokes repository subagents.

### Phase 4 Gate

> Every check below passes before Phase 5 begins. A failure is fixed inside Phase 4.

- [x] [AI] Run `npm run format:check` — acceptance: Prettier reports every file formatted.
- [x] [AI] Run `npm run check:governance` — acceptance: governed documents remain within the word limit.
- [x] [AI] Run `npm run check:harness-parity` — acceptance: supported harness capabilities match.
- [x] [AI] Run `npm run check:markdown-links` — acceptance: every repository-local Markdown link resolves.

> **Pause Safety**: The implemented BDD model is now compulsory and all governance and harness surfaces agree. Safe to
> stop. Resume with `npm run check:governance`.

## Phase 5: Knowledge Capture and Final Verification

- [x] [AI] Review `plans/in-progress/badakmini-layered-bdd/learnings.md` for secrets — acceptance: every entry is marked
      safe or discarded with a reason. No shell command can establish semantic secret safety; inspect every entry
      manually.
- [x] [AI] Review each safe entry in `plans/in-progress/badakmini-layered-bdd/learnings.md` for repository-wide
      relevance — acceptance: every remaining entry is marked generalizable or local with a reason. No shell command can
      classify semantic relevance; inspect every safe entry manually.
- [x] [AI] Route each surviving learning to one durable rule, document, test, code change, or idea — acceptance: every
      entry names exactly one terminal destination, or the file records `No generalizable learnings` with a reason. No
      shell command can choose a semantic destination; inspect the recorded rationale manually.
- [x] [AI] Run `npm audit` — acceptance: npm reports no known vulnerability in the installed dependency tree.
- [x] [AI] Run `npm run check:go-vulnerabilities` — acceptance: govulncheck reports no reachable Go vulnerability.
- [x] [AI] Run `npm run test:quick` — acceptance: the final quick gate exits 0.
- [x] [AI] Run `npm run test:integration` — acceptance: the final local-integration gate exits 0.
- [x] [AI] Run `npm run test:e2e` — acceptance: the final E2E gate exits 0.
- [x] [AI] Run `npm run test:scheduled` — acceptance: after knowledge capture, the final scheduled gate reruns quick
      verification with unit 99% and behavior coverage, then integration 99% coverage, then process E2E and exits 0.
- [x] [AI] Run `npm run check:workflows` — acceptance: the final authoritative GitHub Actions structural and schema
      check exits 0.
- [x] [AI] Run `npm run format:check` — acceptance: the final format gate exits 0.
- [x] [AI] Run `npm run check:governance` — acceptance: the final governance gate exits 0.
- [x] [AI] Run `npm run check:harness-parity` — acceptance: the final harness-parity gate exits 0.
- [x] [AI] Run `npm run check:markdown-links` — acceptance: the final Markdown-link gate exits 0.

### Phase 5 Gate

> Every check below passes before archival begins. A failure is fixed inside Phase 5.

- [x] [AI] Run `npm run check:workflows` — acceptance: exact-pinned actionlint reports the final scheduled workflow
      structurally valid.
- [x] [AI] Run `npm run test:scheduled` — acceptance: the archival gate proves unit 99% coverage, behavior coverage,
      integration 99% coverage, and process E2E remain green after all Phase 5 work.
- [x] [AI] Run `git diff --check` — acceptance: no whitespace error is reported.
- [x] [AI] Run `git status --short` — acceptance: only completed Phase 5 and plan-record changes remain for the phase
      commit.

> **Pause Safety**: Implementation, tests, governance, security checks, captured learnings, workflow schema, both 99%
> coverage slices, behavior coverage, and process E2E are complete and green. Safe to stop. Resume with
> `npm run check:workflows && npm run test:scheduled` to reprove the final workflow, quick, integration-coverage, and
> E2E boundary.
