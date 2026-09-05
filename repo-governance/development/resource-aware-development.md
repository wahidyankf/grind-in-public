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
rtk ./hippo run --class ephemeral --disk-path . -- npm exec -- nx run -p <project> -t <target>
```

The wrapper maps a fixed admitted CPU allocation only to `NX_PARALLEL` and `GOMAXPROCS`. A missing value receives the
allocation, a lower positive value survives, and a larger one is clamped. Inner commands inherit the fixed session;
never add another outer guard or product mapping.

## Classes, Supervision, and Recovery

Use `ephemeral` for restartable builds, tests, and reads; `service` for restartable long-running development; and
`transactional` for authorized indivisible mutations or tracked-output writes. Never change class merely to gain entry.
Under critical pressure HIPPO sheds the newest eligible ephemeral, then the newest service only if none is eligible; it
never sheds a transaction. Only the owning guard signals and reaps its child group before releasing its reservation.

- Exit `73`: free storage safely, then retry.
- Exit `75`: let the capacity, FIFO, compatibility, or pressure deferral finish before retrying the same invocation;
  never duplicate or loop retries.
- Exit `78`: correct configuration, reservation, mapping, or strict-profile planning before retrying.

Never bypass HIPPO, weaken a gate, delete possibly live state, or raise mapped concurrency.

## Consumer Integrity and Evidence

`hippo.lock` pins the public release version, source commit, and SHA-256 for each supported macOS/Linux and amd64/arm64
asset. The wrapper downloads outside the repository, validates the archive and embedded identity, revalidates cached
bytes before execution, publishes atomically, and retains a bounded cache. `hippo.local.json.example` documents schema
2; any machine policy copied to ignored `hippo.local.json` cannot weaken upstream floors.

Use the per-user shared root across checkouts. Override `HIPPO_ROOT` only for isolated tests or a separately
administered domain. Corrupt or unverifiable coordination state fails closed; follow upstream recovery guidance instead
of deleting state based on diagnostic PIDs.

Evidence contains capacity and process-health measurements, never file contents, command arguments, origins,
credentials, or user data. Verify bootstrap and pressure behaviour with isolated synthetic releases, state, and fake
pressure; never create real host pressure to prove shedding. Scheduled Linux/macOS smoke verifies identity, schema,
mappings, shared-root policy, and cleanup while ordinary hosted project jobs retain runner-native limits.
