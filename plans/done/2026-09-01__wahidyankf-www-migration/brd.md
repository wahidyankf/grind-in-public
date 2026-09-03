# Business Requirements

## Why This Work Exists

`wahidyankf-www` is the owner's personal site: the public CV, the personal-projects listing, and the search over both.
It currently lives in `ose-public`, a repository built around a different premise — automation-first, multi-domain, six
applications sharing a design system and an F# tooling CLI. The site is none of those things. It is one person's public
face, and it belongs in the repository that is explicitly about that person's practice.

The second reason is sharper and specific to this repository. `cv/` here already holds `cv-raw.md`, declared the factual
source of truth, alongside an ATS export and a Python generator. The site holds `data.ts`, 544 lines of the same career
record, and it is the more current one. Two hand-maintained CVs in two repositories is not a structural annoyance; it is
a guarantee that the owner will eventually publish a stale one. Consolidation is the point, and the move is what makes
it possible.

## Who It Affects

The owner is the only maintainer and the only person whose workflow changes. Readers of the site are affected only if
the migration breaks something, which is why the plan refuses to close a gate below 99% coverage and defers the domain
cutover entirely.

Every future agent working in this repository is also affected. This becomes the first TypeScript project here, so it
establishes what a TypeScript project looks like under these rules — a precedent that will be copied rather than
re-reasoned.

## What Success Means

The site builds, tests, and lints inside this repository at 99% coverage with real unit, integration, and process-E2E
layers bound to a canonical Gherkin corpus. There is exactly one CV record in the repository. `libs/` is still empty and
no F# or .NET toolchain has appeared. `ose-public` is untouched and still serving production.

Success explicitly does not include the site being served from this repository. That is a later decision with a later
plan.

## Non-Goals

- Removing anything from `ose-public`, including the app, its specs, or its Vercel project.
- Switching the production domain, or promoting the `prod-wahidyankf-www` branch in this repository.
- Redesigning the site. The user interface is ported as-is; this is a relocation and a conformance exercise, not a
  redesign.
- Porting `libs/web-ui` as a design system. Four components are consumed and four components are inlined.
- Generalizing anything for a hypothetical second web application here.

## Risks

**The toolchain conformance may not hold.** TypeScript 6, Biome, and the ESLint commentary check are all required by
policy but none has ever been executed in this repository, and the app is Next 16 with React 19. The owner has set the
fallback: retreat on whichever component breaks, alone, and record the exception. The risk is not failure, it is silent
partial conformance, so the fallback has to reach `tooling.md` rather than only a commit message.

**Phase 3 is large.** Porting the application, inlining three libraries, and reaching 99% land in one phase because no
intermediate state is allowed below the floor. A phase this size is harder to bisect if it goes wrong. It is accepted
deliberately, and mitigated by ordering its checklist so each feature area reaches green before the next begins.

**Deleting `cv/` is irreversible in the working tree.** It is recoverable from Git, but the material includes evidence
and drafting notes that nothing else regenerates. The plan absorbs the files before deleting the directory, and proves
the absorption in the same phase.

**Pre-push gets slower.** Keeping `static-routes:validation` inside `test:quick` means every push affecting this project
runs a full uncached Next production build, in a repository whose pre-push currently costs seconds of Go tests. This is
a judgment call the owner made knowingly, in favour of catching a broken route at push time. If it proves intolerable in
practice, that belongs in `learnings.md` and then in a follow-up, not in a quiet edit to the target.

**A second CV consumer may be missed.** `cv/README.md` states that `cv-raw.md` is read before any CV-related claim is
made. Agents and prompts elsewhere may reference that path. The plan searches for and repairs those references rather
than assuming the directory is unreferenced.
