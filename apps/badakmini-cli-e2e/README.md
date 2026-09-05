# Badak Mini Process E2E

This dedicated Nx project executes the owner's canonical [behaviour corpus](../../specs/apps/badakmini-cli/behaviours/)
through the built `badak-mini` executable. It owns no feature files, unit layer, integration layer, or numeric coverage
threshold.

The project imports Godog directly and owns the public-process scenario lifecycle and registrations. Its driver rejects
an unset, relative, missing, or non-executable `BADAKMINI_BIN`, gives each scenario an isolated filesystem and Git
fixture, and observes command arguments, exit status, standard output, standard error, and filesystem effects.

`test:coverage:behaviour:e2e` statically proves the E2E bindings against the owner's recursively discovered corpus.
`test:coverage:behaviour` delegates to the owner aggregate. Runtime E2E depends on that aggregate and the owner build.
Unit implements every scenario. Integration and process E2E may each be omitted through their layer-specific exemption;
both tags may annotate one scenario when each has its own valid justification and substantive alternative proof.

Run from the repository root:

```sh
npm exec nx -- run badakmini-cli-e2e:test:quick
npm exec nx -- run badakmini-cli-e2e:test:e2e
```
