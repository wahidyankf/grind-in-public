---
tldr: "Uses nullable expansion, resumable key-range backfill, validation, and delayed contraction."
when_to_use: "Use after attempting Case 007."
---

# Debrief 007: Billion-Row Migration

Add nullable `region_code`, deploy code that writes it while old readers ignore it, then backfill stable primary-key
ranges with checkpoints, rate limits, short transactions, and retryable updates. Verify null count, per-tenant mapping,
sample hashes, WAL/lag, and application parity. Add `CHECK (region_code IS NOT NULL) NOT VALID`, validate later, and
build needed indexes concurrently. Switch reads behind a reversible flag; make the column required only when every
writer is new. Remove fallback after two release/rollback windows. Pause automatically on I/O, lag, or p99 thresholds.
