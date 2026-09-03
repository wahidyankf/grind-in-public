# wahidyankf-www Specifications

The canonical as-built boundaries for `wahidyankf-www` live in [architecture.md](architecture.md), and its executable
behavior lives in [behavior/](behavior/README.md). Eleven feature files carry 53 scenarios.

Two projects share this corpus, as the
[architecture specification policy](../../../repo-governance/development/architecture-specifications.md) provides for:
the application at `apps/wahidyankf-www` and the dedicated E2E project at `apps/wahidyankf-www-e2e`.

## Adapters

Three adapters run against this corpus, and they do not all reach the same scenarios.

| Adapter           | Where                                    | Reaches                                                                           |
| ----------------- | ---------------------------------------- | --------------------------------------------------------------------------------- |
| Behavior          | `apps/wahidyankf-www/tests/bdd/`         | All eleven feature files, under `@amiceli/vitest-cucumber` with jsdom.            |
| Local integration | `apps/wahidyankf-www/tests/integration/` | `cv-export.feature`, against the real filesystem and build output.                |
| Process E2E       | `apps/wahidyankf-www-e2e/steps/`         | The browser-observable features, under `playwright-bdd` against a started server. |

The E2E adapter deliberately leaves four features unbound — `env-loader.feature`, `tier-env-loading.feature`,
`port-resolver.feature`, and `cv-export.feature` — because they specify Node-process environment concerns and a
build-time export that no browser reaches. That gap is not left to be noticed: `apps/wahidyankf-www-e2e` records a
generated skip baseline and fails when the number moves, so a newly broken binding cannot hide among the intended skips.

## Targets

| Target                                                | What it proves                                                                               |
| ----------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `npx nx run wahidyankf-www:test:coverage:behavior`    | Every feature file is loaded by exactly one binding and every scenario in it is implemented. |
| `npx nx run wahidyankf-www:test:coverage:unit`        | Unit and behavior together reach the 99% line floor.                                         |
| `npx nx run wahidyankf-www:test:coverage:integration` | The integration adapter reaches the 99% line floor.                                          |
| `npx nx run wahidyankf-www-e2e:test:e2e`              | The browser suite passes against `next start`.                                               |
| `npx nx run wahidyankf-www-e2e:specs:e2e:baseline`    | The count of generated skipped scenarios still matches the recorded baseline.                |

## Directory Map

- [Architecture](architecture.md) is the current as-built C4 model, its boundaries, and behavior traceability.
- [Behavior](behavior/README.md) contains the canonical executable Gherkin corpus.
