# Specification Fixtures

Fixture corpora this repository runs its validators against. Unlike the Gherkin under `specs/apps/`, nothing here
describes this repository's own behaviour: each corpus is a set of synthetic inputs with a declared expected result.

## Directory Map

- [`plan-structure/`](plan-structure/README.md) — the shared plan-structure corpus every implementation of the
  [plan validator contract](../../repo-governance/conventions/plan-validator-contract.md) runs against.

## Adopted, Not Authored

`plan-structure/` was adopted from the catalog rather than written here, and this copy is owned here: nothing pins it
upstream and nothing obliges this repository to notice that the catalog moved. Its bytes still matter — a comparison
between two validators is only meaningful over the same input — so it is exempt from this repository's formatter, its
Markdown link rule, and its directory-index rule, each exemption recorded where that rule lives.

Verify it before trusting a run against it:

```sh
cd specs/fixtures/plan-structure && shasum -a 256 -c SHA256SUMS
```

That is a local integrity check over this copy, never a comparison against the catalog. A mismatch is a stop, not a
warning.
