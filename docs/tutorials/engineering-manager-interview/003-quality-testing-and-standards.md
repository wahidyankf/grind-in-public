---
tldr: "Builds quality through risk-based tests, reviewable changes, standards, feedback loops, and ownership."
when_to_use:
  "Use for questions about code quality, testing strategy, release confidence, and raising the engineering bar."
---

# Quality, Testing, and Engineering Standards

Quality is fitness for purpose across correctness, security, reliability, operability, maintainability, and delivery. A
manager builds a system in which defects are prevented or found at the cheapest useful boundary.

## Quality loop

```text
requirements -> design review -> small change -> automated evidence -> progressive release
     ^                                                          |
     +------------ production signals + retrospective ----------+
```

## Risk-based test portfolio

- pure unit tests for domain decisions and edge cases;
- contract tests for service/event/schema compatibility;
- integration tests for database, filesystem, queue, and model boundaries;
- end-to-end tests for a small set of critical user outcomes;
- load/resilience tests for capacity and declared failure modes;
- canary and observability for risks that appear only in production conditions.

More end-to-end tests are not automatically safer. They are slower, less diagnostic, and often brittle. Push most
behaviour to the narrowest boundary that can prove it, then keep a thin critical-path suite.

## Standards adoption

Start from recurring defects or friction, co-author a concrete standard, provide an example/tooling, trial it, and
measure outcome. A written standard without review, automation, education, and ownership decays.

```text
incident pattern -> rule -> example/tool -> adoption -> audit -> outcome -> revise
```

Examples: typed public interfaces, structured error taxonomy, idempotent consumers, migration expand/contract, bounded
timeouts, or required SLO/runbook before launch.

## Code and design review

Keep changes small, state intent and risk, and separate blocking correctness concerns from suggestions. Reviewers should
cite an invariant or standard. Track review age and rework patterns without optimizing for superficial speed.

## Quality disagreement

When delivery pressure conflicts with testing, offer explicit options: reduce scope, stage exposure, add a temporary
guard with owner/expiry, or accept a quantified risk at the appropriate authority. Do not silently lower the bar or use
"quality" to block all change.

## Metrics

Balance leading and outcome signals: change failure, escaped severity, rollback, flaky tests, review age, lead time,
SLO, vulnerability age, and recurrence. Test count and coverage are supporting signals, not goals.
