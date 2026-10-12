# Muse Policy Hooks

These scripts connect Muse's native tool events to the repository's existing policy endpoint.

Registrations live in [`hooks.json`](../hooks.json). The adapter forwards the original stdin JSON unchanged to
[`scripts/agent-policy-hook.sh`](../../scripts/agent-policy-hook.sh) with `--harness muse`.

## Files

- [`run-policy-hook.sh`](run-policy-hook.sh) — Forwards the Muse `PreToolUse` payload to the repository policy endpoint.
  Empty stdout is an allow and stays silent with exit 0; a denial prints to stderr with exit 2; an endpoint failure
  prints to stderr with exit 2, so the harness fails closed.
- [`run-policy-hook.test.sh`](run-policy-hook.test.sh) — Six cases: wrong arguments, empty stdin, garbage stdin, an
  allowed read of a public repository file, a denied `write_file` to `.env`, and a denied shell reference to `.env`.
