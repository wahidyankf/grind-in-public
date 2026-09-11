---
tldr: "Specifies the six documents every plan folder contains and what each owns."
when_to_use: "Use when scaffolding a plan folder or deciding which file a section belongs in."
---

# Required Documents

Every formal plan — one in `plans/backlog/` or `plans/in-progress/` — contains six documents. Not five, and not five
plus an optional sixth.

| Document            | Answers                                                                                      |
| ------------------- | -------------------------------------------------------------------------------------------- |
| `README.md`         | status, context, scope with affected projects named, approach, and how to navigate the plan  |
| `brd.md`            | the business goal, the roles it serves, the outcomes, the non-goals, the business risks      |
| `prd.md`            | personas, user stories, testable acceptance criteria, product scope, product risks           |
| one technical shape | how it will actually be built — see [Technical Shape](004-technical-shape-and-companions.md) |
| `delivery.md`       | the ordered, granular, execution-grade checklist and its proof                               |
| `learnings.md`      | what was discovered while executing, held until it is routed somewhere durable               |

The technical shape is one of the six, not an addition to them. A plan root therefore holds six documents when the
technical shape is a single file, and five files plus one directory when it is a directory.

`evidence/` is optional and is not a seventh document: it holds command output and artifacts that `delivery.md` links.

## What Each Document Owns

**`README.md`** is the first file the [plan quality gate](../../workflows/plan-quality-gate.md) reads for scope, so name
the affected projects explicitly rather than describing them.

**`brd.md`** carries the business rationale. In a personal repository the "business" is the owner's own goals, so write
real reasoning and label a judgment call as one. Never invent a metric to fill a heading.

**`prd.md`** carries user stories in `As a … I want … So that …` form and acceptance criteria in Gherkin, per the
[specs policy](../../development/specs-policy.md).

**`delivery.md`** is the phased checklist execution reads and the quality gate verifies; see the
[delivery contract](005-delivery-contract.md).

**`learnings.md`** is transient by design; see [Knowledge Capture and Archival](009-knowledge-capture-and-archival.md).

## Why Six

The first three separate concerns that are genuinely different and are routinely conflated: why the work is worth doing,
what the result must do, and how it will be built. Collapsing them produces a document that argues for itself, which is
exactly the document nobody can review.

`delivery.md` exists because a plan that cannot be executed step by step has not finished being a plan.

`learnings.md` is the sixth and the one most often dropped. Execution always discovers things — a wrong assumption, a
tool that does not behave as documented, a rule that turned out to be load-bearing. Without a place to put them, those
discoveries are lost at exactly the moment they are most valuable, and the next plan rediscovers them.

## No Five-Document Rule

`learnings.md` is not optional, and no artifact in this repository may say otherwise. This applies to every place a
plan's contents can be described: governance prose, skills, agent definitions, workflows, templates, validators, and
repository instructions.

This repository previously counted five, treating `learnings.md` as a transient log outside the set. The artifacts did
not change; the count and the status of that file did. A five-document description surviving in a live artifact is a
defect regardless of how old it is — a live rule contradicting another live rule, and an executor following it produces
a plan that fails validation for a reason the text told it was correct.

`plans/done/` is the one exception, and it is not an escape. An archived plan records what happened under the rules of
its day, so its prose is left alone and only its links are repaired when a rule file it points at is renamed. Rewriting
history to agree with a newer rule would make the archive useless for the one thing it is for.

## No Single-File Exception

An owner-requested plan uses all six documents. Do not delete or skip a requested plan because its size appears small;
ask the owner whether to amend or cancel it instead.

The plan quality gate requires all six documents and verifies their distinct reader jobs.
