---
tldr: "Defines the mirrored specification tree, required application C4 models, and optional detail folders."
when_to_use: "Use when adding or organizing specifications under specs/."
---

# Specs Structure

Specs mirror the workspace: `specs/apps/<name>/` and `specs/libs/<name>/`. Each carries the folders its subject actually
needs, and no others:

```text
specs/apps/<name>/
+-- README.md          index for this subject
+-- architecture.md    canonical as-built C4 model for an application
+-- product/           the problem, its users, and scope
+-- system-context/    the boundary with the outside world
+-- containers/        the running parts
+-- components/        the internals of a part
+-- behaviours/          Gherkin acceptance scenarios
```

Every folder in this tree carries its own `README.md` indexing what sits in it, `behaviours/` included, as the
[documentation index policy](../../documentation-index-policy.md) requires; a detail folder added later gets one in the
same change.

`behaviours/` is the only mandatory folder for every subject. Every non-drill application also owns the root
`architecture.md` required by the [architecture specification policy](../architecture-specifications.md). Add a detail
folder only when the subject is complex enough to need it, not in advance.
