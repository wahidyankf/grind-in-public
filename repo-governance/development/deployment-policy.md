---
tldr: "Sets main as the delivery target and defines the prod- promotion branch each deployed project gates on."
when_to_use: "Use when adding or changing deploy configuration, promoting a build, or cutting a domain over."
---

# Deployment Policy

## Delivery Target

`main` is the delivery target. Work lands there directly, as the [plans organization policy](../conventions/plans-organization-policy.md) states of every plan, and no branch stands between a merge and `main`.

Landing on `main` is not deploying. A commit on `main` is delivered; it is promoted only by the step below.

## Promotion Branches

A deployed project is promoted by a branch named `prod-<project>`, where `<project>` is the Nx project name. For the personal site that branch is `prod-wahidyankf-www`.

The branch is a pointer to the commit currently promoted, and nothing advances it automatically. No plan, workflow, hook, or CI job may move it: a promotion is the owner's decision about a specific commit, made at a moment they choose, and an automatic advance would make every merge to `main` a deploy.

Because the branch is a pointer, it may lag `main` by any distance without that being a defect. A `prod-` branch behind `main` means the newer commits are delivered and not yet promoted, which is the normal state.

## Build Gating

Each deployed project gates its own build on its own `prod-` branch. The project's `vercel.json` carries an `ignoreCommand` that compares the building ref against that branch name and cancels the build when they differ, so a push to `main` or to any other branch does not spend a deploy.

One project's promotion therefore never triggers another's. That is the reason the branch name carries the project name rather than being a shared `prod`.

## Domain Cutover

Pointing a live domain at a project in this repository is a separate, explicitly authorized act. Porting or landing a project's deploy configuration does not authorize it, and neither does a green gate: configuration that is present and correct is still inert until the owner cuts the domain over.

Deploy configuration may therefore be committed for a project whose promotion branch does not yet exist here. That is the expected state for a project mid-migration, and the absence of the branch is what keeps the configuration inert.

## Secrets

Deploy configuration names a secret's variable and where it is set. It never carries the value, and neither does any tracked file; the [plan document safety](../conventions/plans-organization-policy/plan-document-safety.md) rules apply to every path in this repository, not only to plan documents.
