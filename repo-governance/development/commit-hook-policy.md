---
tldr: "Defines required Git hooks and the narrow exception for bypassing them."
when_to_use: "Use before committing, pushing, changing hooks, or considering a hook bypass."
---

# Commit Hook Policy

This policy applies to every contributor and AI agent working in this repository.

## Authorization

Committing and pushing are separate permissions, and neither is implied by the work that produced the changes. A request
to edit, fix, finish, test, or stage authorizes none of it.

- Commit only when the owner asks for a commit, or when an owner-approved plan's checklist reaches a commit step.
- Push only when the owner asks for a push, or an approved plan's checklist reaches one. "Commit these changes" does not
  authorize a push.
- Authorization covers the current task and the scope stated with it: the files, the theme, the branch, and the remote.
  Ask before widening it, and never carry it into later work.
- Without authorization, leave the work uncommitted, say so plainly, and name what is waiting.

A plan authorizes only the steps it states, and a plan the owner has not approved authorizes nothing.

## Required Hooks

Do not bypass Git hooks with `--no-verify`, including for `git commit` and `git push`. Hooks enforce formatting,
Conventional Commit messages, and the repository checks that [workspace commands](workspace-commands.md#hooks) lists for
each hook. A failure must be investigated and resolved before continuing.

## Public Repository Safety

Treat every committed byte as public. Before committing, inspect the complete staged diff and remove credentials, secret
values, authentication material, private machine or network identifiers, and user-specific absolute paths. This includes
passwords, tokens, private keys, cookies, authorization headers, connection strings, home-directory user names,
hostnames, and private addresses. Never place a real sensitive value in a fixture, example, comment, document, generated
artifact, or commit message.

Write only values the [leak classes](../workflows/quality/push-leak-review/001-leak-classes.md#what-is-not-a-leak)
clear, and leave an uncertain value unstaged until the owner confirms it is public.

Immediately before every commit, run `git diff --cached --check`, inspect `git diff --cached --no-ext-diff`, and search
the staged additions for credential-shaped values and user-specific home paths. These checks supplement judgment; a
clean pattern search never proves that arbitrary content is safe.

Immediately before every push, run the [push leak review](../workflows/quality/push-leak-review.md) over each outgoing
commit; a finding blocks the push. The staged-diff review above does not replace it.

A leak that reaches history follows
[push review remediation](../workflows/quality/push-leak-review/002-push-review.md#remediation).

## Markdown Link Validation

Pre-push validates every tracked Markdown file; [workspace commands](workspace-commands.md#repository-checks) records
the check.

## Commit Attribution

Commits and pull requests carry no AI attribution: no assistant `Co-Authored-By` trailer, no generated-with footer, no
session link. The author and committer remain the repository owner. Each agent suppresses its own attribution; the
per-harness settings live in [agent harness support](../conventions/agent-harness-support.md). This repository does not
use pull requests for integration; follow the [integration path policy](../conventions/integration-path-policy.md).

## Exception

A bypass is allowed only with explicit approval from the repository owner. Record the approval and reason in the commit
record so the skipped validation is visible and can be rerun promptly.

## Agent Secret Access

Agent `.env*` access lives in the [agent harness support policy](../conventions/agent-harness-support.md).

1. Agents must not open or read private-key contents anywhere, including outside every repository.
2. Each secret step is exactly one copy-ready command carrying its working directory, placed by the agent on the
   operator clipboard; the operator runs it in their own terminal, entering each secret at a hidden prompt only, and
   replies with the success word the step names; the agent continues only on it.
3. Agents must never request, handle, or record secret values, full outputs, or screenshots, and must avoid any script
   or interaction that prints, echoes, logs, or otherwise exposes a secret to the harness session; diagnosis uses exit
   and error codes alone, and only local-development secrets (loopback services, dummy/example values such as
   .env.example) may appear in session records.
4. The secret session opens only inside the phase that needs it and locks before any review, pause, or PR handoff.
5. Agent-assisted delivery applies only with a proper pipeline and secrets management (tracked wrapper, closed one-word
   outputs, leak screens, green gates); elsewhere the agent stops at the secret boundary.
6. An agent may operate the leak review after reading every commit at the pinned head against the three leak classes; a
   record without that reading is fabrication.

Reason: regtech [HUMAN] removes value exposure from the session, and proper pipelines earn faster delivery. Disposition:
unenforced by decision; review checks the observations.
