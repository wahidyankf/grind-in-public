---
tldr: "Derives normalized tables from dependencies, then denormalizes only behind an explicit consistency contract."
when_to_use: "Use when turning business concepts into relational schemas."
---

# Modelling and Normalization

Start with facts and functional dependencies. In an account record, `account_id -> customer_id, currency, status`;
customer name depends on `customer_id`, not the account. Repeating names in every transaction introduces update, insert,
and delete anomalies.

```text
raw record
  |
  +-- customer facts --------> customers
  +-- account facts ---------> accounts
  +-- event facts -----------> transactions
  +-- many-to-many relation -> case_alerts
```

First normal form gives atomic values; second removes dependency on part of a composite key; third removes transitive
dependencies between non-key attributes. Boyce-Codd normal form strengthens the determinant rule. Treat these as
reasoning tools, not a ritual: the goal is one authoritative location per fact.

Denormalize when a measured read path needs it and define:

- authority: which normalized row wins;
- propagation: transaction, outbox, or rebuildable batch;
- freshness: a measurable bound;
- repair: how drift is detected and rebuilt;
- failure: what readers see during lag.

`jsonb` is useful for variable metadata, but core join keys, money, status, and timestamps deserve typed columns and
constraints. A document column is not permission to avoid modelling.
