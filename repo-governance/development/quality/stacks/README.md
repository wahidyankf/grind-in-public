---
tldr: "Indexes the adopted stack standards and the repository adapter that records their local decisions."
when_to_use:
  "Use when a project here builds with a stack whose standard was adopted, or when deciding whether a rule belongs to
  one stack or to every project."
---

# Stack Standards

Each standard governs one stack: it records the normative choices that stack's projects enforce, and links the
language-neutral standards instead of restating them. The stack's skill defers to it. Placement, inheritance, and the
inventory that selects them follow [Stack Packs](../../../conventions/structure/stack-packs.md), and the
[repository adapter](repository-adapter.md) records this repository's decisions and deviations.

## Directory Map

- [Repository Adapter](repository-adapter.md) — the adopted packs, each open decision's local choice, the stronger local
  rules, and a link to every project README
- [Go Standards](golang-standards.md) — gofmt, vet, and linter gates, every error handled or returned, consumer-declared
  interfaces, owned goroutines, and statement coverage
- [JavaScript Standards](javascript-standards.md) — JSDoc type checking under strict options, runtime validation at
  external input, handled promises, and type check, lint, and format gates
- [Next.js Standards](nextjs-standards.md) — App Router with static rendering and server components by default, declared
  caching, validated server actions, and loading and error states
- [Next.js Standards Modules](nextjs-standards/README.md) — the version line and hosting decisions, and framework
  markers
- [Nx Standards](nx-standards.md) — over each project's packs: tagged graph boundaries, affected runs from a correct
  base, real inferred targets, and complete cache inputs
- [Python Standards](python-standards.md) — annotated signatures checked by Pyright in strict mode, formatter and linter
  gates, validated boundaries, narrow exceptions, and branch coverage
- [React Standards](react-standards.md) — typed function components, hooks and accessibility lint, state in its
  narrowest home, server data in a query cache, and client security
- [React Standards Modules](react-standards/README.md) — store, query, and form library decisions, and example tools
- [Shell Standards](shell-standards.md) — beyond Shell Scripts: a supported dialect, analyser and formatter gates,
  quoted and validated input, and behaviour tests without coverage
- [TypeScript Standards](typescript-standards.md) — compiler floor, strict compiler options, no any or unchecked
  assertions, Result values for expected failures, and the type check, lint, and format gates
