---
tldr:
  "Leads architecture through explicit constraints, decision ownership, review, standards, and reversible evolution."
when_to_use: "Use for architecture, technical strategy, migration, and senior-engineer partnership questions."
---

# Technical Leadership and Architecture

Technical leadership is not taking every design from the team. Set outcome and constraints, assign a decision owner,
ensure relevant reviewers, require evidence, and make consequences visible.

## Decision ladder

```text
reversible/local ------------------------------ irreversible/systemic
engineer decides   tech lead reviews   RFC/ADR + cross-team review   executive risk
```

Push decisions to the lowest level with the context and capability to own them. Escalate based on blast radius,
irreversibility, security, data migration, cross-team coupling, and operating cost—not seniority preference.

## Design review questions

1. Which functional and non-functional requirements are prioritized?
2. Which numbers or experiments establish expected scale?
3. What is the smallest architecture that meets them?
4. Which invariants cross boundaries and who owns the data?
5. How does overload, partial failure, rollback, and recovery work?
6. How will on-call observe and operate it?
7. What evidence permits the next migration state?

Require alternatives and rejection reasons. A review that only adds boxes does not improve the decision.

## Working with senior engineers

Agree on decision rights. The staff engineer may own technical coherence and technical mentoring; the manager owns team
health, staffing, prioritization, performance, and accountability for delivery/operation. Both share context and should
disagree in private with concrete evidence, then communicate one decision.

## Technical debt portfolio

Classify debt by current consequence:

| Class                | Evidence and response                                               |
| -------------------- | ------------------------------------------------------------------- |
| Reliability/security | active risk; prioritize with incident/threat evidence               |
| Delivery friction    | repeated lead-time/rework cost; automate or simplify the bottleneck |
| Architecture limit   | blocks a named requirement; evolve at the seam                      |
| Cosmetic preference  | no demonstrated outcome; do not displace higher-value work          |

Reserve capacity is useful only with transparent selection and outcomes. A permanent percentage without a debt portfolio
can become unaccountable work.

## Migration leadership

For modular-monolith extraction, require baseline metrics, a justified seam, one data writer, shadow/canary comparison,
rollback, operational ownership, and decommissioning. Celebrate retired complexity, not new service count.

## Checkpoint

Prepare one story where you rejected a fashionable architecture and one where new evidence changed your position. Both
should show how the team learned rather than how you won an argument.
