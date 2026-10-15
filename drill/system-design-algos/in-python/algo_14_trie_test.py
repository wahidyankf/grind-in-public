"""Tests for algo_14_trie.py."""

from algo_14_trie import Trie


def _sample() -> Trie:
    trie = Trie()
    trie.insert("car", 5)
    trie.insert("cart", 2)
    trie.insert("care", 5)
    trie.insert("cat", 1)
    return trie


def test_search_only_matches_whole_words() -> None:
    trie = _sample()
    assert trie.search("car") is True
    assert trie.search("ca") is False


def test_starts_with_matches_prefixes() -> None:
    trie = _sample()
    assert trie.starts_with("ca") is True
    assert trie.starts_with("dog") is False


def test_autocomplete_orders_by_frequency_then_word() -> None:
    assert _sample().autocomplete("ca") == ["car", "care", "cart", "cat"]


def test_autocomplete_respects_limit() -> None:
    assert _sample().autocomplete("ca", limit=2) == ["car", "care"]


def test_autocomplete_includes_the_prefix_itself() -> None:
    assert _sample().autocomplete("car", limit=10) == ["car", "care", "cart"]


def test_autocomplete_unknown_prefix_is_empty() -> None:
    assert _sample().autocomplete("dog") == []


def test_empty_prefix_lists_every_word() -> None:
    assert len(_sample().autocomplete("", limit=10)) == 4
