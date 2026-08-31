---
tldr: "Requires each commit to carry exactly one purpose that can be reviewed and reverted alone."
when_to_use: "Use when staging and committing, especially after a session that touched several concerns."
---

# Thematic Commits Policy

## Scope

This policy governs how a working tree is divided into commits. The [commit hook policy](../development/commit-hook-policy.md) owns when committing is authorized, what the hooks enforce, and the Conventional Commits message format; this one owns where the boundary between two commits falls.

## One Purpose per Commit

Every commit represents one coherent purpose, and carries everything that purpose needs: the code, its tests, the documentation that describes it, and the configuration that enables it. A commit missing one of those is not smaller, it is broken — it names a state the repository never actually worked in.

Just as firmly, a commit excludes anything serving a different purpose, even work done in the same session and sitting in the same working tree. A theme is defined by intent, not by file type or directory: a feature and its tests and its documentation are one theme, and the unrelated typo fixed along the way is another.

## Splitting

Inspect the working tree and the staged diff before committing, and stage one theme at a time when several are present. Order dependent commits so each one stands on the one before it.

"Commit everything" is an instruction to make as many commits as the work needs, not to make one. Answering it with a single mixed commit discards exactly the information the request assumed would be kept.

## Why

The value shows up later. History is reviewed, reverted, bisected, and quoted, and every one of those operates on whole commits. A commit with two purposes cannot be reverted for one of them, and a bisect that lands on it names two suspects instead of one.

## Verification

Verified in review, by reading the commit against its own diff: does the message describe everything in it, and does the diff contain nothing the message does not mention.
