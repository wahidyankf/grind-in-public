---
tldr: "Defines the one-document bug-fix plan, what makes each section useful, and the plan rules it is exempt from."
when_to_use:
  "Use when fixing a defect through a plan in the repository that owns the defective code, or when reviewing, executing,
  or archiving such a plan."
---

# Bug-Fix Plan

A bug-fix plan is the compressed formal plan for one defect. The defect already answers why the work is worth doing and
what the result must do — the behaviour should not happen — so separate business and product documents would only
restate the report. What a cold executor still needs is the evidence, the cause, and the change.

It is reserved for a defect that blocks the work in hand with no workaround. Any other defect is filed as an
[idea brief](013-two-pager-template.md) and waits for grooming, per
[Upstream Tool Defects](../../development/upstream-tool-defects.md) where that standard applies.

## Shape

One document, `plans/in-progress/fix-<slug>/README.md`, with these sections in this order:

| Section         | Contains                                                                                      |
| --------------- | --------------------------------------------------------------------------------------------- |
| Bug Report      | every field [Bug Reports](../bug-reports.md) requires, the version and commit included        |
| Duplicate Check | each search of open issues, pull requests, plans, and idea briefs, with its query and result  |
| Root Cause      | the mechanism that produces the symptom, with evidence a reviewer can check                   |
| Solution        | the change, why it removes the cause, and every reference consulted, cited by URL             |
| Delivery        | an Execution Record, then labelled items: failing regression test, fix, verification, release |
| Learnings       | what executing it taught, routed before archival                                              |

The slug starts with `fix-` and names the symptom, not the suspected cause: `fix-gate-skips-renamed-files`.

## What Makes Each Section Useful

**Root Cause** separates the symptom from the mechanism. It names where the behaviour comes from — a file and line, a
commit, a condition — and shows why: a minimal reproduction, a trace, or the failing assertion. "The parser is buggy" is
a symptom restated; "an empty path reaches the matcher at this line and matches nothing" is a cause. A cause that is
still a hypothesis says so and states the test that would confirm it.

**Solution** says why the change removes the cause rather than the symptom, and which conditions it covers beyond the
reported one. Research before choosing: the tool's own documentation, the specification or upstream behaviour it
implements, and prior reports of the same failure, found by searching the web as well as the repository. Cite each with
its URL, primary sources first. An approach copied from a source names the source.

**Delivery** starts with a regression test that fails for the reported reason, so the fix is proven rather than
asserted. Each item carries its executor label, as the [Delivery Contract](005-delivery-contract.md) requires, and cites
the Bug Report's expected behaviour where another plan cites an acceptance identifier.

## What It Is Exempt From

- **[Required Documents](003-required-documents.md):** it is the one plan with one document; its sections take the six
  documents' roles.
- **[Lifecycle and Folders](001-lifecycle-and-folders.md):** it starts in `in-progress/`, with no idea or backlog stage,
  and moves to `done/` like any plan.
- **The [decision gates](006-workflows-and-skills.md#decision-gates):** a defect has one correct behaviour, so there is
  nothing to interview about.
- **[Authorization](../plans-organization-policy.md#when-a-plan-is-allowed):** it needs no separate request when written
  under an adopted [Upstream Tool Defects](../../development/upstream-tool-defects.md) standard, or when the owner asks.
  That standing request also directs its plan quality gate, its execution, its release, its plan execution check, and
  its archival: once the plan lands, run the gate on it, execute on a passing verdict, release the fix once its tests
  pass, then run the [plan execution check](../../workflows/plan/plan-execution-check.md) and archive on a verdict that
  permits it, without a further prompt.

Every other plan rule holds: slug rules, one root, executor labels, no time estimates, the
[execution record](012-execution-record.md), knowledge capture, and archival.

## When It Stops Being One

A fix that needs a design decision, changes an interface its consumers rely on, or needs more than one delivery unit is
not a bug-fix plan. Expand it to the six documents in place, keeping the slug, and run the decision gates before
continuing.

## Landing the Plan First

The plan lands on the owning repository's trunk alone, through that repository's route, before any fix is committed. A
parallel finder's duplicate check then sees it, and the fix follows in its own change.
