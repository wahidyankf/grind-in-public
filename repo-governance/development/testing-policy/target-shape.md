---
tldr: "States how a project.json target declares its cache, outputs, working directory, and shared inputs."
when_to_use: "Use when adding or changing any Nx target, not only a testing one."
---

# Target Shape

The [testing policy](../testing-policy.md) owns which targets a project exposes. This document owns what each one
declares. It binds **every** target a project declares, not only the ten in the contract: a rule that reached the
contract targets alone would leave the others to be inferred by comparing files, which is the drift it exists to end.

## Cache

Every target declares `cache` explicitly wherever the root `nx.json` `targetDefaults` does not already reach it. An
omitted `cache` is not a neutral default — Nx reads it as uncached — so an undeclared target and a deliberately uncached
one look identical in the file and behave identically at runtime, and nothing distinguishes the decision from the
oversight. Do not restate `cache` on a target `targetDefaults` already covers; the duplicate is what later disagrees
with it.

## Outputs

A cached target that writes an artifact declares `outputs` naming every path it writes. Nx replays a cache hit without
running the command, so a cached target with no declared output restores nothing: it reports success and produces the
file it claimed to produce only on the runs that missed. An uncached target declares no `outputs` at all, because the
declaration does nothing there and an inert declaration reads as a decision.

An artifact is something a later target or a person consumes. A compiler's own incremental state is not one —
`tsconfig.tsbuildinfo` is read by nothing but the compiler that wrote it and is regenerated on demand — so an
incremental `typecheck` needs no `outputs`.

## Working Directory

A single-command target declares `options.cwd`, normally `{projectRoot}`, rather than encoding its own project path in
the command string. A path written into the command is duplicated knowledge that survives a directory rename silently.

Converting a target to `cwd` is not always a matter of deleting the path. Anything the command resolves relative to the
working directory moves with it: a `mkdir` prefix, a `$PWD` expansion, a search path given as `.`. Each of those fails
quietly rather than loudly — writing to a directory nobody reads, or scanning a tree far wider than intended — so read
the whole command before trusting the strip.

## Shared Inputs

Follow the Beaver Nest descriptor shape: each affected target names its canonical Gherkin corpus directly in `inputs`.
An E2E target also includes `^default`; an owner behaviour aggregate directly includes its E2E adapter. This repetition
is intentional because each target remains auditable without resolving a named indirection.

The deterministic project-contract validator checks ownership, target presence, one-way E2E dependencies, cache
decisions, and aggregate delegation. Before deleting an input because something broader already covers it, search for a
test that names it. "Covered by a wider rule" and "safe to remove" are different claims.

Verify input reach behaviourally as well as structurally. Change the content of a file the input should cover and
confirm a cached target misses; Nx hashes content rather than timestamps, so a miss is only possible if the input truly
reaches that file.
