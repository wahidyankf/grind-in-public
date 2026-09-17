---
tldr: "Uses HIPPO as the checksum-pinned CPU-and-memory admission boundary while preserving safe parallel work."
when_to_use: "Use before running or wiring builds, tests, repository checks, services, or other compute-heavy work."
---

# Resource-Aware Development

Consume [HIPPO](https://github.com/wahidyankf/hippo) through root checksum-pinned `./hippo`; its implementation,
specifications, and releases stay upstream.

## Admission and Parallelism

Give each independent compute node one outer guard. HIPPO admits CPU-and-memory reservations when shared capacity and
host pressure permit. Serialize only for dependencies, shared output, indivisible mutation, runner limits, or proven
races—not repository boundaries.

Every class consumes capacity. Automatic balanced, constrained, and minimal requests use four, two, and one safe shares.
Explicit reservations require at least one CPU and 256 MiB. Admission is FIFO and pressure-aware.

Canonical execution is:

```sh
rtk ./hippo run --class ephemeral --resource-tier standard --disk-path . -- npm exec -- nx run -p <project> -t <target>
```

The wrapper maps admitted CPU only to `NX_PARALLEL` and `GOMAXPROCS`: missing values receive it, smaller values survive,
and larger ones clamp. Never nest guards or add product mappings.

## Classes, Supervision, and Recovery

Use `ephemeral` for restartable work, `service` for long-running development, and `transactional` for indivisible
mutations or tracked-output writes. Never change class to gain entry. Use `light` for narrow static checks, `standard`
for ordinary work, and `heavy` for full builds, suites, and gates. Schema 3 is FIFO and at-most-once. Critical pressure
sheds eligible ephemeral, then service, with transactions last at the emergency floor. Only the owner reaps its group.

- Exit `73`: free storage safely, then retry.
- Exit `75`: requeue only when a new schema-1 receipt proves `never-started`; pressure-shed, storage-shed,
  `started-safety-stop`, and child-owned `75` require payload-specific recovery.
- Exit `76`: never retry. Inspect `./hippo status`, drain or upgrade the incompatible peer, then retry the original
  command. A legacy client without distinct exit `76` can still report the mismatch as `75`, but has no qualifying
  receipt.
- Exit `78`: correct configuration, reservation, mapping, or strict-profile planning before retrying.
- Exit `1`: diagnose malformed shared state or Hippo-owned post-launch cleanup; never classify it as capacity.

Child codes pass through, including `75` and `76`; task-failed evidence without a new `never-started` receipt keeps them
child-owned.

Never bypass HIPPO, weaken a gate, delete possibly live state, or raise mapped concurrency.

## Enforcement

An unguarded sibling Nx fan-out once exhausted memory and forced a restart. HIPPO cannot shed unknown work.

[The boundary hook](../../.claude/hooks/require-hippo-boundary.sh) rejects unguarded compute before spawn. All three
harnesses bind the byte-identical shared consumer hook.

The hook checks presence, not class judgment, and matches compute verbs only in command position.

## Consumer Integrity and Evidence

`hippo.lock` pins release, commit, and each platform asset's SHA-256. The wrapper revalidates bytes, publishes
atomically, and bounds its cache. `hippo.local.json.example` documents schema 3; ignored machine policy cannot weaken
upstream floors. A worktree without it inherits the primary checkout's policy.

The lock pins executable identity, not governance semantics. Before changing consumer behaviour, read the Hippo
repository at the commit in `hippo.lock`, especially its exit-code and recovery references, then reconcile this rule,
Gherkin, and harness checks with the capabilities that commit actually provides. Never infer capability from SemVer
ordering or copy a release number into the rule.

Use the shared per-user root. Override `HIPPO_ROOT` only for isolated tests. Unverifiable coordination state fails
closed; follow upstream recovery guidance rather than deleting state from diagnostic PIDs.

Evidence contains capacity and process health, never contents, arguments, origins, credentials, or user data. Test
pressure only with isolated synthetic state. Scheduled Linux/macOS smoke verifies identity, schema, mappings, root, and
cleanup; ordinary hosted jobs retain runner-native limits.

`hippo.identity.json` labels this repository in live status and the bounded 30-day history. Identity discovery works
from nested paths and an explicitly authorized contained `worktrees/<task>` checkout; add privacy-safe
`--tag checkout=worktree --tag plan=<slug>` values per run. Use `./hippo status`,
`./hippo watch --source grind-in-public`, and `./hippo history --since 30d --source grind-in-public` directly. Every
checkout uses the shared default root; set `HIPPO_ROOT` only for isolated tests. Raw evidence rolls for seven days;
compacted daily summaries roll for 30 days under byte caps.
