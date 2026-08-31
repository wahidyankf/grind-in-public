# Node Type Stripping

Attempted 2026-09-01, Phase 3, before relying on the `tsx` pin. Node 24.16.0.

## Command and Result

```sh
cd apps/wahidyankf-www && node scripts/generate-cv-pdf.ts
```

Exit status **1**. Nothing was written to stdout.

## Error

```text
(node:...) [MODULE_TYPELESS_PACKAGE_JSON] Warning: Module type of
  .../apps/wahidyankf-www/scripts/generate-cv-pdf.ts is not specified and it doesn't
  parse as CommonJS. Reparsing as ES module because module syntax was detected.

Error [ERR_MODULE_NOT_FOUND]: Cannot find module
  .../apps/wahidyankf-www/src/features/cv/core/data
  imported from .../apps/wahidyankf-www/scripts/generate-cv-pdf.ts
    at finalizeResolution (node:internal/modules/esm/resolve:271:11)
```

## What This Proves

Node did strip the types. It failed one step later, at module resolution: the script imports its CV record as `.../core/data`, with no file extension, which is ordinary TypeScript and which Node's ESM resolver refuses — it resolves a specifier to a literal path and will not try `.ts` on its own. `tsx` resolves extensionless TypeScript specifiers, which is the specific capability being relied on here.

So the dependency is not kept out of habit or because nobody checked. It is kept because the alternative fails on a real, reproduced error, and the two ways around it are worse: rewriting the import to `.../core/data.ts` puts an emitted-file extension in a source module for the benefit of one script, and `--experimental-specifier-resolution` was removed from Node.

The `MODULE_TYPELESS_PACKAGE_JSON` warning above is a separate, non-fatal matter — the app manifest declares no `"type"`. It is noted only so a future reader does not mistake it for the cause.

## Consequence

`tsx` stays pinned at `4.23.13` in the root manifest, and `generate:cv-pdf` keeps its `npx tsx scripts/generate-cv-pdf.ts` command. The `delivery.md` item that would have removed the pin is **not triggered** and takes a dated `Not triggered` disposition in Phase 7 rather than a tick.

Re-test this when Node changes its resolver or when the script stops importing an extensionless specifier; the pin exists for exactly one reason and should not outlive it.
