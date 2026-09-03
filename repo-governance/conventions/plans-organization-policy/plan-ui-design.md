---
tldr:
  "Makes material UI plans compare accessible directions and prove the selected experience on every affected device."
when_to_use: "Use when a formal plan creates or materially changes user interface behavior."
---

# Plan UI Design

A material UI plan creates `tech-docs/ui-design.md` and `tech-docs/assets/README.md`, both mapped from
`tech-docs/README.md`. The design document states the user's job, real product copy, states, three alternatives,
selected direction, rationale, trade-offs, reusable components, and keyboard, focus, error, empty, loading,
reduced-motion, responsive, and accessibility behavior.

Create three distinct lo-fi alternatives for desktop, tablet, and mobile: nine assets total. Compare usability,
accessibility, implementation cost, and product fit, then name one selected alternative. Create selected-direction hi-fi
assets for desktop, tablet, and mobile: three assets total. The plan README previews at least one selected hi-fi asset
and links the full comparison.

Store assets under `tech-docs/assets/` as `ui-<option>-<fidelity>-<device>.svg`, where fidelity is `lofi` or `hifi` and
device is `desktop`, `tablet`, or `mobile`. SVG is the default: each has unique `<title>` and `<desc>`, and every
Markdown embed has useful alt text. Use raster only when bitmap fidelity is material. Do not use color alone, include
real accounts or private data, or use Mermaid; architectural diagrams remain ASCII.

`delivery.md` traces exploration, selection, implementation, accessibility checks, and affected-device manual proof to
the relevant `[AC-…]` labels. `tech-docs/file-impact.md` names every implementation, test, specification, and asset
path.
