# Command Code Policy Hooks

These scripts connect Command Code's native tool events to the repository's existing policy checks.

Registrations live in [`settings.json`](../settings.json). The adapter invokes each delegate by its policy name.

## Files

- [`run-policy-hook.sh`](run-policy-hook.sh) — Normalizes native payload fields and invokes the existing policy.
- [`policy-hooks.test.sh`](policy-hooks.test.sh) — Exercises mappings and policy delegation in synthetic repositories.

## Native payload mapping

Settings match native display IDs such as `SHELL`, `READ`, `WRITE`, and `EDIT`.

Payload identifiers include `shell_command`, `read_file`, `write_file`, and `edit_file`.

These `tool_name` values map to the names the existing policies expect. Original native fields remain available.

Path aliases supply `file_path`. Multi-file reads check every path before returning permission.

The first delegate response is returned unchanged, so a later allowed path cannot hide an earlier denial.

Shell commands combine `command` and optional `args`, preserving argument quoting.

The tool directory comes from native `cwd`, `directory`, `workdir`, or the payload's session `cwd`.

The policy runs there. The owning checkout supplies its delegates and `CLAUDE_PROJECT_DIR`.

The adapter also recognizes `format-lint-markdown`. Markdown formatting receives transactional HIPPO admission.

The current settings determine which policy routes run for each event.

## Registered policies

The current settings register these existing policies:

- [HIPPO boundary guard](../../.claude/hooks/require-hippo-boundary.sh) for native shell tools with `failClosed`.
- [Rule-change check](../../scripts/check-rule-change.mjs) for native write and edit tools.

The adapter and settings contain no model or effort selection. Policy hooks follow the active session's tool calls.

These files do not establish native agent discovery or a live session probe.

Authenticated Command Code execution was not exercised in this documentation pass.

The rollout executor records synthetic test results and live evidence.
