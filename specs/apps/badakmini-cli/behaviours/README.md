# Badak Mini Behaviour

The canonical Gherkin corpus for Badak Mini. Every feature here is run unchanged by the unit, local-integration, and
public-process E2E adapters, so a scenario added to this directory has to be bound in all three or
`test:coverage:behaviour` fails.

Scenario shape and cardinality belong to the [specs policy](../../../../repo-governance/development/specs-policy.md);
how a scenario reaches a test belongs to the [TDD policy](../../../../repo-governance/development/tdd-policy.md). The
traceability table against the C4 model lives in [architecture.md](../architecture.md).

## Directory Map

- [cli-contract.feature](cli-contract.feature) — help output without repository discovery.
- [instruction-size.feature](instruction-size.feature) — valid and oversized governance documents.
- [markdown-links.feature](markdown-links.feature) — valid and broken tracked local links.
- [capability-parity.feature](capability-parity.feature) — matching and missing harness capabilities.
- [rule-change.feature](rule-change.feature) — staged notices and pre-edit hook notices.
