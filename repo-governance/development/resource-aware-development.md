---
tldr: "Guards compute-bearing Nx work with a checksum-pinned, cross-platform resource admission boundary."
when_to_use: "Use before running or wiring builds, tests, repository checks, services, or other compute-heavy work."
---

# Resource-Aware Development

Apply this standard to compute-bearing Nx work under `apps/`, `libs/`, and repository tools. It reduces host pressure
without guaranteeing process survival.

## Required Behaviour

- Run compute-bearing work through the guard. Exit `75` is transient capacity or a held lease, not a test failure:
  confirm the earlier run exited, then retry the same command serially.
- Never bypass the guard, retry concurrently, background a retry loop, weaken a gate, or change admission class merely
  to obtain admission.
- Exit `73` means storage is blocked; safely free space before retrying. Exit `78` means invalid configuration or an
  incompatible strict profile; replan instead of looping.
- Use `transactional` only for an authorized mutation that must not be interrupted after starting, never for ordinary
  build or test work.
- The guard signals only its child process group and never a production service or unverified PID.

Canonical execution is:

```sh
rtk ./resource-guard run --class ephemeral --disk-path . -- npm exec -- nx run -p <project> -t <target>
```

The wrapper maps upstream concurrency to `NX_PARALLEL` and `GOMAXPROCS`. Ordinary work can fall from `balanced` through
`constrained` to `minimal`; strict transactional work cannot. Resource recovery is infrastructure handling and never a
semantic quality-gate retry.

## Consumer Integrity

`resource-guard.lock` pins the public release version, source commit, and SHA-256 for each supported macOS/Linux and
amd64/arm64 asset. The wrapper downloads outside the repository, verifies the archive and embedded identity, publishes
the binary atomically, and retains a bounded cache. Copy `resource-guard.local.json.example` to the ignored
`resource-guard.local.json` for machine policy; local configuration cannot weaken upstream safety floors.

Evidence contains capacity and process-health measurements, never file contents, command arguments, origins,
credentials, or user data. Verify bootstrap behaviour with deterministic fake releases and fake pressure; never create
real host pressure to prove shedding.
