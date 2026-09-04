# Badak Mini Specifications

Badak Mini's canonical as-built boundaries live in [architecture.md](architecture.md), and its executable behaviour
lives in [behaviours/](behaviours/). Every feature is run unchanged by its unit, local-integration, and public-process
E2E adapters.

## Directory Map

- [Architecture](architecture.md) is the current as-built C4 model, its constraints, and behaviour traceability.
- [Behaviour](behaviours/README.md) contains the canonical executable Gherkin corpus.

| Feature                                                   | Covers                                     |
| --------------------------------------------------------- | ------------------------------------------ |
| [CLI contract](behaviours/cli-contract.feature)           | Help without repository discovery.         |
| [Instruction size](behaviours/instruction-size.feature)   | Valid and oversized governance documents.  |
| [Markdown links](behaviours/markdown-links.feature)       | Valid and broken tracked local links.      |
| [Capability parity](behaviours/capability-parity.feature) | Matching and missing harness capabilities. |
| [Rule change](behaviours/rule-change.feature)             | Staged notices and pre-edit hook notices.  |

The owner's `test:coverage:behaviour` target verifies recursive discovery, expression binding, and adapter parity. The
dedicated `apps/badakmini-cli-e2e` project uses the same corpus and owns no features of its own.
