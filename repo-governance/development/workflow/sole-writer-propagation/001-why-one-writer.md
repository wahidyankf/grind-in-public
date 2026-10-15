---
tldr:
  "Gives the reasons one writer per family suffices: a fixed input, idempotent repairs, and judging kept apart from
  writing."
when_to_use:
  "Use when justifying the sole-writer design to an adopter, or when a proposed change would let an auditor also write."
---

# Why One Writer

A single writer bound to a frozen ledger is what lets three cycles suffice:

- **The input is fixed.** The writer receives a frozen ledger and may touch only what its rows require. A repair cannot
  grow into an edit made "while I'm here" that the next audit reports as a new finding, so nitpicks do not cascade.
- **It is idempotent.** The same ledger against the same state produces the same change. Running it again on an
  already-repaired state changes nothing, so the open count can only fall.
- **Judging and writing stay separate.** The checker never writes, so each audit after a repair is independent of it.
