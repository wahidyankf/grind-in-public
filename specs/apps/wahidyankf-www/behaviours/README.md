# wahidyankf-www Behaviour

The canonical Gherkin corpus for `wahidyankf-www`. Twelve feature files carry 55 scenario declarations and 70 expanded
scenarios. A scenario added here must have one Unit row, one Integration row or valid `@integration-exempt`, and one
browser row or valid `@e2e-exempt`; `test:coverage:behaviour` enforces that contract.

Nine of the twelve describe the site itself and bind under `apps/wahidyankf-www/tests/bdd/`. Two more describe the
environment loader that was inlined into the application during the migration and would otherwise have arrived with no
corpus at all: `tier-env-loading.feature` and `port-resolver.feature`. They are here rather than under a `libs/` corpus
because this repository publishes no libraries and the module they specify now lives inside the app.

Local-boundary features execute the same thin adapter in two modes. Unit uses injected filesystem or environment seams;
Integration uses isolated real temporary directories, `process.env`, or child processes under the Node environment.
Browser-only scenarios carry `@integration-exempt`; local-only scenarios carry `@e2e-exempt` and name their Integration
proof. A scenario may carry both tags when each omitted boundary is independently justified and the mandatory Unit
adapter provides substantive alternative proof.

Scenario shape and cardinality belong to the [specs policy](../../../../repo-governance/development/specs-policy.md);
how a scenario reaches a test belongs to the [TDD policy](../../../../repo-governance/development/tdd-policy.md). The
as-built boundaries these scenarios exercise live in [architecture.md](../architecture.md).

The dedicated `apps/wahidyankf-www-e2e/` project recursively consumes this corpus. It has no skip baseline: static
compliance validates exact owner rows, exemption syntax and alternative references, then Playwright generation rejects
undefined, ambiguous, or unused browser bindings.

## Directory Map

- [accessibility.feature](accessibility.feature) — keyboard reachability and assistive-technology labelling across the
  site.
- [cv-export.feature](cv-export.feature) — the CV PDF export at the filesystem boundary; bound from
  `apps/wahidyankf-www/tests/integration/` with an in-memory Unit driver and real Integration driver.
- [cv.feature](cv.feature) — the CV page and its rendered record.
- [env-loader.feature](env-loader.feature) — the application's own tier-aware environment loading.
- [home.feature](home.feature) — the landing page and its entry points.
- [personal-projects.feature](personal-projects.feature) — the projects listing and its detail routes.
- [port-resolver.feature](port-resolver.feature) — runtime listener port resolution, from the inlined loader module.
- [responsive.feature](responsive.feature) — layout across viewport widths.
- [search.feature](search.feature) — site search and its result set.
- [static-filterable-routes.feature](static-filterable-routes.feature) — the statically generated filtered portfolio
  routes.
- [theme.feature](theme.feature) — the light and dark theme toggle.
- [tier-env-loading.feature](tier-env-loading.feature) — tier env-file loading in the inlined loader module.
