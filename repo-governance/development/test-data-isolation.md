---
tldr: "Keeps every automated and manual test synthetic, isolated, marked, and unable to fall back to production data."
when_to_use: "Use when a test creates identities, files, processes, browser state, databases, or retained evidence."
---

# Test Data Isolation

## Iron Rule

No automated or manual test may read, write, authenticate against, migrate, derive fixtures from, or otherwise touch a
production user, production data, production credentials, or a production user context. Every test uses synthetic data
inside a validated per-run boundary. If isolation cannot be proved before the subject starts, fail closed. No exemption,
debug session, local environment, or convenient fixture relaxes this rule.

## Identity and Boundary

- Give each run a unique identifier and an isolated filesystem/database root, process namespace or port lease, and
  browser context when applicable. Mark owned resources with that identifier before use.
- Prefix synthetic usernames with `test-user-` when a system accepts usernames. Never mutate an existing identity to
  simulate another one.
- Reject a root or context that resolves inside production state, lacks its expected marker, or is shared unexpectedly.
- Keep parallel workers on distinct identities and user-owned paths. Never assert mutable aggregate state another worker
  can change.
- Use synthetic payloads only. Never copy production content, sessions, cookies, tokens, credentials, or identifiers
  into fixtures, snapshots, logs, or reports.

## Production Inspection

Production schemas may be inspected only through read-only structural checks that compare public schema versions, record
types, field names, and value types with synthetic records. Never authenticate as a real user, copy a production record,
or read payload values to derive a fixture. Evidence contains only value-free pass/fail results.

## Cleanup

Cleanup runs in guaranteed failure-safe teardown after browsers, servers, and child processes stop. It removes only the
exact validated resources owned by the current run. Cleanup failure fails the test or gate and reports only a safe,
synthetic location. Never delete by a broad glob, username prefix, or age alone.
