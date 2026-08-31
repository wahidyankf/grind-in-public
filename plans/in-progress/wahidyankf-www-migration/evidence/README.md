# Evidence

Command output and measurements this plan's `delivery.md` items produce and later items read back. The [five-document structure](../../../../repo-governance/conventions/plans-organization-policy/five-document-structure.md) names this directory for exactly that, which keeps [`learnings.md`](../learnings.md) to what [knowledge capture](../../../../repo-governance/conventions/plans-organization-policy/knowledge-capture.md) defines it as: one short paragraph per surprise, wrong assumption, or rule that failed to prevent what it targets.

The split is not bookkeeping. Phase 7 triages every `learnings.md` entry to exactly one durable home, and a pasted `rg` result has no durable home to reach. Raw output belongs here, where a later phase can read it; the lesson drawn from it, if any, belongs in `learnings.md`.

Two rules bind every file here, both from [plan document safety](../../../../repo-governance/conventions/plans-organization-policy/plan-document-safety.md), because this directory is committed and public. Name a secret's variable and location, never its value. And rewrite an absolute machine path as `$SRC` or a repository-relative path before pasting output that contains one; `delivery.md` keeps `$SRC` out of its own prose for the same reason.

Scratch output that nothing reads back belongs in the gitignored `local-tmp/` instead, which is where this plan puts its step-cardinality report and its `cv/` digest records.

## Directory Map

This index is created with the plan; each file below is added by the `delivery.md` item that writes it, and that same item converts the entry to a relative link, so an entry is unlinked only while the file it names does not yet exist. A Phase 7 item in [`delivery.md`](../delivery.md) confirms the result: eight linked entries and no ninth file. The confirmation is worth an item of its own because `npm run check:markdown-links` validates the links a document carries rather than the entries it is missing, so neither an unlinked line nor an unlisted file reaches a gate on its own.

- [phase-0-baseline.md](phase-0-baseline.md) — the six Phase 0 baseline commands, their exit statuses, and the two source-repository results.
- [phase-1-toolchain.md](phase-1-toolchain.md) — the resolved versions of TypeScript, Biome, ESLint, and `eslint-plugin-jsdoc`.
- [phase-2-background-coverage.md](phase-2-background-coverage.md) — the `Given=0` file list compared against the files carrying a `Background`.
- [unused-importers.txt](unused-importers.txt) — the importer search behind the Phase 3 dependency and file removals.
- [node-type-stripping.md](node-type-stripping.md) — the `node apps/wahidyankf-www/scripts/generate-cv-pdf.ts` attempt, its exit status, and any error text.
- [phase-3-measurements.md](phase-3-measurements.md) — the starting unit coverage percentage and the `static-routes:validation` wall-clock duration.
- [cv-references.txt](cv-references.txt) — every reference to the repository-root `cv/` directory outside the ported application, recorded before any Phase 4 edit.
- `vercel-json-digest.txt` — the SHA-256 of `vercel.json` as it arrives from `ose-public`, taken in Phase 6 before this repository's Prettier reformats it. The reformat changes the file's bytes and not its parsed configuration, so this digest is the provenance proof that a later `shasum` against the delivered file can no longer supply.
