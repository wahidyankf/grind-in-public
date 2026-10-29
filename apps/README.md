# Applications

This directory contains runnable workspace applications. Each application owns its source, tests, project configuration,
and application-specific documentation; shared packages belong in [`libs/`](../libs/README.md).

## Current Applications

- `edutech-be` — the FastAPI backend (Python 3.13, uv, Postgres), with auth endpoints.
- `edutech-web` — the Next.js frontend for the edutech backend.
- `edutech-web-e2e` — the Playwright suite for `edutech-web`.
