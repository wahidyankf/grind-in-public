# Phase 3 Measurements

Two numbers this phase produces and a later decision reads back: the unit coverage the ported application arrives with,
and how long `static-routes:validation` takes.

## Starting Unit Coverage

Measured 2026-09-01, immediately after the corpus was bound and before any coverage-raising test was written.
`npx nx run wahidyankf-www:test:coverage:unit`:

```text
All files          |   97.66 |       90 |   97.66 |   97.99 |
ERROR: Coverage for lines (97.99%) does not meet global threshold (99%)
```

**97.99% lines** against a `src/**` denominator, 1.01 points below the floor. The gap is small because the ported corpus
is thorough, and it is real because the denominator is explicit: `coverage.include: ["src/**"]` makes a module no test
imports count as zero rather than disappear from the table.

Every uncovered line at this point:

| File                                       | Uncovered    | Why it is uncovered                                                       |
| ------------------------------------------ | ------------ | ------------------------------------------------------------------------- |
| `src/env-loader.ts`                        | 26           | Composition root. Imported only by `next.config.ts`, which no test loads. |
| `src/env.ts`                               | 3            | Same.                                                                     |
| `src/app/head.tsx`                         | 7            | Route metadata component; no binding renders it.                          |
| `src/features/cv/shell/cv-content.tsx`     | 437-440, 534 | Branches the CV scenarios do not reach.                                   |
| `src/features/home/shell/home-content.tsx` | 173          | One branch the home scenarios do not reach.                               |

`src/app/robots.ts` and `src/app/sitemap.ts` are already at 100%, reached through the `static-filterable-routes`
scenarios, even though the checklist anticipated needing tests for them. Those tests are still written: the checklist
asks for them, and a route-metadata module covered only as a side effect of a crawler scenario has no test that fails
when its output changes.

Two entries in the table are not source at all. `src/app/favicon.ico` and `src/features/ui/shell/index.ts` both report
`0 | 0 | 0 | 0`, because the `src/**` glob is a path glob rather than a language filter. Neither contributes a statement
or a line, so neither moves the percentage in either direction; they are noted here so a future reader does not mistake
them for uncovered code.

## `static-routes:validation` Duration

Measured 2026-09-01, three consecutive runs with `--skip-nx-cache` so each one pays the full cost:

```text
real 5.14
real 4.38
real 4.24
```

**Roughly 4.2 to 5.1 seconds**, the first run slowest. Almost all of it is the `next build` the target runs first; the
`validate-static-routes.mjs` check that reads the build output is the cheap half, and it reports
`Verified static build output for /, /cv, /personal-projects, /robots.txt, /sitemap.xml.` on every run.

This is what a later decision about the target's place in `test:quick` would need. `test:quick` declares
`static-routes:validation` in `dependsOn`, so every quick gate pays this. Against that: an Nx cache hit skips it
entirely when nothing the build reads has changed, and the alternative — discovering that a route stopped being
statically generated at deploy time — costs far more than five seconds. The measurement is recorded rather than acted
on; no item in this plan changes the target's placement.
