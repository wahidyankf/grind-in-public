# Badak Mini Behaviour

The canonical Gherkin corpus for Badak Mini. Unit implements every scenario. A scenario may omit Integration, E2E, or
both only when each omitted boundary fundamentally cannot express the behaviour, using independently documented
`@integration-exempt` and `@e2e-exempt` tags from the repository
[BDD policy](../../../../repo-governance/development/behaviour-driven-development-policy.md).

Scenario shape and cardinality belong to the [specs policy](../../../../repo-governance/development/specs-policy.md);
how a scenario reaches a test belongs to the [TDD policy](../../../../repo-governance/development/tdd-policy.md). The
traceability table against the C4 model lives in [architecture.md](../architecture.md).

## Directory Map

- [cli-contract.feature](cli-contract.feature) — help output without repository discovery.
- [instruction-size.feature](instruction-size.feature) — valid and oversized governance documents.
- [markdown-links.feature](markdown-links.feature) — valid and broken tracked local links.
- [capability-parity.feature](capability-parity.feature) — canonical harness content and native adapter divergence.
- [rule-change.feature](rule-change.feature) — staged notices and pre-edit hook notices.
