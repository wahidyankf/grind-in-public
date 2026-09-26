---
tldr: "Design multilingual entity candidate generation, scoring, thresholding, feedback, and evidence."
when_to_use: "Use to practise normalization, search indexes, edit distance, ranking, and model/rule evaluation."
---

# Case 009: Multilingual Entity Matching

Design a service that compares people and organizations against large reference lists using names in multiple scripts,
aliases, identifiers, dates, countries, and addresses.

### Functional requirements

| ID   | Requirement                                                                                   |
| ---- | --------------------------------------------------------------------------------------------- |
| FR-1 | Normalize input without losing original text, script, field boundaries, or provenance.        |
| FR-2 | Generate a bounded candidate set using exact identifiers and approximate text retrieval.      |
| FR-3 | Score candidates with field-aware features and return ranked matches plus reason codes.       |
| FR-4 | Apply tenant/version-specific thresholds for no-match, review, and match.                     |
| FR-5 | Record reviewer feedback and evaluation labels without immediately corrupting training truth. |

### Non-functional requirements

| ID    | Requirement                                                                                    |
| ----- | ---------------------------------------------------------------------------------------------- |
| NFR-1 | Search 50 million entities within 120 ms p99 at 3,000 queries/s.                               |
| NFR-2 | Candidate retrieval must meet a validated recall floor by language and entity type.            |
| NFR-3 | A result is reproducible from input, list, normalizer, feature, model, and threshold versions. |
| NFR-4 | Bound adversarial text length, token count, candidate count, and scoring work.                 |
| NFR-5 | Protect sensitive query text and segregate tenant-specific reference additions.                |

## Assumptions and exclusions

Human specialists define accepted normalization and transliteration behaviour. The system assists decisions; it does not
claim that string similarity proves identity.

## Interview prompts

1. Separate normalization, blocking/candidate generation, scoring, and policy.
2. Compare n-grams, phonetic keys, edit distance, token similarity, and learned ranking.
3. How are thresholds evaluated with precision/recall and asymmetric costs?
4. How are list index versions rolled out and monitored for language-specific regressions?

Solve before reading [the worked solution](../solutions/009-multilingual-entity-matching.md).
