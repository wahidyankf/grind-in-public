# Learnings

The Markdown-link checker enumerates Git-tracked Markdown paths, so a plan lifecycle move must be staged before link verification. Running the check first makes it inspect the deleted source path even when the destination exists; stage the move, then validate links before committing.

Phase 0 baseline on 2026-08-23: `npm install` exited 0 with 201 audited packages, no vulnerabilities, and no `package-lock.json` change; `npm run test:quick` exited 0 with existing aggregate statement coverage at 95.6%; `npm run format:check` exited 0; and `npm run check:markdown-links` exited 0.

Phase 1's reviewed checklist originally created the repository boundary-policy tests before migrating the known real-filesystem and Git cases out of the unit target. That ordering could not meet its own green acceptance criterion, so the policy-test items moved after the boundary migrations; future plans should place enforcement after the state it enforces is attainable, while keeping fixture-level policy tests earlier only when their scope says so explicitly.

The first Phase 1 push was blocked because the legacy aggregate coverage target measured new BDD and CLI packages before Phase 3's reviewed coverage slices existed, while migrated external integration tests did not instrument their production packages. A phase that must push needs its current pre-push denominator to stay attainable without hiding runtime code or pulling environment-dependent integration into a cacheable hook, so the reviewed all-runtime 99% unit slice moved into Phase 1; Phase 3 still adds the separate integration denominator and final aggregate composition.

Phase 1 gate on 2026-08-23: `npm exec nx -- run badakmini-cli:test:quick` exited 0 after typecheck, strict lint, unit tests, and 99.2% statement coverage across all `internal/...` runtime code; `npm run format:check`, `npm run check:markdown-links`, and `go -C apps/badakmini-cli test ./tests/integration` also exited 0. The exact production-runtime binding test and all five named integration-owner proofs passed, and the structural assertion found no `os` or `os/exec` import below `internal/`.
