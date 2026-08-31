# wahidyankf-www — Architecture

The as-built C4 model for `wahidyankf-www`, a personal portfolio and CV site. Its executable behavior lives in [behavior/](behavior/README.md); this document describes the boundaries those scenarios exercise.

The site is fully static. No server-side code runs at request time, every page is rendered at build time, and no user data is collected, stored, or transmitted beyond the access logs the CDN keeps. That single constraint is why several boundaries a portfolio site might be expected to have are absent below rather than merely undocumented: there is no API, no database, no session, and no authentication to draw.

## Scope

One deployable application and one dedicated E2E project, sharing this model as the [architecture specification policy](../../../repo-governance/development/architecture-specifications.md) requires of projects that share a corpus.

| Project | Role |
| --- | --- |
| `apps/wahidyankf-www` | The Next.js application, its unit and behavior adapters, and its local integration adapter. |
| `apps/wahidyankf-www-e2e` | The process E2E adapter, driving a real browser against a started server. |

## People and External Systems

| Actor | Description |
| --- | --- |
| Visitor | An anonymous browser user. There is nothing to authenticate as; every page is public. |

| External system | Interaction |
| --- | --- |
| Vercel | Builds the site from `main` and serves it, terminating TLS and holding the edge cache. A build runs only when the `ignoreCommand` in `vercel.json` sees the promotion branch. |

## System Context

```text
   +-----------+                          +-------------------------+
   |  Visitor  |  browses the portfolio   |     wahidyankf-www      |
   | [Person]  | -----------------------> |   [Software System]     |
   +-----------+          HTTPS           |  Static Next.js site    |
                                          +-------------------------+
                                                       |
                                                built and served by
                                                       v
                                          +-------------------------+
                                          |         Vercel          |
                                          |   [External System]     |
                                          |  Build, TLS, edge cache |
                                          +-------------------------+
```

## Containers

One container. There is no backend, no database, and no message bus; all content is TypeScript bundled at build time.

| Container | Technology | Deployment | Description |
| --- | --- | --- | --- |
| `web` | Next.js 16 | Vercel, static | The whole site: every route, the client-side search index, and the CV record. |

```text
   +-----------+                +----------------------------------+
   |  Visitor  | -- HTTPS ----> |               web                |
   | [Person]  |                |  [Container: Next.js 16]         |
   +-----------+                |  Statically rendered at build    |
                                +----------------------------------+
                                                 ^
                                                 | starts and drives
                                                 | over a local port
                                  +--------------------------------+
                                  |     wahidyankf-www-e2e         |
                                  |  [Container: Playwright]       |
                                  |  Process E2E adapter           |
                                  +--------------------------------+
```

The E2E adapter is drawn because it is a real process boundary: it starts the application through `next start` and drives it over HTTP through a browser, which is what makes it a different toolchain from the in-process behavior adapter. It exists only at test time and is never deployed.

## Components

The application is organized by feature, under `src/features/`, with routes in `src/app/`.

| Component | Path | Responsibility |
| --- | --- | --- |
| `app-shell` | `src/features/app-shell/` | Navigation chrome: sidebar on desktop, bottom tab bar on mobile. |
| `home` | `src/features/home/` | The landing page — top skills, languages, frameworks, and the about-me summary. |
| `cv` | `src/features/cv/` | The CV page and the CV record itself. |
| `personal-projects` | `src/features/personal-projects/` | The projects listing and its technology-tag filter. |
| `search` | `src/features/search/` | Client-side search across every other component's content. |
| `ui/shell` | `src/features/ui/shell/` | The presentation primitives the shell renders — inlined during the migration. |
| `env/core` | `src/features/env/core/` | Tier-aware environment loading and listener-port resolution — inlined during the migration. |

```text
                        +-----------------+
                        |    app-shell    |
                        | navigation only |
                        +-----------------+

   +--------+  customer/supplier   +---------------------+
   |  home  | -------------------> | personal-projects   |
   +--------+                      +---------------------+
       ^                                     ^
       | conformist                          | conformist
       |            +----------+             |
       +----------- |  search  | ------------+
                    +----------+
                         |  conformist
                         v
                    +----------+
                    |    cv    |
                    +----------+
```

`search` is a conformist to the three components it indexes: it reads their content as they publish it and imposes no shape on them, so adding a field to the CV record cannot break search's contract, only its results.

## Data Stores

One, and it is not a database.

| Store | Location | Description |
| --- | --- | --- |
| CV record | `src/features/cv/core/data.ts` | The single authoritative CV record in this repository. It is TypeScript compiled into the bundle, not data fetched at runtime. |

That it is the _only_ one is load-bearing rather than incidental. Before the migration this repository also kept a `cv/` directory at its root, and two records of the same career are a guarantee of eventually publishing the stale one. The migration deleted `cv/` and left this store as the single source.

## Boundaries

**Process.** One at build time: Next.js reads the TypeScript sources and emits static output. One at test time: the E2E adapter starts a server process and speaks to it over a local port. None at request time — a served page runs no code of ours.

**Network.** The visitor reaches Vercel's edge over HTTPS and nothing else. The application makes no outbound calls: no analytics, no API, no font or asset fetch from a third party. Fonts are bundled.

**Trust.** There is no privilege boundary inside the application, because there is no authentication, no session, and no user input that reaches a server. The one boundary that matters is deployment: `vercel.json` gates every build on the `prod-wahidyankf-www` branch, so `main` moving does not move the site.

**Environment.** `env/core` resolves which tier's env file to read and refuses to fall through to a developer's local file when the tier is anything other than `local`. The E2E adapter pins its tier explicitly for the same reason.

## Behavior Traceability

| Feature | Component it exercises |
| --- | --- |
| [accessibility.feature](behavior/accessibility.feature) | `app-shell`, and every page it frames |
| [cv.feature](behavior/cv.feature) | `cv` |
| [env-loader.feature](behavior/env-loader.feature) | `env/core` |
| [home.feature](behavior/home.feature) | `home` |
| [personal-projects.feature](behavior/personal-projects.feature) | `personal-projects` |
| [port-resolver.feature](behavior/port-resolver.feature) | `env/core` |
| [responsive.feature](behavior/responsive.feature) | `app-shell` |
| [search.feature](behavior/search.feature) | `search` |
| [static-filterable-routes.feature](behavior/static-filterable-routes.feature) | `personal-projects`, at the routing boundary |
| [theme.feature](behavior/theme.feature) | `ui/shell` |
| [tier-env-loading.feature](behavior/tier-env-loading.feature) | `env/core` |
