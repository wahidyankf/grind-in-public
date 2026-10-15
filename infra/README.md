# Infrastructure

This directory holds local infrastructure that applications depend on, such as databases. Each stack gets its own
folder, so a new stack can be added without changing the others.

## Directory Map

- [edutech/compose.yaml](edutech/compose.yaml) — the local Postgres used by `edutech-be`. Start it with
  `npx nx db:up edutech-be`.
