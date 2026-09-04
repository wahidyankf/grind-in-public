# wahidyankf-www Specifications

The canonical as-built boundaries for `wahidyankf-www` live in [architecture.md](architecture.md), and its executable
behaviour lives in [behaviours/](behaviours/README.md). Twelve feature files carry the executable scenarios.

The owner application hosts Unit and Integration adapters. `apps/wahidyankf-www-e2e` owns the browser adapter and
depends one-way on the owner; it owns no feature files.

## Adapters

Three adapters run against this corpus, and they do not all reach the same scenarios.

| Adapter           | Where                                                      | Reaches                                                                    |
| ----------------- | ---------------------------------------------------------- | -------------------------------------------------------------------------- |
| Unit behaviour    | `apps/wahidyankf-www/tests/bdd/`                           | All 70 expanded scenarios, with injected seams for OS-facing dependencies. |
| Local integration | `apps/wahidyankf-www/tests/integration/` + shared adapters | All 34 local-boundary scenarios against isolated real resources.           |
| Browser E2E       | `apps/wahidyankf-www-e2e/tests/steps/`                     | All 36 browser-observable scenarios against a started production server.   |

Scenario-level `@integration-exempt` and `@e2e-exempt` tags document genuine boundary mismatches and name an alternative
Nx target plus scenario. Static compliance rejects missing, malformed, broad, doubled, and operationally motivated
exemptions; the semantic implementation review verifies their substance.

## Targets

| Target                                                      | What it proves                                                      |
| ----------------------------------------------------------- | ------------------------------------------------------------------- |
| `npx nx run wahidyankf-www:test:coverage:behaviour`         | Unit, Integration, E2E rows, exemptions, and bindings are complete. |
| `npx nx run wahidyankf-www:test:coverage:unit`              | Unit and behaviour together reach the 99% line floor.               |
| `npx nx run wahidyankf-www:test:coverage:integration`       | The integration adapter reaches the 99% line floor.                 |
| `npx nx run wahidyankf-www-e2e:test:coverage:behaviour:e2e` | Browser corpus, exemptions, and bindings are complete.              |
| `npx nx run wahidyankf-www-e2e:test:e2e`                    | The browser suite passes against `next start`.                      |

## Directory Map

- [Architecture](architecture.md) is the current as-built C4 model, its boundaries, and behaviour traceability.
- [Behaviour](behaviours/README.md) contains the canonical executable Gherkin corpus.
