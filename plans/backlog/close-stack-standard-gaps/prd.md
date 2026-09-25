# Product Requirements

## Scope Note on Gherkin

Every acceptance criterion below is a **plan-only operational outcome**: it changes gates, configuration, and
governance, and adds no production code path or scenario under `specs/`.
[tech-docs.md](tech-docs.md#specification-changes) labels them plan-only, and each criterion names its proof. Where a
unit changes a validator's behaviour (U4), that change is built test-first in the validator's own `node --test` suite
per the [TDD policy](../../../repo-governance/development/tdd-policy.md).

## User Stories

**US-1** — As the owner, I want every adopted stack gate enforced by a command, so that a defect the standard names
fails a gate instead of waiting for a reviewer.

**US-2** — As an agent session auditing code, I want the adapter to list no known gaps, so that a checker finding points
at a real regression rather than at an accepted omission.

**US-3** — As the owner, I want project boundaries declared as tags, so that a new dependency between projects is a
visible, checked decision.

## Acceptance Criteria

### [AC-1] TypeScript lint is type-aware

```gherkin
Scenario: A floating promise fails TypeScript lint
  Given a TypeScript file in wahidyankf-www with an unawaited promise-returning call
  When the project's lint target runs
  Then it exits non-zero and names the floating promise
  And with the probe file removed, lint exits 0 for wahidyankf-www and wahidyankf-www-e2e
```

Proof: a temporary probe file, never committed, then the lint target of each project.

### [AC-2] Authored JavaScript is type-checked through JSDoc

```gherkin
Scenario: A JSDoc type error in an authored script fails the type check
  Given the recorded check scope covers every tracked .js, .mjs, and .cjs file outside the vendored copy
  When a probe passes a number where a JSDoc-typed string parameter is declared
  Then the JavaScript type check exits non-zero
  And with the probe reverted, the check exits 0 under strict and noUncheckedIndexedAccess
```

Proof: the JavaScript type-check command named in [tech-docs.md](tech-docs.md#u2-javascript-type-check) and a file-count
comparison against `git ls-files '*.js' '*.mjs' '*.cjs'`.

### [AC-3] The Python pilot has format, lint, and branch-coverage gates

```gherkin
Scenario: The pilot's quick gate formats, lints, and measures branches
  Given apps/forum-be-python
  When its test:quick target runs
  Then it runs the formatter in check mode, the linter with a committed rule selection, pyright, and the unit suite under branch coverage
  And a probe with a misformatted line or an unused import makes it exit non-zero
```

Proof: the pilot's `test:quick` target, with and without a probe edit.

### [AC-4] Every Nx project carries tags from a recorded vocabulary

```gherkin
Scenario: The project-contract validator enforces tags and their constraints
  Given the tag vocabulary and dependency constraints recorded in the adapter
  When node scripts/check-project-contract.mjs runs
  Then every discovered Nx project carries at least one tag from that vocabulary
  And a project with no tag, an unknown tag, or a dependency the constraints forbid is a finding
```

Proof: `node --test scripts/project-contract.test.mjs` with a failing test first for each finding kind, then the
validator over the real workspace.

### [AC-5] The CI test scripts run under Bash in strict mode

```gherkin
Scenario: Both CI test scripts declare Bash and strict mode
  Given .github/scripts/test-hippo-bootstrap.sh and .github/scripts/test-pre-push-contract.sh
  When their first lines and strict-mode line are read
  Then each starts with a Bash shebang and sets -euo pipefail
  And both still pass when run as scripts/check-repo.sh runs them
```

Proof: `head -n 2` on both files and `npm run test:repo`.

### [AC-6] The adapter records no known gap

```gherkin
Scenario: The repository adapter lists no known gap
  Given all five units have landed
  When the repository adapter is read
  Then it contains no "Known gaps" sentence and no Adopter Decisions row whose choice is "gap"
  And each former gap names the gate that now enforces it
```

Proof: `grep -nE 'Known gaps|\| gap ' repo-governance/development/quality/stacks/repository-adapter.md` prints nothing,
paired with `grep -c '## Adopter Decisions' <same file>` printing `1` to prove the search read the real file.
