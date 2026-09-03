---
tldr: "Requires English in everything the repository authors, while leaving conversation unrestricted."
when_to_use: "Use when writing code, documentation, specifications, tests, or commit messages."
---

# Language Policy

## Scope

English is the authored language of this repository. That covers source identifiers, comments, documentation,
specifications and their Gherkin, test names and descriptions, configuration labels, commit messages, plans, and every
rule under `repo-governance/`.

## Conversation Is Not Covered

This policy governs repository artifacts, not people. The owner may talk to an agent in Bahasa Indonesia, English, or a
mix, and an agent answers in the language it was addressed in. Only what lands in a file is bound.

The split matters because the two have different readers. A conversation has one reader who chose the language; a
committed file has every future reader, including the owner months later and every agent that loads it as context.

## The Exceptions

Use another language when it is the subject rather than the medium: a term being defined, a quotation reproduced
faithfully, a fixture that must contain non-English input to test what the code does with it, or an external interface
that dictates its own strings. Keep the surrounding explanation in English so a reader who does not know the other
language can still follow what the passage is for.

## Verification

No check reads for language, so this is verified in review. The failure it prevents is a mixed corpus: a policy written
half in one language and half in another cannot be searched, compared, or propagated reliably, and this repository's
rules are read far more often by search than by browsing.
