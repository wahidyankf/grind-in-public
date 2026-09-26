---
tldr: "A control plane compiles immutable policies into bounded decision artifacts consumed atomically by evaluators."
when_to_use: "Use after attempting Case 003 to compare configuration publication, evaluation, and rollback."
---

# Solution 003: Versioned Rules Engine

## Requirement traceability

| Requirements       | Design response                                                                                          |
| ------------------ | -------------------------------------------------------------------------------------------------------- |
| FR-1, FR-3         | Draft/review workflow publishes immutable artifact; active pointer can move to prior compatible version. |
| FR-2, NFR-1, NFR-3 | Typed DSL compiles to bounded indexed decision plan; no user code executes.                              |
| FR-4               | Simulation reads historical snapshots and writes a separate comparison dataset.                          |
| FR-5, NFR-2        | Result records artifact/compiler/schema/reference versions; evaluator pins one snapshot per request.     |
| NFR-4, NFR-5       | Signed last-known-good artifacts, tenant-scoped storage/cache, and role-separated authoring.             |

## Control and data planes

```text
author -> draft -> validate -> fixture tests -> review -> compile -> immutable artifact
                                                                  |
                                                           active pointer
                                                                  |
request -> evaluator -> pinned in-memory artifact -> result + reasons + versions
```

Compilation validates field types, rejects cycles, expands reference sets, detects contradictory/unreachable rules, and
builds indexes by event type and cheap predicates. Each operation has a static cost; policy publication rejects an
artifact above tenant limits for rules, depth, regex complexity, and reference cardinality.

## Atomic version activation

Evaluators watch a signed manifest. They download, hash, parse, and warm the candidate off-path, then atomically swap an
immutable pointer. An in-flight request keeps its acquired object; new requests see the new one.

```text
time ------->
request A: [----------- version 11 -----------]
pointer:               v11 | v12
request B:                     [-- version 12 --]
```

Publication succeeds only after a quorum of target cells reports artifact readiness; lagging cells keep version 11 and
placement routes a request to one cell. The result always declares the executed version.

## Evaluation plan

1. Validate event schema once.
2. Select rules indexed by event type, tenant policy, and required fields.
3. Evaluate cheap exact/range predicates before expensive set or text operations.
4. Short-circuit only when declared policy semantics permit it.
5. Accumulate stable reason codes and actions with deterministic conflict resolution.

Store the authored representation separately from the compiled artifact. Compiler version is part of artifact identity,
allowing reproducibility after compiler changes.

## Failure and rollout

- control-plane outage: evaluate with last-known-good; authoring/publishing pauses;
- invalid signature/hash: reject candidate and page publication owner;
- memory pressure: admission rejects oversized artifact before rollout;
- bad semantics: shadow historical/live traffic, canary tenants, compare action distribution and reason codes, then
  expand; rollback moves the pointer without mutating history;
- cache miss after restart: startup probe remains false until the assigned tenant set has safe baseline artifacts.

## Alternatives rejected

Executing tenant-provided Python makes cost, security, determinism, and upgrades unbounded. Reading rules from a remote
database during every evaluation violates latency and availability goals. Mutating a live policy in place destroys
reproducibility and safe rollback.
