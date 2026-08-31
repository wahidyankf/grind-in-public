# wahidyankf-www Migration

Bring the `wahidyankf-www` Next.js application from `ose-public` into this repository as a first-class Nx project that meets this repository's standards rather than the ones it arrives with. This plan covers only the receiving side. `ose-public` is not touched: it keeps its copy, keeps serving the site, and its removal belongs to a separate, separately authorized plan.

The app arrives on TypeScript 5.8.3 with oxlint, an 80% coverage gate, `echo` no-op integration and E2E targets, two `@open-sharia-enterprise/*` workspace dependencies in its manifest — `web-ui` and `ts-env-loader`, with `web-ui-token` reaching it only through two CSS imports — and three specification validators that shell out to an F# CLI this repository does not have. None of that survives the move. What lands is one Nx application at 99% coverage with real unit, integration, and process-E2E layers bound to a canonical Gherkin corpus of twelve feature files and 55 scenarios, laid out the way `badakmini-cli` lays its out.

This plan is staged in `plans/in-progress/` by owner direction because execution follows immediately, which is exactly what the [plan-planning](../../../repo-governance/workflows/plan-planning.md) workflow's Prerequisites permit that stage to mean.

## Affected Projects

- `apps/wahidyankf-www` — new Nx application, ported from `ose-public/apps/wahidyankf-www`.
- `apps/wahidyankf-www-e2e` — new dedicated Playwright E2E project, ported from `ose-public/apps/wahidyankf-www-fe-e2e` and **renamed on arrival** by owner direction. The two projects here are `wahidyankf-www` and `wahidyankf-www-e2e`; the source's `fe-` infix does not come across. Every `$SRC/` path in `delivery.md` therefore keeps the old name while every destination path uses the new one, and the copied `package.json` `"name"` field and `README.md` heading are repaired in Phase 5.
- `specs/apps/wahidyankf-www` — new canonical corpus and as-built C4 model. Nine feature files come from the application's own corpus and two from `libs/ts-env-loader`, whose behavior is inlined with its code.
- `cv/` — deleted; the application absorbs the CV material it duplicated.
- `scripts/` — gains `next-with-port.mjs`, which is currently an empty directory with a `.gitkeep`.
- `repo-governance/development` — gains a deployment rule, and `testing-policy/tooling.md` gains one certain amendment plus any conditional ones. The certain amendment is the language-target deviation both new projects carry on `module`, `moduleResolution`, and `target`, which is known before execution starts because Next 16 leaves no alternative. A conditional one is added for each toolchain component that turns out not to conform, if any does.
- `.github/workflows/full-bdd.yml` and root `package.json` — extended to cover the new project.

## Source Provenance

The port copies the working tree; it does not graft Git history. The source is `ose-public` at commit `e74818fc06c4c104725383384d2aa38305a503ef` (`2026-08-31`, branch `main`, clean working tree for both ported paths). Recording the SHA is what makes a later divergence between the two repositories diagnosable without a shared history.

## Decisions Already Settled

The owner resolved these before authoring. They are not open, and execution does not revisit them.

| Decision | Resolution |
| --- | --- |
| CV source of truth | `apps/wahidyankf-www/src/features/cv/core/data.ts` wins and is the more current record. `cv/` is deleted entirely and the application absorbs it. |
| Reaching 99% coverage | Port the whole application in one phase, but no phase gate passes while coverage is below 99% or a required behavior layer is absent. |
| The three shared libraries | Inline all three into the application. `libs/` stays empty. |
| Specification validators | No `rhino-cli`, no new Badak Mini subcommand. BDD is enforced through this repository's existing specs structure and the Gherkin binding suite. |
| Toolchain | Conform fully: TypeScript 6, Biome, exact pins. If one component proves incompatible with Next 16, fall back on that component alone and record the exception. |
| Process E2E | Port the Playwright project, drop its Docker runner, run it against `next start`. |
| Deployment | Port `vercel.json` unchanged, keep the `prod-wahidyankf-www` branch gate, and document that branch in governance. The domain cutover is deliberately out of scope. |
| Local integration boundary | The filesystem the application genuinely owns: CV data and font reads, and PDF generation to a real file. |
| Static route validation | Kept exactly as `ose-public` has it, inside `test:quick` with `--skip-nx-cache`. |

## Two Reconciliations Worth Reading Before Executing

The owner selected an integration layer that included static route manifest validation, and separately chose to keep `static-routes:validation` exactly as `ose-public` has it. Those overlap. This plan resolves it by leaving route validation where it already works — a `test:quick` dependency running a full uncached build — and giving the integration layer only the filesystem boundary that has no other home. Nothing is dropped; one boundary is simply not asserted twice.

Inlining `ts-env-loader` brings `dotenv` with it. This repository's [dependency selection policy](../../../repo-governance/development/dependency-selection-policy.md) prefers a standard-library facility, so `process.loadEnvFile` was examined first and rejected on evidence: verified on Node 24.16, it takes only a path — there is no target record to load into, which the ported scenarios require — and it throws `ENOENT` on a missing file, which breaks the loader's "absence is not an error" rule. `dotenv` therefore stays, added to the application manifest at an exact pin, with the requirement, the rejected alternative, and that evidence recorded in [technical design](tech-docs/README.md#selected-decisions). This is settled, not an assumption Phase 3 revisits.

## Directory Map

- [Business Requirements](brd.md)
- [Product Requirements](prd.md)
- [Technical Design](tech-docs/README.md)
- [Delivery Checklist](delivery.md)
- [Learnings](learnings.md)
- [Evidence](evidence/README.md) — command output and measurements that later phases read back, kept out of `learnings.md` so Phase 7 triages lessons rather than exit statuses.

## Quality Gate

The [plan-quality-gate](../../../repo-governance/workflows/plan-quality-gate.md) check-fix loop ran at **strict** level before this plan was committed. It ended on its **seven-cycle bound**, not on two consecutive clean runs, and the workflow provides for exactly that: seven cycles end the loop too, with the remaining findings reported.

| Cycle | Critical | High | Medium | Low |
| ----- | -------- | ---- | ------ | --- |
| 1     | 1        | 10   | 15     | 8   |
| 2     | 0        | 5    | 7      | 11  |
| 3     | 0        | 4    | 8      | 6   |
| 4     | 0        | 3    | 4      | 5   |
| 5     | 0        | 3    | 4      | 3   |
| 6     | 0        | 1    | 2      | 1   |
| 7     | 0        | 1    | 2      | 3   |

Every finding of every cycle was resolved, including all six of cycle 7; none was accepted, deferred, or waived. The count did not reach zero because each round of repairs is itself new text a strict reader can find something in — cycle 7's findings are all consequences of cycle 6's, and four of the six are one sentence somewhere else in the plan still describing `vercel.json` as byte-identical after the phase stopped delivering it that way. No cycle ever challenged the approach, a phase order, or a decision. Cycle 5's checker stated the structural verdict directly: **structurally sound, with residual imprecision; I would not rewrite a phase.**

The residue is attributed to size rather than to unsoundness. This is a cross-repository port with seven phases, roughly four hundred checklist items, two new projects, and the first TypeScript project this repository has held; the surface a strict reader can examine is larger than the loop bound was written for. Execution is what tests the rest, and [`learnings.md`](learnings.md) is where anything the plan got wrong is recorded as it is found.

The Phase 7 item that re-runs this workflow at archival appends its result below rather than replacing this record.
