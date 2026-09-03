# Phase 0 Baseline

Run 2026-09-01, before anything was copied. `$SRC` is the executor's `ose-public` working copy; its absolute path is
deliberately not written here, per
[plan document safety](../../../../repo-governance/conventions/plans-organization-policy/plan-document-safety.md).

## Repository Gates

| Command                        | Exit | Observed                                                                                                                                                                                                                                                             |
| ------------------------------ | ---- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `npm install`                  | 0    | Locked dependencies installed; `git diff --stat package-lock.json` reported no change. npm warned that `nx@23.1.1` has install scripts not yet covered by `allowScripts` — a pre-existing condition of this workspace, unchanged by this plan and not acted on here. |
| `npm run test:quick`           | 0    | 1 of 1 task, served from the Nx cache.                                                                                                                                                                                                                               |
| `npm run format:check`         | 0    | All matched files use Prettier code style.                                                                                                                                                                                                                           |
| `npm run check:markdown-links` | 0    | Repository-local Markdown links are valid.                                                                                                                                                                                                                           |

## Source Repository

| Command                                                 | Exit | Observed                                                                                                                                               |
| ------------------------------------------------------- | ---- | ------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `git -C "$SRC" rev-parse HEAD`                          | 0    | `e74818fc06c4c104725383384d2aa38305a503ef`, matching the provenance in the plan [`README.md`](../README.md).                                           |
| `git -C "$SRC" status --porcelain <eight source paths>` | 0    | Empty. Paired with `git ls-files --error-unmatch` over the same eight paths, which printed no `MISSING` line, so every path is both clean and tracked. |

The first attempt at the cleanliness check named `apps/wahidyankf-www-e2e` — this repository's name for the E2E project,
which does not exist in `ose-public`. It passed on empty output while matching nothing. The existence half was added to
the checklist item as a result, and the re-run above is what this table records. [`learnings.md`](../learnings.md)
carries the lesson.
