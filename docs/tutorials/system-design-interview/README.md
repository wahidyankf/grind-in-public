---
tldr: "An ordered, terminal-first curriculum for designing reliable multi-tenant decision and investigation systems."
when_to_use:
  "Use to prepare for a system-design interview from first principles through timed, production-grounded cases."
---

# System Design Interview

This course teaches a repeatable interview method, not a catalogue of fashionable components. Read the foundations
first, then the datastore and pattern guides, and only then attempt the cases. Every choice must connect a requirement
to a measurable property, a failure mode, and an operating cost.

```text
requirements -> estimates -> API/data -> architecture -> deep dives -> failures -> evolution
      ^                                                               |
      +----------------------- verify traceability -------------------+
```

## Reading order

1. [Foundations](foundations/README.md) builds the interview method, distributed-systems vocabulary, Kubernetes model,
   security, ML boundaries, and migration method.
2. [Datastores](datastores/README.md) compares relational, wide-column, document, cache, log, search, graph, and
   analytical storage through access patterns.
3. [Patterns](patterns/README.md) connects service boundaries, messaging, consistency, tenancy, resilience, and
   deployment patterns to production trade-offs.
4. [Case-study prompts](case-studies/README.md) contains 24 timed exercises with explicit functional and non-functional
   requirements.
5. [Worked solutions](solutions/README.md) provides one defensible answer to each prompt. Attempt a prompt before
   reading its solution.

## Interview loop

Use this clock for a 45-minute exercise. Adapt it explicitly when the interviewer gives a different limit.

```text
minute  0       5        10              22             35        42   45
        | clarify | scale | broad design | two deep dives | failures | recap |
```

At every transition, summarize the decision in one sentence. A strong answer is not the largest diagram; it is a small
architecture whose important claims can be traced back to numbered requirements.

## Grounding rule

For each technology or algorithm, ask four questions:

1. Which production problem does it solve?
2. Which property makes it fit the access pattern or failure model?
3. What operational and correctness costs does it introduce?
4. Under which conditions should it be rejected?

The examples use Python when executable logic clarifies the boundary. They are teaching slices with complete imports,
not a hidden application that must be run from this repository.
