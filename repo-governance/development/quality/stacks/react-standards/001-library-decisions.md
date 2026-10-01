---
tldr: >-
  Records the React library choices an adopter makes once per application, a store, a query cache, and form handling,
  with each option's gains and costs, and illustrative tools for the React rules.
when_to_use: >-
  Use when choosing a store, query, or form library for a React application, or when looking for tools that apply the
  React lint, test, sanitisation, and accessibility rules.
---

# Library Decisions

## Adopter Decisions

Each decision's options, with gains and costs:

- **store**
  - a store library. Gains: selective subscriptions. Costs: one more dependency and idiom.
  - context and reducers. Gains: nothing beyond React. Costs: broad re-renders as shared state grows.
- **query**
  - a full-featured cache library. Gains: mutations and optimistic updates built in. Costs: a larger surface to
    configure.
  - a lighter revalidating library. Gains: a small API. Costs: mutations assembled by hand.
- **forms**
  - a form library with schema validation. Gains: field and error state handled once. Costs: every form couples to it.
  - controlled components with a schema. Gains: plain React. Costs: field and error state rewritten per form.

Record each choice once per application, selecting libraries as
[Dependency Selection](../../../dependency-selection-policy.md) requires. Either store option satisfies the fourth state
home in [React Standards](../react-standards.md).

## Illustrative Example

As an illustration only: `eslint-plugin-react-hooks` and `eslint-plugin-jsx-a11y` carry the lint rules; Testing Library
with `user-event` gives role queries and user events; TanStack Query or SWR is a query cache; Zustand is a store;
DOMPurify sanitises raw HTML; and `vitest-axe` runs the accessibility check.
