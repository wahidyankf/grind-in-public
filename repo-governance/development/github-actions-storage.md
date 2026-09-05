---
tldr: "Keeps GitHub Actions artifacts, Packages, and caches inside their separate free storage allowances."
when_to_use: "Use when changing GitHub Actions uploads, package publishing, caches, or repository storage settings."
---

# GitHub Actions Storage

## Purpose

Keep recurring GitHub Actions storage within the owner's free allowances and prevent paid Actions overage. Artifacts and
GitHub Packages share an owner-wide allowance; Actions caches use a separate per-repository allowance. This policy does
not change the repository's integration path.

## Standards

### Workflow Artifacts

Every `actions/upload-artifact` step must declare `retention-days`. Use `1` for handoff between jobs in one run and a
value from `1` through `7` for failure triage. A value above `7` requires a comment immediately above `retention-days`
in the same `with:` mapping explaining why shorter retention cannot meet the operational need.

This check passes when every upload uses its required value or justification. A missing value, handoff value other than
`1`, triage value above `7`, or unjustified longer value violates the policy.

Keep repository artifact/log retention at or below 7 days as a backstop. A higher setting violates the policy even when
every current upload declares a shorter value; explicit per-upload retention remains the primary control.

### GitHub Packages

Before a GitHub Packages publisher lands, its delivery evidence must state its lifecycle, retained version count,
average compressed size, and owner-wide steady-state estimate for Packages plus artifacts. The total must not exceed the
owner's 500 MB GitHub Free allowance.

No package-publication obligation applies while no publisher exists. Once one exists, it passes only with every
declaration and a total at or below 500 MB. Calculate artifact bytes as compressed bytes per run multiplied by runs
retained at once, and package bytes as retained versions multiplied by average compressed version bytes. Sum every
retained item owned by the account.

### Actions Caches

Keep cache storage at or below 10 GB per repository and retention at or below 7 days. Forecast demand as:

```text
average saved entry bytes x distinct save-capable keys during the retention window
```

Count matrix variants, toolchains, operating systems, lock-file keys, and Git refs separately when they create distinct
entries. If the forecast can exceed 10 GB, pull-request and non-default refs must restore caches without saving; only
the default branch may save.

This check passes when live settings meet both limits and either the forecast is at most 10 GB or every non-default/PR
path is restore-only. A higher setting, missing forecast, or forecast above 10 GB with non-default saves violates the
policy. Cache storage is separate from the 500 MB calculation.

### Paid-Usage Stop

The personal-account owner must maintain an Actions budget of `$0`; user-level budgets always hard-stop automatically
and do not offer a stop-usage toggle. Organization owners using this rule must also enable **Stop usage when budget
limit is reached**. This check passes only with current external evidence of the applicable hard stop.

Repository content and automation cannot read or enforce a personal-account billing budget. Keep billing data outside
the repository. This control is intentionally **unenforced by repository**; owner verification remains mandatory.

## Examples

```yaml
- uses: actions/upload-artifact@v4
  with:
    name: generated-report
    path: generated-report.json
    retention-days: 1
```

A 4 MB triage artifact produced twice daily for 7 days uses `4 MB x 2 x 7 = 56 MB`. Add it to all owner artifacts and
Packages before comparing with 500 MB.

## Validation

Search workflow and reusable-action YAML for uploads, package publication, and cache saves. Inspect each step. Record
calculation inputs, result, date, and source inventory in delivery evidence. Record billing verification separately.

Confirm live repository settings:

```bash
rtk gh api repos/{owner}/{repo}/actions/permissions/artifact-and-log-retention
rtk gh api repos/{owner}/{repo}/actions/cache/storage-limit
rtk gh api repos/{owner}/{repo}/actions/cache/retention-limit
```

As of 2026-09-05, this public repository has no active artifacts or detected package publisher; package inventory was
not verified because the active token lacks visibility. It uses about 856 MiB of main-only cache, has live cache limits
of 10 GB and 7 days, and retains artifacts/logs for 7 days. The scheduled workflow has no pull-request trigger. This
passes repository-local checks, not owner-wide inventory or `$0` budget verification.

Do not add a validator while no artifact or package publisher exists and cache use remains far below its cap. Mandatory
review and live evidence are proportionate. Reassess automation when a publisher appears, save-capable refs expand, or
repeated violations prove a validator's recurring value.
