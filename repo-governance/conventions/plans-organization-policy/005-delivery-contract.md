---
tldr: "Fixes what a delivery.md checkbox must carry, who owns it, and what the file declares up front."
when_to_use: "Use when writing or reviewing any checkbox in a plan's delivery.md."
---

# Delivery Contract

`delivery.md` is the executable part of a plan, read literally by an agent that cannot reconstruct omitted details.

The dated log it opens with, and how results are recorded so an executor with no memory of the plan can resume, belong
to the [execution record](012-execution-record.md).

## One Checkbox, One Action

A checklist item represents exactly one independently verifiable action, and it names:

- the **paths** it touches — exactly, when known; otherwise the parent directory, the naming pattern, and a sibling to
  imitate;
- the **command** that performs or proves it, verbatim, where one exists;
- its **executor** label;
- the **proof** that it is done — an observable outcome, never a bare "implement" or "configure"; and
- the **acceptance criteria** it satisfies, as `[AC-…]` identifiers traced to `prd.md`.

A criterion satisfied by finding nothing pairs that with a check proving the command looked at something real, and one
reading a tool's output names a command no shell wrapper rewrites. A behaviour cycle names one canonical Gherkin path
and scenario, without duplicating its body, and carries separate RED, GREEN, and REFACTOR checkboxes with the evidence
the [TDD policy](../../development/tdd-policy.md) requires.

"Independently verifiable" is the test. If finishing an item leaves no observable difference, it is a thought rather
than an item. If proving it needs the next item finished first, the split is wrong. An item that takes a long session
hides several, and the first failure inside it has no checkbox to fail against.

`- [ ] Add caching` fails that test. This passes:

````markdown
- [ ] [AI] Edit `apps/example/internal/detect.go`: preserve one rule-change path after normalization. Verify with
  ```sh
  rtk ./hippo run --class ephemeral --resource-tier standard --disk-path . -- npm exec -- \
    nx run -p example -t test:quick
  ```
  — the suite exits 0.
````

## Deterministic and Executable

**Deterministic proof.** Every checklist item's proof is a deterministic check: a command with its expected exit status
or output, or a file whose presence or content a command tests. "Holds on read", "looks right", and "satisfy review" are
not proofs.

**Execution-tier executable.** An agent at the `execution` tier can carry out every item from the plan's own documents,
without consulting other context or making a policy choice. An item labelled `[HUMAN]` for one of the four reasons below
is the only exception.

A plan item is a goal an executor works toward, and only a check that passes or fails the same way for everyone can say
the goal is met.

## Executor Labels and AI-First Ownership

Every checkbox states who can execute it, tagged `[AI]` or `[HUMAN]`. The default is `[AI]`, and an untagged checkbox is
a defect rather than an `[AI]` one. A `delivery.md` opens with a one-line legend naming both, so a reader meets them
before the first checkbox.

`[HUMAN]` is permitted for exactly four reasons:

1. a credential or access the executor does not have;
2. a physical action;
3. an external authority — someone else's approval or action; or
4. a decision genuinely unavailable, because the information it needs does not exist yet.

Two labels, no third. Work an agent prepares and the owner approves splits across them rather than sitting between: the
preparation is `[AI]` with the authorization recorded once, and only the action the owner must perform is `[HUMAN]`. A
step the owner does by hand to learn it is `[HUMAN]` by choice.

**Significance is never a reason.** Important, irreversible, expensive, or public-facing does not transfer an item to a
human; they call for care, evidence, and an authorization recorded once. Git-mechanical steps — committing, pushing,
moving a plan folder — are `[AI]` unless a specific reason says otherwise.

Where the outcome is uncertain, the item becomes a bounded checkpoint with a predeclared fallback (what is tried, how
many attempts, what happens at the ceiling) decided when the plan is written.

Recovery work follows [Resolution Is Not a Tick](008-evidence-and-quality.md#resolution-is-not-a-tick).

## Structure

`delivery.md` declares, before its phases:

- the **execution checkout** — which working copy and branch the work happens on;
- the **delivery units** — the transaction boundaries, each with one owner, one testable outcome, one rollback; and
- **pause safety** — what is recorded at a pause so work resumes without re-deriving it.

Phases, their gates, and the archival section follow [Phases and Gates](011-phases-and-gates.md).
