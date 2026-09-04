# Phase 2 Background Coverage

Checked 2026-09-01 over the eleven feature files in `specs/apps/wahidyankf-www/behaviours/`, against the
step-cardinality report in the gitignored `local-tmp/step-cardinality.txt`.

## The Report

53 lines, one per scenario, matching the 53 scenarios the corpus carries. `grep -cvE 'Given=[01] When=1 Then=1'` printed
`0`: every scenario has exactly one primary `When` and exactly one primary `Then`, and none has more than one primary
`Given`.

## Given=0 Against Background

The [Gherkin cardinality rule](../../../../repo-governance/development/behaviour-driven-development-policy.md) exempts a
scenario whose `Given` comes from a `Background`. The two file lists are not merely in a subset relation, they are
equal:

| Carries a `Given=0` scenario       | Carries a `Background`             |
| ---------------------------------- | ---------------------------------- |
| `accessibility.feature`            | `accessibility.feature`            |
| `cv.feature`                       | `cv.feature`                       |
| `home.feature`                     | `home.feature`                     |
| `personal-projects.feature`        | `personal-projects.feature`        |
| `responsive.feature`               | `responsive.feature`               |
| `search.feature`                   | `search.feature`                   |
| `static-filterable-routes.feature` | `static-filterable-routes.feature` |
| `theme.feature`                    | `theme.feature`                    |

`comm -23` over the two sorted lists printed nothing, so no `Given=0` scenario sits in a file without a `Background`.

The three files absent from both columns are the loader corpora and the app's own env-loader feature —
`env-loader.feature`, `tier-env-loading.feature`, and `port-resolver.feature`. Each of their scenarios states its own
`Given`, which is why they need no `Background` and appear in neither list.

Nothing was flagged, so the splitting item that follows this check in `delivery.md` is a no-op.
