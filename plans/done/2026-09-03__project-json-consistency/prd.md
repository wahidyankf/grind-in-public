# Product Requirements

## Scope Note on Gherkin

Every acceptance criterion below is a **plan-only operational outcome**, not durable application behaviour. This plan
changes build configuration, project layout, and one governance document; it adds no production code path and changes no
scenario in `specs/`. The [specs policy](../../../repo-governance/development/specs-policy.md) permits a plan to retain
such outcomes when [tech-docs/specification-changes.md](tech-docs/specification-changes.md) labels them plan-only, gives
the reason, and names the delivery proof. That document does so for all seven, and each criterion below names its own
proof command.

Because nothing here is durable behaviour, no criterion becomes a RED → GREEN → REFACTOR cycle and none lands in
`specs/apps/*/behaviours/`. The [TDD policy](../../../repo-governance/development/tdd-policy.md) cycle rule binds
behaviour changes in `apps/` and `libs/`; there are none.

## User Stories

**US-1** — As the repository owner, I want every Nx project to expose the same target names, so that a workspace-wide
command reaches every project without special-casing one of them.

**US-2** — As an agent session adding a target, I want one written rule describing what a `project.json` target must
declare, so that I do not have to infer a convention by comparing two files that disagree.

**US-3** — As the repository owner, I want the browser suite to belong to the application it tests, so that editing a
step file is covered by the same push-time gate as editing the application.

## Acceptance Criteria

### [AC-1] Every project exposes the full target contract

```gherkin
Scenario: Both projects expose the ten contract targets
  Given the workspace after this plan lands
  When the target list of every discovered project is inspected
  Then each project declares typecheck, lint, test:unit, test:integration, test:e2e, test:coverage, test:coverage:unit, test:coverage:integration, test:coverage:behaviour, and test:quick
```

Proof: `npx nx show project <name> --json` for each of `badakmini-cli` and `wahidyankf-www`.

### [AC-2] Every target declares its cache state explicitly

```gherkin
Scenario: No target resolves to an undeclared cache state
  Given the workspace after this plan lands
  When every target's resolved cache value is read
  Then no target reports undefined, because each is set in project.json or reached by the root targetDefaults
  And no uncached target declares outputs, because Nx restores nothing for a target it never caches
  And every cached target that writes an artifact declares the path it writes
```

Proof: two inspection commands in `delivery.md`, run once per project. The first reads every target's resolved `cache`
and prints `all N targets declare cache` with a non-zero N, or `UNDECLARED:` and the offending target names; it never
prints a per-target `cache=` line, so the acceptance is that first form. The second reads the two `outputs` halves of
this criterion against a written-out artifact map and prints `outputs rule holds for all N targets`, or `VIOLATIONS:`
naming each uncached target that declares `outputs` and each cached target that writes a mapped path without declaring
it.

Both run three times. The Phase 1 gate runs them over all three projects while all three exist. The Phase 2 gate runs
them again over the two that survive, with the `wahidyankf-www` artifact map extended to carry the `specs:e2e:baseline`
target that phase adds; running them there rather than later is what keeps a failure repairable in the same phase, and
the same commit theme, as the `project.json` that causes it. The Phase 3 gate re-runs that post-merge form once more
against the rule it writes — and that post-merge form is what Phase 4 reconciles, because the Phase 1 form names a
deleted project and its map predates a target the merge introduced.

### [AC-3] A shared input glob is declared once

```gherkin
Scenario: The behaviour corpus is named by a shared input
  Given the workspace after this plan lands
  When a project target depends on the canonical Gherkin corpus
  Then it references a behaviourCorpus namedInput declared once in that project's own project.json rather than repeating the glob
```

Proof, in two halves that fail differently. **Declared once**: `grep -c` for the raw corpus glob over each
`project.json` prints `1`, the single remaining occurrence being the `namedInputs` declaration itself. **Resolves to
something real**: a cache probe, because `npx nx show project --json` reports the declared `inputs` array verbatim and
expands neither `default` nor a named input, so it cannot show what `behaviourCorpus` resolves to. Nx hashes file
content, so appending a Gherkin comment line to a corpus feature file must turn a cache hit into a miss on every target
that names `behaviourCorpus`, and `git checkout --` of that file must turn it back into a hit. `delivery.md`'s Phase 1
gate runs that probe for `badakmini-cli:test:unit`, `wahidyankf-www:test:unit`, and `wahidyankf-www-e2e:typecheck`, and
runs the same probe over `scripts/next-with-port.mjs` for the `workspaceScripts` name.

### [AC-4] Project-relative commands resolve through a declared working directory

```gherkin
Scenario: No target hardcodes its own project path
  Given the workspace after this plan lands
  When every target command is read
  Then none contains a literal apps/<name> path where options.cwd would resolve it
```

Proof: `grep -n 'apps/badakmini-cli' apps/badakmini-cli/project.json` and
`grep -n 'apps/wahidyankf-www' apps/wahidyankf-www/project.json` each return no line inside a `"command"` string, and
`npx nx run badakmini-cli:test:quick` and `npx nx run wahidyankf-www:static-routes:validation` both exit 0. Both files
are read because this criterion binds every target, not only the ten in the contract.

### [AC-5] One form invokes Nx from inside a target

```gherkin
Scenario: A nested Nx invocation uses the workspace-local binary
  Given the workspace after this plan lands
  When a target command invokes another Nx target
  Then it uses npm exec nx -- run rather than a bare nx run
```

Proof: `grep -nE '"command": *"([^"]*[^-] )?nx run' apps/*/project.json` returns nothing after the change, and the same
pattern run before it prints the one `static-routes:validation` command line, so the proof is shown capable of failing.
The group is optional because the bare invocation opens its command string: a pattern requiring a character and a space
before `nx run` never matches it, and would be silent whether or not the defect were present.

### [AC-6] The browser E2E adapter is co-located in the application

```gherkin
Scenario: The application owns its process E2E target
  Given the workspace after this plan lands
  When the discovered project list is read
  Then wahidyankf-www-e2e is absent and wahidyankf-www owns test:e2e against the same corpus
```

Proof: `npx nx show projects` lists exactly `badakmini-cli` and `wahidyankf-www`; `npx nx run wahidyankf-www:test:e2e`
passes; `npx nx run wahidyankf-www:specs:e2e:baseline` exits 0, which is its assertion that the generated `test.fixme`
count still equals the 34 recorded in the moved baseline file. That target is silent on success and prints the two
counts only when they differ, so exit 0 is the whole observation.

### [AC-7] The contract is written down

```gherkin
Scenario: The testing policy states the full target contract
  Given the repository after this plan lands
  When testing-policy.md is read
  Then it names all ten targets, states which are eligibility-dependent, and states when dependsOn rather than options.commands expresses ordering
```

Proof: the Phase 3 gate item that reads `testing-policy.md` against `apps/badakmini-cli/project.json` and
`apps/wahidyankf-www/project.json` and records, rule by rule, which target in which file it was checked against. That
review is what observes this criterion. `npm run check:governance` and `npm run check:markdown-links` run alongside it
and are constraints the edit must respect — the word limit and the new links — but neither observes the criterion: a
policy that named none of the ten targets would pass both.

## Out of Scope

- Renaming any target. The `check:*` family considered during planning is a recorded non-goal in [brd.md](brd.md).
- Any Badak Mini command, target, or pre-push wiring.
- Moving `static-routes:validation` out of `test:quick`'s `dependsOn`.
- Any change to a coverage threshold, denominator, exclusion, or Gherkin scenario.
