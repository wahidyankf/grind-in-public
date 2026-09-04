# Application Specifications

Each application in `apps/` that carries specifications has a directory here under the same name, holding its as-built
C4 model and its executable Gherkin corpus. A directory appears here when the application it mirrors gains a
specification, not before; see the [specs policy](../../repo-governance/development/specs-policy.md) for when one is
required.

## Directory Map

- [Badak Mini](badakmini-cli/README.md) — the repository validation CLI: its C4 model and its five-feature behaviour
  corpus.
- [wahidyankf-www](wahidyankf-www/README.md) — the personal portfolio and CV site: its C4 model and twelve-feature
  behaviour corpus, bound by owner Unit/Integration adapters and a dedicated browser E2E project.
