# Learnings

The Markdown-link checker enumerates Git-tracked Markdown paths, so a plan lifecycle move must be staged before link
verification. Running the check first makes it inspect the deleted source path even when the destination exists; stage
the move, then validate links before committing.

Phase 0 baseline on 2026-08-23: `npm install` exited 0 with 201 audited packages, no vulnerabilities, and no
`package-lock.json` change; `npm run test:quick` exited 0 with existing aggregate statement coverage at 95.6%;
`npm run format:check` exited 0; and `npm run check:markdown-links` exited 0.

Phase 1's reviewed checklist originally created the repository boundary-policy tests before migrating the known
real-filesystem and Git cases out of the unit target. That ordering could not meet its own green acceptance criterion,
so the policy-test items moved after the boundary migrations; future plans should place enforcement after the state it
enforces is attainable, while keeping fixture-level policy tests earlier only when their scope says so explicitly.

The first Phase 1 push was blocked because the legacy aggregate coverage target measured new BDD and CLI packages before
Phase 3's reviewed coverage slices existed, while migrated external integration tests did not instrument their
production packages. A phase that must push needs its current pre-push denominator to stay attainable without hiding
runtime code or pulling environment-dependent integration into a cacheable hook, so the reviewed all-runtime 99% unit
slice moved into Phase 1; Phase 3 still adds the separate integration denominator and final aggregate composition.

Phase 1 gate on 2026-08-23: `npm exec nx -- run badakmini-cli:test:quick` exited 0 after typecheck, strict lint, unit
tests, and 99.2% statement coverage across all `internal/...` runtime code; `npm run format:check`,
`npm run check:markdown-links`, and `go -C apps/badakmini-cli test ./tests/integration` also exited 0. The exact
production-runtime binding test and all five named integration-owner proofs passed, and the structural assertion found
no `os` or `os/exec` import below `internal/`.

The first Phase 2 feature addition returned a stale cached `go test` pass because the canonical `specs/` directory is
outside the Go module and a newly created sibling feature was absent from the prior test cache's observed inputs.
BeaverNest avoids this through embedded feature resources; the Go adaptation needs both recursive Nx inputs and
`-count=1` on every corpus-executing Go command so additions, edits, renames, nesting, and deletions always re-execute
compliance and adapters.

Rules-propagation inventory on 2026-08-23 found that the existing specs and testing rules named Gherkin and generic unit
or integration tests but had no canonical role matrix, no automated corpus/binding/adapter verification, and no required
public-process E2E for applications. The idempotency gate therefore failed its observable-verification criterion. The
durable destination is the new BDD policy, with concise references from specs, TDD, testing, Badak Mini, workspace
commands, indexes, and `AGENTS.md`.

## Phase 5 Review

- Markdown-link staging behaviour: safe, generalizable; terminal destination is
  [Workspace Commands](../../../repo-governance/development/workspace-commands.md#repository-checks), which already
  records the Git-tracked-file caveat.
- Phase 0 baseline: safe, local; terminal destination is this plan record because its command versions and results are
  dated evidence rather than a repository rule.
- Phase 1 boundary-policy ordering: safe, generalizable; terminal destination is
  [delivery checklists](../../../repo-governance/conventions/plans-organization-policy/delivery-checklists.md), which
  owns dependency-aware task ordering.
- Phase 1 coverage-denominator ordering: safe, generalizable; terminal destination is
  [Testing](../../../repo-governance/development/testing-policy.md), whose explicit target roles prevent a quick gate
  from acquiring an unattainable denominator.
- External-spec cache invalidation: safe, generalizable; terminal destination is the
  [BDD policy](../../../repo-governance/development/behaviour-driven-development-policy.md), which requires recursive Nx
  inputs and uncached execution when native cache keys cannot include the corpus.
- Rules-propagation idempotency result: safe, generalizable; terminal destination is the
  [BDD policy](../../../repo-governance/development/behaviour-driven-development-policy.md), the canonical role and
  enforcement rule created from that failed criterion.
