# Badak Mini E2E

This dedicated Nx application executes the canonical [Badak Mini behavior corpus](../../specs/apps/badakmini-cli/behavior/) through the built `badak-mini` executable. It owns no `.feature` files, unit layer, or numeric coverage threshold. Its test package imports Godog directly and owns the process adapter's scenario lifecycle and `Given`, `When`, and `Then` registrations; it does not execute through the owner application's BDD suite.

`test:e2e` first builds its owner application and runs the owner behavior-completeness gate. That gate compares both the owner and E2E direct Godog bindings with the same recursive catalog, including undefined, ambiguous, and unused binding failures. The E2E target passes the absolute executable path through `BADAKMINI_BIN`; the process driver rejects unset, relative, missing, and non-executable paths. Each scenario uses an isolated local filesystem and Git fixture, observes only command arguments, exit status, standard output, standard error, and filesystem effects, and has a 30-second command timeout.

Run these targets from the repository root:

```sh
npm exec nx -- run badakmini-cli-e2e:typecheck
npm exec nx -- run badakmini-cli-e2e:lint
npm exec nx -- run badakmini-cli-e2e:test:coverage:behavior
npm exec nx -- run badakmini-cli-e2e:test:e2e
npm exec nx -- run badakmini-cli-e2e:test:quick
```

Unit and integration tests remain in the owner application: E2E exists solely to prove its public process boundary.
