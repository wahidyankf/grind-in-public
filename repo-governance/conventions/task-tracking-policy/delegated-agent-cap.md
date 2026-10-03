---
tldr: "Caps every session at three live delegated agents besides its main thread, at any depth and in every harness."
when_to_use: "Use before spawning a subagent or background agent, and when running a workflow step in parallel."
---

# Delegated Agent Cap

[Task Tracking Policy](../task-tracking-policy.md)

A delegated agent is any agent a session spawns to carry part of its work: a [subagent](../agent-vocabulary.md), a
background agent, or an unnamed general-purpose one.

## The Cap

At most N delegated agents run at once in a session, and N is 3. The main thread that delegates is the `+1` and takes no
slot, so at most four agents work at once.

The count covers every delegated agent alive in the session, foreground or background, at any depth: an agent that a
delegated agent spawns takes a slot of its own. It binds in every harness whose session can spawn delegated agents,
whatever that harness calls them — among the [supported harnesses](../agent-harness-support.md), the Claude Code `Agent`
tool, Codex spawned agents, and the opencode `task` tool. Work beyond N waits until a running unit returns; it is never
launched over the cap.

A workflow step that runs in parallel, such as a review's specialists or a gate's per-harness research, therefore runs
at most N units at once and queues the rest.

Only the owner, or the plan being executed, declares another N, and only for that task. An agent never changes N on its
own judgement.

Reason: every live agent spends against the same model quota and the same machine at once, and nested fan-out multiplies
that cost beyond what one session can follow or recover from.

## Verification

Followed: at no point does a session have more than N delegated agents alive. Violated: a session launches a delegated
agent while N are already running, counting the ones its delegated agents spawned.

Unenforced by decision: the owner declined mechanical enforcement, so no hook or harness setting carries the cap and
none is added for it. Review verifies it against the agents a session's transcript or task list shows it launched.
