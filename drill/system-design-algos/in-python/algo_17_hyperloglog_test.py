"""Tests for algo_17_hyperloglog.py."""

import pytest

from algo_17_hyperloglog import HyperLogLog


def test_empty_sketch_estimates_zero() -> None:
    assert HyperLogLog().estimate() == 0


def test_small_count_is_close() -> None:
    sketch = HyperLogLog(precision=12)
    for i in range(100):
        sketch.add(f"user-{i}")
    assert abs(sketch.estimate() - 100) <= 5


def test_large_count_is_within_five_percent() -> None:
    sketch = HyperLogLog(precision=12)
    for i in range(10_000):
        sketch.add(f"user-{i}")
    assert abs(sketch.estimate() - 10_000) / 10_000 < 0.05


def test_duplicates_do_not_change_the_estimate() -> None:
    once = HyperLogLog(precision=10)
    twice = HyperLogLog(precision=10)
    for i in range(1000):
        once.add(f"item-{i}")
        twice.add(f"item-{i}")
        twice.add(f"item-{i}")
    assert once.estimate() == twice.estimate()


def test_merge_counts_the_union() -> None:
    left = HyperLogLog(precision=12)
    right = HyperLogLog(precision=12)
    for i in range(0, 5000):
        left.add(f"item-{i}")
    for i in range(2500, 7500):
        right.add(f"item-{i}")
    left.merge(right)
    assert abs(left.estimate() - 7500) / 7500 < 0.05


def test_invalid_precision_is_rejected() -> None:
    with pytest.raises(ValueError):
        HyperLogLog(precision=3)


def test_merge_rejects_different_precision() -> None:
    with pytest.raises(ValueError):
        HyperLogLog(precision=10).merge(HyperLogLog(precision=11))
