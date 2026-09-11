# Specification Fixtures

Fixture corpora this repository runs its validators against. Unlike the Gherkin under `specs/apps/`, nothing here
describes this repository's own behaviour: each corpus is a set of synthetic inputs with a declared expected result.

## Directory Map

- [`plan-structure/`](plan-structure/README.md) — the shared plan-structure corpus every implementation of the
  [plan validator contract](../../repo-governance/conventions/plan-validator-contract.md) runs against.

## Vendored, Not Authored

`plan-structure/` is consumed by digest, not maintained here. Its bytes are identical in every repository that carries
it, and that identity is the whole point: two validators agreeing means they agreed on the same input. It is therefore
exempt from this repository's formatter, its Markdown link rule, and its directory-index rule, each exemption recorded
where that rule lives.

Verify it before trusting a run against it:

```sh
cd specs/fixtures/plan-structure && shasum -a 256 -c SHA256SUMS
```

A mismatch is a stop, not a warning.
