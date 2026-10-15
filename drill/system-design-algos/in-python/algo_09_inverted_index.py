"""Inverted index.

Used in: full-text search engines (Lucene, Elasticsearch, OpenSearch) and any "find documents containing these words"
feature.

Complexity: add is O(t log p) for t tokens and posting lists of length p. An AND query with terms of length p1..pn
costs about O(sum of the lengths) with the two-pointer merge.

Each term maps to a sorted list of document ids, called a posting list.
"""

import re
from bisect import insort


def tokenize(text: str) -> list[str]:
    return re.findall(r"[a-z0-9]+", text.lower())


def _intersect(left: list[int], right: list[int]) -> list[int]:
    result: list[int] = []
    i = 0
    j = 0
    while i < len(left) and j < len(right):
        if left[i] == right[j]:
            result.append(left[i])
            i += 1
            j += 1
        elif left[i] < right[j]:
            i += 1
        else:
            j += 1
    return result


class InvertedIndex:
    def __init__(self) -> None:
        self._postings: dict[str, list[int]] = {}

    def add(self, doc_id: int, text: str) -> None:
        for term in set(tokenize(text)):
            posting = self._postings.setdefault(term, [])
            if doc_id not in posting:
                insort(posting, doc_id)

    def search(self, query: str) -> list[int]:
        """Return the ids of documents that contain every word in the query."""
        terms = tokenize(query)
        if not terms:
            return []
        result = list(self._postings.get(terms[0], []))
        for term in terms[1:]:
            result = _intersect(result, self._postings.get(term, []))
        return result
