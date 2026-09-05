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

`test:quick` is the ordered gate: `typecheck`, `lint`, `test:unit`, `test:coverage:unit`, then
`test:coverage:behaviour`, with `static-routes:validation` ahead of them all. It deliberately stops short of
`test:coverage`, which would pull the integration layer in behind it and write PDFs to the real filesystem on every
push.

## Specifications

The canonical Gherkin corpus is
[specs/apps/wahidyankf-www/behaviours/](../../specs/apps/wahidyankf-www/behaviours/README.md), and the as-built C4 model
is [specs/apps/wahidyankf-www/architecture.md](../../specs/apps/wahidyankf-www/architecture.md).

Three adapters bind that one corpus across the owner and dedicated E2E projects. Which adapter a scenario reaches is a
property of the boundary it touches:

- **Unit behaviour**, `tests/bdd/` plus the CV export adapter — all scenarios under jsdom with injected filesystem and
  environment seams for local-boundary concerns.
- **Integration behaviour**, `tests/integration/` plus the shared env and port adapters — every non-exempt local
  scenario under `node` against isolated real filesystem, environment, or process boundaries.
- **Browser**, `../wahidyankf-www-e2e/tests/steps/` — Playwright against `next start`, with no Docker involved.

The shared env, port, and CV adapters execute once per Unit and Integration project with a boundary-specific driver.
Static compliance requires exactly one Unit marker for every scenario, exactly one Integration marker for every
non-exempt scenario, and no Integration marker for an exempt scenario.

## Browser exemptions

The browser project rejects missing and unused bindings. Scenarios whose concern cannot be observed at the browser
boundary carry a scenario-level `@e2e-exempt` tag and an immediately preceding structured comment naming their
alternative proof. Browser-rendered scenarios use `@integration-exempt` for the inverse boundary mismatch. Both tags may
annotate one scenario when each omitted boundary is independently justified and Unit supplies substantive proof.

```bash
npm exec nx -- run wahidyankf-www-e2e:test:coverage:behaviour:e2e
npm exec nx -- run wahidyankf-www-e2e:test:e2e
```

No skip baseline exists. Static compliance rejects legacy layer tags, malformed or broad exemptions, unconditional
Playwright skips, direct journey specs, undefined bindings, and unused steps.

## Coverage

`test:coverage:unit` and `test:coverage:integration` each enforce a 99% line floor. Both count `src/**` as the
denominator, set explicitly in `vitest.config.ts`: without it only files some test imports would appear at all, and an
untested module would vanish from the measurement rather than count against it.

`test:coverage:unit` runs the `behaviour-unit` project alongside `unit`, because `src/features/env/core/tier-env.ts` and
`port-resolver.ts` are exercised only through their Gherkin bindings and would otherwise report zero against a
whole-`src/**` denominator.

## Code map

Routes in `src/app/` stay thin. The work lives in `src/features/`, where each feature separates pure data and decisions
in `core/` from React and browser behaviour in `shell/`. Three modules that arrived as separate published libraries — a
design system, its tokens, and an environment loader — are inlined here under `features/ui/`, `features/app-shell/`, and
`features/env/`, because this repository publishes no libraries.

## Delivery boundary

Production delivery is automated. Do not push a deployment branch by hand.
