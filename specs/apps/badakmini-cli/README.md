# Badak Mini Specifications

Badak Mini's canonical executable behavior lives in [behavior/](behavior/). Every feature is run unchanged by its unit, local-integration, and public-process E2E adapters.

| Feature | Covers |
| --- | --- |
| [CLI contract](behavior/cli-contract.feature) | Help without repository discovery. |
| [Instruction size](behavior/instruction-size.feature) | Valid and oversized governance documents. |
| [Markdown links](behavior/markdown-links.feature) | Valid and broken tracked local links. |
| [Capability parity](behavior/capability-parity.feature) | Matching and missing harness capabilities. |
| [Rule change](behavior/rule-change.feature) | Staged notices and pre-edit hook notices. |

The owner's `test:coverage:behavior` target verifies recursive discovery, expression binding, and adapter parity. Its co-located E2E package uses the same corpus and owns no features of its own.
