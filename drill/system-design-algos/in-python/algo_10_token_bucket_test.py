"""Tests for algo_10_token_bucket.py."""

import pytest
from algo_10_token_bucket import TokenBucket


class FakeClock:
    def __init__(self, start: float = 0.0) -> None:
        self.now = start

    def __call__(self) -> float:
        return self.now

    def advance(self, seconds: float) -> None:
        self.now += seconds


def test_burst_up_to_capacity_then_denied() -> None:
    clock = FakeClock()
    bucket = TokenBucket(capacity=3, refill_rate=1, clock=clock)
    assert [bucket.allow() for _ in range(4)] == [True, True, True, False]


def test_refill_allows_requests_again() -> None:
    clock = FakeClock()
    bucket = TokenBucket(capacity=2, refill_rate=1, clock=clock)
    bucket.allow()
    bucket.allow()
    assert bucket.allow() is False
    clock.advance(1.0)
    assert bucket.allow() is True
    assert bucket.allow() is False


def test_refill_never_exceeds_capacity() -> None:
    clock = FakeClock()
    bucket = TokenBucket(capacity=5, refill_rate=10, clock=clock)
    clock.advance(100.0)
    assert bucket.tokens == 5


def test_cost_takes_several_tokens_at_once() -> None:
    clock = FakeClock()
    bucket = TokenBucket(capacity=10, refill_rate=1, clock=clock)
    assert bucket.allow(cost=7) is True
    assert bucket.allow(cost=7) is False
    assert bucket.tokens == 3


def test_rejects_non_positive_settings() -> None:
    with pytest.raises(ValueError):
        TokenBucket(capacity=0, refill_rate=1)
    with pytest.raises(ValueError):
        TokenBucket(capacity=1, refill_rate=0)
