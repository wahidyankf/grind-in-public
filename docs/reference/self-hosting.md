---
tldr: "Explains how to run the wahidyankf-www site, how its port and environment tiers work, and where secrets belong."
when_to_use: "Use when you run, host, or configure the wahidyankf-www site outside local development."
---

# Self-hosting

This guide explains how to run the `wahidyankf-www` site yourself and how its configuration works. Read all of it before
you change config or secrets. It has no short summary on purpose: a wrong setting can expose the site.

## Run the site

- Build: `npx nx build wahidyankf-www`.
- Start: `npx nx start wahidyankf-www`. The site listens on port 3201 by default.

## Port

The port is chosen in this order:

1. The `--port` flag, if you pass one. It wins over everything else.
2. `WAHIDYANKF_WWW_PORT`, if it is set.
3. 3201, the built-in default.

A bare `PORT` variable is not read. The wrapper in `scripts/next-with-port.mjs` resolves the port and passes it to
Next.js.

## Environment tiers

`APP_ENV` picks one env file, `.env.<APP_ENV>`. When `APP_ENV` is unset, the tier is `local`.

- Values already in the process environment win over values in the file.
- A missing file is not an error. This is the normal case in CI.
- At any tier except `local`, the load fails if a stray `.env`, `.env.production`, or `.env.local` sits next to the tier
  file. Next.js loads those files on its own, which would break the one-file rule.

## Secrets

The app declares no secrets today. Its config schema, `apps/wahidyankf-www/src/env.ts`, is empty.

When you add a secret:

- Declare it in `env.ts` as a required variable, so the app fails fast in production if it is missing or invalid.
- Document it in this file: its name, what it does, and what happens if it is missing.
- Never commit a real value. Git ignores `.env` and `.env.*` files.
