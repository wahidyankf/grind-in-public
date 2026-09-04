# wahidyankf-www Migration

Bring the `wahidyankf-www` Next.js application from `ose-public` into this repository as a first-class Nx project that
meets this repository's standards rather than the ones it arrives with. This plan covers only the receiving side.
`ose-public` is not touched: it keeps its copy, keeps serving the site, and its removal belongs to a separate,
separately authorized plan.

The app arrives on TypeScript 5.8.3 with oxlint, an 80% coverage gate, `echo` no-op integration and E2E targets, two
`@open-sharia-enterprise/*` workspace dependencies in its manifest — `web-ui` and `ts-env-loader`, with `web-ui-token`
reaching it only through two CSS imports — and three specification validators that shell out to an F# CLI this
repository does not have. None of that survives the move. What lands is one Nx application at 99% coverage with real
unit, integration, and process-E2E layers bound to a canonical Gherkin corpus of twelve feature files and 55 scenarios,
laid out the way `badakmini-cli` lays its out.

This plan is staged in `plans/in-progress/` by owner direction because execution follows immediately, which is exactly
what the [plan-planning](../../../repo-governance/workflows/plan-planning.md) workflow's Prerequisites permit that stage
to mean.

## Affected Projects

- `apps/wahidyankf-www` — new Nx application, ported from `ose-public/apps/wahidyankf-www`.
- `apps/wahidyankf-www-e2e` — new dedicated Playwright E2E project, ported from `ose-public/apps/wahidyankf-www-fe-e2e`
  and **renamed on arrival** by owner direction. The two projects here are `wahidyankf-www` and `wahidyankf-www-e2e`;
  the source's `fe-` infix does not come across. Every `$SRC/` path in `delivery.md` therefore keeps the old name while
  every destination path uses the new one, and the copied `package.json` `"name"` field and `README.md` heading are
  repaired in Phase 5.
- `specs/apps/wahidyankf-www` — new canonical corpus and as-built C4 model. Nine feature files come from the
  application's own corpus and two from `libs/ts-env-loader`, whose behaviour is inlined with its code.
- `cv/` — deleted; the application absorbs the CV material it duplicated.
- `scripts/` — gains `next-with-port.mjs`, which is currently an empty directory with a `.gitkeep`.
- `repo-governance/` and the three harness directories — `development/` gains a deployment rule, and
  `testing-policy/tooling.md` gains two certain amendments plus any conditional ones. **Execution reached further than
  this bullet anticipated**, because Phase 7's learnings triage routes each lesson to a durable home and which homes
  those are cannot be known before the lessons exist: it also edited `development/code-style-policy.md`,
  `development/dependency-selection-policy.md`, `development/behaviour-driven-development-policy.md`,
  `conventions/plans-organization-policy/delivery-checklists.md`, and all three `plan-checker` prompts under `.claude/`,
  `.codex/`, and `.opencode/`. The archival quality gate then reached four more, for a reason of its own: closing the
  word limit policy's headroom band on `conventions/plans-organization-policy/delivery-checklists.md` required
  relocating a section, so this plan also authors `conventions/plans-organization-policy/execution-record.md` and edits
  `conventions/plans-organization-policy.md`, `conventions/plans-organization-policy/README.md`, and
  `workflows/plan-execution/02-phase-loop.md`. [File impact](tech-docs/file-impact.md) maps all eleven. Two amendments
  are certain, not one, and both are known before execution starts: the language-target deviation both new projects
  carry on `module`, `moduleResolution`, and `target`, because Next 16 leaves no alternative, and Biome running as a
  linter only, because Prettier stays the formatting source of truth. `tooling.md`'s `## Recorded Deviations` is the
  register. A conditional one is added for each toolchain component that turns out not to conform, if any does.
- `.github/workflows/full-bdd.yml` and root `package.json` — extended to cover the new project.

## Source Provenance

The port copies the working tree; it does not graft Git history. The source is `ose-public` at commit
`e74818fc06c4c104725383384d2aa38305a503ef` (`2026-08-31`, branch `main`, clean working tree for both ported paths).
Recording the SHA is what makes a later divergence between the two repositories diagnosable without a shared history.

## Decisions Already Settled

The owner resolved these before authoring. They are not open, and execution does not revisit them.

| Decision                   | Resolution                                                                                                                                                         |
| -------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| CV source of truth         | `apps/wahidyankf-www/src/features/cv/core/data.ts` wins and is the more current record. `cv/` is deleted entirely and the application absorbs it.                  |
| Reaching 99% coverage      | Port the whole application in one phase, but no phase gate passes while coverage is below 99% or a required behaviour layer is absent.                             |
| The three shared libraries | Inline all three into the application. `libs/` stays empty.                                                                                                        |
| Specification validators   | No `rhino-cli`, no new Badak Mini subcommand. BDD is enforced through this repository's existing specs structure and the Gherkin binding suite.                    |
| Toolchain                  | Conform fully: TypeScript 6, Biome, exact pins. If one component proves incompatible with Next 16, fall back on that component alone and record the exception.     |
| Process E2E                | Port the Playwright project, drop its Docker runner, run it against `next start`.                                                                                  |
| Deployment                 | Port `vercel.json` unchanged, keep the `prod-wahidyankf-www` branch gate, and document that branch in governance. The domain cutover is deliberately out of scope. |
| Local integration boundary | The filesystem the application genuinely owns: CV data and font reads, and PDF generation to a real file.                                                          |
| Static route validation    | Kept exactly as `ose-public` has it, inside `test:quick` with `--skip-nx-cache`.                                                                                   |

## Two Reconciliations Worth Reading Before Executing

The owner selected an integration layer that included static route manifest validation, and separately chose to keep
`static-routes:validation` exactly as `ose-public` has it. Those overlap. This plan resolves it by leaving route
validation where it already works — a `test:quick` dependency running a full uncached build — and giving the integration
layer only the filesystem boundary that has no other home. Nothing is dropped; one boundary is simply not asserted
twice.

Inlining `ts-env-loader` brings `dotenv` with it. This repository's
[dependency selection policy](../../../repo-governance/development/dependency-selection-policy.md) prefers a
standard-library facility, so `process.loadEnvFile` was examined first and rejected on evidence: verified on Node 24.16,
it takes only a path — there is no target record to load into, which the ported scenarios require — and it throws
`ENOENT` on a missing file, which breaks the loader's "absence is not an error" rule. `dotenv` therefore stays, added to
the application manifest at an exact pin, with the requirement, the rejected alternative, and that evidence recorded in
[technical design](tech-docs/README.md#selected-decisions). This is settled, not an assumption Phase 3 revisits.

## Directory Map

- [Business Requirements](brd.md)
- [Product Requirements](prd.md)
- [Technical Design](tech-docs/README.md)
- [Delivery Checklist](delivery.md)
- [Learnings](learnings.md)
- [Evidence](evidence/README.md) — command output and measurements that later phases read back, kept out of
  `learnings.md` so Phase 7 triages lessons rather than exit statuses.

## Quality Gate

- 2026-08-31 — strict — 7 cycles — settled (nothing open; the last fixes were applied after the final check)

The pre-execution loop ended on its **seven-cycle bound** rather than on two consecutive clean runs, which the workflow
provides for. Every finding of every cycle was resolved; none was accepted, deferred, or waived. The status is `settled`
rather than `pass` because cycle 7's fixes were applied after the last check, so no cycle read them. As
[plan quality gate](../../../repo-governance/workflows/plan-quality-gate.md) requires of a settled plan, the archival
run begins by verifying those fixes, and records its own result below.

The Phase 7 item that re-runs this workflow at archival appends its result below rather than replacing this record.

- 2026-09-01 — strict — 7 cycles — settled (nothing open; the last fixes were applied after the final check)

The archival run ended on the seven-cycle bound too. It verified the settled pre-execution fixes first, as a settled
plan's next run must, then read the plan execution had actually produced. Every finding of every cycle was resolved;
none was accepted, deferred, or waived. Three of them were defects in the delivered repository rather than in the plan,
which is why this run is recorded as work rather than as a formality.
