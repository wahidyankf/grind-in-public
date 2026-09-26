---
tldr:
  "A two-stage matcher preserves original text, retrieves high-recall candidates, then applies bounded field-aware
  scoring."
when_to_use: "Use after attempting Case 009 to compare normalization, retrieval, scoring, and threshold governance."
---

# Solution 009: Multilingual Entity Matching

## Requirement traceability

| Requirements       | Design response                                                                                        |
| ------------------ | ------------------------------------------------------------------------------------------------------ |
| FR-1, NFR-3        | Versioned language-aware normalizer emits derived forms beside immutable original fields.              |
| FR-2, NFR-1, NFR-2 | Exact-id maps and script/language-specific search indexes retrieve a capped high-recall candidate set. |
| FR-3               | Field-aware scorer combines name, alias, date, country, address, and identifier features with reasons. |
| FR-4               | Tenant policy maps score/evidence to no-match, review, or match with versioned thresholds.             |
| FR-5, NFR-5        | Governed feedback store separates raw reviewer action from validated label and protects query text.    |
| NFR-4              | Input, token, variant, candidate, edit-distance, and total-compute limits.                             |

## Pipeline

```text
query -> validate -> normalize/transliterate variants
                      |
          +-----------+------------+
          v                        v
      exact ids             text candidate index
          +-----------+------------+
                      v
              deduplicate candidates -> feature scoring -> threshold policy
```

Normalization performs Unicode normalization, whitespace/punctuation policy, script detection, and configured name
ordering without deleting originals. Transliteration is a derived retrieval aid, not an identity truth.

## Candidate and scoring design

Exact identifiers can short-circuit only when source quality and policy permit. Text retrieval uses character n-grams
for spelling variation and token indexes for reordered multi-token names. Query several narrow indexes in parallel,
union candidate ids with a hash set, and cap per strategy so one noisy query cannot crowd out exact candidates.

Complete normalized edit similarity for a bounded pair:

```python
def normalized_edit_similarity(left: str, right: str) -> float:
    if left == right:
        return 1.0
    if not left or not right:
        return 0.0

    previous = list(range(len(right) + 1))
    for left_index, left_character in enumerate(left, start=1):
        current = [left_index]
        for right_index, right_character in enumerate(right, start=1):
            substitution = previous[right_index - 1] + (left_character != right_character)
            deletion = previous[right_index] + 1
            insertion = current[right_index - 1] + 1
            current.append(min(substitution, deletion, insertion))
        previous = current

    distance = previous[-1]
    return 1.0 - distance / max(len(left), len(right))


print(round(normalized_edit_similarity("maria", "marie"), 2))  # 0.8
```

This uses `O(len(right))` memory but still costs `O(n*m)` time; production caps lengths and runs it only after
retrieval. Feature weights or a learned ranker are evaluated by language/entity cohort. Reason codes cite matched fields
and transformations without claiming causality.

## Release and monitoring

Build immutable list indexes, verify counts and recall suites, warm serving nodes, then switch a version pointer. Shadow
new normalizer/retriever/scorer versions and measure candidate recall, precision/recall at policy thresholds, latency,
candidate counts, and outcome rates by language/script. Roll back pointers independently.

## Alternatives rejected

All-pairs comparison is quadratic and cannot meet scale. A single Latin transliteration index erases script-specific
signal. One global threshold hides different field availability, costs, and language performance.
