---
tldr: "Defines automated and manual proof required when a change affects a public API, and its hand-run collections."
when_to_use:
  "Use for REST, GraphQL, webhook, RPC, streaming, subscription, or other public API changes, or a request collection."
---

# API Testing

Apply this standard only when a change can affect an externally reachable API.

## Automated Proof

- Unit tests prove business rules, validation, authorization, mapping, and error behaviour through injected dependencies
  without OS or network access.
- Integration tests exercise routing, parsing, schema validation, serialization, middleware, and isolated local stores
  in-process without opening a network listener.
- E2E tests exercise representative operations through the exact served public origin.
- Contract assertions cover method or operation, path, headers, content type, payload or variables, status, response
  shape, declared errors, and observable side effects.
- Cover success, invalid input, expected failure, authentication or authorization, and promised idempotency. For
  GraphQL, assert both HTTP and `data`/`errors`; HTTP `200` alone is not success.

## Manual Public-Boundary Proof

Before completing an API-affecting change, invoke every affected HTTP operation with `curl` against the exact isolated
served origin. Exercise successful and materially changed error or authorization paths. Use only synthetic state under
[test-data isolation](test-data-isolation.md).

Record the redacted command shape, exact non-sensitive origin, operation, observed status and response shape, side
effect, and pass/fail. Never record secrets or private payloads. Use a protocol-capable client after `curl` for a
subscription, WebSocket, or stream whose lifecycle cannot be proved by its handshake. When no API is affected, record
`API impact: none` rather than running an unrelated probe.

## Request Collections

A project with a public API may keep a collection that people run by hand to sanity-check the API and explore it. Write
it in the [Kulala](https://github.com/mistweaverco/kulala.nvim) `.http` format under the project's `http/` directory,
beside a tracked `http-client.env.json` that holds only shareable values. Keep a secret or a personal override only in
`http-client.private.env.json`, which Git ignores at any depth. It mirrors the tracked file's structure, its values win,
and Kulala saves fetched auth tokens in it, so never commit it.

A collection is a manual aid, not automated proof, and it never replaces the `curl` proof above. It targets a local
isolated origin with synthetic state under [test data isolation](test-data-isolation.md). An API change should update
the affected requests in the same change.
