---
tldr: "Runs one observable behaviour increment through an evidenced red-green-refactor cycle."
when_to_use: "Use for every new or changed application or library behaviour and every bug fix."
---

# Red-Green-Refactor

Use this workflow for each behaviour increment required by the [TDD policy](../development/tdd-policy.md).

## Prerequisites

- Read the relevant canonical Gherkin, tests, and adapters before production code.
- Identify the smallest observable behaviour to add or change.
- Identify the narrowest automated Nx target that can demonstrate it.

## Cycle

```text
RED: expected behavioural failure
 |
 v
GREEN: minimum implementation passes
 |
 v
REFACTOR: improve design; remain green
 |
 `----> next behaviour increment
```

1. **RED:** write or change one test that expresses the intended behaviour. Run the relevant Nx target and confirm it
   fails because the behaviour is absent or incorrect. Correct the test or environment when it fails for another reason.
2. **GREEN:** implement the minimum production change needed for that test. Run the same target and confirm the new test
   and its surrounding tests pass.
3. **REFACTOR:** improve names, structure, duplication, and design without adding behaviour. Keep tests green after each
   meaningful refactor.
4. Repeat for the next behaviour increment.

## Verification

After the final cycle, run the affected project's broader Nx targets and applicable repository checks. Record the test
path, target, expected RED reason, observed RED, GREEN, and REFACTOR-green result. A broken harness or environment must
be repaired or diagnosed first and never counts as RED evidence.
