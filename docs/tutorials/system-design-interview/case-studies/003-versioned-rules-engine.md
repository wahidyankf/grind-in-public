---
tldr: "Design safe authoring, compilation, publication, and deterministic evaluation of versioned rules."
when_to_use: "Use to practise configuration planes, immutable versions, caching, and explainability."
---

# Case 003: Versioned Rules Engine

Design a multi-tenant engine in which policy authors create rules and online services evaluate a published policy
against an event and entity context.

### Functional requirements

| ID   | Requirement                                                                                           |
| ---- | ----------------------------------------------------------------------------------------------------- |
| FR-1 | Create drafts, validate syntax/types, test fixtures, review, and publish an immutable policy version. |
| FR-2 | Evaluate one named published version deterministically and return actions plus reason codes.          |
| FR-3 | Roll back the active pointer to a prior compatible version without deleting history.                  |
| FR-4 | Simulate a candidate version over historical inputs without producing business effects.               |
| FR-5 | Record policy, compiler, input-schema, and reference-data versions for every result.                  |

### Non-functional requirements

| ID    | Requirement                                                                        |
| ----- | ---------------------------------------------------------------------------------- |
| NFR-1 | Evaluate 5,000 rules within 40 ms p99, excluding external feature retrieval.       |
| NFR-2 | Publish globally within 60 seconds while an evaluation uses exactly one version.   |
| NFR-3 | Prevent arbitrary code execution and bound CPU, memory, recursion, and regex work. |
| NFR-4 | Provide 99.99% evaluation availability from last-known-good configuration.         |
| NFR-5 | Isolate policy visibility and authoring authority by tenant and role.              |

## Assumptions and exclusions

Rules use a deliberately limited declarative language. Full general-purpose scripting and model training are out of
scope.

## Interview prompts

1. Which work happens at authoring, compile, publish, and evaluation time?
2. How does a worker atomically switch versions while in-flight requests finish?
3. Which indexes or decision structures avoid scanning every rule?
4. How do simulation, rollout, rollback, audit, and bad-version containment work?

Solve before reading [the worked solution](../solutions/003-versioned-rules-engine.md).
