# forum-be-python

A mini forum REST API, built as a small pilot of a Python backend in this workspace. People register and log in with a
username and password, then create topics, threads inside them, and replies that nest to any depth. It is FastAPI on
PostgreSQL with raw SQL, checked with pyright (strict) and pytest. It has no frontend, no other consumers, and no
deployment.

## Prerequisites

- [uv](https://docs.astral.sh/uv/) with Python 3.13, which `.python-version` pins. The first `uv run` installs the
  locked dependencies from `uv.lock`.
- Docker, used only to run PostgreSQL from `compose.yaml`. The container binds `127.0.0.1:54329`, trusts every
  connection, and keeps its data in tmpfs, so it holds throwaway development and test data only.

## Commands

Run from the repository root:

```sh
rtk ./hippo run --class ephemeral --resource-tier standard --disk-path . -- npm exec -- nx run -p forum-be-python -t test:quick
rtk ./hippo run --class ephemeral --resource-tier standard --disk-path . -- npm exec -- nx run -p forum-be-python -t test:e2e
rtk ./hippo run --class ephemeral --resource-tier standard --disk-path . -- npm exec -- nx run -p forum-be-python -t dev
```

| Target       | What it does                                                                                              |
| ------------ | --------------------------------------------------------------------------------------------------------- |
| `typecheck`  | pyright in strict mode over `src/`                                                                        |
| `test:unit`  | pytest over `src/tests/unit`: pure functions, no database                                                 |
| `test:e2e`   | starts PostgreSQL through `db:up`, then runs `src/tests/e2e`: the whole app in process on a real database |
| `test:quick` | `typecheck`, then `test:unit`                                                                             |
| `dev`        | serves the API with reload on `http://127.0.0.1:3202`; run `db:up` first                                  |
| `db:up`      | starts the PostgreSQL container and waits until it is healthy                                             |
| `db:down`    | stops and removes the container                                                                           |

The database-backed suite is `test:e2e`, not an integration target, because the
[quality gates](../../repo-governance/development/quality-gates.md) reserve integration tests for network-free code and
these tests reach a real PostgreSQL over loopback.

## Configuration

| Variable               | Default                                          | Purpose                                                                 |
| ---------------------- | ------------------------------------------------ | ----------------------------------------------------------------------- |
| `FORUM_DATABASE_URL`   | `postgresql://postgres@127.0.0.1:54329/forum`    | Database the API connects to                                            |
| `FORUM_TEST_ADMIN_URL` | `postgresql://postgres@127.0.0.1:54329/postgres` | Admin connection the E2E suite uses to create and drop its own database |

The E2E suite creates a uniquely named `forum_test_*` database per run, refuses a non-loopback host, and drops only that
database. The schema in `src/forum_be_python/schema.sql` is applied idempotently at startup.

## API

Interactive documentation is served at `/docs` while `dev` runs. Reading is public; writing needs a bearer token from
`/auth/login`. Only the author of a topic, thread, or post may change or delete it.

| Method             | Path                   | Purpose                                                      |
| ------------------ | ---------------------- | ------------------------------------------------------------ |
| GET                | `/health`              | Liveness plus a database round trip                          |
| POST               | `/auth/register`       | Create a user; 409 when the username is taken                |
| POST               | `/auth/login`          | Exchange credentials for a 24-hour bearer token              |
| GET                | `/auth/me`             | The signed-in user                                           |
| POST, GET          | `/topics`              | Create a topic; list topics                                  |
| GET, PATCH, DELETE | `/topics/{id}`         | Read, partially update, or delete a topic                    |
| POST, GET          | `/topics/{id}/threads` | Start a thread with its opening post; list a topic's threads |
| GET, PATCH, DELETE | `/threads/{id}`        | Read a thread with its reply tree; rename; delete            |
| POST               | `/threads/{id}/posts`  | Reply to the thread or, with `parent_id`, to another post    |
| GET, PATCH, DELETE | `/posts/{id}`          | Read a post with its subtree; edit; delete with its replies  |

## Design notes

- Passwords are hashed with the standard library's scrypt. A token is random and opaque, and only its SHA-256 is stored,
  so a database leak does not leak usable tokens. An unknown username and a wrong password get the same 401.
- A reply's parent must live in the same thread, which a composite foreign key enforces. Deleting a post cascades to its
  replies.
- Endpoints are synchronous and share a `psycopg_pool` connection pool.

## Scope

This is the workspace's
[recorded pilot](../../repo-governance/development/testing-policy/tooling.md#recorded-deviations), kept lighter than an
owner application on purpose. It has no Gherkin corpus, no `specs/` C4 model, no dedicated E2E project, and no `build`,
`lint`, integration, or coverage targets, so no 99% coverage floor either. Its unit and E2E suites are written
test-first, and the manual `curl` proof for the public API still applies to every change.

## Structure

- `src/forum_be_python/` holds the application; `main.py` builds the app and each `routes_*.py` owns one resource.
- `src/tests/unit/` and `src/tests/e2e/` hold the two suites; `src/tests/e2e/conftest.py` owns the per-run database.
- `pyproject.toml` and `uv.lock` pin the toolchain; `compose.yaml` runs PostgreSQL.
