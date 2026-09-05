---
tldr: "Requires canonical skills and agents with thin, semantically equivalent harness adapters."
when_to_use: "Use when adding, changing, renaming, or removing a shared skill, agent, adapter, or harness capability."
---

# Harness Capability Parity Policy

Codex, Claude Code, and opencode receive the same repository-owned skill workflows, custom-agent intent, safety
constraints, and supported capabilities. Matching names or counts is insufficient: every harness must reach the same
canonical content, and native adapters must preserve it without adding prompt instructions or weakening restrictions.

## Canonical Sources

- `.agents/skills/<name>/SKILL.md` and every regular supporting file below its directory form one canonical skill
  bundle. Its directory and frontmatter name match.
- `.agents/agents/<name>.md` is the canonical prompt and semantic capability contract for one custom agent. Its
  frontmatter records `name`, `description`, `mode`, `requires`, `denies`, and `constraints`.
- Root `AGENTS.md` and its exact Claude import remain governed by the
  [instruction alignment policy](agent-instruction-alignment-policy.md).

Claude receives one thin skill adapter under `.claude/skills/<name>/SKILL.md`. It repeats only the exact canonical name
and description, then directs the harness to read the complete canonical bundle. Codex and opencode discover
`.agents/skills/` directly; no opencode copy is allowed.

Every canonical custom agent has exactly one native adapter under `.claude/agents/`, `.codex/agents/`, and
`.opencode/agents/`. An adapter contains only native identity, mode, tools or permission metadata, and the fixed route
to its canonical definition. It must preserve the canonical description and strongest native representation of every
required capability, denial, and constraint. When a harness cannot express one natively, the canonical prompt retains
the restriction and the supported-harness policy records the limitation; never silently omit it.

Commands remain harness-native and are outside the canonical skill and agent contract. No repository capability is
required merely because a harness supports it; add one only after a separate rule establishes recurring need.

## Deterministic Verification

`npm run check:harness-parity` validates exact instruction routing, canonical manifests, complete skill bundles, agent
adapter coverage, descriptions, routes, and the repository's known native permission mappings. It rejects missing,
extra, stale, malformed, prompt-extending, or permission-weakening adapters and competing instruction sources.

The success result reports dynamic harness, skill, and agent counts plus a SHA-256 digest over normalized canonical
instruction, skill, and agent content in ordinal path order. Findings are stable and path-specific. The validator is
read-only, network-free, process-free, and excludes links, generated trees, user-global state, and local overrides.

Follow [Harness Alignment](../workflows/harness-alignment.md) for the canonical edit, adapter reconciliation,
verification, and recovery sequence.
