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

Use environment-variable names without values, unmistakably fake placeholders, and repository-relative paths. Standard
non-identifying system paths such as `/dev/null` and an owner's already-public identity in a public repository URL are
not machine-secret disclosures. When context makes a value uncertain, leave it unstaged until the owner confirms it is
public.

Immediately before every commit, run `git diff --cached --check`, inspect `git diff --cached --no-ext-diff`, and search
the staged additions for credential-shaped values and user-specific home paths. These checks supplement judgment; a
clean pattern search never proves that arbitrary content is safe.

## Markdown Link Validation

Pre-push validates every Git-tracked repository Markdown file, not only changed documents, because a rename breaks links
in files the change never touched. The [workspace commands](workspace-commands.md#repository-checks) reference states
what the check reads and the `git add -N` caveat for a new document.

## Commit Attribution

Commits and pull requests must carry no AI attribution. Do not add a `Co-Authored-By` trailer for an assistant, a
generated-with footer, or a session link, by hand or by tool default. The author and committer remain the repository
owner. Configure each agent to suppress its own attribution; for Claude Code, `.claude/settings.json` sets
`attribution.commitTrailers` to `false`, empties the `commit` and `pr` text, and turns `sessionUrl` off.

## Pull Request Content

This repository does not use pull requests for integration. Follow the
[integration path policy](../conventions/integration-path-policy.md).

## Exception

A bypass is allowed only with explicit approval from the repository owner. Record the approval and reason in the commit
record so the skipped validation is visible and can be rerun promptly.
