# wahidyankf-www

The personal site of Wahidyan Kresna Fridayoka: a CV, a projects listing, and search across both. It is a small,
accessible Next.js application, and it holds the repository's single authoritative CV record in
`src/features/cv/core/data.ts`.

## Run it locally

```bash
# Start at http://localhost:3201
npx nx run wahidyankf-www:dev

# Create a production build
npx nx run wahidyankf-www:build

# Serve a completed local build
npx nx run wahidyankf-www:start
```

The listener port resolves in one order everywhere — an explicit `--port` flag, then `WAHIDYANKF_WWW_PORT`, then the
compiled-in `3201`. A bare `PORT` never moves it. See `src/features/env/core/port-resolver.ts` for why the wrapper at
`scripts/next-with-port.mjs` exists at all.

## Check your changes

```bash
npx nx run wahidyankf-www:test:quick
npx nx run wahidyankf-www:test:coverage
npx nx run wahidyankf-www:static-routes:validation
```

`test:quick` is the ordered gate: `typecheck`, `lint`, `test:unit`, `test:coverage:unit`, then `test:coverage:behavior`,
with `static-routes:validation` ahead of them all. It deliberately stops short of `test:coverage`, which would pull the
integration layer in behind it and write PDFs to the real filesystem on every push.

## Specifications

The canonical Gherkin corpus is
[specs/apps/wahidyankf-www/behavior/](../../specs/apps/wahidyankf-www/behavior/README.md), and the as-built C4 model is
[specs/apps/wahidyankf-www/architecture.md](../../specs/apps/wahidyankf-www/architecture.md).

Three adapters bind that one corpus, and which adapter a feature file reaches is a property of what the scenario touches
rather than a preference:

- **Unit and behavior**, `tests/bdd/` — eleven feature files, under jsdom.
- **Integration**, `tests/integration/` — `cv-export.feature` only, under `node` against the real filesystem, because
  its scenarios write a PDF.
- **Browser**, `tests/e2e/steps/` — Playwright against `next start`, with no Docker involved.

## The browser layer and its skip baseline

`playwright.config.ts` sets `missingSteps: "skip-scenario"`, so `bddgen` renders every unbound scenario as `test.fixme`
rather than refusing to generate anything. That keeps the suite runnable, but it also means the suite exits 0 whether
the gap is the intended one or a binding someone just broke. `tests/e2e/e2e-skip-baseline.json` is what tells those two
apart:

```bash
npm exec nx -- run wahidyankf-www:specs:e2e:baseline
```

Four feature files carry scenarios the browser adapter deliberately does not bind, because they are Node-process
environment concerns or a build-time export no browser reaches:

| Feature                    | Scenarios | Bound instead by                    |
| -------------------------- | --------- | ----------------------------------- |
| `env-loader.feature`       | 4         | the application's unit layer        |
| `tier-env-loading.feature` | 5         | the application's unit layer        |
| `port-resolver.feature`    | 8         | the application's unit layer        |
| `cv-export.feature`        | 2         | the application's integration layer |

The baseline regenerates the tests, counts the `test.fixme` entries, and fails if the number moved. The recorded number
is **34**, and it counts **generated tests, not scenarios**. The two differ because `playwright-bdd` generates one test
per `Examples` row and three of the four unbound features are Scenario Outlines: `env-loader` produces 6,
`tier-env-loading` 7, `port-resolver` 19, and `cv-export` 2. Nineteen scenarios, thirty-four generated tests.

Raise the number only when a scenario is deliberately left unbound, and say here why.

## Coverage

`test:coverage:unit` and `test:coverage:integration` each enforce a 99% line floor. Both count `src/**` as the
denominator, set explicitly in `vitest.config.ts`: without it only files some test imports would appear at all, and an
untested module would vanish from the measurement rather than count against it.

`test:coverage:unit` runs the `behavior` project alongside `unit`, because `src/features/env/core/tier-env.ts` and
`port-resolver.ts` are exercised only through their Gherkin bindings and would otherwise report zero against a
whole-`src/**` denominator.

## Code map

Routes in `src/app/` stay thin. The work lives in `src/features/`, where each feature separates pure data and decisions
in `core/` from React and browser behavior in `shell/`. Three modules that arrived as separate published libraries — a
design system, its tokens, and an environment loader — are inlined here under `features/ui/`, `features/app-shell/`, and
`features/env/`, because this repository publishes no libraries.

## Delivery boundary

Production delivery is automated. Do not push a deployment branch by hand.
