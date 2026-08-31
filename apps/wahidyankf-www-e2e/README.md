# wahidyankf-www-e2e

This project checks `wahidyankf-www` in a real browser: navigation, search, CV content, theme, responsive behavior, and accessibility. It gives the public site a release-confidence signal that no jsdom test can, because it drives the built application over HTTP rather than rendering components in a simulated DOM.

## Why this is a separate project

The [BDD policy](../../repo-governance/development/behavior-driven-development-policy.md) co-locates behavior tests with the code they describe, and grants an exception where the adapter needs a different toolchain. That exception is the only ground for this project, and it is met: the application's behavior adapter is `@amiceli/vitest-cucumber` running under Vitest in jsdom, while this adapter is `playwright-bdd` running under the Playwright runner against a downloaded Chromium binary, with its own generated test directory and its own browser-install step. Two runners cannot share one Vitest project, so co-location is not available.

Because it exists under that exception and nothing more, it owns no separate corpus, no unit layer, and no numeric coverage gate. Its scenarios come from `wahidyankf-www`'s corpus; the application keeps the 99% floor.

## Run the suite

```bash
# Install Chromium once on this machine
npm exec nx -- run wahidyankf-www-e2e:install

# Run the scenarios against a real `next start`
npm exec nx -- run wahidyankf-www-e2e:test:e2e
```

There is no container. `playwright.config.ts` carries a `webServer` block that starts `wahidyankf-www:start` and waits for it, and the `test:e2e` target declares `dependsOn` on `wahidyankf-www:build`, because `next start` serves a `.next` directory it will not build itself. A cold checkout therefore builds, starts, and tests in one command.

`test:e2e` also refuses to run while any step file contains an unconditional `test.skip()`. Use `test.skip(condition, reason)` for a real environment guard, or remove the line.

## Check a deployed environment

Set `BASE_URL` only when deliberately running the suite against an already-running staging or production site. Keep any access values out of tracked files.

## What this suite deliberately does not bind

Eight of the corpus's twelve feature files bind here — `accessibility`, `cv`, `home`, `personal-projects`, `responsive`, `search`, `static-filterable-routes`, and `theme` — for 36 scenarios in all. The other four bind elsewhere, because their behavior has no browser to drive:

| Feature                    | Scenarios | Bound by                            |
| -------------------------- | --------- | ----------------------------------- |
| `env-loader.feature`       | 4         | the application's unit layer        |
| `tier-env-loading.feature` | 5         | the application's unit layer        |
| `port-resolver.feature`    | 8         | the application's unit layer        |
| `cv-export.feature`        | 2         | the application's integration layer |

The first three are Node-process environment concerns — which tier file loads, which variable wins, which port resolves — settled before a browser exists. The CV export is a build-time script that writes a PDF to disk, so it binds at the filesystem boundary instead.

## The skip baseline

`playwright.config.ts` sets `missingSteps: "skip-scenario"`, so `bddgen` renders every unbound scenario as `test.fixme` rather than refusing to generate anything. That keeps the suite runnable, but it also means the suite exits 0 whether the gap is the intended one above or a binding someone just broke. `e2e-skip-baseline.json` is what tells those two apart:

```bash
npm exec nx -- run wahidyankf-www-e2e:specs:e2e:baseline
```

It regenerates the tests, counts the `test.fixme` entries, and fails if the number moved. The recorded number is **34**, and it counts **generated tests, not scenarios**. The two differ because `playwright-bdd` generates one test per `Examples` row and three of the four unbound features are Scenario Outlines: `env-loader` produces 6, `tier-env-loading` 7, `port-resolver` 19, and `cv-export` 2. Nineteen scenarios, thirty-four generated tests.

Raise the number only when a scenario is deliberately left unbound, and say here why.

## Checks

```bash
npm exec nx -- run wahidyankf-www-e2e:typecheck
npm exec nx -- run wahidyankf-www-e2e:lint
```

`lint` is an ordered aggregate of `lint:biome` and `lint:commentary`, matching the shape the application uses.

The behavior source of truth is [the wahidyankf-www Gherkin corpus](../../specs/apps/wahidyankf-www/behavior/README.md).
