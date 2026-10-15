"""Tests for algo_20_count_min_sketch.py."""

import pytest

from algo_20_count_min_sketch import CountMinSketch, HeavyHitters


def test_estimate_never_undercounts() -> None:
    sketch = CountMinSketch(width=200, depth=4)
    truth = {f"item-{i}": i % 7 + 1 for i in range(300)}
    for item, count in truth.items():
        sketch.add(item, count)
    assert all(sketch.estimate(item) >= count for item, count in truth.items())


def test_single_item_is_counted_exactly() -> None:
    sketch = CountMinSketch(width=1000, depth=4)
    for _ in range(7):
        sketch.add("x")
    assert sketch.estimate("x") == 7


def test_unseen_item_estimate_is_small() -> None:
    sketch = CountMinSketch(width=1000, depth=4)
    sketch.add("seen", 3)
    assert sketch.estimate("never-seen") == 0


def test_from_error_builds_a_working_sketch() -> None:
    sketch = CountMinSketch.from_error(epsilon=0.01, delta=0.01)
    sketch.add("a", 3)
    assert sketch.estimate("a") >= 3


def test_heavy_hitters_keep_the_most_frequent_items() -> None:
    hitters = HeavyHitters(CountMinSketch(width=2000, depth=5), k=2)
    for _ in range(50):
        hitters.observe("hot")
    for _ in range(20):
        hitters.observe("warm")
    for i in range(30):
        hitters.observe(f"cold-{i}")
    assert [name for name, _ in hitters.top()] == ["hot", "warm"]


def test_invalid_sizes_are_rejected() -> None:
    with pytest.raises(ValueError):
        CountMinSketch(width=0, depth=1)
    with pytest.raises(ValueError):
        HeavyHitters(CountMinSketch(width=10, depth=1), k=0)
