---
tldr: "Keeps plan documents secret-free and their diagrams terminal-readable."
when_to_use: "Use when authoring, reviewing, or repairing any document or asset under plans/."
---

# Plan Document Safety

Plans are committed records. Name a secret's variable and location, never its value; do not include credentials,
cookies, private identifiers, real accounts, user data, or runtime payloads in prose, evidence, diagrams, or assets.

All plan diagrams are terminal-first ASCII in fenced `text` blocks. Mermaid is not used in a plan, including a plan that
studies a system which uses it elsewhere. UI mockup SVG is an accessibility-reviewed design asset, not an architecture
diagram, and remains subject to the UI-design rule.

The [plan quality gate](../../workflows/plan-quality-gate.md) verifies the ASCII and secrets rules directly.
