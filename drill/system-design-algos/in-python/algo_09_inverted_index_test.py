"""Tests for algo_09_inverted_index.py."""

from algo_09_inverted_index import InvertedIndex, tokenize


def _index() -> InvertedIndex:
    index = InvertedIndex()
    index.add(1, "The quick brown fox")
    index.add(2, "The lazy dog sleeps")
    index.add(3, "A quick dog runs quickly")
    return index


def test_tokenize_lowercases_and_splits_on_non_alphanumeric() -> None:
    assert tokenize("Hello, World-42!") == ["hello", "world", "42"]


def test_single_term_search() -> None:
    assert _index().search("dog") == [2, 3]


def test_and_query_intersects_postings() -> None:
    assert _index().search("quick dog") == [3]


def test_search_is_case_insensitive() -> None:
    assert _index().search("QUICK") == [1, 3]


def test_unknown_term_returns_nothing() -> None:
    assert _index().search("unicorn") == []


def test_empty_query_returns_nothing() -> None:
    assert _index().search("   ") == []


def test_adding_same_document_twice_does_not_duplicate() -> None:
    index = InvertedIndex()
    index.add(7, "cache cache")
    index.add(7, "cache")
    assert index.search("cache") == [7]
