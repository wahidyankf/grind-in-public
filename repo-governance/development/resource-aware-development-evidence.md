---
tldr: "Keeps the HIPPO consumer pin, shared-root policy, and retained evidence contract together."
when_to_use:
  "Use when changing HIPPO consumer identity, reconciling pinned capabilities, or inspecting retained evidence."
---

# Consumer Integrity and Evidence

Admission, supervision, recovery, and harness enforcement remain in
[Resource-Aware Development](resource-aware-development.md).

`hippo.lock` pins release, commit, and each platform asset's SHA-256. The wrapper revalidates bytes, publishes
atomically, and bounds its cache. `hippo.local.json.example` documents schema 3; ignored machine policy cannot weaken
upstream floors. A worktree without it inherits the primary checkout's policy.

The lock pins executable identity, not governance semantics. Before changing consumer behaviour, read the Hippo
repository at the commit in `hippo.lock`, especially its exit-code and recovery references, then reconcile this rule,
Gherkin, and harness checks with the capabilities that commit actually provides. Never infer capability from SemVer
ordering or copy a release number into the rule.

Use the shared per-user root. Override `HIPPO_ROOT` only for isolated tests. Unverifiable coordination state fails
closed; follow upstream recovery guidance rather than deleting state from diagnostic PIDs.

Evidence contains capacity and process health, never contents, arguments, origins, credentials, or user data. Test
pressure only with isolated synthetic state. Scheduled Linux/macOS smoke verifies identity, schema, mappings, root, and
cleanup; ordinary hosted jobs retain runner-native limits.

`hippo.identity.json` labels this repository in live status and the bounded 30-day history. Identity discovery works
from nested paths and an explicitly authorized contained `worktrees/<task>` checkout; add privacy-safe
`--tag checkout=worktree --tag plan=<slug>` values per run. Use `./hippo status`,
`./hippo watch --source grind-in-public`, and `./hippo history --since 30d --source grind-in-public` directly. Every
checkout uses the shared default root; set `HIPPO_ROOT` only for isolated tests. Raw evidence rolls for seven days;
compacted daily summaries roll for 30 days under byte caps.
