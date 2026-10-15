"""Tests for algo_05_lru_cache.py."""

import pytest
from algo_05_lru_cache import LRUCache


def test_get_returns_stored_value() -> None:
    cache: LRUCache[str, int] = LRUCache(2)
    cache.put("a", 1)
    assert cache.get("a") == 1
    assert cache.get("missing") is None


def test_evicts_least_recently_used() -> None:
    cache: LRUCache[str, int] = LRUCache(2)
    cache.put("a", 1)
    cache.put("b", 2)
    cache.put("c", 3)
    assert cache.get("a") is None
    assert cache.get("b") == 2
    assert cache.get("c") == 3
    assert len(cache) == 2


def test_get_refreshes_recency() -> None:
    cache: LRUCache[str, int] = LRUCache(2)
    cache.put("a", 1)
    cache.put("b", 2)
    assert cache.get("a") == 1
    cache.put("c", 3)
    assert cache.get("b") is None
    assert cache.get("a") == 1


def test_put_on_existing_key_updates_and_refreshes() -> None:
    cache: LRUCache[str, int] = LRUCache(2)
    cache.put("a", 1)
    cache.put("b", 2)
    cache.put("a", 10)
    cache.put("c", 3)
    assert cache.get("a") == 10
    assert cache.get("b") is None
    assert len(cache) == 2


def test_capacity_one_keeps_only_newest() -> None:
    cache: LRUCache[int, int] = LRUCache(1)
    cache.put(1, 1)
    cache.put(2, 2)
    assert cache.get(1) is None
    assert cache.get(2) == 2


def test_rejects_zero_capacity() -> None:
    with pytest.raises(ValueError):
        LRUCache[int, int](0)
