---
tldr: "Requires every Nx project to maintain a useful project-level README."
when_to_use: "Use when creating or changing an application or library project under apps/ or libs/."
---

# Project README Policy

Every application or library project rooted under `apps/` or `libs/` has a `README.md` at its project root. Add it in
the same change that creates the project.

The README is that project's human entry point. It states the project's purpose and ownership boundary, relevant setup
and configuration, resolved Nx commands run from the repository root, important source and test paths, interfaces and
consumers when applicable, testing guidance, and links to its canonical specifications and repository documentation.
Omit inapplicable sections instead of adding empty boilerplate, and link to shared guidance instead of copying it.

Every project change includes a README impact check. Update the README in the same change when purpose, ownership,
public interfaces, prerequisites, configuration, targets, structure, or operating procedure changes. Leave accurate
content untouched when none of those facts changes.

The [README refresh workflow](../workflows/readme-refresh.md) performs the review before a thematic commit.
