---
tldr: "Sets main as the delivery target and defines the prod- promotion branch each deployed project gates on."
when_to_use: "Use when adding or changing deploy configuration, promoting a build, or cutting a domain over."
---

# Deployment Policy

## Delivery Target

`main` is the delivery target. Work lands there directly, as the
[plans organization policy](../conventions/plans-organization-policy.md) states of every plan, and no branch stands
between a merge and `main`.

Landing on `main` is not deploying. A commit on `main` is delivered; it is promoted only by the step below.

## Promotion Branches

A deployed project is promoted by a branch named `prod-<project>`, where `<project>` is the Nx project name. For the
personal site that branch is `prod-wahidyankf-www`.

The branch is a pointer to the commit currently promoted, and nothing advances it automatically. No plan, workflow,
hook, or CI job may move it: a promotion is the owner's decision about a specific commit, made at a moment they choose,
and an automatic advance would make every merge to `main` a deploy.

Because the branch is a pointer, it may lag `main` by any distance without that being a defect. A `prod-` branch behind
`main` means the newer commits are delivered and not yet promoted, which is the normal state.

## Build Gating

Each deployed project gates its own build on its own `prod-` branch. The project's `vercel.json` carries an
`ignoreCommand` that compares the building ref against that branch name and cancels the build when they differ, so a
push to `main` or to any other branch does not spend a deploy.

One project's promotion therefore never triggers another's. That is the reason the branch name carries the project name
rather than being a shared `prod`.

The gate is only half the wiring. The hosting project's own production branch setting must name the same
`prod-<project>` branch, and that setting lives in the provider's dashboard rather than in this repository. Leaving it
at `main` deadlocks the project in a way that looks like nothing happening: pushes to `main` become production builds
that the `ignoreCommand` immediately cancels, and pushes to the `prod-` branch become previews that never reach the
domain. Every push is then accounted for and none of them ships.

Check that setting when a project is first connected, and again after any change to which repository it deploys from,
because a reconnect can reset it to the repository's default branch. It cannot be read or written from a tracked file,
so a green repository is not evidence that it is right.

## Domain Cutover

Pointing a live domain at a project in this repository is a separate, explicitly authorized act. Porting or landing a
project's deploy configuration does not authorize it, and neither does a green gate: configuration that is present and
correct is still inert until the owner cuts the domain over.

Deploy configuration may therefore be committed for a project whose promotion branch does not yet exist here. That is
the expected state for a project mid-migration, and the absence of the branch is what keeps the configuration inert.

## Secrets

Deploy configuration names a secret's variable and where it is set. It never carries the value, and neither does any
tracked file; the [plan document safety](../conventions/plans-organization-policy/plan-document-safety.md) rules apply
to every path in this repository, not only to plan documents.
