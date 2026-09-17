---
tldr: "Uses HIPPO as the checksum-pinned CPU-and-memory admission boundary while preserving safe parallel work."
when_to_use: "Use before running or wiring builds, tests, repository checks, services, or other compute-heavy work."
---

# Resource-Aware Development

This repository consumes [HIPPO](https://github.com/wahidyankf/hippo) through the root checksum-pinned `./hippo`
bootstrap. HIPPO implementation, specifications, conformance, and releases remain upstream; never vendor or fork them
here.

## Admission and Parallelism

Run each independent compute-bearing node through one outer guard. Nodes may enter concurrently; HIPPO atomically admits
CPU-and-memory reservations only while shared capacity and current host pressure permit. Serialize only for a
dependency, shared or tracked output, an indivisible mutation, a runner limit, or a demonstrated correctness race. A
repository or project boundary alone is not a serial edge.

All `ephemeral`, `service`, and `transactional` owners consume capacity. Automatic balanced, constrained, and minimal
requests use four, two, and one shares of safe capacity. Explicit reservations cannot fall below one CPU or 256 MiB.
Admission is strict FIFO; a fitting vector can still wait under pressure.

Canonical execution is:

```sh
rtk ./hippo run --class ephemeral --resource-tier standard --disk-path . -- npm exec -- nx run -p <project> -t <target>
```

The wrapper maps a fixed admitted CPU allocation only to `NX_PARALLEL` and `GOMAXPROCS`. A missing value receives the
allocation, a lower positive value survives, and a larger one is clamped. Inner commands inherit the fixed session;
never add another outer guard or product mapping.

## Classes, Supervision, and Recovery

Use `ephemeral` for restartable builds, tests, and reads; `service` for restartable long-running development; and
`transactional` for authorized indivisible mutations or tracked-output writes. Never change class merely to gain entry.
Choose `light` for narrow static checks, `standard` for ordinary checks and writers, and `heavy` for complete gates,
full builds, full suites, and browser suites. Schema 3 keeps one FIFO waiter and launches each payload at most once.
Ordinary critical pressure sheds the newest eligible ephemeral, then service; a transaction is eligible last only at the
emergency floor. Only the owning guard signals and reaps its child group before releasing its reservation.

- Exit `73`: free storage safely, then retry.
- Exit `75`: inspect the receipt or outcome. Requeue only `never-started`; pressure-shed, storage-shed, and
  `started-safety-stop` require payload-specific recovery. Never duplicate or loop retries.
- Exit `78`: correct configuration, reservation, mapping, or strict-profile planning before retrying.

Never bypass HIPPO, weaken a gate, delete possibly live state, or raise mapped concurrency.

## Enforcement

The contract above was prose only, and prose did not hold: an unguarded Nx fan-out in a sibling repository exhausted
host memory and forced a restart. HIPPO cannot shed work it was never told about, so pressure went critical while the
scheduler still reported `normal`.

[`.claude/hooks/require-hippo-boundary.sh`](../../.claude/hooks/require-hippo-boundary.sh) refuses a compute-bearing
command carrying no outer guard, before the process spawns. All three harnesses bind it, byte-identical to every other
consuming repository's copy so a hardening fix cannot land in one and quietly miss the rest.

It decides only whether a guard is present, never which class is right; class choice needs intent and stays a judgment.
Verbs match only in command position, so searching for a verb string is not refused — a guard that blocks ordinary
searching is one that gets switched off.

## Consumer Integrity and Evidence

`hippo.lock` pins the public release version, source commit, and SHA-256 for each supported macOS/Linux and amd64/arm64
asset. The wrapper validates archive identity, revalidates cached bytes, publishes atomically, and retains a bounded
cache. `hippo.local.json.example` documents schema 3; any machine policy copied to ignored `hippo.local.json` cannot
weaken upstream floors. A contained worktree with no local copy uses the primary checkout's ignored policy.

Use the per-user shared root across checkouts. Override `HIPPO_ROOT` only for isolated tests or a separately
administered domain. Corrupt or unverifiable coordination state fails closed; follow upstream recovery guidance instead
of deleting state based on diagnostic PIDs.

Evidence contains capacity and process-health measurements, never file contents, command arguments, origins,
credentials, or user data. Verify bootstrap and pressure behaviour with isolated synthetic releases, state, and fake
pressure; never create real host pressure to prove shedding. Scheduled Linux/macOS smoke verifies identity, schema,
mappings, shared-root policy, and cleanup while ordinary hosted project jobs retain runner-native limits.

`hippo.identity.json` labels this repository in live status and the bounded 30-day history. Identity discovery works
from nested paths and an explicitly authorized contained `worktrees/<task>` checkout; add privacy-safe
`--tag checkout=worktree --tag plan=<slug>` values per run. Use `./hippo status`,
`./hippo watch --source grind-in-public`, and `./hippo history --since 30d --source grind-in-public` directly. Every
checkout uses the shared default root; set `HIPPO_ROOT` only for isolated tests.
