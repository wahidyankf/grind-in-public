"""Tests for algo_19_leaky_bucket.py."""

import pytest

from algo_19_leaky_bucket import LeakyBucketQueue


class FakeClock:
    def __init__(self, start: float = 0.0) -> None:
        self.now = start

    def __call__(self) -> float:
        return self.now

    def advance(self, seconds: float) -> None:
        self.now += seconds


def test_full_queue_rejects_new_items() -> None:
    clock = FakeClock()
    bucket: LeakyBucketQueue[str] = LeakyBucketQueue(capacity=3, leak_rate=2, clock=clock)
    assert [bucket.offer(x) for x in ["a", "b", "c", "d"]] == [True, True, True, False]


def test_drain_releases_items_at_the_leak_rate_in_order() -> None:
    clock = FakeClock()
    bucket: LeakyBucketQueue[str] = LeakyBucketQueue(capacity=3, leak_rate=2, clock=clock)
    for item in ["a", "b", "c"]:
        bucket.offer(item)
    clock.advance(1.0)
    assert bucket.drain() == ["a", "b"]
    assert len(bucket) == 1
    clock.advance(0.5)
    assert bucket.drain() == ["c"]


def test_idle_bucket_does_not_save_up_a_burst() -> None:
    clock = FakeClock()
    bucket: LeakyBucketQueue[str] = LeakyBucketQueue(capacity=5, leak_rate=1, clock=clock)
    clock.advance(100.0)
    for item in ["x1", "x2", "x3"]:
        bucket.offer(item)
    assert bucket.drain() == []
    clock.advance(2.0)
    assert bucket.drain() == ["x1", "x2"]


def test_invalid_settings_are_rejected() -> None:
    with pytest.raises(ValueError):
        LeakyBucketQueue[int](capacity=0, leak_rate=1)
    with pytest.raises(ValueError):
        LeakyBucketQueue[int](capacity=1, leak_rate=0)
