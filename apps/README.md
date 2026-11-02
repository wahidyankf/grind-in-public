# Applications

This directory contains runnable workspace applications. Each application owns its source, tests, project configuration,
and application-specific documentation; shared packages belong in [`libs/`](../libs/README.md).

## Current Applications

- `edutech-be` — the FastAPI backend (Python 3.13, uv, Postgres), with auth endpoints.
- `edutech-web` — the Next.js frontend for the edutech backend.
- `edutech-web-e2e` — the Playwright suite for `edutech-web`.
- `knowledge-base-be` — the Fastify REST backend (Node 24, TypeScript, Postgres), with health and auth endpoints.
- `knowledge-base-web` — the React frontend (TanStack Start in SPA mode, TanStack Router and Query) for the
  knowledge-base backend.
- `knowledge-base-web-e2e` — the Playwright suite for `knowledge-base-web` and its backend.
