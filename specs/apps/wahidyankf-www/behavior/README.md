# wahidyankf-www Behavior

The canonical Gherkin corpus for `wahidyankf-www`. Eleven feature files carry 53 scenarios, and each file is loaded by exactly one binding, so a scenario added here has to be bound or `test:coverage:behavior` fails.

Nine of the eleven describe the site itself and bind under `apps/wahidyankf-www/tests/bdd/`. The other two describe the environment loader that was inlined into the application during the migration and would otherwise have arrived with no corpus at all: `tier-env-loading.feature` and `port-resolver.feature`. They are here rather than under a `libs/` corpus because this repository publishes no libraries and the module they specify now lives inside the app.

Scenario shape and cardinality belong to the [specs policy](../../../../repo-governance/development/specs-policy.md); how a scenario reaches a test belongs to the [TDD policy](../../../../repo-governance/development/tdd-policy.md). The as-built boundaries these scenarios exercise live in [architecture.md](../architecture.md).

Four of these features carry scenarios the Playwright adapter deliberately does not bind, because they are Node-process environment concerns with no browser equivalent, or a build-time export no browser reaches. `apps/wahidyankf-www-e2e/README.md` names them and records the generated skip baseline that keeps the gap from widening unnoticed.

## Directory Map

- [accessibility.feature](accessibility.feature) — keyboard reachability and assistive-technology labelling across the site.
- [cv.feature](cv.feature) — the CV page and its rendered record.
- [env-loader.feature](env-loader.feature) — the application's own tier-aware environment loading.
- [home.feature](home.feature) — the landing page and its entry points.
- [personal-projects.feature](personal-projects.feature) — the projects listing and its detail routes.
- [port-resolver.feature](port-resolver.feature) — runtime listener port resolution, from the inlined loader module.
- [responsive.feature](responsive.feature) — layout across viewport widths.
- [search.feature](search.feature) — site search and its result set.
- [static-filterable-routes.feature](static-filterable-routes.feature) — the statically generated filtered portfolio routes.
- [theme.feature](theme.feature) — the light and dark theme toggle.
- [tier-env-loading.feature](tier-env-loading.feature) — tier env-file loading in the inlined loader module.
