# Badak Mini Layered BDD

Adopt BeaverNest's executable-specification model for `apps/badakmini-cli`, add the dedicated `apps/badakmini-cli-e2e` harness, and make the same role-based model compulsory for future projects below `apps/` and `libs/`.

The canonical corpus will live below `specs/apps/badakmini-cli/behavior/`. Unit, local-only integration, and process E2E adapters will consume that same recursive corpus. A fast static behavior gate will reject malformed features, corpus drift, undefined or ambiguous steps, unused bindings, incomplete drivers, and layer filtering.

## Directory Map

- [Business Requirements](brd.md)
- [Product Requirements](prd.md)
- [Technical Design](tech-docs/README.md)
- [Delivery Checklist](delivery.md)
- [Learnings](learnings.md)

## Quality Gate

- 2026-08-23 — strict — 7 cycles — partial (2 HIGH open: dependency ordering and Phase 3 Pause Safety)
- 2026-08-23 — strict — 2 cycles — pass (0 findings on two consecutive runs after rewriting the affected checklist section)
- 2026-08-23 — strict — 2 cycles — pass (0 findings on two consecutive runs after correcting Phase 1 boundary-enforcement ordering)
- 2026-08-23 — strict — 6 cycles — pass (0 findings on two consecutive runs after moving the all-runtime unit coverage slice and host-adapter boundary into Phase 1)
- 2026-08-23 — strict — 3 cycles — pass (0 findings on two consecutive runs after making external-spec cache invalidation explicit)
- 2026-08-23 — strict — 4 cycles — pass (0 findings on two consecutive runs after resolving the application integration rule, restoring word-limit headroom, and indexing relocated policies)
