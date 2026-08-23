# Learnings

The Markdown-link checker enumerates Git-tracked Markdown paths, so a plan lifecycle move must be staged before link verification. Running the check first makes it inspect the deleted source path even when the destination exists; stage the move, then validate links before committing.

Phase 0 baseline on 2026-08-23: `npm install` exited 0 with 201 audited packages, no vulnerabilities, and no `package-lock.json` change; `npm run test:quick` exited 0 with existing aggregate statement coverage at 95.6%; `npm run format:check` exited 0; and `npm run check:markdown-links` exited 0.
