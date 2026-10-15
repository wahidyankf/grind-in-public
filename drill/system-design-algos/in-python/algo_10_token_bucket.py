"""Token bucket rate limiter.

Used in: API gateways and per-client limits (AWS API Gateway, Stripe-style limits, Nginx-style limits). It allows short
bursts up to the bucket size, while the long-run rate stays at the refill rate.

Complexity: allow is O(1).

The clock is passed in, so tests can move time forward without sleeping.
"""

import time
from collections.abc import Callable


class TokenBucket:
    def __init__(
        self,
        capacity: float,
        refill_rate: float,
        clock: Callable[[], float] = time.monotonic,
    ) -> None:
        if capacity <= 0 or refill_rate <= 0:
            raise ValueError("capacity and refill_rate must be positive")
        self._capacity = capacity
        self._refill_rate = refill_rate
        self._clock = clock
        self._tokens = capacity
        self._last = clock()

    @property
    def tokens(self) -> float:
        self._refill()
        return self._tokens

    def allow(self, cost: float = 1.0) -> bool:
        """Take cost tokens if there are enough. Returns False and takes nothing otherwise."""
        self._refill()
        if self._tokens < cost:
            return False
        self._tokens -= cost
        return True

    def _refill(self) -> None:
        now = self._clock()
        elapsed = now - self._last
        self._last = now
        self._tokens = min(self._capacity, self._tokens + elapsed * self._refill_rate)
