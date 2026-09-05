---
tldr: "Keeps every non-drill application's C4 model current with its implemented boundaries."
when_to_use:
  "Use before changing a non-drill application's production code, configuration, tests, or specifications that can
  affect architecture."
---

# Architecture Specifications

## Scope

Every non-drill application corpus under `specs/apps/<name>/` contains a canonical root `architecture.md` C4 model. The
model describes only the current, as-built system; proposed or unimplemented designs belong in a plan or other
explicitly prospective document. Implementation and dedicated E2E projects that share one corpus share its architecture
model.

## Required Coverage

Each model identifies its scope, people, software systems, runtime or deployment containers, external interfaces, and
material relationships. It shows persistent or temporary data stores plus material process, network, and trust
boundaries.

Include a system-context view and every useful container view. Include a component view only when it materially
clarifies internal responsibilities. Add searchable prose for constraints that a diagram cannot communicate safely, and
link to the corpus's `behaviours/` directory. Each implementing project's README links back to the model.

Use English and follow the [Markdown style policy](../conventions/markdown-style-policy.md): C4 diagrams are
terminal-first ASCII in fenced `text` blocks.

## Scaling

Keep one `architecture.md` while a reader can move from context to the relevant boundary without scanning unrelated
detail and each diagram remains legible at normal Markdown width. Split only when independent views, constraints, or
behaviour traceability make the entry point difficult to use, a diagram cannot stay legible when simplified, or
independently evolving areas create recurring review noise.

When splitting, retain `architecture.md` as the entry point for scope, system context, shared constraints, and an index.
Put detail below `architecture/`, give every statement and diagram one canonical home, and link each detail document
back to the entry point, relevant behaviour, and implementing project.

## Change Discipline

Before changing non-drill application production code, configuration, tests, or specifications, read the relevant
architecture model, Gherkin, and tests, then assess architectural impact. Update the model in the same change when the
implementation changes a documented actor, system, container, component responsibility, relationship, interface, runtime
or deployment boundary, data store, data flow, or security boundary.

Behaviour-only changes and implementation detail below the documented component boundary do not require diagram churn
when every architectural statement remains accurate. Architecture models complement executable Gherkin and TDD; they do
not replace either.

## Verification

Confirm the model describes the final implemented state, every required link resolves, and affected project tests pass.
Run `rtk ./hippo run --class ephemeral --disk-path . -- npm run format:check` and
`rtk ./hippo run --class ephemeral --disk-path . -- npm run check:markdown-links` after changing the model or this
policy.
