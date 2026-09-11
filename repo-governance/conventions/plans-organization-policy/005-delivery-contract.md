---
tldr: "Fixes what a delivery.md checkbox must carry, who owns it, and what the file declares up front."
when_to_use: "Use when writing or reviewing any checkbox in a plan's delivery.md."
---

# Delivery Contract

`delivery.md` is the executable part of a plan. Everything else describes intent; this is the part someone works
through, read literally by an agent that cannot reconstruct what a checkbox left out.

The dated log it opens with is a separate artifact; the [execution record](012-execution-record.md) owns it.

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

```markdown
- [ ] [AI] Edit `apps/example/internal/detect.go`: preserve one rule-change path after normalization. Verify with
      `rtk ./hippo run --class ephemeral --disk-path . -- npm exec -- nx run -p example -t test:quick` — the suite
      exits 0.
```

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
human; they call for care, evidence, and an authorization recorded once. A plan that marks items `[HUMAN]` because they
matter has stopped being executable and become a request for supervision. Git-mechanical steps — committing, pushing,
moving a plan folder — are `[AI]` unless a specific reason says otherwise.

Where the outcome is uncertain, the item becomes a bounded checkpoint with a predeclared fallback — what is tried, how
many attempts, what happens at the ceiling — decided when the plan is written.

Recovery work names its trigger and stays dormant until triggered, read against its wording rather than the presence of
a failure. At final reconciliation a dormant item receives a dated, evidence-backed `Not triggered` disposition, never a
false completion mark.

## Structure

`delivery.md` declares, before its phases:

- the **execution checkout** — which working copy and branch the work happens on;
- the **delivery units** — the transaction boundaries, each with one owner, one testable outcome, one rollback; and
- **pause safety** — what is recorded at a pause so work resumes without re-deriving it.

Phases follow in dependency order and each ends at a gate; see [Phases and Gates](011-phases-and-gates.md). Archival
items live in their own section after every substantive phase, because a plan can be finished without being filed.

## Cold-Executor Resumability

The test for `delivery.md` is whether an executor with no memory of the plan can open it and know what to do next. That
means recording results, not only ticks: a tick says an action happened, not what it produced.
