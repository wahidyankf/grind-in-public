---
tldr: "Defines automated and manual proof required when a change affects a public API."
when_to_use: "Use for REST, GraphQL, webhook, RPC, streaming, subscription, or other public API changes."
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
