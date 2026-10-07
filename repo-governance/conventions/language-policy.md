---
tldr: "Requires English in repository artifacts and clear, simple English when an agent replies in English."
when_to_use: "Use when writing code, documentation, plans, tests, commit messages, or an English agent reply."
---

# Language Policy

## Scope

English is the authored language of this repository. That covers source identifiers, comments, documentation,
specifications and their Gherkin, test names and descriptions, configuration labels, commit messages, plans, and every
rule under `repo-governance/`.

Use the Beaver Nest vocabulary for behaviour-related terminology. Spell `behaviour` and `behavioural`, including source
identifiers, directory names, and Nx targets; do not introduce the American variant. This deliberate British spelling
also matches Elixir's `@behaviour` terminology and keeps cross-repository searches deterministic.

## Writing Style

Use clear, simple, and natural English that is easy for non-native speakers to understand. Apply this to all English
repository writing and to agent replies when the agent uses English.

Keep comments, commit messages, notes, plans, and other documents concise; remove detail that does not help readers.
Explain why code or a choice is needed when the reason is not evident, without retelling visible operations. Keep
behaviour, usage, and verification details that readers need.

Review these judgments in context: a repeated operation, unnecessary detail, or wording that obscures the reason
violates the style. No mechanical check can judge the meaning reliably, so this rule is unenforced by decision.

## Conversation Language

The owner may talk to an agent in Bahasa Indonesia, English, or a mix, and an agent answers in the language it was
addressed in. The English clarity guidance applies when an agent answers in English.

## The Exceptions

Use another language when it is the subject rather than the medium: a term being defined, a quotation reproduced
faithfully, a fixture that must contain non-English input to test what the code does with it, or an external interface
that dictates its own strings. Keep the surrounding explanation in English so a reader who does not know the other
language can still follow what the passage is for.

## Verification

Language meaning is verified in review. `npm run test:repo` additionally rejects the American variant in maintained
surfaces. The failure it prevents is a mixed corpus that cannot be searched, compared, or propagated reliably.
