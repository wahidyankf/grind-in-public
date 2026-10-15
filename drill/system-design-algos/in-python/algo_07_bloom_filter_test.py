"""Tests for algo_07_bloom_filter.py."""

import pytest
from algo_07_bloom_filter import BloomFilter


def test_added_items_are_always_found() -> None:
    bloom = BloomFilter.from_capacity(500, 0.01)
    items = [f"item-{i}" for i in range(500)]
    for item in items:
        bloom.add(item)
    assert all(bloom.might_contain(item) for item in items)


def test_empty_filter_contains_nothing() -> None:
    bloom = BloomFilter(1024, 3)
    assert not any(bloom.might_contain(f"x-{i}") for i in range(100))


def test_false_positive_rate_stays_near_target() -> None:
    bloom = BloomFilter.from_capacity(1000, 0.01)
    for i in range(1000):
        bloom.add(f"in-{i}")
    false_positives = sum(1 for i in range(10_000) if bloom.might_contain(f"out-{i}"))
    assert false_positives / 10_000 < 0.03


def test_from_capacity_sizes_grow_with_items() -> None:
    small = BloomFilter.from_capacity(100, 0.01)
    large = BloomFilter.from_capacity(10_000, 0.01)
    assert large.bit_count > small.bit_count


def test_invalid_arguments_are_rejected() -> None:
    with pytest.raises(ValueError):
        BloomFilter.from_capacity(0, 0.01)
    with pytest.raises(ValueError):
        BloomFilter.from_capacity(10, 1.5)
    with pytest.raises(ValueError):
        BloomFilter(0, 1)
