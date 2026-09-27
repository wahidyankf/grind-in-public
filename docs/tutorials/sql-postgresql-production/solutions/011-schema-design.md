---
tldr: "Separates versioned configuration, detections, workflow, relations, and append-only evidence."
when_to_use: "Use after Drill 011."
---

# Solution 011: Schema Design

Use tenant-prefixed keys for `policy_versions`, `alerts`, `cases`, and junction `case_alerts`. Alerts reference the
exact policy version used, not merely a mutable policy identity. Append evidence with event identity, actor, reason,
source hash, and timestamp; deny application updates/deletes and archive by retention policy. Typed columns own keys,
state, money, and time; `jsonb` holds versioned rule detail/evidence payload. One document cannot enforce all
cross-document references, queue indexes, concurrent workflow transitions, and relational reports cleanly.
