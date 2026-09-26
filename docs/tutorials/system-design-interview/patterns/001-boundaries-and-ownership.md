---
tldr:
  "Chooses module and service boundaries from domain cohesion, data authority, team ownership, and runtime pressure."
when_to_use: "Use for modular-monolith design, service extraction, or reviewing an over-fragmented architecture."
---

# Boundaries and Ownership

Use one vocabulary and one source of truth inside a bounded context. A service boundary should enclose a cohesive
capability, own its data writes, and be operated by one accountable team.

## Boundary test

```text
                  same deploy?   same data?   same owner?   same scaling?
policy authoring       maybe          yes          yes           no
online evaluation      maybe         read         yes           yes
case workflow           no            no          other          no
```

Differences are signals, not automatic splits. Extract only when the benefit exceeds network, deployment, observability,
and consistency costs.

## API composition

An edge or backend-for-frontend may compose independently owned queries for one client view. Bound fan-out and set
per-call deadlines:

```text
client -> composer --+-> profile
                     +-> recent decisions
                     `-> open cases
```

Reject this when a view needs a strongly consistent join; build a projection owned by the view instead of calling many
services inside a transaction.

## Orchestration and choreography

An orchestrator makes workflow state, timeout, compensation, and operator visibility explicit. Event choreography
reduces central coordination but can hide the end-to-end process across many consumers.

```text
orchestrated: order workflow -> reserve -> screen -> finalize
choreographed: order-created -> listeners react independently
```

Use orchestration for long-running business processes with strict states. Use choreography for independent reactions.
Reject a distributed saga when one local database transaction can still satisfy the requirement.

## Anti-corruption layer

During migration, translate a legacy model at the boundary rather than leaking it into the new service. The adapter is
temporary but production-critical: version it, test it with recorded contract fixtures, and measure unmapped fields.

## Ownership checklist

Every service needs an owning team, on-call route, SLO, deployment pipeline, data classification, capacity plan,
dependency contract, and retirement path. A repository directory is not ownership.
