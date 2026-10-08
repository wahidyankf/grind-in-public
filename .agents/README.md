# Shared Agent Directory

This directory holds canonical capabilities shared by every supported harness. Codex and opencode read skills here
directly; custom-agent adapters route to the prompts here. See the
[agent harness support policy](../repo-governance/conventions/agent-harness-support.md).

## Contents

- [`agents/`](agents/README.md) — canonical custom-agent prompts and semantic capability contracts.
- [`skills/`](skills/README.md) — the shared skills available in this repository.

Nothing here is harness-specific by nature. Native adapters contain only the metadata and route their harness needs.

## Repository Tool Policy

[`agent-policy.json`](agent-policy.json) declares protected and permitted path patterns for the repository-owned
[`agent-policy-hook.sh`](../scripts/agent-policy-hook.sh) endpoint. Governing obligations remain in repository security
and resource conventions. Claude Code, OpenCode, and Command Code native bindings and cross-repository routing invoke
this endpoint. Codex retains its existing guards; its new routing and Serena registration are deferred.
